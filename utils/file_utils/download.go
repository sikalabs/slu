package file_utils

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DownloadFile downloads the content from url into path.
// It creates parent directories as needed and writes atomically.
func DownloadFile(path, url string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	if url == "" {
		return errors.New("url is empty")
	}

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	// HTTP client with a sane timeout
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Read a small snippet to include in the error (optional)
		return errors.New("download failed: " + resp.Status)
	}

	// Write to a temp file in the same dir, then rename (atomic-ish)
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	tmp, err := os.CreateTemp(dir, "."+base+".*.part")
	if err != nil {
		return err
	}

	// Ensure cleanup on error
	tmpName := tmp.Name()
	defer func() {
		tmp.Close()
		_ = os.Remove(tmpName)
	}()

	// Stream copy
	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return err
	}

	// Flush to disk
	if err := tmp.Sync(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}

	// Set final permissions (temp files default to 0600)
	if err := os.Chmod(tmpName, 0o644); err != nil {
		return err
	}

	// Atomic rename within same filesystem
	if err := os.Rename(tmpName, path); err != nil {
		return err
	}

	return nil
}
