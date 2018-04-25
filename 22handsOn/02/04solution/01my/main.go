package main

import (
	"log"
	"net"
	"io"
	"bufio"
	"fmt"
)

func main()  {
	
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handle(conn)
	}
	
}
func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	fmt.Println("scanning ends")

	io.WriteString(conn, "I see you conected")
	conn.Close()
}