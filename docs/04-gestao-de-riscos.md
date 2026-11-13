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
| **D-R14** | Expurgo remoto após apreensão por autoridade: o que o sistema faz (padrão proposto: suspender e revogar, sem destruir). Proposta da etapa 1b: erros de desbloqueio e pânico sem apagar até a decisão; lease vencido não apaga (ADR-0006); apagamento remoto com duas autorizações (D23) | Jurídico e comitê; ver R63, R66 e anexo de validação seção 3.7 |
| **D-R6 a D-R13** | Câmera, local, coação, triagem de spyware, atestação, pós-quântico, check-in. **D-R8 decidida pela GD-18** (verificação facial de quem acessa; falta o parecer jurídico) | Ver documento 07 seção 13; riscos R45–R56, R66–R71, R79–R83 (R81–R83 não se aplicam na v1), R96–R97 e R116–R117 no documento 07 seção 10 |
| **D-R15** | **Definido pelo responsável (GD-1 a GD-6; ADR-0003 revisada); falta o jurídico.** Custódia só em sala sob custódia, por flag da instituição; cópia HPKE para chave pública; abertura por quórum presencial de no mínimo 3 comandantes em HSM offline, com CI (sempre) e ordem judicial ou processo administrativo com validação jurídica; guarda até 5 anos da conclusão; revinculação sem custódia; com a flag ligada, custódia é o padrão (GD-14) | Chefia + jurídico + segurança; doc 07 seções 7.3 e 7.4 |
| **D-R29** | ~~Exceção à custódia com a flag ligada~~ **Decidido (GD-14, opção B):** custódia é o padrão; sala comum só com motivo registrado e segunda assinatura; salas de sistema sempre comuns | Ver documento 07 seção 13; R111 |
| **D-R30** | Foto do termo de revinculação (regras 12 e 13) e prazo do termo | Ver documento 07 seção 13 |
| **D-R35** | Prazo padrão de guarda do segredo de lease selado na revogação ("prazo da apuração", GD-13) | Responsável pelo produto + jurídico; ADR-0006 item 8; R110 |
| **D-R36** | Quem dá a segunda assinatura da sala comum com a flag ligada; motivo em lista fechada ou livre (GD-14; proposta: categoria fechada na trilha, texto livre só no grupo, sem hash, 01e-E09b) | Responsável pelo produto + chefia; R111 |
| **D-R37** | Investigação sigilosa (ADR-0010): quem dá a segunda autorização e o que conta como "independente"; quando a suspeita alcança os detentores, quem é a autoridade **de fora dessa classe** (GD-22; proposta: chefe do jurídico ou autoridade do escalão superior); substitutos nomeados de antemão no `instituicao.yaml`, inclusive do Auditor, do Gestor, do responsável pela homologação e do jurídico; detentores de reserva com cartão emitido de antemão quando os impedidos deixam menos de M (01g-G04b) | Responsável pelo produto + corregedoria + jurídico; R112, R115 |
| **D-R38** | Investigação sigilosa (ADR-0010): prazo da investigação (proposta: 90 dias; **no máximo duas renovações**, cada uma com quórum, GD-22) e duração máxima da sessão; prazo do registro selado (nunca antes da revisão final, GD-22, nem do prazo de responsabilização por abuso **[conferir]**); revisão de controle por quórum e base dela; acesso do investigado no contraditório; forma do **aviso no arquivamento**, que é o padrão (GD-22), e de seu adiamento; órgão de controle que recebe a comunicação do encerramento **[conferir]** | Jurídico + corregedoria; R112 |
| **D-R39** | Gatilhos da destruição imediata do nó de campo capturado (GD-18) e o que não aciona | Segurança + comando; doc 10; R118 |
| **D-R40** | Escada por perfil (GD-18; **em parte decidida pela GD-24**: duas capturas para "outra pessoa", vivacidade fora do gatilho até medida à parte, número de PINs errados e PIN de coação no nível 4 só depois de medir). Falta: valores medidos; nível 4 por tentativas no perfil missão; perfil missão na polícia; limites de erro do reconhecimento facial (inclui a parte facial da antiga U-3 do doc 11) | Responsável pelo produto + segurança + jurídico; doc 07; R67, R116 |
| **D-R41** | ADR-0008 nível 2 (GD-15): base mínima da concessão por perfil (CI, autorização do comandante ou as duas); concessão retroage ou não (proposta: não); agregados derivados também exigem concessão (proposta: sim); aprovação dupla para criar a sala de análise | Responsável pelo produto + chefia; R103–R107 |
| **D-R42** | Ferramenta do protótipo navegável e local dele no repositório (doc 11 seção 5.2; antes U-1) | Responsável pelo produto, com segurança; regra 8 |
| **D-R43** | Metas de tempo, erro e SUS definitivas do V09, formativas com amostra pequena (doc 11 seções 6.2, 6.5 e 6.8; antes U-2) | Responsável pelo produto, após a rodada 1; R120 |
| **D-R44** | Termo de participação no teste de usabilidade, responsável pelo teste e rito do rosto real no teste facial (doc 11 seção 6.6; antes U-4; 01g-G18) | Jurídico e encarregado; R128 |
| **D-R45** | Política de notificação por perfil: genérica ou nenhuma (doc 11 PI-2; antes U-6) | Segurança, com a chefia; R125 |
| **D-R46** | Texto aprovado das telas de trava, revogação e marca de custódia (doc 11; antes U-7) | Responsável pelo produto e jurídico; R122, R124 |
| **D-R47** | Margem aceitável de diferença de tempo entre PIN normal e PIN de coação (doc 11 PI-4; parte da antiga U-3; a parte facial está na D-R40) | Segurança, com o jurídico; R121 |
| **D-R28** | Orquestração da instalação: Compose ou systemd (proposta) ou Kubernetes | Segurança + infraestrutura da instituição; doc 02 seção 10.1 |
| **D-R25 a D-R27** | Rádios e redes de campo, teto do lease de campo, cifra do rádio como segunda camada | Ver documento 10 seção 9; riscos R89–R93, R100 e R118 no documento 10 seção 8 |
| **D-R21 a D-R24** | Instância dedicada por instituição, operação em campo, rito militar do termo, detentores do quórum | Ver documento 09 seção 7; riscos R84–R88 no documento 09 seção 6 |
| **D-R17 a D-R20** | Grau máximo, algoritmo de Estado, aparelho e sistema, proteção de metadados. **D-R17 e D-R18 ficam para depois dos spikes, no portão G3; nenhum grau classificado até lá (GD-9)**. A D-R20 cobre também quem vê o mapa da missão e se a posição vai a um grupo MLS de comando (doc 11, antiga U-5; 01g-G19) | Ver documento 08 seção 7; riscos R72–R78 no documento 08 seção 6 |
| **D-R16** | Persistência de matrícula depois do reset e rastreio, frente à regra 10 | Segurança + chefia; doc 07 seção 7.4 |
| **D-R31** | iPhone: estudo de viabilidade na Fase 0 e ADR sobre APNs, App Attest e MDM da Apple antes de produção (ADR-0007) | Responsável pelo produto + segurança |
| **D-R32** | Painel e IA de análise: limite mínimo de contagem do nível 1 (proposta: 5), quem cria a sala de análise, prazo do escopo (ADR-0008) | Chefia + segurança |
| **D-R33** | Papéis por instituição e combinações proibidas no `instituicao.yaml` (ADR-0009) | Chefia + jurídico |
| **D-R34** | Base legal e prazo para dados de terceiros (presos, investigados) nos indicadores nominais (LGPD art. 4º, III) [conferir] | Jurídico |

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
| R40 | Exposição de dados de cadastro (matrícula ↔ codinome ↔ aparelho) | Baixo | Alto | Matrícula cifrada para a chave de quórum de pseudonimato; índice de consulta com chave em HSM, não exportável, consulta registrada e limitada (D09; 01c-C15; doc 02 seção 4.2); quebra de pseudonimato só pelo rito do documento 05. **Limite:** matrícula é enumerável, e o quórum protege o vínculo em repouso e em massa, não o que o gestor sabe por função ao cadastrar | Baixo |
| R41 | Erro humano: expurgo acidental, convite para pessoa errada, exportação indevida | Médio | Alto | Confirmação reforçada (RF51); dupla custódia opcional; revisão de convite antes de enviar; trilha | Médio |
| R42 | Uso desviado da análise de IA (investigar pessoas ou assuntos fora da finalidade) | Baixo | Alto | Pedido assinado e registrado com finalidade; Auditor revê pedidos; relatório com marca d'água | Médio |
| R43 | Relógio do servidor ou da rede adulterado (afeta lease, validade e ordens de expurgo) | Baixo | Alto | Fonte de tempo interna redundante; tolerância configurada; contador de sequência nas ordens (não depende só de hora) | Baixo |
| R44 | Vazamento por captura de cópia do material em trânsito entre sistemas oficiais (exportação autorizada) | Médio | Alto | Marca d'água; autorização individual; auditoria; canal oficial de entrega | Médio |
| R60 | **Expurgo configura eliminação ilícita de documento público** (Lei 8.159, arts. 9º e 25; Código Penal, art. 314) | Médio | Alto | Parecer, tabela de temporalidade aprovada ou ordem de preservação **antes** da Fase 4; exportação do que for documento oficial; registro do fundamento de cada expurgo | **Alto até o parecer**; bloqueia a Fase 4 |
| R61 | Prova extraída do sistema anulada por falta de base normativa, ciência ou rito | Médio | Alto | Norma de uso e termo prevendo a coleta; contraditório; cadeia de custódia; indício, não prova (doc 05) | Médio |
| R62 | **Ordem judicial ou requisição externa que o sistema não pode cumprir** (conteúdo inexistente no servidor por desenho) → multa, pedido de bloqueio, tensão institucional | Médio | Médio | Rito e responsável institucional para requisições; resposta padrão do que existe (cifrado, metadados); jurídico no circuito; playbook P11 (doc 05) | Médio |
| R63 | **Expurgo remoto após apreensão por autoridade** lido como destruição ou supressão de prova | Baixo | Alto | Em apreensão por autoridade: suspender e revogar, sem destruir, até decisão do jurídico (D-R14); revogar nunca destrói o segredo de lease, qualquer que seja o motivo declarado (GD-13); apagamento remoto e destruição do segredo só com duas autorizações (D23); registro do evento | **Alto até D-R14** |
| R64 | **Print ou cópia como prova de vazamento** contestada (adulterável, sem cadeia) | Alto | Médio | Pacote de evidência assinado (RF67); proibir print como prova; cadeia de custódia (doc 05) | Médio |
| R65 | **Segundo cliente ou espelhamento** criado no futuro (web, desktop, aparelho companheiro) reabre a controvérsia jurídica e o vetor de ataque | Baixo | Alto | Requisito de que não existe vínculo de segundo cliente; mudança exige ADR e novo parecer | Baixo |
| R94 | Pacote de instalação adulterado no caminho até a instituição | Baixo | Crítico | Assinatura em HSM com duas aprovações; verificador antes de instalar; build reproduzível; mídia controlada | Baixo |
| R95 | Segredo vazado em arquivo de configuração | Médio | Alto | Nenhum segredo no YAML ou no pacote; varredura de segredos no pipeline; chaves no HSM | Baixo |
| R57 | Conta gov.br de agente tomada por terceiro, que assina o termo em nome dele | Baixo | Alto | Assinatura na presença do gestor, com documento conferido; nível mínimo de conta definido; ativação exige também aparelho íntegro | Baixo |
| R58 | Termo ou fluxo de assinatura revela vínculo agente ↔ sistema, ou detalhes das defesas | Médio | Alto | Termo sem codinome, sem identificador de aparelho e sem detalhe técnico; PDF em guarda segura do gestor; nada operacional enviado ao gov.br | Médio |
| R59 | Indisponibilidade do gov.br impede cadastro ou renovação de termo | Médio | Médio | Plano de contingência em três vias (gov.br, ICP-Brasil, presencial com testemunha e acesso provisório), com gatilho, limites e ratificação (doc 06 seção 7.1); versão do termo muda raramente; carência na renovação | Baixo |
| R98 | **Bifurcação de grupo** por servidor hostil ou corrida de commits; diretório divergente do estado MLS autoriza quem não é membro (D16) | Médio | Crítico | Um commit por época; hash do transcript nos dados autenticados em claro, comparado sem decifrar, e alerta com mensagem da mesma época que diverge ou não decifra (ADR-0005 item 2; 01c-C02); chaves de administrador no estado do grupo; declaração de membros assinada verificada antes de gravar; teste no spike MLS com os dois casos; sinais classificados (SEV1 só com bifurcação provada, isto é, commit conflitante de administrador verificado ou importação de cadeia; pacotes que não decifram, em qualquer quantidade, são anomalia não autenticada, sem recomendar recriar a sala; 01g-G13) e confirmação de época **[proposta]** (01e-E02, E03). **Limite:** partição total e commit retido só se detectam na reunificação, pelos sinais 4 e 5 ou pela confirmação de época | **Médio** até os testes da Fase 1 (01e-E03) |
| R99 | `segredo_servidor` capturado uma vez abre o cache para sempre, junto com o hardware (D21) | Médio | Alto | Lease criptográfico: reembrulho a cada renovação com geração, servidor guarda só o vigente, expiração pelo relógio do servidor (ADR-0006). **Limites declarados (01c-C13):** chave do banco fixa (quem a copiou da memória, ou tem embrulho antigo com o segredo daquela geração e o hardware, abre o banco futuro; rekey medido na Fase 2); offline, o prazo continua sendo o app esquecer; segredos antigos em backup do diretório | Baixo |
| R110 | Segredo de lease guardado na revogação (GD-13) mais posse do aparelho revogado e desbloqueio do agente abrem o cache dele; acervo de segredos selados é alvo | Baixo | Alto | Selado para a chave de quórum da classe, no período de vencimento, com todas as gerações guardadas (01e-E05, E06); abertura só por quórum presencial com CI e base registrada (regra 14); sozinho não abre nada; prazo da apuração com padrão e piso (D-R35); destruição e revogação em lote com duas autorizações; hash do selado na trilha, conferido pelo Auditor (01e-E08); nunca volta ao `server/` em claro (ADR-0006 item 8). Aparelho achado volta por revinculação e `history_share`, não por cerimônia | Baixo |
| R111 | Sala comum criada com a flag de custódia ligada, ou sala de sistema usada para conversa, para fugir da custódia e da apuração | Baixo | Alto | Motivo registrado e segunda assinatura de pessoa distinta (D-R36), conferida contra a lista assinada de papéis (01e-E10a); marca visível aos membros e irreversível; categoria do motivo na trilha, sem hash do texto livre (01e-E09b); sala de sistema só pelo fluxo próprio do sistema; Auditor revê salas comuns por amostragem (GD-14; ADR-0003 seção 1) | Baixo |

