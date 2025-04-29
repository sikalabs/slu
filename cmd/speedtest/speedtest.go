package speedtest

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

func init() {
	root.RootCmd.AddCommand(Cmd)
}

var Cmd = &cobra.Command{
	Use:   "speedtest",
	Short: "Run a speedtest",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		example()
	},
}

func example() {
	var speedtestClient = speedtest.New()
	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()
		fmt.Printf("Latency: %s, Download: %s, Upload: %s\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset()
	}
}
