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
| **G1** | Aparelho Android físico de teste (spike de atestação) e iPhone dedicado (estudo de viabilidade do iOS, ADR-0007) | Infraestrutura |
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
| 3 | Spike MLS: core Rust + CLI; no mínimo quatro clientes (dois administradores, dois membros), inclusão e remoção; caminho do Keystore em emulador; sem React (ADR-0007 aceita) | Opus 5.5 (`xhigh`) | core/, docs/02 | `core/spikes/mls/` + `docs/fase0/03-spike-mls.md` | Teste automático passando; dependências justificadas com hash (regra 8) | 3b |
| 3b | Revisão de segurança do spike MLS | Opus 5.5 (`xhigh`) | saída 3 | `docs/fase0/03b-revisao-spike-mls.md` | Nenhum achado alto aberto | **G1** |
| 5 | Spike de atestação (aparelho real) | Opus 5.5 (`xhigh`) | aparelho de G1 + arquivo da ADR-0002 (atestação sem Play Integrity) | `docs/fase0/05-spike-atestacao.md` | Cadeia válida aceita; emulador e root recusados; desembrulho pelo Keystore em hardware real, chamado do core (item 7b do doc 03) | **G2** |
| 5b | Estudo de viabilidade do iPhone (ADR-0007 aceita; 01e-E19a) | Opus 5.5 (`high`) | iPhone de G1 + protótipo do doc 11 | `docs/fase0/05b-estudo-ios.md` | Relatório sobre APNs, App Attest, MDM da Apple, bloqueio de captura de tela, série atestada e conexão em segundo plano; **sem código de produto**; insumo da ADR própria (D-R31) | **G2** (em paralelo com a 5) |
| 6 | Spike de inferência (Qwen, BGE-M3, PaddleOCR) | Sonnet 5.5 (`medium`) | servidor de G2 | `docs/fase0/06-spike-inferencia.md` | Memória e tempos medidos | 7 |
| 7 | Revisão final da Fase 0 | Fable 5.2 (`high`) | todas as saídas | `docs/fase0/07-revisao-final.md` | Lista do que confirma ou derruba a arquitetura | **G3** |

Cadeia: **1 → 4 → GD → 1b → G0 → 2 → 3 → 3b → G1 → 5 e 5b → G2 → 6 → 7 → G3 → GC**.

O critério da etapa 3 foi ampliado pela revisão (D33) e pelas rodadas da 1b (`01d`, `01f`): no mínimo quatro clientes da CLI (dois administradores, dois membros); cópia de custódia por HPKE com `info` ligando sala, período (só por tempo, hipótese do spike) e classe, na suíte candidata **DHKEM(P-256, HKDF-SHA256), HKDF-SHA256, AES-256-GCM** (GD-20); assinatura do remetente por dentro da cópia com rótulo da biblioteca; um commit por época; **bifurcação induzida pelo servidor detectada nos dois lados** pelo hash do transcript nos dados autenticados em claro e por mensagem da mesma época que não decifra; **bifurcação provada** (SEV1) só com commit conflitante cuja assinatura de administrador registrado o cliente verifica, ou importação de cadeia, e pacotes forjados pelo servidor ou por um membro, em qualquer quantidade, só como **anomalia não autenticada**, sem SEV1 e sem recomendar recriar a sala; se a biblioteca deixa verificar commit de época já aplicada (01g-G13); **partição total** declarada como não detectável em banda e detectada na reunificação; matriz de falso positivo (duplicata, fora de ordem, época anterior atrasada, byte corrompido); commit retido e custo da confirmação de época; membro que mente no campo; remoção de aparelho offline; caminho da chamada ao Keystore a partir do core **em emulador**, com a saída zerada no quadro da chamada (regra 2), ficando o hardware real para a etapa 5 (01e-E04a); levantamento do suporte a propostas por membro (01c-C16); checagem da API da biblioteca (`authenticated_data` por mensagem lido antes de processar, hash do transcript, `epoch_authenticator`); o que o commit público exporia (01e-E09). Texto completo no doc 03, Fase 0.

## Estado

