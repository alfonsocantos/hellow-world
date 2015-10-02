package main

import {
	"net/http"

	"github.com/russross/backfriday"	
}

func main () {
	http.HandleFunc ("/markdown", PrintString)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func PrintString (rw http.ResponseWriter, r *http.Request){
	rw.Write ("Hola Mundo!")
}
