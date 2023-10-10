//go:build windows
// +build windows

package free_space

import "golang.org/x/sys/windows"

func getFreeSpaceOrDie() uint64 {
	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64

	err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr("C:"),
		&freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	handleError(err)
	return totalNumberOfFreeBytes
}
