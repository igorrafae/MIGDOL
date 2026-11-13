# Revisão da terceira rodada — Fase 0, etapa 1g

**Modelo:** Claude Fable 5.2 (pipeline da Fase 0, revisão da terceira rodada da etapa 1b). Nenhum arquivo foi alterado além deste; nada de commit; nada enviado para fora.
**Entradas:** `CLAUDE.md` (regras 2, 4 e 14), `GD-respostas.md` (GD-15 a GD-20 prevalecem), `01e-revisao-rodada2.md`, `01f-correcoes-rodada3.md`, ADRs 0003 a 0010 e README, doc 11 (Sonnet, primeira revisão), `git diff` dos docs 01 a 10, do termo, de `app/CLAUDE.md`, `admin-panel/CLAUDE.md` e do `PIPELINE.md`.
**Limites:** leitura de texto; nada foi executado. O diff contra o último commit junta as três rodadas da 1b; nos docs 01 a 10 li as seções que o `01f` aponta como alteradas e as que um achado exigiu, não o texto inteiro. O que afirmo sobre MLS, Android Keystore, HSM e biometria vem de conhecimento geral: **[verificar no spike]** onde depende da biblioteca ou do hardware. Todo ponto jurídico é hipótese para o jurídico: **[conferir]**.
**Convenção:** achados desta revisão são `01g-Gnn`. Os anteriores seguem como `01e-Enn` e `01c-Cnn`.

## 1. Veredito

| Peça | Veredito | Condição |
|---|---|---|
| **Regra 2** | **Aprovar** (sem mudança nesta rodada) | 01e-E15 tratado no doc 02 e no critério do spike |
| **Regra 4** (nova) | **Aprovar** | Fiel à ADR-0007 item 5. Ressalva baixa: a lista de APIs se lê como exaustiva (G24) |
| **Regra 14** | **Aprovar** (sem mudança) | A ADR-0010 não a enfraquece na divisão 1a/1b/2. Dois pontos a decidir: N3 (redação do 01c-C23) e a leitura de R3-2. O achado G01 mostra que a ADR-0010 ainda não diz como cumpre a regra entre sessões |
| **ADR-0003** (custódia) | **Aprovar com ressalvas** | 01e-E06, E09b, E10a, E17 e E21e fechados. A conferência automática das ordens de preservação na cerimônia de eliminação precisa ser descrita (G01c). 01c-C06 e C11 seguem com o HSM |
| **ADR-0004** (campo) | **Aprovar com ressalvas** | 01e-E07 resolvido por decisão do responsável; a exceção à GD-13 está bem escrita. G20 (falha de boot) e 01c-C17 antes de aceitar. Fase 8 |
| **ADR-0005** (commits) | **Aprovar com ressalvas** | 01e-E02 e E03 fechados no texto. G13 entra no critério do spike antes da etapa 3 |
| **ADR-0006** (lease) | **Aprovar com ressalvas** | 01e-E05, E06, E08, E12 e E13 fechados. Faltam só decisões (Q5, D-R35) |
| **ADR-0007** (plataformas) | **Aceita (GD-19): coerente** | Regra 4, README e doc 02 seção 11 batem |
| **ADR-0008** (painel e IA de análise) | **Aprovar com ressalvas, como proposta para depois da v1** | **01e-E01 fechado no texto** (seção 3). Ressalva baixa G23; D-R41 |
| **ADR-0009** (RBAC + ABAC) | **Aprovar com ressalvas** | 01e-E10 fechado; E18 em parte (D-R33). A lista assinada de papéis precisa do vínculo que G05 pede. Decidir antes da Fase 1 |
| **ADR-0010** (investigação sigilosa) | **Não aceitar como está** | Direção correta e fiel às GD-16 e GD-17. **Dois altos (G01, G02)** e oito médios (G03 a G10). Fase 4: não trava G0 a G3 |
| **Doc 11** (protótipo e V09) | **Aprovar como rascunho, com ressalvas** | Coerente no essencial. Corrigir G14 a G19 e G25 antes da primeira sessão com pessoas |

Contagem dos achados novos: **0 críticos, 2 altos, 17 médios, 7 baixos** (26). Os dois altos estão na ADR-0010.

## 2. ADR-0010 — investigação sigilosa

### 2.1 O que está certo
- **Sem gravação oculta:** sem membro oculto, sem troca de marca, sem uso do nó de IA, sem "modo investigação", sem reter conteúdo nos aparelhos. É a leitura correta da GD-17 e da ADR-02.
- **Divisão 1a / 1b / 2:** correta e mais restritiva que a GD-16. Registro de visualização, registro de acessos, evidência, matrícula e segredo de lease são dado cifrado sob quórum: sem a base da regra 14 não se abrem. Conteúdo só com ordem judicial.
- **Assinatura presencial na estação, com chaves da lista assinada de papéis:** coerente com "sem aprovação remota".
- **Regra 11:** saída estruturada, duas revisões humanas, nenhuma medida automática, indício de conteúdo que não eleva o grau sozinho.

### 2.2 As cinco perguntas do pedido

