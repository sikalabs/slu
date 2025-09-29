package rabbitmq

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"math/big"
	"net"
	"time"
)

func ConnectToRabbitMQ(ssl bool, user, password, host string, port int, vhost, certPath, keyPath string) (*amqp091.Connection, *amqp091.Channel, error) {
	amqpURL := createUrl(ssl, user, password, host, port, vhost)

	var err error
	var conn *amqp091.Connection
	var ch *amqp091.Channel

	if ssl {
		tlsConfig, err := newTLSConfig(certPath, keyPath)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to create TLS config: %s", err)
		}
		conn, err = amqp091.DialTLS(amqpURL, tlsConfig)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to connect to RabbitMQ: %s", err)
		}
	} else {
		conn, err = amqp091.Dial(amqpURL)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to connect to RabbitMQ: %s", err)
		}
	}

	ch, err = conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open a channel: %s", err)
	}
	return conn, ch, nil

}

func createUrl(ssl bool, user, password, host string, port int, vhost string) string {
	if ssl {
		return fmt.Sprintf("amqps://%s:%s@%s:%d/%s",
			user,
			password,
			host,
			port,
			vhost,
		)
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		user,
		password,
		host,
		port,
		vhost,
	)
}

func newTLSConfig(clientCertPath, clientKeyPath string) (*tls.Config, error) {
	var clientCert tls.Certificate
	var err error

	if clientCertPath != "" && clientKeyPath != "" {
		clientCert, err = tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load client certificate and key: %v", err)
		}
	} else {
		certOut, keyOut, err := generateSelfSignedCert()
		if err != nil {
			return nil, fmt.Errorf("failed to generate Self signed ertificate: %v", err)
		}
		clientCert, err = tls.X509KeyPair(certOut, keyOut)
		if err != nil {
			return nil, fmt.Errorf("failed to create TLS certificate from Let's Encrypt cert: %v", err)
		}
	}

	return &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{clientCert},
	}, nil
}

func generateSelfSignedCert() (certOut, keyOut []byte, err error) {
	// Create private key
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Set up certificate
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My Organization"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),

		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		IsCA:        false,
	}

	// Create certificate
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	// Convert certificate to PEM format
	certOut = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}
	keyOut = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

	return certOut, keyOut, nil
}
