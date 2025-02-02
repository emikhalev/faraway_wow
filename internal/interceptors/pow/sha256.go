package pow

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/emikhalev/faraway_wow/internal/config"
	"github.com/emikhalev/faraway_wow/internal/logger"
	"github.com/emikhalev/faraway_wow/internal/server"
	"github.com/emikhalev/faraway_wow/pkg/utils/pow"
)

func Sha256(ctx context.Context, conn net.Conn, handler server.Handler) error {
	tokenSize := config.Get(ctx).PoW.TokenSize
	difficulty := byte(config.Get(ctx).PoW.Difficulty)

	challenge, err := pow.Challenge(tokenSize)
	if err != nil {
		return fmt.Errorf("error generating challenge: %w", err)
	}

	_, err = conn.Write(append([]byte{byte(config.Get(ctx).PoW.Difficulty)}, challenge...))
	if err != nil {
		return fmt.Errorf("error sending challenge: %w", err)
	}

	buf := make([]byte, 8)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return fmt.Errorf("error reading nonce: %w", err)
	}
	nonce := binary.BigEndian.Uint64(buf)

	if !pow.Verify(challenge, nonce, int(difficulty)) {
		return fmt.Errorf("PoW not accepted")
	}

	logger.Infof(ctx, "PoW accepted.")
	return handler(ctx, conn)
}
