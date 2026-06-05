package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/terrence-beckham/hello"
)

func main() {
	hello.PrintTo(os.Stdout)
	fmt.Println("Listening on http://localhost:9001")
	http.ListenAndServe(":9001", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			hello.PrintTo(w)
		},
	))
}
