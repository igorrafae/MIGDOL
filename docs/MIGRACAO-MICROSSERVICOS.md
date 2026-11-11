# Plano de migração para microsserviços

Objetivo: extrair um serviço de `server/services/` para um processo próprio sem reescrevê-lo. Pré-requisitos já garantidos pela estrutura: módulo próprio, contrato em `api/`, dados e migrações próprios, grafo de dependências verificado.

## Quando extrair (gatilhos)
- Escala diferente do restante (ex.: delivery por volume de salas).
- Isolamento de segurança ou de falha (ex.: audit com credenciais próprias).
- Ciclo de implantação independente.
Registrar a decisão em novo ADR antes de começar.

## Roteiro por serviço
1. **Contrato:** congelar `api/`. Se tipos de transporte forem necessários, definir o esquema (protobuf ou JSON) a partir dele.
2. **Servidor:** criar `services/<x>/cmd/<x>d/` que expõe o contrato por rede (mTLS interno) usando a implementação existente.
3. **Cliente:** criar `services/<x>/client/` que implementa a mesma interface de `api/` chamando o servidor, com timeout, repetição limitada e política de **negar por padrão** em falha.
4. **Composição:** em `cmd/server`, trocar `x.New(...)` pelo cliente. Nenhum outro serviço muda.
5. **Dados:** subir banco próprio, aplicar `migrations/` e migrar as tabelas do serviço. Como não há FK entre serviços, o corte é limpo.
6. **Implantação:** imagem própria em `deploy/`, fixada por digest; rede e credenciais próprias.
7. **Observabilidade:** métricas e eventos sem conteúdo; identificadores de sala só em hash.
8. **Testes:** rodar tests/security (C01–C13) contra a composição remota.

## Pontos de atenção por serviço
| Serviço | Atenção |
|---|---|
| audit | Envio assíncrono com fila local; não perder evento; hash externo continua publicado |
| directory | Chamado em toda publicação: cache com TTL curto no cliente; negar se indisponível |
| delivery | Fan-out entre múltiplas instâncias do gateway exige barramento; expurgo limpa por chamada |
| blobs | Quase sem estado; separa fácil |
| gateway | Continua como borda; passa a usar os clientes remotos |

## O que NÃO muda
Cifra ponta a ponta, servidor sem conteúdo em claro, autorização por sala, assinatura antes de gravar estado. A separação não pode ampliar o que cada serviço enxerga.
