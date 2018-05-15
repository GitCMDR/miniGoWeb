package main

import (
	"net"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}