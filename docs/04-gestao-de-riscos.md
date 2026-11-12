# 04 — Gestão de Riscos

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** responsável pela homologação, área de segurança, comitê de governança, equipe de desenvolvimento

> Este documento transforma a matriz do documento 03 (seção 2) em um **processo contínuo** e acrescenta os riscos que os documentos 01–03 não cobriam. Resposta a incidentes e investigação de vazamentos estão no documento 05; proteção de dados pessoais, no documento 06.

---

## 1. Lacunas encontradas nos documentos anteriores

A análise dos documentos 01–03 mostrou boa cobertura de **controles técnicos preventivos** e pouca cobertura de **processo**. O que faltava:

| # | Lacuna | Onde passa a ser tratada |
|---|---|---|
| L1 | Não havia processo de gestão de risco: a matriz era uma fotografia, sem dono, revisão nem prazo de aceite | Seções 2 a 5 deste documento |
| L2 | Nenhum plano de **resposta a incidentes** (só "procedimento para perda de aparelho" como item de checklist) | Documento 05, seção 5 |
| L3 | Nenhum método para **identificar a origem de um vazamento** depois que ele ocorre; a marca d'água existia, o procedimento de uso não | Documento 05, seções 2 a 4 |
| L4 | **Notificações**: "sem conteúdo" ainda exige um gatilho; o caminho padrão do Android usa serviço de terceiro (FCM), que vê metadados e contraria a premissa de infraestrutura própria | Risco R27; decisão D-R1 (seção 4) |
| L5 | **Insider privilegiado**: gestor do sistema que cadastra aparelho fantasma; administrador de infraestrutura | R28, R38; seção 6 |
| L6 | **Ciclo de vida do usuário**: desligamento, transferência, afastamento | R39 |
| L7 | **Gestão de chaves da instituição**: cerimônia da raiz da CA, rotação, recuperação | R31; seção 6 |
| L8 | **Continuidade e recuperação**: backup, RTO/RPO, teste de restauração | Seção 8 |
| L9 | **Cadeia de entrega do software**: pipeline de compilação, assinatura do app, SBOM | R33, R34 |
| L10 | **Gestão de vulnerabilidades e de mudanças** com prazos | Seção 7 |
| L11 | **Análise de tráfego e metadados** (quem fala com quem, quando) | R35 |
| L12 | **Falsa atribuição**: risco de acusar a pessoa errada por causa da marca d'água | R36 |
| L13 | **Proteção de dados pessoais** (LGPD) | Documento 06 |
| L14 | **Monitoramento e detecção** de ataques em andamento (só havia prevenção) | Documento 05, seção 4 |
| L15 | **Classificação da informação** e regra de compartimentação (quem precisa saber) | Seção 6 |
| L16 | **Exercícios**: nenhum teste de mesa, simulação de vazamento ou teste de restauração | Documento 05, seção 8 |

---

## 2. Papéis

| Papel | Responsabilidade |
|---|---|
| **Responsável pela homologação** (dono do risco) | Aceita riscos residuais por escrito; responde pelo sistema perante a instituição |
| **Comitê de segurança** | Revisa o registro de riscos; aprova mudanças que alterem o perfil de risco |
| **Encarregado de dados** | Opina sobre riscos a dados pessoais; ver documento 06 |
| **Dono do ativo** | Mantém o inventário e os controles do seu ativo (servidor, app, nó de IA, CA) |
| **Auditor** | Verifica, de forma independente, se os controles existem e funcionam |

Quem desenvolve **não** aceita risco do que desenvolveu.

---

## 3. Método

Base: ISO/IEC 27005 e NIST SP 800-30, simplificados. A escala é a do documento 03: probabilidade (P) e impacto (I) em Baixo / Médio / Alto / Crítico.

### 3.1 Ciclo

1. **Identificar** — a cada fase do roteiro, a cada mudança de arquitetura, a cada incidente e em revisão trimestral.
2. **Analisar** — P, I e os controles existentes; calcular o residual.
3. **Tratar** — uma das quatro opções, sempre registrada:
   - **Mitigar:** novo controle, com dono e prazo.
   - **Evitar:** remover a funcionalidade ou o requisito.
   - **Transferir:** seguro, contrato ou outra unidade responsável.
   - **Aceitar:** apenas com aceite formal e validade.
4. **Monitorar** — indicadores (seção 9) e revisão.

