# 06 — LGPD e Conformidade

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** jurídico, encarregado de dados, corregedoria, área de segurança, equipe de desenvolvimento

> **Aviso.** Este documento organiza o assunto e traduz obrigações em requisitos técnicos. **Não é parecer jurídico.** Toda referência a artigo, resolução ou prazo deve ser conferida pelo jurídico da instituição na versão vigente; os pontos que dependem de interpretação estão marcados com **[validar]**.

---

## 0. Premissa institucional

**A instituição é a controladora e assume a responsabilidade pela privacidade dos usuários.** O projeto não usa a privacidade do usuário como motivo para deixar de oferecer um mecanismo de segurança, monitoramento ou atribuição, desde que o mecanismo tenha base legal e finalidade escritas pela instituição.

Os limites que **permanecem** não vêm da privacidade, e sim de quatro motivos práticos:

1. **Segurança operacional.** Todo dado coletado sobre agentes (imagem, posição, quem leu o quê) é alvo. Um banco com rostos e trajetos da equipe é o que o adversário mais quer, e pode custar integridade física. Por isso o que se coleta é **cifrado, com acesso restrito e prazo**, mesmo quando a coleta é permitida.
2. **Validade jurídica e probatória.** Prova obtida sem base legal, sem ciência prévia ou sem rito pode ser anulada num procedimento disciplinar ou penal. Transparência (seção 7) e rito (doc 05 seção 3.4) protegem a **própria apuração**.
3. **Obrigações da instituição** que existem independentemente do projeto (seção 1, **[validar]**).
4. **Terceiros no quadro ou nas conversas** (investigados, passantes), que não são usuários e não consentem.

Onde antes o texto dizia "evitar por privacidade", leia-se agora "decidir pela instituição, com os controles acima".

5. **Limites constitucionais que a instituição não renuncia.** A proteção de dados pessoais é direito fundamental (CF, art. 5º, LXXIX; STF, ADI 6387) e o STF afirmou que não há livre fluxo irrestrito de dados no Poder Público, exigindo finalidade, necessidade e proporcionalidade (ADPF 695 e ADI 6649). Portanto a premissa **não é "sem privacidade"**: é "a instituição responde, documenta e justifica cada tratamento". Verificação em docs/anexos/validacao-juridica-preliminar.md.

---

## 1. Enquadramento

### 1.1 Dois grupos de titulares, dois regimes possíveis

| Grupo | Quem | Dados | Regime provável |
|---|---|---|---|
| **A — Usuários do sistema** | Agentes da instituição (gestores, administradores, operadores, auditores) | Matrícula, codinome, aparelho, atestação, eventos de integridade, trilha de uso, registro de visualização, marca d'água | **LGPD aplica**, com a instituição como controlador no poder público (Capítulo IV da lei) **[validar]** |
| **B — Terceiros citados nas operações** | Investigados, testemunhas, vítimas, quem aparece em foto ou texto | Conteúdo das conversas, fotos, OCR, relatórios de IA | O art. 4º, III, da LGPD exclui da lei o tratamento para **segurança pública, defesa nacional, segurança do Estado e atividades de investigação e repressão de infrações penais**, que deve ser regido por **legislação específica** **[validar]** |

### 1.2 O que o art. 4º exige mesmo na exclusão
Segundo a própria lei, a legislação específica deve prever medidas proporcionais e estritamente necessárias ao atendimento do interesse público, observados o **devido processo legal**, os **princípios gerais de proteção** e os **direitos do titular**. A lei também limita o tratamento por pessoa de direito privado nesses casos **[validar]**.

### 1.3 Postura recomendada
Não depender da exclusão. **Aplicar a LGPD como padrão de referência para os dois grupos**, porque:
- o desenho do projeto (minimização, finalidade, segurança, retenção, auditoria) já a atende em grande parte;
- facilita o diálogo com corregedoria, controle externo e eventual regulamentação futura;
- reduz risco se o enquadramento do grupo B for contestado.

**Verificado em 2026-11-11** (docs/anexos/validacao-juridica-preliminar.md): não há lei específica em vigor; o PL 1515/2022 segue em tramitação na Câmara. A ANPD sustenta que os princípios de proteção de dados valem mesmo nas exceções do art. 4º (fonte secundária). Normas da atividade de inteligência a conferir: Lei 9.883/1999, Decretos 3.695/2000, 8.793/2016, 10.777/2021 e 10.778/2021.

**Pontos a confirmar com o jurídico:** quais dessas normas impõem exigências sobre sigilo e sistemas de informação; manifestações da ANPD aplicáveis ao tema **[validar]**.

---

## 2. Papéis

