// Package audit implementa api.Recorder. Esqueleto da Fase 0: ainda não grava nada.
package audit

import (
	"context"
	"errors"

	"github.com/igorrafae/MIGDOL/server/services/audit/api"
)

var ErrNotImplemented = errors.New("audit: não implementado")

type Service struct{}

func New() *Service { return &Service{} }

var _ api.Recorder = (*Service)(nil)

func (*Service) Record(context.Context, api.Event) error { return ErrNotImplemented }
