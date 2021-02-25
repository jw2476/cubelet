package tcp

import (
	"fmt"
	"github.com/jw2476/cubelet/client"
	"github.com/jw2476/cubelet/event"
	"github.com/jw2476/cubelet/net/prot"
	"net"
)

func HandleConn(conn net.Conn) {
	c := client.NewClient(conn)

	for {
		length, err := c.ReadVarIntLive()
		if err != nil {
			fmt.Println(err)
			return
		}

		buf, err := c.ReadBytesLive(length)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.AddBytes(buf)

		opcode, err := c.ReadVarInt()
		if err != nil {
			fmt.Println(err)
			return
		}

		var packet event.Event

		switch c.GetState() {
		case 0: packet, err = prot.DecodeHandshake(&c)
		case 1: {
			switch opcode {
			case 0: packet = prot.DecodeRequest(&c)

			}
		}
		case 2: {
			switch opcode {
			case 0: packet, err = prot.DecodeLoginStart(&c)
			}
		}
		default: {
			fmt.Println("Invalid State", c.GetState())
			return
		}
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%T { %+v }\n", packet, packet)
		event.PacketEventBus.Publish(packet.GetName(), packet)

		c.ResetBuffer()
	}
}
