package ssh_utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"

	"github.com/sikalabs/slu/utils/exec_utils"
	"golang.org/x/crypto/ssh"
)

// https://stackoverflow.com/a/64178933/5281724

func MakeSSHKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return "", "", err
	}

	// generate and write private key as PEM
	var privKeyBuf strings.Builder

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(&privKeyBuf, privateKeyPEM); err != nil {
		return "", "", err
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	var pubKeyBuf strings.Builder
	pubKeyBuf.Write(ssh.MarshalAuthorizedKey(pub))

	return pubKeyBuf.String(), privKeyBuf.String(), nil
}

func MakeSSHKeyPairSSHKeyGen(length int) (string, string, error) {
	var err error
	err = exec_utils.ExecNoOut("rm", "-rf", "/tmp/slu_id_rsa", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	err = exec_utils.ExecNoOut("ssh-keygen", "-t", "rsa", "-b", strconv.Itoa(length), "-N", "", "-C", "", "-f", "/tmp/slu_id_rsa")
	if err != nil {
		return "", "", err
	}
	pub, err := exec_utils.ExecStr("cat", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	priv, err := exec_utils.ExecStr("cat", "/tmp/slu_id_rsa")
	if err != nil {
		return "", "", err
	}
	err = exec_utils.ExecNoOut("rm", "-rf", "/tmp/slu_id_rsa", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	return pub, priv, nil
}

func MakeSSHKeyPairSSHKeyGenECDSA() (string, string, error) {
	var err error
	err = exec_utils.ExecNoOut("rm", "-rf", "/tmp/slu_id_rsa", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	err = exec_utils.ExecNoOut("ssh-keygen", "-t", "ecdsa", "-b", "521", "-N", "", "-C", "", "-f", "/tmp/slu_id_rsa")
	if err != nil {
		return "", "", err
	}
	pub, err := exec_utils.ExecStr("cat", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	priv, err := exec_utils.ExecStr("cat", "/tmp/slu_id_rsa")
	if err != nil {
		return "", "", err
	}
	err = exec_utils.ExecNoOut("rm", "-rf", "/tmp/slu_id_rsa", "/tmp/slu_id_rsa.pub")
	if err != nil {
		return "", "", err
	}
	return pub, priv, nil
}
