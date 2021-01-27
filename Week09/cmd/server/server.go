package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	messageChan := make(chan string, 32)
	go handleWrite(conn, messageChan)
	for {
		receiveMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			log.Println("Read Finish")
			break
		} else if err != nil {
			log.Println("Read Error: ", err)
			break
		}
		log.Printf("Receive Message: %s", receiveMessage)
		messageChan <- receiveMessage
	}
	close(messageChan)
}

func handleWrite(conn net.Conn, messageChan chan string) {
	writer := bufio.NewWriter(conn)
	for message := range messageChan {
		message = fmt.Sprintf("%s, too\n", message[:len(message)-1])
		writer.Write([]byte(message))
		writer.Flush()
	}
	log.Println("Writer Done")
	log.Println("Number of active goroutines ", runtime.NumGoroutine())
}

func main() {
	port := "9000"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	log.Println("tcp server running on: ", port)
	for {
		conn, err := listen.Accept()
		log.Println("Accept TCP request from ", conn.RemoteAddr().String())
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}