### 3.2 Regras de aceite

| Residual | Quem aceita | Validade |
|---|---|---|
| Baixo | Dono do ativo | Revisão anual |
| Médio | Responsável pela homologação | 6 meses |
| Alto | Responsável pela homologação **e** comitê, com plano de redução | 3 meses |
| Crítico | **Não pode ser aceito.** Bloqueia o uso real | — |

### 3.3 Campos do registro de riscos

`ID · descrição · ativo · ameaça · P · I · controles existentes · residual · tratamento · dono · prazo · aceite (quem, data, validade) · última revisão · incidentes relacionados`

O registro vive em `docs/riscos/registro.md` (ou ferramenta da instituição) e é versionado. A matriz do documento 03 seção 2 e a seção 5 deste documento são a carga inicial.

---

## 4. Decisões pendentes que viram risco

| ID | Decisão | Por que importa |
|---|---|---|
| **D-R1** | Como entregar o aviso de "nova mensagem" | Ver R27. Opções: (a) conexão persistente mantida por serviço em primeiro plano, gerenciado pelo MDM, sem terceiros; (b) push via FCM com carga vazia, aceitando que o Google vê quem recebe aviso e quando; (c) gateway de push próprio (exige que o aparelho aceite). Recomendação inicial: **(a)**. Medir impacto em bateria na Fase 0 |
| **D-R2** | Quebra de pseudonimato (codinome → pessoa real) | Quem pode pedir, quem autoriza, quantas assinaturas. Proposta no documento 05, seção 3.4 |
| **D-R3** | Variação de conteúdo por destinatário (rastreamento textual) | Ajuda a achar quem vazou texto copiado, mas altera o conteúdo entregue. Ver documento 05, seção 3.2 |
| **D-R4** | Retenção da trilha de auditoria e dos logs de integridade | Conflito possível entre minimização (LGPD) e controle. Ver documento 06 |
| **D-R5** | Conflito entre expurgo e dever de guarda documental | Jurídico precisa decidir o que pode ser destruído. Ver documento 06 |
| **D-R14** | Expurgo remoto após apreensão por autoridade: o que o sistema faz (padrão proposto: suspender e revogar, sem destruir) | Jurídico e comitê; ver R63 e anexo de validação seção 3.7 |
| **D-R6 a D-R13** | Câmera, local, coação, triagem de spyware, atestação, pós-quântico, check-in | Ver documento 07 seção 13; riscos R45–R56, R66–R71 e R79–R83 no documento 07 seção 10 |
| **D-R15** | **Definido: quórum de no mínimo 3 comandantes (ADR-0003, proposta); falta o jurídico.** Custódia institucional de histórico: recuperação por administrador em tela e entrega só por ordem judicial, com prazo máximo | Chefia + jurídico + segurança; doc 07 seções 7.3 e 7.4; exige ADR-0003 |
| **D-R28** | Orquestração da instalação: Compose ou systemd (proposta) ou Kubernetes | Segurança + infraestrutura da instituição; doc 02 seção 10.1 |
| **D-R25 a D-R27** | Rádios e redes de campo, teto do lease de campo, cifra do rádio como segunda camada | Ver documento 10 seção 9; riscos R89–R93 no documento 10 seção 8 |
| **D-R21 a D-R24** | Instância dedicada por instituição, operação em campo, rito militar do termo, detentores do quórum | Ver documento 09 seção 7; riscos R84–R88 no documento 09 seção 6 |
| **D-R17 a D-R20** | Grau máximo, algoritmo de Estado, aparelho e sistema, proteção de metadados | Ver documento 08 seção 7; riscos R72–R78 no documento 08 seção 6 |
| **D-R16** | Persistência de matrícula depois do reset e rastreio, frente à regra 10 | Segurança + chefia; doc 07 seção 7.4 |

---

