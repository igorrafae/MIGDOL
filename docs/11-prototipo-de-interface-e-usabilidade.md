# 11 — Protótipo de interface e protocolo de usabilidade

**Status:** rascunho para validação; corrigido na quarta rodada da Fase 0 (`docs/fase0/01h-correcoes-rodada4.md`: 01g-G14 a G19, G25; fusão dos riscos e decisões no doc 04) · **Origem:** GD-19 (docs/fase0/GD-respostas.md) · **Base:** ADR-0007 (plataformas), ADR-0008 (painel e IA), ADR-0009 (papéis), doc 01 (RF e RNF-F), doc 03 seção 4.12 (V09), doc 05 (playbooks), doc 07 seção 7 (coação, perda, revinculação), doc 10 (campo)

> Nota: a ADR-0007 está **aceita** (GD-19). As ADR-0008, 0009 e 0010 são propostas. Telas que dependem de ADR ainda não aceita estão marcadas como **[condicionada]**. Telas do painel usam o prefixo **PW** (PW01 a PW12), para não colidir com os playbooks P1 a P11 e P1-S do doc 05.

**Todos os nomes, codinomes, salas, operações e locais deste documento e dos protótipos são fictícios.** Nenhum dado real entra (regra 10).

---

## 1. Objetivo e escopo

**Objetivo:** mostrar e testar a interface dos três projetos antes de escrever código de produção, e provar com agentes de verdade (V09) que os ritos de segurança não empurram o usuário para um mensageiro comercial (R120, seção 7).

| Projeto | Tecnologia (ADR-0007) | Papéis que usam | Entra em protótipo na fase | Observação |
|---|---|---|---|---|
| **Android** | Jetpack Compose | Agente, comandante, subcomandante, administrador de operação, equipe de inteligência, analista [condicionada] | **Fase 0** (papel e navegável) e **Fase 1** (nativo, mensagens de texto) | Telas de proteção do aparelho entram na Fase 2; anexos, na Fase 3 |
| **iPhone** | SwiftUI | Os mesmos do Android | **Estudo de viabilidade (Fase 0):** só protótipo em papel e navegável, **sem código de produto**. Nativo só depois da ADR própria (D-R31) | Sem macOS como cliente. A interface é feita duas vezes (ADR-0007, consequências) |
| **Painel web Linux** | Web, só rede interna | Gestor, equipe de inteligência (cadastro), auditor (leitura de trilha de metadados) | **Fase 1** (mínimo: cadastro, revogação, estado dos aparelhos). Indicadores e nós de campo na fase do nó de IA e na Fase 8 | Nível 1a apenas (ADR-0008). Sem conteúdo, sem posição |

**Dentro do escopo:** mapa de telas, fluxos críticos, fidelidade do protótipo, protocolo de teste, riscos, critérios de aceite.
**Fora do escopo:** identidade visual final, textos jurídicos do termo (anexo próprio), desenho de HSM e cerimônia de quórum (a cerimônia é presencial e fora do app; o protótipo cobre só o que o app mostra).

---

## 2. Princípios de interface

Cada princípio vem de uma regra do projeto e vira critério de revisão de toda tela.

| ID | Princípio | Origem | Como se verifica no protótipo |
|---|---|---|---|
| PI-1 | **Nenhum dado fora do core.** A interface só mostra o que o core entrega e não guarda estado próprio: sem rascunho em SharedPreferences, DataStore, arquivos, UserDefaults, Core Data. Rascunho de mensagem vive no core, cifrado | Regras 2 e 4; ADR-0007 item 5 | Revisão da lista de estados da tela: todo campo tem dono no core. Teste estático em `tests/security` (ADR-0007) |
| PI-2 | **Notificação sem conteúdo.** Nunca nome, sala, trecho ou codinome de remetente. Texto genérico ("Há novidades") ou nenhuma, conforme perfil. Sem push de terceiro (regra 10; D-R1: conexão própria) | Regras 3 e 10; D-R1 | Lista de notificações possíveis, todas revisadas; captura da tela de bloqueio nos testes |
| PI-3 | **Marca de custódia sempre visível.** Em toda sala e no cabeçalho da conversa: "sob custódia" ou "comum", com o motivo (ex.: investigação militar). A marca é assinada, vem do estado do grupo e **não se desliga** na sala (GD-3). O indicador não deve ser removível por configuração | Regra 14; GD-3; RF82; RNF-F3 | Todas as telas de sala mostram a marca; teste com marca adulterada: a tela mostra erro, não "comum" |
| PI-4 | **Telas de coação indistinguíveis do normal.** A conta isca usa a mesma interface, os mesmos tempos de resposta, as mesmas mensagens de erro e as mesmas animações. Nenhum selo, cor ou aviso diferente. O alerta silencioso não aparece na tela | Doc 07 seção 7 (PIN de coação "nunca revela que é isca") | Comparação lado a lado e medição de tempo de abertura (PIN normal x coação) com diferença dentro de margem a definir (D-R47), medida por instrumento |
| PI-5 | **Sem captura de tela e sem gravação de tela** nas telas de conteúdo (Android: bandeira de janela segura; iPhone: sem garantia equivalente, tratar na ADR do iOS: **[a verificar]**). Marca d'água visível em tela (RF22) | RF22; doc 02 | Tentativa de captura nos testes; no iPhone, o resultado entra no estudo de viabilidade |
| PI-6 | **Modo escuro como padrão e modo de campo.** Escuro (RNF-F2); "campo" = brilho baixo, contraste alto, sem cor vermelha pura isolada para sinalizar estado, sem som por padrão | RNF-F2; doc 10 | Teste de pouca luz (seção 6.4) |
| PI-7 | **Uso com luva e uma mão.** Alvos de toque grandes (meta inicial: 56 dp ou equivalente, a validar no teste), ações principais na metade inferior, sem gesto de precisão como único caminho | RNF-F2 | Teste com luva fina e grossa |
| PI-8 | **Acessibilidade mínima.** Texto escalável até 200% sem quebrar layout; contraste mínimo AA; nenhum estado só por cor (ícone + texto); leitor de tela nas telas de rotina; sem dependência de som. As telas de coação e pânico não anunciam nada por leitor de tela | WCAG 2.1 AA como referência [conferir exigência do edital] | Revisão de acessibilidade por tela |
| PI-9 | **O rito é curto e explica o motivo.** Todo bloqueio diz o que aconteceu, o que fazer e quem procurar, sem expor detalhe técnico que ajude um atacante. Menos passos que o mensageiro comercial é impossível; o objetivo é o mínimo de passos que o rito permite | V09; R120 | Contagem de toques por tarefa (seção 6.3) |
| PI-10 | **Estado de segurança discreto e permanente** (RNF-F3): estado do aparelho, presença do nó de IA como membro visível, prazo de retenção. Nunca alarmante sem ação possível | RNF-F3 | Revisão de cada tela |
| PI-11 | **A face nunca decide sozinha, e a tela não sugere que decide.** Falha facial isolada não mostra "destruição" nem "ameaça". Só pede o PIN | GD-18; regra 11 | Fluxo 4.8 |
| PI-12 | **Nenhuma decisão contra pessoa aparece como automática.** Indicador, IA ou marca d'água sempre rotulados "indício para análise humana" | Regra 11; RF87 | Revisão de textos |
| PI-13 | **Linguagem do usuário.** Sem "época", "commit", "lease", "HPKE" na tela do agente. Termos internos ficam na documentação. O gestor e o auditor veem termos de negócio ("segredo guardado sob quórum até DD/MM/AAAA", nunca "destruído": D27, doc 07 seção 7.1) | RNF-F1 | Revisão de texto por tela |

---

## 3. Mapa de telas

### 3.1 Android e iPhone (mesmas telas)

