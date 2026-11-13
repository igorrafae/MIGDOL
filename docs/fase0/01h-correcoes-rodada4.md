# Correções da revisão — Fase 0, etapa 1b, quarta rodada (passe curto)

**Modelo:** Claude Opus 5.5 (passe curto pedido no `01g` seção 9, condição 1). **Conferência seguinte:** Fable 5.2, só em 01g-G01 e G02.
**Entradas:** `CLAUDE.md`; `GD-respostas.md` (GD-21 a GD-25 prevalecem); `01g-revisao-rodada3.md` (achados 01g-G01 a G26; seção 6, fusão dos R-U); `01f-correcoes-rodada3.md`; ADR-0010; doc 11; seções citadas pelos achados.
**Limites:** só texto; nenhum código, commit ou push; nada enviado para fora. Ponto jurídico segue **[conferir]**. Nada de cripto à mão: o texto descreve recursos da biblioteca MLS, da pilha HPKE e do hardware, e o que depende deles ficou **[verificar no spike]** ou **[verificar]**. Nenhum risco é aceito aqui (doc 04 seção 3.2). Única regra do `CLAUDE.md` alterada: a **12** (GD-25), autorizada pelo responsável.
**Convenção:** achados do `01g` citados como `01g-Gnn` ou só `Gnn`.

## Resumo

| Severidade | Total | Fechado no texto | Em parte (resto pendente de resposta) | Adiado |
|---|---|---|---|---|
| Alta | 2 | 2 (G01, G02) | 0 | 0 |
| Média | 17 | 16 | 1 (G19) | 0 |
| Baixa | 7 | 4 | 3 (G23, G24, G25) | 0 |
| **Total** | **26** | **22** | **4** | **0** |

- **Altos e críticos do `01g` abertos: nenhum.** G01 e G02 estão fechados **no texto** da ADR-0010 pelas GD-21 e GD-22, com uma leitura a confirmar ("CI validada", seção 4 item 1). A ADR-0010 continua **proposta**.
- "Fechado" quer dizer **fechado no texto**. G13 depende de prova no spike MLS; G09 e parte de G01, G02 e G07 dependem do jurídico; G10a fica como proposta até o spike de inferência; os valores de G11 e G12 vêm das medições exigidas pela GD-24.
- Riscos novos **R120 a R128** e decisões **D-R42 a D-R47** registrados no doc 04 (regra 11). Telas do painel renomeadas **PW01 a PW12**. IDs conferidos por busca: cada um definido uma só vez; antes desta rodada só apareciam como plano no `01f` e no `01g`.

## 1. Achados 01g-G01 a G26