## 5. Riscos adicionais (continuação da matriz do documento 03)

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R27 | Aviso de nova mensagem passa por serviço de terceiro e revela quem recebe e quando | Alto | Médio | Decisão D-R1; carga sempre vazia; sem identificador de sala | **Médio** até D-R1 |
| R28 | Gestor do sistema cadastra aparelho ou usuário fantasma para obter acesso | Baixo | Crítico | Cadastro presencial com duas pessoas; inventário do MDM cruzado com o diretório; Auditor revisa cadastros por amostragem; alerta de cadastro fora do padrão | Médio |
| R29 | Administrador de operação inclui ou remove membros de forma abusiva | Baixo | Alto | Convites e remoções na trilha; notificação ao membro; revisão pelo Auditor | Baixo |
| R30 | Engenharia social ou phishing contra gestor ou administrador | Médio | Crítico | Painel só em estação gerenciada; chave física de autenticação; chave de administrador condicionada a biometria; treinamento; confirmação fora de banda para ações críticas | Médio |
| R31 | Chave raiz da CA ou de emissão comprometida | Baixo | Crítico | Raiz offline em HSM; cerimônia com testemunhas e registro; emissão por intermediária revogável; plano de rotação ensaiado | Baixo |
| R32 | Ransomware ou destruição dos servidores | Médio | Alto | Backup do banco do backend, isolado e imutável; restauração testada; reconstrução automatizada; nó de IA sem backup por projeto (perda aceita) | Médio |
| R33 | Pipeline de compilação comprometido (app adulterado antes da assinatura) | Baixo | Crítico | Compilação reprodutível; CI isolado; duas aprovações para release; chave de assinatura em HSM; comparação de hash do binário distribuído | Médio |
| R34 | Vulnerabilidade em dependência crítica (biblioteca MLS, WebView, Go, Rust, runtime Python) sem correção a tempo | Médio | Alto | SBOM; monitoramento de CVE; prazos da seção 7; atualização testada em homologação | Médio |
| R35 | Correlação de metadados: tamanho, horário e volume de mensagens revelam a operação | Médio | Médio | Preenchimento (padding) para tamanhos fixos; envio em lotes; tráfego de cobertura opcional; testes B05 | Médio |
| R36 | **Acusação indevida**: marca d'água aponta aparelho que não é de quem vazou (aparelho emprestado, esquecido ou usado sob coação) | Médio | Alto | Marca é indício, não prova; corroboração obrigatória; dois analistas; contraditório; graus de confiança (documento 05) | Médio |
| R37 | Vazamento por transcrição manual ou memorização | Alto | Alto | Compartimentação por canal; norma de uso; variação por destinatário (D-R3); retenção curta | Alto — **não eliminável** |
| R38 | Administrador de infraestrutura exfiltra dados (dump de memória, backup, tráfego) | Baixo | Crítico | Servidor só tem texto cifrado; separação de funções; acesso privilegiado com aprovação e registro; HSM; nó de IA com acesso restrito e gravado | Médio |
| R39 | Usuário desligado, transferido ou afastado mantém acesso | Médio | Alto | Integração com RH ou rito manual com prazo de revogação (proposto: 4 h úteis); revisão mensal de usuários ativos × quadro | Baixo |
| R40 | Exposição de dados de cadastro (matrícula ↔ codinome ↔ aparelho) | Baixo | Alto | Tabela com acesso mínimo; cifra de coluna; consulta registrada; quebra de pseudonimato só pelo rito do documento 05 | Baixo |
| R41 | Erro humano: expurgo acidental, convite para pessoa errada, exportação indevida | Médio | Alto | Confirmação reforçada (RF51); dupla custódia opcional; revisão de convite antes de enviar; trilha | Médio |
| R42 | Uso desviado da análise de IA (investigar pessoas ou assuntos fora da finalidade) | Baixo | Alto | Pedido assinado e registrado com finalidade; Auditor revê pedidos; relatório com marca d'água | Médio |
| R43 | Relógio do servidor ou da rede adulterado (afeta lease, validade e ordens de expurgo) | Baixo | Alto | Fonte de tempo interna redundante; tolerância configurada; contador de sequência nas ordens (não depende só de hora) | Baixo |
| R44 | Vazamento por captura de cópia do material em trânsito entre sistemas oficiais (exportação autorizada) | Médio | Alto | Marca d'água; autorização individual; auditoria; canal oficial de entrega | Médio |
| R60 | **Expurgo configura eliminação ilícita de documento público** (Lei 8.159, arts. 9º e 25; Código Penal, art. 314) | Médio | Alto | Parecer, tabela de temporalidade aprovada ou ordem de preservação **antes** da Fase 4; exportação do que for documento oficial; registro do fundamento de cada expurgo | **Alto até o parecer**; bloqueia a Fase 4 |
| R61 | Prova extraída do sistema anulada por falta de base normativa, ciência ou rito | Médio | Alto | Norma de uso e termo prevendo a coleta; contraditório; cadeia de custódia; indício, não prova (doc 05) | Médio |
| R62 | **Ordem judicial ou requisição externa que o sistema não pode cumprir** (conteúdo inexistente no servidor por desenho) → multa, pedido de bloqueio, tensão institucional | Médio | Médio | Rito e responsável institucional para requisições; resposta padrão do que existe (cifrado, metadados); jurídico no circuito; playbook P11 (doc 05) | Médio |
| R63 | **Expurgo remoto após apreensão por autoridade** lido como destruição ou supressão de prova | Baixo | Alto | Em apreensão por autoridade: suspender e revogar, sem destruir, até decisão do jurídico (D-R14); registro do evento | **Alto até D-R14** |
| R64 | **Print ou cópia como prova de vazamento** contestada (adulterável, sem cadeia) | Alto | Médio | Pacote de evidência assinado (RF67); proibir print como prova; cadeia de custódia (doc 05) | Médio |
| R65 | **Segundo cliente ou espelhamento** criado no futuro (web, desktop, aparelho companheiro) reabre a controvérsia jurídica e o vetor de ataque | Baixo | Alto | Requisito de que não existe vínculo de segundo cliente; mudança exige ADR e novo parecer | Baixo |
| R94 | Pacote de instalação adulterado no caminho até a instituição | Baixo | Crítico | Assinatura em HSM com duas aprovações; verificador antes de instalar; build reproduzível; mídia controlada | Baixo |
| R95 | Segredo vazado em arquivo de configuração | Médio | Alto | Nenhum segredo no YAML ou no pacote; varredura de segredos no pipeline; chaves no HSM | Baixo |
| R57 | Conta gov.br de agente tomada por terceiro, que assina o termo em nome dele | Baixo | Alto | Assinatura na presença do gestor, com documento conferido; nível mínimo de conta definido; ativação exige também aparelho íntegro | Baixo |
| R58 | Termo ou fluxo de assinatura revela vínculo agente ↔ sistema, ou detalhes das defesas | Médio | Alto | Termo sem codinome, sem identificador de aparelho e sem detalhe técnico; PDF em guarda segura do gestor; nada operacional enviado ao gov.br | Médio |
| R59 | Indisponibilidade do gov.br impede cadastro ou renovação de termo | Médio | Médio | Plano de contingência em três vias (gov.br, ICP-Brasil, presencial com testemunha e acesso provisório), com gatilho, limites e ratificação (doc 06 seção 7.1); versão do termo muda raramente; carência na renovação | Baixo |

