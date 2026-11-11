# 06 — LGPD e Conformidade

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** jurídico, encarregado de dados, corregedoria, área de segurança, equipe de desenvolvimento

> **Aviso.** Este documento organiza o assunto e traduz obrigações em requisitos técnicos. **Não é parecer jurídico.** Toda referência a artigo, resolução ou prazo deve ser conferida pelo jurídico da instituição na versão vigente; os pontos que dependem de interpretação estão marcados com **[validar]**.

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

**Pontos a confirmar com o jurídico:** quais normas específicas se aplicam hoje à atividade de inteligência policial da instituição; status atual de qualquer lei de proteção de dados para segurança pública e persecução penal **[validar]**; manifestações e relatórios da ANPD aplicáveis ao tema **[validar]**.

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
6. Em que situações a **identidade real** pode ser vinculada ao codinome e por qual rito.
7. Quanto tempo cada dado é guardado.
8. Como falar com o encarregado.

Ciência do usuário **não substitui a base legal** e não deve ser tratada como consentimento livre, dada a relação funcional **[validar]**.

---

## 8. Retenção, eliminação e dever de guarda

### 8.1 Conflito a resolver (D-R5)
O sistema foi desenhado para **destruir** conteúdo (expurgo criptográfico). Pode haver normas de **guarda documental e de arquivos públicos**, de **prestação de contas** e de **cadeia de custódia** que exijam preservar certas informações **[validar]**. O documento 01 já define o sistema como canal de coordenação, não repositório de prova; essa posição precisa ser **confirmada por escrito** pelo jurídico, junto com:
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
- Incidente que possa acarretar risco ou dano relevante a titulares exige avaliação para **comunicação à ANPD e aos titulares** (art. 48). Existe regulamento da ANPD sobre o tema, com prazo em dias úteis; **conferir o texto e o prazo atuais** **[validar]**.
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

---

## 13. Checklist LGPD para a homologação

Acrescentar à seção 4 do documento 03:

- [ ] Parecer jurídico sobre o enquadramento dos grupos A e B (art. 4º)
- [ ] Encarregado indicado e contato divulgado
- [ ] Inventário de tratamento preenchido e aprovado
- [ ] RIPD elaborado, revisado pelo encarregado e aprovado pelo responsável
- [ ] Bases legais definidas por tratamento
- [ ] Texto de transparência aprovado e entregue aos usuários, com registro de ciência
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
