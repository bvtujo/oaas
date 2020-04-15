package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var otterStatuses = []string{
	"frolicking",
	"smashing open a clam",
	"holding hands",
	"diving",
	"sunning on the pier",
	"splashing",
	"playing with a beach ball like a water dog",
}

var numOtterStatuses = len(otterStatuses)

func getRandomStatus() string {
	return otterStatuses[rand.Intn(numOtterStatuses)]
}

func bytef(in string, args ...interface{}) []byte {
	return []byte(fmt.Sprintf(in, args...))
}

func GetOtterStatus(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	stat := getRandomStatus()
	w.Write(bytef("the otter is %s", stat))
}

func HealthCheck(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := httprouter.New()
	r.GET("/healthcheck", HealthCheck)
	r.GET("/otter/status", GetOtterStatus)
	log.Fatal(http.ListenAndServe(":8080", r))
}
