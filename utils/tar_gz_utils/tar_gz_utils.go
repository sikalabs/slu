package tar_gz_utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sikalabs/slu/internal/error_utils"
)

func WebTarGzToBin(url, inTarGzFileName string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	error_utils.HandleError(err, "Failed to create HTTP request")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	uncompressedStream, err := gzip.NewReader(resp.Body)
	error_utils.HandleError(err, "Failed to create gzip reader")
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		error_utils.HandleError(err, "Failed to read tar header")

		if header.Typeflag == tar.TypeReg &&
			(header.Name == inTarGzFileName || header.Name == "./"+inTarGzFileName) {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			error_utils.HandleError(err, "Failed to create output file")
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			error_utils.HandleError(err, "Failed to copy file from tar")
			return
		}
	}

	error_utils.HandleError(fmt.Errorf("file \"%s\" not found in tar.gz", inTarGzFileName), "File not found in tar.gz")
}
