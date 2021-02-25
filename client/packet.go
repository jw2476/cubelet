package client

type Packet struct {
	buffer []byte
}


func NewPacket(opcode int) Packet {
	packet := Packet{buffer: make([]byte, 0)}
	packet.WriteVarInt(opcode)
	return packet
}

func (p *Packet) WriteVarInt(number int) {
	for {
		value := number & 0b01111111
		number >>= 7

		if number != 0 {
			value |= 0b10000000
		}
		p.buffer = append(p.buffer, byte(value))

		if number == 0 {
			return
		}
	}
}

func (p *Packet) writeBytes(buf []byte) {
	p.buffer = append(p.buffer, buf...)
}

func (p *Packet) WriteString(value string) {
	p.WriteVarInt(len(value))
	p.writeBytes([]byte(value))
}

func (p *Packet) finish() []byte {
	packet := NewPacket(len(p.buffer))
	packet.writeBytes(p.buffer)
	return packet.buffer
}