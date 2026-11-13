# 01 — Negócio e Requisitos

**Projeto:** MIGDOL
**Versão do documento:** 0.2 — rascunho para validação
**Público:** patrocinador do projeto, chefia da unidade de inteligência, jurídico/corregedoria, equipe de desenvolvimento

> Documentos relacionados: `02-arquitetura-engenharia.md` (como construir), `03-seguranca-homologacao-roadmap.md` (como provar que está seguro e em que ordem entregar), `04-gestao-de-riscos.md` (processo de gestão de risco), `05-vazamentos-incidentes-e-atribuicao.md` (detecção, resposta e investigação de vazamentos) `06-lgpd-e-conformidade.md` (proteção de dados e conformidade) e `07-resistencia-a-clonagem-e-endurecimento-avancado.md` (clonagem, aparelho e local autorizados, câmera, coação), `08-informacao-classificada-e-nivel-inteligencia.md` (graus de sigilo, Decreto 7.845/2012, requisitos de nível de inteligência), `09-perfil-exercito-e-modelo-de-implantacao.md` (inteligência militar e instância dedicada por instituição), `10-operacao-em-campo-por-radio.md` (nó de campo sobre rádio tático) e `anexos/validacao-juridica-preliminar.md` (verificação de lei e jurisprudência, não substitui parecer).

**Sobre o nome.** MIGDOL (hebraico: "torre") vem de Provérbios 18:10, "Torre Forte". O nome remete a firmeza e proteção, adequado a um ambiente policial e operacional.

---

## 1. Problema e objetivo

**Problema.** A equipe de inteligência usa o WhatsApp para coordenar operações. Isso tem gerado vazamento de dados e de capturas de tela, porque o aplicativo:

- vincula a identidade ao número de telefone (SIM swap, exposição do número pessoal);
- guarda histórico no aparelho e em backups de nuvem fora do controle da instituição;
- permite print, encaminhamento e download livre de anexos;
- não oferece ao responsável pela operação nenhum controle sobre o ciclo de vida das conversas.

**Objetivo.** Um aplicativo de mensagens institucional, cifrado ponta a ponta, sem histórico permanente, em que o administrador da operação controla retenção, exportação e destruição dos dados, e em que cada vazamento seja dificultado e, quando ocorrer, atribuível a um usuário.

**Exceção: sala sob custódia** (ADR-0003, GD-1). Onde a instituição liga a custódia no `instituicao.yaml` (Exército: ativa; polícia: desligada por padrão), a sala guarda uma cópia institucional cifrada para a chave de quórum, por até 5 anos a contar da conclusão. Nas salas comuns, vale o objetivo acima sem exceção.

**O que o projeto não promete.** Nenhum software impede que alguém fotografe a tela com outro aparelho ou copie à mão o que leu. O projeto reduz a superfície de vazamento, desencoraja e permite atribuir a origem. Essa limitação precisa estar clara para o patrocinador desde o início.

---

## 2. Decisões já tomadas

| # | Decisão | Motivo |
|---|---|---|
| D1 | App nativo, **não PWA**. A ADR-0007, **aceita** (GD-19), troca o Capacitor por telas nativas (Compose; SwiftUI depois do estudo do iOS) a partir da Fase 1 | Bloqueio de captura, detecção de depuração e chaves em hardware só existem em app instalado |
| D2 | **Android gerenciado** (MDM) na v1; iOS em fase posterior | No iOS só é possível detectar captura, não bloquear |
| D3 | Identidade **sem número de telefone**: chave gerada no hardware do aparelho, cadastro presencial por QR, vínculo com matrícula | Elimina SIM swap e exposição do número pessoal |
| D4 | Cifra ponta a ponta com **MLS (RFC 9420)** e AES-256-GCM | Grupos/canais grandes com remoção eficiente de membros; biblioteca auditada em vez de protocolo próprio |
| D5 | IA roda em **nó on-premises declarado como membro da sala** | É a única forma honesta de ter OCR, busca semântica e análise sem quebrar a cifra em segredo |
| D6 | Correção de texto **sob demanda**; sem autocompletar a cada tecla na v1 | Autocompletar no servidor exporia cada tecla digitada |
| D7 | Expurgo por **destruição de chaves** (crypto-shredding), não só por exclusão de registros | Exclusão em banco deixa resíduos em disco, WAL e backups |
| D8 | **Trilha de auditoria selada de metadados** (sem conteúdo) | Sem ela, corregedoria e jurídico tendem a barrar a homologação |
| D9 | PostgreSQL como único banco; sem OpenSearch | Menos componentes dentro do perímetro |
| D10 | Backend como monólito modular, com serviços já separados em módulos e prontos para virar microsserviços | Baixo custo operacional agora, sem reescrita depois (docs/adr/0001) |

---

## 3. Brainstorm

