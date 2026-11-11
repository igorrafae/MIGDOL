# services/gateway

## Responsabilidade
Borda: termina mTLS 1.3, valida certificado e lease, mantém WebSocket, limita taxa, verifica Origin.

## Não faz
Não interpreta conteúdo e não expõe contrato a outros serviços; só consome.

## Dono dos dados
Nenhum.

## Depende de (somente via api/)
directory/api, delivery/api

## Contrato exposto
Nenhum (é cliente dos demais)

## Regras do módulo
- Código em `internal/` é privado; outros módulos só importam `api/`.
- Nunca ler nem escrever tabela de outro serviço. Nunca criar FK para tabela de outro serviço.
- Nunca importar tipo de implementação de outro serviço, só o contrato.
- Regras globais: CLAUDE.md da raiz (autorização por sala, assinatura antes de gravar estado, logs sem conteúdo).

## Plano de extração
Já é o ponto de entrada. Ao separar serviços, passa a falar com eles por rede usando os clientes remotos das mesmas interfaces.

Fase: 1
