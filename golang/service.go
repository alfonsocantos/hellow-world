package main

import (
  "fmt"
  "net/http"
  "io"
  "github.com/satori/go.uuid"
  "math/rand"
  "os"
  "strconv"
  "time"
  "encoding/json"
  "crypto/md5"
)


// HelloResponse struct
type HelloResponse struct {
  Message string  `json:"message"`
}

// UUIDResponse struct
type UUIDResponse struct {
  UUID string  `json:"uuid"`
}

// AccountResponse struct
type AccountResponse struct {
  Code string `json:"code"`
}

func main () {

  if b, err := ComputeMd5(os.Args[0]); err == nil {
    fmt.Printf("version: %x \n", b)
  }

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  fmt.Printf("Listening on port %s...\n", port)

//  http.HandleFunc ("/", BasePath)
  http.HandleFunc ("/hello", PrintString)
  http.HandleFunc ("/uuid", PrintUUID)
  http.HandleFunc ("/accounts", Accounts)
  http.HandleFunc ("/account", Account)
  http.HandleFunc ("/404", Error404)
  http.HandleFunc ("/500", Error500)
	http.ListenAndServe(":" + port, nil)
}

// PrintString just prints the json Hola: Mundo! string
func PrintString (rw http.ResponseWriter, r *http.Request){
  if RandomizeResponse (rw, r) {

    msg, _ := json.Marshal (&HelloResponse{Message: "Hello World!"})

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    io.WriteString (rw, string(msg))
  }
}

// PrintUUID just prints an UUID string in json
func PrintUUID (rw http.ResponseWriter, r *http.Request){
  if RandomizeResponse (rw, r) {

    u1, _ := json.Marshal (&UUIDResponse{UUID: uuid.NewV4().String()})

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    io.WriteString (rw, string(u1))
  }
}

// Account just prints an UUID string in json
func Account (rw http.ResponseWriter, r *http.Request){
  if RandomizeResponse (rw, r) {

    msg, _ := json.Marshal (&AccountResponse{Code: "ok"})

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    io.WriteString (rw, string (msg))
  }
}


// Accounts serves the accounts.json file
func Accounts(rw http.ResponseWriter, r *http.Request) {
  if RandomizeResponse (rw, r) {
    fp := "accounts.json"

    rw.Header().Set("Content-Type", "application/json")
    http.ServeFile(rw, r, fp)
  }
}

// Error500 serves the accounts.json file
func Error500(rw http.ResponseWriter, r *http.Request) {
  if RandomizeResponse (rw, r) {
    // 500
    http.Error(rw, "500 Interval Server Error", http.StatusInternalServerError)
  }
}

// Error404 serves the accounts.json file
func Error404(rw http.ResponseWriter, r *http.Request) {
  if RandomizeResponse (rw, r) {
    http.NotFound(rw, r)
  }
}

// BasePath tal
//func BasePath (rw http.ResponseWriter, r *http.Request) {
//  http.Redirect (rw, r, "/hello", http.StatusFound)
//}

// RandomizeResponse prints the body, or erros in a random fashion
func RandomizeResponse (rw http.ResponseWriter, r *http.Request) bool {

  // Environment variables:
  // RANDOM404: How many responses are going to be 404
  // RANDOM500: How many responses are going to be 500
  // RANDOMSLEEP: Sleep time in miliseconds
  random404, _ := strconv.Atoi(os.Getenv ("RANDOM404"))
  random500, _ := strconv.Atoi(os.Getenv ("RANDOM500"))
  randomMaxTime, _ := strconv.Atoi(os.Getenv ("RANDOMMAXTIME"))
  randomMinTime, _ := strconv.Atoi(os.Getenv ("RANDOMMINTIME"))

  random500 += random404
  randomResponse := rand.Intn(100)

  if randomMaxTime > 0 {
      time.Sleep(time.Duration (rand.Intn(randomMaxTime - randomMinTime) + randomMinTime) * time.Millisecond)
  }

  result := true
  if (random404 != 0) && randomResponse <= random404 {
    // 404
    http.NotFound(rw, r)
    result = false
  } else if (random500 != 0) && randomResponse <= random500{
    // 500
    http.Error(rw, "500 Interval Server Error", http.StatusInternalServerError)
    result = false
  }

  return result
}

//ComputeMd5 does it
func ComputeMd5(filePath string) ([]byte, error) {
  var result []byte
  file, err := os.Open(filePath)
  if err != nil {
    return result, err
  }
  defer file.Close()

  hash := md5.New()
  if _, err := io.Copy(hash, file); err != nil {
    return result, err
  }

  return hash.Sum(result), nil
}
