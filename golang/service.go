package main

import (
	"crypto/md5"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

// HelloResponse struct
type HelloResponse struct {
	Message string `json:"message"`
}

// UUIDResponse struct
type UUIDResponse struct {
	UUID string `json:"uuid"`
}

// AccountResponse struct
type AccountResponse struct {
	Code string `json:"code"`
}

func main() {

	if b, err := ComputeMd5(os.Args[0]); err == nil {
		log.Printf("version: %x \n", b)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s...\n", port)

	http.HandleFunc("/basic/hello", getHelloWorld)
	http.HandleFunc("/basic/uuid", getUUID)
	http.HandleFunc("/basic/accounts", getAccounts)
	http.HandleFunc("/basic/account", postAccount)
	http.HandleFunc("/basic/echo", postEcho)
	http.HandleFunc("/", getDefault)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println(err)
	}
}

func getHelloWorld(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	randomSleepTime()

	msg, _ := json.Marshal(&HelloResponse{Message: "Hello World!"})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(msg))

	elapsed := time.Since(start)
	log.Printf("%s took %s\n", r.URL, elapsed)
}

func getUUID(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	randomSleepTime()

	u1, _ := json.Marshal(&UUIDResponse{UUID: uuid.NewV4().String()})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(u1))

	elapsed := time.Since(start)
	log.Printf("%s took %s\n", r.URL, elapsed)
}

func postAccount(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	randomSleepTime()

	msg, _ := json.Marshal(&AccountResponse{Code: "ok"})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	io.WriteString(rw, string(msg))

	elapsed := time.Since(start)
	log.Printf("%s took %s\n", r.URL, elapsed)
}

func getAccounts(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	randomSleepTime()

	data, err := ioutil.ReadFile("accounts.json")
	if err != nil {
		log.Println(err)
		http.Error(rw, "Error sent from golang service", http.StatusTeapot)
		return
	}

	u1 := uuid.NewV4().String()
	dataString := strings.Replace(string(data), "XXXXXXXXX", string(u1)[:rand.Intn(31)], -1)
	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, dataString)

	elapsed := time.Since(start)
	log.Printf("%s took %s\n", r.URL, elapsed)
}

func randomSleepTime() {
	randomMaxTime, _ := strconv.Atoi(os.Getenv("RANDOMMAXTIME"))
	randomMinTime, _ := strconv.Atoi(os.Getenv("RANDOMMINTIME"))

	if randomMaxTime > 0 {
		time.Sleep(time.Duration(rand.Intn(randomMaxTime-randomMinTime+1)+randomMinTime) * time.Millisecond)
	}
}

//ComputeMd5 does it to know the version of the service
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

func getDefault(rw http.ResponseWriter, r *http.Request) {
	msg := "Sorry, no path registrered in " + r.URL.String() + ". Default Response!	"

	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusNotFound)
	io.WriteString(rw, msg)

	log.Printf("not_found %s. \n", r.URL)
}

func postEcho(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()

	randomSleepTime()

	dataString, _ := ioutil.ReadAll(r.Body)
	rw.Header().Set("Content-Type", "application/json")
	io.WriteString(rw, string(dataString[:]))

	elapsed := time.Since(start)
	log.Printf("%s took %s\n", r.URL, elapsed)

}