Legenda de papéis: **AG** agente; **AO** administrador de operação (comandante ou substituto na missão); **EI** equipe de inteligência; **SO** solicitante; **CH** chefe da inteligência e superior da área; **AN** analista **[condicionada ADR-0008/0009]**.

| ID | Tela | Papéis | Requisito ligado | Fase |
|---|---|---|---|---|
| T01 | Ativação do aparelho (leitura do QR de uso único, geração de chaves, atestação) | AG | RF01 | 1 |
| T02 | Termo de Ciência e Autorização: texto, versão e estado que o core entrega; bloqueio até a assinatura constar. **A assinatura no gov.br é feita fora do app, na estação do gestor** (regra 4; doc 06 seção 7) | AG | RF06; RN16 | 1 |
| T03 | Desbloqueio (biometria ou PIN) | todos | RF02 | 1 |
| T04 | Lista de salas (marca de custódia, estado: ativa ou concluída, prazo de retenção) | todos | RF10; RF82; RNF-F3 | 1 |
| T05 | Conversa (texto, confirmação de entrega e leitura, busca local) | AG, AO | RF12; RF30; RNF-F1 | 1 |
| T06 | Membros da sala (inclui nó de IA, sempre visível) | AG, AO | RNF-F3; doc 02 | 1 |
| T07 | Criar sala: nome, política de retenção, membros, **marca de custódia** | AO | RF10; RF82; GD-14 | 1 |
| T08 | Criar sala comum com a flag ligada: motivo + segunda assinatura | AO + segundo assinante | GD-14; D-R36 | 1 |
| T09 | Ciclo da sala: marcar ativa (EI), pedir conclusão, três assinaturas, número e hash do relatório | EI, AO, SO | RF83; RN18; GD-3 | 4 |
| T10 | Câmera interna e envio de foto (sem galeria) | AG | RF13; RF23 | 3 |
| T11 | Visualizador de anexo (imagem, PDF, áudio), com marca d'água | AG | RF20; RF22 | 3 |
| T12 | Pedido de exportação e aprovação | AG, AO | RF21 | 3 |
| T13 | Assistência: "Corrigir texto" (sugestão aceita ou descartada) | AG | RF40 | 5 |
| T14 | "Analisar operação" e relatório (rótulo permanente de apoio à decisão) | AO | RF42–RF44 | 6 |
| T15 | Expurgo: confirmação reforçada (biometria + digitar o nome da sala) e situação por aparelho | AO | RF50–RF52 | 4 |
| T16 | Aviso de trava, nível 1 ou 3 (genérico, com a quem procurar) | AG | doc 07 seção 7.1 | 2 |
| T17 | Gesto ou botão de pânico (discreto, com confirmação curta que não comprometa o disfarce) | AG | doc 07 seção 7 | 2 |
| T18 | Conta isca (mesmas telas T03 a T06, conteúdo genérico; **não existe tela própria**) | AG | doc 07 seção 7; PI-4 | 2 |
| T19 | Verificação facial **durante a digitação do PIN** em T03 (câmera interna; sem tela própria e sem passo que dependa do PIN; igual com o PIN de coação; nenhuma tela mostra o resultado) | AG | GD-18; GD-24; doc 07 seções 7.1a e 7.2 | 2 [condicionada ao parecer jurídico e à D-R40] |
| T20 | Aviso "você foi informado depois": registro do que o aparelho fez, mostrado na próxima autenticação legítima | AG | RN15 | 2 |
| T21 | Código de segurança entre membros (comparação presencial ou QR) | AG | RF04 | 1 |
| T22 | Aprovação dupla de revinculação (aprovador assina o pedido, com resumo legível e hash visível) | CH | RN19; GD-4; R69 | 4 (doc 07 seção 14) |
| T23 | Termo de ciência e aceite da revinculação (assinatura na tela, foto com a câmera interna, hash, exportação) | AG (proprietário) | RF84; GD-4; GD-5 | 4 |
| T24 | Envio de histórico ao aparelho novo (membro indicado vê o que será enviado, dentro dos intervalos de pertença) | AG membro | doc 07 seção 7.4 | 4 |
| T25 | Sala de missão: mapa de posições dos agentes, exibido no aparelho do comandante e do substituto (e de quem o grupo definir **[pergunta aberta]**). **Limite:** é restrição de interface; todo aparelho membro decifra a posição (ADR-0008 seção 3; 01g-G19; D-R20) | AO | RF88; ADR-0008 seção 3 | 8 |
| T26 | Estado do enlace de campo (nó de campo, rádio, banda estreita, lease de campo) | AG, AO | doc 10 seções 3 a 6 | 8 |
| T27 | Mensagem em quadro de banda estreita (limite de tamanho visível, sem anexo) | AG | doc 10 seção 6 | 8 |
| T28 | Sala de análise: escopo, prazo, revisão; pergunta à IA; resposta com referência e rótulo "indício para análise humana" | AN | RF86; RF87; ADR-0008 | depois da Fase 6 [condicionada] |
| T29 | Pedido e concessão de escopo de análise (CI e/ou autorização do comandante, registrada) | AO, AN | GD-15 | depois da Fase 6 [condicionada] |
| T30 | Aviso de transparência: "o conteúdo desta sala pode entrar numa sala de análise" | AG | doc 06 seção 7; ADR-0008 item 2 | depois da Fase 6 [condicionada] |

**Diferenças de plataforma anotadas:**

| Tema | Android | iPhone |
|---|---|---|
| Captura de tela | Bandeira de janela segura | Sem equivalente garantido; resultado do estudo de viabilidade decide (PI-5) |
| Notificação | Conexão própria em segundo plano (D-R1); texto genérico | Depende do APNs (dependência de terceiro, D-R1 e ADR-0007); sem conteúdo no payload; pode ser inviável (D-R31) |
| Chave de hardware | StrongBox / Keystore; a interface segura só a referência (GD-12) | Secure Enclave; idem |
| Biometria | BiometricPrompt, classe forte exigida | Face ID e Touch ID; PIN próprio, nunca só o código do aparelho **[a verificar]** |
| Câmera interna | CameraX sem galeria | AVFoundation sem galeria |
| Atestação | Sem Play Integrity (regra 12) | App Attest consulta servidor da Apple (conflita com a regra 12): ponto do estudo de viabilidade |
| Gesto de pânico | Sequência de botão de volume ou gesto na tela **[a testar]** | Gesto na tela; botões físicos têm menos opções **[a verificar]** |
| MDM / apagamento | Android Enterprise, aparelho totalmente gerenciado | MDM da Apple (também depende do APNs) |
| Gerar telas | Compose | SwiftUI; mesma lista T01 a T30, mesmo texto, mesmo fluxo (R102) |

### 3.2 Painel web Linux (rede interna, token de hardware)

Só indicadores de **nível 1a** (ADR-0008 item 1): agregados dos metadados que o servidor já tem, **sem nome, sala ou operação identificável**, contagens **arredondadas ou em faixas**, sem filtro que reduza o grupo abaixo do limite mínimo (D-R32; 01e-E14), datas por semana ou mês. O painel **nunca** mostra conteúdo, posição, nome de sala ou de operação, indicador nominal nem preservação selada (regra 3; GD-15; ADR-0010). **O nível 1b não existe** (GD-15). No perfil Exército, missões ativas e nós de campo ficam fora do painel até a D-R20.

