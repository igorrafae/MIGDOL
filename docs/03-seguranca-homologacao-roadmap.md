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
| **Comprometimento de sistema operacional ou hardware** | Exploração de kernel, spyware de nível estatal | **Fora do escopo** — nenhuma aplicação resiste |
| **Coação do usuário** | Obriga a desbloquear e mostrar | Fora do escopo técnico; tratado por procedimento |

### 1.3 Limites declarados

1. Foto da tela por outro aparelho não é impedida, apenas atribuída.
2. Quem lê pode memorizar e contar.
3. O nó de IA vê texto em claro das salas em que está. Se ele for comprometido, essas salas estão expostas.
4. O servidor conhece metadados: quais aparelhos estão em quais salas, horários e volume.
5. Detecção de root e de instrumentação é uma corrida; um atacante dedicado com o aparelho na mão eventualmente contorna. A atestação por hardware e o MDM elevam esse custo.

---

## 2. Matriz de riscos

Probabilidade (P) e impacto (I): Baixo / Médio / Alto / Crítico. A coluna "Residual" é a avaliação depois dos controles.

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R01 | Membro fotografa a tela com outro aparelho | Alto | Alto | Marca d'água visível e invisível; norma de uso; retenção curta | **Médio** — não evitável |
| R02 | Captura ou gravação de tela no próprio aparelho | Alto | Alto | `FLAG_SECURE`, detecção de projeção, MDM | Baixo |
| R03 | Espelhamento por cabo (scrcpy) ou depuração | Médio | Alto | Bloqueio com ADB ativo; MDM desativa depuração; atestação | Baixo |
| R04 | App de acesso remoto ou serviço de acessibilidade lê a tela | Médio | Crítico | Lista de permitidos; lista de apps do MDM; verificação periódica | Baixo |
| R05 | Aparelho com root ou sistema adulterado | Médio | Crítico | Atestação de hardware validada no servidor; Play Integrity | Baixo |
| R06 | Instrumentação do app (Frida) para extrair texto | Baixo | Crítico | Detecção no núcleo nativo; ofuscação; verificação de assinatura do app | Médio |
| R07 | Aparelho perdido com sessão ativa | Médio | Alto | Bloqueio por biometria e inatividade; lease de 12 h; revogação remota | Baixo |
| R08 | Extração forense do armazenamento do aparelho | Baixo | Alto | SQLCipher com chave dividida; nada fora do app; sem backup | Baixo |
| R09 | Interceptação de tráfego | Médio | Alto | TLS 1.3 mútuo, pinning, cifra ponta a ponta por baixo | Baixo |
| R10 | Servidor de mensagens comprometido lê conteúdo | Baixo | Crítico | Ponta a ponta; servidor sem chaves | Baixo |
| R11 | Servidor comprometido insere membro falso na sala | Baixo | Crítico | Clientes só aceitam inclusão assinada por administrador; código de segurança | Baixo |
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
- [ ] Plano de resposta a incidentes e playbooks P1–P10 aprovados; contatos e acesso de emergência testados
- [ ] Comandante de incidente designado, com substituto
- [ ] Monitoramento de eventos de segurança funcionando, sem conteúdo (doc 05 seção 4)
- [ ] Rito de quebra de pseudonimato aprovado e testado (dupla assinatura)
- [ ] Procedimento de apuração de vazamento aprovado pela corregedoria, com graus de confiança
- [ ] Teste cego de atribuição realizado e taxas registradas (F01–F03)
- [ ] Exercício de mesa de incidente realizado
- [ ] Teste de restauração do backup do backend realizado
- [ ] Separação de funções implementada (gestor, infraestrutura, Auditor, desenvolvimento)
- [ ] Pipeline de release com duas aprovações, SBOM e assinatura em HSM
- [ ] Prazos de correção de vulnerabilidade definidos e aprovados
- [ ] Revogação por desligamento testada (≤ 4 h úteis)

### 4.10 LGPD

Ver checklist completo no documento 06, seção 13. Todos os itens são condição para o parecer de homologação.

---

## 5. Roteiro de entrega

Cada fase termina com algo demonstrável e com critérios de aceite verificáveis. Prazos dependem do tamanho da equipe e ficam em aberto.

### Fase 0 — Fundação e provas de conceito

**Objetivo:** eliminar as incertezas que podem invalidar a arquitetura.

