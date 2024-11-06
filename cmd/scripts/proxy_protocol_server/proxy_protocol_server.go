package proxy_protocol_server

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"

	"net"
	"net/http"
	"time"

	"github.com/pires/go-proxyproto"
	h2proxy "github.com/pires/go-proxyproto/helper/http2"
)

var Cmd = &cobra.Command{
	Use:   "proxy-protocol-server",
	Short: "Start Proxy Protocol HTTP Server on port 8000",
	Run: func(c *cobra.Command, args []string) {
		server()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func server() {
	server := http.Server{
		Addr: ":8000",
		ConnState: func(c net.Conn, s http.ConnState) {
			if s == http.StateNew {
				log.Printf("[ConnState] %s -> %s\n", c.LocalAddr().String(), c.RemoteAddr().String())
			}
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[Handler] remote ip %q\n", r.RemoteAddr)
			w.Write([]byte(fmt.Sprintf("remote ip %q\n", r.RemoteAddr)))
		}),
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	proxyListener := &proxyproto.Listener{
		Listener:          ln,
		ReadHeaderTimeout: 10 * time.Second,
	}
	defer proxyListener.Close()

	// Create an HTTP server which can handle proxied incoming connections for
	// both HTTP/1 and HTTP/2. HTTP/2 support relies on TLS ALPN, the reverse
	// proxy needs to be configured to accept "h2".
	h2proxy.NewServer(&server, nil).Serve(proxyListener)
}
