package zip_utils

import (
	"archive/zip"
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

func WebZipToBin(url, inZipFileName, outFileName string) {
	var err error

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	tmpInFile, err := os.CreateTemp("", "go-zip-example")
	handleError(err)

	_, err = io.Copy(tmpInFile, resp.Body)
	handleError(err)

	r, err := zip.OpenReader(tmpInFile.Name())
	handleError(err)
	defer r.Close()

	for _, f := range r.File {
		if f.Name == inZipFileName {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			handleError(err)
			defer outFile.Close()

			zipFile, _ := f.Open()
			defer zipFile.Close()

			_, err = io.Copy(outFile, zipFile)
			handleError(err)
			return
		}
	}
	handleError(fmt.Errorf("file \"%s\" not found in ZIP", inZipFileName))
}
