package main

import (
	"fmt"
	"net/http"
	"os"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "G'day")
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
