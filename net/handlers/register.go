package handlers

import "github.com/jw2476/cubelet/event"

func Register() error {
	err := event.PacketEventBus.Subscribe("handshake", handleHandshake)
	if err != nil {
		return err
	}

	err = event.PacketEventBus.Subscribe("request", handleRequest)
	if err != nil {
		return err
	}

	return nil
}

