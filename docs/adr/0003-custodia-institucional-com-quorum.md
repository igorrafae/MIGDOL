# ADR-0003 — Custódia institucional em sala sob custódia, com quórum presencial de comandantes

**Status:** proposta, **revisada na etapa 1b da Fase 0** (achados D01–D12, D27, D28 de `docs/fase0/01-revisao-desenho.md`; decisões GD-1 a GD-6) e na **segunda rodada** (`docs/fase0/01d-correcoes-rodada2.md`: GD-13, GD-14; 01c-C05, C07, C08, C09, C10, C18, C23, C27) e na **terceira** (`docs/fase0/01f-correcoes-rodada3.md`: 01e-E06, E09b, E10a, E17, E21e; ADR-0010) · **Contraria** o objetivo "sem histórico permanente" (doc 01) e a retenção só por operação (RN03) **somente em sala sob custódia** · **Complementa** ADR-04 (expurgo por destruição de chave), ADR-10 (anexos cifrados no cliente), ADR-15 (pacote de instalação e `instituicao.yaml`), ADR-0005 (estado autenticado do grupo), ADR-0006 (segredo de lease guardado) e doc 04 seção 6.4 · **Depende de** parecer do jurídico (D-R5, D-R15) e do levantamento do HSM (pergunta 14 da revisão; pergunta 11 do 01c)

## Contexto
O responsável pelo produto definiu (GD):
1. Todo dado é cifrado. A instituição só decifra com **quórum mínimo de 3 comandantes**, cada um com uma parte, em HSM.
2. Custódia **só em sala marcada como sob custódia**. O risco admitido é este: comprometer o HSM ou o quórum expõe os períodos guardados dessas salas (GD-1). **Sala comum é MLS puro.**
3. A flag de custódia é **por instituição**, no `instituicao.yaml` assinado (ADR-15). Exército: **ativa**, motivo investigação militar. Polícia: **desligada por padrão**; a instituição pode ligar, com motivo corregedoria ou inquérito, quando o edital exigir. Mudar exige cerimônia registrada e **não retroage** (GD-2).
4. Ciclo da sala: a equipe de inteligência marca **ativa**; a **conclusão** exige relatório no sistema corporativo e três assinaturas no app (GD-3; doc 01 RN18).
5. Sala sob custódia guarda a cópia **até 5 anos a contar da conclusão**; depois, eliminação com recibo (GD-1).
6. **Sem HSM online e sem aprovação remota por quórum na v1.** A revinculação não usa custódia: o histórico vem de outro membro pelo MLS (GD-4; doc 07 seção 7.4).
7. Abrir custódia exige **quórum presencial** e uma **base**: ordem judicial, ou processo administrativo com validação jurídica **[conferir]**. A CI registra e tramita **sempre**; **não substitui** ordem judicial, que é cumprida (GD-6).
8. Com a flag ligada, a custódia é o **padrão**; sala comum só com motivo registrado e segunda assinatura; salas de sistema sempre comuns (GD-14, D-R29 opção B).
9. Revogar aparelho não destrói o segredo de lease: ele fica guardado para a chave de quórum (GD-13; ADR-0006 item 8).

A versão anterior desta ADR (chave simétrica de custódia conhecida pelos membros, HSM online de recuperação, aprovação remota por três celulares) foi rejeitada pela revisão de desenho (D01–D05). Está em "Alternativas rejeitadas".

## Decisão (proposta)
**Leitura normal não muda:** membros leem pelo MLS no próprio aparelho. **Membros não decifram cópia de custódia.**

