package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "backup <file>",
	Short: "Create a numbered backup of a file in the same directory",
	Long: `Create a backup of a file with format: .<basename>.<number>.<YYYY-mm-dd_HH-MM>.ext.backup

Example:
  slu scripts backup /etc/nginx/nginx.conf
  Creates: /etc/nginx/.nginx.1.2025-11-05_14-30.conf.backup`,
	Args: cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		filePath := args[0]
		backupFile, err := createBackup(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Backup file created:", backupFile)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func createBackup(filePath string) (string, error) {
	// Check if source file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Parse the file path
	dir := filepath.Dir(filePath)
	filename := filepath.Base(filePath)

	// Split filename into name and extension
	ext := filepath.Ext(filename)
	basename := strings.TrimSuffix(filename, ext)

	// Get current timestamp
	now := time.Now()
	timestamp := now.Format("2006-01-02_15-04")

	// Find the next backup number
	backupNum := getNextBackupNumber(dir, basename, ext)

	// Create backup filename
	backupFilename := fmt.Sprintf(".%s.%d.%s%s.backup", basename, backupNum, timestamp, ext)
	backupPath := filepath.Join(dir, backupFilename)

	// Read source file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read source file: %w", err)
	}

	// Write backup file
	err = os.WriteFile(backupPath, content, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write backup file: %w", err)
	}

	return backupPath, nil
}

func getNextBackupNumber(dir, basename, ext string) int {
	// Pattern to match existing backups: .<basename>.<number>.*<ext>.backup
	pattern := fmt.Sprintf(".%s.*%s.backup", basename, ext)
	matches, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil || len(matches) == 0 {
		return 1
	}

	maxNum := 0
	prefix := fmt.Sprintf(".%s.", basename)
	suffix := fmt.Sprintf("%s.backup", ext)

	for _, match := range matches {
		filename := filepath.Base(match)
		if !strings.HasPrefix(filename, prefix) || !strings.HasSuffix(filename, suffix) {
			continue
		}

		// Extract the number part: .basename.<number>.<timestamp>.ext.backup
		afterPrefix := strings.TrimPrefix(filename, prefix)
		parts := strings.Split(afterPrefix, ".")
		if len(parts) < 2 {
			continue
		}

		var num int
		_, err := fmt.Sscanf(parts[0], "%d", &num)
		if err == nil && num > maxNum {
			maxNum = num
		}
	}

	return maxNum + 1
}
