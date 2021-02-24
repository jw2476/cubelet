package tcp

import (
	"fmt"
	"github.com/jw2476/cubelet/client"
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

		_, err = c.ReadVarInt()
		if err != nil {
			fmt.Println(err)
			return
		}

		var packet interface{}

		switch c.GetState() {
		case 0: packet, err = prot.DecodeHandshake(c)

		}

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%T { %+v }\n", packet, packet)

		c.ResetBuffer()
	}
}
