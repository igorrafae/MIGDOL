# Correções da revisão — Fase 0, etapa 1b, terceira rodada

**Modelo:** Claude Opus 5.5 (pipeline da Fase 0, etapa 1b, rodada 3). **Revisão seguinte:** Fable 5.2.
**Entradas:** `GD-respostas.md` (GD-15 a GD-20 prevalecem), `01e-revisao-rodada2.md` (achados 01e-E01 a E21, recomendações e perguntas), `01d-correcoes-rodada2.md`, ADRs 0003 a 0009 e README, docs 01 a 10 nas seções afetadas, termo de ciência.
**Limites:** nenhum código, commit ou push; nada enviado para fora. Fato jurídico não conferido segue **[conferir]**. Nada de cripto à mão: o texto só descreve recursos da biblioteca MLS, da pilha HPKE e do hardware; o que depende de verificação virou critério do spike. Nenhum risco é aceito aqui: GD é decisão do responsável pelo produto, não aceite formal (doc 04 seção 3.2). `docs/11-*` não foi editado (outro agente); só é citado. Única regra do `CLAUDE.md` alterada: a 4 (GD-19).
**Convenção:** achados do 01e citados como `01e-Enn`. Perguntas novas desta rodada: `R3-n`.

## Resumo

| Severidade | Total | Fechado | Pendente de resposta (em parte fechado) | Adiado |
|---|---|---|---|---|
| Alta | 1 | 1 (E01) | 0 | 0 |
| Média | 12 | 11 | 1 (E03) | 0 |
| Baixa | 8 | 7 | 1 (E18) | 0 |
| **Total** | **21** | **19** | **2** | **0** |

- **Altos e críticos do 01e abertos: nenhum.** O 01e-E01 está fechado **no texto** pela GD-15 e pelo redesenho do nível 2 da ADR-0008; depende da revisão do Fable.
- "Fechado" quer dizer **fechado no texto**. E04, E13, E15 e E16 dependem de prova no spike ou na Fase 2; E06 e E09 têm valores em decisões já abertas (D-R35, D-R36).
- **ADR-0010 é nova** (investigação sigilosa): precisa de revisão do Fable antes de qualquer aceite. Não achei conflito real com a regra 14 (seção 4.2); registrei duas leituras como pergunta.
- **Regra 12:** o reconhecimento facial da GD-18 cabe na letra da regra; dois pontos que ela não cobre viraram pergunta (R3-3), sem mudar a regra.

## 1. Achados 01e-E01 a E21

