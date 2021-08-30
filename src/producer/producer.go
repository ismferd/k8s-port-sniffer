package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/semaphore"
)

// PortScanner definition
type PortScanner struct {
	Hostname         string
	Ip               string
	Lock             *semaphore.Weighted
	WhitelistedPorts []string
}

func NewPortScanner(hostname string, ip string, lock *semaphore.Weighted, whitelistedPorts []string) *PortScanner {
	return &PortScanner{Hostname: hostname, Ip: ip, Lock: lock, WhitelistedPorts: whitelistedPorts}
}

// PortScanner method, where it returns ports opened or 0. It takes care about whitelisting ports.
func (ps *PortScanner) ScanPort(ip string, port int, timeout time.Duration) int {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ps.ScanPort(ip, port, timeout)
		}
		return 0
	}

	conn.Close()
	fmt.Println(port, "open")
	for _, p := range ps.WhitelistedPorts {
		intPort, _ := strconv.Atoi(p)
		if intPort == port {
			return 0
		}
	}
	return port
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) []int {
	var portList []int
	scanPort := 0
	for port := f; port <= l; port++ {
		ps.Lock.Acquire(context.TODO(), 1)
		scanPort = ps.ScanPort(ps.Ip, port, timeout)
		if scanPort != 0 {
			portList = append(portList, scanPort)

		}
	}
	return portList
}

// Convert from map to string
func MapToString(node map[string][]int) string {
	mJson, err := json.Marshal(node)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(mJson)
}
