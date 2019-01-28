package main

import (
	"os"
	"strconv"
	"fmt"
	"net"
)

func main() {
	var port = "8080"
	if (len(os.Args) > 1) {
		port = os.Args[1]
		intport, err := strconv.Atoi(port)
		if (err != nil || intport > 65535 || intport < 0) {
			fmt.Printf("Invalid port: %q\n", port)
			os.Exit(1)
		}
	}
	sock, err := net.Listen("tcp", ":" + port)
	fmt.Println("Attempting to listen on port " + port)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for connections on port " + port + " ...")
	for {
		conn, err := sock.Accept();
		if err != nil {
			panic(err)
		}
		ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())
		if err != nil {
			panic(err)
		}
		fmt.Println(ip)
		conn.Close()
	}
}
