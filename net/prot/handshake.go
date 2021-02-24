package prot

import (
	"github.com/jw2476/cubelet/client"
)

type Handshake struct {
	protocolVersion int
	serverAddress string
	serverPort uint16
	nextState int
}

func DecodeHandshake(c client.Client) (Handshake, error) {
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
		protocolVersion: protocolVersion,
		serverAddress:   serverAddress,
		serverPort:      serverPort,
		nextState:       nextState,
	}, nil
}
