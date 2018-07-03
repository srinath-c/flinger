package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give 2 ports to listen (with ':' prefix) as  arguments")
		return
	}
	if len(os.Args) != 3 {
		fmt.Println("Error not enough arguments")
		return
	}

	portA, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		fmt.Println("Could not listen in the given port :", os.Args[1])
		return
	}
	fmt.Println("Listening on port : ", os.Args[1])
	portB, err := net.Listen("tcp", os.Args[2])
	if err != nil {
		fmt.Println("Could not listen in the given port :", os.Args[2])

	}
	fmt.Println("Listening on port : ", os.Args[2])
	ever := true
	for ever {
		con1, err := portA.Accept()
		if err != nil {
			fmt.Println("Could not accept connection :", err)
			continue
		}
		fmt.Println("Con accepted on port : ", os.Args[1], "from :", con1.RemoteAddr())
		con2, err := portB.Accept()
		if err != nil {
			fmt.Println("Could not accept connection :", err)
			return
		}
		fmt.Println("Con accepted on port : ", os.Args[2], "from :", con2.RemoteAddr())
		handleConn(con1, con2)
		continue
	}
}

func handleConn(in net.Conn, out net.Conn) {
	go io.Copy(out, in)
	go io.Copy(in, out)
}
