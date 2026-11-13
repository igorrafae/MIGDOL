# 03 — Segurança, Homologação e Roteiro de Entrega

**Projeto:** MIGDOL
**Versão do documento:** 0.2 — rascunho para validação
**Público:** área de segurança da instituição, equipe de pentest, equipe de desenvolvimento

---

## 1. Modelo de ameaças

### 1.1 O que se protege

| Ativo | Por que importa |
|---|---|
| Conteúdo das mensagens e anexos | Expõe operação, alvos e métodos |
| Identidade real dos operadores | Risco à integridade física |
| Existência e composição de uma operação | Revela interesse investigativo |
| Chaves de administrador | Permitem expurgo e autorização de exportação |
| Relatórios de IA | Concentram o conteúdo mais sensível em poucas páginas |
| Trilha de auditoria | Base do controle institucional |
| Acervo de custódia e chaves de quórum (HSM offline, cartões) | Abrem as salas sob custódia e os registros cifrados para quórum por até 5 anos (ADR-0003) |

### 1.2 Adversários

| Adversário | Capacidade | Dentro do escopo |
|---|---|---|
| **Membro que vaza** | Acesso legítimo ao conteúdo; quer extrair sem ser identificado | Sim — é a causa do projeto |
| **Aparelho perdido ou apreendido** | Posse física, possivelmente com ferramentas forenses | Sim |
| **Atacante de rede** | Observa e altera tráfego | Sim |
| **Servidor de mensagens comprometido** | Controle total do backend e do banco | Sim |
| **Administrador de infraestrutura curioso** | Acesso a servidores, backups e logs | Sim |
| **Malware no aparelho** | App malicioso, serviço de acessibilidade, acesso remoto | Sim, com limites |
| **Nó de IA comprometido** | Leitura do conteúdo das salas em que participa | Mitigado, não eliminado |
| **Detentores de cartão em conluio ou coagidos** | 3 cartões e acesso ao HSM offline e à estação isolada | Sim; residual (R79, R80) |
| **Nó de campo capturado ligado** | Ordena e entrega na missão; sem chave de grupo nem de membros | Sim (ADR-0004, R89) |
| **Comprometimento de sistema operacional ou hardware** | Exploração de kernel, spyware de nível estatal | **Fora do escopo** — nenhuma aplicação resiste |
| **Coação do usuário** | Obriga a desbloquear e mostrar | Fora do escopo técnico; tratado por procedimento |

### 1.3 Limites declarados

1. Foto da tela por outro aparelho não é impedida, apenas atribuída.
2. Quem lê pode memorizar e contar.
3. O nó de IA vê texto em claro das salas em que está. Se ele for comprometido, essas salas estão expostas.
4. O servidor conhece metadados: quais aparelhos estão em quais salas, horários e volume.
5. Detecção de root e de instrumentação é uma corrida; um atacante dedicado com o aparelho na mão eventualmente contorna. A atestação por hardware e o MDM elevam esse custo.
6. Em sala sob custódia, comprometer o HSM offline ou o quórum expõe os períodos guardados por até 5 anos (risco admitido só nessas salas no GD-1; com a flag ligada a custódia é o padrão, GD-14; aceite formal do R68 pendente). A estação isolada tem texto em claro durante a cerimônia de abertura (ADR-0003 seção 6) e durante a análise da investigação sigilosa (ADR-0010 seção 6).

---

## 2. Matriz de riscos