| Pergunta | Resposta curta |
|---|---|
| É sólida contra "um detentor do quórum é o vazador"? | **Só quando o detentor suspeito é conhecido.** Com suspeito declarado, o impedimento por chave funciona. Com "um dos detentores, não se sabe qual", qualquer quórum tem boa chance de incluir o vazador, a segunda autorização proposta para a corregedoria é justamente um detentor, e o iniciador nunca é conferido (G03, G04). O limite não está declarado |
| O impedimento é implementável sem revelar a investigação? | **Em parte.** A conferência por chave na estação é implementável. Faltam: como a estação lê a lista de impedidos numa segunda sessão sem abrir o registro selado (G01); o vínculo entre a chave do papel e a identidade nas salas, sem o qual o impedimento por posição passa em branco (G05); e os rastros laterais que mostram ao impedido que houve cerimônia sem ele (G06) |
| O registro selado garante prestação de contas? | **Nos níveis 1b e 2, razoavelmente** (o HSM obriga o quórum). **No nível 1a, não:** os metadados estão em claro, duas pessoas fazem tudo, renovam sozinhas, e a revisão por quórum é semestral e ainda proposta (G02) |
| Os dois níveis são coerentes com a regra 14? | **Sim na tabela da seção 4.** A incoerência está fora dela: o registro e os retratos preservados são dado cifrado para quórum, e a ADR os lê em sessões que diz não precisarem de quórum (G01, G07) |
| Abuso, vazamento da investigação, nulidade | G02 (abuso), G06 e G07 (vazamento), G08 e G09 (indício forjado e prova) |

### 2.3 Respostas às pendências do Opus (01f seção 6, item 1)
- **(a) Três fontes de impedimento:** boas. Faltam o iniciador (G03) e o vínculo chave ↔ identidade (G05).
- **(b) Divisão e leitura da regra 14:** correta. Recomendo confirmar R3-1.
- **(c) Retrato do backup de rotina:** não basta dizer "de rotina". Quem retira a mídia e o registro de acesso ao backup são visíveis ao administrador de infraestrutura (G07).
- **(d) Registro de tamanho fixo:** o tamanho não distingue, mas a entrada fica órfã na trilha, e isso distingue (G06).
- **(e) Modelo de inferência na estação:** pode, sem virar "terceiro lugar", se for a mesma classe de estação e a mesma sessão de cerimônia. Recomendo máquina separada da que fala com o HSM (G10).

## 3. 01e-E01 e o nível 2 da ADR-0008

**Fechado no texto pela GD-15.** Conferi cada ponto do achado:

| Ponto do 01e-E01 | Como ficou |
|---|---|
| Sem ato da sala de origem | Concessão assinada pelo administrador **dentro do grupo**, visível aos membros, retirável; o nó de IA só cruza sala cuja concessão ele viu no grupo |
| Sala comum alimentando acervo de custódia | Regra de tipo (A97) |
| Expurgo da origem não alcançava a sala de análise | Expurgo em cascata (A98) |
| Analista pergunta o que quiser | Saída estruturada, registro e limite de perguntas, revisão amostral |
| Chefia sozinha decide o escopo | Concessão da origem; aprovação dupla para criar a sala (proposta, D-R41) |

Regra 3 intacta: o painel só tem o nível 1a, e o 1b não existe. Restam ajustes baixos (G23) e a D-R41. R104 em Médio até o teste de injeção e R98 em Médio até a Fase 1 estão certos (pendência 7 do Opus).

## 4. Achados novos

