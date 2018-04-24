package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("handle end")

}