Todos os riscos com residual **Médio** ou **Alto** exigem aceite conforme a seção 3.2. O **R37** hoje fica em **Alto** e precisa de decisão do comitê: o projeto não tem como eliminá-lo, só reduzir o alcance e aumentar a chance de atribuição.

---

## 6. Controles organizacionais exigidos

### 6.1 Classificação e compartimentação
- Toda operação tem nível de sigilo definido pelo administrador na criação entre os graus legais (reservado, secreto; ultrassecreto fora do escopo, D-R17), conforme a Lei 12.527/2011 e o doc 08.
- **Necessidade de saber:** operador só entra em canais de que precisa. Quanto menos pessoas veem uma mensagem, menor o conjunto de suspeitos num vazamento. Canais com membros distintos dentro da mesma operação já estão previstos para a v2; avaliar antecipar.
- Relatórios de IA herdam o nível mais alto das salas analisadas.

### 6.2 Separação de funções
| Função | Não pode acumular com |
|---|---|
| Gestor do sistema (cadastra) | Administrador de operação; Auditor |
| Administrador de infraestrutura | Gestor do sistema; Auditor |
| Auditor | Qualquer papel operacional |
| Quem desenvolve | Quem aprova release; quem opera produção |

### 6.3 Acesso privilegiado à infraestrutura
- Sem contas compartilhadas; acesso individual e nominal.
- Acesso a produção por solicitação aprovada, com prazo e gravação de sessão.
- Operações destrutivas (apagar banco, alterar autoridade certificadora) exigem duas pessoas.

