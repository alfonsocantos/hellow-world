package main

import (
  "fmt"
	"net/http"
	"io"
  "github.com/satori/go.uuid"
)

func main () {
  port = ":8080"

  fmt.Printf("Listening on port %s\n", port)

	http.HandleFunc ("/", PrintString)
  http.HandleFunc ("/uuid", PrintUUID)
	http.ListenAndServe(port, nil)
}

// PrintString just prints the json Hola: Mundo! string
func PrintString (rw http.ResponseWriter, r *http.Request){
	io.WriteString (rw, "{\"Hola\": \"Mundo!\"}\n")
}

// PrintUUID just prints an UUID string in json
func PrintUUID (rw http.ResponseWriter, r *http.Request){
  // Creating UUID Version 4
  u1 := uuid.NewV4()
  // fmt.Printf("UUIDv4.2: %s\n", u1)
  io.WriteString (rw, "{\"uuid\": \""+ u1.String() +"\"}")
}
