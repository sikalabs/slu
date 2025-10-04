//go:build windows
// +build windows

package free_space

import (
	"github.com/sikalabs/slu/internal/error_utils"
	"golang.org/x/sys/windows"
)

func getFreeSpaceOrDie() uint64 {
	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64

	err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr("C:"),
		&freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	error_utils.HandleError(err, "Failed to get disk free space")
	return totalNumberOfFreeBytes
}