Todos os riscos com residual **Médio** ou **Alto** exigem aceite conforme a seção 3.2. O **R37** hoje fica em **Alto** e precisa de decisão do comitê: o projeto não tem como eliminá-lo, só reduzir o alcance e aumentar a chance de atribuição.

### 5.1 Mudanças de perfil de risco da revisão de desenho (Fase 0, etapa 1b)

Regra 11: mudança que altera o perfil de risco atualiza este documento. Correções em `docs/fase0/01b-correcoes.md`.

| Risco | Onde | Mudança | Motivo |
|---|---|---|---|
| R68 | doc 07 | Residual "A avaliar" → **Médio**, proposto para aceite (o GD-1 limita o risco à sala sob custódia; aceite formal pelo responsável pela homologação pendente); controles reescritos | GD-1; D01–D07; 01c-C18 |
| R69 | doc 07 | Controles: aprovação dupla, termo com foto, só operação ativa, histórico por outro membro | GD-4; D03, D05 |
| R81, R82, R83 | doc 07 | **Não se aplicam na v1** (sem HSM online, sem aprovação remota) | GD-4 |
| R52 | doc 07 | "Cifra para o Auditor" → chave de quórum de evidência | D32 |
| R66 | doc 07 | Na primeira rodada: P Médio → Baixo; residual Alto → Médio. **Desfeito na segunda rodada:** volta a P Médio e residual **Alto até D-R14**, porque rebaixar sem o comitê contraria a seção 3.2 | D23; 01c, seção 6 |
| R63 | este doc | Controle novo: duas autorizações para apagar. Residual mantido | D23 |
| R40 | este doc | Matrícula cifrada para quórum; índice com chave fora do `server/` | D09 |
| R11 | doc 03 | Controles de ADR-0005 | D16 |
| R96 | doc 07 | **Novo:** termo de revinculação com foto exposto | GD-4, GD-5 |
| R97 | doc 07 | **Novo:** cópia de custódia inválida gravada pelo remetente (residual Baixo, proposto para aceite) | D06; 01c-C18 |
| R98 | este doc | **Novo:** bifurcação de grupo e pertença divergente | D16 |
| R99 | este doc | **Novo:** segredo de lease reutilizável | D21 |
| R89, R90, R92 | doc 10 | Controles do nó que só ordena, central congelada, segredos de missão e teto monotônico | D13, D14, D22 |
| R100 | doc 10 | **Novo:** aparelho do administrador da missão capturado desbloqueado inclui membro | D13, GD-11 |