### 1. Sala comum, sala sob custódia e sala de sistema
- O `instituicao.yaml` assinado traz `custodia.ativa` e `custodia.motivo`. O app mostra a configuração vigente aos membros.
- **Flag desligada:** toda sala nasce **comum**.
- **Flag ligada (GD-14):** a sala nasce **sob custódia** por padrão. Nasce **comum** só se o criador registrar o **motivo** e houver **segunda assinatura** de pessoa distinta do criador, com papel definido no `instituicao.yaml` (quem assina: D-R36). **Motivo (01e-E09b) [proposta; D-R36]:** uma **categoria** de lista fechada no `instituicao.yaml`, assumida como metadado e registrada na trilha; texto livre opcional só na marca, cifrado no estado do grupo, **sem hash na trilha** (com lista fechada o hash seria o próprio motivo, e texto curto sai por dicionário). O cliente recusa marca "comum" com a flag ligada sem as duas assinaturas válidas. **A segunda assinatura é conferida contra a lista assinada de papéis** (cerimônia; ADR-0009 seção 3), e não contra o diretório nem contra o estado montado pelo criador, que poderia inventar o segundo signatário (01e-E10a).
- **Sala de sistema** (hoje só a sala de resposta, doc 07 seção 7.4): sempre **comum**, também com a flag ligada (GD-14). Nasce só pelo fluxo próprio do sistema, não pelo criar sala do usuário. O resto do desenho (membros, ciclo, batimento) está pendente (01c-C04).
- A **marca de custódia** vai no **estado autenticado do grupo MLS** (extensão do contexto do grupo), assinada na criação. Contém: tipo (comum, sob custódia ou sistema), motivo e segunda assinatura quando comum com a flag ligada, hash do `instituicao.yaml` vigente, papéis da sala (doc 01 RN18) e a primeira chave pública de período.
- Os clientes conferem a marca contra a configuração assinada. Para conferir sala criada antes de uma mudança de flag, o pacote traz o **histórico assinado das versões** do `instituicao.yaml` (01c-C27). **Recusam** commit que altere a marca: ela é **irreversível** enquanto a sala existir.
- Mudar a flag é cerimônia registrada e **não retroage**: sala comum continua comum; sala sob custódia continua sob custódia até a eliminação.

### 2. Cópia de custódia por HPKE para chave pública (D01, D02)
- Em sala sob custódia, cada mensagem leva uma cópia cifrada **no remetente** (core) com **HPKE (RFC 9180), modo base, para a chave pública de custódia do período**. A chave efêmera é descartada.
- O parâmetro `info` do HPKE liga a cópia a **sala, período e classe** (recurso da biblioteca; nada à mão) (01c-C05).
- Anexo: a cópia da mensagem `attachment` contém a chave do arquivo; o objeto cifrado é retido pelo prazo de guarda.
- **A chave privada existe só no HSM offline.** Membros, nó de IA, servidor e nó de campo têm só a pública. Comprometer um aparelho-membro **não** abre a custódia (corrige D01).
- **Período (hipótese do spike, 01e-E21e):** **só por tempo** (proposta: mês), como propôs o 01c-C06. A chave pública do novo período entra no estado do grupo por commit do administrador. Como os membros não têm a chave privada, trocar o período a cada mudança de membros só mudaria a granularidade do escopo de abertura, ao custo de uma chave no HSM por mudança; essa alternativa (texto anterior: a cada commit que altera membros e, no máximo, a cada 30 dias) fica para decidir com o HSM (capacidade de chaves; pergunta 11 do 01c). O spike trabalha com o período por tempo.
- **Origem das chaves públicas (01c-C05):** o HSM offline gera, em cerimônia com o **mesmo M de N** da abertura, um **lote** de pares. Exporta só as públicas, cada uma numa **estrutura assinada** pela **chave de cerimônia de custódia** com: classe, identificador do lote, validade, identificador da chave e suíte. A pública da chave de cerimônia é fixada no app pelo pacote assinado. O cliente **recusa** chave sem essa assinatura, fora da validade ou de outra classe. O serviço `custody` distribui o lote.
- **Atribuição chave ↔ sala:** assinada pelo administrador no commit que instala a chave de período, e guardada pelo `custody`. Conferida em cada cerimônia contra o estado do grupo. Servidor que entregue a mesma chave a duas salas só alarga o escopo de abertura; não expõe nada, e a atribuição assinada denuncia o reuso.
- Lote esgotado: a sala fica no período corrente até a próxima cerimônia.
- **Troca da chave de cerimônia:** em cerimônia com M de N; a nova pública vai num pacote assinado novo; a antiga continua válida só para conferir lotes já emitidos, até a eliminação deles; revogação registrada e avisada aos clientes pelo pacote.
- **Suíte (01c-C11):** escolhida **junto com o HSM** (curva suportada, decapsulação interna, roteiro de ML-KEM) e fixada com hash na Fase 1 (regra 8). **Candidata até a escolha do HSM (GD-20):** DHKEM(P-256, HKDF-SHA256), HKDF-SHA256 e AES-256-GCM, a curva mais comum em HSM; o spike MLS usa essa suíte. KEM híbrido pós-quântico quando a pilha e o HSM tiverem suíte estável (D08, D-R12). Nada à mão (regra 1).

