package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my awesome site. Lets learn GO! </h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Fprintln(os.Stdout, "hello world")
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)

}
