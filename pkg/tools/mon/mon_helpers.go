package mon

import (
	"log"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskUsage struct {
	Total       uint64
	TotalGb     float64
	Used        uint64
	UsedGb      float64
	Free        uint64
	FreeGb      float64
	UsedPercent float64
}

func getDiskUsage() (DiskUsage, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		log.Printf("Error getting disk usage: %v", err)
		return DiskUsage{}, err
	}

	return DiskUsage{
		Total:       usage.Total,
		TotalGb:     float64(usage.Total) / 1e9,
		Used:        usage.Used,
		UsedGb:      float64(usage.Used) / 1e9,
		Free:        usage.Free,
		FreeGb:      float64(usage.Free) / 1e9,
		UsedPercent: usage.UsedPercent,
	}, nil
}
