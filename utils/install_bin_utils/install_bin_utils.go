package install_bin_utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/sikalabs/slu/utils/tar_gz_utils"
	"github.com/sikalabs/slu/utils/zip_utils"
)

func InstallBin(url, source, binDir, name string) {
	if strings.HasSuffix(url, "zip") {
		zip_utils.WebZipToBin(
			url,
			source,
			path.Join(binDir, name),
		)
		return
	}
	if strings.HasSuffix(url, "tar.gz") || strings.HasSuffix(url, "tgz") {
		tar_gz_utils.WebTarGzToBin(
			url,
			source,
			path.Join(binDir, name),
		)
		return
	}
	webToBin(
		url,
		path.Join(binDir, name),
	)
}

func webToBin(url, outFileName string) {
	var err error

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
	handleError(err)
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
