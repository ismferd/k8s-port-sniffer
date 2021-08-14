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
	iterationTime, _ := strconv.Atoi(os.Getenv("ITERATION_TIME"))
	region := os.Getenv("AWS_REGION")
	s3Endpoint := os.Getenv("AWS_ENDPOINT")
	bucketName := os.Getenv("BUCKET_NAME")
	s3Session := aws.S3Session(region, s3Endpoint)
	portsWhitelisted := strings.Split(os.Getenv("PORTS"), ",")

	s3cli := aws.NewS3Cli(bucketName, s3Session, region)

	go consumer.Consumer(s3cli)
	node := make(map[string][]int)
	for {
		time.Sleep(time.Duration(iterationTime) * time.Second)
		log.Print("executing")

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ps := producer.NewPortScanner(hostname, host, semaphore.NewWeighted(1048576), portsWhitelisted)
		//lock: semaphore.NewWeighted(Ulimit()),
		node[hostname] = ps.Start(1, 65535, 500*time.Millisecond)
		s3cli.PutObjectToS3(producer.MapToString(node))
	}
}