| ID | Sev. | Onde | Problema | Proposta |
|---|---|---|---|---|
| 01g-G01 | **Alta** | ADR-0010 seções 3, 4, 5, 8; ADR-0003 seções 4 e 7; regra 14 | **A investigação não tem como continuar entre sessões sem abrir o próprio registro, e abri-lo é acesso sob a regra 14.** (a) A lista de impedidos "existe só dentro do registro selado", e a estação tem memória volátil. Para "conferir antes de cada ato" numa segunda sessão, ela precisa decifrar o registro: quórum presencial e base registrada. A tabela da seção 4 diz que o nível 1a não exige quórum. Como está, ou o 1a exige quórum a partir da segunda sessão, ou existe uma cópia de trabalho fora do quórum (não descrita), ou o impedimento não é conferido. (b) A seção 5 manda abrir o registro "com base registrada (regra 14)" para a revisão periódica; a revisão de controle não tem ordem judicial nem processo administrativo. (c) Em toda cerimônia de eliminação "a estação abre as ordens de preservação vigentes": é uso da chave da série de investigação por um quórum de rotina, que pode incluir o impedido | Decidir e escrever uma das duas: **(1, recomendada)** toda sessão da investigação, em qualquer nível, é cerimônia com M detentores não impedidos; os níveis diferem só pela **base** (1a: CI e segunda autorização; 1b: mais a base da regra 14; 2: mais ordem judicial). Resolve também G02. Custo: mais gente sabe. **(2)** Cópia de trabalho do pedido e da lista de impedidos cifrada para as chaves pessoais do iniciador e da segunda autorização (HPKE da pilha), com declaração de que não é "registro" da regra 14; exige decisão do responsável, porque toca a regra. Em (b): a revisão de controle entra como base própria na redação da regra 14 ou corre em processo administrativo de controle **[conferir]**. Em (c): descrever na ADR-0003 seção 7 a conferência automática, sem exibição e registrada |
| 01g-G02 | **Alta** | ADR-0010 seções 2, 4, 6, 10; R112; D-R38 | **No nível 1a o controle do rito não é independente de quem o iniciou.** (a) "Sessão com duas pessoas não impedidas" e "ao menos um revisor distinto do iniciador" permitem que iniciador e segunda autorização sejam as duas pessoas da sessão e os dois revisores. (b) A renovação de 90 dias é dada pelas mesmas duas, sem limite de vezes. (c) A revisão por quórum é semestral e proposta. (d) O investigado pode nunca ser avisado, e o registro é eliminado no prazo da D-R38: a prova de um abuso some. (e) Nada obriga a criar o registro: os metadados do 1a estão em claro no backup, e quem tem a mídia faz a correlação em qualquer máquina. O controle real é o acesso ao backup, que a ADR não trata. É vigilância de metadados de agentes por duas pessoas, em segredo (regra 13) | (1) Quórum não impedido em toda sessão (G01, opção 1), ou no mínimo na abertura e em cada renovação. (2) Um revisor que não seja o iniciador nem a segunda autorização. (3) Limite de renovações (proposta: duas); depois, decisão do quórum ou do MP. (4) Encerramento obrigatório com ato registrado e comunicação a órgão de controle fora da cadeia que conduziu **[conferir qual]**. (5) O registro não é eliminado antes da revisão final nem antes do prazo de responsabilização por abuso **[conferir]**. (6) Aviso ao investigado no arquivamento como padrão; adiar exige decisão escrita e revista. (7) O retrato só sai do backup por ato de liberação com duas pessoas, registrado na trilha (G07) |
| 01g-G03 | Média | ADR-0010 seção 3; doc 04 seção 6.2 | A estação confere as chaves da segunda autorização, dos cartões e dos revisores, **não a do iniciador**. Quem é membro da sala do escopo pode iniciar, operar a sessão, rever e receber o relatório. O próprio vazador pode iniciar, definir escopo e impedidos e conduzir | O impedimento por posição vale para o iniciador em todos os atos seguintes: quem é membro ou administrador do escopo **pode denunciar e iniciar**, mas não opera a estação, não revê e não recebe o relatório (doc 05, princípio 5). A segunda autorização confirma escopo e lista de impedidos como ato próprio |
| 01g-G04 | Média | ADR-0010 seções 2 e 3; R115; D-R24, D-R37 | (a) **Detentor suspeito não identificado:** não há como impedir todos; com 3 de 5, o quórum inclui o vazador na maioria dos casos; e a segunda autorização proposta para a corregedoria é "um detentor não impedido". O limite não está declarado. (b) "Reemissão de cartões a detentores de reserva, em cerimônia com M de N livres" é circular: se há M livres, não precisa de reserva; se não há, não há cerimônia. (c) Se os detentores são comandantes membros de salas (D-R24 aberta), o impedimento por posição esvazia o quórum com frequência | (a) Declarar o limite: nesse caso só o nível 1a guarda sigilo perante os detentores, e a segunda autorização vem de **fora da classe dos detentores** (D-R37; proposta: chefe do jurídico ou autoridade do escalão superior). (b) Reservas com cartão **emitido de antemão** na cerimônia original (N maior), guardado lacrado. (c) Reforça a recomendação para a D-R24: detentor de cartão não é membro de sala operacional |
| 01g-G05 | Média | ADR-0010 seção 3 item 2; ADR-0009 seção 3; doc 05 seção 3.4 | O impedimento por posição compara chaves da lista assinada de papéis com membros das salas. As salas conhecem aparelho e codinome; a ligação com a pessoa é o pseudonimato, que é classe sob quórum (nível 1b). A identidade dos detentores fica fora do diretório. Sem um vínculo, a estação não liga a chave do cartão ao membro, e a conferência passa em branco | A lista assinada de papéis (offline ou selada) leva, para cada chave, o identificador opaco do usuário no diretório. A estação compara por esse identificador, sem quebrar pseudonimato de terceiros. Participante sem vínculo conferível é tratado como impedido (falha fechada). Teste em A85 |
| 01g-G06 | Média | ADR-0010 seção 5; R113; doc 11 P09 | Rastros que revelam a investigação e que a ADR não trata: (a) a entrada de tamanho fixo fica **órfã**: revogação, lote e auditoria amostral têm evento legível ao lado; quem lê a trilha conta as órfãs e vê as datas; (b) log e contadores do HSM, vistos na cerimônia seguinte; (c) registro de acesso à sala, gravação e testemunhas (ADR-0003); (d) o doc 11 põe no painel "suspensões de prazo (preservação) só como contagem" | (a) Entradas de cobertura em toda cerimônia de rotina, ou declarar o limite no R113. (b) e (c) Sessões da investigação junto das cerimônias de rotina, com operadores e testemunhas não impedidos; gravação selada para a série de investigação. (d) O painel não mostra preservação selada (G16). Declarar que, com N pequeno, a ausência de um detentor é percebida |
| 01g-G07 | Média | ADR-0010 seções 4, 6 e 8; doc 06 seções 3 e 8.2; termo item 14 | (a) "Retrato do backup de rotina" não esconde nada se quem retira a mídia é o administrador investigado, ou se o acesso ao backup fica registrado. (b) "Retrato dos metadados **e dos cifrados** do escopo" não exclui os cifrados MLS de sala comum. (c) Os retratos "para frente" são cifrados para a chave de quórum: abri-los é regra 14 (quórum e base), e a tabela do 1a não diz. (d) Horários e volume ficam hoje "até a entrega ou 72 h" (doc 06 seção 8.2); o retrato a cada backup cria uma série longa que não existia. É coleta nova. (e) A destruição antecipada de segredo de lease selado, com duas autorizações, não passa pela estação: o `custody` apaga registro preservado | (a) Retrato por **ato de liberação do backup** com duas pessoas, registrado, casado com o teste trimestral de restauração (doc 05 seção 8). (b) Sala comum: só metadados, nunca cifrado MLS. (c) Dizer que retrato preservado só abre com quórum e base da regra 14. (d) Declarar no inventário e no termo que a preservação pode guardar metadados além do prazo normal. (e) Dizer que a cópia no retrato cobre esse caso |
| 01g-G08 | Média | ADR-0010 seções 6 e 9; R114 | O retrato vem do banco do servidor, que é hostil e pode ser administrado pelo investigado. Horários, volumes e estados **sem assinatura** podem ter sido forjados ou apagados para apontar outra pessoa. O relatório não distingue o que é atestado do que é só registro do servidor | Cada sinal rotulado como **atestado** (declaração de membros assinada, cadeia de hash da auditoria, evento assinado pelo aparelho) ou **registro do servidor sem assinatura**. Este não eleva o grau sozinho enquanto administradores de infraestrutura não estiverem afastados da suspeita. Comparar retratos de backups de datas diferentes |
| 01g-G09 | Média | ADR-0010 seções 2, 4, 9, 10; doc 05 seção 3.6a | Riscos de nulidade **[conferir todos]**: (a) nível 1a por 90 dias renováveis só com CI, sem procedimento instaurado (sindicância, IPM); (b) "acesso ao relatório e ao trecho do registro que o fundamenta" pode ser menos que a defesa tem direito sobre o já documentado (Súmula Vinculante 14); (c) indício de conteúdo "não lido por pessoa" não é verificável sem reprodução do método; (d) a ordem judicial do nível 2 precisa cobrir análise automatizada sem leitura; (e) cadeia de custódia do retrato (CPP arts. 158-A e seguintes); (f) metadados de comunicação sem ordem judicial, mesmo em sistema e aparelho da instituição | (a) O pedido cita o ato que instaurou o procedimento, ou o jurídico confirma que a CI basta. (b) Texto neutro: "acesso conforme decidir a autoridade do procedimento, no mínimo ao que estiver documentado contra ele". (c) O registro guarda hash do modelo e do software, parâmetros, semente e hash do retrato, para um perito repetir. (d) Modelo de pedido ao juízo. (e) Hash do retrato na trilha no ato da liberação. (f) Parecer antes da Fase 4 |
| 01g-G10 | Média | ADR-0010 seção 6; ADR-0003 seções 5 e 6; doc 02 seção 1 | (a) Pôr GPU, pesos e pilha de inferência na estação da cerimônia aumenta a base confiável da máquina que recebe chaves do HSM, e ela processa texto que o adversário pode ter plantado. (b) Análise de anos de custódia dura horas ou dias: a ADR não limita a sessão nem diz quem fica presente. (c) Não diz **o que os detentores veem**: quórum que insere o cartão sem saber o quê não é controle; quórum que vê o escopo aumenta o vazamento | (a) Duas máquinas da mesma classe (mesmos requisitos da ADR-0003 seção 6): a da cerimônia decifra o escopo e passa por mídia só de leitura à de análise. Ajustar o doc 02 para "estações isoladas de cerimônia". Decidir depois do spike de inferência. (b) Duração máxima e presença mínima. (c) Mínimo que cada detentor vê antes do cartão: CI, base, nível, classes e quantidade de períodos, sem identificar salas; pode recusar |
| 01g-G11 | Média | doc 07 seções 7.1a e 7.2; R67, R116; D-R40 | (a) "Falsificação detectada" foi equiparada a "outra pessoa". A GD-18 diz "face errada". Vivacidade passiva em câmera frontal comum erra com pouca luz, reflexo e rosto pintado: no perfil missão, falha de vivacidade do dono mais **um** PIN errado destrói o aparelho. (b) Quem cobre a câmera cai em "inconclusivo" e entra só com o PIN (limite já declarado): contra adversário informado o reconhecimento não reforça o acesso; serve de armadilha para o desavisado, e o custo recai no dono. (c) Qualquer colega destrói o aparelho de outro em segundos (rosto dele e um PIN errado) | (a) Para o nível 4 conta só "rosto vivo, de boa qualidade, que não corresponde", em **duas capturas independentes**; falha de vivacidade vale como inconclusiva para destruição (segue nível 2 ou 3) até ser medida à parte. (b) Dizer na seção 7.2 que a função é detectar, não barrar, e fixar na D-R40 limite de erro perto de zero para o dono, em condição de campo (pintura, capacete, ferimento, noite). (c) Registrar em R67 |
| 01g-G12 | Média | doc 07 seções 7 e 7.1a; R55; doc 11 seção 4.7 | No perfil missão o PIN de coação aciona o nível 4 por baixo da isca. (a) PIN de coação digitado **por engano**, sob pressa, luva e escuro, destrói o aparelho em campo; o R55 segue Médio/Médio. (b) Ordem dos atos não descrita: alerta e autoexclusão precisam sair antes de zerar as chaves; sem rede, o alerta se perde. (c) A isca precisa existir fora do cache destruído, e o apagamento não pode mudar tempo nem aquecimento (doc 11 PI-4). (d) O doc 11 diz que o PIN de coação sem coação tem "o mesmo efeito (alerta)" e que o nível 2 é reversível: não vale por perfil | (a) Reavaliar o R55 com impacto Alto no perfil missão; medir o acionamento acidental sob estresse (V09) **antes** de ligar o nível 4 ao PIN de coação; resultado entra na D-R40. (b) Ordem: enfileirar alerta, enviar se houver enlace, zerar. (c) Teste de tempo em A79. (d) Corrigir o doc 11 (G15) |
| 01g-G13 | Média | ADR-0005 item 2c; doc 03 Fase 0 item 4 | O sinal "forte" (não decifra e traz hash diferente) é tão forjável quanto o fraco: o servidor, ou um **membro**, envia dois pacotes e obtém "forte repetido" e SEV1. O limite de alertas reduz a fadiga, mas não impede induzir "recriar a sala". Erro meu de classificação no 01e-E02 | Separar **bifurcação provada** (commit conflitante cuja assinatura de administrador registrado o cliente verifica, ou importação de cadeia) de **anomalia não autenticada** (qualquer quantidade de pacotes que não decifram). A segunda gera alerta "servidor ou membro anômalo", sem recomendar recriar a sala. Critério do spike: a biblioteca deixa verificar um commit de época já aplicada? **[verificar no spike]** |
| 01g-G14 | Média | doc 11 T02 e fluxo 4.1; regra 4; doc 06 seção 7 | A tela T02 assina o termo "pelo gov.br" no app. Isso exige rede do aparelho para um serviço externo, fora do core (regra 4), num aparelho gerenciado, e CPF digitado nele. Os docs 01, 02 e 06 dizem "na presença do gestor" e validação "no servidor ou na estação do gestor", sem dizer onde se assina | A assinatura no gov.br ocorre **fora do app**, na estação do gestor. T02 mostra o texto, a versão e o estado que o core entrega. Registrar no doc 06 seção 7 |
| 01g-G15 | Média | doc 11 fluxos 4.7 e 4.8; doc 07 seções 7.1a e 7.2 | Divergências: (a) o doc 11 pede a câmera **depois** do PIN correto; a matriz do doc 07 tem "outra pessoa com PIN errado", que exige captura durante a digitação; (b) "rosto não reconhecido com PIN correto": o doc 11 pede o PIN de novo; o doc 07 abre a isca (nível 2); (c) "vivacidade falha = face não reconhecida" (ver G11); (d) falta dizer que a verificação aparece **igual** com o PIN de coação; (e) o fluxo 4.7 não tem os efeitos por perfil; (f) o dono com falso negativo cai na isca e o doc 11 não descreve como volta | Reescrever 4.7 e 4.8 a partir da matriz do doc 07 seção 7.1a, depois de G11 e G12. Captura durante a digitação, sem tela própria que dependa do PIN. Caminho de volta do dono que caiu na isca (gestor, sem destruição) |
| 01g-G16 | Média | doc 11 seção 3.2 (P08, P09, P10), seção 8, ponto 8, R-U16 | O painel do doc 11 diverge da ADR-0008 e da ADR-0010: (a) P09 mostra a contagem de preservações (revela investigação; o servidor nem conhece a preservação selada); (b) P08 e P10 mostram nós de campo sem a condição da D-R20; (c) "<5" é o desenho anterior ao 01e-E14 (hoje: faixas ou arredondamento); (d) "sem o nível 1b até decisão": a GD-15 já decidiu que ele não existe; (e) P07 não diz que a atribuição de papel privilegiado é objeto assinado conferido contra a lista de cerimônia | Corrigir os cinco. Em (a), só preservação comum (RF58), nunca a selada |
| 01g-G17 | Média | doc 11 seção 6 | Lacunas do protocolo V09: seção 5 desta revisão | Seção 5 |
| 01g-G18 | Média | doc 11 seções 5.3, 6.2 (U09), 6.5; regra 10; doc 06 | O teste facial precisa do rosto real do participante e grava um **gabarito biométrico** dele no aparelho de teste. O doc 11 só trata da foto do termo e diz "sem gravação de rosto". Dado pessoal sensível em teste, sem rito (regra 10; LGPD art. 11) | Rito escrito antes do teste: consentimento específico, gabarito só no aparelho de teste, zerado e conferido ao fim da sessão, parecer do encarregado. A taxa de erro se mede em A76 e A78, com muitas tentativas; o V09 mede só aceitação e tempo |
| 01g-G19 | Média | ADR-0008 seção 3; doc 11 T25, fluxo 4.10, U-5; doc 06 seção 3 | "Posição só no app do comandante" é restrição de **interface**: a posição vai por MLS à sala da missão, e todo aparelho membro a decifra e guarda. Um aparelho de agente capturado desbloqueado tem a posição de toda a equipe. O doc 11 apresenta como garantia | Decidir (Fase 8; D-R20): posição enviada a um **grupo MLS de comando** (comandante, substituto e quem a missão definir), para a restrição ser criptográfica, ou declarar o limite e tirar "só" do texto |
| 01g-G20 | Baixa | ADR-0004 item 13; doc 10 seção 3.1; D-R39 | "Falha do boot verificado" como gatilho de destruição: atualização corrompida ou defeito destrói custódia e trilha não sincronizadas | Falha de boot **não libera** a chave do disco (reversível na base); destruição só por botão, ordem assinada ou violação do gabinete |
| 01g-G21 | Baixa | doc 10 seção 3.2 | Afirma o perfil missão (nível 4) para todo aparelho em missão; o doc 07 seção 7.1a e a D-R40 condicionam ao perfil da instituição (polícia desligada por padrão) | Acrescentar "conforme o perfil da instituição (D-R40)" |
| 01g-G22 | Baixa | doc 07 seção 7.2 | (a) Desembrulhar o gabarito pelo Keystore faz o gabarito em claro passar pela JVM. (b) O perfil missão precisa ser lido antes do desbloqueio, fora do cache cifrado | (a) O hardware embrulha uma chave; o core decifra o gabarito em Rust. (b) Marca de missão protegida por chave de hardware, gravada na saída e no retorno |
| 01g-G23 | Baixa | ADR-0008 seção 2; D-R41 | (a) "CI e/ou autorização do comandante": quando o administrador é o comandante, o "ou" vira autoconcessão sem rastro externo. (b) "Não retroage" só é imposto pelo nó de IA e não tem teste. (c) "Prazo do indicador nunca maior que o da origem" não vale para o que entrou na cópia de custódia da sala de análise | (a) CI sempre; autorização do comandante quando o administrador não é ele. (b) Caso novo em A96. (c) Declarar a exceção no termo |
| 01g-G24 | Baixa | regra 4; `admin-panel/CLAUDE.md` | A lista entre parênteses se lê como exaustiva: ficam de fora Room e SQLite direto, Keychain, sockets, WebSocket, Ktor, e o log (Logcat, `os_log`). O painel web não é coberto por regra de armazenamento no navegador | Trocar "sem" por "por exemplo, sem" nos dois parênteses e pôr a lista completa no teste estático, com log. No `admin-panel/CLAUDE.md`: sem dado de cadastro em armazenamento do navegador e sem recurso externo |
| 01g-G25 | Baixa | doc 11; doc 03 seção 4.12; `CLAUDE.md` | (a) Telas P10 e P11 colidem com os playbooks P10 e P11 do doc 05; etapas E1 a E3 confundem com os cenários E01 a E09; U01 a U20 com U-1 a U-7. (b) A nota inicial diz que a ADR-0007 é "proposta". (c) Doc 03 seção 4.12 e a lista de leitura do `CLAUDE.md` não citam o doc 11. (d) T22 a T24 "Fase 2 ou 4": o doc 07 seção 14 já diz Fase 4. (e) Fluxo 4.5: "a tela diz só concluída" contraria a ADR-0003 seção 7 (dizer que a cópia de custódia fica guardada até a data) | Telas do painel `PW01` a `PW12`; etapas "papel, navegável, nativo" sem sigla; decisões com D-R definitivo (seção 6). Corrigir (b) a (e) |
| 01g-G26 | Baixa | ADR-0010 seção 7; `PIPELINE.md` | (a) Limite não declarado: o investigado que administra uma sala comum pode expurgá-la durante a investigação, e nada impede (consequência aceita da GD-17). (b) `PIPELINE.md`: falta a linha de estado da etapa 5b | Declarar (a) na ADR e no R119; acrescentar (b) |

