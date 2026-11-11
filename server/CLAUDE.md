# server/ — Monólito modular em Go, pronto para separar em microsserviços

Contexto: docs/02 seções 7 e 10; docs/adr/0001-monolito-modular.md; docs/MIGRACAO-MICROSSERVICOS.md.

## Estrutura
- `services/<nome>/`: um módulo Go por serviço (go.mod próprio), unidos por `go.work`.
  - `api/`: contrato público (interfaces e tipos). Única coisa que outros serviços importam.
  - `<nome>.go`: construtor `New(...)`, usado só pela composição.
  - `internal/`: implementação privada.
  - `migrations/`: esquema do serviço (dono único dos dados).
  - `CLAUDE.md`: responsabilidade, dados, dependências, plano de extração. Leia antes de mexer.
- `platform/ids`: único código compartilhado (tipos de identificador). Mantenha mínimo.
- `cmd/server`: composição. Liga implementações locais aos contratos. Sem regra de negócio.

## Grafo de dependências permitido (somente via api/)
| Serviço | Pode depender de |
|---|---|
| audit | — |
| directory | audit |
| delivery | directory, audit |
| blobs | directory |
| gateway | directory, delivery |

`scripts/check-boundaries.sh` falha se um go.mod declarar dependência fora desta tabela.
Mudar a tabela exige atualizar o script, este arquivo e o CLAUDE.md do serviço.

## Regras
- Nunca importar `internal/` ou implementação de outro serviço. Só `api/`.
- Cada tabela tem um dono. Sem consulta nem FK entre serviços.
- Sem dependência circular. Sem tipo de implementação vazando pelo contrato.
- Contratos recebem `context.Context` e devolvem erro: preparados para virar chamada de rede.
- Autorização por sala em toda operação; assinatura antes de gravar estado; logs sem conteúdo.
- Verificação: `cd server && ../scripts/check-boundaries.sh && go build -buildvcs=false ./...` por módulo (ver script).
