// Package api é o contrato público do serviço de anexos.
package api

import (
	"context"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
)

type Store interface {
	// NewUploadURL emite URL temporária para envio de um anexo já cifrado no cliente.
	NewUploadURL(ctx context.Context, room ids.RoomID, dev ids.DeviceID, size int64) (string, error)
}
