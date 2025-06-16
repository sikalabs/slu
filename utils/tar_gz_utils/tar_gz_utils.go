package tar_gz_utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WebTarGzToBin(url, inTarGzFileName string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	handleError(err)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	uncompressedStream, err := gzip.NewReader(resp.Body)
	handleError(err)
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		handleError(err)

		if header.Typeflag == tar.TypeReg &&
			(header.Name == inTarGzFileName || header.Name == "./"+inTarGzFileName) {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			handleError(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			handleError(err)
			return
		}
	}

	handleError(fmt.Errorf("file \"%s\" not found in tar.gz", inTarGzFileName))
}
