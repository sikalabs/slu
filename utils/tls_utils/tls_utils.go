package tls_utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"
)

func PrintCertificateFromServer(addr string, serverName string) {
	var conf *tls.Config
	if serverName == "" {
		conf = &tls.Config{
			InsecureSkipVerify: true,
		}
	} else {
		conf = &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         serverName,
		}
	}

	conn, err := tls.Dial("tcp", addr, conf)
	if err != nil {
		log.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Subject Name: %s\n", cert.Subject)
		fmt.Printf("Subject Common Name: %s\n", cert.Subject.CommonName)
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Issuer Common Name: %s \n", cert.Issuer.CommonName)
		fmt.Printf("Created: %s \n", cert.NotBefore.Format(time.RFC3339))
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format(time.RFC3339))
		fmt.Println()
	}
}

func PrintCertificateFromFile(certFile, keyFile string) {
	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	certs, err := x509.ParseCertificates(tlsCert.Certificate[0])
	if err != nil {
		log.Fatal(err)
	}
	for _, cert := range certs {
		fmt.Printf("Subject Name: %s\n", cert.Subject)
		fmt.Printf("Subject Common Name: %s\n", cert.Subject.CommonName)
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Issuer Common Name: %s \n", cert.Issuer.CommonName)
		fmt.Printf("Created: %s \n", cert.NotBefore.Format(time.RFC3339))
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format(time.RFC3339))
		fmt.Println()
	}
}
