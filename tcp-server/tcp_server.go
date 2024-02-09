package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func handlingRequest(conn net.Conn, msg []byte) {

	for {
		reqString, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println("Received ", reqString, " from ", conn.RemoteAddr().String())

		sendResponse(conn, msg)

	}

}

func sendResponse(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func main() {
	port := "10080"

	fmt.Println("Listening on", port)

	/* Hostname */
	hostname, err := os.Hostname()
	fmt.Println("Hostname=" + hostname)
	checkError(err)

	/* Lets prepare a address at any address at port 10080*/
	serverAddr := ":" + port
	checkError(err)

	/* Now listen at selected port */
	listener, err := net.Listen("tcp", serverAddr)
	checkError(err)
	defer listener.Close()

	msg := []byte(hostname)

	for {
		tcpconn, err := listener.Accept()
		checkError(err)
		defer tcpconn.Close()

		go handlingRequest(tcpconn, msg)

	}
}