Probabilidade (P) e impacto (I): Baixo / Médio / Alto / Crítico. A coluna "Residual" é a avaliação depois dos controles.

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R01 | Membro fotografa a tela com outro aparelho | Alto | Alto | Marca d'água visível e invisível; norma de uso; retenção curta | **Médio** — não evitável |
| R02 | Captura ou gravação de tela no próprio aparelho | Alto | Alto | `FLAG_SECURE`, detecção de projeção, MDM | Baixo |
| R03 | Espelhamento por cabo (scrcpy) ou depuração | Médio | Alto | Bloqueio com ADB ativo; MDM desativa depuração; atestação | Baixo |
| R04 | App de acesso remoto ou serviço de acessibilidade lê a tela | Médio | Crítico | Lista de permitidos; lista de apps do MDM; verificação periódica | Baixo |
| R05 | Aparelho com root ou sistema adulterado | Médio | Crítico | Atestação de hardware validada no servidor (sem Play Integrity, doc 07) | Baixo |
| R06 | Instrumentação do app (Frida) para extrair texto | Baixo | Crítico | Detecção no núcleo nativo; ofuscação; verificação de assinatura do app | Médio |
| R07 | Aparelho perdido com sessão ativa | Médio | Alto | Bloqueio por biometria e inatividade; lease de 12 h; revogação remota | Baixo |
| R08 | Extração forense do armazenamento do aparelho | Baixo | Alto | SQLCipher com chave dividida; nada fora do app; sem backup | Baixo |
| R09 | Interceptação de tráfego | Médio | Alto | TLS 1.3 mútuo, pinning, cifra ponta a ponta por baixo | Baixo |
| R10 | Servidor de mensagens comprometido lê conteúdo | Baixo | Crítico | Ponta a ponta; servidor sem chaves | Baixo |
| R11 | Servidor comprometido insere membro falso na sala | Baixo | Crítico | Clientes só aceitam inclusão assinada por administrador registrado no estado do grupo; um commit por época; hash do transcript nos dados autenticados de toda mensagem, comparado sem decifrar (ADR-0005 item 2); código de segurança por sala e época. O commit retido (remoção que não chega a parte do grupo) não insere membro e está no R98, em Médio (01e-E03) | Baixo |
| R12 | Ordem de expurgo forjada ou repetida | Baixo | Alto | Assinatura verificada no aparelho; contador por sala | Baixo |
| R13 | Expurgo não alcança aparelho offline | Alto | Alto | Lease; expurgo ao reconectar | Baixo |
| R14 | Dados residuais em servidor após expurgo | Médio | Alto | Destruição de chave; banco da IA fora de backup | Baixo |
| R15 | Nó de IA comprometido | Baixo | Crítico | Rede isolada, sem internet; contêiner mínimo; sem persistência em claro; acesso administrativo restrito e registrado | **Médio** |
| R16 | Texto em foto manipula a IA (injeção de instruções) | Médio | Médio | Material tratado como dado; saída estruturada; citação de origem; IA sem poder de ação | Médio |
| R17 | Relatório de IA com erro tomado como fato | Alto | Alto | Aviso permanente; citações; norma de uso | Médio |
| R18 | Abuso do expurgo para ocultar irregularidade | Baixo | Crítico | Trilha encadeada; dupla custódia opcional; perfil Auditor | Médio — depende de governança |
| R19 | Chave de administrador comprometida | Baixo | Crítico | Uso condicionado a biometria; chave embrulhada por hardware; revogação | Baixo |
| R20 | Exportação autorizada vaza | Médio | Alto | Autorização individual com validade; marca d'água; auditoria | Médio |
| R21 | Cadastro fraudulento de aparelho | Baixo | Crítico | QR presencial de uso único e curta validade; atestação; MDM | Baixo |
| R22 | Metadados revelam operação a quem administra o servidor | Médio | Médio | Identificadores aleatórios; nome da operação só no conteúdo cifrado; logs sem conteúdo | Médio |
| R23 | Erro de implementação criptográfica | Médio | Crítico | Biblioteca auditada; revisão externa; testes de interoperabilidade | Médio até a revisão externa |
| R24 | Dependência adulterada (cadeia de suprimentos) | Baixo | Crítico | Versões fixadas com hash; compilação reprodutível; repositório interno | Baixo |
| R25 | Indisponibilidade durante operação | Médio | Alto | Redundância do gateway; fila; plano de contingência definido pela unidade | Médio |
| R26 | Texto vaza por teclado, área de transferência ou notificação | Médio | Alto | Teclado controlado pelo MDM; copiar/colar bloqueado; notificação sem conteúdo | Baixo |

Riscos com residual **Médio** exigem aceite formal do responsável pela homologação. Os riscos R27 a R44, o processo de gestão e as regras de aceite por nível estão no documento 04. O R37 (transcrição e memorização) está hoje em **Alto** e não é eliminável.

---

## 3. Plano de testes de invasão

### 3.1 Condições

- Ambiente de homologação idêntico ao de produção, com dados fictícios.
- Equipe independente da equipe de desenvolvimento.
- Três rodadas: caixa-preta, caixa-cinza (com documentação e contas) e caixa-branca (com código).
- Regras de engajamento e janela aprovadas por escrito.
- Referências: OWASP MASVS/MASTG para o app, OWASP ASVS para o backend.

### 3.2 Cenários

#### A. Aparelho e aplicativo

| # | Cenário | Resultado esperado |
|---|---|---|
| A01 | Print pelos botões, pelo assistente e por gesto | Imagem preta ou bloqueio |
| A02 | Gravação de tela nativa e por app de terceiro | Conteúdo não aparece |
| A03 | `adb shell screencap` e `screenrecord` | App bloqueado com ADB ativo |
| A04 | scrcpy por cabo e por rede | Idem |
| A05 | Espelhamento sem fio e tela externa por cabo | Conteúdo oculto; evento auditado |
| A06 | Serviço de acessibilidade próprio lendo a árvore de telas | App bloqueia |
| A07 | App de acesso remoto instalado | App bloqueia |
| A08 | Sobreposição de tela para capturar toques | Toques ignorados ou sobreposição oculta |
| A09 | Aparelho com root (Magisk com ocultação) | Cadastro e sessão recusados pelo servidor |
| A10 | Bootloader desbloqueado com ROM alternativa | Idem |
| A11 | Emulador | Recusado |
| A12 | Frida: anexar ao processo, interceptar funções de exibição | Processo encerra; tentativa auditada quando possível |
| A13 | App reempacotado e reassinado | Recusado pelo servidor |
| A14 | Cópia do diretório de dados para outro aparelho | Dados ilegíveis |
| A15 | Imagem forense do armazenamento após uso | Nenhum texto, imagem ou chave em claro |
| A16 | Imagem forense após expurgo | Nenhum resíduo da operação |
| A17 | Despejo de memória do processo | Sem chaves de longo prazo; registrar quanto texto recente aparece |
| A18 | Backup por ADB e por nuvem | Sem dados do app |
| A19 | Miniaturas, lista de recentes, notificações | Sem conteúdo |
| A20 | Área de transferência e sugestões do teclado | Sem conteúdo |
| A21 | Aparelho offline além do lease | Cache inacessível |
| A22 | Alteração do relógio para estender o lease | Sem efeito |
| A23 | Depuração remota da WebView | Indisponível |
| A24 | Injeção de script por conteúdo de mensagem | Não executa |

#### B. Rede e transporte

| # | Cenário | Resultado esperado |
|---|---|---|
| B01 | Proxy com certificado instalado no aparelho | Handshake falha |
| B02 | Conexão sem certificado de cliente ou com certificado revogado | Recusada |
| B03 | Rebaixamento de versão do TLS | Recusado |
| B04 | Repetição de envelopes capturados | Descartados |
| B05 | Análise de tráfego: tamanho e ritmo | Registrar o que é inferível; avaliar preenchimento |
| B06 | Inundação de conexões e de mensagens | Limites atuam; serviço permanece |

