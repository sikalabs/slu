package install_bin_utils

import (
	"fmt"
	"log"
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
	log.Fatal(fmt.Errorf("unknown suffix (no .zip, .tar.gz or .tgz)"))
}
