// Package api é o contrato público do serviço de entrega.
package api

import (
	"context"
	"time"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
)

// Envelope é tudo o que o servidor enxerga de uma mensagem (docs/02 5.3).
type Envelope struct {
	Room       ids.RoomID
	Epoch      uint64
	Ciphertext []byte
	ExpiresAt  time.Time
}

type Publisher interface {
	Publish(ctx context.Context, from ids.DeviceID, env Envelope) error
}
