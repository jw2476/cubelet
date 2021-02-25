package event

import "github.com/asaskevich/EventBus"

var PacketEventBus EventBus.Bus

func MakePacketEventBus() {
	PacketEventBus = EventBus.New()
}