### 3.1 Ideias adotadas

- Estrutura tipo Discord: **Operação → Canais → Mensagens**, com tela de conversa simples no estilo WhatsApp/Telegram.
- Convite por **QR de uso único**, com validade de poucos minutos, lido presencialmente.
- **Marca d'água forense** em toda a tela de conversa e em todo anexo exibido (codinome + identificador do aparelho + hora).
- **Foto com OCR automático**: ao enviar uma foto, o texto extraído acompanha a mensagem e passa a ser pesquisável.
- **Busca textual rápida no aparelho**, sobre o que está no cache local cifrado.
- **Botão "Analisar operação"** exclusivo do administrador: resumo, linha do tempo, riscos.
- **Anexos apenas para visualização**; exportar exige autorização assinada do administrador.
- **Kill-switch**: o administrador apaga a operação e tudo ligado a ela é destruído. Em sala sob custódia, a cópia de custódia segue o prazo de guarda, contado do expurgo quando não houve conclusão (RN04; ADR-0003 seção 7). Se o kill-switch de um administrador continua ao lado da conclusão com três assinaturas: pergunta 5 do 01c.
- **Chaves com prazo de validade**: aparelho que fica offline além do prazo perde acesso sozinho.

### 3.2 Ideias para avaliar depois

- Autocompletar de próxima palavra com modelo pequeno **rodando no próprio aparelho**.
- Expurgo com **dupla custódia** (duas assinaturas) para operações sensíveis.
- Mensagem de voz com transcrição local.
- Modo "sala de leitura": conteúdo só aparece enquanto o dedo está na tela.
- Cliente desktop em estação gerenciada.

### 3.3 Ideias descartadas

| Ideia | Por que saiu |
|---|---|
| Identificação por número de telefone | SIM swap, interceptação de SMS, expõe o número pessoal |
| AES-128 | Sem ganho prático frente ao AES-256; padrão institucional pede 256 |
| PWA puro | Não bloqueia captura nem acessa hardware seguro |
| OpenSearch junto com PostgreSQL | Duplica dados sensíveis em dois motores |
| Embeddings em claro com índice permanente | Permitem reconstruir boa parte do conteúdo original |
| Configurar o banco para "não alertar monitoramento" | Expurgo em órgão público precisa ser auditável, não invisível |

### 3.4 Construir ou adaptar

Antes de iniciar a Fase 1, vale registrar uma comparação com adaptar uma base existente (Matrix/Element, Wire ou similar). A parte cara e arriscada do projeto é a criptografia e sua auditoria, não a tela de chat. A recomendação deste documento é construir o produto próprio **usando biblioteca MLS auditada**, porque os requisitos de expurgo, marca d'água, controle de exportação e nó de IA não existem prontos. Essa escolha deve ser confirmada no spike da Fase 0.

---

## 4. Atores e papéis

| Papel | Quem é | Pode |
|---|---|---|
| **Gestor do sistema** | Responsável institucional pela plataforma | Cadastrar e revogar usuários e aparelhos; criar administradores de operação. Não lê conteúdo |
| **Administrador de operação** | Chefe da operação: o **comandante da missão**; substituto, o **subcomandante** (GD-11; se vale para toda operação ou só para a missão de campo: pergunta 9 do 01c) | Criar operação e canais; convidar e remover membros; definir retenção; autorizar exportação; pedir análise de IA; acionar expurgo; assinar a conclusão |
| **Operador** | Agente da equipe | Enviar e ler mensagens e anexos; buscar; pedir correção de texto |
| **Auditor** | Corregedoria / controle interno | Consultar a trilha de metadados; a corregedoria pode **iniciar** investigação sigilosa (ADR-0010). Não lê conteúdo |
| **Nó de IA** | Serviço on-premises | Membro técnico da sala; processa OCR, embeddings, correção e análise. Sempre visível na lista de membros |
| **Equipe de inteligência** | Analistas da unidade de inteligência | Marcar a sala como ativa (RN18) |
| **Solicitante** | Quem pediu a abertura da operação | Assinar a conclusão (RN18). Se o solicitante for o comandante ou o subcomandante, assina no lugar dele o **superior imediato** do solicitante, definido na criação |
| **Analista de inteligência** | Analista com credencial em todo o escopo de uma sala de análise (ADR-0008 e ADR-0009, propostas; depois da v1) | Ver indicadores do nível 2 e perguntar à IA, como membro da sala de análise, só sobre salas que concederam entrada no escopo (GD-15). Não lê conversas das salas do escopo |
| **Chefe da inteligência e superior da área** | Chefias | Aprovar a revinculação (RN19). Não leem conteúdo |
| **Comandantes detentores de parte** | N comandantes com cartão (D-R24) | Abrir dado cifrado da instituição só em quórum presencial de no mínimo 3, em HSM offline (RN20); iniciar investigação sigilosa (ADR-0010), da qual ficam fora se impedidos. Não leem nada no dia a dia |