| ID | Sev. | Status | Onde | Nota |
|---|---|---|---|---|
| E01 | Alta | **Fechado no texto** (revisar) | ADR-0008 (status, contexto, §1, §2, §4, §5, alternativas, consequências); ADR-0009 §1, §3; doc 08 §5.1; doc 01 RN09, atores, RF85–RF87, RF92; doc 02 §9.2, §11; doc 06 §3, §7 item 3; termo item 3; doc 04 §5.4; `admin-panel/CLAUDE.md` | GD-15: nível 2 só por **concessão** da sala de origem (mensagem de controle assinada pelo administrador, com CI e/ou autorização do comandante, visível aos membros, registrada e retirável); nó de IA só cruza sala com concessão vista no grupo; sem concessão, só o rito da regra 14; **regra de tipo** (sala de análise sob custódia só com salas sob custódia); **expurgo em cascata**; saída estruturada; não retroage e agregados também exigem concessão **[proposta; D-R41]**; limite declarado (só salas ativas). Nível 1b fora |
| E02 | Média | **Fechado** | ADR-0005 item 2c, 2d; doc 02 §5.5; doc 07 §5 | Tabela de classes: autenticado (remetente inconsistente, SEV2), forte, fraco, não conta (duplicata, fora de janela, corrupção, retransmissão, com erro próprio da biblioteca **[verificar no spike]**). SEV1 só com sinal forte repetido ou dois sinais distintos; limite de alertas por sala e época |
| E03 | Média | **Pendente de resposta** (em parte fechado) | ADR-0005 item 2b (sinais 4 e 5), 2e, 2f, Consequências; doc 02 §5.5; doc 04 R98 e §5.4; doc 03 R11 | Commit retido declarado como caso normal de servidor hostil; sinais 4 e 5; **confirmação de época** como **[proposta]**, medida no spike (pergunta Q4 do 01e, ainda sem resposta); R98 → **Médio** até a Fase 1; R11 mantido em Baixo, com nota (commit retido não insere membro) |
| E04 | Média | **Fechado** (prova no spike) | doc 03 Fase 0 (spike MLS e spike de atestação); `PIPELINE.md` etapas 3 e 5 e parágrafo do critério | (a) Keystore em duas partes: caminho em emulador na etapa 3 (7a), hardware real na etapa 5 (7b); (b) suíte candidata P-256 (GD-20); (c) CLI com no mínimo quatro clientes, sem React, igual nos dois textos; (d) matriz de falso positivo, sinal forjado sem SEV1 sozinho, commit retido e confirmação de época, membro que mente, API da biblioteca (`authenticated_data`, hash do transcript, `epoch_authenticator`), commit público |
| E05 | Média | **Fechado** | ADR-0006 item 8; doc 04 R110 | Sela **todas** as gerações guardadas (`n` e `n+1`), com o número da geração; `n+1` idempotente (item 3) |
| E06 | Média | **Fechado** (valores em D-R35) | ADR-0006 item 8 (prazo, painel); ADR-0003 §4; doc 06 §8.2; doc 04 §6.4, R110 | Padrão e piso no `instituicao.yaml` (proposta do 01e: 5 anos e 90 dias **[conferir]**); abaixo do padrão, só com as duas autorizações; selado na chave do **período de vencimento**; prorrogação = ordem de preservação; painel distingue "registro excluído" de "chave eliminada" |
| E07 | Média | **Fechado** pela GD-18 | ADR-0004 item 7 e item 13 (novo), Alternativas, Consequências; doc 10 §1 nota, §3.1, §4, §5, §7, R89, R118, D-R39, RF97, RNF-R2, F-R13, F-R14; doc 04 §8 (auditoria), §5.4; doc 05 §9 item 7 | **Exceção escrita à GD-13**, só para o nó: destruição imediata, um ato ou automática; **perda aceita** registrada (cache de missão de quem reiniciar, custódia e trilha não sincronizadas, selados de removidos, lista de revogação); gatilhos **[proposta; D-R39]**; mitigações sem mudar a decisão; "selar antes de destruir" registrado como alternativa não adotada (R3-6) |
| E08 | Média | **Fechado** | ADR-0006 item 8 (revogação em lote, rastro do selado, volta); doc 04 R110 | (a) Revogação em lote com duas autorizações; na dúvida, suspender; aparelho achado volta por revinculação e `history_share`. (b) Hash do selado na trilha, conferência pelo Auditor, RPO 0 |
| E09 | Média | **Fechado** (b com D-R36) | ADR-0005 item 4 (condição do commit público); ADR-0003 §1 (motivo); doc 04 R111, D-R36; doc 07 D-R36 | (a) Commit público só se nenhum dado sensível transitar em proposta. (b) **[proposta]** categoria do motivo em lista fechada na trilha; texto livre só no grupo, sem hash |
| E10 | Média | **Fechado** | ADR-0009 §3 (raiz de confiança, papéis privilegiados, detentores); ADR-0003 §1; doc 08 §5.1; doc 04 §6.2 | Atribuição de papel privilegiado como **objeto assinado** pelas duas pessoas, com chaves da lista assinada em cerimônia; segunda assinatura da sala comum e concessão conferidas contra essa lista; combinações com detentor conferidas na cerimônia de emissão, identidade dos detentores fora do diretório em claro; atribuição dupla só para papéis privilegiados (RN05 e D20 para o agente) |
| E11 | Média | **Fechado** | ADR-0008 §2 e §5; doc 01 RF87; doc 04 R104, R106; doc 07 A99 | Saída estruturada validada no código; texto livre com limite e verificação de sobreposição; teste de injeção dentro do escopo; limite de perguntas e revisão amostral; R104 → **Médio** até o teste |
| E12 | Média | **Fechado** | ADR-0006 item 9; ADR-0004 item 7; doc 10 §4; doc 10 F-R15 | Passagem do lease pelo protocolo de duas fases nos dois sentidos; o nó apaga só depois da confirmação; aparelho que reinicia com o nó destruído: perda aceita (GD-18) |
| E13 | Média | **Fechado no texto** (proposta; confirmar Q5) | ADR-0006 item 11 (novo); doc 02 §7.1; doc 03 Fase 2 | Fase 2: `directory` sela com a chave pública lida do pacote, em tabela própria, chave de teste até a cerimônia; Fase 4: `custody` busca no `directory` (sentido permitido) ou a tabela fica lá; sem ciclo (regra 9) |
| E14 | Baixa | **Fechado** | ADR-0008 §1 (nível 1a); doc 01 RF85; doc 04 R105; `admin-panel/CLAUDE.md` | Contagens arredondadas ou em faixas, sem filtro abaixo do limite; lista fechada no `instituicao.yaml`; missões ativas e nós de campo fora do painel no perfil Exército até D-R20 |
| E15 | Baixa | **Fechado** (prova no spike) | doc 02 §4.1, §6.3; doc 03 Fase 0 item 7a | Zerar na JVM é melhor esforço; buffer direto alocado pelo core **[verificar]**; JVM no limite conhecido. A regra 2 não mudou |
| E16 | Baixa | **Fechado** (verificar no spike) | doc 01 RF04; ADR-0005 item 2e; doc 03 Fase 0 item 9 | Código de segurança por sala e época, derivado do `epoch_authenticator` |
| E17 | Baixa | **Fechado** | ADR-0003 §3 | R97 (residual Baixo): aceite do dono do ativo (`custody`), revisão anual |
| E18 | Baixa | **Pendente de resposta** (em parte fechado) | ADR-0009 §1, §3; doc 04 §6.2 | (a) Gestor: "no painel de indicadores". (c) Analista × gestor proibido. (b) Papéis acrescentados (segunda assinatura, autorizadores, criador da sala de análise, plantão, papéis da ADR-0010); **quem ocupa** cada um depende de D-R33, D-R36, D-R37 e da pergunta 8 do 01c |
| E19 | Baixa | **Fechado** | `PIPELINE.md` etapa 5b (nova) e cadeia; ADR-0007 item 1 e Consequências; doc 03 Fase 0 (spike sem React; estudo do iPhone); ADR-0009 status ("decidir antes da Fase 1") | Etapa do estudo do iPhone com saída, modelo e critério (APNs, App Attest, MDM, captura de tela, série atestada, segundo plano; protótipo do doc 11) |
| E20 | Baixa | **Fechado** | doc 02 §6.4; doc 06 §6.3; doc 05 P11 | "Salvo o cache de aparelho revogado em poder da instituição, pelo rito da regra 14, dentro do prazo do segredo guardado" |
| E21 | Baixa | **Fechado** | (a) doc 06 §3; (b) ADR-0008 §4; (c) ADR-0006 item 8; (d) ADR-0006 Consequências; (e) ADR-0003 §2, doc 02 §5.2, doc 07 §7.3, doc 04 §6.4; (f) `PIPELINE.md` estado | (a) Linhas do nível 2, da concessão e do registro de perguntas; (b) relatório da operação só ao administrador (RN09), respostas da sala de análise aos analistas; (c) "depois que o aparelho esquece o segredo"; (d) R110 citado; (e) período **só por tempo** como hipótese do spike; (f) linha da 1b rodada 3 |

