package prot

import (
	"github.com/jw2476/cubelet/client"
)

type Handshake struct {
	Client *client.Client
	ProtocolVersion int
	ServerAddress string
	ServerPort uint16
	NextState int
}

func DecodeHandshake(c *client.Client) (Handshake, error) {
	protocolVersion, err := c.ReadVarInt()
	if err != nil {
		return Handshake{}, err
	}

	serverAddress, err := c.ReadString()
	if err != nil {
		return Handshake{}, err
	}

	serverPort, err := c.ReadUnsignedShort()
	if err != nil {
		return Handshake{}, err
	}

	nextState, err := c.ReadVarInt()
	if err != nil {
		return Handshake{}, err
	}

	return Handshake{
		Client: c,
		ProtocolVersion: protocolVersion,
		ServerAddress:   serverAddress,
		ServerPort:      serverPort,
		NextState:       nextState,
	}, nil
}

func (h Handshake) GetName() string {
	return "handshake"
}