| Papel (LGPD) | Quem | Observações |
|---|---|---|
| **Controlador** | A instituição (pessoa jurídica de direito público) | Decide finalidade e meios |
| **Operador** | Quem trata em nome do controlador: equipe de desenvolvimento externa, fornecedor de MDM, empresa de auditoria | Contrato com cláusulas de proteção de dados |
| **Encarregado** | Pessoa indicada pela instituição | Canal com titulares e ANPD; assessora o projeto. Contato divulgado |
| **Autoridade** | ANPD | Quando aplicável ao grupo A |

Equipe de desenvolvimento **não** acessa dado de produção. Ambientes de teste usam dados fictícios (doc 03 seção 3.1).

---

## 3. Inventário de tratamento (registro das operações)

Preencher e manter atualizado. Base legal e retenção são **propostas** a validar.

| Dado | Titular | Finalidade | Base legal proposta | Onde fica | Retenção | Quem acessa |
|---|---|---|---|---|---|---|
| Matrícula, nome real | A | Vincular aparelho à pessoa para controle e responsabilização | Execução de políticas públicas / cumprimento de dever legal **[validar]** | Diretório: cifrada para a chave de quórum de pseudonimato; índice de consulta com chave fora do `server/` (D09) | Vínculo + prazo de apuração | Gestor, pelo índice, com consulta registrada; vínculo com o codinome só por quórum presencial (doc 05 seção 3.4) |
| Codinome | A | Identificação dentro das salas, sem expor a pessoa | Idem | Diretório; mensagens cifradas | Idem | Membros da sala |
| Aparelho, atestação, certificado | A | Segurança do acesso | Idem | Diretório | Vida do aparelho + prazo | Gestor; segurança |
| Eventos de integridade | A | Detectar uso inseguro | Idem; legítimo interesse não se aplica ao poder público em regra **[validar]** | Auditoria | **A definir (D-R4)** | Segurança; Auditor |
| Trilha de auditoria (metadados) | A | Controle institucional | Cumprimento de obrigação / controle | Auditoria | **A definir (D-R4)** | Auditor |
| Registro de visualização | A | Apuração de vazamento | Idem | Cifrado para a chave de quórum (ADR-0003) | Curta; **a definir** | Quórum presencial, pelo rito do doc 05, conduzido pelo Auditor |
| Marca d'água (identificador do aparelho) | A | Atribuição de vazamento | Idem | Em exibição; no conteúdo exportado | Enquanto o conteúdo existir | Corregedoria, sob rito |
| Biometria de desbloqueio | A | Autenticar uso | — | **Fica no sistema do aparelho**; o app só recebe sucesso ou falha | — | Ninguém no app |
| Metadados de mensagens (room_id aleatório, horários, volume) | A (e indiretamente B) | Entrega e operação | Execução de política pública | Backend, fila temporária | Até entrega ou 72 h | Sistema |
| Conteúdo das mensagens, fotos, anexos | B (e A) | Coordenação de atividade de inteligência | Atividade de investigação / legislação específica **[validar]** | Aparelhos, cache cifrado; fila cifrada | Política da operação (ex.: 24 h, 7 dias) | Membros da sala |
| OCR, trechos, vetores | B | Busca e apoio | Idem | Banco da IA, cifrado | Igual à operação | Nó de IA |
| Imagem da câmera frontal (detecção de presença e de câmera apontada) | A e terceiros no quadro | Segurança contra fotografia da tela | Processada em memória; **não persistida**; só evento (doc 07 seção 6) | Evento: **a definir** | Segurança, Auditor |
| Eventos de local (zona) e série atestada | A | Controle de aparelho e local autorizados | Sem coordenadas; retenção curta (doc 07 seções 3 e 4) | **A definir** | Segurança, Auditor |
| Relatórios de IA | B | Apoio à decisão | Idem | Cache do administrador | Igual à operação | Administrador |
| Cópia de custódia (só sala sob custódia) | B (e A) | Atender ordem judicial, ou processo administrativo com validação jurídica, sempre com CI | Legislação específica; investigação militar, corregedoria ou inquérito, conforme o perfil **[validar]** | Serviço `custody`, cifrada para a chave de quórum (ADR-0003) | Até 5 anos da conclusão **[validar]** | Quórum presencial |
| Termo de ciência e aceite da revinculação, com foto e assinatura | A | Registrar ciência e aceite da reposição do aparelho | Cumprimento de dever legal **[validar]**; foto é dado biométrico sensível (art. 11) **[validar]** | Cifrado para a chave de quórum do termo; hash na trilha; cópia exportada ao sistema corporativo, **sem a foto** até D-R30 | **A definir (D-R30)** | Quórum presencial, para responder a CI |
| Posição em modo perda | A | Recuperar o aparelho e proteger o agente | **[validar]** | Tempo real pela sala de resposta (MLS); cópia guardada só cifrada para a chave de quórum | Prazo da perda | Membros da sala de resposta; quórum presencial |
| Posição em missão ativa (ADR-0008, proposta; Fase 8) | A | Coordenação e segurança da equipe em campo | **[validar]** | Mensagens MLS da sala da missão; em sala sob custódia, também a cópia de custódia | Retenção da sala; cópia de custódia até 5 anos | Membros da sala (comandante); quórum presencial |
| Segredo de lease guardado de aparelho revogado | A (indireto) | Preservar o cache para apuração (GD-13) | **[validar]** | Serviço `custody`, cifrado para a chave de quórum | Prazo da apuração (D-R35) | Quórum presencial, com o aparelho em mãos |
| Número e hash do relatório de conclusão | — | Fundamentar a conclusão da sala (RN18) | — | Estado do grupo; hash na trilha | Vida da custódia | Membros; Auditor |
| Indicadores do nível 2: agregados derivados e nominais (ADR-0008, proposta; depois da v1) | B (presos, investigados) e A | Planejamento de operações | D-R34 **[validar]** | App do analista e sala de análise (MLS); índice do nó de IA | Nunca maior que o da sala de origem e o do escopo; expurgo em cascata | Analistas credenciados do escopo, só de salas com concessão (GD-15) |
| Concessão de escopo (ADR-0008) | A (administrador que concede) | Autorizar a entrada da sala num escopo de análise | Idem | Estado do grupo (cifrado); hash na trilha | Vida da sala | Membros da sala; Auditor (hash) |
| Registro de perguntas à IA de análise (ADR-0008) | A (analista) e B (citados) | Controle do uso da análise (R42) | **[validar]** | Cifrado para a chave de quórum de registros | Prazo do escopo + apuração | Quórum presencial |
| Gabarito facial do dono (GD-18; doc 07 seção 7.2) | A | Verificar quem acessa o aparelho | Dado biométrico sensível (art. 11) **[validar]** | Só no aparelho, cifrado por chave não exportável do hardware seguro; nunca no servidor, backup ou exportação | Vida do aparelho; zerado no nível 4 e na devolução | Ninguém: o core compara |
| Registro selado de investigação sigilosa (ADR-0010) | A (investigados, participantes) e B | Apuração de vazamento ou infiltrado sem alertar o investigado | **[validar]** | `custody`, cifrado para a chave de quórum de registros (série de investigação); só hash na trilha | D-R38 | Quórum presencial |
| Retrato de metadados preservado e relatório de rastro (ADR-0010) | A (e B) | Preservação e resultado da apuração | **[validar]** | Retrato: mídia offline cifrada para a chave de quórum; em sala comum, só metadados, nunca o cifrado MLS (GD-23). Relatório: pacote de evidência (RF67) ao procedimento e cópia no registro selado | D-R38. **Coleta nova (01g-G07d):** a preservação pode guardar pertença, horários e volumes além do prazo da seção 8.2 | Quórum presencial com a base da regra 14; procedimento |

