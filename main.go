package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3001", nil) //ip address and port
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("menu.txt") // get content from file and ignore error
	io.Copy(w, f) // take content and copy to ResponseWriter

}