| ID | Sev. | Status | Onde | Nota |
|---|---|---|---|---|
| G01 | Alta | **Fechado no texto** (conferir pelo Fable) | ADR-0010 status, contexto, §3 item 1, §4, §5, §8, §10, alternativas, consequências; ADR-0003 §7; doc 07 A86, A89 | GD-21: **toda sessão** é cerimônia com M de N detentores não impedidos; o quórum abre o registro e a estação lê pedido e impedidos antes de qualquer ato; não há cópia de trabalho fora do quórum (a). Como toda sessão abre o registro selado (classe "registros"), a base mínima é a da regra 14: **CI validada** (leitura na seção 4). (b) Revisão de controle em processo administrativo de controle **[conferir]**; nomeá-la na regra 14 é a pergunta N3. (c) **Lista de bloqueio** assinada, só com identificadores opacos de chaves de período, conferida na cerimônia de eliminação sem abrir o registro nem usar a chave da série de investigação **[proposta]** |
| G02 | Alta | **Fechado no texto** (conferir pelo Fable) | ADR-0010 §2, §6, §10; doc 04 R112, §6.2, D-R37, D-R38; doc 07 A84, A86, A92 | (1) Quórum em toda sessão (GD-21). (2) Revisor não é o iniciador nem a segunda autorização (GD-22). (3) No máximo duas renovações, cada uma com quórum; depois, encerramento, e pedido novo sobre o mesmo escopo só com ordem judicial **[proposta]**. (4) Encerramento por ato registrado, com comunicação a órgão de controle fora da cadeia **[proposta; conferir qual]**. (5) Registro guardado até a revisão final (GD-22) e o prazo de responsabilização **[conferir]**. (6) Aviso no arquivamento como padrão; adiar exige decisão escrita e revista. (7) Retrato só por ato de liberação do backup com duas pessoas. (e) Limite declarado: quem tem acesso legítimo ao backup correlaciona fora do rito (R38) |
| G03 | Média | Fechado | ADR-0010 §2, §3 item 2, §10; doc 04 §6.2; doc 07 A85 | Iniciador na posição pode denunciar e iniciar, mas não opera a estação, não revê e não recebe; segunda autorização confirma escopo e impedidos como ato próprio |
| G04 | Média | Fechado no texto | ADR-0010 §2, §3; doc 04 R115, D-R37 | (a) Limite declarado do detentor suspeito não identificado; segunda autorização de fora da classe (GD-22). (b) Reservas com cartão emitido de antemão, lacrado. (c) Recomendação reforçada para a D-R24 (detentor não é membro de sala operacional), que segue aberta |
| G05 | Média | Fechado | ADR-0010 §3 item 2; ADR-0009 §3; doc 04 R115, §6.2; doc 07 A85 | Lista assinada de papéis leva o identificador opaco do usuário; comparação sem quebrar pseudonimato; falha fechada |
| G06 | Média | Fechado | ADR-0010 §5, §6; doc 04 R113; doc 11 PW09 | (a) Entradas de cobertura **[proposta]** e limite no R113. (b, c) Sessões junto da rotina, operadores e testemunhas conferidos, gravação selada. (d) Painel sem preservação selada. N pequeno declarado |
| G07 | Média | Fechado | ADR-0010 §6, §8; doc 06 §3, §7 item 13, §8.3; termo item 14 | (a) Ato de liberação do backup casado com o teste trimestral. (b) Sala comum só metadados (GD-23). (c) Retrato abre só com quórum e base da regra 14. (d) Guarda alongada declarada no inventário, na transparência e no termo. (e) Cópia do selado no retrato |
| G08 | Média | Fechado | ADR-0010 §9; doc 04 R114 | Rótulo atestado × registro do servidor sem assinatura, que não eleva o grau sozinho; comparação de retratos |
| G09 | Média | Fechado no texto (jurídico) | ADR-0010 §2, §6, §9, §10; doc 03 Fase 4 | (a) Pedido cita o ato do procedimento. (b) Texto neutro sobre o acesso. (c) Hashes, parâmetros e semente no registro. (d) Modelo de pedido ao juízo a criar. (e) Hash da mídia na liberação. (f) Parecer antes da Fase 4. Tudo **[conferir]** |
| G10 | Média | Fechado no texto | ADR-0010 §4, §6; doc 02 §1 | (a) Duas máquinas da mesma classe **[proposta; decidir depois do spike de inferência]**. (b) Presença mínima e sessão longa dividida; duração em D-R38. (c) O que o detentor vê antes do cartão; pode recusar |
| G11 | Média | Fechado (GD-24) | doc 07 §6.4, §7.1, §7.1a, §7.2, R67, R116, A76, A77, D-R40; doc 04 D-R40 | Duas capturas para "outra pessoa"; falsificação fora do gatilho de destruição até medida à parte; "detecta, não barra"; limite de falso negativo perto de zero em condição de campo; colega que destrói o aparelho de outro no R67 |
| G12 | Média | Fechado (GD-24) | doc 07 §7, §7.1, §7.1a, R55, A79, RF96; doc 10 §3.2; doc 09 §5; doc 11 4.7 | (a) R55 com impacto Alto no perfil missão; PIN de coação no nível 4 só depois de medir o acionamento acidental. (b) Ordem: alerta, envio, zerar. (c) Isca fora do cache e sem diferença de tempo e aquecimento (A79). (d) Doc 11 corrigido |
| G13 | Média | Fechado no texto (prova no spike) | ADR-0005 status, item 2c, 2d, consequências; doc 02 §5.5; doc 03 Fase 0 item 4; `PIPELINE.md`; doc 07 §5; doc 04 R98 | **Bifurcação provada** (commit conflitante com assinatura de administrador registrado verificada, ou importação de cadeia) → SEV1. **Anomalia não autenticada** (pacotes que não decifram, em qualquer quantidade) → "servidor ou membro anômalo", sem recomendar recriar a sala, nunca SEV1. Spike: a biblioteca deixa verificar commit de época já aplicada? |
| G14 | Média | Fechado | doc 11 T02, 4.1; doc 06 §7 | Assinatura no gov.br fora do app, na estação do gestor; T02 mostra texto, versão e estado que o core entrega |
| G15 | Média | Fechado | doc 11 T19, 4.7, 4.8; doc 07 §7.2 | Fluxos reescritos da matriz do doc 07 §7.1a com a GD-24; captura durante a digitação; igual com PIN de coação; efeitos por perfil; caminho de volta do dono que caiu na isca |
| G16 | Média | Fechado | doc 11 §3.2, PW07–PW10, §8, §9 | (a) PW09 só preservação comum; (b) nós de campo com D-R20; (c) faixas em vez de "<5"; (d) "o nível 1b não existe"; (e) papel privilegiado como objeto assinado |
| G17 | Média | Fechado | doc 11 §6.4, §6.5, §6.6, §6.8 (nova) | Os nove itens: linha de base, metas formativas, teste cego do observador, retenção, recuperação, tarefas novas, estresse de campo, ética, rodada por perfil |
| G18 | Média | Fechado | doc 11 §5.3, U09, §6.5, §6.6; doc 04 R128, D-R44 | Sem reconhecimento real no papel e no navegável; no nativo só com rito (consentimento específico, gabarito zerado e conferido, parecer do encarregado); taxa de erro em A76 e A78 |
| G19 | Média | **Em parte** | ADR-0008 §3; doc 11 T25, 4.10, §8; doc 04 R127, D-R20 | Limite declarado ("só" retirado). **Pendente:** grupo MLS de comando ou limite aceito (Fase 8; D-R20; pergunta 9 do `01c`) |
| G20 | Baixa | Fechado | ADR-0004 item 13; doc 10 §3.1, D-R39 | Falha de boot não destrói; só não libera a chave do disco |
| G21 | Baixa | Fechado | doc 10 §3.2 | "Conforme o perfil da instituição (D-R40)" |
| G22 | Baixa | Fechado (dúvida registrada) | doc 07 §7.2, R117 | (a) Chave de dados embrulhada pelo hardware, gabarito decifrado em Rust, nunca na JVM. (b) Marca de missão protegida por chave de hardware, fora do cache. Dúvida: a chave de dados abre o gabarito sem outro segredo; conferir na Fase 2 com o "sozinho não abre nada" da regra 2 |
| G23 | Baixa | **Em parte** | doc 07 A96; termo item 3 | (b) Caso novo em A96. (c) Exceção declarada no termo. **Pendente (a):** "CI sempre" muda a GD-15; é a D-R41 e a R3-7, que o responsável não respondeu |
| G24 | Baixa | **Em parte** | `admin-panel/CLAUDE.md` | Sem armazenamento do navegador nem recurso externo no painel. **Pendente:** "por exemplo" na regra 4 do `CLAUDE.md` exige autorização para alterar a regra 4 |
| G25 | Baixa | **Em parte** | doc 11 (nota, PW01–PW12, etapas sem sigla, T22–T24 na Fase 4, 4.5, §8, §9); doc 03 §4.12 V09; `admin-panel/CLAUDE.md` | (a), (b), (d), (e) feitos; U-1 a U-7 viraram D-R. (c) Doc 03 feito. **Pendente:** doc 11 na lista de leitura do `CLAUDE.md` raiz (fora da regra 12; exige autorização) |
| G26 | Baixa | Fechado | ADR-0010 §7; doc 04 R119; `PIPELINE.md` | (a) Limite do expurgo de sala comum pelo investigado. (b) Linha da etapa 5b |

