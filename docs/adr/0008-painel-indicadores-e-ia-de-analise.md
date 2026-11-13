# ADR-0008 — Painel de gestão, indicadores e IA de análise

**Status:** proposta (decisão do responsável pelo produto, 12/11/2026; integrada na segunda rodada da etapa 1b, `docs/fase0/01d-correcoes-rodada2.md`; **redesenhada na terceira rodada**, `docs/fase0/01f-correcoes-rodada3.md`, pela GD-15 e pelo 01e-E01) · **Nível 1a** proposto (aprovado pelo Fable no 01e) · **Nível 1b fora** (GD-15: a regra 3 não muda; o painel mostra só o nível 1a) · **Nível 2** redesenhado: só por MLS e só de salas que **concederam** entrada no escopo (GD-15), com regra de tipo, expurgo em cascata e saída estruturada; aguarda revisão do Fable · **Complementa** ADR-02 (nó de IA como membro), doc 02 seção 9, doc 08 seção 5.1 (compartimento de análise), ADR-0009 (papéis) · **Contraria** doc 01 RN09 (relatório de IA só para o administrador da operação) na sala de análise · **Fase:** depois da v1 (após a Fase 6), condicionada ao spike de inferência (etapa 6)

## Contexto
A chefia quer um painel com KPIs, uma barra de perguntas à IA local sobre riscos das operações, vínculos entre pessoas e presos em operações diferentes, padrões semelhantes e sugestões para mitigar riscos. O perfil inteligência já tem acesso a **indicadores para planejar operações, não às conversas**. O painel web roda no servidor, que nunca vê texto em claro (regra 3). A posição guardada dos agentes é dado sob quórum (regra 14). Nenhuma decisão contra pessoa sai de IA (regra 11).

A revisão do Fable (01e-E01, alta) mostrou que o nível 2 da versão anterior entregava a quem não é membro um conhecimento nominal derivado do conteúdo, por um membro que a instituição controla (o nó de IA), **sem ato da sala de origem**. Isso contrariava a GD-1 ("sala comum é MLS puro e expurgada, sem cópia"), levava indícios de salas comuns para o acervo de custódia da sala de análise, escapava do expurgo da origem (RN04) e deixava a chefia sozinha decidir o escopo.

Decisão do responsável (**GD-15**): o analista lê indicadores nominais por **concessão** da sala de origem, com CI e/ou autorização do comandante; a concessão é registrada. A regra 3 **não muda**: o painel mostra só o nível 1a; agregados derivados de conteúdo vão ao app por MLS.

## Decisão (proposta)

### 1. Níveis de indicador
| Nível | Conteúdo | Onde aparece | Proteção |
|---|---|---|---|
| **1a. Agregado de metadados** | Contagens e prazos que o servidor já tem para funcionar, de uma **lista fechada** no `instituicao.yaml`: aparelhos por estado, revinculações, alertas de integridade, salas por estado do ciclo (RN18), duração das salas, prazos de custódia | Painel web (rede interna, token de hardware) | Calculado pelo servidor a partir dos próprios metadados; sem nome, sala ou operação identificável; contagens **arredondadas ou em faixas**, sem filtro (unidade, período) que reduza o grupo abaixo do limite mínimo (D-R32), para que a diferença entre duas consultas não revele um caso (01e-E14); datas por semana ou mês. "Missões ativas" e nós de campo: no perfil Exército, só com D-R20, e fora do painel até lá. **Não** precisa de exceção à regra 3 |
| **1b. Agregado derivado no painel** | (era: operações por tipo e região; tendências; riscos por categoria) | **Não existe** (GD-15) | Tipo e região não vão ao servidor: continuam só no estado cifrado do grupo. Riscos por categoria derivam do conteúdo e não vão ao painel. Se a chefia precisar deles, vão ao app por MLS, como nível 2 sem nomes |
| **2. Por MLS, no app** | **2-agregado:** riscos por categoria, tendências, sem nomes. **2-nominal:** pessoas e presos em comum entre operações, vínculos, padrões semelhantes, sugestões de mitigação | Só no app do analista de inteligência, como membro da sala de análise, por MLS. **Nunca** no painel | Só de salas com **concessão vigente** (seção 2); só para quem tem credencial em todo o escopo (doc 08 seção 5.1); o servidor nunca vê |

