package stdin_utils

import (
	"bufio"
	"log"
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

func ReadFromPipeOrDie() string {
	var jwtToken string

	// Ensure input is from a pipe
	fi, err := os.Stdin.Stat()
	if err != nil || fi.Mode()&os.ModeNamedPipe == 0 {
		log.Fatalln("No input from pipe.")
	}

	// Read the input from stdin (pipe)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		jwtToken = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	return jwtToken
}
