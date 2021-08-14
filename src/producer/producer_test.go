package producer

import (
	"log"
	"net"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestScanPort(t *testing.T) {

	ps := NewPortScanner("hostname", "host", semaphore.NewWeighted(1048576), []string{"50"})
	output := ps.ScanPort("ip", 0, 500*time.Millisecond)
	expectedOutput := 0
	if output == expectedOutput {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed ! got %v want %c", output, expectedOutput)
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	output = ps.ScanPort("127.0.0.1", 8080, 500*time.Millisecond)
	expectedOutput = 8080
	if output == expectedOutput {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	}
	l.Close()

}

func TestMapToString(t *testing.T) {
	node := make(map[string][]int)
	node["foo"] = []int{1, 2, 3}
	expectedOutput := "{\"foo\":[1,2,3]}"
	output := MapToString(node)
	if output == expectedOutput {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	}
}
