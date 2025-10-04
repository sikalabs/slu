//go:build linux || darwin
// +build linux darwin

package free_space

import (
	"syscall"

	"github.com/sikalabs/slu/internal/error_utils"
)

func getFreeSpaceOrDie() uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	error_utils.HandleError(err, "Failed to get disk free space")
	freeSpace := stat.Bavail * uint64(stat.Bsize)
	return freeSpace
}
