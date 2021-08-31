package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ismferd/k8s-port-sniffer/src/aws"
	"github.com/ismferd/k8s-port-sniffer/src/consumer"
	"github.com/ismferd/k8s-port-sniffer/src/producer"
	"golang.org/x/sync/semaphore"
)

func main() {

	host := os.Getenv("HOST")
	iterationTime, err := strconv.Atoi(os.Getenv("ITERATION_TIME"))
	if err != nil {
		log.Print("Iteration time needs to be an integer")
		os.Exit(1)
	}
	region := os.Getenv("AWS_REGION")
	s3Endpoint := os.Getenv("AWS_ENDPOINT")
	bucketName := os.Getenv("BUCKET_NAME")
	s3Session := aws.S3Session(region, s3Endpoint)
	portsWhitelisted := strings.Split(os.Getenv("PORTS"), ",")

	// Controlling input envars
	if iterationTime < 60 {
		log.Print("The loop time must be higher than 60")
		os.Exit(1)
	}

	if region == "" {
		log.Print("Var region must be informed")
		os.Exit(1)
	}

	if s3Endpoint == "" {
		log.Print("Var s3Endpoint must be informed. eg: https://s3.us-east-2.amazonaws.com")
		os.Exit(1)
	}

	if bucketName == "" {
		log.Print("Var bucketName must be informed")
		os.Exit(1)
	}
	// Showing the current config as STDOUT
	log.Printf("Current Configuration: \n iterationTime %d, region: %s, s3Endpoint: %s bucketName: %s", iterationTime, region, s3Endpoint, bucketName)

	// Making a s3 object
	s3cli := aws.NewS3Cli(bucketName, s3Session, region)

	// Go routine with the listener at 8080
	go consumer.Consumer(s3cli)
	node := make(map[string][]int)
	// Main loop
	for {
		log.Printf("Executing")
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Making a new producer to scan ports
		ps := producer.NewPortScanner(hostname, host, semaphore.NewWeighted(1048576), portsWhitelisted)
		// Starting to scan ports
		node[hostname] = ps.Start(1, 65535, 500*time.Millisecond)
		go s3cli.PutObjectToS3(producer.MapToString(node))
		time.Sleep(time.Duration(iterationTime) * time.Second)
	}
}
