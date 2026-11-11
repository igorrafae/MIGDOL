// Package delivery implementa api.Publisher. Esqueleto da Fase 0.
package delivery

import (
	"context"
	"errors"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
	auditapi "github.com/igorrafae/MIGDOL/server/services/audit/api"
	"github.com/igorrafae/MIGDOL/server/services/delivery/api"
	dirapi "github.com/igorrafae/MIGDOL/server/services/directory/api"
)

var ErrForbidden = errors.New("delivery: não é membro da sala")

type Service struct {
	dir   dirapi.Directory
	audit auditapi.Recorder
}

func New(dir dirapi.Directory, audit auditapi.Recorder) *Service {
	return &Service{dir: dir, audit: audit}
}

var _ api.Publisher = (*Service)(nil)

// Publish verifica autorização por sala antes de qualquer outra coisa (regra 5).
func (s *Service) Publish(ctx context.Context, from ids.DeviceID, env api.Envelope) error {
	ok, err := s.dir.IsMember(ctx, env.Room, from)
	if err != nil || !ok {
		return ErrForbidden
	}
	return errors.New("delivery: fila não implementada")
}
