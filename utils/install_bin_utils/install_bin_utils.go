package install_bin_utils

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/tar_bz2_utils"
	"github.com/sikalabs/slu/utils/tar_gz_utils"
	"github.com/sikalabs/slu/utils/zip_utils"
)

func InstallBin(url, source, binDir, name string, exeSuffix bool) {
	nameTmp := name + ".new"
	if exeSuffix {
		source = source + ".exe"
		name = name + ".exe"
		nameTmp = nameTmp + ".exe"
	}
	if strings.HasSuffix(url, "zip") {
		zip_utils.WebZipToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, nameTmp),
		)
		os.Rename(path.Join(binDir, nameTmp), path.Join(binDir, name))
		return
	}
	if strings.HasSuffix(url, "tar.gz") || strings.HasSuffix(url, "tgz") {
		tar_gz_utils.WebTarGzToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, nameTmp),
		)
		os.Rename(path.Join(binDir, nameTmp), path.Join(binDir, name))
		return
	}
	if strings.HasSuffix(url, "tar.bz2") || strings.HasSuffix(url, "tbz2") {
		tar_bz2_utils.WebTarBz2ToBin(
			url,
			source,
			map[string]string{},
			path.Join(binDir, nameTmp),
		)
		os.Rename(path.Join(binDir, nameTmp), path.Join(binDir, name))
		return
	}
	webToBin(
		url,
		map[string]string{},
		path.Join(binDir, nameTmp),
	)
	os.Rename(path.Join(binDir, nameTmp), path.Join(binDir, name))
}

func webToBin(url string, headers map[string]string, outFileName string) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	error_utils.HandleError(err, "Failed to create HTTP request")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	error_utils.HandleError(err, "Failed to execute HTTP request")
	defer resp.Body.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
	error_utils.HandleError(err, "Failed to create output file")
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	error_utils.HandleError(err, "Failed to copy response to file")
}
