package error_utils

import (
	"log"
)

func HandleError(err error, message ...string) {
	if len(message) == 0 {
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		if err != nil {
			log.Fatalln(message[0], err)
		}
	}
}

func HandleNotOK(ok bool, message string) {
	if !ok {
		log.Fatalln(message)
	}
}