## 2. Decisões GD-15 a GD-20 aplicadas

| Decisão | O que mudou | Onde |
|---|---|---|
| **GD-15** (nível 2 por concessão; regra 3 sem mudança) | Painel só com o nível 1a; nível 1b fora; agregados derivados e nominais só por MLS, de salas que concederam; concessão assinada no grupo de origem, com CI e/ou autorização do comandante, registrada | ADR-0008; ADR-0009 §1; doc 01 RN09, atores, RF85–RF87, RF92; doc 02 §9.2, §11; doc 08 §1, §5.1; doc 06 §3, §7; termo item 3; doc 04 R103–R107, D-R41; `admin-panel/CLAUDE.md`; doc 07 A96–A99. **Regra 3 não foi tocada** |
| **GD-16** (investigação sigilosa) | ADR-0010 nova; playbook P1-S; impedimento; registro selado; análise só na estação isolada; relatório de rastro; revisão humana | ADR-0010; ADR-0003 §4, §5 item 5, §7; doc 05 princípios 5 e 6, §3.4 item 6, §3.5, §5.4 (P1, P1-S), §8, §9; doc 01 atores, RF93, RF94, §11; doc 02 §1, §6.4, §7.1, §11; doc 03 §1.3, §4.9, Fase 4, depois da v1; doc 04 R112–R115, D-R37, D-R38, §6.2; doc 06 §3, §5.2, §6.1, §7 item 13, §8.2, §8.3, §12, §14; doc 07 A84–A95; termo item 14; ADR-0009 §1 |
| **GD-17** (preservação; sem gravação oculta) — *pendente de confirmação do responsável* | Preservação selada das cópias de custódia e dos registros selados do escopo, com o `custody` apagando cifrados só com recibo de eliminação; retratos de metadados cifrados para o quórum; **sem** membro oculto, sem troca de marca, sem uso do nó de IA, sem suspender expurgo nos aparelhos | ADR-0010 §7, §8; ADR-0003 §7; doc 06 §8.2, §8.3; doc 04 R119; doc 07 A87–A89, A95 |
| **GD-18** (destruição na captura; verificação facial) | Nó capturado: destruição imediata como **exceção escrita à GD-13**, com a perda aceita registrada. Celular: verificação facial de quem acessa, no aparelho, com vivacidade, gabarito cifrado pelo hardware seguro; a face nunca decide sozinha; escada por perfil (missão de campo destrói; urbano suspende e revoga, D-R14). D-R8 fechada nesse sentido, falta o parecer | ADR-0004 item 13; doc 10 (§1, §3.1, §3.2, §4, §5, §7, R89, R118, D-R39, RF97, RNF-R2, F-R13–F-R15); doc 07 §6.1, §6.3, §6.4, §7, §7.1, **§7.1a (nova)**, §7.2 (reescrita), R66, R67, R116, R117, A51, A76–A83, RF95, RF96, RNF-G2, RNF-G3, D-R8, D-R40, §14, §15; doc 09 §5; doc 04 §4, §5.4, §8, §10; doc 06 §3, §5.2, §7 itens 5 e 12, §8.2, §12, §14; termo itens 6 e 13 |
| **GD-19** (ADR-0007 aceita; doc 11) | ADR-0007 **aceita**; regra 4 alterada; diretórios **não** renomeados (Fase 1); notas "proposta" trocadas; estudo do iPhone como etapa 5b | ADR-0007; `CLAUDE.md` regra 4; `docs/adr/README.md`; doc 02 §1, §2, §3, §7.2, §8.1, §11; doc 01 D1; doc 08 §1; `app/CLAUDE.md`; `PIPELINE.md`; doc 03 Fase 0. O doc 11 é citado, não editado |
| **GD-20** (P-256) | Suíte candidata do spike: DHKEM(P-256, HKDF-SHA256), HKDF-SHA256, AES-256-GCM, até a escolha do HSM | doc 03 Fase 0 item 1; `PIPELINE.md`; ADR-0003 §2 |

