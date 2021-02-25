package client

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	buffer *bytes.Buffer
	state int
}

func NewClient(conn net.Conn) Client {
	return Client{
		conn: conn,
		buffer: new(bytes.Buffer),
		state: 0,
	}
}

func (c *Client) ReadVarIntLive() (int, error) {
	result, i := 0, 0

	for {
		buf := make([]byte, 1)
		n, err := c.conn.Read(buf)
		if err != nil {
			return 0, err
		}
		if n == 0 {
			return 0, errors.New("EOF")
		}

		read := buf[0]
		value := read & 0b01111111
		result |= int(value) << (i * 7)
		i++

		if read & 0b10000000 == 0 {
			return result, nil
		}
	}
}

func (c *Client) ReadVarInt() (int, error) {
	result, i := 0, 0

	for {
		read, err := c.buffer.ReadByte()
		if err != nil {
			return 0, err
		}

		value := read & 0b01111111
		result |= int(value) << (i * 7)
		i++

		if read & 0b10000000 == 0 {
			return result, nil
		}
	}
}

func (c *Client) ReadBytesLive(length int) ([]byte, error) {
	buf := make([]byte, length)
	read := 0

	for read < length {
		n, err := c.conn.Read(buf[read:])
		if err != nil {
			return buf, err
		}
		if n == 0 {
			return buf, errors.New("EOF")
		}

		read += n
	}

	return buf, nil
}

func (c *Client) AddBytes(buf []byte) {
	c.buffer.Write(buf)
}

func (c *Client) ResetBuffer() {
	c.buffer.Reset()
}

func (c *Client) ReadBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	n, err := c.buffer.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != length {
		return buf, errors.New("not enough bytes")
	}

	return buf, nil
}

func (c *Client) ReadString() (string, error) {
	length, err := c.ReadVarInt()
	if err != nil {
		return "", err
	}

	buf, err := c.ReadBytes(length)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *Client) ReadUnsignedShort() (uint16, error) {
	buf, err := c.ReadBytes(2)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(buf), nil
}

func (c Client) GetState() int {
	return c.state
}

func (c *Client) SetState(state int) {
	c.state = state
}

func (c *Client) Send(packet Packet) error {
	fmt.Println(packet.finish())
	_, err := c.conn.Write(packet.finish())
	return err
}