# ADR-0005 — Um commit por época, detecção de bifurcação e pertença verificável

**Status:** proposta (etapa 1b da Fase 0; achado D16, base de D13, D14 e D19; **segunda rodada**: item 2 reescrito pelo 01c-C02; **terceira rodada**: sinais classificados, commit retido e confirmação de época, 01e-E02, E03, E09a, E16; **quarta rodada**: bifurcação provada separada de anomalia não autenticada, 01g-G13) · **Complementa** ADR-01 (MLS) e doc 02 seções 5 e 7 · **Provar** no spike MLS (etapa 3) antes de aceitar

## Contexto
O doc 07 seção 5 prevê "detecção de bifurcação de época" sem dizer como. Três problemas (D16):
- servidor hostil entrega commits diferentes a partes do grupo (bifurcação);
- dois administradores fazem commit na mesma época (corrida);
- o diretório guarda membros por sala separado do estado MLS e pode divergir dele: furo na autorização por sala (regra 5).

A primeira versão do item 2 punha um resumo do estado **dentro do conteúdo MLS**. Não detecta nada (01c-C02): em bifurcação, cada ramo tem chaves de época diferentes, e quem está num ramo não decifra as mensagens do outro nem lê o resumo; e quem decifra já concorda no estado, porque a chave da época deriva do contexto do grupo.

## Decisão (proposta)
1. **Um commit por época:** o serviço de entrega (ou o nó de campo, ADR-0004) aceita, para cada grupo e época, **um único** commit, o primeiro que chegar, e recusa os demais com "época já avançou". Quem perdeu busca o commit vencedor, aplica e refaz o seu sobre a época nova.
2. **Detecção de bifurcação sem decifrar (01c-C02).** Só com recursos da biblioteca MLS; nada à mão (regra 1).
   - **(a) Hash do transcript em claro, autenticado:** cada mensagem leva, nos **dados autenticados** do quadro MLS (campo em claro coberto pela cifra; RFC 9420 `authenticated_data`), a época e o **hash do transcript confirmado** dessa época, como a biblioteca o expõe **[verificar no spike]**. Quem recebe compara com o próprio, **antes** de tentar decifrar.
   - **(b) Sinais de bifurcação:** (1) mensagem da **mesma época** com hash do transcript diferente do local; (2) mensagem da mesma época que **não decifra**; (3) commit para uma época já aplicada que **difere** do aplicado, ou commit que não aplica sobre o estado local; (4) mensagem de época **futura** sem que o commit dessa época tenha chegado depois de um prazo T; (5) par que continua enviando em época **antiga** depois de T do commit aplicado (01e-E03b).
   - **(c) Classificação dos sinais (01e-E02; revista no 01g-G13):** pacote que não decifra ou traz hash diferente é tão forjável repetido quanto isolado: o servidor, ou um membro, envia quantos quiser. Por isso a quantidade não prova nada, e só o que o cliente **autentica** prova bifurcação.

     | Classe | Quando | Resposta |
     |---|---|---|
     | **Bifurcação provada** | Commit conflitante para época já aplicada (sinal 3) cuja assinatura de **administrador registrado** no estado autenticado do grupo o cliente verifica; ou divergência mostrada pela **importação de cadeia** (ADR-0004 seção 5) | **SEV1** ao administrador da sala e à segurança, com época e hashes (sem conteúdo). O administrador decide: recriar a sala ou refazer a inclusão a partir do ramo legítimo |
     | **Autenticado** | A mensagem decifra e o hash do transcript nela diverge do local | Falha atribuível ao **remetente** (membro inconsistente ou que mente no campo), não bifurcação. Alerta SEV2 com o remetente; o administrador decide |
     | **Anomalia não autenticada** | Qualquer quantidade de pacotes que não decifram ou trazem hash diferente; sinais 4 e 5 | Alerta **"servidor ou membro anômalo"** (SEV2), **sem** recomendar recriar a sala; nunca SEV1 sozinho. Limite de alertas por sala e época, contra fadiga. Conta para a confirmação de época (item f) e para a comparação do código de segurança (RF04) |
     | **Não conta** | Duplicata (chave da mensagem já consumida), fora de ordem além da janela, retransmissão de rádio, byte corrompido, quando a biblioteca dá erro próprio para o caso **[verificar no spike]** | Só registro local |

     O aparelho continua no próprio ramo, não aplica o commit divergente e não exibe nada do outro ramo. **[verificar no spike]:** a biblioteca deixa processar e verificar a assinatura de um commit para uma época já aplicada (retenção dos segredos da época anterior, ou commit como mensagem pública)? Se não deixar, a bifurcação provada só vem pela importação de cadeia, e a detecção em banda fica como anomalia não autenticada.
   - **(d) Valor falso do servidor:** o servidor pode trocar o hash em claro, mas a troca quebra a autenticação e a mensagem não decifra; o efeito é **alarme forjado** (anomalia não autenticada), nunca bifurcação escondida nem SEV1.
   - **(e) Limite declarado: partição total.** Ramos que nunca trocam mensagens não são detectáveis em banda. Com servidor hostil, que controla o roteamento, é o caminho **normal**, e o caso mais provável é o **commit retido**: a remoção chega só a parte do grupo, quem ficou na época antiga recebe só mensagens antigas e o hash confere (01e-E03). A detecção vem na reunificação, na **importação de cadeia** (ADR-0004 seção 5), pelos sinais 4 e 5, pela confirmação de época (item f) e pela comparação do código de segurança entre membros (RF04). Esse código é por sala e por época (muda a cada commit) e deriva do `epoch_authenticator` do MLS, que existe para isso **[verificar no spike]**; o hash do transcript fica para o campo em claro (01e-E16).
   - **(f) Confirmação de época (01e-E03a) [proposta; pergunta Q4 do 01e]:** depois de aplicar um commit, cada membro envia (ou leva de carona numa mensagem) uma confirmação curta na época nova. O cliente do administrador mostra "membros que ainda não confirmaram" e alerta depois de um prazo. O servidor não forja a confirmação, porque ela é cifrada na época nova. Transforma a partição silenciosa em ausência visível; custo de uma mensagem curta por membro por commit, a medir no spike e em banda estreita (Fase 8).
   - O hash do transcript em claro revela ao servidor só um identificador da época, que ele já distingue pela ordem dos commits.