## 2. Decisões GD-21 a GD-25 aplicadas

| Decisão | O que mudou | Onde |
|---|---|---|
| **GD-21** (quórum em toda sessão) | Toda sessão é cerimônia com M de N detentores não impedidos; sem M livres não há sessão; níveis só pela base (nível 1 com CI validada; nível 2 com ordem judicial); a divisão 1a/1b do `01f` deixa de existir | ADR-0010 §3, §4, §5, §8, §10, alternativas, consequências; doc 04 R112, R113, R115, §6.2, §5.5; doc 05 §3.5, P1-S; doc 03 Fase 4; doc 02 §11 (ADR-0010); doc 01 RF93; doc 07 A86; termo item 14 |
| **GD-22** (controle externo do rito) | Segunda autorização de fora dos detentores quando a suspeita os alcança; no máximo duas renovações com quórum; revisor distinto do iniciador e da segunda autorização; aviso no arquivamento como padrão; registro até a revisão final | ADR-0010 §2, §5, §10; doc 04 D-R37, D-R38, R112, §6.2; doc 07 A84, A86, A92; doc 06 §7 item 13; termo item 14; doc 01 RF93; doc 02 §11; doc 05 P1-S |
| **GD-23** (GD-17 confirmada; R3-2) | "Expurgo suspenso" = eliminação da cópia de custódia; metadados de sala comum preservados sem cifrado MLS, retrato selado e aberto pela regra 14; transparência da guarda alongada | ADR-0010 status, contexto, §4, §6, §7, §8; doc 06 §3, §7 item 13, §8.3; termo item 14; doc 04 R119; ADR README |
| **GD-24** (gatilhos no perfil missão) | Duas capturas para "outra pessoa"; vivacidade fora do gatilho de destruição até medida à parte; número de PINs errados só depois de medir (proposta: um); PIN de coação no nível 4 só depois de medir o acionamento acidental; até lá, nível 3 | doc 07 §6.4, §7, §7.1, §7.1a, §7.2, R55, R67, RF96, A76, A77, A79, D-R40; doc 04 D-R40, §5.5; doc 10 §3.2; doc 09 §5; doc 11 T19, 4.7, 4.8 |
| **GD-25** (regra 12) | Câmera só em primeiro plano, na tela de conteúdo e no desbloqueio; gabarito só no aparelho, cifrado por chave de hardware, nunca enviado | `CLAUDE.md` regra 12; doc 07 §7.2 (parágrafo da regra 12), R117 |

