package main

import "github.com/jw2476/cubelet/net/tcp"

func main() {
	go tcp.Listen(25565, tcp.HandleConn)

	select {}
}