---

## 5. Regras de negócio

| ID | Regra |
|---|---|
| RN01 | Toda mensagem e todo anexo são cifrados no aparelho do remetente. O servidor de mensagens nunca tem acesso a texto em claro. |
| RN02 | O servidor guarda apenas mensagens cifradas **ainda não entregues**, por prazo máximo configurável (padrão: 72 h). Entregue ou vencido, o registro é apagado. **Exceção:** em sala sob custódia, o serviço `custody` guarda a cópia de custódia (cifrada para a chave de quórum, ilegível para o servidor) até o prazo de guarda; e os registros cifrados para quórum (evidência, visualização, termo) seguem o prazo próprio (ADR-0003). |
| RN03 | Cada operação tem uma política de retenção definida pelo administrador na criação (ex.: 24 h, 7 dias, até o encerramento). Vencido o prazo, o conteúdo some do aparelho e do nó de IA. Na **conclusão** (RN18), tudo sai dos aparelhos e do nó de IA. |
| RN04 | Quando o administrador apaga a operação, ou ela é concluída, são destruídos: chaves da operação, mensagens em cache nos aparelhos, anexos, textos de OCR, embeddings e relatórios de IA. **Sala comum:** permanece somente o registro de auditoria de que a operação existiu e foi expurgada. **Sala sob custódia:** permanece também a cópia de custódia até o fim do prazo de guarda; o painel e os recibos dizem isso, e nunca "destruído" para o que continua guardado (D27). |
| RN05 | Um usuário só entra no sistema por cadastro presencial feito pelo gestor, com QR de uso único e vínculo à matrícula. |
| RN16 | O aparelho só é ativado depois que o agente assina, via gov.br e na presença do gestor, o **Termo de Ciência e Autorização de Acesso** (docs/anexos/termo-de-ciencia-e-autorizacao.md). Cada nova versão do termo exige nova assinatura; sem ela, o acesso é suspenso. |
| RN06 | Um usuário só entra em uma operação por convite do administrador daquela operação. |
| RN07 | Anexos são exibidos dentro do app. Salvar fora do app exige autorização do administrador, individual por arquivo e por usuário, com validade. A cópia exportada leva marca d'água e a exportação é registrada. |
| RN08 | O app não abre conteúdo se o aparelho estiver em estado inseguro (ver RNF-D). O evento é registrado e o administrador é avisado. |
| RN09 | Relatórios de análise de IA são visíveis apenas ao administrador da operação. Todos os membros sabem que o nó de IA participa da sala. A ADR-0008 (proposta) acrescenta a **sala de análise**: indícios entre operações, com referência e sem trecho de conversa, para analistas com credencial no escopo, só de salas cujo administrador concedeu a entrada no escopo, com CI e/ou autorização do comandante, de forma visível aos membros (GD-15); se aceita, esta regra e a transparência (doc 06 seção 7) mudam. |
| RN10 | Aparelho que não renova sua autorização dentro do prazo (padrão: 12 h) perde acesso ao cache local, mesmo sem receber ordem de expurgo. |
| RN11 | A trilha de auditoria registra quem, quando e o quê em nível de evento (cadastro, convite, remoção, exportação, alerta de integridade, expurgo). Nunca registra conteúdo de mensagem. |
| RN12 | Remoção de um membro troca as chaves da sala; o removido não lê nada do que for enviado depois. |
| RN13 | Copiar e colar texto para fora do app é bloqueado. Notificações do sistema não exibem conteúdo. |
| RN14 | Nenhuma decisão contra pessoa é tomada automaticamente a partir de marca d'água, alerta de integridade ou IA. Vazamento é apurado por pessoas, com contraditório (doc 05). |
| RN15 | O usuário é informado, antes do primeiro uso, de tudo que o sistema observa e de como a identidade real pode ser ligada ao codinome (doc 06 seção 7). |
| RN17 | **Tipo da sala** (ADR-0003, GD-1, GD-2, GD-14): com a flag de custódia da instituição desligada, toda sala nasce **comum**. Com a flag ligada, a sala nasce **sob custódia** por padrão; nasce comum só com **motivo registrado** e **segunda assinatura** de pessoa distinta do criador (D-R36). **Salas de sistema** (sala de resposta, doc 07 seção 7.4) são sempre comuns. A marca é feita na criação, assinada, visível aos membros e irreversível. Mudar a flag é cerimônia registrada e não retroage. |
| RN18 | **Ciclo da sala** (GD-3): criada → **ativa** (marcada pela equipe de inteligência) → **concluída** → expurgada. A conclusão exige o número e o hash do relatório interno assinado eletronicamente no sistema corporativo (o relatório não entra no MIGDOL) e as assinaturas, no app, de três pessoas distintas: **comandante da missão**, **subcomandante** e **solicitante**. Quando o solicitante é o comandante ou o subcomandante, a terceira assinatura é do **superior imediato do solicitante**, definido na criação. Os papéis ficam no estado autenticado do grupo desde a criação; quem autoriza trocar o titular de um papel durante a operação está pendente (pergunta 3 do 01c). Signatário que não é membro assina sobre o identificador opaco da sala, o número e o hash do relatório, sem ver conteúdo. Sala comum: expurgo total na conclusão. Sala sob custódia: expurgo da cópia de trabalho na conclusão; a cópia de custódia fica até 5 anos a contar da conclusão (ou do expurgo, se não houve conclusão) e depois é eliminada com recibo. |
| RN19 | **Revinculação** (GD-4): só em operação ativa; aprovação dupla (chefe da inteligência e superior da área); termo de ciência e aceite do proprietário, com foto e assinatura no app; histórico enviado pelo aparelho de outro membro pelo MLS. Não usa custódia nem decifra dado da instituição (doc 07 seção 7.4). |
| RN20 | **Abertura institucional** de dado cifrado (custódia, evidência, registros, termo, localização, matrícula, segredo de lease guardado): só por quórum presencial de no mínimo 3 comandantes, com cartões, em HSM offline, com **CI** (registra e tramita o pedido em todos os casos) e base registrada: ordem judicial (cumprida), ou processo administrativo com validação jurídica **[conferir]**. A CI não substitui ordem judicial (GD-6; 01c-C23). |

