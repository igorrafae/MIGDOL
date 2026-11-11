// Composition root da v1: liga os serviços em processo, passando contratos (api), nunca implementações.
// Para extrair um serviço, troque aqui a implementação local por um cliente remoto da mesma interface.
package main

import (
	"github.com/igorrafae/MIGDOL/server/services/audit"
	"github.com/igorrafae/MIGDOL/server/services/blobs"
	"github.com/igorrafae/MIGDOL/server/services/delivery"
	"github.com/igorrafae/MIGDOL/server/services/directory"
	"github.com/igorrafae/MIGDOL/server/services/gateway"
)

func main() {
	aud := audit.New()
	dir := directory.New(aud)
	del := delivery.New(dir, aud)
	_ = blobs.New(dir)
	_ = gateway.New(dir, del)
}
