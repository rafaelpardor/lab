package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide a port number")
		return
	}
	PORT := ":" + args[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()
}
