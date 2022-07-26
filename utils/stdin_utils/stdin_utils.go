package stdin_utils

import (
	"bufio"
	"os"
)

func ReadAll() string {
	var all string
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		all = all + text
	}
	return all
}
