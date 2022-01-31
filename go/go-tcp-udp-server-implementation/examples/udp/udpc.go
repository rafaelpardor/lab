package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Provide host:port")
	}

	CONNECT := args[1]
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String)
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString("\n")
		data := []byte(text)
		_, err = c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("ADIOS")
			return
		}

		if err != nil {
			fmt.Prinln(err)
			return
		}
		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