---

## 6. Requisitos funcionais

### 6.1 Identidade e acesso

| ID | Requisito | Prioridade |
|---|---|---|
| RF01 | Cadastro de aparelho por QR de uso único, com geração de chaves no hardware seguro e atestação do aparelho | Essencial |
| RF02 | Desbloqueio do app por biometria ou PIN a cada abertura e após inatividade | Essencial |
| RF03 | Exibição de codinome; matrícula e nome real nunca trafegam nas salas | Essencial |
| RF04 | Verificação entre membros por código de segurança **por sala e por época** (muda a cada commit), derivado do `epoch_authenticator` do MLS **[verificar no spike]**, por comparação presencial ou por QR; também revela partição total na reunificação (ADR-0005 item 2e; 01e-E16) | Importante |
| RF05 | Revogação imediata de aparelho pelo gestor, sem destruir: o segredo de lease fica selado para a chave de quórum (RF91) | Essencial |
| RF06 | Termo de Ciência e Autorização assinado eletronicamente (gov.br) no cadastro; assinatura validada, versão e resumo guardados; revalidação a cada nova versão. Contingência em três vias, com acesso provisório limitado e ratificação (doc 06 seção 7.1) | Essencial |

### 6.2 Operações, canais e mensagens

| ID | Requisito | Prioridade |
|---|---|---|
| RF10 | Criar operação com política de retenção e lista inicial de membros | Essencial |
| RF11 | Criar canais dentro da operação (ex.: geral, campo, análise) | Essencial |
| RF12 | Enviar e receber texto em tempo real, com confirmação de entrega e leitura | Essencial |
| RF13 | Enviar foto tirada no app, sem passar pela galeria do aparelho | Essencial |
| RF14 | Enviar documento e áudio | Importante |
| RF15 | Responder, mencionar e reagir a mensagens | Desejável |
| RF16 | Apagar mensagem para todos (remetente ou administrador). Em sala sob custódia, a cópia de custódia continua (ADR-0003 seção 3) | Importante |
| RF17 | Mensagem com prazo próprio de autodestruição, menor que o da operação. Em sala sob custódia, a cópia de custódia continua | Desejável |

### 6.3 Anexos e controle de exportação

| ID | Requisito | Prioridade |
|---|---|---|
| RF20 | Visualizador interno de imagens, PDF e áudio, sem gravar arquivo em área compartilhada | Essencial |
| RF21 | Pedido de exportação pelo operador e aprovação/negação pelo administrador | Essencial |
| RF22 | Marca d'água visível em tela e em anexos; marca invisível em imagens exportadas | Essencial |
| RF23 | Remoção de metadados (GPS, modelo do aparelho, hora do sensor) antes do envio | Essencial |

### 6.4 Busca e OCR

| ID | Requisito | Prioridade |
|---|---|---|
| RF30 | Busca textual instantânea no cache local do aparelho | Essencial |
| RF31 | OCR automático de fotos enviadas; texto extraído anexado à mensagem para todos os membros | Essencial |
| RF32 | Busca semântica por significado dentro da operação, respondida pelo nó de IA | Importante |

### 6.5 Assistência de IA