| ID | Tela | Papel | Requisito ligado | Fase |
|---|---|---|---|---|
| PW01 | Entrada com token de hardware e escolha de instituição | todos | RNF-T3; ADR-0008 | 1 |
| PW02 | Cadastro de usuário (codinome; matrícula e nome real **não** aparecem em sala) | Gestor | RF03; RF05 | 1 |
| PW03 | Cadastro e estado de aparelhos: série, estado (ativo, travado, revogado), atestação, inventário | Gestor | RF01; RF59; RF60 | 1 |
| PW04 | Revogar aparelho: motivo, texto "cache ilegível, segredo guardado sob quórum até DD/MM/AAAA", segunda autorização para destruir | Gestor + segunda pessoa | RF05; RF91; GD-13 | 1 e 2 |
| PW05 | Declaração de perda, roubo ou apreensão (motivo; trava imediata; apagamento só com segunda autorização) | Gestor | doc 07 seção 7.4 | 2 |
| PW06 | Termo de Ciência: situação por usuário (vigente, vencido, nova versão pendente) | Gestor | RF06; RN16 | 1 |
| PW07 | Atribuição de papéis: duas pessoas, motivo, prazo, recusa de combinação proibida; papéis vencidos. A atribuição de papel privilegiado é **objeto assinado** pelas duas pessoas, conferido contra a lista assinada em cerimônia, nunca só contra o diretório (ADR-0009 seção 3) | Gestor + segunda pessoa | RF90; ADR-0009 | 1 |
| PW08 | KPIs nível 1a, de lista fechada: aparelhos por estado, revinculações, alertas de integridade, salas por estado do ciclo (RN18), duração das salas; nós de campo só com a D-R20 no perfil Exército | Gestor, chefia | RF85 | nó de IA / depois da Fase 6; versão mínima na Fase 1 |
| PW09 | Prazos de custódia: salas sob custódia por faixa de prazo e vencimentos próximos, **sem identificar sala**; expurgos pendentes; ordens de preservação **comuns** (RF58), em faixa. A preservação selada da investigação sigilosa **nunca** aparece: o servidor nem a conhece (ADR-0010 seção 8; 01g-G16) | Gestor, chefia | GD-1; RF52; RF58 | 4 |
| PW10 | Nós de campo: estado, lease de campo, última sincronização, lista de revogação de campo. No perfil Exército, só com a D-R20 | Gestor | doc 10 seção 3; RF das seções 10 | 8 |
| PW11 | Trilha de auditoria de metadados (consulta); botão "pedir abertura pelo rito" que abre um pedido, não abre dado | Auditor | RF53; regra 14 | 4 |
| PW12 | Resumo do estado do exercício: expurgos com recibo, aparelhos sem atestação | Gestor | RF52 | 4 |

**Não existe no painel (e o protótipo mostra o buraco, de propósito):** lista de salas com nome, lista de membros por sala, posição de agente, pergunta à IA, indicador nominal, botão "abrir custódia", botão de aprovar revinculação (a aprovação é no app, T22).

---

## 4. Fluxos críticos

Convenção: **Passo** é o que a pessoa faz ou vê; **Erro** é o estado de erro que o protótipo precisa ter. Tela entre colchetes.

### 4.1 Cadastro e termo (AG, gestor) [T01, T02, T03; PW02, PW03, PW06]
1. Gestor cadastra o usuário e o aparelho no painel, presencialmente, e emite o QR de uso único.
2. Agente abre o app e lê o QR [T01]. O aparelho gera as chaves no hardware e envia a atestação.
3. Agente define o PIN e escolhe a biometria. O app pede também o **PIN de coação** (segundo PIN), com texto de treinamento curto.
4. Agente lê o Termo de Ciência no app [T02] e o **assina pelo gov.br na estação do gestor**, fora do app, na presença dele (doc 06 seção 7). A assinatura é validada no servidor ou na estação do gestor; o app só mostra o estado que o core recebe (assinatura constando, versão). O aparelho não fala com o gov.br e o CPF não é digitado nele (regra 4; 01g-G14).
5. App ativa e mostra a lista de salas vazia [T04].

| Erro | Tela e texto |
|---|---|
| QR vencido ou já usado | "Código inválido. Procure o gestor para um novo." Nada mais |
| Atestação falhou | "Este aparelho não pode ser ativado. Procure o gestor." |
| Assinatura gov.br recusada ou fora do ar (na estação do gestor) | O app mostra "Assinatura do termo pendente. Procure o gestor." e permanece inativo; a contingência é a do doc 06 seção 7.1 |
| PIN comum igual ao PIN de coação | Recusa com mensagem neutra: "Escolha outro PIN." (não explicar o porquê na tela) |
| Nova versão do termo | Acesso suspenso até assinar (doc 03 Fase 1) |

### 4.2 Sala e envio de mensagem (AO, AG) [T07, T04, T05, T06]
1. AO cria a sala: nome fictício, retenção, membros, marca de custódia [T07].
2. Cada membro recebe o convite pelo próprio app. A tela mostra a marca de custódia e o motivo antes de aceitar.
3. AG abre a sala [T04], vê a marca, a retenção e o nó de IA na lista de membros [T06], e digita a mensagem [T05].
4. Confirmação de entrega e de leitura aparecem ao lado da mensagem.

| Erro | Tela e texto |
|---|---|
| Sem rede | Mensagem fica "aguardando envio", cifrada no core; nenhum texto em notificação |
| Lease vencido | Cache ilegível; "Reconecte para continuar." Nada é apagado |
| Aparelho removido da sala | "Você não está mais nesta sala." Sem detalhe |
| Mensagem maior que o limite | Recusa antes de enviar, com o limite |

### 4.3 Anexo (AG, AO) [T10, T11, T12]
1. AG abre a câmera do app [T10], tira a foto. Não há acesso à galeria.
2. O app remove metadados (RF23) e mostra "Foto sem localização" antes de enviar.
3. O destinatário vê a imagem no visualizador interno com marca d'água [T11].
4. Para exportar, pede a exportação [T12]; AO aprova ou nega; a saída é pelo pacote de evidência assinado (RF67).

| Erro | Tela e texto |
|---|---|
| Captura de tela tentada | Tela fica preta; nenhuma mensagem que denuncie o bloqueio ao terceiro que olha |
| Anexo grande demais | Recusa com limite |
| Exportação negada | "Pedido negado." |

### 4.4 Marca de custódia na criação, sala comum com motivo e segunda assinatura [T07, T08]
Com a flag ligada, a custódia é o padrão (GD-14).
1. AO escolhe o tipo na criação. O padrão é "sob custódia", com motivo (ex.: "investigação militar").
2. Para criar **sala comum** com a flag ligada, o app exige **motivo registrado** e **segunda assinatura** de outra pessoa [T08]. Quem assina: **D-R36**.
3. A marca é assinada e entra no estado do grupo; fica visível aos membros e **não há botão para mudá-la**.
4. Sala de sistema é sempre comum (GD-14) e não passa por este fluxo.

| Erro | Tela e texto |
|---|---|
| Segundo assinante recusa ou não responde | Sala não é criada como comum; o app oferece "sob custódia" |
| Tentativa de mudar a marca depois | Não há controle na tela; se houver tentativa por outro caminho, o app mostra erro e registra |
| Marca divergente do pacote da instituição | "Esta sala não pode ser aberta." A tela não oferece forma de ignorar |

### 4.5 Conclusão com três assinaturas [T09]
1. EI marca a sala como **ativa** (RN18).
2. Para concluir, o relatório interno precisa estar assinado no sistema corporativo. O app pede o **número e o hash** do relatório (RF83).
3. Assinam, cada um no próprio app, o **comandante da missão**, o **subcomandante** e o **solicitante**. Se o solicitante é comandante ou subcomandante, assina o **superior imediato** definido na criação (01c).
4. Cada assinante vê o resumo (sala, relatório, quem já assinou) e confirma com biometria.
5. Concluída: sai dos aparelhos. Sala comum é expurgada com recibo. Sala sob custódia: a cópia de trabalho sai dos aparelhos e a tela diz "concluída; cópia de custódia guardada até DD/MM/AAAA" (ADR-0003 seção 7; GD-1; 01g-G25e), nunca "destruída".

