# ADR-0001 — Monólito modular com módulos Go separados por serviço

**Status:** aceito · **Complementa** docs/02 seção 7.1 e 7.2 (não contraria os ADR-01 a ADR-10)

## Contexto
O documento 02 define a v1 como um único binário com fronteiras de pacote claras e prevê separar em processos "quando houver necessidade operacional". Fronteiras só de pacote não impedem acoplamento: nada barra um import indevido nem uma consulta a tabela alheia.

## Decisão
Um único binário (`cmd/server`) compondo **módulos Go independentes** (um `go.mod` por serviço, unidos por `go.work`):
1. Cada serviço expõe um contrato em `api/` e esconde a implementação em `internal/`.
2. Serviços dependem uns dos outros só por `api/`, segundo um grafo fixo, verificado por `scripts/check-boundaries.sh`.
3. Cada serviço é dono único das suas tabelas, com migrações próprias. **Sem chave estrangeira entre serviços.**
4. Contratos usam `context.Context` e erros, como se fossem chamadas de rede.
5. O repositório `store/` central do rascunho foi removido: cada serviço tem o seu acesso a dados.

## Alternativas rejeitadas
- **Microsserviços desde já:** custo operacional (rede, tracing, implantação) sem necessidade comprovada; contraria docs/02 7.1.
- **Pacotes num só módulo:** não impede dependência indevida.

## Consequências
- Ajuste ao esquema do doc 02 7.5: `outbox` e `blobs` perdem a FK para `rooms`; a limpeza por expurgo passa a ser explícita (chamada/evento), não `ON DELETE CASCADE`.
- Pequeno custo de cerimônia: contratos e construtores.
- Dependências externas (pgx, WebSocket etc.) só entram na Fase 1, por serviço, com versão fixada e hash.
- Revisar quando surgir necessidade real de escala ou isolamento por serviço.