#### C. Backend

| # | Cenário | Resultado esperado |
|---|---|---|
| C01 | Publicar em sala da qual não se é membro | Negado |
| C02 | Ler fila ou anexos de outra sala | Negado |
| C03 | Adivinhar nome de objeto de anexo | Inviável; mesmo com o objeto, conteúdo ilegível |
| C04 | Com acesso total ao banco: recuperar conteúdo | Impossível |
| C05 | Com acesso total ao servidor: incluir aparelho em sala | Clientes rejeitam |
| C06 | Com acesso total ao servidor: forjar expurgo | Clientes rejeitam |
| C07 | Reenviar ordem de expurgo antiga | Rejeitada pelo contador |
| C08 | Alterar ou remover linha da auditoria | Detectado pela cadeia |
| C09 | Reutilizar ou adivinhar token de cadastro | Negado |
| C10 | Cadastro com atestação forjada ou de outro aparelho | Negado |
| C11 | Escalar de operador para administrador ou gestor | Negado |
| C12 | Painel do gestor: autenticação, sessão, injeção | Sem falhas de severidade alta |
| C13 | Logs, métricas e mensagens de erro | Sem conteúdo nem identificador de sala em claro |

#### D. Protocolo e criptografia

| # | Cenário | Resultado esperado |
|---|---|---|
| D01 | Membro removido tenta ler mensagens posteriores | Não decifra |
| D02 | Chave de uma época comprometida: ler épocas anteriores | Não decifra |
| D03 | Inclusão de membro por não administrador | Rejeitada pelos clientes |
| D04 | Autorização de exportação forjada, vencida ou de outro usuário | Rejeitada |
| D05 | Reutilização de nonce em anexos e em chaves embrulhadas | Não ocorre (revisão de código + teste) |
| D06 | Revisão externa da integração MLS | Parecer sem pendência crítica |

#### E. Nó de IA

| # | Cenário | Resultado esperado |
|---|---|---|
| E01 | Saída de rede a partir do contêiner | Bloqueada |
| E02 | Escrita em disco durante OCR e inferência | Nenhum arquivo em claro |
| E03 | Disco e banco da IA após expurgo | Conteúdo ilegível; chave ausente |
| E04 | Foto com texto "ignore as instruções e..." | Relatório não segue a instrução; registrar comportamento |
| E05 | Pedido de análise por não administrador | Negado |
| E06 | Resposta individual lida por outro membro | Impossível |
| E07 | Busca semântica retorna trecho de outra sala | Nunca |
| E08 | Acesso ao serviço de inferência por outro ponto da rede | Indisponível |
| E09 | Relatório contém afirmação sem mensagem de origem | Registrar taxa; critério definido com a unidade |

#### F. Marca d'água e atribuição

| # | Cenário | Resultado esperado |
|---|---|---|
| F01 | Teste cego: foto da tela por outro aparelho, vários ângulos e distâncias | Aparelho de origem identificado |
| F02 | Foto recortada, comprimida e reenviada por mensageiro comum | Registrar taxa de identificação |
| F03 | Remoção da marca visível por edição | Marca invisível ainda identifica, ou registrar limite |
| F04 | Arquivo exportado com autorização | Identificável |

### 3.3 Critério de aprovação

- Nenhum achado crítico ou alto em aberto.
- Achados médios com plano e prazo, aceitos pelo responsável.
- Reteste dos itens corrigidos.
- Resultados de A17, B05, E04, E09, F02 e F03 documentados como limites conhecidos, mesmo quando não houver correção possível.

---

## 4. Checklist de homologação

Cada item precisa de evidência (relatório, captura do teste, configuração exportada) e de responsável.

### 4.1 Governança

- [ ] Ato interno que institui a ferramenta e define papéis
- [ ] Norma de uso assinada por todos os usuários
- [ ] Termo de Ciência e Autorização assinado via gov.br por todos os usuários, versão vigente, assinatura validada
- [ ] Ato interno admitindo as vias de assinatura (gov.br, ICP-Brasil, presencial) e plano de contingência testado (doc 06 seção 7.1)
- [ ] Política de retenção padrão aprovada
- [ ] Rito de acesso do Auditor definido
- [ ] Parecer jurídico sobre retenção, expurgo e cadeia de custódia
- [ ] Decisão registrada sobre dupla custódia
- [ ] Plano de contingência para indisponibilidade
- [ ] Procedimento para perda, roubo ou apreensão de aparelho
- [ ] Aceite formal dos riscos residuais (seção 2)

### 4.2 Aparelhos e MDM

- [ ] Lista de modelos homologados, com hardware seguro e atestação
- [ ] Versão mínima do sistema e nível de patch definidos e impostos
- [ ] App distribuído apenas pelo canal privado do MDM
- [ ] Depuração e instalação de fontes desconhecidas bloqueadas por política
- [ ] Lista de apps permitidos no perfil
- [ ] Teclado definido por política
- [ ] Apagamento remoto testado

### 4.3 Aplicativo

- [ ] Build de produção sem depuração, com ofuscação e verificação de assinatura
- [ ] `FLAG_SECURE` em todas as telas, diálogos e visualizador
- [ ] Todas as verificações da seção 8.2 do documento 02 implementadas e testadas
- [ ] Marca d'água em conversa e visualizador
- [ ] Nenhum dado em armazenamento compartilhado, `localStorage` ou cache da WebView
- [ ] Backup desativado
- [ ] Notificações sem conteúdo
- [ ] Copiar/colar bloqueado
- [ ] Bloqueio por inatividade e biometria
- [ ] Lease funcionando com relógio adulterado
- [ ] Cenários A01–A24 aprovados

