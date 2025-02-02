package client

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/emikhalev/faraway_wow/pkg/utils/pow"
)

func WoW(ctx context.Context, conn net.Conn) (string, error) {
	difficultyByte := make([]byte, 1)
	_, err := io.ReadFull(conn, difficultyByte)
	if err != nil {
		return "", fmt.Errorf("error reading difficulty byte: %w", err)
	}

	challenge := make([]byte, 32)
	_, err = io.ReadFull(conn, challenge)
	if err != nil {
		return "", fmt.Errorf("error reading challenge: %w", err)
	}

	nonce := pow.Solve(challenge, int(difficultyByte[0]))
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, nonce)
	_, err = conn.Write(buf)
	if err != nil {
		return "", fmt.Errorf("error sending solution: %w", err)
	}

	quote, err := getText(conn)
	if err != nil {
		return "", fmt.Errorf("error get quote: %w", err)
	}

	return quote, nil
}

func getText(conn net.Conn) (string, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(conn, lengthBuf)
	if err != nil {
		return "", fmt.Errorf("error reading text length: %w", err)
	}
	quoteLength := binary.BigEndian.Uint32(lengthBuf)

	quoteBuffer := make([]byte, quoteLength)
	_, err = io.ReadFull(conn, quoteBuffer)
	if err != nil {
		return "", fmt.Errorf("error reading text: %w", err)
	}

	return string(quoteBuffer), nil
}