## 3. Mudança no `CLAUDE.md`: regra 4 (GD-19)

Antes:
> 4. app/src não usa localStorage, sessionStorage, IndexedDB nem fetch direto.
>    Todo acesso a dados e rede passa pelo plugin que fala com o core.

Depois (texto da ADR-0007 item 5, sem alteração):
> 4. A interface (Compose, SwiftUI e, enquanto existir, app/src) não grava dado fora do core (sem
>    localStorage, sessionStorage, IndexedDB, SharedPreferences, DataStore, arquivos próprios,
>    UserDefaults nem Core Data) e não abre rede própria (sem fetch, OkHttp nem URLSession).
>    Todo acesso a dados e rede passa pelo core.

O Fable sugeriu (01e §5.1) a regra em termos gerais e a lista de APIs só no teste estático. Não apliquei: o pedido era o texto da ADR. Fica como sugestão (R3-10). Nenhuma outra regra foi alterada.

## 4. ADR-0010 resumida

### 4.1 O que decide
| Ponto | Decisão (proposta) |
|---|---|
| Iniciador | Qualquer detentor de cartão ou a corregedoria |
| Autorização | CI (tramita como sigilosa **[conferir]**) **e** segunda autorização independente (D-R37), assinadas **presencialmente na estação isolada**, com chaves da lista assinada de papéis |
| Impedimento | Três fontes: lista declarada no pedido (por chave, nunca por nome no servidor); posição, calculada pela estação (membros, administradores e quem atribuiu papéis nas salas do escopo no período); declaração de cada participante. A estação recusa ato com chave impedida. Quórum exige N ≥ M + impedidos; senão, detentores de reserva (D-R37) ou só nível 1a |
| Níveis | **1a** metadados em claro (CI + segunda autorização); **1b** registros selados de atribuição (com a base da regra 14 e quórum); **2** conteúdo (com ordem judicial e quórum) |
| Registro selado | Na classe "registros", série própria; trilha só com hash de tamanho fixo, igual aos demais selados; aviso ao Auditor dentro do registro (exceção à ADR-0003 §5 item 5) |
| Onde roda | Só na estação isolada da ADR-0003 §6, em sessão de cerimônia, sobre retrato do backup de rotina (sem consulta ao servidor vivo); não é o nó de IA; ninguém lê conversa |
| Sem gravação oculta | Sem membro oculto, sem troca da marca, sem nó de IA, sem "modo investigação"; gravação para frente só onde já há custódia |
| Preservação | Suspende eliminação e prazo das cópias de custódia e dos registros selados do escopo; `custody` só apaga com recibo de eliminação; retratos de metadados cifrados para o quórum; **não** suspende o expurgo nos aparelhos nem no nó de IA |
| Saída | Relatório de rastro estruturado, sem texto de conversa, com grau de confiança do doc 05; indício de conteúdo rotulado e sem elevar o grau sozinho; conteúdo para o processo só por ordem judicial e rito comum (RF67) |
| Revisão (regra 11) | Duas pessoas não impedidas assinam; sem as duas, não há pacote; nenhuma medida automática; MP (MPM no Exército) pede cautelares; contraditório no procedimento |
| Controle do rito | Prazo de 90 dias renovável com as duas autorizações; revisão por quórum a cada 6 meses; estatística ao comitê **[proposta; D-R38]** |
| Fase | Rito na Fase 4 (com a custódia e a estação); análise automatizada depois da Fase 6 |

