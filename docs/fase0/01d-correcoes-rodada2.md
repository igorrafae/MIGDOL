# Correções da revisão — Fase 0, etapa 1b, segunda rodada

**Modelo:** Claude Opus 5.5 (pipeline da Fase 0, etapa 1b, rodada 2). **Revisão seguinte:** Fable 5.2.
**Entradas:** `01c-revisao-correcoes.md` (achados C01–C27, recomendações e perguntas), `GD-respostas.md` (GD-12 a GD-14 prevalecem), `01b-correcoes.md`, ADRs 0003 a 0009 e README, docs 01 a 10.
**Limites:** nenhum código, commit ou push; nada enviado para fora. Fato jurídico não conferido segue **[conferir]**. Nada de cripto à mão: o texto só descreve recursos da biblioteca MLS e da pilha HPKE; o que depende de verificação virou critério do spike. Nenhum risco é aceito aqui: "aceito" só onde há aceite formal (doc 04 seção 3.2), e não há nenhum.
**Convenção nova:** nos documentos normativos, achados do 01c são citados como `01c-Cnn`, para não colidir com os cenários C01–C08 do doc 03.

## Resumo

| Severidade | Total | Fechado | Adiado | Pendente de resposta |
|---|---|---|---|---|
| Alta | 2 | 2 | 0 | 0 |
| Média | 18 | 10 | 2 (C06, C11) | 6 (C03, C04, C10, C15, C16, C17) |
| Baixa | 7 | 4 | 1 (C21) | 2 (C24, C25) |
| **Total** | **27** | **16** | **3** | **8** |

- **Altos abertos: nenhum.** C01 fechado no texto (GD-13). C02 fechado no texto; a detecção precisa ser provada no spike MLS (etapa 3), que tem critério novo.
- C04, C10 e C15 tiveram a parte de texto fechada; ficam como "pendente de resposta" pela parte que é decisão do responsável.
- "Fechado" quer dizer **fechado no texto**. C02, C08 e C12 dependem de prova no spike ou na Fase 2.
- **Ponto de parada (regra do pedido):** a ADR-0008 declarava "exceção à regra 3". A regra não admite exceção e mudá-la está fora do autorizado. Não mudei a regra 3: dividi o nível 1 em 1a (sem exceção) e 1b (só se o responsável mudar a regra). Seção 4.

## 1. Achados C01–C27