| Erro | Tela e texto |
|---|---|
| Mesma pessoa tenta assinar duas vezes | Recusa: "Você já assinou." |
| Assinatura ausente | A conclusão fica pendente e mostra quem falta |
| Hash do relatório não confere | "Relatório não confere." Não conclui |
| Papel vencido (R109) | "Seu papel venceu. Procure o gestor." |

### 4.6 Perda do aparelho [PW05, PW04; T16; do ponto de vista do agente]
1. Agente avisa o gestor **por outro canal** (a norma de uso define: a lista de canais e quem responde: doc 05 seção 9).
2. Gestor abre PW05 e declara o motivo: perdido, roubado, suspeita de comprometimento ou apreendido por autoridade.
3. Perdido, roubado ou comprometido: trava imediata (nível 3); servidor para de entregar o segredo de lease; cache ilegível. **Nada é destruído.** Apreensão por autoridade: suspender e revogar sem apagar (D-R14).
4. Apagamento exige **segunda autorização** de pessoa distinta; pode vir depois.
5. O painel mostra "cache ilegível, segredo guardado sob quórum até DD/MM/AAAA" (PW04).
6. Agente recebe o aviso pelos canais anteriores (RN15).

| Erro | Tela e texto |
|---|---|
| Segunda autorização não chega | O painel segue em "travado", com o prazo e o alerta (doc 07 seção 7.4); nada é apagado |
| Aparelho offline | Painel diz "ordem pendente" e "cache ilegível por lease"; não promete apagamento |
| Aparelho revogado reconecta | Recebe as ordens pendentes, nunca o segredo (doc 07 seção 7.1) |

### 4.7 Coação (PIN de coação) e escada de resposta [T03, T17, T18, T16]
Fonte: matriz do doc 07 seção 7.1a, com as condições da GD-24. O efeito depende do **perfil**: urbano (e fora de missão) ou missão de campo (só durante missão em andamento, no perfil da instituição; D-R40).
1. Sob coação, o agente digita o **PIN de coação** em T03. A verificação facial (4.8) aparece **igual** à do PIN normal.
2. O app abre a **conta isca** com a mesma interface [T18]: conversas e salas genéricas. O alerta silencioso sai ao Auditor e à segurança; evidência de câmera, se houver alerta, só cifrada para a chave de quórum (regra 12).
3. Efeito por baixo da isca: **urbano**, nível 3 (suspende; nada se destrói). **Missão de campo**, nível 4 (apagamento local), **só depois de medido o acionamento acidental do PIN de coação** (GD-24; tarefa U05); até lá, nível 3. No nível 4, o alerta é enfileirado e enviado, se houver enlace, antes de zerar.
4. Nada na tela, no som, na vibração, no tempo, no aquecimento ou em notificação diferencia o PIN de coação do normal (PI-4; margem: D-R47).
5. **Pânico** [T17]: gesto discreto oculta tudo, alerta pela sala de resposta e aciona os níveis 2 e 3 da escada. O nível 4 por pânico só com decisão escrita do jurídico (D-R14, D23).
6. Depois, o agente legítimo é informado na próxima autenticação legítima [T20].

Escada (doc 07 seção 7.1) como o usuário a vê:

| Nível | O que o usuário vê | Reversível |
|---|---|---|
| 1. Travar | Tela de desbloqueio; conteúdo oculto | Sim |
| 2. Isca e alerta silencioso | Conta isca, sem sinal | Sim, pelo gestor |
| 3. Trava criptográfica | "Procure o gestor." (T16) | Sim, pelo gestor |
| 4. Apagamento local (só perfil missão, nas condições da GD-24) | Nada: o aparelho já foi zerado, ou a isca continua na tela | Não, para o aparelho; volta por revinculação |

| Erro | Comportamento |
|---|---|
| Pânico acionado sem querer | Confirmação curta que não compromete; o retorno exige o gestor e a norma de uso; o protótipo testa a taxa de acionamento acidental (U06) |
| PIN de coação digitado quando não há coação | Mesmo efeito do caso real, **conforme o perfil**: urbano, isca e nível 3, volta pelo gestor; missão de campo com o nível 4 ligado, o aparelho é zerado e volta por revinculação. O agente avisa o gestor; sem consequência contra pessoa (regra 11). Por isso o nível 4 pelo PIN de coação só liga depois da medição (GD-24; R55) |
| Alerta sem rede | Fica na fila do core e sai na próxima conexão; no nível 4 sem enlace, o alerta se perde com o aparelho (limite declarado, doc 07 seção 7.1) |
| Rede, nenhum receptor responde | Alerta de segundo nível ao gestor (D17) |

### 4.8 Verificação facial de quem acessa [T03, T19] **[condicionada ao parecer jurídico e à D-R40]**
Reconhecimento no aparelho, gabarito cifrado por chave de hardware, sem terceiros (regra 12; GD-18; doc 07 seção 7.2). **A face nunca decide sozinha.** A função é **detectar**, não barrar (doc 07 seção 7.2).
1. Enquanto o agente digita o PIN em T03, o core faz as capturas com a câmera interna [T19]. Não há tela própria nem passo que dependa de o PIN estar certo, e o comportamento é igual com o PIN de coação. Instrução mínima ("olhe para a tela").
2. O resultado ("dono", "outra pessoa" em duas capturas, "falsificação detectada" ou "inconclusivo") e o PIN definem o efeito pela matriz do doc 07 seção 7.1a:

| Situação | Perfil urbano (e fora de missão) | Perfil missão de campo | O que a tela mostra |
|---|---|---|---|
| Dono e PIN correto | Abre | Abre | Abre normalmente |
| Inconclusivo (sem rosto, coberto, pouca luz) e PIN correto | Abre, com evento | Idem | Abre normalmente |
| Dono ou inconclusivo, e PIN errado | Conta tentativa; N erros → nível 3 | Idem (nível 4 por tentativas: D-R40) | "PIN incorreto" |
| Outra pessoa (duas capturas) ou falsificação, e PIN correto | Nível 2: isca e alerta | Idem | A isca, sem sinal |
| Outra pessoa (duas capturas) e PIN errado | Nível 3 | Nível 4, só depois das medições da GD-24; até lá, nível 3 | "PIN incorreto" ou nada |
| Falsificação, ou outra pessoa numa só captura, e PIN errado | Conta tentativa; N erros → nível 3 | Idem: nunca nível 4 até a vivacidade ser medida à parte | "PIN incorreto" |
| PIN de coação, com qualquer rosto | Ver 4.7 | Ver 4.7 | A isca, sem sinal |

3. Nenhuma tela mostra "rosto não reconhecido", "pontuação", "falha facial", "destruição" nem "ameaça" (PI-11).

| Erro | Comportamento |
|---|---|
| Falso negativo do dono (luz, luva, barba, óculos, pintura, capacete) com PIN correto | Lido como "inconclusivo": abre. Se a leitura der "outra pessoa", o dono cai na **isca** (nível 2) |
| **Dono caiu na isca** por falso negativo | Percebe pelo conteúdo genérico; procura o gestor por outro canal (norma de uso); o gestor confere a identidade presencialmente, registra o falso negativo e libera o aparelho. **Sem destruição e sem consequência contra pessoa** (regra 11). O tempo até voltar a trabalhar é medido no protótipo (seção 6.8) |
| Câmera indisponível ou coberta | "Inconclusivo": cai no PIN; o app registra |
| Falha de vivacidade | Nível 2 com PIN correto; com PIN errado, só conta tentativa (nunca nível 4 até medição à parte; GD-24) |