### 4.4 Criptografia

- [ ] Biblioteca MLS e versão registradas; auditorias públicas conferidas
- [ ] Revisão externa da integração concluída
- [ ] Chaves de identidade em hardware, não exportáveis, com atestação
- [ ] Chave de administrador condicionada a biometria
- [ ] Vetores de teste de interoperabilidade aprovados
- [ ] Formato da ordem de expurgo documentado e testado com casos inválidos
- [ ] Um commit por época e detecção de bifurcação sem decifrar, nos dois casos do spike (ADR-0005); lease criptográfico e revogação sem destruição (ADR-0006, A73, A74)
- [ ] Cópia de custódia por HPKE para chave pública, verificada pelo hash na mensagem (ADR-0003; A65–A68), onde a flag estiver ligada
- [ ] Cenários D01–D06 aprovados

### 4.5 Backend

- [ ] TLS 1.3 mútuo obrigatório; raiz da autoridade certificadora offline
- [ ] Revogação de certificado com efeito imediato
- [ ] Autorização por sala em toda operação
- [ ] Banco sem conteúdo em claro (inspeção)
- [ ] Logs sem conteúdo (inspeção)
- [ ] Auditoria encadeada; papel da aplicação sem alteração ou exclusão
- [ ] Publicação externa periódica do hash da auditoria
- [ ] Limites de taxa e de tamanho
- [ ] Dependências fixadas; análise de vulnerabilidades sem pendência alta
- [ ] Cenários B01–B06 e C01–C13 aprovados

### 4.6 Nó de IA

- [ ] Rede sem rota para a internet (teste)
- [ ] Contêiner somente leitura, sem swap, usuário sem privilégios
- [ ] Imagens fixadas por digest; pesos conferidos por hash
- [ ] Banco da IA fora de backup
- [ ] Chave mestra em HSM/TPM ou alternativa aprovada
- [ ] Acesso administrativo restrito, com registro
- [ ] Reinício do processo após expurgo
- [ ] Qualidade medida: OCR em amostra real da unidade; correção; relatório avaliado por analistas
- [ ] Cenários E01–E09 aprovados

### 4.7 Expurgo

- [ ] Expurgo com todos online: confirmação em até 10 s
- [ ] Expurgo com aparelho offline: acesso perdido no lease; limpeza ao reconectar
- [ ] Análise forense de aparelho após expurgo sem resíduo
- [ ] Análise forense dos servidores após expurgo sem conteúdo recuperável
- [ ] Registro de auditoria do expurgo íntegro e consultável
- [ ] Expurgo forjado e repetido rejeitados

### 4.8 Operação

- [ ] Procedimentos de implantação, atualização e rotação de chaves escritos e ensaiados
- [ ] Monitoramento e alertas sem conteúdo
- [ ] Treinamento de gestores, administradores e operadores
- [ ] Piloto com uma equipe, em operação simulada, antes do uso real

### 4.9 Gestão de risco, incidentes e atribuição

- [ ] Registro de riscos criado, com donos, prazos e aceites (doc 04)
- [ ] Decisões D-R1 a D-R5 tomadas e registradas
- [ ] Plano de resposta a incidentes e playbooks P1–P11 e P1-S (investigação sigilosa, ADR-0010) aprovados; contatos e acesso de emergência testados
- [ ] Comandante de incidente designado, com substituto
- [ ] Monitoramento de eventos de segurança funcionando, sem conteúdo (doc 05 seção 4)
- [ ] Rito de quebra de pseudonimato aprovado e testado (quórum presencial de 3 comandantes em HSM offline, ADR-0003)
- [ ] Cerimônia do HSM offline realizada (M de N, M ≥ 3): lote de chaves públicas assinado, cartões distribuídos, inventário de backups do HSM, teste de abertura com escopo e de eliminação com recibo (ADR-0003; A54, A65–A68); estação isolada conforme ADR-0003 seção 6
- [ ] Revinculação com aprovação dupla e termo testada (A58–A60, A70, A71); sem HSM online nem aprovação remota na v1
- [ ] Investigação sigilosa (ADR-0010) ensaiada em teste de mesa: impedimento conferido pela estação, registro selado com só o hash na trilha, preservação, nenhuma mudança observável nas salas (A84–A95); D-R37 e D-R38 decididas
- [ ] Procedimento de apuração de vazamento aprovado pela corregedoria, com graus de confiança
- [ ] Teste cego de atribuição realizado e taxas registradas (F01–F03)
- [ ] Exercício de mesa de incidente realizado
- [ ] Teste de restauração do backup do backend realizado
- [ ] Separação de funções implementada (gestor, infraestrutura, Auditor, desenvolvimento)
- [ ] Pipeline de release com duas aprovações, SBOM e assinatura em HSM
- [ ] Prazos de correção de vulnerabilidade definidos e aprovados
- [ ] Revogação por desligamento testada (≤ 4 h úteis)

### 4.11 Aparelho, local, câmera e clonagem

