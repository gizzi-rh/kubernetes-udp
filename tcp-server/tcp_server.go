package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func sendResponse(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	conn.Close()
}

func main() {
	port := "10080"

	fmt.Println("Listening on", port)

	/* Hostname */
	hostname, err := os.Hostname()
	fmt.Println("Hostname=" + hostname)
	checkError(err)

	/* Lets prepare a address at any address at port 10001*/
	//serverAddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	serverAddr := ":" + port
	checkError(err)

	/* Now listen at selected port */
	listener, err := net.Listen("tcp", serverAddr)
	checkError(err)
	defer listener.Close()

	msg := []byte(hostname)
	//buf := make([]byte, 2048)

	for {
		tcpconn, err := listener.Accept()
		checkError(err)

		reqString, err := bufio.NewReader(tcpconn).ReadString('\n')
		checkError(err)

		fmt.Println("Received ", reqString, " from ", tcpconn.RemoteAddr().String())

		go sendResponse(tcpconn, msg)

	}
}
