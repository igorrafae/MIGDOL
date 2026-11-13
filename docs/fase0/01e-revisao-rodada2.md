# Revisão da segunda rodada — Fase 0, etapa 1e

**Modelo:** Claude Fable 5.2 (pipeline da Fase 0, revisão da segunda rodada da etapa 1b). Nenhum arquivo foi alterado além deste; nada de commit; nada enviado para fora.
**Entradas:** `CLAUDE.md` (regras 2 e 14), `GD-respostas.md` (GD-1 a GD-14, prevalece), `01c-revisao-correcoes.md` (C01–C27), `01d-correcoes-rodada2.md`, ADRs 0003 a 0009 e README, `git diff` dos docs 01 a 10, dos anexos, de `app/CLAUDE.md`, `admin-panel/CLAUDE.md` e do `PIPELINE.md`.
**Limites:** leitura de texto; nada foi executado. O que afirmo sobre MLS (RFC 9420), HPKE (RFC 9180), OpenMLS, Android Keystore e HSM vem de conhecimento geral e está marcado **[verificar no spike]** onde depende da biblioteca. Trechos que o diff não tocou foram lidos só onde um achado dependia deles.
**Convenção:** achados desta revisão são `01e-Enn` (o doc 03 já usa E01–E09 para cenários). Achados do 01c seguem como `01c-Cnn`.

## 1. Veredito

| Peça | Veredito | Condição |
|---|---|---|
| **Regra 2** (`CLAUDE.md`) | **Aprovar** | Fiel à GD-12 e ao ajuste de 01c-C19. Ressalva baixa: 01e-E15 (zerar na JVM é melhor esforço) |
| **Regra 14** (`CLAUDE.md`) | **Aprovar** | Fiel à GD-13. Recomendo aprovar também a redação de 01c-C23 (pergunta N3). A frase "membros leem as salas pelo MLS" precisa de decisão explícita sobre o nível 2 da ADR-0008 (01e-E01) |
| **ADR-0003** (custódia) | **Aprovar com ressalvas** | 01e-E06, E09, E10 no texto; 01c-C06 e C11 seguem com o HSM; D-R36 |
| **ADR-0004** (campo) | **Aprovar com ressalvas** | 01e-E07 (destruição do nó) e 01e-E12 (passagem do lease) antes de aceitar; 01c-C17 pendente. Fase 8: não trava nada agora |
| **ADR-0005** (commits e bifurcação) | **Aprovar com ressalvas** | **01c-C02 fechado no texto.** 01e-E02, E03 e E04 entram no critério do spike antes da etapa 3; 01c-C16 pendente |
| **ADR-0006** (lease) | **Aprovar com ressalvas** | **01c-C01 fechado.** 01e-E05, E06, E08 e E13 no texto antes da Fase 2 |
| **ADR-0007** (plataformas) | **Aprovar com ressalvas** | Coerente. Decidir N2 antes da etapa 3 ou, no máximo, antes da Fase 1; 01e-E19 |
| **ADR-0008** (painel e IA de análise) | **Aprovar só o nível 1a.** Nível 1b: não (manter a regra 3). Nível 2: **não aceitar como está** | 01e-E01 (alto) e 01e-E11. É proposta para depois da v1: não trava G0 (seção 9) |
| **ADR-0009** (RBAC + ABAC) | **Aprovar com ressalvas** | 01e-E10 e E18. Precisa estar decidida antes da Fase 1 (RF89 e RF90 são "Essencial, Fase 1") |

Contagem dos achados novos: **0 críticos, 1 alto, 12 médios, 8 baixos** (seção 3).
Nenhuma ADR volta a ser reprovada por inteiro. O único alto (01e-E01) está na ADR-0008, que é nova, é proposta e só se implementa depois da Fase 6.

## 2. Conferência de 01c-C01 e 01c-C02

### 2.1 C02 — ADR-0005 item 2: **fechado no texto; prova no spike**

| Ponto | Conferido | Observação |
|---|---|---|
| `epoch_hash` nos dados autenticados em claro | Sim. ADR-0005 item 2a; doc 02 §5.3 e §5.5; doc 07 §5; R98; R11; ADR-0004 item 3; F-R12 | Em `PrivateMessage` o campo `authenticated_data` vai em claro e entra nos dados associados da cifra; a época também já vai em claro no quadro. O hash do transcript confirmado faz parte do contexto do grupo, e dois ramos da mesma época têm valores diferentes. Desenho correto **[verificar no spike: a biblioteca deixa gravar o campo por mensagem e lê-lo antes de processar]** |
| Detecção sem decifrar | Sim, item 2b | Os três sinais existem de fato. O sinal 2 (mesma época que não decifra) já bastaria para detectar; o sinal 1 serve para **diagnóstico** e evidência (qual ramo), não para autenticar. Ver 01e-E02 |
| Valor falso do servidor | Sim, item 2d | Correto: trocar o campo quebra a autenticação. Mas "alarme falso" não é inofensivo (01e-E02) |
| Partição total como limite declarado | Sim, item 2e; doc 02 §5.5; R98; R90 | Declarado com honestidade. O limite é mais largo do que o texto sugere e dá para estreitá-lo (01e-E03) |
| `epoch_auth` removido | Sim | Só resta a explicação da retirada (doc 02 §5.3; ADR-0005 Contexto) |
| Custo de expor o hash | Aceitável | Com commit privado, o servidor ganha um valor opaco por época, que ele já distingue pela ordem. Não vejo vazamento novo. Se o spike escolher commit **público** (item 4, alternativa), ver 01e-E09 |