### 5.2 Riscos de plataformas, painel, IA de análise e papéis (ADR-0007 a ADR-0009)

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R101 | iPhone expõe à Apple quando o aparelho recebe algo (APNs) e depende de App Attest e MDM da Apple | Alto | Alto | iOS fora da produção até o estudo de viabilidade e ADR própria (D-R31); carga vazia; Android na v1 | Médio (enquanto fora da v1) |
| R102 | Clientes nativos divergem no comportamento de segurança (Android × iOS), inclusive na interface: mesma tela com comportamento distinto confunde ritos (doc 11, antes R-U15) | Médio | Alto | Núcleo único em Rust; regras de segurança no core; mesma bateria de testes `tests/security` nas duas plataformas; mesma lista de telas T01 a T30, mesmo texto e testes de equivalência de interface (doc 11) | Baixo |
| R103 | Nó de IA com escopo entre operações vira o alvo mais valioso (mapa das operações num ponto só) | Médio | Crítico | Escopo mínimo com prazo; isolamento sem saída de rede; sem rastro em disco; GPU dedicada; acesso físico por duas pessoas (ADR-0008) | Médio |
| R104 | IA vaza dado de um compartimento para outro, ou copia trecho de conversa na resposta (busca sem filtro, ou injeção de prompt por mensagem de sala, inclusive dentro do escopo) | Médio | Crítico | Filtro de permissão antes da busca, no código; índice por sala; só salas com concessão (GD-15); IA sem ferramentas que agem; saída estruturada validada no código, com verificação de sobreposição com a fonte (01e-E11); teste de vazamento entre escopos e de injeção dentro do escopo (A99) | **Médio** até o teste de injeção dentro do escopo (01e-E11) |
| R105 | Agregado pequeno no painel permite deduzir uma operação ou pessoa | Médio | Alto | Contagens arredondadas ou em faixas, sem filtro que reduza o grupo abaixo do limite (D-R32; 01e-E14); lista fechada de indicadores no `instituicao.yaml`; datas por semana ou mês; sem sala nem nome; nível 1b fora do painel (GD-15); contra a pressão por mais dado no painel (doc 11, antes R-U16): texto "o que este painel não mostra", tarefa U20 do V09 e mudança só por ADR | Baixo |
| R106 | Vínculo errado da IA entre pessoas vira suspeita indevida; analista lê a referência da IA como fato (doc 11, antes R-U10) | Médio | Alto | Fonte citada; rótulo fixo "indício para análise humana"; tarefa U15 do V09 e treino; saída estruturada; regra 11; revisão humana registrada; limite de perguntas e revisão amostral do que foi perguntado | Médio |
| R107 | Indicadores nominais tratam dados de terceiros sem base ou prazo definidos | Médio | Alto | D-R34; expurgo junto com a sala; acesso só no escopo | Médio até D-R34 |
| R108 | Pessoa acumula papéis conflitantes (ex.: administrador de sistema e detentor de quórum) | Baixo | Crítico | Combinações proibidas recusadas pelo sistema; atribuição por duas pessoas (ADR-0009) | Baixo |
| R109 | Papel vencido ou de pessoa transferida continua ativo | Médio | Alto | Papel com prazo; negação por padrão; revisão periódica; KRI de papéis vencidos | Baixo |