Dado pessoal sensível (art. 5º, II) pode aparecer no conteúdo (origem racial, saúde, biometria de investigados etc.); o sistema **não classifica** o conteúdo, e por isso trata tudo como sigiloso por padrão.

---

## 4. Princípios (art. 6º) e como o projeto os atende

| Princípio | Como é atendido | Lacuna a fechar |
|---|---|---|
| **Finalidade** | Uso restrito a operações da unidade; norma de uso | Finalidade escrita no ato interno e no RIPD |
| **Adequação** | Funcionalidades ligadas à coordenação | Rever periodicamente novas funções contra a finalidade |
| **Necessidade** | Sem número de telefone; codinome; sem histórico permanente (exceto a cópia de custódia em sala sob custódia, com prazo máximo); retenção curta | Revisar campos do cadastro; cifrar matrícula e nome (D09) |
| **Livre acesso** | Usuários consultam seu cadastro e eventos próprios | Procedimento da seção 6 |
| **Qualidade** | Relatório de IA com citações e aviso de erro | Medir taxa de afirmações sem origem (E09) |
| **Transparência** | Aviso no convite sobre nó de IA; norma de uso | **Aviso explícito de monitoramento** (seção 7) |
| **Segurança** | Documentos 02 e 03 | Seção 9 |
| **Prevenção** | Matriz de riscos; testes | Doc 04 |
| **Não discriminação** | IA não decide sobre pessoas; saída revisada por humano | Teste de viés em amostra, quando houver dados reais |
| **Responsabilização** | Trilha; registro de tratamentos; RIPD | Manter inventário e RIPD vivos |

---

## 5. Relatório de Impacto à Proteção de Dados (RIPD)

Recomendado **antes da Fase 1 em produção e obrigatoriamente antes do piloto**. Nesta fase o projeto só guarda o roteiro; o texto é redigido pelo encarregado com a equipe.

### 5.1 Conteúdo mínimo
1. Descrição do tratamento e finalidades.
2. Necessidade e proporcionalidade de cada coleta (inventário da seção 3).
3. Riscos aos titulares e medidas (doc 04).
4. Salvaguardas e mecanismos de mitigação.
5. Revisão periódica.

