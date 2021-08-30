package consumer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ismferd/k8s-port-sniffer/src/aws"
)

// Create a Listener in 8080 and call to ListObjectFromS3 to show in HTTP the results
func Consumer(s3Cli *aws.S3Cli) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s3Values := s3Cli.ListObjectFromS3()
		fmt.Fprintf(w, "%s\n%s", time.Now(), s3Values)
	})
	http.ListenAndServe(":8080", nil)
}
