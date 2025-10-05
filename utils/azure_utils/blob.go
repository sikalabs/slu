package azure_utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func AzureBlobClient(accountName, accountKey string) (*azblob.Client, error) {
	if accountName == "" || accountKey == "" {
		return nil, errors.New("missing Azure accountName/accountKey")
	}
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
	return azblob.NewClientWithSharedKeyCredential(url, cred, nil)
}

func FindLatestBlob(ctx context.Context, client *azblob.Client, container, prefix, suffix string) (string, int64, error) {
	pager := client.NewListBlobsFlatPager(container, &azblob.ListBlobsFlatOptions{
		Prefix: to.Ptr(prefix),
	})

	var (
		bestName    string
		bestSize    int64
		bestTime    time.Time
		suffixLower = strings.ToLower(suffix)
	)

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return "", 0, err
		}

		for _, b := range resp.Segment.BlobItems {
			if b.Name == nil || b.Properties.LastModified == nil {
				continue
			}
			name := *b.Name
			mod := *b.Properties.LastModified

			if suffixLower != "" && !strings.HasSuffix(strings.ToLower(name), suffixLower) {
				continue
			}
			if !strings.HasSuffix(strings.ToLower(name), ".gz") {
				continue
			}

			if mod.After(bestTime) {
				bestTime = mod
				bestName = name
				if b.Properties.ContentLength != nil {
					bestSize = *b.Properties.ContentLength
				}
			}
		}
	}

	if bestName == "" {
		return "", 0, fmt.Errorf("no matching blobs found (prefix=%q suffix=%q)", prefix, suffix)
	}
	return bestName, bestSize, nil
}

func OpenBlobStream(ctx context.Context, client *azblob.Client, container, blobName string) (io.Reader, func(), error) {
	resp, err := client.DownloadStream(ctx, container, blobName, nil)
	if err != nil {
		return nil, nil, err
	}
	rc := resp.Body // io.ReadCloser
	closer := func() { _ = rc.Close() }
	return rc, closer, nil
}