| ID | Requisito | Prioridade |
|---|---|---|
| RF40 | Botão "Corrigir texto" no campo de digitação; mostra sugestão e o usuário aceita ou descarta | Essencial |
| RF41 | "Interpretar/explicar" uma mensagem ou trecho, a pedido do usuário | Importante |
| RF42 | "Analisar operação" (somente administrador): resumo executivo, linha do tempo, entidades citadas, riscos e recomendações | Essencial |
| RF43 | Relatório de análise exibido apenas no app, sujeito à mesma retenção e expurgo | Essencial |
| RF44 | Aviso permanente de que o relatório é apoio à decisão e pode conter erros | Essencial |

### 6.6 Expurgo e auditoria

| ID | Requisito | Prioridade |
|---|---|---|
| RF50 | Expurgo da operação por ordem assinada pelo administrador, verificada por cada aparelho | Essencial |
| RF51 | Confirmação reforçada antes do expurgo (biometria + digitar o nome da operação) | Essencial |
| RF52 | Painel do administrador com situação do expurgo por aparelho (confirmado / pendente / expirará por prazo) e, em sala sob custódia, a data até a qual a cópia de custódia fica guardada | Importante |
| RF53 | Consulta à trilha de auditoria pelo perfil Auditor | Essencial |
| RF54 | Expurgo com dupla custódia, configurável por operação | Desejável |
| RF55 | Registro de visualização cifrado para a chave de quórum (ADR-0003), para reduzir o conjunto de suspeitos em vazamentos (adotado; doc 05 seção 3.3) | Importante |
| RF56 | Quebra de pseudonimato (codinome → pessoa) somente por quórum presencial de no mínimo 3 comandantes em HSM offline (ADR-0003), com registro (doc 05 seção 3.4); matrícula cifrada para a chave de quórum no diretório | Essencial |
| RF57 | Alertas de segurança sem conteúdo para monitoramento (doc 05 seção 4) | Essencial |
| RF58 | Ordem de preservação que suspende o expurgo de uma operação, se o jurídico decidir que é necessária (doc 06 seção 8.3) | A decidir |
| RF67 | Pacote de evidência assinado para qualquer material que sair do sistema (hash, assinatura, cadeia de custódia); print não é meio de prova | Proposto |
| RF68 | Rito de requisição externa (doc 06 seção 6.3) e requisito de que não existe vínculo de segundo cliente nem espelhamento oficial | Proposto |
| RF82 | Marca de custódia no estado autenticado do grupo: assinada na criação, conferida contra o `instituicao.yaml` assinado (e o histórico assinado das versões), visível aos membros, irreversível; com a flag ligada, marca "comum" só com motivo e duas assinaturas, e sala de sistema sempre comum; cópia de custódia HPKE em toda mensagem de sala sob custódia, sem exceção por tipo (RN17, ADR-0003) | Essencial onde a flag estiver ligada |
| RF83 | Ciclo da sala: marcação de ativa pela equipe de inteligência; conclusão com número e hash do relatório e três assinaturas distintas (superior imediato no lugar do solicitante que for comandante ou subcomandante); expurgo na conclusão conforme o tipo; alerta de sala sob custódia ativa além do teto (RN18) | Essencial |
| RF84 | Termo de ciência e aceite da revinculação gerado no app, com foto e assinatura, cifrado para a chave de quórum do termo, hash na trilha e exportação ao sistema corporativo (sem a foto até D-R30) (RN19, doc 07 seção 7.4) | Proposto |
| RF85 | Painel web com indicadores nível 1a, de lista fechada: agregados dos metadados que o servidor já tem, sem nome, sala ou operação identificável, contagens arredondadas ou em faixas, sem filtro abaixo do limite, datas por semana ou mês. **Sem nível 1b no painel** (GD-15; ADR-0008) | Depois da v1 |
| RF86 | Sala de análise com escopo de operações, prazo e revisão definidos pela chefia; nó de IA só cruza salas do escopo de que é membro **e com concessão vigente** (RF92), com filtro antes da busca (ADR-0008) | Depois da v1 |
| RF87 | Resposta da IA de análise em **saída estruturada** validada no código, com referência à origem e sem trecho de conversa, rotulada "indício para análise humana"; toda pergunta registrada e cifrada para a chave de quórum de registros; limite de perguntas (ADR-0008; 01e-E11) | Depois da v1 |
| RF88 | Posição de agentes em missão só pelo MLS na sala da missão; painel e servidor não veem posição (ADR-0008) | Fase 8 |
| RF89 | Autorização por papel e atributos (sala, grau, compartimento, prazo, instituição) num ponto só, com negação por padrão (ADR-0009) | Essencial (Fase 1) |
| RF90 | Atribuição de papel por duas pessoas, com prazo, motivo e recusa de combinações proibidas (ADR-0009) | Essencial (Fase 1) |
| RF91 | Revogar aparelho não destrói: em toda revogação o segredo de lease vigente é selado para a chave de quórum do segredo de lease, com prazo da apuração (D-R35); destruí-lo exige duas autorizações e registro; painel nunca diz "destruído" enquanto ele existir (GD-13; ADR-0006 item 8) | Essencial (Fase 2) |
| RF92 | Concessão de escopo pela sala de origem (GD-15): mensagem de controle assinada pelo administrador (comandante, em missão), com CI e/ou autorização do comandante, visível aos membros, registrada e retirável; regra de tipo (sala de análise sob custódia só com salas sob custódia); expurgo em cascata (ADR-0008) | Depois da v1 |
| RF93 | Investigação sigilosa (ADR-0010): pedido com CI validada e segunda autorização independente, assinado na estação isolada; toda sessão com quórum presencial de detentores não impedidos (GD-21); no máximo duas renovações (GD-22); impedimento conferido pela estação; registro selado para a chave de quórum de registros, com só o hash na trilha; preservação das cópias de custódia e dos registros selados do escopo; sem membro oculto e sem uso do nó de IA | Proposto (Fase 4) |
| RF94 | Relatório de rastro de suspeita (ADR-0010): análise só na estação isolada; saída estruturada sem trecho de conversa, com grau de confiança; conteúdo só com ordem judicial; revisão e assinatura de duas pessoas não impedidas; pacote de evidência (RF67); nenhuma medida automática | Proposto (rito na Fase 4; análise automatizada depois da Fase 6) |