| ID | Sev. | Status | Onde | Nota |
|---|---|---|---|---|
| C01 | Alta | **Fechado** | ADR-0006 itens 8–9; ADR-0003 §4; regra 14; doc 07 §7.1, §7.4, A59, A74, RF69, §15; doc 02 §4.3, §6.3, §6.4, §7.1; doc 05 P2; doc 01 RF05, RF91, RN20; ADR-0004 itens 6–7, 10; doc 10 §5; doc 06 §3, §7 item 11, §8.2; termo item 10; doc 04 R63, R110 | GD-13: revogar e destruir são atos separados; **toda** revogação sela o segredo para a chave de quórum, com prazo da apuração; destruir exige duas autorizações. Prazo padrão: D-R35. Caminho de volta marcado **[proposta]** |
| C02 | Alta | **Fechado** (prova no spike) | ADR-0005 item 2 (reescrito); doc 02 §5.3 (`epoch_hash`; `epoch_auth` removido), §5.5, §11; doc 07 §5; doc 04 R98; doc 03 R11, Fase 0, Fase 1, §4.4; ADR-0004 item 3; doc 10 R90, F-R12; `PIPELINE.md` | Hash do transcript confirmado nos dados autenticados em claro; sinais: mesma época com hash diferente, mesma época que não decifra, commit divergente. Valor falso do servidor gera alarme falso. Partição total declarada como limite. Critério do spike com os dois casos |
| C03 | Média | Pendente de resposta | doc 07 §7.4 (sala de resposta, registrado) | Estado mínimo fora do lease e credencial de rastreio dependem da pergunta 8 |
| C04 | Média | Pendente de resposta (em parte fechado) | ADR-0003 §1; doc 01 RN17; doc 07 §7.3, §7.4; doc 02 §7.5 (`room_kind`) | GD-14 fecha a parte da custódia: sala de sistema sempre comum, só pelo fluxo do sistema. Membros, ciclo, batimento e missão: pergunta 8 |
| C05 | Média | **Fechado** | ADR-0003 §2; doc 02 §5.2; doc 04 §6.4 | Estrutura assinada (classe, lote, validade, identificador, suíte); `info` do HPKE com sala, período e classe; atribuição chave ↔ sala assinada pelo administrador; lote com M de N; troca da chave de cerimônia |
| C06 | Média | Adiado (HSM) | ADR-0003 §2 (nota); doc 07 §7.3; doc 02 §5.2; doc 04 §6.4 | Período só por tempo é a proposta do 01c; decide-se com a capacidade de chaves do HSM (pergunta 11) |
| C07 | Média | **Fechado** | ADR-0003 §4, §5 item 3; doc 04 §6.4 | Chave por classe **e** período; eliminação por período; pseudonimato como exceção declarada; decapsulação dentro do HSM como requisito de compra |
| C08 | Média | **Fechado** (verificar no spike) | ADR-0003 §3; doc 03 Fase 0 | Credencial do remetente com cadeia até a AC dentro da cópia; assinatura com rótulo da biblioteca |
| C09 | Média | **Fechado** | ADR-0003 §3; doc 02 §5.3; doc 01 RF16, RF17, RF82; doc 07 §7.3; doc 06 §7 item 9; termo item 2 | Todo envelope de aplicação leva cópia, sem exceção por tipo; apagar mensagem e autodestruição não tiram a cópia |
| C10 | Média | Pendente de resposta (em parte fechado) | ADR-0003 §7; doc 01 §3.1, RN18, RF83; doc 02 §7.5 (`expunged_at`); doc 07 §7.3; doc 06 §8.2; termo item 2 | Prazo conta do expurgo quando não há conclusão; alerta de sala ativa além do teto (proposta: 12 meses). Kill-switch: pergunta 5 |
| C11 | Média | Adiado (HSM) | ADR-0003 §2; doc 03 Fase 0 | Suíte escolhida com o HSM; spike usa a suíte candidata |
| C12 | Média | **Fechado** (prova na Fase 2) | ADR-0006 itens 3 e 10; doc 02 §4.3; doc 04 §8; doc 03 Fase 2, V05; A73 | Geração no embrulho; `n` e `n+1` até a confirmação; gravação atômica; RPO 0 **[proposta]** |
| C13 | Média | **Fechado** (medição na Fase 2) | ADR-0006 "Limites declarados"; doc 04 R99; doc 03 Fase 2 | Chave do banco fixa, prazo offline e backup do diretório declarados; rekey medido na Fase 2 |
| C14 | Média | **Fechado** | ADR-0006 item 8; ADR-0003 §4; regra 14; doc 01 RN20; doc 02 §6.3, §6.4; doc 07 §7.1, §15; doc 06 §3, §8.2 | Classe nova com prazo, rito, volta, painel e ordem na reconexão |
| C15 | Média | Pendente de resposta (em parte fechado) | doc 02 §4.2, §7.5; doc 04 R40 | Chave do índice em HSM (opção mais restritiva) e limite do D09 registrados. Custo para o gestor: pergunta 12 |
| C16 | Média | Pendente de resposta | ADR-0005 seção "Pendente"; doc 03 Fase 0 item 8 | O spike só levanta o que a biblioteca oferece; a decisão espera o responsável |
| C17 | Média | Pendente de resposta | ADR-0004 item 8 (nota) | Pergunta 10 |
| C18 | Média | **Fechado** | doc 07 R66, R68, R97, §7.3; ADR-0003 §3 e Consequências; doc 04 §5.1, §5.3; doc 01 RNF-A4; doc 03 §1.3; doc 02 §11 | "Aceito" trocado por "proposto para aceite", com quem aceita e validade. Os aceites formais seguem pendentes (pergunta 7) |
| C19 | Média | **Fechado** | `CLAUDE.md` regra 2; doc 02 §4.1, §9.1; doc 03 Fase 4 | D30 fechado pela GD-12; D28 virou critério de entrada da Fase 4, com dono; D31: lista corrigida com a inferência local, decisão da regra 7 pendente (antes da Fase 5) |
| C20 | Média | **Fechado** | doc 07 §7.4, A70 | Intervalos de pertença no pedido assinado; reinclusão depois da perícia pelo mesmo rito da revinculação **[proposta]** |
| C21 | Baixa | Adiado | — | Renomear D01–D34 para `RD-Dnn` mexe em cerca de 180 citações em 15 arquivos e colide com os cenários D01–D06 do doc 03. Fica para um passe mecânico separado, revisado, antes do commit. Nesta rodada, citações novas do 01c já usam `01c-Cnn` |
| C22 | Baixa | **Fechado** | doc 02 §4.1, §6.2 (diagrama), §6.3 | Chave do cache aleatória e embrulhada; "lease vence: ilegível, não apagado"; cache sem chave por sala |
| C23 | Baixa | **Fechado** nos documentos | ADR-0003 Contexto 7, §5; doc 01 RN20, §9; doc 07 §7.3, A54, D-R15; doc 04 D-R15; doc 05 §3.4; doc 06 §3, §6.3, §14; doc 09 §3; doc 02 §6.4; doc 03 §4.11 | CI obrigatória nos dois casos; base é ordem judicial ou processo administrativo com validação jurídica. A regra 14 **não** foi mudada (fora do autorizado): texto proposto na seção 3 |
| C24 | Baixa | Pendente de resposta | — | Depende de D-R30 (pergunta 2) |
| C25 | Baixa | Pendente de resposta | doc 01 §4 (nota); ADR-0009 §1 (nota) | Pergunta 9 |
| C26 | Baixa | **Fechado** | `PIPELINE.md` (15 altos; ADR-0002 como entrada da etapa 5); `docs/adr/README.md`; doc 03 Fase 1 | Estado autenticado do grupo e ADR-0005 desde a Fase 1, para a Fase 4 não exigir migração |
| C27 | Baixa | **Fechado** | ADR-0003 §1, §6; doc 02 §10.1; doc 01 RF82 | Histórico assinado das versões do `instituicao.yaml`; mídia de entrada só de leitura com hash conferido |

