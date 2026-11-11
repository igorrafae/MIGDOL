// Package blobs implementa api.Store. Esqueleto da Fase 0.
package blobs

import (
	"context"
	"errors"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
	"github.com/igorrafae/MIGDOL/server/services/blobs/api"
	dirapi "github.com/igorrafae/MIGDOL/server/services/directory/api"
)

type Service struct{ dir dirapi.Directory }

func New(dir dirapi.Directory) *Service { return &Service{dir: dir} }

var _ api.Store = (*Service)(nil)

func (s *Service) NewUploadURL(ctx context.Context, room ids.RoomID, dev ids.DeviceID, size int64) (string, error) {
	if ok, err := s.dir.IsMember(ctx, room, dev); err != nil || !ok {
		return "", errors.New("blobs: acesso negado")
	}
	return "", errors.New("blobs: não implementado")
}