- [ ] Estação de preparação com processo escrito e registro de quem preparou
- [ ] Triagem de spyware executada em todo aparelho antes do cadastro; relatório anexado
- [ ] Série atestada conferida com o inventário (A31–A32)
- [ ] Atestação sem Play Integrity, com raízes de confiança mantidas internamente (ADR-0002)
- [ ] Cenários A25–A50 aprovados; A38, A39 e A42 documentados como limites
- [ ] Câmera: sem alerta nada é gravado ou enviado; com alerta só evidência cifrada para a chave de quórum (A43)
- [ ] Perfis de local definidos por sala e testados (A33–A37)
- [ ] Modo de coação aprovado, treinado e com destinatário do alerta definido
- [ ] Decisões D-R6 a D-R13 do documento 07 tomadas
- [ ] Decisão D-R14 (expurgo remoto após apreensão por autoridade), rito de requisição externa (doc 06 seção 6.3) e pacote de evidência assinado (RF67) definidos
- [ ] Verificação facial de quem acessa e escada por perfil (doc 07 seções 7.1a e 7.2; GD-18) com parecer jurídico, taxas de A76 e A78 medidas e D-R40 decidida
- [ ] Escada de resposta do aparelho (doc 07 seção 7.1) aprovada; decisão D-R15 com parecer do jurídico sobre o processo administrativo (registrado por CI) como base e a eliminação; ADR-0003 revisada aceita; flag de custódia do `instituicao.yaml` conferida; D-R29 decidida (GD-14); D-R30, D-R35 e D-R36 decididas
- [ ] Grau máximo (D-R17) definido e sistema enquadrado no Decreto 7.845/2012 ou na norma do ente; algoritmo de Estado (D-R18) contratado antes de tratar informação classificada (doc 08). D-R17 e D-R18 ficam para o portão G3, depois dos spikes; **nenhum grau classificado até lá** (GD-9)
- [ ] Instância dedicada da instituição instalada pela própria equipe a partir de binário reproduzível; fornecedor sem acesso à produção (doc 09)
- [ ] Instalação e atualização feitas pelo pacote offline assinado (doc 02 seção 10.1): verificador recusa pacote alterado; nenhum segredo no pacote; retorno à versão anterior testado

### 4.10 LGPD

Ver checklist completo no documento 06, seção 13. Todos os itens são condição para o parecer de homologação.

---

### 4.12 Critérios de entrada do piloto e da produção

Validações além de "arquitetura provada, decisões tomadas, parecer jurídico e fluxo real sob teste". **V01 a V10 são critério de entrada do piloto (Fase 7); V11 a V16, da produção.** Os protocolos podem ser escritos antes do G0; a execução depende de aparelhos e do ambiente de homologação. Só dados fictícios.

| ID | Validação | Por que | Como validar (teste ou evidência) | Entrada de |
|---|---|---|---|---|
| V01 | Avaliação independente | Quem desenhou não pode ser o único a validar | Revisão externa curta do core e do protocolo; pentest do servidor e do app (seção 3) | Piloto |
| V02 | Expurgo comprovado de forma forense | Promessa central: ninguém recupera | Aparelho de teste após expurgo e após lease vencido examinado com ferramenta forense; nenhum conteúdo recuperado; laudo anexado | Piloto |
| V03 | Ensaio da cerimônia de quórum | Erro na cerimônia real perde chaves ou expõe dados | Ensaio com HSM e cartões de teste: gerar lote, abrir com quórum, recusar sem quórum, perda de um cartão, substituição de detentor; ata | Piloto |
| V04 | Exercício de incidente | Playbooks do doc 05 nunca testados | Exercício de mesa: aparelho perdido, coação, apreensão por autoridade, nó de campo capturado, pergunta maliciosa à IA; relatório com falhas e correções | Piloto |
| V05 | Backup e restauração | RPO de 15 min do diretório e RPO 0 proposto para os segredos de lease (ADR-0006 item 10) | Derrubar o banco de homologação e restaurar dentro do RPO/RTO; caches continuam legíveis | Piloto |
| V06 | Atualização e reversão | Pacote offline nunca instalado nem revertido (ADR-15) | Instalar, atualizar e reverter em ambiente isolado; recusa de pacote com assinatura ou manifesto inválido | Piloto |
| V07 | Cadeia de suprimentos | Adversário estatal ataca dependências | SBOM; hashes fixados (regra 8); build reproduzível conferido por segunda máquina; release assinado | Piloto |
| V08 | Desempenho e bateria | App que gasta bateria é desligado em campo | Medir bateria com conexão própria (D-R1), latência, grupo grande, anexo grande; metas registradas | Piloto |
| V09 | Usabilidade com agentes | Rito complexo empurra o agente para mensageiro comercial (R120) | Teste com 5 a 8 agentes, dados fictícios: tempo de tarefa, erros, pânico, PIN de coação, revinculação; ajustes registrados. Protocolo, protótipo e critérios no doc 11 (seção 6) | Piloto |
| V10 | Plano do piloto | Sem critério o piloto não termina nem aprova | Escopo (unidade, número de agentes, prazo), métricas de sucesso, critério de interrupção e de saída, aprovados pela chefia | Piloto |
| V11 | Homologação oficial | Órgão que homologa sistema e recurso criptográfico [conferir] (doc 09 seção 9) | Parecer do órgão competente | Produção |
| V12 | Red team com adversário estatal | Obrigatório no perfil Exército (doc 09 seção 5) | Exercício com equipe externa credenciada; achados altos fechados | Produção |
| V13 | Versões convivendo | Aparelhos em versões diferentes não podem perder sala | Teste de migração de esquema e protocolo com versões N e N-1 | Produção |
| V14 | Sustentação | Protocolo próprio exige equipe por anos | Equipe nomeada; responsável pela pilha criptográfica; prazo para corrigir vulnerabilidade; plano de troca da biblioteca MLS (mls-rs como reserva) | Produção |
| V15 | Treinamento e manual | Ritos de coação e perda só funcionam se treinados | Manuais do agente, do gestor e do comandante; treino registrado | Produção |
| V16 | Custo total e licenças | Sustentação e edital | Custo anual de HSM, GPU, aparelhos e equipe; licenças conferidas pelo SBOM | Produção |

