// Package api é o contrato público do serviço de diretório.
package api

import (
	"context"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
)

type Directory interface {
	// IsMember informa se o aparelho pertence à sala. Negar por padrão em caso de erro.
	IsMember(ctx context.Context, room ids.RoomID, dev ids.DeviceID) (bool, error)
	// DeviceActive informa se o aparelho está ativo e com lease válido.
	DeviceActive(ctx context.Context, dev ids.DeviceID) (bool, error)
}
