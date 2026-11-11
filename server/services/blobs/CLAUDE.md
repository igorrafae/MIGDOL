# services/blobs

## Responsabilidade
Emite URLs temporárias de envio/leitura, controla cota e expiração dos anexos cifrados.

## Não faz
Não tem chaves; não decifra nada.

## Dono dos dados
blobs (room_id sem FK) e o armazenamento de objetos.

## Depende de (somente via api/)
directory/api

## Contrato exposto
api.Store

## Regras do módulo
- Código em `internal/` é privado; outros módulos só importam `api/`.
- Nunca ler nem escrever tabela de outro serviço. Nunca criar FK para tabela de outro serviço.
- Nunca importar tipo de implementação de outro serviço, só o contrato.
- Regras globais: CLAUDE.md da raiz (autorização por sala, assinatura antes de gravar estado, logs sem conteúdo).

## Plano de extração
Quase sem estado; separável cedo. Armazenamento de objetos já é externo.

Fase: 3