---

## 7. Requisitos não funcionais

### A. Criptografia

| ID | Requisito |
|---|---|
| RNF-A1 | Cifra ponta a ponta de grupo com MLS (RFC 9420), implementação auditada, executada em código nativo |
| RNF-A2 | Cifra simétrica AES-256-GCM (ou ChaCha20-Poly1305); assinaturas Ed25519; troca de chaves X25519 |
| RNF-A3 | Chaves privadas de identidade geradas e mantidas em hardware seguro (StrongBox/TEE), não exportáveis |
| RNF-A4 | Sigilo futuro e recuperação pós-comprometimento garantidos pelo protocolo **para os membros**. Exceção declarada: em sala sob custódia, a cópia para a chave de quórum não tem sigilo futuro; comprometer o HSM offline ou o quórum expõe os períodos guardados (ADR-0003; risco admitido só nessas salas pelo GD-1; aceite formal do R68 pendente) |
| RNF-A5 | Anexos cifrados no aparelho com chave própria por arquivo antes do envio |

### B. Transporte e servidor

| ID | Requisito |
|---|---|
| RNF-B1 | TLS 1.3 com autenticação mútua; certificado do cliente emitido no cadastro, com chave em hardware |
| RNF-B2 | Fixação de certificado (pinning) no app |
| RNF-B3 | Servidor de mensagens não registra conteúdo, nem em log |
| RNF-B4 | Toda a infraestrutura on-premises ou em nuvem privada da instituição; nó de IA sem saída para a internet |

### C. Armazenamento

| ID | Requisito |
|---|---|
| RNF-C1 | Cache local em banco cifrado; chave composta por parte em hardware e parte entregue pelo servidor a cada sessão |
| RNF-C2 | Backup do sistema operacional desativado para os dados do app |
| RNF-C3 | Nomes de objetos no armazenamento de anexos são aleatórios, sem relação com operação, usuário ou arquivo |
| RNF-C4 | Texto de OCR, trechos e embeddings ficam cifrados em repouso com chave da operação |

### D. Integridade do aparelho

| ID | Requisito |
|---|---|
| RNF-D1 | Bloqueio de captura e gravação de tela; conteúdo oculto na lista de apps recentes |
| RNF-D2 | Detecção de: depuração USB/sem fio, opções de desenvolvedor, root, bootloader desbloqueado, emulador, instrumentação (Frida/Xposed) |
| RNF-D3 | Detecção de: serviços de acessibilidade não autorizados, projeção/espelhamento de tela, sobreposições, apps de acesso remoto, tela externa |
| RNF-D4 | Atestação do aparelho validada no servidor no cadastro e periodicamente |
| RNF-D5 | Em estado inseguro: conteúdo oculto, sessão encerrada, evento auditado |
| RNF-D6 | Somente aparelho da instituição, com número de série atestado pelo hardware e presente no inventário (doc 07 seção 3) |
| RNF-D7 | Preparação do aparelho em estação controlada, com triagem de spyware antes do cadastro (doc 07 seção 2) |
| RNF-D8 | Detecção por câmera, no aparelho, de ausência, terceiros e câmera apontada; imagens de evidência apenas em alerta, cifradas para a chave de quórum (doc 07 seção 6) |
| RNF-D9 | Salas com perfil de local autorizado: acesso só por rede ou proximidade autorizadas (doc 07 seção 4) |

### E. Desempenho e disponibilidade

