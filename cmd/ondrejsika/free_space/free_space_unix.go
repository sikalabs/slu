//go:build linux || darwin
// +build linux darwin

package free_space

import (
	"syscall"
)

func getFreeSpaceOrDie() uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	handleError(err)
	freeSpace := stat.Bavail * uint64(stat.Bsize)
	return freeSpace
}
