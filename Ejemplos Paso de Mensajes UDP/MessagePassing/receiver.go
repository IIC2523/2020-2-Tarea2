package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

const (
	address = "127.0.0.1:9999"
)

const (
	maxDatagramSize = 8192
)

func main() {
	// Set Address
	fmt.Println("Set Address")
	addr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	// Open up a connection
	fmt.Println("Open Connection")
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(conn, err)
	}
	defer conn.Close()

	fmt.Println("Listening")
	for {
		buffer := make([]byte, maxDatagramSize)
		numBytes, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		log.Println(hex.Dump(buffer[:numBytes]))
	}
}