Riscos que continuam abertos por decisão adiada: R72 e R78 (sem grau classificado até G3; KEM híbrido na custódia, D08, com portão de produção no doc 03 seção 4.12); R73 (prazos do registro de acessos, D28, entrada da Fase 4).

### 5.3 Mudanças de perfil de risco da segunda rodada da etapa 1b

Correções em `docs/fase0/01d-correcoes-rodada2.md`. Nenhuma mudança abaixo é aceite de risco: quem desenvolve não aceita (seção 2).

| Risco | Onde | Mudança | Motivo |
|---|---|---|---|
| R66 | doc 07 | Volta a **Alto até D-R14** (P Médio); reavaliação proposta para Médio depois de A72–A74 aprovados e do comitê | Recomendação do 01c, seção 6 |
| R63 | este doc | Controle novo: revogar nunca destrói o segredo de lease | GD-13; 01c-C01 |
| R68 | doc 07 | "Aceito" → proposto para aceite; nota de que, com a flag ligada, a custódia é o padrão | 01c-C18; GD-14 |
| R97 | doc 07; ADR-0003 | "Aceito" → proposto para aceite do dono do ativo | 01c-C18 |
| R98 | este doc; doc 03 R11 | Detecção sem decifrar (hash do transcript nos dados autenticados); limite da partição total declarado | 01c-C02 |
| R99 | este doc | Limites declarados; renovação com geração | 01c-C12, C13 |
| R40 | este doc | Chave do índice em HSM; limite do D09 registrado | 01c-C15 |
| R110 | este doc | **Novo:** segredo de lease guardado na revogação | GD-13 |
| R111 | este doc | **Novo:** sala comum ou de sistema usada para fugir da custódia | GD-14 |
| R101–R109 | seção 5.2 | ADR-0008 dividida em nível 1a e 1b; nenhum residual muda | Integração das ADRs 0007–0009 |