### 4.2 Regra 14: há conflito?
**Não encontrei conflito real.** A GD-16 diz quem **inicia** e quem **autoriza o início**; não dispensa o quórum. Por isso:
- todo acesso a dado cifrado (registros selados, cópia de custódia, cache com segredo guardado) continua com **quórum presencial** de no mínimo 3 detentores, em HSM offline, com base registrada, e ninguém decifra sozinho;
- o nível 2 exige **ordem judicial**, mais restritivo que a regra 14 (que também admite processo administrativo);
- o nível 1 da GD-16 ("sem ordem judicial, só metadados e sinais de atribuição") foi dividido: os sinais que já são dado cifrado (registro de visualização, de acessos, evidência, matrícula) ficam no **nível 1b** e exigem a **base da regra 14** (processo administrativo com validação jurídica, ou ordem judicial). Sem essa base, a investigação fica no nível 1a. A GD-16 não nomeia base para o nível 1 além da CI: esta é a leitura mais restritiva, e vai como pergunta (R3-1);
- o registro selado entra na classe "registros", que já está na regra 14: nenhuma classe nova fora da regra;
- a assinatura do pedido é presencial, coerente com "não há aprovação remota por quórum".

Duas leituras que podem tocar a regra 14 e ficaram como pergunta, sem mudar a regra:
- **Preservação × "sala comum é expurgada ao concluir"** (R3-2): a preservação não suspende o expurgo do conteúdo nos aparelhos; guarda só metadados (retrato do backup), cifrados para o quórum. Se o responsável entender que metadados de sala comum também precisam sumir na conclusão, o nível 1a de sala comum concluída deixa de existir.
- **GD-17 pendente de confirmação** (R3-5).

## 5. IDs novos (conferidos por busca, sem colisão)

