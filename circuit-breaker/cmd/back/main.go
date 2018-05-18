package main

import (
	"fmt"
	"log"
	"net/http"
)

type profile struct{}

func handle(w http.ResponseWriter, res *http.Request) {
	fmt.Fprint(w, "12345")
}

func main() {
	http.HandleFunc("/account", handle)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
