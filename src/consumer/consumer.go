package consumer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ismferd/k8s-port-sniffer/src/aws"
)

func Consumer(s3Cli *aws.S3Cli) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s3Values := s3Cli.ListObjectFromS3()
		fmt.Fprintf(w, "%s\n%s", time.Now(), s3Values)
	})
	http.ListenAndServe(":8080", nil)
}
