package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
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
		ext, err := net.DialTimeout("tcp", os.Args[1], 1*time.Second)
		if err != nil {
			fmt.Println("Could not connect to remote server : ", err)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("Connection to :", os.Args[1], " Successful")

		internal, err := net.DialTimeout("tcp", os.Args[2], 1*time.Second)
		if err != nil {
			fmt.Println("Could not connect to remote server : ", err)
			return
		}
		handleConn(ext, internal)
	}
}

func handleConn(in net.Conn, out net.Conn) {
	go io.Copy(out, in)
	go io.Copy(in, out)
}
