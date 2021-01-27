package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("tcp client is running")
	times := 1
	for {
		msg := fmt.Sprintf("i love you %d times\n", times)
		fmt.Fprintf(conn, msg)

		log.Printf("send: %s", msg)
		reply, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("receive: %s\n", reply)
		times++
		time.Sleep(time.Second)
	}
}
