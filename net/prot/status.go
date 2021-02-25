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