### Adiamentos da primeira rodada (condições do 01c, seção 3)

| ID | Situação agora | Onde |
|---|---|---|
| D08 | Adiado **com portão escrito:** nenhuma sala sob custódia em produção antes de KEM híbrido ou de aceite formal do R78 para a custódia (pergunta 13) | doc 03 §4.12; ADR-0003 Consequências; doc 04 §5.2 |
| D28 | Adiado **com critério de entrada da Fase 4** e dono (responsável pela homologação) | doc 03 Fase 4; doc 04 §10 |
| D30 | **Fechado** pela GD-12 | `CLAUDE.md` regra 2; doc 02 §4.1 |
| D31 | Adiado: lista corrigida (01c-C19); regra 7 não mudou | doc 02 §9.1 |
| D26 | Adiado, sem mudança (protótipo de D-R1) | ADR-0006 Consequências |

## 2. Decisões aplicadas

| Decisão | O que mudou | Onde |
|---|---|---|
| **GD-12** (regra 2) | Texto aprovado aplicado, com os ajustes de C19 e Swift | `CLAUDE.md` regra 2; doc 02 §1 (princípio 5), §4.1, §8.1; ADR-0007 item 5; `app/CLAUDE.md` (nota) |
| **GD-13** (revogação) | Revogar ≠ destruir; selagem em toda revogação; destruição com duas autorizações; classe na regra 14 | Ver C01 e C14; R110 novo; RF91 novo; A74 novo; D-R35 nova |
| **GD-14** (D-R29, opção B) | Custódia é o padrão com a flag ligada; sala comum com motivo e segunda assinatura; salas de sistema sempre comuns. A leitura "toda sala nova sob custódia, sem exceção" saiu de todos os lugares | ADR-0003 Contexto 8, §1, Alternativas, Revisar; doc 01 RN17, RF82, §11 item 11; doc 07 §7.3, §7.4, RF70, D-R29, A75; doc 04 D-R29, R111; doc 02 §4.1, §5.2, §7.5, §10.1, §11; doc 03 Fase 4, §4.11; doc 06 §7 item 9; doc 09 §3. D-R36 nova (quem dá a segunda assinatura) |