| ID | Requisito | Meta inicial |
|---|---|---|
| RNF-E1 | Entrega de mensagem de texto | ≤ 500 ms no percentil 95, em rede 4G |
| RNF-E2 | Busca local | ≤ 200 ms |
| RNF-E3 | OCR de uma foto | ≤ 5 s |
| RNF-E4 | Correção de texto | ≤ 3 s |
| RNF-E5 | Relatório de análise | ≤ 60 s |
| RNF-E6 | Propagação do expurgo a aparelhos online | ≤ 10 s |
| RNF-E7 | Disponibilidade do serviço de mensagens | 99,5% |

As metas são ponto de partida para validar com a unidade; nenhuma foi medida ainda.

### F. Usabilidade

| ID | Requisito |
|---|---|
| RNF-F1 | Tela de conversa no padrão WhatsApp/Telegram; operador novo envia a primeira mensagem sem treinamento |
| RNF-F2 | Tema escuro como padrão; uso com uma mão |
| RNF-F3 | Indicadores discretos e permanentes: estado de segurança do aparelho, presença do nó de IA, prazo de retenção |

---

## 8. Processos (BPMN simplificado)

Os diagramas usam Mermaid. Raias são representadas por subgrafos.

### 8.1 Cadastro de usuário e aparelho

```mermaid
flowchart LR
  subgraph Gestor
    A([Início]) --> B[Registra matrícula e codinome]
    B --> C[Gera QR de uso único]
  end
  subgraph Operador
    C --> D[Lê o QR presencialmente]
    D --> D2[Assina o Termo de Ciência e Autorização via gov.br, na presença do gestor]
    D2 --> E[App gera chaves no hardware]
    E --> F[App envia atestação e pedido de certificado]
  end
  subgraph Servidor
    F --> G{Aparelho íntegro, gerenciado e termo válido?}
    G -- Não --> H[Recusa e registra em auditoria]
    G -- Sim --> I[Emite certificado e ativa o aparelho]
  end
  I --> J([Aparelho ativo])
  H --> K([Fim])
```

### 8.2 Criação de operação e convite

```mermaid
flowchart LR
  subgraph Administrador
    A([Início]) --> B[Cria operação e define retenção]
    B --> C[Seleciona membros]
    C --> D[App cria o grupo cifrado e inclui o nó de IA]
  end
  subgraph Servidor
    D --> E[Distribui convites cifrados]
    E --> F[Registra evento em auditoria]
  end
  subgraph Operador
    E --> G[Recebe convite]
    G --> H{Aceita?}
    H -- Sim --> I[Entra na operação]
    H -- Não --> J([Fim])
  end
  I --> K([Operação ativa])
```

### 8.3 Envio de foto com OCR

```mermaid
flowchart LR
  subgraph Remetente
    A([Tira foto no app]) --> B[Remove metadados]
    B --> C[Cifra o anexo e a mensagem]
  end
  subgraph Servidor
    C --> D[Guarda anexo cifrado com nome aleatório]
    D --> E[Encaminha mensagem cifrada]
  end
  subgraph Nó de IA
    E --> F[Decifra em memória]
    F --> G[OCR]
    G --> H[Gera embedding e guarda cifrado]
    H --> I[Envia texto do OCR como mensagem cifrada]
  end
  subgraph Destinatários
    E --> J[Exibe foto com marca d'água]
    I --> K[Exibe e indexa o texto localmente]
  end
```

### 8.4 Pedido de exportação de anexo

```mermaid
flowchart LR
  subgraph Operador
    A([Solicita exportar]) --> B[Informa justificativa]
  end
  subgraph Administrador
    B --> C{Autoriza?}
    C -- Não --> D[Nega]
    C -- Sim --> E[Assina autorização com validade]
  end
  subgraph App do Operador
    E --> F[Verifica assinatura]
    F --> G[Gera cópia com marca d'água]
  end
  subgraph Servidor
    D --> H[Registra em auditoria]
    G --> H
  end
  H --> I([Fim])
```

### 8.5 Análise da operação por IA

```mermaid
flowchart LR
  subgraph Administrador
    A([Toca em Analisar]) --> B[Confirma com biometria]
  end
  subgraph Nó de IA
    B --> C[Reúne mensagens e textos de OCR em memória]
    C --> D[Gera relatório estruturado]
    D --> E[Cifra o relatório só para o administrador]
  end
  subgraph Servidor
    E --> F[Encaminha e registra o evento]
  end
  F --> G([Administrador lê no app])
```

### 8.6 Expurgo da operação