## 3. Regra 12 do `CLAUDE.md` (GD-25)

Antes:
> 12. Câmera: sem alerta, nenhum quadro é gravado nem enviado; com alerta, só imagens de evidência
>     cifradas para a chave de quórum (docs/07). Atestação sem Play Integrity nem serviço de terceiro.

Depois:
> 12. Câmera: só com o app em primeiro plano, na tela de conteúdo e no desbloqueio. Sem alerta, nenhum
>     quadro é gravado nem enviado; com alerta, só imagens de evidência cifradas para a chave de quórum
>     (docs/07). O gabarito facial fica só no aparelho, cifrado por chave de hardware, e nunca é enviado.
>     Atestação sem Play Integrity nem serviço de terceiro.

Texto da GD-25, sem acréscimo. Nenhuma outra regra do `CLAUDE.md` foi alterada. A frase de escopo da D-R30 (foto do termo de revinculação, T23, e câmera do app para foto, T10, fora de alerta; pergunta 2 do `01c`) **não** foi decidida e não entrou: lida ao pé da letra, a regra nova não cobre a câmera fora da tela de conteúdo e do desbloqueio, como na foto do termo, e a frase antiga já não cobria foto enviada sem alerta.

## 4. Leituras aplicadas e dúvidas registradas

1. **"CI validada" (GD-21)** lida como CI em procedimento instaurado, com validação jurídica: a base que a regra 14 já nomeia. Motivo: a partir da segunda sessão, toda sessão abre o registro selado (classe "registros"), e a regra 14 exige base registrada. Consequência: não há investigação só com a CI e a segunda autorização (o antigo nível 1a). Opção mais restritiva; se o responsável quiser o nível com CI simples, a regra 14 precisa nomear essa base (pergunta N3 do `01g`, que pode nomear também a revisão de controle).
2. **Lista de bloqueio** (G01c): mecanismo novo, **[proposta]**, para a cerimônia de eliminação respeitar a preservação sem abrir o registro selado. A conferir pelo Fable.
3. **"Outra pessoa" com duas capturas** aplicado a todo uso do resultado, inclusive o nível 2, como diz a GD-24 sem distinguir. Reduz a chance de o dono cair na isca; reduz também a detecção do adversário com PIN correto. Confirmar.
4. **Gabarito e regra 2** (G22a): a chave de dados desembrulhada abre o gabarito guardado no aparelho sem outro segredo, porque a matriz precisa do resultado facial também com PIN errado. Conferir na Fase 2.
5. **Renovações esgotadas** (G02): pedido novo sobre o mesmo escopo só com ordem judicial **[proposta]**, para que o limite da GD-22 não se contorne com pedido novo.

## 5. Pendências