**Resposta à pendência 1 do Opus (negação de serviço):** comparar antes de decifrar **não** abre canal novo de indisponibilidade (o servidor já pode reter ou corromper). Abre um canal de **alarme forjado**: um pacote com a época corrente e 32 bytes quaisquer dispara SEV1 em todos os membros. O risco é fadiga de alerta e indução a "recriar a sala". Tratar em 01e-E02.

**O desenho é sólido e verificável no spike?** Sim, com três complementos: classificar os sinais (E02), cobrir o commit retido e a confirmação de época (E03) e ajustar o critério (E04).

### 2.2 C01 — ADR-0006 itens 3, 8 e 10: **fechado**

| Ponto | Conferido | Observação |
|---|---|---|
| Revogar e destruir são atos separados | Sim. Item 8; doc 07 §7.1 e §7.4; doc 02 §4.3 e §6.3; doc 05 P2; RF05, RF91; R63, R66, R110; A74 | Fiel à GD-13 |
| Selagem em **toda** revogação, sem depender do motivo | Sim | Falta o caso da renovação pendente, com `n` e `n+1` (01e-E05) |
| Destruição com duas autorizações e registro | Sim | Quem são os dois não está dito (01e-E18). A destruição é exclusão de registro no `server/`; o controle é procedimental (01e-E08) |
| Prazo da apuração | Sim, com D-R35 | Quem fixa, piso, prorrogação e casamento com a eliminação por período: 01e-E06 |
| Classe na regra 14 e na ADR-0003 §4 | Sim | — |
| "Sozinho não abre nada" | Sim | Correto: exige aparelho e desbloqueio do agente. Ver a nuance de 01e-E20 |
| Item 3 (geração, `n` e `n+1`, gravação atômica) | Sim | Protocolo correto nos três pontos de queda que conferi (antes de gravar, depois de gravar e antes de confirmar, confirmação perdida). Acrescentar: pedido repetido devolve o **mesmo** `n+1` (idempotente) |
| Item 10 (RPO 0) | Sim, como proposta | Coerente com doc 04 §8 e V05. Recomendo aceitar (N6) |
| Campo (item 9) | Sim | Passagem do lease entre central e nó não descrita (01e-E12); fase e grafo de módulos (01e-E13) |

### 2.3 Critério do spike (doc 03 Fase 0 e `PIPELINE.md`)

Cobre o que o 01c pediu para C02 (os dois casos, o valor falso e a declaração de membros). **Não está completo.** Faltas em 01e-E04: itens que a etapa 3 não consegue cumprir como está (Keystore antes de G1; "suíte candidata do HSM" sem HSM), divergência entre `PIPELINE.md` e doc 03, e cinco verificações que decidem se o item 2 da ADR-0005 funciona na prática.

## 3. Achados novos