### 4.9 Revinculação com aprovação dupla e termo com foto [T22, T23, T24; PW03]
Vale só para operação **ativa** (GD-4). Não decifra nada da instituição.
1. Aparelho antigo revogado; aparelho novo preparado e triado; série no inventário; cadastro presencial; Termo vigente.
2. Gestor registra o pedido com agente, aparelho novo, aparelho antigo e salas.
3. **Chefe da inteligência e superior da área** assinam no app [T22]. Veem o resumo legível. Duas pessoas distintas entre si, do agente e do gestor que cadastrou.
4. Agente (proprietário) assina o **termo de ciência e aceite** no próprio app, com assinatura na tela e foto com a câmera interna [T23]. O app cifra o termo para a chave de quórum do termo, grava o hash na trilha e exporta o documento para o sistema da corporação. **Dúvida D-R30:** até a decisão, o documento exportado leva o hash da foto, não a foto.
5. O administrador da sala (ou substituto) inclui o aparelho novo; os clientes conferem as duas assinaturas.
6. Um membro indicado envia o histórico dentro da retenção e dos intervalos de pertença [T24]. O aparelho novo recusa o que estiver fora.

| Erro | Tela e texto |
|---|---|
| Só um aprovador assina | Fica pendente: "Falta uma aprovação." |
| Aprovador é a mesma pessoa que cadastrou | Recusa |
| Sala concluída | "Esta sala não aceita novo aparelho." |
| Foto recusada (tremida, escura) | Pede nova foto sem guardar a anterior |
| Nenhum membro com histórico em cache | Mostra "sem histórico a recuperar" |

### 4.10 Missão de campo (mapa no app do comandante) [T25, T26, T27; PW10]
1. Na base, com conexão à central, o grupo da missão é montado; comandante e subcomandante registrados (GD-11).
2. Em campo, a **central fica congelada**: não altera membros nem entrega aos grupos da missão. O nó de campo só ordena.
3. Os agentes compartilham a posição pelo MLS na sala da missão. O mapa é exibido no app do comandante (e do substituto) [T25]. Servidor e painel não veem posição (ADR-0008 seção 3; RF88). **Limite (01g-G19):** a restrição é de interface; todo aparelho membro decifra e guarda a posição, e um aparelho de agente capturado desbloqueado a tem. Grupo MLS de comando: decisão da Fase 8 (D-R20).
4. Banda estreita [T27]: só texto curto; sem anexo; o app mostra o limite.
5. Em sala sob custódia, a posição é conteúdo e leva cópia de custódia (ADR-0008 item 3).
6. Remoção de membro em campo é commit do comandante da missão, via nó.
7. Nó de campo capturado: destruição imediata, exceção à GD-13 (GD-18).

| Erro | Tela e texto |
|---|---|
| Sem enlace | "Sem enlace." Mensagens aguardam |
| Posição sem permissão | A tela nunca pede localização do aparelho a terceiros; usa o que o core entrega |
| Comandante perde o aparelho | O substituto assume pela lista registrada (sem tela nova) |
| Pessoa não administradora tenta ver o mapa | Não há mapa para ela; não existe tela de "sem permissão" que confirme a existência |

### 4.11 Pergunta à IA e indicador com referência sem texto [T28] **[condicionada]**
1. Analista abre a sala de análise (escopo, prazo, revisão definidos pela chefia).
2. Digita a pergunta. O filtro de escopo vem **antes** da busca, no código, nunca por instrução ao modelo.
3. A resposta traz **referência**: "operação A, data X" (fictícios), **sem trecho de conversa**, com o rótulo fixo "indício para análise humana" e a fonte. Não há botão de ação automática.
4. Ler a origem exige ser membro da sala ou pedir pelo rito.
5. A pergunta fica registrada (quem, quando, escopo), cifrada para a chave de quórum de registros.

| Erro | Tela e texto |
|---|---|
| Pergunta fora do escopo | "Fora do escopo desta sala de análise." Sem revelar se há dado |
| Sala sem IA (por decisão da sala ou do grau) | Fica fora do escopo; a tela não lista |
| Nó de IA indisponível | "Análise indisponível." |
| Escopo vencido | Pede revisão à chefia |
| Sala de grau classificado | Recusa até o portão G3 (GD-9) |

### 4.12 Concessão de escopo de análise [T29] **[condicionada]**
1. O analista pede acesso ao nível 2 de uma sala. O pedido leva a CI e/ou a autorização do comandante (GD-15).
2. O comandante (ou quem a CI nomeia) vê o resumo do pedido: escopo, prazo, analista, base.
3. Concede por assinatura; o app **registra** a concessão. O painel só mostra a contagem (nível 1a).
4. A concessão tem prazo; a renovação é explícita (ADR-0009).
5. Os membros da sala veem o aviso de transparência [T30].

| Erro | Tela e texto |
|---|---|
| Sem CI nem autorização | Recusa: "Falta a base do pedido." |
| Analista sem credencial em todo o escopo | Recusa; o escopo não é recortado por conta própria |
| Concessão vencida | O acesso termina na hora; o app avisa |

---

## 5. Protótipo

### 5.1 Fidelidade por etapa

| Etapa | Fidelidade | Para quê | Fase | Saída |
|---|---|---|---|---|
| Papel | **Papel** (desenho em folha, cartões por tela) | Validar o mapa de telas e os fluxos 4.1 a 4.12 com 2 a 3 agentes, sem tecnologia | Fase 0 | Fotos das folhas (sem rosto), lista de ajustes |
| Navegável | **Navegável** (telas ligadas, sem lógica, dados fictícios) | Testes V09 de tempo e erro dos fluxos principais; teste dos ritos de coação e revinculação | Fase 0, antes de gravar Compose | Protótipo em arquivos de texto no repositório (ver 5.2) |
| Nativo | **Nativo** (Compose ligado ao core de teste, dados fictícios) | Repetir V09 com a interface real; medir tempo e bateria | Fase 1 (texto), Fase 2 (proteção), Fase 3 em diante | Interface de produção |
| iPhone | Papel e navegável apenas, no estudo de viabilidade; nativo só depois da ADR (D-R31) | Conhecer a diferença de plataforma sem escrever código de produto | Fase 0 | Relatório do estudo |
| Painel | Navegável em arquivos estáticos; nativo na Fase 1 mínimo | Cadastro e revogação com gestor | Fase 0 e 1 | Painel mínimo |

### 5.2 Ferramenta de prototipagem

Regra 10: nenhuma ferramenta em nuvem de terceiros, nenhuma telemetria, nenhum serviço de conta externo. O protótipo roda **local, offline**, e vive no repositório.

| Opção | Descrição | Vantagens | Limites |
|---|---|---|---|
| **A. HTML e CSS estáticos, sem script externo e sem fontes remotas** | Telas ligadas por links, abertas num navegador local ou no emulador | Sem dependência; versionável como texto; fácil de revisar em PR | Sem toque real; luva e pouca luz só por simulação |
| **B. Compose Preview e emulador Android** | Telas em Compose com dados fixos, sem o core | Já é código reaproveitável; toque real no aparelho de teste | Só Android; já é código (Fase 1) |
| **C. Ferramenta de desenho de código aberto, instalada localmente** (ex.: editor de vetores local) | Telas como arquivos de desenho | Rápido para o papel digital | Nova dependência: exige justificativa e versão fixada com hash (regra 8); verificar que não há telemetria |
| **D. Papel e cartões** | Folhas e recortes | Sem nenhuma dependência; ideal no papel | Não mede tempo com precisão |

Proposta: **D no papel, A no navegável, B no nativo.** O iPhone usa o mesmo protótipo A, com ajustes de plataforma. **DECISÃO PENDENTE: D-R42** (ferramenta do navegável e quem a mantém). Nenhuma opção é adotada até a decisão; qualquer dependência nova segue a regra 8.

