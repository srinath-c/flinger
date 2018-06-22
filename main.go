package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give a port to listen (with ':' prefix) and remote ip with port as arguments")
		return
	}
	if len(os.Args) != 3 {
		fmt.Println("Error not enough arguments")
		return
	}

	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	for {
		out, err := net.Dial("tcp", os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		in, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		handleConn(in, out)
		defer out.Close()
		defer in.Close()
	}

}

func handleConn(in net.Conn, out net.Conn) {
	go io.Copy(out, in)
	io.Copy(in, out)
}