| ID | Sev. | Onde | Problema | Proposta |
|---|---|---|---|---|
| 01e-E01 | **Alta** | ADR-0008 §1 (nível 2) e §2; ADR-0009 §1 e §3; doc 08 §5.1; doc 02 §6.4 e §9.2; RN04, RN09; GD-1 | O nível 2 entrega a quem **não é membro** conhecimento nominal derivado do conteúdo, por um membro que a instituição controla (o nó de IA), sem ato da sala de origem. Contradições concretas: (a) GD-1 e doc 02 §6.4: sala comum é "inacessível à instituição por desenho" e "expurgada, sem cópia", mas o indício derivado fica no cache dos analistas; (b) com a flag ligada, a sala de análise nasce sob custódia e leva para o acervo de 5 anos indícios derivados inclusive de **salas comuns** (as que alguém justificou tirar da custódia); (c) RN04: o expurgo da sala de origem não alcança o que já foi entregue à sala de análise; (d) ADR-0009 §3 "nenhum papel dá leitura de conversa": o analista pergunta o que quiser sobre salas que não lê; (e) o escopo é criado "pela chefia", sem participação de quem responde pelo compartimento. Não fere a letra da regra 14 (a lista de classes não inclui análise); fere o princípio "membros leem as salas pelo MLS" e a compartimentação. É conflito entre duas decisões do responsável (GD-1 e o nível 2): só ele resolve | Seção 5.2. Em resumo: tratar o nível 2 como **exportação autorizada pela sala de origem** (concessão de escopo assinada pelo administrador da sala, visível aos membros, mais aprovação dupla da chefia); regra de tipo (sala de análise sob custódia só inclui salas sob custódia); expurgo em cascata; saída estruturada (01e-E11). Sem a concessão da sala, a instituição só chega ao conteúdo pelo rito da regra 14 |
| 01e-E02 | Média | ADR-0005 item 2b–2d; doc 02 §5.5; doc 07 §5 | Os sinais não são classificados. (a) O sinal 1 vindo de outro ramo **não é autenticado** por quem recebe: servidor forja SEV1 com um pacote. (b) Membro que mente no campo gera o sinal 1 em mensagem que **decifra**: é falha atribuível ao remetente, não bifurcação. (c) O sinal 2 tem causas benignas: duplicata (a chave da mensagem já foi consumida), fora de ordem além da janela, corrupção, retransmissão de rádio. SEV1 a cada uma inutiliza o alerta, sobretudo em campo | Tabela de classificação na ADR: **autenticado** (decifra e o hash diverge: remetente inconsistente); **forte** (não decifra nem os dados do remetente, com hash diferente; ou segundo commit para época já aplicada); **fraco** (um pacote isolado). SEV1 só com sinal forte repetido ou dois sinais distintos; sinal fraco conta e vira SEV2. Duplicata e fora de janela têm erro próprio da biblioteca e não contam **[verificar no spike]**. Limite de alertas por sala e época |
| 01e-E03 | Média | ADR-0005 item 2e e Consequências; R98; R11; doc 07 §7.1 (D17) | "Partição total" é o caminho **normal** de um servidor hostil, porque ele controla o roteamento; e o caso mais provável não é o de dois commits, é o **commit retido** (remoção entregue só a parte do grupo). Nenhum sinal do item 2b cobre: quem ficou na época antiga recebe só mensagens antigas, e o hash confere. O alerta D17 só avisa que o administrador não fez o commit; não avisa que os membros não o receberam. R98 com residual "Baixo (depois do spike)" e R11 "Baixo" ficam otimistas para P Médio e I Crítico com esse limite | (a) **Confirmação de época:** depois de aplicar um commit, cada membro envia (ou pega carona em) uma mensagem na época nova; o cliente do administrador mostra "membros que ainda não confirmaram" e alerta depois de um prazo. O servidor não forja essa confirmação. Transforma a partição silenciosa em ausência visível; custo de uma mensagem curta por membro por commit (medir em banda estreita). (b) Sinais 4 e 5: mensagem de época **futura** sem commit entregue em T; par que segue enviando em época **antiga** depois de T. (c) R98 em **Médio** até os testes da Fase 1; revisar R11 |
| 01e-E04 | Média | doc 03 Fase 0; `PIPELINE.md` etapa 3 | Critério do spike incompleto ou impossível de cumprir: (a) o item 7 (desembrulho pelo Keystore) exige Android, e a etapa 3 vem **antes** de G1; (b) o item 1 pede "suíte candidata do HSM", mas o HSM não foi levantado (pergunta 11); (c) `PIPELINE.md` diz "core Rust + CLI; dois clientes", o doc 03 diz "dois aparelhos, ponte Kotlin e chamada do React"; (d) faltam verificações de que o item 2 depende | (a) Item 7 em duas partes: caminho da chamada em emulador na etapa 3; hardware real na etapa 5. (b) Fixar a candidata por padrão: DHKEM P-256 com HKDF-SHA256 e AES-256-GCM (a curva mais comum em HSM), a confirmar. (c) Alinhar os dois textos; mínimo de **quatro** clientes (dois administradores, dois membros). (d) Acrescentar: API da biblioteca (gravar e ler `authenticated_data` antes de processar; hash do transcript; `epoch_authenticator`); matriz de falso positivo (duplicata, fora de ordem, época anterior atrasada, byte corrompido); sinal forjado pelo servidor não vira SEV1 sozinho; commit retido e confirmação de época (E03); membro que mente no campo; o que o commit público expõe (E09) |
| 01e-E05 | Média | ADR-0006 itens 3 e 8 | Revogação com renovação pendente: o servidor guarda `n` e `n+1`, e o aparelho pode estar em qualquer um. O item 8 manda selar "o segredo vigente". Se selar só um e o aparelho estiver no outro, o cache fica ilegível **para sempre**: o defeito de 01c-C01 num canto | "Sela **todas** as gerações que guarda para o aparelho (`n` e, se houver, `n+1`), com o número da geração no registro." Teste em A74 |
| 01e-E06 | Média | ADR-0006 item 8 (prazo); ADR-0003 §4 e §7; doc 06 §8.2; painel | (a) Quem fixa o "prazo da apuração" é a mesma pessoa que revoga: prazo de um dia vira destruição por uma pessoa, com atraso. (b) Não há prorrogação descrita. (c) A eliminação é "pela chave do período da classe", mas cada segredo tem prazo próprio: ou a chave do período espera o maior prazo (o painel promete uma data que não é a da destruição criptográfica), ou apaga-se registro (sem garantia em backup) | (a) Prazo padrão e **piso** no `instituicao.yaml` (D-R35); abaixo do padrão, só com as duas autorizações. (b) Selar para a chave do **período de vencimento** (mês em que o prazo termina), como a ADR-0003 §7 já sugere para a custódia: a eliminação por período fica exata. (c) Prorrogação é ordem de preservação sobre a chave do período (RF58), declarada como tal. (d) O painel distingue "registro excluído" de "chave eliminada" |
| 01e-E07 | Média | ADR-0004 item 7; doc 10 §3.1, R89, RNF-R2 | Destruição do nó de campo (escada própria e botão físico) é destruição irreversível por **um ato**, sem duas autorizações, e atinge terceiros: cache de missão de todo aparelho não apreendido (depois de reinício ou de 72 h), **cópias de custódia e trilha ainda não sincronizadas**, segredos já selados dos removidos e a lista de revogação. Os gatilhos da escada do nó não estão definidos. Seção 6 | "Selar antes de destruir": (a) o nó sela cada segredo de missão na **emissão** e entrega o selado ao próprio aparelho; (b) a chave do disco do nó tem cópia selada para uma chave pública de quórum, gravada no próprio disco; o botão destrói só a cópia operacional. Com isso a destruição do nó deixa de ser irreversível e cabe na GD-13 sem exceção. Sem isso, pedir decisão explícita ao responsável |
| 01e-E08 | Média | ADR-0006 item 8; R71; R110; doc 02 §7.1 | (a) Revogar é ato de uma pessoa e, depois da selagem, o cache só volta por cerimônia de quórum: revogação errada ou **em lote** por uma conta de gestor (ou conta comprometida) tranca a frota. (b) O registro selado vive no `server/`: administrador de banco apaga sem autorização nenhuma e ninguém nota | (a) Revogação em lote com duas autorizações, como o apagamento (R71); padrão em caso de dúvida é suspender (nível 3, reversível). Declarar em R110 que aparelho achado volta por revinculação e `history_share`, não por cerimônia. (b) Hash do registro selado na trilha de auditoria no ato da selagem; conferência periódica pelo Auditor; backup com RPO 0 (item 10) |
| 01e-E09 | Média | ADR-0005 item 4 (alternativa) e item 6; ADR-0003 §1; doc 02 §7.5; R111 | (a) Se o spike escolher commit como mensagem **pública**, as propostas de extensão do contexto (marca, motivo, papéis, chave de período e, se estiver lá, o nome da operação) vão em claro ao servidor: contraria "o motivo vai cifrado no estado do grupo" e a regra 3. (b) "Hash do motivo na trilha": com lista fechada, o hash **é** o motivo; com texto livre curto, sai por dicionário | (a) Condição na alternativa: commit público só se nenhum dado sensível transitar em proposta; senão, commit privado e declaração separada (texto atual). Critério do spike (E04). (b) Categoria do motivo em lista fechada, assumida como metadado; texto livre só no estado do grupo, sem hash na trilha (ou hash com sal guardado no grupo) |
| 01e-E10 | Média | ADR-0009 §2 e §3; ADR-0003 §1 (segunda assinatura, GD-14); doc 08 §5 e §5.1; doc 04 §6.2 | A dupla conferência servidor e cliente está certa para ações **dentro do grupo**. Fora dele, papel, grau, compartimento e "atribuição por duas pessoas" são registro do diretório, que é hostil. (a) Segunda assinatura da sala comum: a sala ainda não existe, e quem monta o estado do grupo é o criador; sem raiz fora do diretório, ele inventa o segundo signatário e o R111 fica sem controle. (b) Credencial do analista e escopo: conferidos por quem? (c) Registrar no diretório quem é detentor de cartão, para recusar combinações, cria a lista de quem coagir (R80). (d) "Atribuição por duas pessoas" para todo papel contradiz RN05 (um gestor cadastra o agente) | (a) Atribuição de papel privilegiado vira **objeto assinado** pelas duas pessoas, com chaves da lista assinada em cerimônia; clientes e nó de IA conferem. (b) Segunda assinatura da sala comum e concessão de escopo conferidas contra essa lista. (c) Combinações com detentor de cartão conferidas na **cerimônia de emissão**, contra a lista assinada de papéis; identidade dos detentores fora do diretório, ou selada. (d) Atribuição dupla vale para papéis privilegiados; o cadastro do agente segue RN05 e D20 |
| 01e-E11 | Média | ADR-0008 §2 ("referência, não texto"), §5; R104, R106 | Com pergunta livre e resposta gerada, "sem trecho de conversa" não é garantível: o modelo pode copiar o contexto, e uma mensagem plantada numa sala do escopo pode induzi-lo a copiar ou a fabricar vínculo. O filtro antes da busca protege **entre** escopos, não o conteúdo **dentro** do escopo. R104 com residual Baixo é otimista sem teste | Saída **estruturada** e validada no código: entidade, tipo de vínculo, referência (sala e data), grau; texto livre do modelo só como explicação com limite de tamanho e verificação de sobreposição com a fonte. Teste de injeção dentro do escopo. R104 em Médio até o teste. Limite de perguntas por analista e revisão por segunda pessoa do que foi perguntado |
| 01e-E12 | Média | ADR-0006 item 9; ADR-0004 item 7; doc 10 §4 e §5 | A passagem do lease entre a central e o nó não está descrita: gerações na saída e no retorno, queda no meio, quem confirma. É o problema de 01c-C12 em dois servidores. Aparelho removido em campo: o segredo da central fica obsoleto, e o que abre o cache é o de missão selado no nó, que se perde se o nó for destruído antes de sincronizar (E07) | Descrever a passagem com o mesmo protocolo de duas fases (o nó entra como `n+1` na saída; a central como `n+1` no retorno). Com a selagem na emissão (E07a), o selado útil já está no aparelho |
| 01e-E13 | Média | doc 03 Fases 2 e 4; RF91; A74; doc 02 §7.1; `server/CLAUDE.md` | (a) RF91 e A74 são **Fase 2**, mas a selagem precisa do serviço `custody`, do lote de chaves públicas da classe e de cerimônia de HSM, que são **Fase 4**. "Uma fase por vez" fica sem saída. (b) Os leases são do `directory`; os selados, do `custody`. O grafo só permite `custody → directory`. Para selar, o `directory` precisa da chave pública e de entregar o selado: seria `directory → custody`, um ciclo (regra 9) | (a) Na Fase 2, o `directory` sela com a chave pública da classe lida do pacote assinado (chave pública não é segredo) e guarda o selado em tabela própria; chave de teste até a cerimônia. (b) Na Fase 4, o `custody` **busca** os selados no `directory` (sentido permitido) ou a tabela fica no `directory`. Registrar a escolha na ADR-0006 e no doc 02 §7.1 |
| 01e-E14 | Baixa | ADR-0008 §1 (nível 1a); R105 | "<5" não resiste a diferença entre semanas nem a filtro por unidade pequena. "Missões ativas" e nós de campo são metadados operacionais sensíveis no perfil Exército (D-R20) | Arredondar ou agrupar contagens, sem filtro que reduza o grupo abaixo do limite; lista fechada de indicadores no `instituicao.yaml`; rever com D-R20 |
| 01e-E15 | Baixa | Regra 2; doc 02 §4.1 | "É zerado em seguida": na JVM, zerar vetor é melhor esforço (cópias do coletor e do JNI). A regra afirma como fato | No spike: saída do Keystore em `ByteBuffer` direto alocado pelo core, sem passar pelo heap **[verificar]**. Incluir a JVM no "limite conhecido" do doc 02 §6.3 |
| 01e-E16 | Baixa | ADR-0005 item 2e; doc 01 RF04 | A ADR diz que o código de segurança "passa a incluir o hash do transcript", mas a RF04 não mudou. O código passa a ser por sala e por época (muda a cada commit), o que altera o uso | Atualizar a RF04. Para o código comparado por pessoas, usar o `epoch_authenticator` do MLS, que existe para isso **[verificar no spike]**; o hash do transcript fica para o campo em claro |
| 01e-E17 | Baixa | ADR-0003 §3; doc 07 R97; doc 04 §3.2 | R97 (residual Baixo): a ADR diz "aceite pelo responsável pela homologação"; o doc 07 diz "dono do ativo (`custody`)". Pela regra do doc 04, Baixo é do dono do ativo | Corrigir a ADR-0003 §3 |
| 01e-E18 | Baixa | ADR-0009 §1; doc 04 §6.2 | (a) "Gestor: nunca pode ver nomes" contradiz RN05 e R40 (ele cadastra com identificação). (b) Faltam papéis: segunda assinatura da sala comum (D-R36), autorizadores do apagamento remoto e da destruição do segredo selado, criador de sala de análise (D-R32), plantão da sala de resposta. (c) Falta a combinação proibida analista × gestor (o gestor conhece o vínculo matrícula e codinome) | Ajustar a linha do Gestor ("no painel de indicadores"); completar a tabela junto com D-R33 |
| 01e-E19 | Baixa | ADR-0007 item 1; `PIPELINE.md` G1; doc 03 Fase 0 e §6.1; RF89, RF90 | (a) O estudo de viabilidade do iOS tem portão (G1) e não tem etapa, saída nem critério. O escopo cita só APNs, App Attest e MDM; faltam bloqueio de captura de tela (doc 01 D2), série atestada (doc 07 §3) e conexão em segundo plano (D-R1). (b) A linha do spike MLS no doc 03 ainda pede "chamada a partir do React": se a ADR-0007 for aceita, é esforço jogado fora. (c) RF89 e RF90 são "Essencial (Fase 1)" e dependem da ADR-0009, que é proposta | Criar a etapa do estudo no `PIPELINE.md` (depois de G1), condicionada a N2. Decidir N2 antes da etapa 3. Decidir a ADR-0009 antes da Fase 1 |
| 01e-E20 | Baixa | doc 02 §6.4; doc 06 §6.3; doc 05 P11; ADR-0006 item 8 | "Sala comum: conteúdo inacessível à instituição por desenho" ganhou uma exceção que só a ADR-0006 declara: com aparelho revogado em mãos, segredo selado, quórum **e** desbloqueio do agente, a instituição lê o cache dele, inclusive de sala comum já expurgada nos outros aparelhos. Resposta a juízo que omita isso fica inexata (R62) | Nota nos três textos: "salvo o cache de aparelho revogado em poder da instituição, pelo rito da regra 14, dentro do prazo do segredo guardado" |
| 01e-E21 | Baixa | vários | Miúdos: (a) doc 06 §3 sem linha para indícios nominais do nível 2 (dados de terceiros) nem para o registro de perguntas à IA; (b) ADR-0008 §4 "relatórios entregues na sala" contra RN09 (só o administrador); (c) ADR-0006 item 8: "ilegível também para aparelho offline" vale só depois que o aparelho esquece o segredo (limite já declarado em outro ponto); (d) ADR-0006 não cita R110; (e) ADR-0003 §2 mantém "período a cada commit de membros" como texto normativo com nota "em revisão": o spike precisa saber qual vale; (f) `PIPELINE.md` sem linha de estado para esta revisão | Correções de texto. Em (e), marcar o período só por tempo como hipótese do spike |