### 2. Análise entre operações
- **Sala de análise:** sala MLS criada pela chefia (quem cria: D-R32), com **escopo**, **prazo** e **revisão**. É um **compartimento de análise** (doc 08 seção 5.1), com credencial própria. Proposta: a criação exige **aprovação dupla** da chefia, como objeto assinado conferido contra a lista assinada de papéis (ADR-0009 seção 3; 01e-E10) **[proposta; D-R41]**.
- **Concessão da sala de origem (GD-15):** uma sala só entra no escopo se houver **concessão**: mensagem de controle assinada **dentro do grupo de origem** pelo administrador da sala (em missão, o comandante da missão), com:
  - a base: número da **CI** e/ou **autorização do comandante** (GD-15; qual é o mínimo por perfil: D-R41);
  - o identificador opaco da sala de análise, a finalidade, o prazo (nunca maior que o do escopo) e o tipo de indicador concedido (2-agregado, 2-nominal ou os dois).
- A concessão é **registrada** (hash na trilha de auditoria; conteúdo no estado autenticado do grupo, cifrado) e **visível a todos os membros** da sala de origem, como evento na conversa. O administrador pode **retirá-la** a qualquer momento, pelo mesmo meio.
- O nó de IA só cruza sala cuja concessão vigente ele **viu no próprio grupo**, conferida contra as chaves de administrador do estado autenticado (ADR-0005). O servidor e a chefia sozinhos não alargam o escopo.
- **Sem concessão, a sala fica fora.** A instituição só chega ao conteúdo dela pelo rito da regra 14 (custódia, quórum presencial, base registrada) ou, na investigação sigilosa, pela ADR-0010.
- **Agregados também exigem concessão** (2-agregado): eles também derivam do conteúdo. Leitura mais restritiva da GD-15, que fala só dos nominais; dúvida registrada em D-R41.
- **Não retroage** **[proposta; D-R41]**: a concessão vale para o que for dito depois dela, como a flag de custódia (GD-2).
- **Regra de tipo:** sala de análise **sob custódia** só inclui salas sob custódia. **Sala comum nunca alimenta acervo de custódia.** Sala de análise comum (flag desligada, ou exceção da GD-14) pode incluir salas comuns e sob custódia; seus indícios são expurgados com ela.
- **Expurgo em cascata:** cada indicador leva o identificador opaco da sala de origem. Conclusão, expurgo ou retirada da concessão na origem fazem o nó de IA apagar os indícios dela do próprio índice e emitir, na sala de análise, mensagem de controle de **retirada**; os apps dos analistas apagam os indicadores daquela origem. O prazo do indicador nunca é maior que o da origem nem que o do escopo. Em sala de análise sob custódia, o que já entrou na cópia de custódia segue o prazo dela, e isso está no termo (item 3).
- O nó de IA só cruza salas de que **é membro** (ADR-02). Sala sem IA (por decisão da sala ou do grau; pergunta 15 da revisão) fica fora de qualquer escopo.
- O **filtro de permissão vem antes da busca**, no código, nunca por instrução ao modelo: um índice por sala (doc 02 seção 9.4) e montagem do contexto só com salas do escopo com concessão vigente.
- **Saída estruturada, validada no código (01e-E11):** entidade, tipo de vínculo, referência (sala opaca e data), grau de confiança. Texto livre do modelo só como explicação, com limite de tamanho e verificação de sobreposição com a fonte, que bloqueia trecho de conversa copiado. Ler a origem exige ser membro da sala ou pedir pelo rito.
- Toda resposta é rotulada **"indício para análise humana"**, cita a fonte e nunca vira decisão automática (regra 11). Relatório herda o grau mais alto das salas analisadas (doc 04 seção 6.1). Até o portão G3, nenhuma sala de análise trata grau classificado (GD-9).
- **Registro e limite de perguntas:** toda pergunta fica registrada (quem, quando, escopo), cifrada para a chave de quórum de registros; limite de perguntas por analista e por período no `instituicao.yaml`; uma segunda pessoa revê, por amostragem, o que foi perguntado (R42).
- A sala de análise segue a regra de custódia da instituição: com a flag ligada, nasce sob custódia (GD-14), e as respostas entram na cópia de custódia, com o prazo dela (sujeito à regra de tipo acima).
- **Transparência:** os membros das salas sabem que a sala só entra num escopo por concessão do administrador, que a concessão aparece na sala e o que a sala de análise devolve (doc 06 seção 7; termo, item 3). A RN09 do doc 01 muda com a aceitação desta ADR.
- **Limite declarado:** a análise só alcança salas **ativas**, com IA e com concessão vigente; sala concluída já foi expurgada do nó. "Presos em operações diferentes" ao longo dos anos **não** sai deste desenho.