| # | Estado |
|---|---|
| 1 | **concluída**: `01-revisao-desenho.md` (5 críticos, 15 altos, como o `01b` conferiu; ADR-0003 e ADR-0004 não devem ser aceitas como estão) |
| 4 | **concluída**: `04-construir-x-adaptar.md` (construir com OpenMLS; não adaptar Matrix, Wire nem Signal) |
| GD | **respondido**: `GD-respostas.md` |
| 1b | **concluída**: `01b-correcoes.md` (5 críticos fechados; altos D08, D28, D30 adiados) |
| 1c | **concluída** (Fable): `01c-revisao-correcoes.md`, ADRs 0003–0006 e regra 14 "aprovar com ressalvas"; 27 achados novos (0 críticos, 2 altos: C01, C02) |
| — | ADR-0007 (plataformas), 0008 (painel e IA de análise) e 0009 (RBAC + ABAC) registradas como proposta; entram na segunda rodada da 1b e na revisão do Fable |
| — | **parado**: o critério da 1b (altos fechados) não passou por C01 e C02. Segunda rodada da 1b (Opus) depois das perguntas 4 (regra 2) e 6 (revogação) do `01c` |
| GD (2ª parte) | **respondido**: GD-12 (regra 2), GD-13 (revogação), GD-14 (D-R29 opção B), em `GD-respostas.md` |
| 1b (rodada 2) | **concluída** (Opus): `01d-correcoes-rodada2.md`. C01 e C02 fechados no texto; regras 2 e 14 do `CLAUDE.md` alteradas; ADRs 0007–0009 integradas. Conflito registrado: ADR-0008 × regra 3 (nível 1b). Aguarda revisão do Fable |
| 2, 3, 3b | aguardando G0 |
| 5 | aguardando G1 |
| 5b | aguardando G1 (em paralelo com a 5; 01g-G26b) |
| 6 | aguardando G2 |
| 7 | aguardando 6 |
| 1e | **concluída** (Fable): `01e-revisao-rodada2.md`; C01 e C02 fechados; 21 achados novos (0 críticos, 1 alto: 01e-E01, nível 2 da ADR-0008); ADR-0008 só nível 1a aprovado |
| — | ~~parado: decisão do responsável sobre Q1/N1, Q2 e N2/Q3~~ |
| GD (3ª parte) | **respondido**: GD-15 (nível 2 por concessão; regra 3 sem mudança), GD-16 e GD-17 (investigação sigilosa; GD-17 pendente de confirmação), GD-18 (destruição na captura e verificação facial), GD-19 (ADR-0007 aceita; doc 11), GD-20 (P-256) |
| 1b (rodada 3) | **concluída** (Opus): `01f-correcoes-rodada3.md`. ADR-0010 nova (investigação sigilosa); ADR-0008 com nível 2 redesenhado (01e-E01 fechado no texto); ADR-0004 com a exceção escrita à GD-13 para o nó; ADR-0007 aceita e regra 4 alterada; critério do spike revisto (01e-E04, GD-20); 19 achados do 01e fechados, 2 em parte (E03, E18). Aguarda revisão do Fable, depois G0 |
| 1f/1g | 3ª rodada (Opus) `01f-correcoes-rodada3.md` e revisão (Fable) `01g-revisao-rodada3.md`: 01e-E01 fechado; ADR-0010 "não aceitar como está" (2 altos: 01g-G01, G02); doc 11 aprovado como rascunho |
| — | ~~parado: decisões do responsável sobre as 5 perguntas do `01g`~~ respondidas: GD-21 a GD-25 |
| 1h | **concluída** (Opus, passe curto): `01h-correcoes-rodada4.md`. ADR-0010 com GD-21, GD-22 e GD-23 (01g-G01 e G02 fechados no texto); regra 12 alterada (GD-25); GD-24 no doc 07; G13 no critério do spike; doc 11 corrigido e riscos R120–R128 e decisões D-R42–D-R47 no doc 04. Aguarda conferência do Fable só em G01 e G02; depois G0 |
| 1h/1i | 4ª rodada (Opus) `01h-correcoes-rodada4.md` e conferência (Fable) `01i-conferencia-rodada4.md`: G01 e G02 fechados; 0 críticos, 0 altos, 3 médios (01i-I01–I03), 4 baixos. **Critério da 1b cumprido** |
| — | **parado no portão G0** (autorização formal da chefia para enviar código a IA externa). Pendências do responsável no `01i` §6, nenhuma trava G0 |