```mermaid
flowchart LR
  subgraph Administrador
    A([Aciona expurgo]) --> B[Biometria + nome da operação]
    B --> C[Assina a ordem]
  end
  subgraph Servidor
    C --> D{Assinatura e sequência válidas?}
    D -- Não --> E[Descarta e audita tentativa]
    D -- Sim --> F[Difunde a ordem]
    F --> G[Apaga fila, anexos e metadados da operação]
    G --> H[Registra expurgo em auditoria]
  end
  subgraph Nó de IA
    F --> I[Destrói chave da operação e índice em memória]
    I --> J[Apaga registros cifrados]
  end
  subgraph Aparelhos
    F --> K{Online?}
    K -- Sim --> L[Verifica assinatura e apaga chaves e cache]
    K -- Não --> M[Perde acesso ao vencer a autorização]
    L --> N[Confirma ao servidor]
  end
  H --> O([Operação expurgada])
```

Em **sala sob custódia**, o mesmo fluxo vale para a cópia de trabalho. A cópia de custódia não é tocada: fica no serviço `custody` até o fim do prazo de guarda e é eliminada em cerimônia, com recibo (ADR-0003 seção 7). O mesmo fluxo é disparado pela conclusão (RN18).

---

## 9. Governança e conformidade

Este item precisa de validação do jurídico da instituição; o texto abaixo é ponto de partida, não parecer.

- **Base normativa do uso.** Definir o ato interno que autoriza a ferramenta, quem pode ser administrador de operação e quem responde pelo expurgo.
- **LGPD.** O tratamento para segurança pública e investigação tem regime próprio (art. 4º, III), mas os princípios gerais e o devido processo continuam aplicáveis. Registrar finalidade, retenção e responsáveis.
- **Cadeia de custódia.** O sistema é canal de coordenação, não repositório de prova. Informação que precise integrar procedimento formal deve ser levada ao sistema oficial pelo fluxo de exportação autorizada. Isso deve constar na norma de uso. Em sala sob custódia, a cópia institucional existe para atender ordem judicial, ou processo administrativo com validação jurídica, sempre com CI, por quórum presencial (RN20); não substitui a exportação para o sistema oficial.
- **Controle.** A trilha de metadados existe para que corregedoria e controle externo possam verificar *que* houve uma operação, quem participou e quem a encerrou, sem acesso ao conteúdo.
- **Uso da IA.** Relatórios são apoio à decisão. Não substituem análise humana e não devem ser citados como fonte em documento oficial.
- **Transparência interna.** Todo membro é informado, no convite, de que há nó de IA na sala e de que o administrador pode pedir análise do conteúdo.

---

## 10. Critérios de sucesso

| Critério | Como verificar |
|---|---|
| A unidade abandona o WhatsApp para operações | Uso semanal por 90% da equipe após 60 dias de piloto |
| Captura de tela e gravação não funcionam em aparelho homologado | Checklist do documento 03 |
| Vazamento por foto de tela é atribuível | Teste cego: identificar o aparelho de origem a partir de uma foto |
| Expurgo não deixa conteúdo recuperável | Teste forense de aparelho e de servidor após expurgo |
| Servidor comprometido não revela conteúdo | Pentest com acesso total ao servidor de mensagens |
| Homologação pela área de segurança da instituição | Parecer formal |

---

## 11. Pontos em aberto

1. Modelos de aparelho e solução de MDM disponíveis na instituição.
2. Prazo padrão de retenção e de renovação de autorização (propostos: 7 dias e 12 h).
3. Hardware do nó de IA (GPU) e local de hospedagem.
4. Necessidade de dupla custódia para expurgo.
5. Quem exerce o papel de Auditor e com que rito acessa a trilha.
6. Volume esperado: usuários, operações simultâneas, fotos por dia.
7. Necessidade de operação em área sem cobertura (modo offline e seus limites).
8. Como entregar o aviso de nova mensagem sem serviço de terceiro (decisão D-R1, doc 04).
9. Rito de quebra de pseudonimato e variação de conteúdo por destinatário (D-R2 e D-R3, docs 04 e 05).
10. Enquadramento LGPD, base legal, retenção da trilha e compatibilidade do expurgo com dever de guarda (D-R4 e D-R5, doc 06).
11. ~~Exceções à custódia com a flag ligada (D-R29)~~: decidido no GD-14 (RN17); resta quem dá a segunda assinatura (D-R36). Foto do termo de revinculação (D-R30), doc 07 seção 13.
12. ~~Quem assina a conclusão quando o solicitante é o próprio comandante ou subcomandante~~: o superior imediato do solicitante (RN18). Resta quem autoriza trocar o titular de um papel (pergunta 3 do 01c).
13. Investigação sigilosa (ADR-0010): segunda autorização e substitutos (D-R37); prazo, contraditório e aviso no arquivamento (D-R38).
14. Escada por perfil e reconhecimento facial (GD-18): números e perfil missão na polícia (D-R40); parecer jurídico sobre biometria.
15. Concessão de escopo de análise (GD-15): base mínima, retroação, agregados (D-R41).
