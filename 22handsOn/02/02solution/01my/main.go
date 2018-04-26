package main

import (
	"io"
	"log"
	"net"
)

func main() {

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

	io.WriteString(conn, "I see you conected")
	conn.Close()
}