### 6.4 Gestão de chaves da instituição
| Chave | Guarda | Rotação | Recuperação |
|---|---|---|---|
| Raiz da CA | HSM offline, cofre | Longa (ex.: 10 anos); cerimônia registrada | Duas cópias em locais distintos, divisão de segredo entre custodiantes |
| CA intermediária de emissão | HSM online | Anual | Reemissão pela raiz |
| Chave mestra do nó de IA | HSM/TPM | Anual; reembrulha só as chaves das salas | Sem recuperação por projeto: perda = salas da IA ilegíveis (aceito) |
| Chave de assinatura do app | HSM | Conforme política da loja privada | Procedimento escrito |
| Chaves de administrador de operação | Aparelho, por biometria | Ao trocar de aparelho | Nova chave registrada por gestor; ordens antigas continuam válidas só até a revogação |

### 6.5 Ciclo de vida do usuário
- **Entrada:** cadastro presencial (RN05); termo de ciência; treinamento.
- **Transferência:** revisão de operações; remoção das operações sem necessidade.
- **Saída ou afastamento:** revogação de aparelho em até 4 h úteis; devolução do aparelho; remoção das operações; registro.
- **Revisão mensal:** lista de usuários ativos conferida com o quadro.

---

## 7. Gestão de vulnerabilidades e de mudanças

| Severidade (CVSS ou avaliação interna) | Prazo para correção em produção |
|---|---|
| Crítica, com exploração conhecida | 72 h (mitigação em 24 h) |
| Alta | 15 dias |
| Média | 60 dias |
| Baixa | Próximo ciclo |

Os prazos são ponto de partida para validar com a área de segurança.

- **SBOM** gerado a cada release; comparado a bases de vulnerabilidades continuamente.
- Todo componente de segurança (MLS, TLS, SQLCipher, WebView) tem dono e é acompanhado em seus anúncios.
- **Mudanças** em criptografia, autorização, fronteiras de serviço ou fluxo de expurgo exigem ADR, revisão de segurança e atualização do registro de riscos antes de entrar em produção.
- Toda release passa por: testes automatizados, `check-boundaries`, análise estática, verificação de dependências e aprovação de duas pessoas.

---

## 8. Continuidade e recuperação

Valores propostos, a validar com a unidade (ponto em aberto 6 do documento 01):

| Serviço | RTO (voltar a operar) | RPO (perda tolerada) | Observação |
|---|---|---|---|
| Gateway e entrega | 1 h | 0 (fila é reconstruída pelo cliente) | Redundância ativa |
| Diretório | 1 h | 15 min | Banco com replicação |
| Anexos | 4 h | 24 h | Perda de anexo = reenvio |
| Auditoria | 4 h | 0 | Não pode perder evento; fila local no emissor |
| Nó de IA | 8 h | Total aceita | Sem backup por projeto; vetores são reconstruídos a partir das mensagens em cache |

- Teste de restauração do backup do backend **a cada trimestre**, com relatório.
- Plano de contingência para indisponibilidade (checklist 4.1 do documento 03): a unidade define o meio alternativo e **quais informações podem trafegar nele**, para que a queda do sistema não leve de volta ao WhatsApp.

---

## 9. Indicadores de risco (KRI)

Medidos mensalmente e apresentados ao comitê.

| Indicador | Meta inicial |
|---|---|
| Aparelhos com alerta de integridade abertos > 24 h | 0 |
| Tentativas de cadastro recusadas por atestação | Acompanhar tendência |
| Vulnerabilidades críticas/altas fora do prazo | 0 |
| Riscos com aceite vencido | 0 |
| Usuários ativos sem correspondência no quadro | 0 |
| Tempo médio para revogar aparelho após pedido | ≤ 4 h úteis |
| Restaurações testadas no trimestre | ≥ 1 |
| Incidentes com causa humana (erro de convite, expurgo, exportação) | Acompanhar tendência |
| Pedidos de análise de IA por operação | Acompanhar (detecta desvio de finalidade) |

---

## 10. Entradas para o roteiro

- **Fase 0:** decidir D-R1; iniciar o registro de riscos; definir RTO/RPO; primeira versão do RIPD (documento 06).
- **Fase 2:** teste de bateria e de estabilidade da conexão persistente (se D-R1 = a).
- **Fase 4:** rito de quebra de pseudonimato implementado e testado; dupla custódia decidida.
- **Fase 7:** exercício de mesa de incidente; simulação de vazamento com teste cego; teste de restauração; aceite formal de todos os residuais Médio e Alto.
