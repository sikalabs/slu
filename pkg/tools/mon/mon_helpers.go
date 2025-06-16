package mon

import (
	"log"
	"strconv"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskUsage struct {
	Total          uint64
	TotalGb        float64
	TotalGbStr     string
	Used           uint64
	UsedGb         float64
	UsedGbStr      string
	Free           uint64
	FreeGb         float64
	FreeGbStr      string
	UsedPercent    float64
	UsedPercentStr string
}

func getDiskUsage() (DiskUsage, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error getting disk usage: %v", err)
		return DiskUsage{}, err
	}

	return DiskUsage{
		Total:          usage.Total,
		TotalGb:        float64(usage.Total) / 1e9,
		TotalGbStr:     strconv.FormatFloat(float64(usage.Total)/1e9, 'f', 2, 64) + " GB",
		Used:           usage.Used,
		UsedGb:         float64(usage.Used) / 1e9,
		UsedGbStr:      strconv.FormatFloat(float64(usage.Used)/1e9, 'f', 2, 64) + " GB",
		Free:           usage.Free,
		FreeGb:         float64(usage.Free) / 1e9,
		FreeGbStr:      strconv.FormatFloat(float64(usage.Free)/1e9, 'f', 2, 64) + " GB",
		UsedPercent:    usage.UsedPercent,
		UsedPercentStr: strconv.FormatFloat(usage.UsedPercent, 'f', 2, 64) + "%",
	}, nil
}
