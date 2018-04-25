package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

		go serve(conn)
	}

}
func serve(conn net.Conn) {
	defer conn.Close()

	do := true
	var method, URI string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if do {

			sls := strings.Fields(ln)
			method = sls[0]
			URI = sls[1]

			fmt.Println("method:", method)
			fmt.Println("URI:", URI)

			do = false
		}

		if ln == "" {
			fmt.Println("no more lines to scan")
			break
		}

	}

	var body string

	if method == "GET" && URI == "/" {
		body = "<h1>GET INDEX</h1>"
	} else if method == "GET" && URI == "/apply" {
		body = "<h1>GET APPLY</h1>"
	} else if method == "POST" && URI == "/apply" {
		body = "<h1>POST APPLY</h1>"
	} else {
		body = "<h1>PAGE NOT FOUND</h1>"
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