### 5.4 Riscos e mudanças de perfil da terceira rodada da etapa 1b (GD-15 a GD-20; ADR-0010)

Correções em `docs/fase0/01f-correcoes-rodada3.md`. Nenhuma mudança abaixo é aceite de risco; as GD são decisão do responsável pelo produto, não aceite formal (seção 3.2).

**Riscos novos da investigação sigilosa (ADR-0010):**

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R112 | Investigação sigilosa usada para perseguir ou vigiar alguém sem suspeita legítima (desvio de finalidade) | Baixo | Crítico | CI validada em procedimento instaurado e segunda autorização independente, de fora da classe dos detentores quando a suspeita a alcança (GD-22; D-R37); impedimento, inclusive do iniciador na posição; **quórum presencial de detentores não impedidos em toda sessão** (GD-21); ordem judicial para conteúdo; registro selado com hash na trilha, guardado até a revisão final (GD-22); prazo com no máximo duas renovações, cada uma com quórum; revisor que não é o iniciador nem a segunda autorização; encerramento comunicado a órgão de controle fora da cadeia **[conferir]**; aviso ao investigado no arquivamento como padrão; revisão de controle por quórum (D-R38); estatística ao comitê; contraditório | Médio |
| R113 | Existência ou escopo da investigação vaza ao investigado (CI em tramitação, cerimônia percebida, período não eliminado, entrada órfã na trilha, log e contadores do HSM, comportamento de quem sabe) | Médio | Alto | Nada muda nas salas, no nó de IA e no cliente (A87); trilha só com hash de tamanho fixo e entradas de cobertura nas cerimônias de rotina **[proposta]** (01g-G06a); retrato por ato de liberação do backup casado com o teste de restauração, sem consulta ao servidor vivo (01g-G07a); CI sigilosa **[conferir]**; declaração de cada participante; sessões junto das cerimônias de rotina, com operadores e testemunhas não impedidos e gravação selada (01g-G06b, c); lista de bloqueio conferida sem abrir o registro, com só a contagem de períodos preservados (01g-G01c); painel sem preservação selada (01g-G06d). **Limites declarados:** sem cobertura, as entradas órfãs se contam; com N pequeno, a ausência de um detentor numa cerimônia é percebida; com quórum em toda sessão (GD-21), mais gente sabe | Médio |
| R114 | Indício falso da análise (correlação, modelo, mensagem plantada pelo verdadeiro autor, registro do servidor forjado ou apagado por administrador investigado) leva a medida contra inocente | Médio | Alto | Saída estruturada sem conteúdo; grau de confiança do doc 05; cada sinal rotulado **atestado** ou **registro do servidor sem assinatura**, que não eleva o grau sozinho; comparação de retratos de datas diferentes (01g-G08); duas revisões humanas; indício de conteúdo rotulado e sem elevar o grau sozinho; reprodutibilidade por hash do modelo, parâmetros e semente (01g-G09c); decisão do MP e do juiz; contraditório (regra 11; R36, R106) | Médio |
| R115 | Impedimento incompleto (suspeito não declarado aprova ou participa) ou quórum inviável por impedimentos | Médio | Alto | Três fontes de impedimento conferidas pela estação por chave, com o identificador opaco do usuário na lista assinada de papéis e falha fechada para quem não tem vínculo (01g-G05); iniciador na posição impedido nos atos seguintes (01g-G03); declaração de cada participante; N ≥ M + impedidos; detentores de reserva com cartão emitido de antemão (D-R37; 01g-G04b). **Limites declarados:** sem M livres não há sessão, e o caso sai do sistema (GD-21); com detentor suspeito não identificado, o sistema não garante sigilo perante os detentores (01g-G04a) | Médio |
| R119 | Pressão para gravar sala comum às escondidas (membro oculto, troca silenciosa da marca, expurgo suspenso nos aparelhos, uso do nó de IA) | Baixo | Crítico | Proibido no desenho (GD-17, confirmada pela GD-23; ADR-0010 seção 7): MLS sem membro oculto e com aviso de inclusão; marca irreversível e visível; nó de IA fora; preservação de sala comum só de metadados, nunca do cifrado MLS (GD-23); testes A87, A88 e A95. **Limite declarado (01g-G26a):** o investigado que administra uma sala comum pode expurgá-la durante a investigação, e nada o impede | Baixo |

