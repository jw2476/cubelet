package prot

import "github.com/jw2476/cubelet/client"

type Request struct {
	Client *client.Client
}

func (r Request) GetName() string {
	return "request"
}

func DecodeRequest(c *client.Client) Request {
	return Request{Client: c}
}

type Ping struct {
	Client *client.Client
	Payload uint64
}

func (p Ping) GetName() string {
	return "ping"
}

func DecodePing(c *client.Client) (Ping, error) {
	payload, err := c.ReadUnsignedLong()
	if err != nil {
		return Ping{}, err
	}

	return Ping{
		Payload: payload,
		Client: c,
	}, nil
}

