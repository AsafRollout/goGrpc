package speakers

import (
	"context"
	"log"
)

type Server struct {
	UnsafeAudioOutputServer
}

func (b *Server) Play(ctx context.Context, audio *Audio) (*Audio, error) {
	log.Printf("got message from client: %v\n", audio.Body)
	return &Audio{Body: "lalala"}, nil
}
