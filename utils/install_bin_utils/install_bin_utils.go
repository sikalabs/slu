package install_bin_utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/sikalabs/slu/utils/tar_bz2_utils"
	"github.com/sikalabs/slu/utils/tar_gz_utils"
	"github.com/sikalabs/slu/utils/zip_utils"
)

func InstallBin(url, source, binDir, name string, exeSuffix bool) {
	if exeSuffix {
		source = source + ".exe"
		name = name + ".exe"
	}
	if strings.HasSuffix(url, "zip") {
		zip_utils.WebZipToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, name),
		)
		return
	}
	if strings.HasSuffix(url, "tar.gz") || strings.HasSuffix(url, "tgz") {
		tar_gz_utils.WebTarGzToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, name),
		)
		return
	}
	if strings.HasSuffix(url, "tar.bz2") || strings.HasSuffix(url, "tbz2") {
		tar_bz2_utils.WebTarBz2ToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, name),
		)
		return
	}
	webToBin(
		url,
		map[string]string{},
		name,
	)
}

func webToBin(url string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	handleError(err)

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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