### 3. Cópia ligada à mensagem e verificável (D06)
- **Todo envelope de aplicação** em sala sob custódia leva cópia, **sem exceção por tipo**: texto, anexo, resultado de OCR, pedido e resposta da IA (inclusive a resposta individual cifrada só para quem pediu), autorização de exportação, `history_share` e `control` (01c-C09). O servidor não vê o tipo; qualquer exceção viraria desvio.
- **Apagar mensagem para todos** (RF16) e **autodestruição** (RF17) tiram a mensagem dos aparelhos, **não** a cópia de custódia (01c-C09). Os textos de transparência dizem isso (doc 06 seção 7; termo, item 2).
- O conteúdo MLS da mensagem leva o **hash do cifrado de custódia**. Quem recebe confere. Diferença é alerta de segurança (servidor trocou a cópia, ou remetente inconsistente).
- Dentro do cifrado de custódia vão o conteúdo, a **assinatura do remetente** (chave de assinatura MLS, no core) sobre conteúdo, sala, época e identificador da mensagem, e a **credencial do remetente com a cadeia até a AC** (01c-C08). A assinatura vai por dentro para não revelar o remetente ao servidor (doc 02 seção 5.3). Na abertura, a estação confere a assinatura contra a AC, não contra o diretório.
- A assinatura usa a função de assinatura da biblioteca com **rótulo próprio** (separação de domínio), nunca bytes crus **[verificar no spike]** (01c-C08).
- O serviço de entrega **recusa** envelope de sala sob custódia sem cópia de custódia.
- **Limite proposto para aceite (R97):** remetente malicioso pode cifrar lixo na própria cópia; só se nota na abertura. Equivale a vazar por outro meio; não abre conteúdo de terceiros. Auditoria amostral em cada cerimônia. Residual Baixo: aceite pelo **dono do ativo (`custody`)**, com revisão anual (doc 04 seção 3.2; 01e-E17), **pendente** (01c-C18).

### 4. O que exige quórum presencial (regra 14)
Toda abertura institucional de dado cifrado, sem exceção:
- cópia de custódia de sala sob custódia;
- imagens de evidência de câmera (doc 07 seção 6.3);
- registro de visualização (doc 05 seção 3.3) e registro de acessos (doc 08 seção 2.3);
- posição guardada em modo perda (doc 07 seção 7.4);
- termo de ciência e aceite da revinculação, com foto (doc 07 seção 7.4; GD-5);
- matrícula, na quebra de pseudonimato (doc 05 seção 3.4; D09);
- **segredo de lease guardado** de aparelho revogado (ADR-0006 item 8; GD-13), selado para a chave do **período de vencimento** (mês em que termina o prazo da apuração), para que a eliminação por período seja exata (01e-E06);
- **registro de investigação sigilosa** (ADR-0010), na classe "registros" da regra 14, com **série de chaves própria** por período: abrir registros de visualização não abre registros de investigação, e vice-versa.

**Chave por classe e período (01c-C07):** cada classe tem pares próprios, gerados em lote no HSM offline e com a pública assinada em cerimônia (seção 2), **por período** (proposta: mês). A eliminação é **por período**, destruindo a chave privada no HSM e nos backups dele (seção 7), e não apagando registro. Destruir uma classe ou período não depende de outra (D28, em parte).
- **Exceção declarada:** pseudonimato, cuja chave vive enquanto durar o vínculo funcional.
- **Requisito de compra do HSM:** decapsulação HPKE **dentro do HSM** para as classes sem chave por sala (evidência, registros, termo, localização, pseudonimato, segredo de lease), para que abrir um escopo não libere a chave do período inteiro.
- **Destruição antecipada** de um registro (ex.: segredo de lease selado, ADR-0006 item 8) é exclusão do registro, sem garantia nos backups; a destruição criptográfica ocorre na eliminação do período.

**Não há** quórum remoto, chave de quórum em HSM online nem chave de recuperação na v1. (O HSM online da CA intermediária, doc 04 seção 6.4, não guarda chave de quórum.)

