# services/audit

## Responsabilidade
Grava eventos de metadados em cadeia de hash; serve o perfil Auditor; publica o hash periodicamente.

## Não faz
Não recebe nem guarda conteúdo de mensagem.

## Dono dos dados
audit_log (migrations/0001_init.sql).

## Depende de (somente via api/)
nenhum (folha do grafo)

## Contrato exposto
api.Recorder

## Regras do módulo
- Código em `internal/` é privado; outros módulos só importam `api/`.
- Nunca ler nem escrever tabela de outro serviço. Nunca criar FK para tabela de outro serviço.
- Nunca importar tipo de implementação de outro serviço, só o contrato.
- Regras globais: CLAUDE.md da raiz (autorização por sala, assinatura antes de gravar estado, logs sem conteúdo).

## Plano de extração
Candidato mais simples a processo próprio: só recebe eventos. Cliente remoto de api.Recorder com envio assíncrono e fila local para não perder evento.

Fase: 1 (eventos de cadastro), 4 (encadeamento e perfil Auditor)
