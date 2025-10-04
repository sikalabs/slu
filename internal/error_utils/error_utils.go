package error_utils

import (
	"log"
)

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalln(message, err)
	}
}

func HandleNotOK(ok bool, message string) {
	if !ok {
		log.Fatalln(message)
	}
}
