package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

// TODO: get a port from input
var addList [2]string = [2]string{"127.0.0.1:9999", "127.0.0.1:9998"}

const (
	maxDatagramSize = 8192
)

func send(address string, count *int) {
	addr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	// Dial connects to the address on the named network.
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		time.Sleep(2 * time.Second)
		*count++
		conn.Write([]byte("Hello " + address + ", " + strconv.Itoa(*count)))
		fmt.Println("Sent Message to " + address + ", " + strconv.Itoa(*count))
	}
}

func main() {
	count := 0
	for _, addr := range addList {
		go send(addr, &count)
	}

	for {
	}
}
