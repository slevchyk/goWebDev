package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)

	}
}
func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Connction closed. Timed out")
	}

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Fprintln(conn, ln)
	}
	defer conn.Close()

	fmt.Println("handle end")
}
