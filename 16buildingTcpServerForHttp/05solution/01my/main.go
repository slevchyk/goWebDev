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
}

func request(conn net.Conn) {

	do := true
	var path string

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)

		if do && ln != "" {
			fs := strings.Fields(ln)

			handle := fs[1]

			switch handle {
			case "/":
				response(conn, "INDEX PAGE")
			case "/home":
				response(conn, "home")
			default:
				response(conn, "hm, can't find "+path)
			}

			do = false
		}

		if ln == "" {
			break
		}
	}
}

func response(conn net.Conn, text string) {

	bodyText := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Title</title>
		</head>
		<body>` +
		text +
		`</body>
		</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(bodyText))
	fmt.Fprintf(conn, "Content-Type:  text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, bodyText)

}