## 5. Doc 11 — coerência e lacunas do V09

**Coerente:** nada fora do core (PI-1); notificação sem conteúdo; marca sempre visível e sem botão; isca sem tela própria e com teste cego; painel sem conteúdo, posição nem nominal; face nunca decide sozinha (PI-11); rótulo de indício (PI-12); "nunca destruído" (PI-13); ferramenta local, sem nuvem (regra 10). Divergências: G14 a G16, G18, G19, G25.

**Lacunas do protocolo (G17):**
1. **R-U1 medido só por pergunta.** Falta linha de base (mesma tarefa num mensageiro comum) e medida de uso real no piloto. Preferência declarada diante do avaliador vale pouco.
2. **Critérios que a amostra não sustenta.** "0 erros", "100%" e "1 em 20" com 5 a 8 pessoas não medem taxa. Fixar tentativas por pessoa e tratar as metas como formativas.
3. **Teste cego da isca.** Quem precisa não distinguir é o **observador**, não o participante: escolha forçada entre duas sessões (normal e coação), várias rodadas por observador, acerto perto do acaso; tempo medido por instrumento. Repetir nos dois perfis.
4. **Retenção.** O PIN de coação e o gesto de pânico se esquecem: repetir U05 e U06 depois de 2 a 4 semanas sem uso.
5. **Recuperação.** Faltam tarefas do que acontece depois: voltar de pânico acidental, de trava nível 3, de lease vencido, da isca por falso negativo. O tempo para voltar a trabalhar é o que empurra para o canal paralelo.
6. **Tarefas ausentes:** código de segurança entre membros (T21, RF04); perceber mensagem urgente com notificação genérica ou nenhuma (R-U12); banda estreita (T27); o que o membro entende ao ver uma concessão (T30); compreensão dos itens 2, 13 e 14 do termo.
7. **Estresse de campo:** tela molhada, sol direto, veículo em movimento, capacete e pintura no rosto (perfil Exército).
8. **Ética:** recrutar por canal neutro (convite da chefia é pressão); rito do rosto real (G18).
9. **Por perfil:** o PIN de coação e a face têm efeitos diferentes no perfil missão; o protocolo precisa de uma rodada por perfil.