**Recomendações do 01c sem decisão do responsável (aplicadas):**

| Recomendação | Onde |
|---|---|
| R66 volta a **Alto até D-R14**, com reavaliação proposta para Médio depois de A72–A74 e do comitê. Voltei também P para Médio, porque a primeira rodada baixou P sem o comitê | doc 07 R66; doc 04 §5.1, §5.3 |
| Terceira assinatura da conclusão passa ao **superior imediato do solicitante**, definido na criação, quando ele é comandante ou subcomandante. Signatário não membro assina sobre identificador opaco, número e hash, sem ver conteúdo | doc 01 §4, RN18, RF83, §11 item 12; doc 02 §4.1, §6.2; doc 04 §6.2; doc 07 A69; ADR-0009 §1 |
| R97 e R68 sem rótulo "aceito" (C18) | doc 07 R68, R97, §7.3; ADR-0003 §3; doc 04 §5.1 |

## 3. Mudanças no `CLAUDE.md`

### Aplicada: regra 2 (GD-12)

Antes:
> 2. Chaves privadas e chaves de sessão existem apenas em core/ (Rust). Nunca atravessam a ponte para Kotlin ou TypeScript.

Depois:
> 2. Chaves privadas e de sessão de membro (MLS, assinatura, cache) existem apenas em core/ (Rust) e nunca atravessam a ponte para Kotlin, Swift ou TypeScript. Chaves de hardware (StrongBox/TEE, Secure Enclave) não saem do hardware: Kotlin e Swift só seguram a referência, e o uso é chamado a partir do core; o que o hardware devolve (ex.: desembrulho pelo Keystore) transita só no quadro da chamada iniciada pelo core, é zerado em seguida e sozinho não abre nada. Chaves da instituição (AC, quórum, cerimônia) ficam em HSM, fora do app.

### Aplicada: regra 14 (GD-13)

Antes (trecho):
> Todo acesso institucional a dado cifrado (custódia, evidência, registros, termo, localização, pseudonimato) exige quórum presencial [...]

Depois (trecho; o resto da regra não mudou):
> Todo acesso institucional a dado cifrado (custódia, evidência, registros, termo, localização, pseudonimato, segredo de lease guardado) exige quórum presencial [...]

### Não aplicadas (fora do autorizado)

| Regra | Situação | Texto proposto |
|---|---|---|
| 3 | **Conflito registrado** com a ADR-0008 (nível 1b). Não mudei | Nenhum. Opções para o responsável na seção 5 |
| 4 | Pendente da aceitação da ADR-0007 | "A interface (Compose, SwiftUI e, enquanto existir, app/src) não grava dado fora do core (sem localStorage, sessionStorage, IndexedDB, SharedPreferences, DataStore, arquivos próprios, UserDefaults nem Core Data) e não abre rede própria (sem fetch, OkHttp nem URLSession). Todo acesso a dados e rede passa pelo core." |
| 7 | Pendente (D31); lista corrigida no doc 02 §9.1 | "ai-node/ não grava arquivo em claro e só abre conexão para destinos fixos: gateway (mTLS), serviço de anexos da instância, PostgreSQL da IA e serviço de inferência local (localhost). Qualquer outro destino é bloqueado e testado em tests/security (E01)." |
| 14 (base) | C23 aplicado nos documentos; a regra continua "ordem judicial, ou CI em processo administrativo". Os documentos ficaram **mais** restritivos (CI sempre), sem contradizer a regra | Trocar "e base registrada (ordem judicial, ou CI em processo administrativo com validação jurídica)" por "com CI e base registrada (ordem judicial, ou processo administrativo com validação jurídica)" |