## 4. Fidelidade às decisões

| Decisão | Aplicada? | Observação |
|---|---|---|
| **GD-12** (regra 2) | Sim, sem extrapolar | Texto igual ao aprovado, com Swift e o ajuste de 01c-C19. "StrongBox/TEE" segue RNF-A3. A regra 7 não foi tocada, como decidido |
| **GD-13** (revogação) | Sim | Acréscimos do Opus, todos marcados: "com indício de apreensão, só depois de D-R14" (coerente com D23), a volta ao aparelho **[proposta]**, D-R35 (o GD não dá valor ao prazo). Lacunas: 01e-E05, E06, E08. Alcance sobre o nó de campo: seção 6 |
| **GD-14** (D-R29, opção B) | Sim | A leitura "toda sala nova, sem exceção" saiu de todos os lugares (conferi por busca). Acréscimos razoáveis: pessoa distinta do criador, papel no `instituicao.yaml` (D-R36), revisão por amostragem pelo Auditor (R111). "Vale também para sala de compartimento" (doc 07 §7.3) decorre do GD. A sala de análise sob custódia é aplicação do padrão, com o efeito apontado em 01e-E01 |
| GD-1 a GD-6, GD-9, GD-11 | Sim, como no 01c | Sem regressão. GD-1 entra em conflito com o nível 2 da ADR-0008 (01e-E01) |
| Recomendações do 01c acolhidas no GD (R66; terceira assinatura) | Sim | R66 voltou a P Médio e residual "Alto até D-R14". Superior imediato do solicitante em RN18, RF83, A69, doc 02 §6.2, doc 04 §6.2 e ADR-0009 |
| Regras 3, 4 e 7 e a base da regra 14 | Não tocadas | Correto: mudar regra é do responsável. Textos propostos no 01d §3 |