### 5.3 Dados fictícios e versionamento
- **Dados fictícios:** codinomes de lista fixa ("Agente A", "Agente B", "Comandante X"), salas "Operação Alfa" e semelhantes, datas "DD/MM", mapas sem localização real. Nenhum nome, matrícula, local, foto de rosto, número de telefone ou texto de caso real. O protótipo traz um arquivo de dados fictícios único, revisado antes de cada sessão.
- **Fotos de teste:** imagens de objetos, nunca de pessoas, exceto o teste do termo com foto (4.9), em que o participante usa o **próprio rosto só durante a tarefa**, sem gravação: no protótipo navegável a foto não é guardada; no nativo vale o rito da seção 6.6.
- **Rosto e gabarito facial (01g-G18):** no papel e no navegável **não há reconhecimento real**: o avaliador simula o resultado e nenhum rosto é gravado. No nativo, o teste facial (U09) grava um **gabarito biométrico** do participante no aparelho de teste, e por isso só ocorre com o rito da seção 6.6. A taxa de erro do reconhecimento se mede em A76 e A78 (doc 07), com muitas tentativas; o V09 mede só aceitação e tempo.
- **Versionamento:** os protótipos entram no repositório (`docs/prototipo/` ou local a definir **[pendente]**), com tag por etapa e por rodada de teste (ex.: `proto-navegavel-r1`). Cada achado de usabilidade aponta o fluxo e a versão do protótipo.
- **Revisão:** mudança de tela passa por PR, como todo código; mudança no fluxo de coação ou perda revisada por quem cuida de segurança (Opus, doc 03 seção 6.1).

---

## 6. Protocolo de usabilidade (V09)

V09 é critério de entrada do piloto e a prioridade 1 (doc 03 seção 4.12). Este protocolo detalha a validação.

### 6.1 Participantes
| Projeto | Perfil | N mínimo | Observação |
|---|---|---|---|
| Android | Agente | 5 a 8 | Sem treino prévio no navegável; com o treino curto da norma de uso no nativo |
| Android | Comandante / subcomandante / administrador de operação | 2 a 3 | Criação de sala, conclusão, revinculação, missão |
| Android | Analista | 2 a 3 **[condicionada]** | Só quando a ADR-0008 for aceita |
| iPhone | Agente | 3 a 5 (navegável apenas) | Estudo de viabilidade; sem código de produto |
| Painel | Gestor | 2 a 3 | Cadastro, revogação, papéis, prazos |
| Painel | Auditor | 1 a 2 | Trilha e pedido pelo rito |
Total por projeto: 5 a 8 pessoas ao todo (agentes como núcleo), repetindo quem tiver papel duplo apenas uma vez por rodada. Ninguém pode ter participado do desenho da interface.

### 6.2 Tarefas por perfil, meta de tempo e taxa de erro
Metas **iniciais** (a calibrar na primeira rodada; **valores a confirmar**). Tempo é do início ao fim sem ajuda; erro é ação que leva a resultado diferente do pedido ou a pedido de ajuda.

| ID | Perfil | Tarefa | Fluxo | Meta de tempo | Erro tolerado |
|---|---|---|---|---|---|
| U01 | Agente | Ativar o aparelho com QR, definir PIN e PIN de coação, assinar o termo | 4.1 | 5 min | 1 em 10 |
| U02 | Agente | Abrir sala e enviar a primeira mensagem, sem instrução (RNF-F1) | 4.2 | 60 s | 0 sem ajuda |
| U03 | Agente | Enviar foto pela câmera do app | 4.3 | 45 s | 1 em 10 |
| U04 | Agente | Dizer, olhando para a tela, se a sala é sob custódia ou comum e qual o motivo | PI-3 | 10 s | 0 |
| U05 | Agente | Abrir o app com o PIN de coação (cenário simulado S4, seção 6.3) | 4.7 | 10 s | 0 (acerto do PIN certo) |
| U06 | Agente | Acionar o pânico em cenário simulado | 4.7 | 5 s | 0 em ação; acionamento acidental medido em tarefa neutra: no máximo 1 em 20 |
| U07 | Agente | Reconhecer a tela de trava e saber a quem procurar | 4.6 | 15 s | 0 |
| U08 | Agente | Fazer o termo de revinculação (assinatura e foto) | 4.9 | 3 min | 1 em 10 |
| U09 | Agente | Entrar com PIN e rosto num aparelho com luz baixa e luva fina | 4.8 | 10 s | Aceitação e tempo. No navegável, resultado simulado pelo avaliador, sem rosto gravado; no nativo, só com o rito da seção 6.6. Falso negativo medido em A78, não aqui |
| U10 | Comandante | Criar sala sob custódia com membros | 4.4 | 3 min | 1 em 10 |
| U11 | Comandante | Criar sala comum com motivo e segunda assinatura | 4.4 | 3 min | 1 em 10 |
| U12 | Comandante / subcomandante / solicitante | Concluir sala (três assinaturas) | 4.5 | 3 min por pessoa | 0 em assinatura duplicada |
| U13 | Comandante | Ver a posição dos agentes na missão e remover um aparelho | 4.10 | 90 s | 0 |
| U14 | Chefe da inteligência / superior da área | Aprovar uma revinculação com resumo legível | 4.9 | 2 min | 0 aprovação do pedido errado |
| U15 | Analista **[condicionada]** | Fazer a pergunta e achar a referência | 4.11 | 3 min | 0 em ler "texto" como fato |
| U16 | Comandante e analista **[condicionada]** | Conceder e pedir escopo de análise | 4.12 | 3 min | 0 |
| U17 | Gestor | Cadastrar usuário e aparelho | 4.1 | 4 min | 1 em 10 |
| U18 | Gestor | Revogar aparelho, ler o estado ("cache ilegível, segredo guardado") | 4.6 | 2 min | 0 em interpretar como "destruído" |
| U19 | Gestor | Atribuir papel com segunda pessoa e ver recusa de combinação proibida | ADR-0009 | 3 min | 0 |
| U20 | Gestor | Ler KPIs nível 1a e dizer o que **não** aparece (sem nome, sala ou operação) | 3.2 | 2 min | 0 |

### 6.3 Cenários de estresse
Aplicados nas tarefas U02, U05, U06, U07 e U09 e, quando fizer sentido, em U03.
| ID | Cenário | Como simular |
|---|---|---|
| S1 | **Pouca luz** | Sala escura ou lente com filtro; brilho no mínimo do aparelho; modo de campo (PI-6) |
| S2 | **Luva** | Luva fina e luva tática grossa; mede erro de toque (PI-7) |
| S3 | **Pressa** | Tarefa com cronômetro visível e interrupção do avaliador a cada 15 s |
| S4 | **Coação simulada** | O avaliador, com voz neutra e sem contato físico, pede o aparelho e "exige o PIN"; o participante usa o PIN de coação. Sem violência, sem ameaça real, **sem gravação de rosto**, com aviso prévio no termo de participação e direito de parar na hora |
| S5 | **Cansaço** | Tarefa no fim de uma sessão longa ou depois de exercício físico leve **[opcional]** |
| S6 | **Uma mão** | Tarefas com a mão não dominante ocupada |

### 6.4 Métricas
| Métrica | Como mede |
|---|---|
| **Sucesso** | Tarefa concluída sem ajuda, em porcentagem |
| **Tempo** | Cronômetro por tarefa, mediana e pior caso |
| **Erros** | Contagem e tipo (toque errado, termo mal lido, PIN trocado, marca de custódia lida errado) |
| **Acionamentos acidentais** | Pânico e PIN de coação fora do cenário |
| **Compreensão** | Perguntas curtas depois da tarefa: "o que a tela disse?" (marca, trava, revogação) |
| **SUS** (System Usability Scale), por projeto | Questionário de 10 itens ao fim de cada rodada |
| **Desvio para canal paralelo** | Pergunta direta: "você usaria outro aplicativo para isso?", com linha de base e uso real no piloto (seção 6.8). Mede R120 |
| **Observação do avaliador** | Hesitação, atalhos, anotações em papel |