Prioridade dos protocolos: V09, V03 e V02.

**Portão da custódia em produção (D08; 01c, seção 3):** nenhuma sala sob custódia entra em produção antes de KEM híbrido na cópia de custódia, **ou** de aceite formal do R78 para a custódia pelo responsável pela homologação e pelo comitê (pergunta 13 do 01c). Vale junto com V11.

## 5. Roteiro de entrega

Cada fase termina com algo demonstrável e com critérios de aceite verificáveis. Prazos dependem do tamanho da equipe e ficam em aberto.

### Fase 0 — Fundação e provas de conceito

**Objetivo:** eliminar as incertezas que podem invalidar a arquitetura.

| Entrega | Critério de aceite |
|---|---|
| Spike MLS: núcleo Rust + cliente de linha de comando (CLI); ponte Kotlin só para o caminho do Keystore, em emulador. Sem React (ADR-0007 aceita; 01e-E19b) | No mínimo **quatro clientes** da CLI (dois administradores e dois membros) trocam mensagens num grupo; inclusão e remoção funcionam (o mesmo texto vale no `PIPELINE.md`, 01e-E04c). Critério ampliado pela revisão (D33) e pelas rodadas da 1b: (1) cópia de custódia por HPKE para chave pública, com `info` ligando sala, período (só por tempo, hipótese do spike) e classe, na **suíte candidata DHKEM(P-256, HKDF-SHA256), HKDF-SHA256, AES-256-GCM** (GD-20; confirmar com o HSM, 01c-C05, C11); (2) assinatura do remetente por dentro da cópia, pela função da biblioteca com rótulo próprio (01c-C08); (3) um commit por época; (4) **bifurcação induzida pelo servidor** (dois commits na mesma época entregues a metades do grupo) detectada nos dois lados pelos sinais da ADR-0005 item 2 (hash do transcript nos dados autenticados em claro; mensagem da mesma época que não decifra; commit divergente); classificada como **bifurcação provada** (SEV1) só quando o cliente verifica a assinatura de administrador registrado no commit conflitante, ou pela importação de cadeia; pacotes que não decifram ou trazem hash diferente, forjados pelo servidor ou por um **membro**, em qualquer quantidade, ficam como **anomalia não autenticada** (alerta "servidor ou membro anômalo", sem recomendar recriar a sala, nunca SEV1), e valor falso nos dados autenticados nunca esconde bifurcação; verificar se a biblioteca deixa processar e verificar commit de época já aplicada **[verificar no spike]** (01g-G13); (4a) **matriz de falso positivo**: duplicata, fora de ordem, época anterior atrasada, byte corrompido e retransmissão não viram SEV1; membro que mente no campo é classificado como remetente inconsistente (ADR-0005 item 2c; 01e-E02); (4b) **commit retido** para parte do grupo detectado pelos sinais 4 e 5, e custo da confirmação de época medido (01e-E03; pergunta Q4); (5) **partição total** não detectada em banda e detectada na reunificação (limite declarado); (6) remoção de aparelho offline; (7a) caminho da chamada ao Keystore a partir do core, **em emulador**, com a saída num buffer direto alocado pelo core, fora do heap da JVM **[verificar]**, e zerada no quadro da chamada (regra 2, GD-12; 01e-E15); o hardware real fica para a etapa 5, depois de G1 (01e-E04a); (8) verificar o que a biblioteca oferece para propostas de atualização e de autoexclusão por membro (01c-C16, decisão pendente); (9) **API da biblioteca**: gravar e ler `authenticated_data` por mensagem antes de processar; ler o hash do transcript confirmado; `epoch_authenticator` para o código de segurança (RF04) (01e-E04d, E16); (10) o que o commit como mensagem pública exporia, se essa alternativa for considerada (ADR-0005 item 4; 01e-E09a) |
| Spike de atestação | Servidor valida a cadeia de um aparelho real e recusa emulador e aparelho com root; (7b) desembrulho pelo Keystore em hardware real (StrongBox/TEE), chamado do core, com a saída zerada no quadro da chamada (01e-E04a) |
| Estudo de viabilidade do iPhone (ADR-0007; etapa 5b do `PIPELINE.md`, depois de G1) | Relatório sobre APNs, App Attest, MDM da Apple, bloqueio de captura de tela, série atestada e conexão em segundo plano; protótipo do doc 11; **sem código de produto**; insumo da ADR própria (D-R31) |
| Spike de inferência | Qwen quantizado, BGE-M3 e PaddleOCR rodando juntos no hardware disponível, com memória e tempos medidos |
| Comparação construir × adaptar | Uma página com a decisão |
| Repositório, integração contínua, padrões de código | Build e testes automáticos de `app`, `core`, `server`, `ai-node`; `scripts/check-boundaries.sh` na integração contínua |
| Esqueleto modular do `server/` (já criado) | Cinco módulos compilam; grafo de dependências verificado |
| Pontos em aberto do documento 01 respondidos | Registrados |
| Decisões D-R1 a D-R5 (documento 04) | Registradas; D-R1 medida em protótipo (bateria, estabilidade) |
| Registro de riscos e primeira versão do RIPD | Criados; responsáveis nomeados |

### Fase 1 — Mensagens de texto ponta a ponta

