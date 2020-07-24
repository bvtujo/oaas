package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/julienschmidt/httprouter"
)

var otterStatuses = []string{
	"smashing open a clam",
	"holding hands",
	"frolicking",
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

func writeS3() error {
	sess := session.Must(session.NewSession())

	uploader := s3manager.NewUploader(sess)

	svcName := os.Getenv("COPILOT_SERVICE_NAME")
	bucketname := os.Getenv("S3SSL_NAME")
	now := time.Now().Unix()
	fname := fmt.Sprintf("%s-test-startup-%d", svcName, now)

	f, err := os.Create(fname)
	if err != nil {
		log.Printf("error creating file: %w\n", err)
		return err
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketname),
		Key:    aws.String(fname),
		Body:   f,
	})
	if err != nil {
		log.Printf("error uploading: %w\n")
		return err
	}
	log.Printf("success upload: %s", result.Location)
	return nil
}

func main() {
	writeS3()
	r := httprouter.New()
	r.GET("/healthcheck", HealthCheck)
	r.GET("/", GetOtterStatus)
	log.Fatal(http.ListenAndServe(":8080", r))
}
