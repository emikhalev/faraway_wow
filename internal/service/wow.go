package handlers

import (
	"context"
	"math/rand"
	"net"
)

var (
	quotes = []string{
		"Do what you can, with what you have, where you are. — Theodore Roosevelt",
		"The secret to success is to start. — Mark Twain",
		"The most important thing in life is not what you did, but how you did it. — Leonardo da Vinci",
		"If you want to live a happy life, tie it to a goal, not to people or things. — Albert Einstein",
		"Be yourself; everyone else is already taken. — Oscar Wilde",
		"It’s never too late to be what you might have been. — George Eliot",
		"Problems are not something we need to avoid. They are opportunities we should use. — Matt Besley",
		"Strength is not in never falling, but in getting up every time. — Confucius",
		"Life is what happens while you’re busy making other plans. — John Lennon",
		"He who is able to rejoice in the little things will never feel poor. — Antoine de Saint-Exupéry",
	}
)

func (s *Service) WoWHandler(ctx context.Context, conn net.Conn) error {
	return s.SendText(conn, quotes[rand.Intn(len(quotes))])
}
