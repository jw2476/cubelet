package tcp

import (
	"fmt"
	"net"
)

func Listen(port uint16, callback func(conn net.Conn)) {
	listener, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port", port)

	for {
		conn, err := listener.Accept()
		fmt.Println("Connection Received")
		if err != nil {
			fmt.Println(err)
			return
		}

		go callback(conn)
	}
}
