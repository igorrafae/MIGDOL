# ADRs

As decisões ADR-01 a ADR-15 estão resumidas em docs/02 seção 11. Mudança que as contrarie exige um novo ADR aqui, antes do código.
Formato: `NNNN-titulo.md` com Contexto, Decisão, Alternativas rejeitadas, Consequências.

## Registradas aqui

| Arquivo | Decisão | Estado |
|---|---|---|
| 0001 | Monólito modular | aceito |
| 0002 | Atestação sem Play Integrity | proposta no doc 07 seção 8; **arquivo é entrada da etapa 5** (spike de atestação; `docs/fase0/PIPELINE.md`) |
| 0003 | Custódia institucional em sala sob custódia, com quórum presencial | proposta, revisada na etapa 1b, na segunda e na terceira rodada (GD-13, GD-14; série de investigação e preservação selada da ADR-0010); quarta: lista de bloqueio conferida na eliminação (01g-G01c) |
| 0004 | Nó de campo e lease de campo | proposta, revisada na etapa 1b; terceira rodada: destruição imediata do nó capturado como exceção escrita à GD-13 (GD-18) e passagem do lease (01e-E12); quarta: falha de boot não destrói (01g-G20); 01c-C17 pendente |
| 0005 | Um commit por época, detecção de bifurcação e pertença verificável | proposta (etapa 1b); item 2 reescrito na segunda rodada (01c-C02); sinais classificados e commit retido na terceira (01e-E02, E03); quarta: bifurcação provada × anomalia não autenticada (01g-G13) |
| 0006 | Lease criptográfico do cache local | proposta (etapa 1b); revogar não destrói (GD-13); terceira rodada: 01e-E05, E06, E08, E12, E13; substitui ADR-05 se aceita |
| 0007 | Plataformas: Android na v1, iPhone depois do estudo, sem macOS; telas nativas; núcleo Rust único; monorepo | **aceita** (GD-19); regra 4 alterada; `app/` vira `android/` no início da Fase 1 |
| 0008 | Painel de gestão, indicadores e IA de análise com escopo | proposta; nível 1a no painel; nível 1b fora (GD-15); nível 2 só por MLS e por concessão da sala de origem (GD-15), redesenhado na terceira rodada; 01e-E01 fechado no texto (01g); limite da posição declarado (01g-G19) |
| 0009 | Controle de acesso por papéis com atributos (RBAC + ABAC) | proposta; decidir antes da Fase 1 |
| 0010 | Investigação sigilosa de vazamento ou infiltrado | proposta (GD-16; GD-17 confirmada pela GD-23); terceira rodada; passe curto da quarta (`01h`): GD-21 (quórum em toda sessão), GD-22 (controle externo do rito), GD-23; 01g-G01 e G02 fechados no texto, a conferir pelo Fable só nesses pontos |

## Correspondência de numeração (D34)

Os resumos do doc 02 seção 11 usam dois dígitos; os arquivos daqui usam quatro. **Não são a mesma série.**

| Resumo (doc 02 §11) | Arquivo | Atenção |
|---|---|---|
| ADR-11 | 0001 | |
| ADR-12 | 0003 | ADR-03 (vetores cifrados) **não** é ADR-0003 |
| ADR-14 | 0004 | ADR-04 (expurgo por chave) **não** é ADR-0004 |
| ADR-05 | 0006 | 0006 substitui ADR-05 se aceita |

Decisões novas usam só o número de arquivo (quatro dígitos).
