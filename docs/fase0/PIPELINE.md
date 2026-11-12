# Fase 0 — Pipeline automático

Cada etapa tem um modelo (docs/03 seção 6.1), uma entrada, uma saída em arquivo e um critério. Quando a saída passa no critério, o orquestrador (sessão principal) **aciona a próxima etapa sozinho**. Quando chega num **portão humano**, a cadeia para e espera.

Regras do pipeline:
- Nenhuma etapa faz commit ou push: o orquestrador junta as saídas e pede autorização ao responsável.
- Nenhuma etapa envia dado real, segredo ou texto classificado.
- Saída que não passa no critério volta para a mesma etapa uma vez; na segunda falha, a cadeia para e reporta.
- O orquestrador confere cada saída antes de acionar a próxima (não confia só no relatório da etapa).

## Portões humanos

| Portão | O que libera | Quem |
|---|---|---|
| **G0** | Autorização formal para enviar **código** a IA externa (docs/03 seção 6.1) | Chefia |
| **G1** | Aparelho Android físico de teste (spike de atestação) | Infraestrutura |
| **G2** | Servidor com GPU (spike de inferência) | Infraestrutura |
| **G3** | Decisões D-R1–D-R5, D-R17, D-R18; RIPD; responsáveis nomeados | Chefia e jurídico |
| **GD** | Respostas às perguntas 1, 2, 9 e 11 da revisão de desenho (custódia, histórico na v1, ordem D-R17/D-R18, missão de campo) | Responsável pelo produto e chefia |
| **GC** | Commit e push das saídas | Responsável pelo produto |

## Etapas

| # | Etapa | Modelo | Entrada | Saída | Critério | Próxima |
|---|---|---|---|---|---|---|
| 1 | Revisão de desenho: ADR-0003, ADR-0004, ameaças, invariantes | Fable 5.2 (`high`) | docs/ e docs/adr | `docs/fase0/01-revisao-desenho.md` | Achados com arquivo e seção, severidade e proposta | 4 |
| 1b | Corrigir ADR-0003, ADR-0004 e documentos conforme os achados D01–D34 | Opus 5.5 (`xhigh`), revisado por Fable 5.2 | saída 1 + respostas de GD | ADRs e docs revisados + `docs/fase0/01b-correcoes.md` | Cada achado crítico e alto fechado ou aceito por escrito | 3 (depois de G0) |
| 4 | Construir × adaptar (uma página com a decisão) | Opus 5.5 (`high`) | docs/ + saída 1 | `docs/fase0/04-construir-x-adaptar.md` | Opções comparadas contra as regras do CLAUDE.md; recomendação | **G0** |
| 2 | Integração contínua: build, testes, `check-boundaries.sh` | Sonnet 5.5 (`medium`) | server/, scripts/ | workflow de CI + README de CI | Build e fronteiras rodando localmente | 3 |
| 3 | Spike MLS: core Rust + CLI; dois clientes, inclusão e remoção | Opus 5.5 (`xhigh`) | core/, docs/02 | `core/spikes/mls/` + `docs/fase0/03-spike-mls.md` | Teste automático passando; dependências justificadas com hash (regra 8) | 3b |
| 3b | Revisão de segurança do spike MLS | Opus 5.5 (`xhigh`) | saída 3 | `docs/fase0/03b-revisao-spike-mls.md` | Nenhum achado alto aberto | **G1** |
| 5 | Spike de atestação (aparelho real) | Opus 5.5 (`xhigh`) | aparelho de G1 | `docs/fase0/05-spike-atestacao.md` | Cadeia válida aceita; emulador e root recusados | **G2** |
| 6 | Spike de inferência (Qwen, BGE-M3, PaddleOCR) | Sonnet 5.5 (`medium`) | servidor de G2 | `docs/fase0/06-spike-inferencia.md` | Memória e tempos medidos | 7 |
| 7 | Revisão final da Fase 0 | Fable 5.2 (`high`) | todas as saídas | `docs/fase0/07-revisao-final.md` | Lista do que confirma ou derruba a arquitetura | **G3** |

Cadeia: **1 → 4 → GD → 1b → G0 → 2 → 3 → 3b → G1 → 5 → G2 → 6 → 7 → G3 → GC**.

O critério da etapa 3 foi ampliado pela revisão (D33): cópia de custódia por HPKE, um commit por época com teste de bifurcação, remoção de aparelho offline e desembrulho pelo Keystore via core.

## Estado

| # | Estado |
|---|---|
| 1 | **concluída**: `01-revisao-desenho.md` (5 críticos, 16 altos; ADR-0003 e ADR-0004 não devem ser aceitas como estão) |
| 4 | **concluída**: `04-construir-x-adaptar.md` (construir com OpenMLS; não adaptar Matrix, Wire nem Signal) |
| GD | **respondido**: `GD-respostas.md` |
| 1b | **em andamento** (Opus), depois revisão Fable |
| 2, 3, 3b | aguardando G0 |
| 5 | aguardando G1 |
| 6 | aguardando G2 |
| 7 | aguardando 6 |