## 5. ADR-0007, 0008 e 0009

### 5.1 ADR-0007 — coerente
- Bate com a regra 2 (GD-12), mantém a ADR-06 e não renomeia nada antes da aceitação. A correção "sem WebView a regra 4 não vale por construção" está certa.
- O iPhone encosta nas regras 10 (APNs) e 12 (App Attest consulta a Apple): a ADR adia para estudo e ADR própria, que é o caminho certo.
- Texto proposto da regra 4: recomendo a regra em termos gerais ("a interface não grava dado nem abre rede fora do core") e a **lista de APIs proibidas no teste estático**, não na regra, porque a lista envelhece.
- Ressalvas: 01e-E19.

### 5.2 ADR-0008

**(a) Divisão 1a e 1b: está certa.**
- **1a** usa só o que o servidor já guarda para funcionar. Não é conteúdo, não cria metadado novo e não precisa de exceção à regra 3. Aprovar, com 01e-E14.
- **1b** junta duas coisas diferentes, e convém separá-las na pergunta ao responsável:
  - *Tipo e região* não são conteúdo: não ferem a letra da regra 3. São **metadados novos** no servidor hostil, que hoje vivem cifrados no grupo. É mudança de perfil de risco (regra 11; R35; D-R20).
  - *Riscos por categoria* derivam do conteúdo. Publicá-los no servidor fere a regra 3 na leitura mais restritiva, e cria um canal em claro do nó de IA para o servidor, que um nó comprometido ou uma mensagem plantada podem usar para tirar informação.
