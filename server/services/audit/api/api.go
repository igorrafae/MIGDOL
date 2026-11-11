// Package api é o contrato público do serviço de auditoria.
// Hoje implementado em processo; na extração, ganha um cliente HTTP/gRPC com a mesma interface.
package api

import (
	"context"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
)

// Event é um evento de metadados. Nunca contém conteúdo de mensagem.
type Event struct {
	Actor ids.DeviceID
	Name  string
	Room  ids.RoomID
}

type Recorder interface {
	Record(ctx context.Context, e Event) error
}