Os riscos novos do reconhecimento facial (R116, R117) estão no doc 07 seção 10, e o da destruição do nó de campo (R118), no doc 10 seção 8.

**Mudanças de perfil:**

| Risco | Onde | Mudança | Motivo |
|---|---|---|---|
| R98 | este doc | Residual Baixo → **Médio** até os testes da Fase 1; controles: sinais classificados, sinais 4 e 5, confirmação de época **[proposta]** | 01e-E02, E03 |
| R104 | seção 5.2 | Residual Baixo → **Médio** até o teste de injeção dentro do escopo; controles: concessão, saída estruturada | 01e-E11; GD-15 |
| R105, R106 | seção 5.2 | Controles: contagens em faixas, lista fechada, nível 1b fora; saída estruturada e limite de perguntas. Residual mantido | 01e-E14, E11; GD-15 |
| R103, R107 | seção 5.2 | Escopo só com concessão da sala de origem, regra de tipo e expurgo em cascata reduzem o alcance; residual mantido até a revisão do Fable | GD-15; 01e-E01 |
| R110 | este doc | Controles: todas as gerações seladas, período de vencimento, prazo com piso, revogação em lote com duas autorizações, hash na trilha | 01e-E05, E06, E08 |
| R111 | este doc | Controles: segunda assinatura contra a lista assinada de papéis; sem hash do motivo na trilha | 01e-E09b, E10a |
| R66 | doc 07 | Controles: destruição por perfil (missão de campo) aceita pelo responsável; no uso urbano nada muda. Residual mantido (**Alto até D-R14**) | GD-18 |
| R67 | doc 07 | Reescrito como falso negativo, com a destruição possível no perfil missão; residual mantido (Médio) | GD-18 |
| R116, R117 | doc 07 | **Novos:** falso positivo ou falsificação do reconhecimento facial; gabarito facial exposto | GD-18 |
| R89 | doc 10 | Controle: destruição imediata do nó capturado | GD-18 |
| R118 | doc 10 | **Novo:** perda do não sincronizado na destruição do nó (exceção escrita à GD-13) | GD-18; 01e-E07 |
| R112–R115, R119 | este doc | **Novos:** investigação sigilosa | GD-16, GD-17; ADR-0010 |

### 5.5 Riscos e mudanças de perfil da quarta rodada (GD-21 a GD-25; doc 11)

Correções em `docs/fase0/01h-correcoes-rodada4.md`. Riscos de usabilidade do doc 11 seção 7, numerados e fundidos conforme o `01g` seção 6. P e I são propostas a validar no V09; nenhuma mudança abaixo é aceite de risco (seção 3.2).

