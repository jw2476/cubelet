package main

import (
	"github.com/jw2476/cubelet/event"
	"github.com/jw2476/cubelet/net/handlers"
	"github.com/jw2476/cubelet/net/tcp"
	"log"
)

func startup() {
	event.MakePacketEventBus()
	err := handlers.Register()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	startup()
	go tcp.Listen(25565, tcp.HandleConn)

	select {}
}
