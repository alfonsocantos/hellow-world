package main

import (
	"net/http"
	"io"
)

func main () {
	http.HandleFunc ("/markdown", PrintString)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func PrintString (rw http.ResponseWriter, r *http.Request){
	io.WriteString (rw, "Hola Mundo!\n")
}
