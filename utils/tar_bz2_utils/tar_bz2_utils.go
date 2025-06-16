package tar_bz2_utils

import (
	"archive/tar"
	"compress/bzip2"
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

func WebTarBz2ToBin(url, inTarBz2FileName string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	handleError(err)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	uncompressedStream := bzip2.NewReader(resp.Body)
	handleError(err)
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		handleError(err)

		if header.Typeflag == tar.TypeReg &&
			(header.Name == inTarBz2FileName || header.Name == "./"+inTarBz2FileName) {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			handleError(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			handleError(err)
			return
		}
	}

	handleError(fmt.Errorf("file \"%s\" not found in bz2", inTarBz2FileName))
}
