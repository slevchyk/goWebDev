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
		bs := []byte(ln)
		encr := Rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, encr)
	}

}

//Rot13 encrypt by byte shifting
func Rot13(bs []byte) []byte {

	var r13 = make([]byte, len(bs))

	for i, val := range bs {
		if val < 109 {
			r13[i] = val + 13
		} else {
			r13[i] = val - 13
		}
	}

	return r13
}