| Tipo | IDs | Onde |
|---|---|---|
| Riscos | R112–R115, R119 | doc 04 §5.4 |
| | R116, R117 | doc 07 §10 |
| | R118 | doc 10 §8 |
| Requisitos | RF92, RF93, RF94 | doc 01 §6.6 |
| | RF95, RF96, RNF-G3 | doc 07 §12 |
| | RF97 | doc 10 §10 |
| | RL13, RL14 | doc 06 §12 |
| Testes | A76–A83 (face e escada), A84–A95 (investigação sigilosa), A96–A99 (análise entre operações) | doc 07 §11 |
| | F-R13, F-R14, F-R15 | doc 10 §11 |
| Decisões | D-R37, D-R38, D-R41 | doc 04 §4 |
| | D-R39 | doc 10 §9 e doc 04 §4 |
| | D-R40 | doc 07 §13 e doc 04 §4 |
| Playbook | P1-S (não P12, que o doc 11 já usa para tela do painel) | doc 05 §5.4 |
| Seção nova | doc 04 §5.4; doc 07 §7.1a | — |

**Para o orquestrador:** o doc 11 deixou "ID a atribuir" para riscos (R-U1 a R-U17) e decisões (U-1 a U-7). Numerar a partir de **R120** e **D-R42**. R-U4 (falso negativo facial) e R-U5 (falso positivo facial e destruição) se sobrepõem a **R67** e **R116**: convém fundir em vez de criar IDs novos. O limite de falso negativo facial (U-3) está dentro da D-R40.

## 6. Pendências para o Fable

1. **ADR-0010 inteira**, em especial: (a) as três fontes de impedimento e a conferência por chave na estação; (b) a divisão 1a/1b/2 e a leitura da regra 14 (seção 4.2); (c) o retrato do backup de rotina como entrada (é suficiente para não alertar um administrador investigado?); (d) o registro de tamanho fixo na trilha e o vazamento por período não eliminado (R113); (e) se a estação isolada pode carregar modelo de inferência sem virar terceiro lugar com texto em claro.
2. **ADR-0008 nível 2 redesenhado (01e-E01):** concessão no grupo de origem, regra de tipo, expurgo em cascata feito pelo nó de IA (retirada assinada por credencial de serviço), agregados também sob concessão.
3. **ADR-0004 item 13:** a exceção à GD-13, os gatilhos propostos e as mitigações (sobretudo o aparelho guardar a própria cópia de custódia até a central confirmar).
4. **Doc 07 §7.1a e §7.2:** a matriz por perfil; "inconclusivo" que nunca destrói (e o desvio de quem cobre o rosto); o gabarito "no hardware seguro" lido como cifrado por chave do hardware, com comparação no core (R3-4).
5. **ADR-0005 item 2c e 2f:** a tabela de classes e a confirmação de época.
6. **ADR-0006 itens 9 e 11:** passagem do lease nos dois sentidos e o dono dos selados antes da Fase 4.
7. **R98 e R104 em Médio:** conferir se o residual certo é esse até os testes.
8. **01c-C21** (renomear D01–D34) segue adiado para antes do portão GC.

## 7. Perguntas ao responsável

### Novas desta rodada
| # | Pergunta | Recomendação |
|---|---|---|
| R3-1 | ADR-0010: sinais de atribuição que são dado cifrado (registro de visualização, de acessos, evidência, matrícula) exigem, sem ordem judicial, a base da regra 14 (processo administrativo com validação jurídica)? | Sim (aplicado; mais restritivo) |
| R3-2 | Preservação: confirmar que ela **não** suspende o expurgo nos aparelhos e no nó de IA, e que guardar retrato de metadados de sala comum não conflita com "sala comum é expurgada ao concluir" | Confirmar as duas (aplicado como leitura) |
| R3-3 | Regra 12: o gabarito facial gravado (derivado de quadro) e a câmera na tela de desbloqueio estão dentro da regra? | Escrever na regra que cobre também o gabarito e o desbloqueio, ou declarar fora; até lá, leitura restritiva aplicada |
| R3-4 | "Modelo do rosto no hardware seguro": aceitar gabarito cifrado por chave do hardware, com comparação no core, ou exigir comparação dentro do TEE (só com recurso do fabricante, o que restringe modelos)? | Aceitar, com o limite declarado |
| R3-5 | Confirmar a GD-17 | — |
| R3-6 | Nó de campo: adotar como mitigação, sem mudar a GD-18, o aparelho guardar a cópia de custódia que gerou até a central confirmar, e/ou o "selar antes de destruir" do 01e-E07? | Adotar a primeira (barata, sem leitura nova) |
| R3-7 | D-R41: base mínima da concessão por perfil; concessão não retroage; agregados também sob concessão; aprovação dupla para criar a sala de análise | Sim nos três últimos (aplicados como proposta) |
| R3-8 | D-R40: um só PIN errado com rosto de outra pessoa basta para destruir no perfil missão? A polícia liga o perfil missão nas missões de campo dela? | Manter um (GD-18) só com taxa de falso negativo medida; polícia desligada por padrão |
| R3-9 | D-R37 e D-R38: quem dá a segunda autorização; substitutos; detentores de reserva; prazo; aviso ao investigado no arquivamento | — |
| R3-10 | Regra 4 em termos gerais, com a lista de APIs só no teste (sugestão do Fable) | Opcional |

