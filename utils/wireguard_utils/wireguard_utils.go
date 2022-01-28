package wireguard_utils

import (
	"encoding/json"
	"fmt"

	wg "golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func NewKeys() (string, string, string, error) {
	priv, _ := wg.GeneratePrivateKey()
	preshared, _ := wg.GenerateKey()
	pub := priv.PublicKey()
	return priv.String(), pub.String(), preshared.String(), nil
}

func PrintNewKeys() {
	priv, pub, preshared, _ := NewKeys()
	fmt.Printf("priv: %s\n", priv)
	fmt.Printf("publ: %s\n", pub)
	fmt.Printf("pres: %s\n", preshared)
}

func PrintNewKeysJson() {
	priv, preshared, pub, _ := NewKeys()
	data := map[string]string{
		"private":   priv,
		"public":    pub,
		"preshared": preshared,
	}
	dataJsonBin, _ := json.Marshal(data)
	fmt.Println(string(dataJsonBin))
}
