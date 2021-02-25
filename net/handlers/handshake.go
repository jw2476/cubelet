package handlers

import (
	"github.com/jw2476/cubelet/net/prot"
)

func handleHandshake(handshake prot.Handshake) {
	handshake.Client.SetState(handshake.NextState)
}
