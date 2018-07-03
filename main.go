package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("External server to forward and internal ip with port(for both) as arguments")
		return
	}
	if len(os.Args) != 3 {
		fmt.Println("Error not enough arguments")
		return
	}
	ever := true
	for ever {
		ext, err := net.Dial("tcp", os.Args[1])
		if err != nil {
			fmt.Println("Could not connect to remote server : ", err)
			continue
		}
		fmt.Println("Connection to :", os.Args[1], " Successful")
		internal, err := net.Dial("tcp", os.Args[2])
		if err != nil {
			fmt.Println("Could not connect to remote server : ", err)
			return
		}
		handleConn(ext, internal)
		continue
	}
}

func handleConn(in net.Conn, out net.Conn) {
	go io.Copy(out, in)
	io.Copy(in, out)
}
