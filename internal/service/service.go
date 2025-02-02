package handlers

import (
	"encoding/binary"
	"net"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) SendText(conn net.Conn, text string) error {
	quoteLen := uint32(len(text))

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, quoteLen)

	_, err := conn.Write(lengthBuf)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(text))
	if err != nil {
		return err
	}
	return nil
}
