package token_expiration_alert

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagExpirationDate string
var FlagMessage string

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(&FlagExpirationDate, "expiration-date", "e", "", "Token expiration date (eg. 2024-01-01)")
	Cmd.MarkFlagRequired("expiration-date")
	Cmd.Flags().StringVarP(&FlagMessage, "message", "m", "", "Message to send")
	Cmd.MarkFlagRequired("message")
}

var Cmd = &cobra.Command{
	Use:   "token-expiration-alert",
	Short: "Alert about token expiration",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		expirationAlert(FlagExpirationDate, FlagMessage)
	},
}

func expirationAlert(expitationDateString, message string) {
	expirationDate, err := time.Parse("2006-01-02", expitationDateString)
	if err != nil {
		log.Fatalln(err)
		return
	}

	currentDate := time.Now()

	// Check if the token has expired
	if currentDate.After(expirationDate) || currentDate.Equal(expirationDate) {
		printInBox(message)
		os.Exit(1)
	}
}

func printInBox(text string) {
	const maxLineLength = 60
	lines := splitText(text, maxLineLength)

	// Determine the length of the longest line
	maxLength := len(lines[0])
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Create the top and bottom border
	borderLength := maxLength + 8 // Additional 8 to account for padding and border
	border := strings.Repeat("#", borderLength)
	emptyLine := "#" + strings.Repeat(" ", borderLength-2) + "#"

	// Print the box with the text
	fmt.Println(border)
	fmt.Println(emptyLine)
	fmt.Println(emptyLine)
	for _, line := range lines {
		fmt.Printf("#   %-*s   #\n", maxLength, line)
	}
	fmt.Println(emptyLine)
	fmt.Println(emptyLine)
	fmt.Println(border)
}

// Helper function to split text into lines of a specific length
func splitText(text string, length int) []string {
	var lines []string
	for len(text) > length {
		lines = append(lines, text[:length])
		text = text[length:]
	}
	lines = append(lines, text)
	return lines
}