3. **Quem altera membros:** só chaves de administrador registradas no **estado autenticado do grupo** (extensão do contexto), na criação ou por commit de administrador vigente. Os clientes **não** aprendem chave de administrador pelo diretório (D19).
4. **Pertença verificável pelo servidor:** junto com o commit, o cliente envia uma **declaração de mudança de membros** (sala, época nova, aparelhos incluídos e removidos), assinada pela chave de administrador. O diretório verifica a assinatura contra as chaves de administrador da sala **antes** de gravar `memberships` (regra 6). Os clientes conferem que a declaração bate com o commit; divergência é alerta.
   Alternativa a decidir no spike: commits como mensagem MLS pública (o servidor vê e verifica a própria assinatura do commit, sem declaração separada), se isso vazar menos que a declaração. **Condição (01e-E09a):** commit público só se **nenhum dado sensível** transitar em proposta (marca, motivo, papéis, chave de período, nome da operação); senão, commit privado e declaração separada (texto acima). O spike levanta o que o commit público exporia.
5. **Aparelho só aplica commit** que venha de chave de administrador registrada, seja da época seguinte à que tem e tenha declaração conferida.
6. Os dados de controle do grupo (marca de custódia e chave de período, ADR-0003; papéis da sala, doc 01 RN18) vão no mesmo estado autenticado e seguem as mesmas regras.

## Pendente (01c-C16, depende do responsável)
Se só o administrador faz commit, os outros membros só renovam a própria folha por **proposta**. Falta decidir: quais propostas um membro pode enviar (atualização e autoexclusão), se o commit externo fica desligado, o mínimo de administradores por sala (proposta do 01c: dois) e o que fazer se os dois se perderem (proposta: recriar a sala). O spike verifica o que a biblioteca oferece; a decisão fica para depois das respostas.

## Alternativas rejeitadas
- **Qualquer membro faz commit a qualquer momento:** corridas e bifurcação sem dono.
- **Detecção só por comparação manual de código de segurança:** tardia e dependente de pessoas. Fica como complemento para a partição total (item 2e).
- **Diretório como fonte da pertença sem verificar o commit:** servidor hostil autoriza quem quiser a publicar e a receber.
- **Resumo do estado dentro do conteúdo cifrado** (primeira versão do item 2): não detecta bifurcação (01c-C02).

## Consequências
- Commit pode falhar por corrida; o cliente refaz. Custo pequeno com poucos administradores.
- Toda mensagem leva época e hash do transcript nos dados autenticados (algumas dezenas de bytes).
- O servidor escolhe o vencedor da corrida e pode **atrasar** ou **reter** uma remoção para parte do grupo. O alerta de remoção pendente (doc 07 seção 7.1, D17) só avisa que o administrador não fez o commit; quem mostra que os membros não o receberam é a confirmação de época (item 2f) e os sinais 4 e 5. Risco R98 em **Médio** até os testes da Fase 1 (doc 04 seção 5.4).
- O spike MLS (etapa 3) precisa provar: um commit por época; **bifurcação induzida pelo servidor** (dois commits na mesma época entregues a metades do grupo) detectada pelos sinais do item 2b, nos dois lados; **partição total** não detectada em banda e detectada na reunificação; bifurcação provada só com commit conflitante de administrador verificado ou importação de cadeia, e pacotes forjados pelo servidor ou por um membro, em qualquer quantidade, só como anomalia não autenticada, sem SEV1 (01g-G13); se a biblioteca deixa verificar commit de época já aplicada; matriz de falso positivo (duplicata, fora de ordem, época anterior atrasada, byte corrompido) sem SEV1; **commit retido** para parte do grupo e confirmação de época medida; membro que mente no campo classificado como remetente inconsistente; declaração de membros conferida pelo diretório e pelos clientes (doc 03, Fase 0).

## Revisar quando
O spike mostrar o que a biblioteca oferece (dados autenticados por mensagem, acesso ao hash do transcript confirmado, extensões de contexto, mensagens de controle públicas ou privadas, propostas por membro).