**Riscos novos de usabilidade (doc 11):**

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R120 | Rito complexo leva ao uso de mensageiro comercial; o conteúdo sai do sistema, fora da regra 3 e da custódia (antes R-U1) | Alto | Alto | V09 (doc 11 seção 6) com linha de base num mensageiro comum e medida de uso real no piloto (01g-G17); mínimo de passos que o rito permite (PI-9); tarefas U02 e U03; revisão do rito se 20% ou mais disser que usaria outro aplicativo; norma de uso | **Alto** até a V09 aprovada |
| R121 | Tela de coação distinguível: o adversário percebe a isca e reage com violência (antes R-U6) | Médio | Crítico | PI-4; teste cego com observador em escolha forçada, uma rodada por perfil; tempo medido por instrumento dentro da margem (D-R47); apagamento do nível 4 sem mudar tempo nem aquecimento (A79); revisão por segurança | **Alto** até o teste cego e a medição de tempo da Fase 2 |
| R122 | Marca de custódia mal lida: o agente acha que a sala é comum e fala algo sensível, ou o contrário; falha de transparência (antes R-U7) | Médio | Médio | PI-3; tarefa U04; texto padronizado (D-R46) | Baixo |
| R123 | Três assinaturas da conclusão viram "clique e siga": conclusão indevida e expurgo irreversível da sala comum (antes R-U9) | Médio | Alto | Resumo na tela (sala, relatório, quem já assinou); confirmação com biometria; número e hash do relatório conferidos; tarefa U12 | Médio |
| R124 | Gestor lê "revogado" como "destruído": promessa falsa à instituição, ou expectativa de recuperação errada (antes R-U11) | Médio | Médio | Texto padronizado (D27; D-R46); tarefa U18 | Baixo |
| R125 | Notificação genérica excessiva ou ausente: o agente deixa de ver ordem urgente, ou desliga o app por barulho (antes R-U12) | Médio | Médio | Política por perfil (D-R45); tarefa de perceber mensagem urgente (doc 11 seção 6.8); teste em campo | Médio |
| R126 | Em campo, tela clara ou som revela a posição em operação noturna (antes R-U13) | Médio | Alto | PI-6 (modo de campo, sem som por padrão); cenário S1; teste de campo na Fase 8 | Médio |
| R127 | Comandante vê a posição e a repassa por foto da tela ou por fala; a posição sai do MLS (antes R-U14) | Médio | Alto | Sem captura de tela (PI-5); norma de uso; treino. **Limite declarado (01g-G19):** a restrição de quem vê o mapa é de interface; todo aparelho membro decifra a posição; grupo MLS de comando é decisão da Fase 8 (D-R20) | Médio |
| R128 | Sessão de teste de usabilidade vaza dado real, ou grava biometria real de participante (antes R-U17; 01g-G18) | Baixo | Alto | Só dados fictícios e arquivo fictício revisado (doc 11 seções 5.3 e 6.6); sem gravação de rosto, voz ou tela com identificação; teste facial sem reconhecimento real no papel e no navegável; no nativo, só com rito (consentimento específico, gabarito zerado e conferido ao fim, parecer do encarregado; D-R44) | Baixo |

**Fusão no que já existia (01g seção 6):** R-U2 e R-U3 no R55; R-U4 no R67; R-U5 dividido entre R67 (destruição do aparelho do dono) e R116 (adversário que passa); R-U8 no R69 (controle novo); R-U10 no R106; R-U15 no R102; R-U16 no R105.

**Mudanças de perfil:**

| Risco | Onde | Mudança | Motivo |
|---|---|---|---|
| R55 | doc 07 | Reescrito com R-U2 e R-U3; impacto **Alto no perfil missão**; nível 4 pelo PIN de coação só depois de medir o acionamento acidental. Residual mantido (Médio) | 01g-G12a; GD-24 |
| R67 | doc 07 | Inclui R-U4, parte do R-U5 e o colega que destrói o aparelho de outro; controles da GD-24 (duas capturas, vivacidade fora, PINs só depois de medir) e caminho de volta da isca. Residual mantido | 01g-G11; GD-24 |
| R116 | doc 07 | Parte do R-U5; "detecta, não barra" declarado. Residual mantido | 01g-G11b |
| R117 | doc 07 | Gabarito decifrado no core, em Rust, sob chave de dados embrulhada pelo hardware; regra 12 (GD-25). Residual mantido | 01g-G22a; GD-25 |
| R69 | doc 07 | Controle novo: resumo legível com hash e U14 (R-U8) | 01g seção 6 |
| R98 | este doc | SEV1 só com bifurcação provada; anomalia não autenticada sem recomendar recriar a sala. Residual mantido (Médio) | 01g-G13 |
| R102, R105, R106 | seção 5.2 | Controles de interface do doc 11 (R-U15, R-U16, R-U10). Residuais mantidos | 01g seção 6 |
| R112 | seção 5.4 | Quórum em toda sessão; CI validada; segunda autorização de fora dos detentores quando a suspeita os alcança; no máximo duas renovações; revisor distinto do iniciador e da segunda autorização; aviso no arquivamento; registro até a revisão final. Residual mantido (Médio) | GD-21, GD-22; 01g-G01, G02 |
| R113 | seção 5.4 | Entradas de cobertura, ato de liberação do backup, lista de bloqueio, sessões junto da rotina; limites declarados (órfãs, N pequeno, mais gente sabe). Residual mantido | 01g-G06, G07; GD-21 |
| R114 | seção 5.4 | Sinal atestado × registro do servidor sem assinatura; reprodutibilidade. Residual mantido | 01g-G08, G09c |
| R115 | seção 5.4 | Identificador opaco e falha fechada; iniciador impedido por posição; reservas emitidas de antemão; sem M livres não há sessão; limite do detentor suspeito não identificado. Residual mantido | 01g-G03, G04, G05; GD-21 |
| R119 | seção 5.4 | GD-17 confirmada (GD-23); limite do expurgo de sala comum pelo investigado. Residual mantido | 01g-G26a; GD-23 |
| R120–R128 | esta seção | **Novos:** usabilidade (doc 11) | GD-19; 01g seção 6 |

---

## 6. Controles organizacionais exigidos

### 6.1 Classificação e compartimentação
- Toda operação tem nível de sigilo definido pelo administrador na criação entre os graus legais (reservado, secreto; ultrassecreto fora do escopo, D-R17), conforme a Lei 12.527/2011 e o doc 08. **Até o portão G3 (D-R17, D-R18), nenhuma sala trata grau classificado** (GD-9).
- **Necessidade de saber:** operador só entra em canais de que precisa. Quanto menos pessoas veem uma mensagem, menor o conjunto de suspeitos num vazamento. Canais com membros distintos dentro da mesma operação já estão previstos para a v2; avaliar antecipar.
- Relatórios de IA herdam o nível mais alto das salas analisadas.
- **Compartimento de análise** (ADR-0008, proposta): sala de análise com escopo, prazo e revisão; só entra quem tem credencial em todo o escopo (doc 08 seção 5.1), e só entram no escopo salas que concederam (GD-15).

