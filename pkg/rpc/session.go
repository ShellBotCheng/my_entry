package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

const headerLen = 4

type Session struct {
	conn net.Conn
}

func NewSession(conn net.Conn) *Session {
	return &Session{conn}
}

func (t *Session) Send(data []byte) error {
	buf := make([]byte, headerLen+len(data))
	binary.BigEndian.PutUint32(buf[:headerLen], uint32(len(data)))
	copy(buf[headerLen:], data)
	_, err := t.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (t *Session) Read() ([]byte, error) {
	header := make([]byte, headerLen)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(t.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
