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
| Matrícula, nome real | A | Vincular aparelho à pessoa para controle e responsabilização | Execução de políticas públicas / cumprimento de dever legal **[validar]** | Diretório (coluna cifrada) | Vínculo + prazo de apuração | Gestor; Auditor sob rito |
| Codinome | A | Identificação dentro das salas, sem expor a pessoa | Idem | Diretório; mensagens cifradas | Idem | Membros da sala |
| Aparelho, atestação, certificado | A | Segurança do acesso | Idem | Diretório | Vida do aparelho + prazo | Gestor; segurança |
| Eventos de integridade | A | Detectar uso inseguro | Idem; legítimo interesse não se aplica ao poder público em regra **[validar]** | Auditoria | **A definir (D-R4)** | Segurança; Auditor |
| Trilha de auditoria (metadados) | A | Controle institucional | Cumprimento de obrigação / controle | Auditoria | **A definir (D-R4)** | Auditor |
| Registro de visualização | A | Apuração de vazamento | Idem | Cifrado para o Auditor | Curta; **a definir** | Auditor, sob rito do doc 05 |
| Marca d'água (identificador do aparelho) | A | Atribuição de vazamento | Idem | Em exibição; no conteúdo exportado | Enquanto o conteúdo existir | Corregedoria, sob rito |
| Biometria de desbloqueio | A | Autenticar uso | — | **Fica no sistema do aparelho**; o app só recebe sucesso ou falha | — | Ninguém no app |
| Metadados de mensagens (room_id aleatório, horários, volume) | A (e indiretamente B) | Entrega e operação | Execução de política pública | Backend, fila temporária | Até entrega ou 72 h | Sistema |
| Conteúdo das mensagens, fotos, anexos | B (e A) | Coordenação de atividade de inteligência | Atividade de investigação / legislação específica **[validar]** | Aparelhos, cache cifrado; fila cifrada | Política da operação (ex.: 24 h, 7 dias) | Membros da sala |
| OCR, trechos, vetores | B | Busca e apoio | Idem | Banco da IA, cifrado | Igual à operação | Nó de IA |
| Imagem da câmera frontal (detecção de presença e de câmera apontada) | A e terceiros no quadro | Segurança contra fotografia da tela | Processada em memória; **não persistida**; só evento (doc 07 seção 6) | Evento: **a definir** | Segurança, Auditor |
| Eventos de local (zona) e série atestada | A | Controle de aparelho e local autorizados | Sem coordenadas; retenção curta (doc 07 seções 3 e 4) | **A definir** | Segurança, Auditor |
| Relatórios de IA | B | Apoio à decisão | Idem | Cache do administrador | Igual à operação | Administrador |

Dado pessoal sensível (art. 5º, II) pode aparecer no conteúdo (origem racial, saúde, biometria de investigados etc.); o sistema **não classifica** o conteúdo, e por isso trata tudo como sigiloso por padrão.

---

## 4. Princípios (art. 6º) e como o projeto os atende

| Princípio | Como é atendido | Lacuna a fechar |
|---|---|---|
| **Finalidade** | Uso restrito a operações da unidade; norma de uso | Finalidade escrita no ato interno e no RIPD |
| **Adequação** | Funcionalidades ligadas à coordenação | Rever periodicamente novas funções contra a finalidade |
| **Necessidade** | Sem número de telefone; codinome; sem histórico permanente; retenção curta | Revisar campos do cadastro; cifrar matrícula e nome |
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

---

## 6. Direitos dos titulares

### 6.1 Grupo A (usuários)
| Direito (art. 18) | Como atender |
|---|---|
| Confirmação e acesso | Relatório de cadastro, aparelho, eventos de integridade e consultas feitas à sua identidade |
| Correção | Pelo gestor, com registro |
| Informação sobre compartilhamento | Declarar que dados não saem da instituição; quem acessa e em que condições |
| Eliminação e oposição | Limitadas por dever legal e controle; resposta fundamentada |
| Revisão de decisão automatizada | Não há decisão automatizada sobre pessoas; o documento 05 exige apuração humana |

Prazo e formato de resposta: seguir o regulamento aplicável **[validar]**. O encarregado coordena; o gestor executa.

### 6.2 Grupo B (terceiros)
Os direitos ficam condicionados à legislação específica e ao sigilo da investigação **[validar]**. O sistema não oferece canal direto para terceiros; qualquer requisição chega pela instituição e é tratada por procedimento formal.

---

## 7. Transparência com os usuários

O texto de convite e a norma de uso devem informar, **antes do primeiro uso**:
1. Que o aplicativo verifica a integridade do aparelho e envia o resultado ao servidor.
2. Que toda tela leva **marca d'água** que identifica o aparelho e a hora, e para que serve.
3. Que há **nó de IA** na sala e que o administrador pode pedir análise do conteúdo (já previsto).
4. Que há **trilha de auditoria** de metadados e quem pode consultá-la.
5. Que o **registro de visualização** existe, se for adotado (doc 05 seção 3.3).
   Que a **câmera frontal** é usada durante a leitura para detectar ausência, terceiros e câmera apontada, sem gravar imagem, e que há controle de **local** nas salas que o exigem (doc 07).
6. Em que situações a **identidade real** pode ser vinculada ao codinome e por qual rito.
7. Quanto tempo cada dado é guardado.
8. Como falar com o encarregado.

**Como a ciência é registrada:** o agente assina o **Termo de Ciência e Autorização de Acesso** (docs/anexos/termo-de-ciencia-e-autorizacao.md) pelo **gov.br**, na presença do gestor, no momento do cadastro. A ativação do aparelho depende de assinatura válida da **versão vigente**. Nova versão = nova assinatura; sem ela o acesso é suspenso.

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

### 8.3 Suspensão do expurgo (hipótese a decidir)
Se o jurídico concluir que pode haver retenção legal, o projeto precisa de uma **ordem de preservação** assinada que impeça o expurgo automático por prazo daquela operação. Isso é mudança de requisito (afeta RN03, RN04 e o fluxo da seção 6.2 do doc 02) e deve entrar por ADR **antes** de implementar o expurgo (Fase 4).

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
