// Package gateway é a borda: mTLS 1.3 e WebSocket. Não expõe api; só consome contratos.
// Esqueleto da Fase 0.
package gateway

import (
	deliveryapi "github.com/igorrafae/MIGDOL/server/services/delivery/api"
	dirapi "github.com/igorrafae/MIGDOL/server/services/directory/api"
)

type Gateway struct {
	Directory dirapi.Directory
	Publisher deliveryapi.Publisher
}

func New(dir dirapi.Directory, pub deliveryapi.Publisher) *Gateway {
	return &Gateway{Directory: dir, Publisher: pub}
}