### 5. Abertura
1. **Pedido registrado antes** da cerimônia: número da **CI**, que registra e tramita o pedido **em todos os casos** (01c-C23); **base**: número da ordem judicial, ou do processo administrativo com o parecer jurídico **[conferir]**; finalidade; escopo (sala, períodos, classe); solicitante. Para autoridade externa, ordem judicial obrigatória.
2. **Cerimônia presencial:** M de N cartões de comandantes (M ≥ 3; proposta 3 de 5), duas testemunhas, gravação. O HSM libera **só** as chaves do escopo pedido, nunca a chave de cerimônia.
3. **Onde decifra:** dentro do HSM, se o modelo fizer a decapsulação HPKE internamente (exigido para as classes sem chave por sala, seção 4). Para a cópia de custódia, se o HSM não fizer, na **estação isolada** (seção 6), com a chave do escopo carregada só durante a cerimônia e destruída ao fim **[conferir no modelo do HSM, pergunta 14]**.
4. **Saída:** pacote de evidência assinado (RF67), com hash, escopo e cadeia de custódia. Nada volta ao `server/` em claro.
5. **Registro:** pedido, CI, base, comandantes presentes, escopo e resultado na trilha de auditoria (cadeia de hash), com aviso ao Auditor. **Exceção (ADR-0010):** na investigação sigilosa, tudo isso vai no registro selado, e a trilha guarda só o hash, num registro de tamanho fixo igual ao dos demais registros selados; o aviso ao Auditor vai dentro do registro.

### 6. Estação isolada (D12)
É o segundo lugar com texto em claro fora dos aparelhos, além do nó de IA (doc 02 seção 1). Requisitos:
- boot verificado; sistema em mídia somente leitura;
- sem rede física nem sem fio; sem disco persistente (memória volátil);
- duas pessoas presentes; sala com acesso controlado e registro;
- **mídia de entrada** dos cifrados só de leitura, com hash conferido contra o `custody` antes da cerimônia (01c-C27);
- mídia de saída só para o pacote assinado;
- incluída no modelo de ameaças (doc 03 seção 1).

### 7. Prazo, eliminação e recibo (D07, D27)
- **Conclusão** da sala sob custódia (doc 01 RN18): a cópia de trabalho sai dos aparelhos e do nó de IA (expurgo, ADR-04). A **cópia de custódia** fica guardada **até 5 anos a contar da conclusão**.
- **Sala expurgada sem conclusão** (expurgo pelo administrador, doc 01 seção 3.1): o prazo conta **do expurgo** (01c-C10). A cópia de custódia não é tocada pelo expurgo. Se o expurgo por um administrador continua ao lado da conclusão com três assinaturas é pergunta aberta (pergunta 5 do 01c).
- **Sala ativa por tempo longo (01c-C10):** sala sob custódia ativa além de um teto (proposta: 12 meses) gera alerta ao chefe da inteligência; a continuação é registrada com motivo. Sem isso, a guarda não teria fim.
- **Eliminação** no fim do prazo: destruir no HSM as chaves privadas de todos os períodos da sala **e as cópias delas nos backups do HSM**, em cerimônia com registro e **recibo** assinado. Os cifrados no banco e nos backups do banco ficam ilegíveis; apagá-los é complemento.
- Se o modelo de HSM não apagar uma chave dentro de backup, as chaves de período são agrupadas por mês de vencimento sob uma chave de lote, e destrói-se a chave de lote **[conferir no modelo]**.
- **Inventário dos backups do HSM** (onde, quantos, quem guarda), conferido em toda cerimônia.
- Eliminar exige autorização arquivística (Lei 8.159; D-R5) **[conferir]**. Ordem de preservação (RF58) suspende a eliminação. A **ordem selada de preservação** da investigação sigilosa (ADR-0010 seção 8) suspende também a contagem do prazo. **Conferência sem abrir o registro (01g-G01c) [proposta]:** a ordem selada vem com uma **lista de bloqueio** assinada pelos participantes da sessão, só com os identificadores opacos das chaves de período preservadas (sem escopo, sem identificador da investigação, sem nível). Em toda cerimônia de eliminação, a estação confere as listas de bloqueio vigentes de forma automática, sem exibir nada além da contagem, sem abrir o registro selado e sem usar a chave da série de investigação; recusa eliminar chave bloqueada; grava a conferência no registro da cerimônia; e, antes de eliminar, confere que todo hash de registro selado da trilha tem objeto no `custody`, recusando eliminar se faltar algum.
- O `custody` **só apaga cifrados de período que tenham recibo de eliminação** assinado na cerimônia; não apaga por prazo próprio. Assim, a preservação vale sem que o servidor conheça o escopo dela.
- Painel e recibos dizem o que ocorreu: "cópia de trabalho expurgada; cópia de custódia guardada até DD/MM/AAAA" ou "cópia de custódia eliminada em DD/MM/AAAA". **Nunca** "destruído" para o que continua guardado (D27).
- **Sala comum e sala de sistema:** na conclusão ou no expurgo, expurgo total (ADR-04). Não há cópia.

