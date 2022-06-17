package file_server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Server(port int) {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Printf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