### 6.5 Critérios de aprovação por projeto
Valores iniciais, a confirmar na primeira rodada **(D-R43)**. Com 5 a 8 pessoas, "0 erros", "100%" e "1 em 20" não medem taxa: as metas são **formativas**, com número fixo de tentativas por pessoa (seção 6.8).
| Projeto | Critério |
|---|---|
| **Android** | Sucesso média de 90% ou mais nas tarefas de agente; U02 sem ajuda em 100% dos participantes; U04, U05 e U07 com 0 erros; acionamento acidental de pânico dentro do limite; SUS maior ou igual a 70 (referência comum; ajustar); pelo menos 80% dos agentes dizem que **não** usariam outro aplicativo; nenhum achado **crítico** aberto |
| **iPhone** | Etapa navegável: mesmos fluxos, SUS comparável ao Android; lista de diferenças de plataforma e do que a Apple impõe (notificação, captura, atestação) entregue ao estudo de viabilidade. **Não aprova produção**; aprova só a continuação do estudo |
| **Painel** | Gestores concluem U17 a U19 com 0 erro em U18 e U19; U20 com 0 erro; SUS maior ou igual a 70 |
| **Coação (todos)** | Teste cego com **observador** independente em escolha forçada entre duas sessões (normal e coação), várias rodadas por observador: acerto perto do acaso; desvio de tempo medido por instrumento dentro da margem (PI-4; D-R47); uma rodada por perfil (urbano e missão) |
| **Reconhecimento facial** | Aceitação e tempo em S1 e S2 (U09). O falso negativo e o falso positivo se medem em A76 e A78 (doc 07): se acima do limite **(D-R40)**, o recurso não entra no gatilho de destruição e fica só como sinal auxiliar (nunca decide sozinho) |
Achado **crítico:** leva a perda de dado, trava indevida do agente, exposição de dado ou desvio para canal comercial. Achado **alto:** erro frequente ou confusão em tela de segurança. **Médio e baixo:** conforto e texto.

### 6.6 Ética e sigilo da sessão
- **Termo de participação** (anexo próprio **[a criar]**): objetivo, tarefas, direito de parar a qualquer momento, o que é registrado, quem vê, prazo de guarda.
- **Participação voluntária.** Nenhuma consequência funcional ou disciplinar pelo desempenho; resultado individual não vai à chefia (regra 11).
- **Sem gravação de rosto, de voz ou de tela com identificação.** Só anotações do avaliador e cronômetro, por código do participante ("PART-01").
- **Rito do rosto real (01g-G18; regra 10; LGPD art. 11):** o teste facial no nativo só ocorre com rito escrito antes da sessão: consentimento específico para o teste facial, separado do termo de participação; gabarito só no aparelho de teste, cifrado como no produto; gabarito **zerado ao fim da sessão** e a remoção conferida e registrada por duas pessoas; parecer do encarregado. Sem esse rito, o teste facial fica no navegável, com resultado simulado (D-R44).
- **Recrutamento por canal neutro:** convite da chefia é pressão; o convite vem do responsável pelo teste, e a recusa não é comunicada à chefia.
- **Sem dado real.** Só dados fictícios; o aparelho de teste é do projeto, sem conta real, sem rede da instituição.
- **Sem terceiros:** nada de ferramenta de pesquisa em nuvem nem formulário externo; SUS em papel ou arquivo local.
- **Participantes de perfis sensíveis** (comandante, chefia): sessão individual e fechada; o conteúdo da sessão não tem dado real, mas o perfil indica a função; registrar só o papel, não o nome.
- **Coação simulada** (S4): aviso prévio, sem contato físico, parada imediata a qualquer sinal, conversa final ("debriefing").
- **Guarda dos registros:** anotações com código, guardadas pelo prazo do piloto e depois expurgadas; sem vínculo com nome. A tabela código-participante fica só com o responsável do teste, separada, e some ao fim.
- **Parecer** do jurídico e do encarregado antes da primeira sessão com pessoas da instituição.

### 6.7 Registro de achados e ciclo de correção
1. Cada sessão gera um relatório curto: fluxo, tarefa, métrica, observação, versão do protótipo.
2. Cada achado recebe ID (US-001...), **gravidade** (crítico, alto, médio, baixo), tela e fluxo ligados, e dono.
3. Achado que muda fluxo de segurança (coação, perda, revinculação, custódia) passa por revisão de desenho (Fable) e, se contrariar um ADR, vira ADR antes do código (CLAUDE.md).
4. Correção no protótipo, nova rodada com **pelo menos metade dos participantes novos** e todos os crítico/alto reexaminados.
5. Rodadas por etapa: pelo menos uma no navegável e uma no nativo por projeto, mais uma antes do piloto (V09).
6. Registro das decisões na pasta do protótipo, com versão e data; resumo ao comitê.

### 6.8 Complementos do protocolo (01g-G17)
1. **Desvio para canal paralelo (R120):** além da pergunta, linha de base com a mesma tarefa num mensageiro comum (tempo e toques) e medida de uso real no piloto (norma de uso e relato anônimo). Preferência declarada diante do avaliador vale pouco.
2. **Amostra:** com 5 a 8 pessoas, as metas da seção 6.2 são formativas; número fixo de tentativas por pessoa e por tarefa, registrado.
3. **Teste cego da isca:** quem não pode distinguir é o **observador**, não o participante: escolha forçada entre duas sessões gravadas só em tela de protótipo, sem rosto (normal e coação), várias rodadas por observador, acerto perto do acaso; tempo medido por instrumento. Uma rodada por perfil.
4. **Retenção:** repetir U05 e U06 depois de 2 a 4 semanas sem uso.
5. **Recuperação:** tarefas do que vem depois: voltar de pânico acidental, de trava nível 3, de lease vencido e da isca por falso negativo (4.8). Mede o tempo até voltar a trabalhar.
6. **Tarefas acrescentadas:** código de segurança entre membros (T21, RF04); perceber mensagem urgente com notificação genérica ou nenhuma (R125); mensagem em banda estreita (T27); o que o membro entende ao ver uma concessão (T30); compreensão dos itens 2, 13 e 14 do termo.
7. **Estresse de campo:** tela molhada, sol direto, veículo em movimento, capacete e pintura no rosto (perfil Exército).
8. **Ética:** recrutamento por canal neutro e rito do rosto real (seção 6.6).
9. **Por perfil:** o PIN de coação e a face têm efeitos diferentes no perfil missão (4.7, 4.8); uma rodada por perfil.

---

## 7. Riscos de usabilidade que viram risco de segurança

Registrados no doc 04 (regra 11), com a fusão proposta no `01g` seção 6. Os IDs antigos R-U1 a R-U17 ficam só como referência histórica.