| Entrega | Critério de aceite |
|---|---|
| Spike MLS: núcleo Rust + ponte Kotlin + chamada a partir do React | Dois aparelhos e um cliente de linha de comando trocam mensagens num grupo; inclusão e remoção funcionam |
| Spike de atestação | Servidor valida a cadeia de um aparelho real e recusa emulador e aparelho com root |
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
| Cadastro por QR, certificado, painel mínimo do gestor | Aparelho cadastrado presencialmente; revogação efetiva |
| Gateway, entrega, diretório, fila com prazo | C01, C02, C04 aprovados |
| Criar operação, convidar, remover | D01, D03 aprovados |
| Tela de conversa, canais, confirmações | Operador novo envia mensagem sem instrução |
| Cache cifrado e busca local | A14 aprovado; busca em até 200 ms |

### Fase 2 — Proteção do aparelho

| Entrega | Critério de aceite |
|---|---|
| Plugin de segurança completo | A01–A13 aprovados |
| Marca d'água visível | F01 aprovado |
| Lease e chave de cache dividida | A21, A22 aprovados |
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
| Ordem de expurgo assinada, verificada nos aparelhos | C06, C07 aprovados |
| Destruição no servidor e nos aparelhos | A16 e checklist 4.7 |
| Auditoria encadeada e perfil Auditor | C08 aprovado |
| Retenção automática por prazo | Mensagens somem no prazo em todos os pontos |
| Dupla custódia (se decidida) | Expurgo exige duas assinaturas |

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
| Parecer de homologação | Emitido pela área de segurança |

### Depois da v1

- Cliente iOS, com limitações documentadas.
- Autocompletar com modelo no aparelho.
- Canais com membros distintos dentro da mesma operação.
- Transcrição de áudio.
- Cliente desktop em estação gerenciada.

---

## 6. Orientações para o desenvolvimento assistido (Claude Code)

Para colocar no `CLAUDE.md` da raiz do repositório quando ele for criado:

```markdown
# Regras do projeto

## Leitura obrigatória
- docs/01-negocio-requisitos.md
- docs/02-arquitetura-engenharia.md
- docs/03-seguranca-homologacao-roadmap.md
- docs/04-gestao-de-riscos.md
- docs/05-vazamentos-incidentes-e-atribuicao.md
- docs/06-lgpd-e-conformidade.md

## Regras que não podem ser quebradas
1. Nenhuma primitiva ou protocolo criptográfico é implementado à mão.
   Use as bibliotecas definidas em docs/02. Se faltar algo, pare e pergunte.
2. Chaves privadas e chaves de sessão existem apenas em core/ (Rust).
   Nunca atravessam a ponte para Kotlin ou TypeScript.
3. O servidor (server/) nunca recebe, registra ou grava conteúdo em claro.
   Nenhum log contém corpo de mensagem ou room_id em claro.
4. app/src não usa localStorage, sessionStorage, IndexedDB nem fetch direto.
   Todo acesso a dados e rede passa pelo plugin que fala com o core.
5. Todo endpoint e todo tratador de mensagem verifica autorização por sala.
6. Verificação de assinatura vem antes de qualquer gravação de estado.
7. ai-node/ não grava arquivo em claro e não abre conexão de saída.
8. Dependência nova exige justificativa no PR e versão fixada com hash.
9. server/ é monólito modular: serviços só se falam por api/, sem importar internal/
   de outro serviço, sem consultar tabela alheia e sem FK entre serviços.
   Rode scripts/check-boundaries.sh antes de commitar mudanças no server/.
10. Nenhum dado de produção em desenvolvimento ou teste. Nenhuma telemetria, serviço de
    erro, push ou IA de terceiros. Dado pessoal de usuário só por rito documentado (docs/06).
11. Nenhuma decisão contra pessoa é tomada automaticamente a partir de IA, marca d'água
    ou integridade (docs/05). Mudança que altere o perfil de risco atualiza docs/04.

## Como trabalhar
- Uma fase do roteiro por vez (docs/03, seção 5). Não adiantar fases.
- Cada entrega vem com teste automatizado ligado ao critério de aceite.
- Cenários de docs/03 seção 3 viram testes em tests/security quando automatizáveis.
- Mudança que contrarie um ADR de docs/02 seção 11 ou de docs/adr exige novo ADR antes do código.
- Em dúvida sobre segurança, escolher a opção mais restritiva e registrar a dúvida.
```

---

## 7. Próximos passos imediatos

1. Validar os três documentos com a chefia da unidade e com o jurídico, em especial as seções 2 e 9 do documento 01.
2. Responder os pontos em aberto do documento 01, seção 11.
3. Levantar aparelhos e MDM disponíveis e o hardware para o nó de IA.
4. Iniciar a Fase 0 pelos três spikes; eles confirmam ou derrubam as escolhas de MLS, atestação e dimensionamento de GPU.
5. Nomear responsável pela homologação, encarregado de dados e comandante de incidente.
6. Levar ao jurídico os pontos do documento 06 seção 14 e as decisões D-R1 a D-R5 do documento 04.
