package generate_files_utils

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/sikalabs/slu/utils/file_utils"
	"github.com/sikalabs/slu/utils/random_utils"
)

const MB int = 1024 * 1024

func GenerateFile(path string, sizeMB int) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r, _ := os.Open("/dev/urandom")
	io.CopyN(f, r, int64(sizeMB*MB))
}

func GenerateFiles(basePath string, sizeMB int, count int) {
	for i := 0; i < count; i++ {
		dir := path.Join(
			basePath,
			random_utils.RandomString(2, random_utils.LOWER),
			random_utils.RandomString(2, random_utils.LOWER),
		)
		filepath := path.Join(
			dir,
			random_utils.RandomString(6, random_utils.LOWER),
		)
		file_utils.EnsureDir(dir)
		GenerateFile(filepath, sizeMB)
	}
}
