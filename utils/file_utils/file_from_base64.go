package file_utils

import (
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateFileFromBase64 writes the decoded base64 data into the given file path,
// creating all necessary parent directories.
func CreateFileFromBase64(path, b64 string) error {
	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	// Create/truncate the file
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decode base64 and stream to file
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	_, err = io.Copy(f, dec)
	return err
}
