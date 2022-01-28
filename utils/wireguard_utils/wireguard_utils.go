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
	priv, preshared, pub, _ := NewKeys()
	fmt.Printf("priv:\t%s\n", priv)
	fmt.Printf("publ:\t%s\n", pub)
	fmt.Printf("pres:\t%s\n", preshared)
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