### 8. Primitivas
HPKE (RFC 9180) da pilha do core e quórum de cartões nativo do HSM. **Nenhuma divisão de segredo implementada à mão** (regra 1). Se a pilha não tiver HPKE com a suíte exigida, parar e perguntar.

## Comandantes (detentores de parte)
- Cartão próprio com PIN, guardado em local separado; nunca em aparelho conectado.
- Não precisam ser membros das salas; não leem nada no dia a dia. Se podem ser membros ou gestor: pendente (D-R24, pergunta 4 da revisão).
- Credencial de segurança do grau mais alto tratado (doc 08 seção 5).
- Saída de um comandante: reemissão de cartões no HSM e revogação do antigo, registradas.
- Cerimônias de geração de lote, abertura, eliminação e troca da chave de cerimônia documentadas, com duas testemunhas e gravação, como a da AC raiz (doc 04 seção 6.4).
- No Exército, os N detentores vêm da cadeia de comando (D-R24).

## Alternativas rejeitadas
- **Chave simétrica de custódia por sala e período, conhecida pelos membros** (versão anterior): um aparelho comprometido mais o acervo do servidor abre o período inteiro; quebra sigilo futuro e recuperação pós-comprometimento (D01, D02).
- **HSM online de recuperação com aprovação remota por três celulares** (versão anterior): política dependente do diretório, que é hostil; sem tela confiável; aparelhos com falha comum; mesma chave embrulhada para o HSM mais fraco (D03–D05). Se a revinculação com histórico institucional voltar a ser requisito, exige **nova ADR** que atenda D03 (a)–(e) e D05.
- **Custódia em todas as salas, sem marca nem exceção** (D-R29 opção A): põe o perfil inteiro no acervo, inclusive salas de sistema; o responsável escolheu a opção B (GD-14).
- **Sala comum à escolha do criador, sem motivo nem segunda assinatura:** deixa criar sala comum para fugir da apuração.
- **Nenhuma custódia (MLS puro em tudo):** não atende investigação militar nem corregedoria quando o edital exigir.
- **Chave do Auditor ou de um administrador sozinho:** ponto único de abuso e de coação.
- **Divisão de segredo implementada no projeto:** viola a regra 1.
- **Cópia em claro no servidor:** viola a regra 3.

## Consequências
- Expurgo: em sala comum, destrói tudo; em sala sob custódia, a conclusão destrói a cópia de trabalho e a custódia segue o prazo (RN03, RN04 e doc 06 seção 8 atualizados).
- Sigilo futuro e recuperação pós-comprometimento (RNF-A4) valem para os membros. A cópia de custódia **não** tem sigilo futuro: comprometer o HSM offline, ou o quórum, expõe os períodos guardados das salas sob custódia por até 5 anos. Risco **admitido** pelo responsável pelo produto só nessas salas (GD-1); **aceite formal pendente** (R68; doc 04 seção 3.2) (01c-C18). Com a flag ligada, a custódia é o padrão (GD-14): o acervo tende a cobrir quase todas as salas da instituição.
- **Portão de produção (D08):** nenhuma sala sob custódia em produção antes de KEM híbrido na custódia, **ou** de aceite formal do R78 para a custódia pelo responsável pela homologação e pelo comitê (pergunta 13 do 01c).
- Novo serviço `custody` no `server/` (D11; doc 02 seção 7.1): dono das cópias, do lote de chaves públicas, das atribuições chave ↔ sala, dos pedidos, dos recibos e dos registros cifrados para quórum (inclusive segredos de lease guardados).
- Novo lugar com texto em claro: a estação isolada (D12).
- Custo no aparelho: uma cifra HPKE a mais por mensagem em sala sob custódia.
- "Colher agora, decifrar depois": acervo de até 5 anos sob suíte clássica até haver KEM híbrido (D08, R78). Nenhum grau classificado antes de D-R17 e D-R18 (G3, GD-9).
- Conluio de 3 comandantes (R79) e coação de detentores (R80) continuam como residuais. Sala comum criada para fugir da custódia: R111.
- Sem HSM online, os riscos R81, R82 e R83 deixam de se aplicar na v1.

## Revisar quando
O jurídico concluir D-R5 e D-R15; o HSM for escolhido (pergunta 14; 01c-C06, C11); D-R17 e D-R18 forem definidos (G3); D-R36 for decidida; a ADR-0010 for aceita ou mudar; ou o spike MLS (etapa 3) mostrar que a marca e a chave de período não cabem no estado do grupo.
