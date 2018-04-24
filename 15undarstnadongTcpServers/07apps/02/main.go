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

	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln()
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

	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n"+
		"USE:\n"+
		"SET key balue\n"+
		"GET key\n"+
		"DEL key\n"+
		"SHOWALL\n"+
		"EXAMPLE:\n\" +"+
		"SET cat Bonny"+
		"GET cat\n\n\n")

	//read & write

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		//command is a first word in conn message. utils command
		command := fs[0]
		switch command {
		case "GET":
			//keyDB is a key of DB
			keyDB := fs[1]
			//valDB is  a value of key in DB
			valDB := data[keyDB]
			fmt.Fprintf(conn, "%s\n", valDB)
		case "SET":
			//we shoud have 3 words COMMAND KEY VALUE
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE\n")
				continue
			}
			keyDB := fs[1]
			valDB := fs[2]
			data[keyDB] = valDB
		case "DEL":
			keyDB := fs[1]
			delete(data, keyDB)
		case "SHOWALL":
			if len(data) > 0 {
				for key, val := range data {
					fmt.Fprintf(conn, "%s - %s\n", key, val)
				}
			} else {
				fmt.Fprintf(conn, "DB is emty. Nothing to show\n")
			}
		default:
			fmt.Fprintf(conn, "INVALID COMMAND\n")
		}
	}

}