## 4. Integração das ADRs 0007–0009

Mantidas as decisões do responsável: Android na v1, iPhone depois do estudo, sem macOS; painel com indicadores agregados e nível 2 só por MLS; RBAC + ABAC.

### Inconsistências corrigidas nas ADRs

| ADR | Inconsistência | Correção |
|---|---|---|
| 0007 | "Sem WebView, a regra 4 vale por construção": falso. Kotlin e Swift têm armazenamento e rede próprios | Item 3 corrigido; texto novo da regra 4 proposto (item 5), aplicado só quando a ADR for aceita |
| 0007 | Status dizia que a regra 2 cita só Kotlin e TypeScript | Regra 2 já cita Swift (GD-12); status e item 5 atualizados |
| 0007 | `app/` renomeado "antes da Fase 1" sem condição | "Se a ADR for aceita"; até lá, nada se renomeia (`app/CLAUDE.md`, nota) |
| 0007 | Contraria também doc 01 D1 (Capacitor), não citado | Status e doc 01 D1 com nota |
| 0008 | "Exceção declarada à regra 3": a regra não admite exceção | **Ponto de parada.** Nível 1 dividido: **1a** (metadados que o servidor já tem; sem exceção) e **1b** (tipo e região, que seriam metadados novos; riscos por categoria, derivados do conteúdo). O 1b não existe até o responsável decidir |
| 0008 | "Fase do nó de IA, depois da v1" × RF85–RF87 "Fase do nó de IA" (Fases 5 e 6 são v1) | "Depois da v1 (após a Fase 6)" na ADR, em RF85–RF87 e no doc 03 "Depois da v1" |
| 0008 | Sala de análise contraria RN09 (relatório só para o administrador) sem dizer | Status e §2; nota em RN09; transparência no doc 06 §7 item 3 e termo item 3 (condicional) |
| 0008 | Não dizia que o nó só cruza salas de que é membro, nem a custódia da sala de análise e da posição em missão | §2 e §3: só salas do escopo de que o nó é membro; sala de análise segue a flag (GD-14); posição em sala sob custódia leva cópia; inventário do doc 06 |
| 0008 | "Compartimento (doc 08)" sem seção correspondente | Doc 08 §5.1 novo (compartimento de análise; credencial em todo o escopo; grau herdado; nada classificado até G3) e doc 04 §6.1 |
| 0009 | "Corregedoria e auditoria: ver o registro de acessos" contraria a regra 14 (registros só por quórum) | Corregido: consulta a trilha de metadados e **pede** abertura pelo rito |
| 0009 | Faltavam papéis do RN18 (equipe de inteligência, solicitante, superior imediato) | Linhas acrescentadas; coluna de correspondência com os atores do doc 01 |
| 0009 | "Membro da sala (estado MLS)" como atributo do servidor, que não vê o estado MLS | Servidor usa a pertença verificada pela declaração assinada (ADR-0005 item 4); clientes conferem ações do grupo contra o estado autenticado |
| 0009 | Ponto único de autorização no servidor podia ser lido como suficiente | Ações dentro do grupo são conferidas também pelos clientes; servidor hostil não concede papel |

### Coerência com o resto