### Do 01e, ainda sem resposta
| # | Pergunta |
|---|---|
| Q4 | Confirmação de época por membro (01e-E03): aplicada como proposta, a medir no spike |
| Q5 | Selagem na Fase 2 pelo `directory` (01e-E13): aplicada como proposta |
| N3 | Redação da regra 14 conforme 01c-C23 (CI sempre; base ordem judicial ou processo administrativo) |
| N4 | D-R36: quem dá a segunda assinatura; motivo em lista fechada (proposta aplicada) |
| N5, N6 | Posição em missão na cópia de custódia; RPO 0 e teto de 12 meses |

### Do 01c e da etapa 1, ainda abertas
2 (D-R30 e frase de escopo da regra 12, agora junto com R3-3), 3 (troca de titular de papel), 4 (regra 7, antes da Fase 5), 5 (kill-switch), 6 (D-R35), 7 (responsável pela homologação e comitê; aceites de R68, R97, R66, R118), 8 (sala de resposta), 9 (administrador de toda operação), 10 (01c-C17: retorno com nó destruído, agora mais urgente com a GD-18), 11 (HSM), 12 (índice em HSM), 13 (custódia antes do KEM híbrido); da etapa 1: 4 (D-R24), 5 (D-R14), 6 (D-R5), 8 (D28), 15 (IA em sala secreta).

**Jurídico [conferir]:** base do nível 1b sem ordem judicial; CI sigilosa; limite de acesso do investigado; contraditório e aviso no arquivamento; MP ou MPM; Tema 977; biometria (art. 11, ANPD) e câmera no desbloqueio; destruição automática em missão de campo; preservação de metadados de sala comum; os pontos anteriores (processo administrativo como base, corregedoria, prazo de 5 anos, autorização arquivística, dados de terceiros do nível 2).

## 8. Arquivos alterados nesta rodada

| Arquivo | O que mudou |
|---|---|
| `CLAUDE.md` | Regra 4 (GD-19) |
| `app/CLAUDE.md`, `admin-panel/CLAUDE.md` | Notas: ADR-0007 aceita; nível 1b fora e nível 2 por concessão |
| `docs/adr/0010-investigacao-sigilosa-de-vazamento.md` | **Novo** |
| `docs/adr/0003` | E06, E09b, E10a, E17, E21e; GD-20; série de investigação, exceção do §5 item 5, preservação selada |
| `docs/adr/0004` | Item 13 (GD-18), passagem do lease, alternativas, consequências |
| `docs/adr/0005` | E02, E03, E09a, E16 |
| `docs/adr/0006` | E05, E06, E08, E12, E13 (item 11), E21c, E21d |
| `docs/adr/0007` | Aceita (GD-19); regra 4 aplicada; estudo do iPhone; renomeação na Fase 1 |
| `docs/adr/0008` | Reescrita: GD-15, E01, E11, E14, E21b |
| `docs/adr/0009` | E10, E18, concessão, papéis da ADR-0010, decidir antes da Fase 1 |
| `docs/adr/README.md` | Estado das ADRs 0003–0010 |
| docs 01 a 10 | Conforme as seções 1 e 2 |
| `docs/anexos/termo-de-ciencia-e-autorizacao.md` | Itens 3, 6, 13 e 14; autorizações |
| `docs/fase0/PIPELINE.md` | Etapas 3, 5 e 5b; cadeia; critério; estado |
| `docs/fase0/01f-correcoes-rodada3.md` | **Novo** (este arquivo) |
