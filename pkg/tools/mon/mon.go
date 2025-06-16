package mon

import "log"

func Mon() {
	usage, err := getDiskUsage()
	if err != nil {
		log.Printf("Error getting disk usage: %v", err)
	}
	log.Printf(
		"Disk Usage: Total: %.2f GB, Used: %.2f GB, Free: %.2f GB, Used Percent: %.2f%%",
		usage.TotalGb, usage.UsedGb, usage.FreeGb, usage.UsedPercent,
	)
}
