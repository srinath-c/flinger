package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give a port to listen (with ':' prefix) as  arguments")
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("Error not enough arguments")
		return
	}

	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		fmt.Println("Could not listen in the given port :", os.Args[1])
		return
	}
	con1, err := listen.Accept()
	if err != nil {
		fmt.Println("Could not accept connection :", err)
		return
	}
	con2, err := listen.Accept()
	if err != nil {
		fmt.Println("Could not accept connection :", err)
		return
	}
	handleConn(con1, con2)

}

func handleConn(in net.Conn, out net.Conn) {
	go io.Copy(out, in)
	io.Copy(in, out)
}
