package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to your beautiful site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Fprintln(os.Stdout, "Hello World!")
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
