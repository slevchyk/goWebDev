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

	var url string

	request(conn, &url)
	response(conn, url)
}

func request(conn net.Conn, url *string) {

	do := true
	var path string

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)

		if do && ln != "" {
			fs := strings.Fields(ln)

			switch fs[0] {
			case "GET":
				path = fs[1]
				fmt.Println(path)
			case "Host:":
				*url = fs[1] + path
				fmt.Printf(">>>>>host: %s; path: %s", fs[1], path)
				do = false
			}
		}

		if ln == "" {
			break
		}
	}
}

func response(conn net.Conn, url string) {

	bodyText := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Title</title>
		</head>
		<body>` +
		"<a href=" + url + ">" + url + "</>" +
		`</body>
		</html>`

	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(bodyText))
	fmt.Fprintf(conn, "Content-Type:  text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, bodyText)

}