- **Recomendação:** manter a regra 3 como está (opção a de N1). Se a chefia precisar dos agregados derivados, entregá-los **por MLS ao app** de quem tem o papel, como um nível 2 sem nomes, e nunca pelo painel web. Tipo e região ficam fora do servidor; no perfil Exército, não recomendo de forma alguma.

**(b) Nível 2: fere o espírito da regra 14 e a compartimentação, como está (01e-E01).**
O que a ADR já tem e deve ficar: escopo com prazo e revisão, credencial em todo o escopo e necessidade de saber por sala (doc 08 §5.1), filtro antes da busca, registro das perguntas, transparência. O que falta é o **ato da sala de origem**. Hoje a chefia cria o escopo e o nó de IA, que está em todas as salas, passa a servir de leitor institucional.

Solução recomendada, em ordem de importância:
1. **Concessão de escopo pela sala de origem.** Uma sala só entra num escopo se o administrador dela assinar a concessão dentro do grupo (mensagem de controle, como a autorização de exportação), visível a todos os membros. O nó de IA só cruza sala cuja concessão ele viu no grupo. Assim o nível 2 vira exportação autorizada pela sala, e o servidor ou a chefia sozinhos não alargam escopo.
2. **Aprovação dupla da chefia** para criar o escopo (as mesmas duas pessoas da revinculação, ou as de D-R32), com objeto assinado (01e-E10).
3. **Sem concessão, só o rito da regra 14.** Sala cujo administrador não concede fica fora; a instituição só chega ao conteúdo dela pela custódia e pelo quórum.
4. **Regra de tipo.** Sala de análise sob custódia só inclui salas sob custódia. Sala comum nunca alimenta acervo de custódia.
5. **Expurgo em cascata.** Indício leva a identificação da sala de origem; expurgo ou conclusão da origem apaga os indícios dela na sala de análise (prazo do indício nunca maior que o da origem). Em sala de análise sob custódia, o que já entrou na cópia segue o prazo dela, e isso precisa estar no termo.
6. **Não retroagir** (a decidir): a concessão vale para o que for dito depois dela, como a flag de custódia (GD-2).
7. **Saída estruturada** (01e-E11) e analista credenciado no compartimento (já no doc 08 §5.1).

Quórum presencial para criar escopo: **não recomendo** como regra geral. O rito da regra 14 é para abrir dado selado; usá-lo para rotina de análise o banaliza. Ele já é o caminho para o caso sem concessão (item 3).

Limite a declarar na ADR: a análise só alcança salas **ativas** com IA; sala concluída já foi expurgada do nó. "Presos em operações diferentes" ao longo dos anos não sai deste desenho.

**(c) ADR-0009.**
- **Corregedoria só pede abertura:** correto e fiel à regra 14. A correção do Opus era necessária.
- **Autorização dupla, servidor e cliente:** correta para ações dentro do grupo, e a frase "servidor hostil não concede papel" vale ali. Fora do grupo, falta a raiz de confiança (01e-E10).
- Negação por padrão, prazo no papel, ponto único no `directory` chamado por `api/`: coerentes com as regras 5 e 9.
- Ressalvas: 01e-E10 e E18. A pergunta 9 do 01c segue aberta.

## 6. Autodestruição do nó de campo sob a GD-13

**Cabe em parte; como está escrita, não.**

