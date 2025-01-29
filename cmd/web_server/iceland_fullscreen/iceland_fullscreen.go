package iceland_fullscreen

import (
	"fmt"
	"net/http"

	"github.com/ondrejsika/go-iceland"
	parentcmd "github.com/sikalabs/slu/cmd/web_server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "iceland-fullscreen",
	Short: "Webserver with full screen pictures from Iceland",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server()
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}

func server() {
	http.HandleFunc("/photo.jpg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(iceland.ICELAND_RIVER_AT_POOL_2022)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Iceland</title>
<style>
body {
background-image: url(/photo.jpg);
background-size: cover;
background-repeat: no-repeat;
background-attachment: fixed;
background-position: center;
}
</style>
</head>
<html>
`))
	})
	fmt.Println("Listen on 0.0.0.0:8000, see http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