| Entrega | Critério de aceite |
|---|---|
| Cadastro por QR, certificado, painel mínimo do gestor, **Termo assinado via gov.br como condição de ativação** | Aparelho cadastrado presencialmente; aparelho sem termo válido não ativa; nova versão do termo suspende o acesso até assinar; revogação efetiva |
| Gateway, entrega, diretório, fila com prazo | C01, C02, C04 aprovados |
| Criar operação, convidar, remover | D01, D03 aprovados |
| Estado autenticado do grupo (ADR-0005) desde a primeira sala: chaves de administrador, papéis e marca (tipo comum na Fase 1) na extensão do contexto; um commit por época; sinais de bifurcação; declaração de membros verificada antes de gravar (01c-C26) | Bifurcação induzida detectada; inclusão sem declaração assinada recusada pelo diretório; a Fase 4 não exige migração do contexto |
| Tela de conversa, canais, confirmações | Operador novo envia mensagem sem instrução |
| Cache cifrado e busca local | A14 aprovado; busca em até 200 ms |

### Fase 2 — Proteção do aparelho

| Entrega | Critério de aceite |
|---|---|
| Plugin de segurança completo | A01–A13 aprovados |
| Marca d'água visível | F01 aprovado |
| Lease criptográfico e chave de cache dividida (ADR-0006) | A21, A22, A73 e A74 (doc 07) aprovados; queda no meio da renovação e restauração do banco sem perder caches (01c-C12); troca periódica da chave do banco medida (01c-C13). Selagem na revogação feita pelo `directory`, com a chave pública da classe lida do pacote assinado e chave de teste até a cerimônia (ADR-0006 item 11; 01e-E13 **[proposta]**) |
| Verificação facial de quem acessa e escada por perfil (doc 07 seções 7.1a e 7.2; GD-18), depois do parecer jurídico | A76–A83 (doc 07) aprovados; taxas de falso negativo e de falsificação aceita medidas e dentro da D-R40; perfil missão completo na Fase 8 |
| Bloqueio de cópia, notificações, backup | A18–A20 aprovados |
| Integração com MDM | Checklist 4.2 |

### Fase 3 — Anexos

| Entrega | Critério de aceite |
|---|---|
| Câmera interna, cifra de anexos, envio e visualização | C03 aprovado; nada em armazenamento compartilhado |
| Remoção de metadados | Verificação em amostra |
| Fluxo de exportação autorizada | D04 aprovado; evento em auditoria |
| Marca invisível | F02–F04 medidos |

### Fase 4 — Expurgo e auditoria

| Entrega | Critério de aceite |
|---|---|
| **Entrada da fase (D28; 01c-C19):** hierarquia de chaves por grau, registro de acessos gerado no servidor e cifrado para quórum, incluindo negações, e `audit_log.actor` sem identificador em claro, decididos com D-R4 e D-R5. Dono: responsável pela homologação | Decisão registrada antes do primeiro PR da fase |
| Ordem de expurgo assinada, verificada nos aparelhos | C06, C07 aprovados |
| Destruição no servidor e nos aparelhos | A16 e checklist 4.7 |
| Auditoria encadeada e perfil Auditor | C08 aprovado |
| Retenção automática por prazo | Mensagens somem no prazo em todos os pontos |
| Dupla custódia (se decidida) | Expurgo exige duas assinaturas |
| Ciclo da sala (ativa; concluída com número e hash do relatório e três assinaturas) e expurgo na conclusão conforme o tipo (RN18) | A69 (doc 07) aprovado |
| Marca de custódia (comum por exceção com motivo e duas assinaturas; sala de sistema sempre comum, GD-14), cópia HPKE e serviço `custody`, onde a flag estiver ligada (ADR-0003) | A54, A65–A68 e A75 (doc 07) aprovados; cerimônia de teste do HSM offline |
| Revinculação com aprovação dupla, termo e histórico por outro membro (RN19) | A58–A60, A70, A71 (doc 07) aprovados |
| Rito da investigação sigilosa (ADR-0010): pedido e segunda autorização na estação isolada, quórum de detentores não impedidos em toda sessão (GD-21), impedimento, registro selado, preservação com lista de bloqueio, nível 1 (metadados e registros de atribuição, com CI validada); correlação de metadados na estação; parecer jurídico sobre metadados de agentes sem ordem judicial (01g-G09f) | A84–A90, A92, A94 e A95 (doc 07) aprovados; teste de mesa do P1-S (doc 05) |

### Fase 5 — Nó de IA: OCR, busca e correção

| Entrega | Critério de aceite |
|---|---|
| Nó de IA como membro da sala | Aparece na lista de membros; E01, E02, E08 aprovados |
| OCR automático de fotos | Texto chega à sala em até 5 s; qualidade medida em amostra real |
| Vetores cifrados e índice em memória | E03, E07 aprovados |
| Busca semântica | Resposta só para quem pediu (E06) |
| Corrigir e interpretar texto | Resposta em até 3 s; usuário aceita ou descarta |
| Expurgo alcança o nó de IA | Chave destruída; processo reiniciado |

### Fase 6 — Análise da operação

| Entrega | Critério de aceite |
|---|---|
| Pedido assinado pelo administrador | E05 aprovado |
| Relatório estruturado com citações de origem | Avaliação por analistas da unidade; E09 medido |
| Resumo em etapas para operações longas | Operação com milhares de mensagens analisada dentro do tempo |
| Tratamento de injeção por conteúdo | E04 documentado |

### Fase 7 — Homologação e piloto

