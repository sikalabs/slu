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

func WebTarGzToBin(url, inTarGzFileName, outFileName string) {
	var err error

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	uncompressedStream, err := gzip.NewReader(resp.Body)
	handleError(err)
	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		handleError(err)

		if header.Typeflag == tar.TypeReg && header.Name == inTarGzFileName {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 755)
			handleError(err)
			defer outFile.Close()

			_, err = io.Copy(outFile, tarReader)
			handleError(err)
			return
		}
	}

	handleError(fmt.Errorf("file \"%s\" not found in tar.gz", inTarGzFileName))
}
