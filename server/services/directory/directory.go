// Package directory implementa api.Directory. Esqueleto da Fase 0: nega tudo por padrão.
package directory

import (
	"context"

	"github.com/igorrafae/MIGDOL/server/platform/ids"
	"github.com/igorrafae/MIGDOL/server/services/audit/api"
	dirapi "github.com/igorrafae/MIGDOL/server/services/directory/api"
)

type Service struct{ audit api.Recorder }

func New(audit api.Recorder) *Service { return &Service{audit: audit} }

var _ dirapi.Directory = (*Service)(nil)

func (*Service) IsMember(context.Context, ids.RoomID, ids.DeviceID) (bool, error) { return false, nil }
func (*Service) DeviceActive(context.Context, ids.DeviceID) (bool, error)          { return false, nil }