## 6. IDs, referências e fusão dos R-U

- **Sem colisão** em R112 a R119, RF92 a RF97, RL13, RL14, A76 a A99, F-R13 a F-R15, D-R37 a D-R41 e P1-S (busca por definição em tabela). As repetições são linhas de índice do doc 04.
- **Colisão:** telas P10 e P11 do doc 11 com playbooks do doc 05 (G25).
- **Referências:** não achei referência quebrada nas seções citadas pelo `01f`. Faltam as do doc 11 (G25c).
- **Fusão dos R-U (numerar a partir de R120):**

| Doc 11 | Destino |
|---|---|
| R-U2, R-U3 | Fundir no **R55**, com o impacto do perfil missão (G12) |
| R-U4 | Fundir no **R67** |
| R-U5 | Dividir: destruição do dono no **R67**; adversário que passa no **R116** |
| R-U8 | Controle novo no **R69** |
| R-U10 | **R106** |
| R-U15 | **R102** |
| R-U16 | **R105** (pressão por mais dado no painel) |
| R-U1, R-U6, R-U7, R-U9, R-U11, R-U12, R-U13, R-U14, R-U17 | Novos: **R120 a R128**, nessa ordem |

- **Decisões:** U-1, U-2, U-4, U-6 e U-7 viram **D-R42 a D-R46**. U-3 se divide: limite facial dentro da **D-R40**; margem de tempo entre PIN normal e de coação, **D-R47**. U-5 junta-se à D-R20 e à pergunta 9 do 01c (G19).
- **Regressões:** nenhuma nas ADRs 0003 a 0006. Contradições novas: G01, G07c, G14, G15, G16, G21.