### 3. Posição dos agentes
- **Missão ativa:** os agentes compartilham a posição pelo MLS na sala da missão; o comandante vê no próprio app (rádio ou internet, via nó de campo). O servidor e o painel não veem posição.
- **Limite declarado (01g-G19):** mostrar o mapa só ao comandante e ao substituto é restrição de **interface**. A posição vai por MLS à sala da missão, e todo aparelho membro a decifra e guarda: um aparelho de agente capturado desbloqueado tem a posição da equipe. Tornar a restrição criptográfica (posição enviada a um grupo MLS de comando: comandante, substituto e quem a missão definir) é decisão da Fase 8, junto com a D-R20 e a pergunta 9 do `01c`.
- Em sala de missão **sob custódia**, a mensagem de posição é conteúdo como outro qualquer e leva cópia de custódia (ADR-0003 seção 3): fica guardada pelo prazo da custódia e só se abre por quórum presencial. Isso entra na transparência e no inventário do doc 06.
- **Depois da missão:** a posição guardada continua sob quórum (regra 14). O painel não mostra posição nem região.

### 4. Relatórios de operação
Gerados pelo nó de IA. O relatório da operação vai **só ao administrador** dela (RN09). Na sala de análise, as respostas vão aos membros dela (analistas credenciados), o que muda a RN09 se esta ADR for aceita (01e-E21b). Exportação ao sistema corporativo assinada e registrada.

### 5. Controles da IA local contra vazamento interno
| Risco | Controle |
|---|---|
| Vazamento entre compartimentos | Filtro antes da busca; índice por sala; só salas com concessão vigente; teste automático com pergunta de um escopo tentando obter dado de outro (A99) |
| Injeção de prompt por mensagem de sala, inclusive **dentro** do escopo | Autorização fora do modelo; a IA não tem ferramentas que agem, só responde; saída estruturada validada no código; verificação de sobreposição com a fonte; teste de injeção dentro do escopo (A99) |
| Nó de IA comprometido | Isolado, sem saída de rede (regra 7; lista de destinos internos proposta no doc 02 seção 9.1, D31 pendente), escopo mínimo, acesso físico por duas pessoas |
| Rastros em disco | Sem log de conteúdo, sem swap, despejo de memória desligado, sistema de arquivos só leitura |
| Embeddings reversíveis | Vetor tratado como conteúdo: cifrado em repouso e expurgado com a sala de origem (cascata) |
| Treino com dados | **Proibido treinar ou ajustar com dados das salas**; só busca (RAG). Se um dia houver, os pesos viram material classificado e exigem ADR |
| Modelo adulterado | Origem conhecida, hash fixado (regra 8), sem execução de código |
| GPU compartilhada | GPU dedicada ao nó de IA |
| Expurgo incompleto | Expurgo da origem apaga índice, vetores, indícios e relatórios dela, com recibo (RN04); retirada na sala de análise (A98) |

## Alternativas rejeitadas
- **IA no backend com acesso ao texto para o painel:** contraria a regra 3 e a ADR-02.
- **Painel com indicadores nominais:** o servidor passaria a ver nomes e vínculos de todas as operações.
- **Nível 1b no painel** (tipo, região, riscos por categoria): metadado novo no servidor hostil e canal em claro do nó de IA para o servidor; a regra 3 não muda (GD-15).
- **Nível 2 sem ato da sala de origem** (versão anterior): a chefia e o nó de IA viravam leitores institucionais das salas (01e-E01).
- **Análise sem escopo (todas as operações):** quebra a necessidade de conhecer; vira o mapa completo num ponto só.
- **Quórum presencial para criar escopo:** banaliza o rito da regra 14, que é para abrir dado selado; ele continua sendo o caminho para a sala sem concessão.
- **Posição em tempo real no painel web:** contraria a regra 14 e expõe a missão ao servidor.
- **"Exceção à regra 3" declarada na ADR:** a regra não admite exceção.

## Consequências
- O painel fica só com o nível 1a. Agregados derivados e indicadores nominais só existem no app, por MLS, e só de salas que concederam.
- O administrador de cada sala passa a ter um ato a mais (conceder e retirar); quem não concede fica fora da análise.
- O nó de IA ganha papel de analista entre salas, só dentro do escopo concedido; aumenta o seu valor como alvo (R103).
- Riscos: R103 a R107, com R104 em Médio até o teste de injeção dentro do escopo (doc 04 seção 5.4). Requisitos: RF85 a RF88 e RF92. Testes: A96 a A99 (doc 07 seção 11). Decisão: D-R41.

## Revisar quando
O Fable revisar o redesenho do nível 2 (01e-E01); o spike de inferência medir memória e tempo (etapa 6); D-R41 for decidida; ou o jurídico definir base e prazo para dados de terceiros (LGPD art. 4º, III; D-R34) [conferir].