### Do responsável
| # | Pendência | Achado |
|---|---|---|
| 1 | Confirmar a leitura de "CI validada" ou mudar a regra 14 (com N3 e a revisão de controle como base) | G01 |
| 2 | Grupo MLS de comando para a posição, ou limite aceito (Fase 8; D-R20; pergunta 9 do `01c`) | G19 |
| 3 | Base mínima da concessão: "CI sempre" (D-R41; R3-7) | G23a |
| 4 | Autorizar "por exemplo" nos parênteses da regra 4 | G24 |
| 5 | Autorizar o doc 11 na lista de leitura do `CLAUDE.md` | G25c |
| 6 | Frase de escopo da D-R30 na regra 12 (pergunta 2 do `01c`) | GD-25 |
| 7 | D-R24 (detentor de cartão não é membro de sala operacional), D-R37, D-R38, D-R40 (com as medições), D-R42 a D-R47 | G04c, G02, G11, G12, doc 11 |
| 8 | Ainda abertas de rodadas anteriores: Q4, Q5, N3, N4, N5, N6; R3-4, R3-6, R3-10; perguntas 2 a 13 do `01c` e 4, 5, 6, 8 e 15 da etapa 1 (lista do `01f` seção 7) | — |

### Do jurídico [conferir]
Base do nível 1 (procedimento instaurado e validação; se a CI basta); CI sigilosa; órgão de controle que recebe o encerramento; prazo de responsabilização por abuso; acesso do investigado (Súmula Vinculante 14); aviso no arquivamento e seu adiamento; cadeia de custódia do retrato (CPP arts. 158-A e seguintes); ordem judicial que cubra análise automatizada sem leitura e modelo de pedido; metadados de comunicação de agentes sem ordem judicial (parecer antes da Fase 4); guarda alongada de metadados na preservação (LGPD); rito do rosto real no teste de usabilidade; texto do termo, itens 3 e 14.

### Para o Fable (conferência curta)
G01 e G02 na ADR-0010, inclusive a leitura de "CI validada" e a lista de bloqueio (ADR-0003 seção 7). G13 na ADR-0005 item 2c, se quiser, antes da etapa 3.

### Provas
G13 no spike MLS (commit de época já aplicada); G10a e duração da sessão depois do spike de inferência; G11 e G12 nas medições de A76, A78, A79 e no V09 (U05, U06, U09) antes de ligar o nível 4; G22 na Fase 2.

## 6. Arquivos alterados nesta rodada

| Arquivo | O que mudou |
|---|---|
| `CLAUDE.md` | Regra 12 (GD-25) |
| `admin-panel/CLAUDE.md` | Navegador sem armazenamento e sem recurso externo; telas PW (G24, G25) |
| `docs/adr/0010-investigacao-sigilosa-de-vazamento.md` | Reescrita: GD-21 a GD-23; G01 a G10, G26a; continua proposta |
| `docs/adr/0003` | Seção 7: lista de bloqueio (G01c) |
| `docs/adr/0004` | Item 13: falha de boot não destrói (G20) |
| `docs/adr/0005` | Status, item 2c, 2d, consequências (G13) |
| `docs/adr/0008` | Seção 3: limite da posição (G19) |
| `docs/adr/0009` | Seção 3: identificador opaco na lista assinada (G05) |
| `docs/adr/README.md` | Estado das ADRs 0003, 0004, 0005, 0008, 0010 |
| doc 01 | RF93 |
| doc 02 | Seção 1 (estação de análise), seção 5.5 (G13), seção 11 (ADR-0010) |
| doc 03 | Fase 0 item 4 (G13); Fase 4 (ADR-0010); seção 4.12 V09 (doc 11, R120) |
| doc 04 | D-R20, D-R37, D-R38, D-R40, **D-R42 a D-R47**; R98, R102, R105, R106, R112–R115, R119; seção 6.2; **seção 5.5 nova** com **R120 a R128** e mudanças de perfil |
| doc 05 | Seção 3.5 e P1-S |
| doc 06 | Inventário (retrato), seção 7 (item 13; assinatura fora do app), seção 8.3 |
| doc 07 | Seções 5, 6.4, 7, 7.1, 7.1a, 7.2; R55, R67, R69, R116, R117; A76, A77, A79, A84–A86, A89, A92, A96; RF96; D-R40 |
| doc 09 | Seção 5 (perfil missão) |
| doc 10 | Seções 3.1, 3.2; D-R39 |
| doc 11 | G14 a G19, G25; fluxos 4.7 e 4.8 reescritos; seção 6.8 nova; seção 7 com os IDs definitivos; decisões D-R |
| `docs/anexos/termo-de-ciencia-e-autorizacao.md` | Itens 3 e 14 |
| `docs/fase0/PIPELINE.md` | Critério da etapa 3 (G13); estados 5b e 1h |
| `docs/fase0/01h-correcoes-rodada4.md` | **Novo** (este arquivo) |

Não alterados: `GD-respostas.md`, doc 08, `app/CLAUDE.md`, anexo de validação jurídica.