- A favor: a GD-13 fala de **revogação**; a escada do doc 07 §7.1 admite destruição "quando há adversário ativo confirmado", e o nível 4 do aparelho já destrói sem duas autorizações. Nó capturado é esse caso.
- Contra: o nível 4 do aparelho só atinge o próprio aparelho ("não apaga cópia no servidor, nos outros membros, custódia nem trilha"). A destruição do nó atinge **terceiros**:
  1. o cache de missão de todo aparelho não apreendido que reiniciar ou passar de 72 h (quem voltar à base ainda com o segredo em memória reembrulha e não perde);
  2. cópias de custódia e trilha de auditoria ainda não sincronizadas (a auditoria tem RPO 0 no doc 04 §8);
  3. segredos já selados dos removidos e a lista de revogação de campo, se ainda não saíram do nó.
- O ganho é pequeno: o segredo de missão sozinho não abre nada sem o hardware e o desbloqueio do aparelho. É o mesmo argumento que a ADR-0006 usa para guardar o segredo selado.
- Os gatilhos automáticos do nó não estão definidos. No aparelho, cabo de dados não aciona o nível 4; no nó, o texto não diz.

**Recomendação (01e-E07):** selar antes de destruir. O nó sela cada segredo de missão na emissão e entrega o selado ao aparelho; a chave do disco do nó tem cópia selada para o quórum no próprio disco. O botão e a escada destroem só o que está em claro. O adversário não ganha nada (falta o HSM), a instituição não perde nada irreversivelmente, e a GD-13 vale sem exceção. Se o responsável preferir a destruição pura, é exceção dele, por escrito, com R66 e R89 atualizados.

## 7. Riscos, IDs e referências

- **R66:** voltou a P Médio e residual "Alto até D-R14", igual ao estado anterior à etapa 1b. **Voltar os dois está certo** (pendência 7 do Opus): baixar P também muda o perfil e também é do comitê (doc 04 §2).
- **R97 e R68:** sem rótulo de aceite. Busquei "aceito" nos docs e nas ADRs; sobram só usos legítimos. Divergência pequena sobre quem aceita o R97 (01e-E17).
- **R98:** residual otimista (01e-E03). **R104:** idem (01e-E11).
- **IDs novos sem colisão:** R101–R111 (contínuos, sem lacuna), D-R31–D-R36, RF85–RF91, V01–V16 (só no doc 03 §4.12 e uma citação na ADR-0006), A74, A75, F-R12. As repetições que a busca acusa são linhas de índice (doc 04 §4, §5.1 e §5.3), não definições duplicadas.
- **Referências:** doc 03 §4.12, doc 08 §5.1, doc 04 §5.2 e §5.3 existem e são citadas corretamente. Não achei referência quebrada. D-R31 a D-R34 só têm definição no doc 04 §4, o que basta.
- **Convenção `01c-Cnn`:** boa; adotei `01e-Enn`. 01c-C21 (renomear D01–D34) adiado é aceitável, com prazo: antes do portão GC.
- **Contradições novas:** 01e-E01 (nível 2 × GD-1), E13 (fase e grafo), E18a (Gestor), E20 ("inacessível por desenho"), E21b (RN09).

## 8. Pendências que o Opus deixou para esta revisão (01d §5)

| # | Resposta |
|---|---|
| 1 | Sinais e alarme falso: seção 2.1; 01e-E02 e E03 |
| 2 | Item 8: aprovado, com 01e-E05, E06, E08. Nó de campo: seção 6 |
| 3 | Itens 3 e 10: protocolo correto; RPO 0 recomendado; acrescentar idempotência do `n+1` |
| 4 | ADR-0003: 01c-C05, C07, C08 e C10 bem aplicados. GD-14 depende da raiz de confiança da segunda assinatura (01e-E10a) e do tratamento do motivo (E09b) |
| 5 | ADR-0008: divisão certa; nível 2 pede decisão explícita (seção 5.2; 01e-E01) |
| 6 | ADR-0009: linha da corregedoria correta; 01e-E10 e E18 |
| 7 | R66: voltar P e residual está certo |
| 8 | 01c-C21 adiado com prazo (antes de GC); convenção aprovada |

## 9. Recomendações por pergunta

### Novas do 01d

| # | Recomendação |
|---|---|
| N1 | **Opção (a):** manter a regra 3; sem 1b no painel. Agregado derivado, se necessário, por MLS ao app. Tipo e região fora do servidor |
| N2 | **Aceitar a ADR-0007**, com a regra 4 em termos gerais e a lista de APIs no teste. Decidir antes da etapa 3, para o spike não gastar esforço com React |
| N3 | **Aprovar** a redação de 01c-C23. Alinha a regra aos documentos e ao GD-6 (CI sempre; base é ordem judicial ou processo administrativo com validação jurídica) |
| N4 | Segunda assinatura: **chefe da inteligência** (ou o superior imediato do criador, fixado no `instituicao.yaml`), nunca o criador nem membro da sala. Motivo em **lista fechada**, com texto livre opcional só dentro do grupo (01e-E09b) |
| N5 | **Aceitar** a posição na cópia de custódia. Exceção por tipo viraria desvio (01c-C09). Deixar claro no termo |
| N6 | **Aceitar** RPO 0 como meta, com a alternativa `n−1` se a replicação síncrona não for viável. Teto de 12 meses: aceitar, configurável no `instituicao.yaml` |

### Restantes do 01c

