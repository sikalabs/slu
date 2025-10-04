package tar_bz2_utils

import (
	"archive/tar"
	"compress/bzip2"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sikalabs/slu/internal/error_utils"
)

func WebTarBz2ToBin(url, inTarBz2FileName string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	error_utils.HandleError(err, "Failed to create HTTP request")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	uncompressedStream := bzip2.NewReader(resp.Body)
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		error_utils.HandleError(err, "Failed to read tar header")

		if header.Typeflag == tar.TypeReg &&
			(header.Name == inTarBz2FileName || header.Name == "./"+inTarBz2FileName) {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			error_utils.HandleError(err, "Failed to create output file")
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			error_utils.HandleError(err, "Failed to copy file from tar")
			return
		}
	}

	error_utils.HandleError(fmt.Errorf("file \"%s\" not found in bz2", inTarBz2FileName), "File not found in bz2")
}
