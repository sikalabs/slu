package zip_utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sikalabs/slu/internal/error_utils"
)

func WebZipToBin(url, inZipFileName string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	error_utils.HandleError(err, "Failed to create HTTP request")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	tmpInFile, err := os.CreateTemp("", "go-zip-example")
	error_utils.HandleError(err, "Failed to create temp file")

	_, err = io.Copy(tmpInFile, resp.Body)
	error_utils.HandleError(err, "Failed to copy response to temp file")

	r, err := zip.OpenReader(tmpInFile.Name())
	error_utils.HandleError(err, "Failed to open zip file")
	defer r.Close()

	for _, f := range r.File {
		if f.Name == inZipFileName {
			outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
			error_utils.HandleError(err, "Failed to create output file")
			defer outFile.Close()

			zipFile, _ := f.Open()
			defer zipFile.Close()

			_, err = io.Copy(outFile, zipFile)
			error_utils.HandleError(err, "Failed to copy file from zip")
			return
		}
	}
	error_utils.HandleError(fmt.Errorf("file \"%s\" not found in ZIP", inZipFileName), "File not found in ZIP")
}
