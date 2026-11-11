// Package ids define identificadores compartilhados entre serviços.
// É o único código comum permitido entre módulos; mantenha mínimo e sem regra de negócio.
package ids

// RoomID é o identificador aleatório de 128 bits de uma sala (sem significado).
type RoomID [16]byte

// DeviceID identifica um aparelho cadastrado.
type DeviceID [16]byte