| Documento | O que mudou |
|---|---|
| doc 02 §1, §2, §3, §7.2, §8 | Notas da ADR-0007 (proposta): telas nativas, UniFFI para Swift, `app/` → `android/` se aceita. A pilha atual continua valendo até a aceitação |
| doc 02 §9.2, §11 | Linha da análise entre operações; resumos das ADRs 0005, 0006, 0008 e ADR-12 atualizados |
| doc 02 §7.5 | `users.role` vira atribuições com prazo (ADR-0009); `room_kind` e `expunged_at` |
| doc 04 §6.2 | Analista de inteligência entra nas combinações proibidas (com administrador de infraestrutura e detentor de cartão); terceira assinatura; segunda assinatura da sala comum; a tabela vira as combinações proibidas do `instituicao.yaml` (D-R33) |
| doc 08 §1, §5.1 | Decisões até D-R36 e ADRs 0007–0009; compartimento de análise |
| `admin-panel/CLAUDE.md` | Nota: só nível 1a; 1b depende da regra 3; nível 2 nunca no painel; ADR-0009 |
| `app/CLAUDE.md` | Nota: ADR-0007 propõe Compose e `android/`; não renomear; regra 2 já cita Kotlin e Swift |

## 5. Pendências

### Para o Fable

1. **ADR-0005 item 2 (C02):** os três sinais, o argumento do alarme falso e o custo de expor o hash do transcript em claro. Confirmar que comparar antes de decifrar não abre canal de negação de serviço novo além do que o servidor já tem.
2. **ADR-0006 item 8 (GD-13):** selagem em toda revogação; destruição antecipada é só exclusão de registro (limite declarado); "volta" ao aparelho **[proposta]**; segredos antigos em backup do diretório; selagem pelo nó de campo. Conferir se a autodestruição do nó de campo capturado (que apaga primeiro os segredos de missão, ADR-0004 item 7) continua aceitável sob a GD-13: é destruição sem duas autorizações, justificada por adversário ativo, mas deixa ilegível o cache de missão de aparelhos não apreendidos.
3. **ADR-0006 itens 3 e 10 (C12):** protocolo de geração e RPO 0 proposto.
4. **ADR-0003:** estrutura assinada e atribuição chave ↔ sala (C05); chave por classe e período e requisito de decapsulação no HSM (C07); credencial dentro da cópia (C08); teto de sala ativa (C10); implementação da GD-14 (duas assinaturas, sala de sistema, R111).
5. **ADR-0008:** divisão 1a/1b. E uma dúvida de fundo: o nível 2 dá a quem não é membro um conhecimento **derivado** do conteúdo das salas, sem quórum. Não contraria a letra da regra 14 (a lista de classes não inclui análise), mas toca "membros leem as salas pelo MLS". Pede decisão explícita e transparência, não só a ADR.
6. **ADR-0009:** correções acima, em especial a linha da corregedoria.
7. **R66:** voltei P e residual; conferir se só o residual devia voltar.
8. **C21** adiado; e a convenção `01c-Cnn` usada nesta rodada.

### Perguntas que restam ao responsável

Do 01c, seção 7 (as respondidas saíram: 1 pela GD-14; 4, parte da regra 2, pela GD-12; 6 pela GD-13; 3, parte do terceiro signatário, pela recomendação):

| # | Pergunta | Por que importa |
|---|---|---|
| 2 | D-R30: frase de escopo na regra 12; foto fora da exportação; prazo do termo | C24; RF84 não se implementa antes |
| 3 (resto) | Quem autoriza trocar o titular de um papel da sala durante a operação | Senão o administrador escolhe os signatários |
| 4 (resto) | Regra 7 (D31), com a lista corrigida | Antes da Fase 5 |
| 5 | Kill-switch de um administrador ao lado da conclusão com três assinaturas; teto de sala ativa (proposta: 12 meses) | C10 |
| 6 (resto) | Prazo padrão do segredo selado (D-R35) | GD-13 diz "prazo da apuração", sem valor |
| 7 | Quem é o responsável pela homologação e quem compõe o comitê; aceite formal de R68, R97, R66 | Nenhum residual tem aceite hoje |
| 8 | Sala de resposta: membros, tempo de reação, aparelho de função, missão | C03, C04 |
| 9 | Administrador de toda operação é o comandante, ou só na missão de campo; aprovadores distintos do gestor em unidade pequena | C25; ADR-0009 |
| 10 | Recusa de sessão direta em missão com cobertura parcial; nó destruído ou capturado | C17 |
| 11 | HSM: decapsulação interna, curvas, ML-KEM, destruição em backup, capacidade de chaves, quórum para lote | C06, C07, C11 |
| 12 | Chave do índice de matrícula em HSM, com o custo de consulta para o gestor | C15 |
| 13 | Custódia em produção antes do KEM híbrido, com aceite formal do R78, ou espera | Portão D08 |