## 7. GD-17, GD-18 e regra 12

- **GD-17:** a leitura do Opus é a correta e a mais restritiva. A frase da GD ("o expurgo e a contagem do prazo de guarda das cópias de custódia ficam suspensos") admite duas leituras; convém o responsável confirmar que "expurgo" ali é a eliminação da cópia de custódia, não o expurgo nos aparelhos (R3-5).
- **Metadados de sala comum × regra 14 (R3-2):** não há conflito com a letra. "Expurgada ao concluir" trata do conteúdo e das chaves; os metadados seguem o doc 06 seção 8.2. Condições: só metadados que o servidor já tem, nunca cifrado MLS de sala comum, retrato selado para quórum e aberto pela regra 14, prazo, e transparência dizendo que a preservação alonga a guarda (G07).
- **GD-18, nó de campo:** exceção bem escrita, restrita ao nó, com a perda nomeada e sem aceite formal indevido. Ajuste: G20.
- **GD-18, face:** "nunca decide sozinha" está aplicado em todas as linhas da matriz. Vivacidade, gabarito e perfil estão descritos. Ajustes: G11, G12, G22.
- **Regra 12 (R3-3):** processar quadro em memória cabe na letra. O gabarito não é quadro, mas deriva de quadros e é dado biométrico. Recomendo **escrever na regra**, por decisão do responsável: "o gabarito facial do dono fica só no aparelho, cifrado por chave de hardware, e nunca é enviado; a câmera só é usada com o app em primeiro plano, na tela de conteúdo e no desbloqueio". Até lá vale a leitura restritiva que o Opus aplicou.