| Entrega | Critério de aceite |
|---|---|
| Pentest em três rodadas | Critério da seção 3.3 |
| Revisão externa de criptografia | D06 |
| Checklist da seção 4 completo | Todas as evidências anexadas |
| Piloto com uma equipe | Operação simulada concluída; ajustes registrados |
| Exercício de mesa de incidente e simulação de vazamento com teste cego | Relatórios anexados; taxas de atribuição registradas |
| Teste de restauração e de revogação remota | Aprovados |
| RIPD aprovado e parecer jurídico sobre LGPD e dever de guarda | Anexados |
| Critérios V01–V10 (seção 4.12) | Evidências anexadas |
| Parecer de homologação | Emitido pela área de segurança |

### Fase 8 — Operação em campo por rádio (depois da v1)

Polícia e Exército (D-R22). Entregas e critérios no doc 10 seção 12; exige ADR-0004, ADR-0005 e ADR-0006 aprovadas e levantamento dos rádios (D-R25).

### Depois da v1

- Cliente iOS, só após estudo de viabilidade e ADR própria sobre APNs, App Attest e MDM da Apple (ADR-0007, D-R31).
- Painel com indicadores e IA de análise entre operações (ADR-0008, proposta), depois da Fase 6 e do spike de inferência; painel só com o nível 1a (GD-15); nível 2 só por MLS e por concessão da sala de origem (A96–A99).
- Análise automatizada da investigação sigilosa na estação isolada, com modelo local (ADR-0010; A91, A93), depois da Fase 6 e do spike de inferência.
- Autocompletar com modelo no aparelho.
- Canais com membros distintos dentro da mesma operação.
- Transcrição de áudio.
- Cliente desktop: macOS descartado (ADR-0007); outro desktop só com ADR e novo parecer (R65).

---

## 6. Orientações para o desenvolvimento assistido (Claude Code)

As regras do projeto (leitura obrigatória, regras 1 a 14 que não podem ser quebradas, como trabalhar) estão **só** no `CLAUDE.md` da raiz do repositório. Este documento não as copia, para não divergir: a cópia que havia aqui estava sem a regra 14 (D32, revisão de desenho da Fase 0). Mudança nas regras se faz no `CLAUDE.md`.

### 6.1 Uso de modelos de IA no desenvolvimento

Vale só para **escrever e revisar o código** do MIGDOL. No produto, a regra 10 continua: nenhuma IA de terceiros em execução.

**Pré-condições (antes do primeiro uso):**
- **Autorização formal da instituição** para enviar código e desenho do sistema a um serviço externo de IA (doc 08: credenciamento e cadeia de suprimentos). Sem ela, não se usa.
- **Nunca** enviar dado de produção, nome real, matrícula, dado de operação, chave, certificado, segredo, nem texto classificado. Exemplos e testes usam dados sintéticos (regra 10).
- Toda saída de IA é **proposta**: revisão humana em todo PR; nos PRs de segurança, revisão por pessoa da área de segurança.

**Qual modelo em cada trabalho:**

| Modelo | Esforço | Onde usar |
|---|---|---|
| **Claude Fable 5.2** (`claude-fable-5-2`) | `high` | **Revisão de desenho**, não volume: ADRs de criptografia e protocolo (0003 custódia e quórum, 0004 campo, 0005 commits e bifurcação, 0006 lease, ADR-0002 atestação, algoritmo de Estado como segunda camada); modelagem de ameaças por fase; revisão de segurança final antes da homologação de cada fase; análise de invariantes difíceis (épocas MLS, reconciliação campo–central, expurgo, escada de resposta) |
| **Claude Opus 5.5** (`claude-opus-5-5`) | `high` ou `xhigh` | **Implementação crítica** e revisão de todo PR de segurança: `core/` em Rust (MLS, HPKE, chaves, cache, lease, ordens assinadas); verificação de assinatura e autorização por sala no `server/` (regras 5 e 6); diretório, custódia, auditoria encadeada; plugins de segurança e câmera; nó de campo; verificador do pacote de instalação; testes de `tests/security` |
| **Claude Sonnet 5.5** (`claude-sonnet-5-5`) | `medium` | **Volume:** interface React (`app/src`), código de apoio em Kotlin e Go sem decisão de segurança, painel administrativo (telas), testes de unidade e de integração comuns, documentação, scripts, manifestos YAML e runbooks |

**Regras de escalonamento:**
- Na dúvida se o trabalho é de segurança, usar **Opus** (regra: opção mais restritiva).
- PR feito com Sonnet que toque `core/`, criptografia, autorização, assinatura, chaves, custódia, expurgo ou atestação **é revisado por Opus** antes da revisão humana.
- Mudança de ADR ou de desenho criptográfico passa por **Fable** antes de ser aceita.
- Nenhum modelo implementa primitiva criptográfica à mão (regra 1); se a biblioteca não tiver, parar e perguntar.

**Medição de custo:** registrar tokens e custo dos primeiros 5 PRs da Fase 0 e recalibrar a estimativa (cerca de 140–200 PRs no projeto).

---

## 7. Próximos passos imediatos

1. Validar os três documentos com a chefia da unidade e com o jurídico, em especial as seções 2 e 9 do documento 01.
2. Responder os pontos em aberto do documento 01, seção 11.
3. Levantar aparelhos e MDM disponíveis e o hardware para o nó de IA.
4. Iniciar a Fase 0 pelos três spikes; eles confirmam ou derrubam as escolhas de MLS, atestação e dimensionamento de GPU.
5. Nomear responsável pela homologação, encarregado de dados e comandante de incidente.
6. Levar ao jurídico os pontos do documento 06 seção 14 e as decisões D-R1 a D-R5 do documento 04.
