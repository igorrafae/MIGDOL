# services/delivery

## Responsabilidade
Roteia envelopes por sala, mantém a fila com prazo e coleta confirmações.

## Não faz
Não conhece remetente nem conteúdo; não decide quem é membro (pergunta ao diretório).

## Dono dos dados
outbox, outbox_pending (room_id sem FK).

## Depende de (somente via api/)
directory/api, audit/api

## Contrato exposto
api.Publisher

## Regras do módulo
- Código em `internal/` é privado; outros módulos só importam `api/`.
- Nunca ler nem escrever tabela de outro serviço. Nunca criar FK para tabela de outro serviço.
- Nunca importar tipo de implementação de outro serviço, só o contrato.
- Regras globais: CLAUDE.md da raiz (autorização por sala, assinatura antes de gravar estado, logs sem conteúdo).

## Plano de extração
Escala por sala. Ao separar, o fan-out para conexões WebSocket precisa de barramento (ex.: pub/sub) entre instâncias do gateway; decidir em ADR na hora.

Fase: 1