| # | Recomendação |
|---|---|
| 2 | Aprovar a frase de escopo na regra 12 (trata da câmera de vigilância; captura iniciada pelo usuário segue as regras 3 e 13). Foto **fora** da exportação. Prazo do termo: 5 anos, igual ao da custódia **[conferir]** |
| 3 (resto) | Troca de titular de papel: chefe da inteligência **e** um dos outros signatários da sala, pessoa distinta do novo titular; registrada no estado do grupo e visível aos membros. Nunca o administrador sozinho |
| 4 (resto) | Aprovar o texto proposto da regra 7 (destinos fixos, com a inferência local), antes da Fase 5 |
| 5 | Sala sob custódia: manter o kill-switch de um administrador (só some a cópia de trabalho; a custódia fica). Sala comum com a flag ligada: kill-switch com duas assinaturas (dupla custódia obrigatória), senão um administrador destrói tudo sem as três assinaturas. Teto: N6 |
| 6 (resto) | D-R35: padrão de **5 anos**, alinhado à custódia, com piso de 90 dias e destruição antecipada só com duas autorizações **[conferir]**. O segredo sozinho não abre nada, então o custo de guardar é baixo |
| 7 | Nomear o responsável pela homologação e o comitê **antes de G3** (já é entrega da Fase 0). O GD-1 não é aceite formal: registrar R68 (Médio: responsável), R97 (Baixo: dono do ativo) e R66 (Alto: responsável e comitê, com plano) |
| 8 | Sala de resposta: dois ou três **aparelhos de função** do plantão de segurança mais o gestor de serviço; meta de reação de 15 minutos; em missão, sala de resposta da missão com comandante e subcomandante; batimento de tamanho fixo |
| 9 | Não generalizar: comandante e subcomandante administram a **missão de campo** (GD-11). Nas demais, o chefe da operação designado na criação, com no mínimo dois administradores (01c-C16). Aprovadores distintos do gestor: manter; em unidade pequena, sobe um nível hierárquico |
| 10 | Manter a recusa **total** de sessão direta (mais restritiva): aparelho removido em campo não fala com sala nenhuma. Nó perdido: cadeia importável do aparelho do comandante, com a mesma verificação; descongelar com duas pessoas (comandante e chefe da inteligência). Credencial do nó renovada a cada dia quando houver enlace |
| 11 | Requisitos de compra: decapsulação interna, M de N nativo, destruição de chave inclusive em backup, capacidade de chaves, assinatura de lote sob quórum, roteiro de ML-KEM. Fixar já a suíte candidata do spike (01e-E04b) |
| 12 | **Sim**, chave do índice em HSM. A consulta por matrícula é rara (cadastro e revogação); o custo é aceitável |
| 13 | Perfil Exército: custódia em produção **espera** o KEM híbrido (adversário estatal, acervo de 5 anos). Perfil polícia com a flag ligada: aceite formal do R78, com validade curta |

### Novas desta revisão

| # | Pergunta | Recomendação |
|---|---|---|
| Q1 | Nível 2 da ADR-0008: a sala de origem precisa conceder, por assinatura do administrador e visível aos membros, a entrada no escopo de análise? | Sim (seção 5.2) |
| Q2 | Destruição do nó de campo: "selar antes de destruir", ou exceção à GD-13 aceita por escrito? | Selar antes de destruir (seção 6) |
| Q3 | Suíte candidata do spike: P-256 enquanto o HSM não é escolhido? | Sim (01e-E04b) |
| Q4 | Confirmação de época por membro (uma mensagem curta por commit), para tornar visível o commit retido? | Sim, medir no spike e em banda estreita (01e-E03) |
| Q5 | Selagem na revogação na Fase 2, com a chave pública lida do pacote e guarda no `directory` até a Fase 4? | Sim (01e-E13) |

## 10. A cadeia pode seguir para G0?

- **01c-C01 e 01c-C02 estão fechados.** As ADRs 0003 a 0006 e as regras 2 e 14 não têm crítico nem alto aberto. Os médios 01e-E02 a E06, E08, E09 e E13 são correções de texto ou de critério, sem decisão de fundo.
- **Pelo critério literal da 1b há um alto aberto: 01e-E01.** Ele está só na ADR-0008, que é proposta nova, para depois da Fase 6. Nada entre G0 e G3 depende dela: o spike MLS, a CI, a atestação e a inferência não tocam o nível 2.
- **Recomendo seguir para G0 com uma destas duas condições**, à escolha do responsável, sem terceira rodada completa:
  1. o status da ADR-0008 e a linha dela no README passam a dizer "nível 1a proposto; nível 1b fora; **nível 2 volta ao desenho** (01e-E01)"; ou
  2. o responsável registra por escrito, no `GD-respostas.md`, que a ADR-0008 fica fora do critério da 1b até ser redesenhada.
- **Antes da etapa 3** (não antes de G0): aplicar 01e-E04 ao critério do spike e responder N2 e Q3.

## 11. Próximo passo sugerido

1. Responsável: Q1, Q2 e N1 (definem as ADRs 0008 e 0004); N2 e Q3 (definem o spike); N3 e pergunta 7 (regra 14 e nomes para o aceite).
2. Passe curto do Opus, só texto: 01e-E02, E03, E05, E06, E08, E09, E13, E17, E20, E21 e o critério do spike (E04). 01e-E07 e E12 junto com 01c-C17, antes de aceitar a ADR-0004. 01e-E10, E11, E14 e E18 quando as ADRs 0008 e 0009 forem redesenhadas.
3. Orquestrador: registrar esta revisão no `PIPELINE.md` e a condição escolhida da seção 10; depois, G0.