### 6.2 Separação de funções
| Função | Não pode acumular com |
|---|---|
| Gestor do sistema (cadastra) | Administrador de operação; Auditor |
| Administrador de infraestrutura | Gestor do sistema; Auditor |
| Auditor | Qualquer papel operacional |
| Quem desenvolve | Quem aprova release; quem opera produção |
| Detentor de cartão de quórum | Proposta: gestor do sistema, administrador de infraestrutura e analista de inteligência (ADR-0009). Decisão em D-R24 |
| Analista de inteligência (ADR-0008, ADR-0009, propostas) | Administrador de infraestrutura; detentor de cartão de quórum; gestor do sistema, que conhece o vínculo matrícula ↔ codinome (01e-E18c) |
| Aprovador da revinculação | O agente revinculado e o gestor que fez o cadastro do aparelho novo |
| Signatário da conclusão (RN18) | Os outros dois signatários da mesma conclusão (três pessoas distintas). Quando o solicitante é o comandante ou o subcomandante, o terceiro signatário é o superior imediato do solicitante, definido na criação |
| Segunda assinatura da sala comum com a flag ligada (GD-14) | O criador da sala e os membros dela (D-R36) |
| Iniciador e segunda autorização da investigação sigilosa (ADR-0010) | Um ao outro; qualquer pessoa impedida (lista declarada ou posição no escopo); subordinado direto do iniciador, na segunda autorização; detentor de cartão, na segunda autorização, quando a suspeita alcança a classe dos detentores (GD-22; D-R37). O iniciador na posição pode iniciar, mas não opera a estação, não revê e não recebe o relatório (01g-G03) |
| Revisor do relatório de rastro (ADR-0010) | O outro revisor; impedidos; o iniciador e a segunda autorização (GD-22) |
| Detentor de cartão de quórum, operador e testemunha nas sessões da investigação sigilosa | Impedidos; a conferência é feita pela estação, por chave e identificador opaco (01g-G05); toda sessão tem quórum (GD-21) |

Com a ADR-0009 (proposta), esta tabela vira as **combinações proibidas** do `instituicao.yaml`, recusadas pelo sistema na atribuição de papel (D-R33), conferidas contra a lista assinada de papéis e, para detentores de cartão, na cerimônia de emissão (01e-E10). A correspondência entre os nomes desta tabela, os atores do doc 01 seção 4 e os papéis da ADR-0009 está na própria ADR.

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
| Chaves de quórum (custódia por sala e período; evidência, registros, com série própria para a investigação sigilosa, termo, localização e segredo de lease guardado por classe e período; pseudonimato pela vida do vínculo) | HSM offline; M de N cartões (M ≥ 3) | Lotes gerados em cerimônia com o mesmo M de N; período por tempo (proposta: mês) em todas as classes, como hipótese do spike (01c-C06, C07; 01e-E21e); segredo de lease no período de vencimento (01e-E06) | Sem cópia fora do HSM e dos backups dele; eliminação por período destrói no HSM e nos backups, com recibo; decapsulação dentro do HSM exigida para as classes sem chave por sala (ADR-0003) |
| Chave de cerimônia de custódia | HSM offline | Longa; troca em cerimônia com M de N, com a antiga válida só para conferir lotes emitidos (ADR-0003 seção 2) | Assina as estruturas das chaves públicas dos lotes (classe, lote, validade, identificador, suíte); a pública vai fixada no pacote assinado |
| Chaves dos aprovadores da revinculação e dos papéis da sala | Aparelho, por biometria | Ao trocar de aparelho | Registradas em cerimônia ou no estado do grupo; conferidas no core (doc 02 seção 4.1) |

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
| Auditoria | 4 h | 0 | Não pode perder evento; fila local no emissor. **Exceção aceita pelo responsável (GD-18):** trilha de missão ainda não sincronizada quando o nó de campo capturado é destruído (R118) |
| Segredos de lease (diretório e `custody`) | 1 h | 0 **[proposta]** | Perder o segredo vigente deixa o cache ilegível; o MLS não redecifra o passado (ADR-0006 item 10; 01c-C12) |
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
- **Fase 2:** teste de bateria e de estabilidade da conexão persistente (se D-R1 = a); verificação facial e escada por perfil (GD-18), depois do parecer jurídico, com R67, R116 e R117 medidos (A76–A83).
- **Fase 4:** entrada com D28 decidido (doc 03 seção 5); rito de quebra de pseudonimato implementado e testado; dupla custódia decidida; ciclo da sala, custódia (onde a flag estiver ligada; padrão com a flag ligada, GD-14) e revinculação v1; rito da investigação sigilosa (ADR-0010), com D-R37 e D-R38 decididas; aceite formal do R68 e do R97.
- **Fase 7:** exercício de mesa de incidente; simulação de vazamento com teste cego; teste de restauração; aceite formal de todos os residuais Médio e Alto.
