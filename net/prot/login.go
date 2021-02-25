package prot

import "github.com/jw2476/cubelet/client"

type LoginStart struct {
	Client *client.Client
	Username string
}

func (l LoginStart) GetName() string {
	return "login start"
}

func DecodeLoginStart(c *client.Client) (LoginStart, error) {
	username, err := c.ReadString()
	if err != nil {
		return LoginStart{}, err
	}

	return LoginStart{
		Client: c,
		Username: username,
	}, nil
}
