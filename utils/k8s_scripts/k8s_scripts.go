package k8s_scripts

import (
	"fmt"

	"github.com/sikalabs/slu/utils/sh_utils"
)

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