| ID (antigo) | Risco | Como vira risco de segurança | Onde está |
|---|---|---|---|
| **R120** (R-U1) | Rito complexo leva ao uso de mensageiro comercial | Conteúdo sai do sistema (vazamento, fora da regra 3 e da custódia) | doc 04 seção 5.5 |
| **R55** (R-U2, R-U3) | PIN de coação confundido com o normal ou digitado sem querer; pânico acionado sem querer | Alerta falso, desgaste do plantão; na coação real, o PIN errado abre a conta real; no perfil missão, destruição por engano | doc 07 seção 10 (impacto Alto no perfil missão) |
| **R67** (R-U4 e parte do R-U5) | Falso negativo facial bloqueia ou destrói o aparelho do dono | Perda de acesso na hora crítica; atalho; destruição do aparelho do legítimo | doc 07 seção 10 |
| **R116** (parte do R-U5) | Falso positivo facial: o adversário passa | Entrada sem disparar a escada | doc 07 seção 10 |
| **R121** (R-U6) | Tela de coação distinguível | Adversário percebe a isca e reage com violência | doc 04 seção 5.5 |
| **R122** (R-U7) | Marca de custódia mal lida | Agente acha que a sala é comum e fala algo sensível, ou o contrário; falha de transparência (doc 06) | doc 04 seção 5.5 |
| **R69** (R-U8) | Aprovador assina a revinculação sem ler | Reposição do aparelho do adversário | doc 07 seção 10 (controle novo) |
| **R123** (R-U9) | Três assinaturas da conclusão viram "clique e siga" | Conclusão indevida; expurgo irreversível da sala comum | doc 04 seção 5.5 |
| **R106** (R-U10) | Analista lê a referência da IA como fato | Decisão contra pessoa baseada em indício (regra 11) | doc 04 seção 5.2 |
| **R124** (R-U11) | Gestor lê "revogado" como "destruído" | Promessa falsa à instituição, ou expectativa de recuperação errada | doc 04 seção 5.5 |
| **R125** (R-U12) | Notificação genérica excessiva ou ausente | Agente deixa de ver ordem urgente, ou desliga o app por barulho | doc 04 seção 5.5 |
| **R126** (R-U13) | Campo: tela clara ou som revela a posição | Detecção visual em operação noturna | doc 04 seção 5.5 |
| **R127** (R-U14) | Comandante vê a posição e a repassa por foto da tela ou fala | A posição sai do MLS | doc 04 seção 5.5 |
| **R102** (R-U15) | Diferença entre Android e iPhone gera erro | Agente confunde ritos; divergência de segurança | doc 04 seção 5.2 |
| **R105** (R-U16) | Painel parece mostrar mais do que mostra, ou o gestor pede mais | Pressão por dado nominal ou derivado no painel (regra 3) | doc 04 seção 5.2 |
| **R128** (R-U17) | Sessão de teste de usabilidade vaza dado real ou biometria real | Dado real em protótipo, anotação ou gabarito de teste | doc 04 seção 5.5 |

---

## 8. Entregas e critérios de aceite por fase

| Fase | Entrega | Critério de aceite |
|---|---|---|
| **0** | Protótipo em papel de T01 a T08, T16 a T24 e PW01 a PW07; navegável dos fluxos 4.1, 4.2, 4.6, 4.7, 4.9 | Uma rodada V09 no navegável com 5 agentes e 2 comandantes; relatório de achados; critérios dos fluxos de coação e revinculação sem achado crítico aberto |
| **0** | Estudo de viabilidade do iPhone: navegável com a lista de diferenças de plataforma | Relatório sobre APNs, App Attest, captura de tela e MDM; **sem código de produto**; insumo para a ADR própria (D-R31) |
| **0** | Decisão da ferramenta de prototipagem (seção 5.2) | Decidido e registrado; nenhuma dependência sem regra 8 |
| **1** | Interface nativa T01 a T07, T21 e painel PW01 a PW03, PW06, PW07 | Operador novo envia a primeira mensagem sem instrução (doc 03 Fase 1; U02); U17 e U19 aprovadas; SUS parcial |
| **2** | T16 a T20 (trava, pânico, isca, verificação facial se adotada), PW04 e PW05 | U05 a U07 e U09 aprovadas; teste cego da isca (6.5); relatório de falso negativo facial |
| **3** | T10 a T12 | U03 aprovada; captura de tela bloqueada |
| **4** | T09, T15, PW09, PW11, PW12 | U10 a U12 aprovadas; painel de prazos de custódia sem identificar sala |
| **4** | T22 a T24 (revinculação; doc 07 seção 14) | U08 e U14 aprovadas; D-R30 decidida |
| **5, 6** | T13 e T14 | Rótulo de apoio à decisão sempre visível |
| **7 (piloto)** | Rodada final V09 completa com a interface nativa, dados fictícios, em aparelho gerenciado | Critérios da seção 6.5 atendidos; achados críticos e altos fechados; **V09 aprovada** (entrada do piloto) |
| **8** | T25 a T27, PW10 | U13 aprovada; teste de campo com luva e pouca luz; mapa só na interface do comandante e do substituto, com o limite declarado (01g-G19) ou grupo de comando (D-R20) |
| **Depois da Fase 6** [condicionada] | T28 a T30, PW08 completo | U15, U16 e U20 aprovadas; painel só com o nível 1a (o 1b não existe, GD-15) |
| **iPhone nativo** | Depois da ADR própria | Mesmo conjunto T01 a T30; testes de equivalência com o Android |

### Decisões pendentes (registradas no doc 04 seção 4)

| Ref. | Decisão | Quem decide |
|---|---|---|
| D-R42 (antes U-1) | Ferramenta do protótipo navegável e local no repositório (5.2) | Responsável pelo produto, com segurança |
| D-R43 (antes U-2) | Metas de tempo, erro e SUS definitivas (6.2 e 6.5), formativas com amostra pequena | Responsável pelo produto, após a rodada 1 |
| D-R47 (parte da U-3) | Margem aceitável de diferença de tempo entre PIN normal e PIN de coação (PI-4) | Segurança, com o jurídico |
| D-R40 (parte da U-3) | Limite de falso negativo facial, dentro da escada por perfil (doc 07) | Responsável pelo produto + segurança + jurídico |
| D-R44 (antes U-4) | Termo de participação, responsável pelo teste e rito do rosto real (6.6) | Jurídico e encarregado |
| D-R20 (antes U-5) | Quem vê o mapa da missão além do comandante e do substituto, e se a restrição vira grupo MLS de comando (01g-G19) | Responsável pelo produto (pergunta 9 do 01c) |
| D-R45 (antes U-6) | Política de notificação por perfil (genérica ou nenhuma) | Segurança, com a chefia |
| D-R46 (antes U-7) | Texto aprovado das telas de trava, revogação e marca de custódia | Responsável pelo produto e jurídico |
| D-R30, D-R31, D-R36 | Já existentes (foto do termo, iPhone, segunda assinatura da sala comum) | Ver docs 07 e ADR-0007 |

---

## 9. Pontos para validação

1. **Fases das telas:** T22 a T24 (revinculação) ficam na Fase 4, como no doc 07 seção 14; T09 (ciclo) também na Fase 4.
2. **Metas e critérios numéricos** (60 s, SUS 70, limites de erro) são propostas iniciais; confirmar ou trocar.
3. **Captura de tela no iPhone:** aceitar que a garantia pode ser menor que no Android, ou impedir o iPhone na v1 por esse motivo?
4. **Notificação:** texto genérico ou nenhuma notificação por padrão?
5. **Verificação facial:** o protótipo cobre T19 mesmo antes da decisão de D-R8? Proposta: sim, no navegável, para medir aceitação, sem código e sem reconhecimento real (seção 6.2, U09).
6. **Mapa de missão:** só o comandante e o substituto, ou outros papéis (pergunta 9 do 01c)?
7. **Teste de coação simulada (S4):** o jurídico e a chefia aceitam o desenho e o aviso prévio?
8. **Painel de prazos de custódia (PW09):** a faixa de prazo sem identificar sala é suficiente para o gestor agir? (O nível 1b não existe, GD-15; o que faltar vai ao app por MLS.)
9. **Local do protótipo no repositório** e quem o mantém (nenhum dos componentes tem CLAUDE.md para isso hoje).
10. **Referências a este doc:** feitas na quarta rodada no doc 03 seção 4.12 (V09), no doc 04 (riscos R120 a R128, decisões D-R42 a D-R47) e no `admin-panel/CLAUDE.md` (telas PW). Falta a lista de leitura do `CLAUDE.md` raiz, que depende de autorização do responsável.