## 8. Recomendações por pergunta

### R3-1 a R3-10

| # | Recomendação |
|---|---|
| R3-1 | **Sim.** Sinal de atribuição que é dado cifrado exige a base da regra 14 e quórum. Sem a base, a investigação fica no nível 1a |
| R3-2 | **Confirmar as duas**, com as condições da seção 7 (sem cifrado MLS de sala comum; retrato selado sob a regra 14; transparência) |
| R3-3 | **Escrever na regra 12** o gabarito e o desbloqueio (texto na seção 7), junto com a frase de escopo da D-R30 |
| R3-4 | **Aceitar** gabarito cifrado por chave do hardware com comparação no core, com o limite declarado e o ajuste G22a |
| R3-5 | **Confirmar a GD-17** na leitura restritiva, com a frase sobre "expurgo" esclarecida |
| R3-6 | **Adotar a primeira** (o aparelho guarda a cópia de custódia que gerou até a central confirmar). Manter no estudo da Fase 8 a entrega do selado ao aparelho na emissão: reduz a perda sem mudar a GD-18 |
| R3-7 | **Sim nos três.** Base mínima: CI sempre; autorização do comandante quando o administrador não é ele (G23a) |
| R3-8 | **Um PIN errado só depois de medir**, e com G11: duas capturas de "outra pessoa", vivacidade fora do gatilho de destruição até ser medida. Polícia: perfil missão desligado por padrão |
| R3-9 | Segunda autorização: iniciador detentor → corregedoria; iniciador corregedoria → autoridade **fora da classe dos detentores** quando a suspeita alcança essa classe. Substitutos nomeados de antemão no `instituicao.yaml`. Reservas com cartão emitido de antemão. Prazo de 90 dias, no máximo duas renovações, cada uma com quórum não impedido. Registro guardado até a revisão final e o prazo de responsabilização **[conferir]**. Aviso no arquivamento como padrão |
| R3-10 | **Mínimo:** "por exemplo" nos dois parênteses (G24). Não vale reabrir o texto além disso antes da Fase 1 |

