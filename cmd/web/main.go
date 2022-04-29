package main

import (
	"fmt"
	"net/http"

	"github.com/vrodnenok/go-ship/pkg/handlers"
)

const PORT_NUMBER = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Listening at port %s", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
