package main

import (
	"log"
	"net"
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

		go serve(conn)
	}
	
}
func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			fmt.Println("no more lines to scan")
			break
		}

	}

	fmt.Fprintf(conn, "writing to the response")
}