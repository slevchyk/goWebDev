package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	request(conn)
	response(conn)
}

func response(conn net.Conn) {

	bodyText := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Title</title>
		</head>
		<body>
		<h1>Hello world</h1>
		</body>
		</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(bodyText))
	fmt.Fprintf(conn, "Content-Type:  text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, bodyText)

}
func request(conn net.Conn) {

	do := true

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)

		if do {
			x := strings.Fields(ln)[0]
			fmt.Printf(">>>>>>Method:%s\n", x)
			do = false
		}

		if ln == "" {
			break
		}
	}
}
