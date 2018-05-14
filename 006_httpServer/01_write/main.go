package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close() // Defer this instruction to the last part of execution

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello! This is Luis Server \n")
		fmt.Fprintln(conn, "This is my first Go server")
		fmt.Fprintf(conn, "%v", "How cool is it!\n")

		conn.Close()
	}

}