### 5.2 Riscos que o RIPD deve tratar expressamente
| Risco | Observação |
|---|---|
| **Monitoramento de agentes** (integridade, marca d'água, registro de visualização) | Teste de proporcionalidade; transparência; acesso restrito; retenção curta |
| **Atribuição de vazamento** | Risco de acusação indevida (R36); apuração humana; contraditório; **decisão nunca automatizada** (art. 20 e princípio da transparência) |
| **IA analisando conversas** | Apoio à decisão; não substitui análise humana; sem decisão automática sobre pessoas |
| **Expurgo × dever de guarda** | Ver seção 8 |
| **Metadados** | Exposição a administrador de infraestrutura |
| **Terceiros nas conversas (grupo B)** | Impossibilidade de informar o titular por sigilo investigativo; salvaguardas compensatórias |
| **Investigação sigilosa de agentes** (ADR-0010) | Sem aviso ao investigado durante a apuração; CI, segunda autorização, impedimento, quórum e ordem judicial para conteúdo; contraditório depois; sem gravação oculta de sala comum |
| **Reconhecimento facial e destruição por perfil** (GD-18) | Dado biométrico sensível; falso negativo e falso positivo; destruição do aparelho no perfil missão de campo |

---

## 6. Direitos dos titulares

### 6.1 Grupo A (usuários)
| Direito (art. 18) | Como atender |
|---|---|
| Confirmação e acesso | Relatório de cadastro, aparelho, eventos de integridade e consultas feitas à sua identidade, observado o sigilo de investigação em curso, inclusive a sigilosa da ADR-0010 **[validar]** |
| Correção | Pelo gestor, com registro |
| Informação sobre compartilhamento | Declarar que dados não saem da instituição; quem acessa e em que condições |
| Eliminação e oposição | Limitadas por dever legal e controle; resposta fundamentada |
| Revisão de decisão automatizada | Não há decisão automatizada sobre pessoas; o documento 05 exige apuração humana |

Prazo e formato de resposta: seguir o regulamento aplicável **[validar]**. O encarregado coordena; o gestor executa.

### 6.2 Grupo B (terceiros)
Os direitos ficam condicionados à legislação específica e ao sigilo da investigação **[validar]**. O sistema não oferece canal direto para terceiros; qualquer requisição chega pela instituição e é tratada por procedimento formal.

---

### 6.3 Requisições e ordens de autoridades externas

- Toda ordem ou requisição de Ministério Público, polícia ou juízo chega **à instituição** e é tratada pelo responsável designado com o jurídico (playbook P11, doc 05). Nenhum operador ou administrador técnico responde diretamente.
- O servidor **não possui conteúdo em claro** (regras 2 e 3). O que pode existir para entregar: mensagens cifradas, metadados do servidor e registros de auditoria (doc 02 seção 6.4).
- **Sala comum:** o rito registra que o conteúdo é inacessível **por desenho**, não por recusa, **salvo** o cache de aparelho revogado em poder da instituição, pelo rito da regra 14, dentro do prazo do segredo guardado (ADR-0006 item 8; 01e-E20).
- **Sala sob custódia** (ADR-0003): a ordem judicial é **cumprida** por quórum presencial no HSM offline, no escopo pedido, com saída em pacote de evidência assinado. A **CI** é o comunicado interno que registra e tramita o pedido; **não substitui** a ordem judicial. Para pedido interno, a base é o processo administrativo com validação jurídica, registrado pela CI **[validar]**. A CI é obrigatória nos dois casos (01c-C23).
- A jurisprudência encontrada (STJ afastou multa por impossibilidade técnica de interceptar mensagens cifradas) é favorável a essa posição, mas **precisa de confirmação** nos acórdãos e não trata de sistema institucional.
- Conflito a decidir (D-R14): apreensão de aparelho por autoridade versus expurgo remoto.

## 7. Transparência com os usuários

O texto de convite e a norma de uso devem informar, **antes do primeiro uso**:
1. Que o aplicativo verifica a integridade do aparelho e envia o resultado ao servidor.
2. Que toda tela leva **marca d'água** que identifica o aparelho e a hora, e para que serve.
3. Que há **nó de IA** na sala e que o administrador pode pedir análise do conteúdo (já previsto). Se a ADR-0008 for aceita: que a sala só entra no escopo de uma **sala de análise** entre operações por **concessão** do administrador, que aparece na própria sala (GD-15), e que a sala de análise devolve indícios com referência e sem trecho de conversa a analistas credenciados.
4. Que há **trilha de auditoria** de metadados e quem pode consultá-la.
5. Que o **registro de visualização** existe, se for adotado (doc 05 seção 3.3).
   Que a **câmera frontal** é usada durante a leitura para detectar ausência, terceiros e câmera apontada, sem gravar imagem, e no **desbloqueio**, para verificar o rosto de quem acessa, com o gabarito só no aparelho (GD-18); e que há controle de **local** nas salas que o exigem (doc 07).
6. Em que situações a **identidade real** pode ser vinculada ao codinome e por qual rito.
7. Quanto tempo cada dado é guardado.
8. Como falar com o encarregado.
9. Se a sala é **sob custódia**: a marca aparece na sala; com a flag ligada, é o padrão, e sala comum só existe com motivo registrado e segunda assinatura (GD-14); há cópia institucional cifrada de **toda** mensagem, guardada até 5 anos a contar da conclusão (ou do expurgo, se não houver conclusão), aberta só por quórum presencial com CI e ordem judicial ou processo administrativo com validação jurídica; **apagar a mensagem para todos, a autodestruição, apagar ou concluir a operação não eliminam essa cópia** (D27; 01c-C09). Em sala de missão sob custódia, isso inclui a posição compartilhada pelo MLS (ADR-0008).
10. Na **revinculação**: o termo de ciência e aceite tem foto e assinatura, fica cifrado para a chave de quórum, com prazo, e é exportado ao sistema da corporação (doc 07 seção 7.4).
11. Na **revogação** do aparelho: o cache fica ilegível, mas não é destruído; o segredo que o abre fica guardado cifrado para a chave de quórum pelo prazo da apuração e só se abre por quórum presencial, com o aparelho em mãos (GD-13).
12. Em **missão de campo**, o aparelho pode se apagar sozinho quando houver indício de captura (GD-18), sem detalhar os sinais.
13. Que pode haver **investigação sigilosa** de vazamento ou infiltrado (ADR-0010), sem aviso durante a apuração: ela usa metadados, registros de atribuição e, com ordem judicial, a cópia de custódia das salas sob custódia; **nunca** grava sala comum nem acrescenta membro oculto; a preservação pode guardar metadados (não o conteúdo) por mais tempo que o normal (GD-23); quem for apontado tem contraditório no procedimento e, se a investigação for arquivada, é avisado como padrão (GD-22).

**Como a ciência é registrada:** o agente assina o **Termo de Ciência e Autorização de Acesso** (docs/anexos/termo-de-ciencia-e-autorizacao.md) pelo **gov.br**, na presença do gestor, no momento do cadastro. A assinatura ocorre **fora do app**, na estação do gestor: o aparelho não fala com o gov.br e o CPF não é digitado nele (regra 4); o app só mostra o texto, a versão e o estado que o core recebe do servidor (doc 11, tela T02; 01g-G14). A ativação do aparelho depende de assinatura válida da **versão vigente**. Nova versão = nova assinatura; sem ela o acesso é suspenso.

Cuidados de desenho:
- O termo **não leva codinome nem identificador de aparelho**; quem liga o termo ao usuário é a referência interna no diretório (R58, R40).
- Nenhum dado operacional vai ao gov.br: apenas o termo genérico, a matrícula e o CPF para a assinatura.
- O texto **não detalha** métodos, limiares ou respostas de detecção; descreve categorias. Detalhe técnico ajuda o adversário.
- A assinatura é validada pelo servidor (ou na estação do gestor) contra a cadeia de confiança aplicável, sem depender de consulta externa em tempo de uso.
- Nível da assinatura: a Lei 14.063/2020 (verificada) deixa o nível mínimo a **ato de cada ente**; o Decreto 10.543/2020 é **federal**. Como há informação sob sigilo, a assinatura simples não serve; usar **avançada (conta gov.br prata ou ouro) com ato interno e conferência presencial**, ou **qualificada (ICP-Brasil)**. Instituição estadual ou municipal precisa de ato próprio. Decisão do jurídico.

**Transparência pública mínima (LGPD arts. 23 e 41):** o poder público deve informar as hipóteses de tratamento e divulgar o encarregado. Publicar **aviso genérico** (finalidade, categorias de dados dos usuários, retenção, encarregado e canal), sem detalhe de defesa, local ou operação; o jurídico decide o texto e eventual dispensa.

Ciência do usuário **não substitui a base legal** e não deve ser tratada como consentimento livre, dada a relação funcional **[validar]**.


### 7.1 Plano de contingência da assinatura (gov.br fora do ar) — R59

**Regra geral:** o agente assina **sempre pelo gov.br**. A contingência existe só para quando o gov.br estiver indisponível ou a conta do agente não conseguir assinar, e **não** é escolha de conveniência.

| Via | Quando | Como | Validade do acesso |
|---|---|---|---|
| **1. gov.br** (principal) | Sempre que disponível | Conta prata ou ouro; assinatura na presença do gestor; PDF validado | Definitiva (vigência do termo) |
| **2. Assinatura qualificada ICP-Brasil** (contingência preferida) | gov.br fora do ar, ou conta do agente sem nível suficiente | Certificado do agente (token ou cartão funcional) na estação do gestor; validação **sem depender do gov.br** | Definitiva, se o ato interno admitir a qualificada |
| **3. Assinatura presencial com testemunha** (último recurso) | As duas anteriores indisponíveis e necessidade operacional justificada | Termo impresso assinado à mão, com conferência de documento por **dois servidores** (gestor e testemunha), digitalizado e anexado; ata registrada | **Provisória: até 15 dias** (proposta) |

**Gatilho:** o gestor declara a contingência no painel, com motivo; o evento vai à trilha de auditoria e o Auditor é avisado. Declaração automática se o teste de disponibilidade do serviço de assinatura falhar por mais de 4 horas úteis (proposta).

**Limites do acesso provisório (via 3):**
- sem papel de administrador de operação e sem acesso a salas de perfil "compartimento";
- prazo curto, com expiração automática no servidor (`provisional_until`);
- **ratificação**: o agente assina pelo gov.br (ou qualificada) em até 15 dias; sem ratificação, o aparelho volta a `pending_term`;
- cada acesso provisório é listado ao Auditor mensalmente.

**Renovação de nova versão do termo:** carência de 7 dias para assinar; se o gov.br estiver fora do ar nesse período, a carência é prorrogada até 7 dias após o retorno.

**Condições para valer:**
1. **Ato interno** da instituição que admita cada via e o nível (Lei 14.063/2020, art. 5º), a aprovar pelo jurídico **[validar]**; a via 3 exige análise do jurídico quanto à forma e ao arquivamento.
2. Estação do gestor com validação offline da cadeia ICP-Brasil.
3. Teste de cada via **antes** do piloto e semestralmente.
4. Registro de quem declarou, quem conferiu e qual via foi usada.

**O que a contingência não faz:** não relaxa a atestação do aparelho, o cadastro presencial nem a triagem de spyware. Só muda **como** a ciência é assinada.

---

## 8. Retenção, eliminação e dever de guarda

### 8.1 Conflito a resolver (D-R5) — **bloqueia a Fase 4**
O sistema foi desenhado para **destruir** conteúdo (expurgo criptográfico). **Verificado:** a Lei 8.159/1991 (arts. 7º, 9º e 25) trata como documento público o que órgãos produzem no exercício de suas funções, exige **autorização da instituição arquivística** para eliminar e prevê responsabilidade penal, civil e administrativa; o Código Penal, art. 314, pune quem inutiliza documento de que tem a guarda em razão do cargo. Antes de implementar o expurgo é preciso parecer, tabela de temporalidade aprovada ou ordem de preservação (docs/anexos/validacao-juridica-preliminar.md seção 3.3). Normas de **prestação de contas** e de **cadeia de custódia** também podem exigir preservar certas informações **[validar]**. O documento 01 já define o sistema como canal de coordenação, não repositório de prova; essa posição precisa ser **confirmada por escrito** pelo jurídico, junto com:
- o que deve ser levado ao sistema oficial antes de expirar (via exportação autorizada);
- quem é responsável por fazê-lo;
- se há hipótese de **suspensão do expurgo** (guarda legal por determinação) e como ela se implementa sem quebrar a criptografia.

### 8.2 Mapa de eliminação
| Dado | Prazo | Mecanismo |
|---|---|---|
| Fila de entrega | Até entrega ou 72 h | Exclusão |
| Cache do aparelho | Política da operação | Destruição de chave |
| Banco da IA | Igual à operação | Destruição de chave |
| Anexos cifrados | Igual à operação | Exclusão |
| Trilha de auditoria | **A definir** | Política de retenção própria; sem conteúdo |
| Registro de visualização | **A definir** (curta) | Exclusão automática |
| Logs de integridade | **A definir** | Exclusão automática |
| Cadastro de usuário | Vínculo funcional + prazo de apuração | Anonimização ou eliminação após o prazo |
| Cópia de trabalho de sala sob custódia | Conclusão da sala (RN18) | Destruição de chave (expurgo) |
| Cópia de custódia (sala sob custódia) | Até 5 anos da conclusão, ou do expurgo sem conclusão **[validar]** | Destruição das chaves no HSM e nos backups dele, em cerimônia com recibo; com autorização arquivística (D-R5) **[validar]** |
| Termo de revinculação com foto | **A definir (D-R30)** | Destruição da chave de quórum do termo |
| Segredo de lease guardado (aparelho revogado) | Prazo da apuração (D-R35), no período de vencimento | Destruição da chave do período da classe no HSM; destruição antecipada com duas autorizações, que só exclui o registro (ADR-0006) |
| Gabarito facial do dono | Vida do aparelho | Zerado no nível 4 e na devolução; nunca sai do aparelho |
| Registro selado de investigação e retrato preservado (ADR-0010) | D-R38 | Destruição da chave do período da série de investigação, em cerimônia com recibo |
| Qualquer item sob preservação da investigação sigilosa | Suspenso até a decisão | Ordem selada; o prazo volta a correr pelo saldo (ADR-0010 seção 8) |

### 8.3 Suspensão do expurgo (hipótese a decidir)
Se o jurídico concluir que pode haver retenção legal, o projeto precisa de uma **ordem de preservação** assinada que impeça o expurgo automático por prazo daquela operação. Isso é mudança de requisito (afeta RN03, RN04 e o fluxo da seção 6.2 do doc 02) e deve entrar por ADR **antes** de implementar o expurgo (Fase 4).

A custódia em sala sob custódia (ADR-0003) atende parte dessa hipótese sem quebrar a criptografia: a cópia fica guardada até 5 anos da conclusão, e a ordem de preservação suspende a eliminação. Em sala comum, continua valendo o parágrafo acima.

A **investigação sigilosa** (ADR-0010 seção 8) tem preservação própria, por ordem selada: suspende a eliminação e a contagem do prazo das cópias de custódia e dos registros selados do escopo e guarda retratos de metadados cifrados para o quórum. **Não** suspende o expurgo nos aparelhos nem no nó de IA: reter ali, às escondidas, conteúdo que os membros acreditam apagado equivaleria a gravação oculta (GD-17). **Decidido (GD-23):** "expurgo suspenso" é a eliminação da cópia de custódia; metadados de sala comum podem ser preservados, nunca o cifrado MLS, com o retrato selado para quórum e aberto pela regra 14; a transparência diz que a guarda pode se alongar (seção 7 item 13). O jurídico confere a base dessa guarda alongada **[validar]**.

---

## 9. Segurança e incidentes com dados pessoais

### 9.1 Medidas técnicas
Os controles dos documentos 02, 03 e 04 são as medidas de segurança do art. 46. Para o grupo A, acrescentar:
- cifra de coluna para matrícula e nome real;
- consulta a esses campos somente pelo rito do documento 05 seção 3.4;
- registro de toda consulta a dado pessoal.

### 9.2 Comunicação de incidente
- Incidente que possa acarretar risco ou dano relevante a titulares exige avaliação para **comunicação à ANPD e aos titulares** (art. 48, verificado). A Resolução CD/ANPD 15/2024, segundo fontes secundárias, fixa **3 dias úteis a contar do conhecimento** de que dados pessoais foram afetados, com complementação em 20 dias úteis; há divergência entre fontes sobre o prazo ao titular. **Conferir o texto oficial e se se aplica fora do alcance da LGPD** **[validar]**.
- Na avaliação de gravidade, o art. 48, § 3º considera se os dados foram tornados ininteligíveis a terceiros: a cifra ponta a ponta do projeto conta a favor.
- Para o grupo B e conteúdo sob sigilo, o jurídico decide **se, como e quando** comunicar, considerando o sigilo da investigação **[validar]**.
- O playbook de incidente (doc 05) inclui o encarregado desde a triagem e preserva a linha do tempo exigida para a comunicação.

---

## 10. Transferência internacional e terceiros

- Infraestrutura **on-premises ou em nuvem privada da instituição** (RNF-B4): sem transferência internacional por desenho.
- **Proibido:** telemetria externa, serviços de análise de erro de terceiros, push por serviço de terceiro sem decisão formal (D-R1), IA em nuvem, repositório de código com dados reais.
- Lojas de aplicativos públicas não são usadas; distribuição pelo MDM privado.
- Fornecedores (MDM, desenvolvimento, auditoria externa, pentest): contrato com cláusulas de confidencialidade, operador sob instruções do controlador, proibição de subcontratação sem aval, direito de auditoria e devolução ou destruição dos dados.

---

## 11. Monitoramento de usuários: teste de proporcionalidade

Antes de adotar cada mecanismo que observa o comportamento do agente, responder por escrito:

| Pergunta | Aplicação |
|---|---|
| Qual finalidade específica? | Ex.: atribuir vazamento de conteúdo sigiloso |
| É necessário? Há meio menos invasivo? | Ex.: marca d'água em vez de gravar tela |
| É adequado? Funciona para a finalidade? | Medir (testes F01–F03) |
| O usuário foi informado? | Seção 7 |
| Quem acessa e sob qual rito? | Doc 05 seção 3.4 |
| Por quanto tempo guarda? | Seção 8.2 |
| Que decisões decorrem dele? | Apuração humana; nunca punição automática |

O resultado entra no RIPD e é revisto a cada mudança.

---

## 12. Requisitos derivados para o produto

| ID | Requisito | Documento de origem |
|---|---|---|
| RL01 | Cifrar matrícula e nome real no diretório; acesso somente pelo rito de quebra de pseudonimato | Doc 05 seção 3.4 e 9.1 |
| RL02 | Registrar toda consulta a dado pessoal de usuário, com finalidade | 9.1 |
| RL03 | Retenção configurável e automática da trilha, do registro de visualização e dos logs de integridade | 8.2 |
| RL04 | Texto de transparência no convite e na norma de uso, com versão e data de ciência | 7 |
| RL05 | Relatório de dados do usuário (acesso do titular), gerado pelo gestor | 6.1 |
| RL06 | Ordem de preservação que suspende o expurgo de uma operação (se decidido) | 8.3 |
| RL07 | Nenhuma decisão sobre pessoa tomada automaticamente a partir de IA, marca ou integridade | 5.2 |
| RL08 | Nenhum dado de produção em ambiente de desenvolvimento ou teste | 2 |
| RL09 | Inventário de tratamento mantido e versionado em `docs/` | 3 |
| RL10 | Aparelho só ativa com termo assinado da versão vigente; assinatura validada; hash do PDF e versão registrados; revalidação a cada versão | 7 |
| RL11 | Termo de revinculação com foto: cifrado no aparelho para a chave de quórum, prazo definido, hash na trilha, exportação só pelo fluxo autorizado (sem a foto até D-R30) | 3, 8.2 |
| RL12 | Marca de custódia visível na sala e texto de transparência sobre a cópia institucional, o prazo e quem a abre | 7 |
| RL13 | Investigação sigilosa só pelo rito da ADR-0010; transparência genérica na norma e no termo; contraditório no procedimento | 7; ADR-0010 |
| RL14 | Gabarito facial só no aparelho, cifrado por chave do hardware seguro; quadros só em memória | 3; doc 07 seção 7.2 |

---

## 13. Checklist LGPD para a homologação

Acrescentar à seção 4 do documento 03:

- [ ] Parecer jurídico sobre o enquadramento dos grupos A e B (art. 4º)
- [ ] Encarregado indicado e contato divulgado
- [ ] Inventário de tratamento preenchido e aprovado
- [ ] RIPD elaborado, revisado pelo encarregado e aprovado pelo responsável
- [ ] Bases legais definidas por tratamento
- [ ] Texto de transparência aprovado e entregue aos usuários, com registro de ciência
- [ ] Termo de Ciência e Autorização aprovado pelo jurídico; fluxo de assinatura gov.br testado; nível de assinatura definido
- [ ] Termo sem codinome, sem identificador de aparelho e sem detalhe técnico de defesa
- [ ] Retenção definida para trilha, registro de visualização e logs de integridade
- [ ] Posição do jurídico sobre expurgo × dever de guarda documentada
- [ ] Procedimento de atendimento a direitos de titulares
- [ ] Procedimento de comunicação de incidente com decisão sobre prazos e destinatários
- [ ] Cláusulas de proteção de dados nos contratos com fornecedores
- [ ] Confirmação de que não há transferência internacional nem telemetria externa
- [ ] Verificação técnica: matrícula e nome cifrados; consultas registradas; retenção automática funcionando

---

## 14. Pontos para validação jurídica

1. Enquadramento de cada grupo de titulares e a norma específica aplicável à atividade de inteligência.
2. Base legal de cada tratamento do grupo A (especialmente monitoramento de integridade e registro de visualização).
3. Aplicação de obrigações de comunicação de incidente a este sistema.
4. Compatibilidade do expurgo com guarda documental e prestação de contas; hipótese de ordem de preservação.
5. Validade jurídica e probatória da atribuição por marca d'água e do registro de visualização; rito de contraditório.
6. Natureza da ciência do usuário e conteúdo da norma de uso.
7. Retenção da trilha de auditoria e dos logs.
8. Responsabilidade do operador (desenvolvedor, MDM) e cláusulas contratuais.
9. Custódia: processo administrativo com validação jurídica (registrado por CI) como base para abrir; prazo de 5 anos a contar da conclusão; eliminação e autorização arquivística (Lei 8.159); motivo por perfil (investigação militar; corregedoria ou inquérito).
10. Termo de revinculação com foto: base para dado sensível (art. 11), prazo de guarda, exportação ao sistema corporativo (D-R30).
11. Investigação sigilosa (ADR-0010): base para abrir registros selados sem ordem judicial (processo administrativo com validação jurídica); CI que tramita como sigilosa; limite do direito de acesso do investigado durante a apuração; contraditório; aviso no arquivamento (D-R38); MP ou MPM como destinatário; Tema 977 do STF para o cache de aparelho.
12. Reconhecimento facial (GD-18): base para o gabarito (art. 11) e regras da ANPD; uso da câmera no desbloqueio, além das cautelas da seção 3.5 do anexo de validação; destruição automática do aparelho em missão de campo.
13. Preservação de metadados de sala comum na investigação sigilosa frente ao expurgo (ADR-0010 seção 8).
