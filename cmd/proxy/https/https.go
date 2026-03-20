package https

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	proxy_cmd "github.com/sikalabs/slu/cmd/proxy"
	"github.com/spf13/cobra"
)

var CmdFlagLocalAddr string
var CmdFlagRemoteAddr string
var CmdFlagCertFile string
var CmdFlagKeyFile string

var Cmd = &cobra.Command{
	Use:   "https",
	Short: "HTTPS to HTTP Proxy (expose HTTPS, forward to HTTP)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		target, err := url.Parse("http://" + CmdFlagRemoteAddr)
		if err != nil {
			log.Fatalf("error: invalid remote address: %s", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		var tlsCert tls.Certificate
		if CmdFlagCertFile != "" && CmdFlagKeyFile != "" {
			tlsCert, err = tls.LoadX509KeyPair(CmdFlagCertFile, CmdFlagKeyFile)
			if err != nil {
				log.Fatalf("error: failed to load cert/key: %s", err)
			}
		} else {
			tlsCert, err = generateSelfSignedCert()
			if err != nil {
				log.Fatalf("error: failed to generate self-signed cert: %s", err)
			}
			fmt.Println("Using auto-generated self-signed certificate")
		}

		server := &http.Server{
			Addr:    CmdFlagLocalAddr,
			Handler: proxy,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{tlsCert},
			},
		}

		fmt.Printf("Starting HTTPS proxy on %s -> http://%s\n", CmdFlagLocalAddr, CmdFlagRemoteAddr)
		log.Fatal(server.ListenAndServeTLS("", ""))
	},
}

func generateSelfSignedCert() (tls.Certificate, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"slu proxy https"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return tls.Certificate{}, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})

	return tls.X509KeyPair(certPEM, keyPEM)
}

func init() {
	proxy_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagLocalAddr,
		"local",
		"l",
		"",
		"Local address (eg. :8443)",
	)
	Cmd.MarkFlagRequired("local")
	Cmd.Flags().StringVarP(
		&CmdFlagRemoteAddr,
		"remote",
		"r",
		"",
		"Remote address (eg. localhost:8080)",
	)
	Cmd.MarkFlagRequired("remote")
	Cmd.Flags().StringVarP(
		&CmdFlagCertFile,
		"cert",
		"c",
		"",
		"TLS certificate file (optional, auto-generated if not set)",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagKeyFile,
		"key",
		"k",
		"",
		"TLS key file (optional, auto-generated if not set)",
	)
}