### Restantes

| # | Recomendação |
|---|---|
| Q4 | **Sim.** Confirmação de época como proposta, medida no spike e em banda estreita |
| Q5 | **Sim.** Selagem na Fase 2 pelo `directory`, com chave pública do pacote |
| N3 | **Aprovar** a redação do 01c-C23. Se a opção 1 de G01 for aceita, a mesma edição pode nomear a revisão de controle |
| N4 | Segunda assinatura pelo chefe da inteligência ou pelo superior imediato do criador, fixado no `instituicao.yaml`; motivo em lista fechada |
| N5 | **Aceitar** a posição na cópia de custódia; ver G19 |
| N6 | **Aceitar** RPO 0 como meta e o teto de 12 meses configurável |
| Pergunta 10 do 01c (01c-C17) | Manter a recusa total de sessão direta. Com a GD-18 fica urgente: cadeia de commits e lista de revogação importáveis do aparelho do comandante, com a mesma verificação; descongelar com duas pessoas (comandante e chefe da inteligência); credencial do nó renovada a cada dia quando houver enlace. Decidir antes de aceitar a ADR-0004 |
| Demais do 01c | Valem as recomendações do 01e seção 9. Reforço para a D-R24 (G04c) |

## 9. A cadeia pode seguir para G0?

- **Pelo critério literal, não:** há dois altos abertos (G01, G02).
- **Os dois estão só na ADR-0010**, que é nova, é proposta e é da Fase 4. Nada entre G0 e G3 depende dela. As ADRs 0003 a 0009 e as regras 2, 4 e 14 não têm crítico nem alto aberto, e o 01e-E01 está fechado no texto.
- **Recomendo seguir para G0 com uma destas condições**, à escolha do responsável, sem quarta rodada completa:
  1. passe curto do Opus na ADR-0010 para G01 e G02 (texto e duas decisões: quórum em toda sessão, e controle independente), conferido por mim só nesses pontos; ou
  2. o status da ADR-0010 e a linha dela no README passam a dizer "proposta; não aceitar antes de fechar 01g-G01 e G02", e o responsável registra no `GD-respostas.md` que ela fica fora do critério da 1b.
- **Antes da etapa 3:** G13 no critério do spike.
- **Antes da primeira sessão do protótipo:** G14 a G19 e G25 no doc 11 (Sonnet), depois de G11 e G12 (Opus).
- **Antes da Fase 2:** G11, G12, G22 e a D-R40. **Antes da Fase 4:** G03 a G10, D-R37, D-R38 e parecer jurídico (G09).

## 10. Próximo passo sugerido

1. Responsável: opção de G01 (quórum em toda sessão da investigação sigilosa?); R3-5 (GD-17); R3-3 (regra 12); R3-8 com G11 e G12 (gatilhos de destruição no perfil missão); condição da seção 9.
2. Opus, só texto: G01 e G02 na ADR-0010; G13 na ADR-0005 e no critério do spike; G11, G12, G20 a G23 nos docs 07 e 10 e nas ADRs 0004 e 0008.
3. Sonnet: doc 11 (G14 a G19, G25), fusão dos R-U e numeração (seção 6).
4. Orquestrador: registrar esta revisão no `PIPELINE.md`; depois, G0.