Novas desta rodada:

| # | Pergunta | Opções |
|---|---|---|
| N1 | **Regra 3 × ADR-0008 nível 1b** | (a) manter a regra 3 e não ter o 1b (situação atual; mais restritiva); (b) mudar a regra 3 para admitir agregados derivados publicados pelo nó de IA, com limite mínimo e teste; e, à parte, aceitar ou não tipo e região como metadados novos no servidor |
| N2 | Regra 4: aceitar a ADR-0007 e o texto proposto | Muda a regra junto com a aceitação |
| N3 | Regra 14: aprovar a redação de C23 (CI sempre; base ordem judicial ou processo administrativo) | Hoje os documentos são mais restritivos que a regra |
| N4 | D-R36: quem dá a segunda assinatura da sala comum com a flag ligada; motivo em lista fechada ou livre | GD-14 |
| N5 | ADR-0008: aceitar que a posição em missão, em sala sob custódia, fique na cópia de custódia pelo prazo dela | Alternativa: posição como mensagem de sala de sistema da missão (comum), o que exige desenho novo |
| N6 | RPO 0 para segredos de lease (proposta) e teto de 12 meses de sala ativa (proposta) | Valores marcados como proposta |

Continuam abertas, da etapa 1: 4 (D-R24), 5 (D-R14), 6 (D-R5), 8 (D28, agora entrada da Fase 4) e 15 (IA em sala secreta ou de compartimento).

**Jurídico [conferir]** (sem mudança de fundo): processo administrativo com validação jurídica como base; corregedoria sem ordem judicial; prazo de 5 anos e início pelo expurgo; autorização arquivística; base para guardar o segredo de lease (dado indireto do agente); base para posição em missão; dados de terceiros no nível 2 (D-R34).

## 6. Arquivos alterados nesta rodada

| Arquivo | O que mudou |
|---|---|
| `CLAUDE.md` | Regras 2 e 14 |
| `app/CLAUDE.md`, `admin-panel/CLAUDE.md` | Notas das ADRs 0007, 0008, 0009 |
| `docs/adr/0003` | GD-13, GD-14, C05, C07, C08, C09, C10, C11, C18, C23, C27, D08 |
| `docs/adr/0004` | Itens 3, 6, 7, 8 (nota C17), 10 |
| `docs/adr/0005` | Item 2 reescrito (C02); C16 pendente; consequências e critério do spike |
| `docs/adr/0006` | Itens 3, 8, 9, 10; limites declarados (C01, C12, C13, C14) |
| `docs/adr/0007`, `0008`, `0009` | Inconsistências da seção 4 |
| `docs/adr/README.md` | Estado das ADRs; ADR-0002 como entrada da etapa 5 |
| docs 01 a 10 | Conforme a tabela da seção 1 |
| `docs/anexos/termo-de-ciencia-e-autorizacao.md` | Itens 2, 3 (condicional) e 10 |
| `docs/fase0/PIPELINE.md` | Critério da etapa 3, "15 altos", entrada da etapa 5, estado |
| `docs/fase0/01d-correcoes-rodada2.md` | **Novo** (este arquivo) |

**IDs novos** (conferidos sem colisão): R110, R111, RF91, D-R35, D-R36, A74, A75, F-R12. Seção nova: doc 04 §5.3; doc 08 §5.1.
