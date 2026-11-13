# Portão GD — respostas do responsável pelo produto

Respostas às perguntas 1, 2, 9 e 11 da revisão de desenho (`01-revisao-desenho.md`), com os esclarecimentos dados em conversa. Entrada da etapa 1b.

## Decidido

| # | Tema | Decisão |
|---|---|---|
| GD-1 | Custódia por tipo de sala | O risco de expor o período inteiro se o HSM ou o quórum for comprometido **só é aceito em sala marcada como sob custódia**. No perfil Exército, o motivo é investigação militar. **Sala comum:** MLS puro; ao concluir, sai dos aparelhos e é **expurgada**, sem cópia. **Sala sob custódia:** tem cópia para a chave de quórum no HSM offline; ao concluir, sai dos aparelhos e fica guardada **até 5 anos a contar da conclusão**, depois é expurgada com recibo |
| GD-2 | Flag por perfil | A flag de custódia vem do pacote de instalação (`instituicao.yaml`, ADR-15). No perfil Exército ela fica ativa, com o motivo investigação militar. Perfil polícia: **pendente** (ver abaixo) |
| GD-3 | Ciclo da sala | A **equipe de inteligência** marca a sala como ativa. Para marcar como **concluída**, é preciso o relatório interno assinado eletronicamente no sistema corporativo, e assinam no app o **comandante da missão**, o **subcomandante** e o **responsável que pediu a abertura**. A marca de custódia é feita na criação, é assinada, fica visível aos membros e não pode ser revertida enquanto a sala existir |
| GD-4 | Revinculação na v1 | Vale só para operação **ativa**, e o histórico vem do aparelho de outro membro pelo MLS. **Sem HSM online e sem aprovação remota por quórum na v1.** Aprovam o chefe da inteligência e o superior da área. O app gera o **termo de ciência e aceite** do proprietário, com foto e assinatura no próprio app, e exporta o documento para anexar no sistema da corporação |
| GD-5 | Guarda do termo | O termo, que tem foto e portanto dado biométrico sensível (LGPD art. 11), fica guardado no MIGDOL **cifrado para a chave de quórum**, com prazo definido, e o hash entra na auditoria. Ele é aberto para responder a uma CI futura, por quórum presencial (regra 14) |
| GD-6 | CI × ordem judicial | A **CI** é o comunicado interno oficial que registra a demanda (ex.: processo administrativo). Ela **não substitui** a ordem judicial: a ordem judicial é cumprida, e a CI registra e tramita o pedido. Abrir custódia exige quórum presencial e uma base: ordem judicial, ou CI em processo administrativo com validação jurídica [conferir] |
| GD-9 | D-R17 e D-R18 | Ficam para **depois dos spikes**, no portão G3. Até lá, nenhum grau classificado é tratado |
| GD-11 | Missão de campo | Durante a missão a **central fica congelada**: não altera membros nem entrega para os grupos de missão. Só o nó de campo ordena. O administrador da missão é o **comandante da missão**, e o substituto é o **subcomandante** |

## Consequências para a etapa 1b
- Resolve D01–D03 na v1: não há HSM online nem aprovação remota. A custódia fica só em HSM offline e só em sala sob custódia.
- A regra 14 do `CLAUDE.md` precisa mudar: a revinculação deixa de usar quórum remoto de comandantes e passa a ter aprovação dupla (chefe da inteligência + superior), sem decifrar nada da instituição.
- O doc 07 §7.3 e §7.4, a ADR-0003 e a ADR-0004 (D14) precisam ser reescritos conforme as decisões acima.
- Resolve D14: com a central congelada, ela não readiciona aparelho capturado.

## Perfil polícia (decidido)
- A custódia é uma **opção configurável no pacote de instalação**. No perfil Exército vem **ativa**, com o motivo investigação militar. No perfil polícia vem **desligada por padrão**, e a instituição pode ligá-la, com o motivo corregedoria ou inquérito, quando o edital exigir.
- A configuração é **assinada** no `instituicao.yaml` e fica **visível aos membros**. Mudá-la exige cerimônia registrada.
- Mudar a configuração **não retroage**: salas criadas antes da mudança não ganham cópia de custódia.
- Atende ao RNF-T3: a configuração é por instituição, sem mudar o código.

## Decisões após a revisão do Fable (01c), 12/11/2026

O responsável aceitou as três recomendações:

| # | Tema | Decisão |
|---|---|---|
| GD-12 | Regra 2 (D30, pergunta 4 do 01c) | **Aprovado** o texto proposto no `01b` seção 3, com os ajustes de C19 e incluindo Swift (ADR-0007): chaves de membro (MLS, assinatura, cache) só no core Rust; chaves de hardware (StrongBox, Secure Enclave) não saem do hardware, e Kotlin/Swift só seguram a referência, com o uso chamado a partir do core; chaves da instituição (AC, quórum, cerimônia) ficam em HSM. A regra 7 (D31) **não** foi decidida e continua pendente |
| GD-13 | Revogação (C01, pergunta 6) | Revogar credencial e destruir segredo são **atos separados**. Em **toda** revogação o segredo de lease é guardado cifrado para a chave de quórum, com prazo igual ao da apuração. Destruir exige duas autorizações e registro. A classe entra na lista da regra 14 |
| GD-14 | Custódia com a flag ligada (D-R29, pergunta 1) | **Opção B:** com a flag ligada, a custódia é o padrão; sala comum só com motivo registrado e segunda assinatura; salas de sistema sempre comuns |

Também valem as recomendações do 01c sem decisão do responsável: R66 volta a **Alto** até o aceite do comitê; na conclusão, quando o solicitante já é comandante ou subcomandante, a terceira assinatura passa ao superior imediato do solicitante, definido na criação.

## Decisões após a revisão do Fable (01e), 12/11/2026

| # | Tema | Decisão |
|---|---|---|
| GD-15 | Nível 2 da ADR-0008 (planejamento) | O analista lê indicadores nominais por **concessão** da sala de origem, com CI e/ou autorização do comandante; a concessão é registrada. A regra 3 **não muda**: o painel mostra só o nível 1a; agregados derivados de conteúdo vão ao app por MLS |
| GD-16 | Investigação sigilosa de vazamento ou infiltrado | Caminho próprio, para quando a equipe não pode saber (ex.: suspeita sobre um detentor do quórum). **Qualquer** detentor do quórum ou a corregedoria pode **iniciar**; exige CI **e** uma segunda autorização independente; quem está sob suspeita fica **impedido** de aprovar; os membros não são avisados; o ato fica num **registro selado** para o quórum. A análise roda **só no ambiente isolado** (nó de análise), ninguém lê conversa, e a saída é um **relatório de rastro de suspeita, sem conversas**, para o MP pedir medidas cautelares (ordem judicial, afastamento do investigado). Dois níveis [conferir]: sem ordem judicial, só metadados e sinais de atribuição (doc 05); com ordem judicial, também conteúdo, dentro do ambiente isolado. Relatório revisado e assinado por pessoa antes do envio (regra 11) |
| GD-17 | Gravação para estes casos | Aberta a investigação, as salas do escopo entram em **preservação**: o expurgo e a contagem do prazo de guarda das cópias de custódia ficam suspensos, por ordem selada, até a decisão. Gravação **para frente** só existe onde há custódia (com a flag ligada, a custódia é o padrão, GD-14). **Sala comum não é gravada às escondidas:** isso exigiria membro oculto no grupo MLS, que é porta dos fundos e contraria o desenho (ADR-02, regra 14). Pendente de confirmação do responsável |
| GD-18 | Destruição na captura | **Nó de campo capturado:** destruição imediata, aceitando perder o que não foi sincronizado; **exceção escrita à GD-13**, só para o nó de campo. **Celular:** reconhecimento facial de quem acessa, modelo no aparelho, sem terceiros (regra 12), com vivacidade e modelo do rosto no hardware seguro. A face **nunca decide sozinha**: destrói com face errada **e** PIN errado, ou com PIN de coação. **Por perfil:** destruição imediata ativa em missão de campo (Exército, captor adversário); no uso policial urbano vale suspender e revogar sem destruir (D-R14) |
| GD-19 | Interface e plataformas | **ADR-0007 aceita:** interface nativa (Compose, SwiftUI) e painel web Linux. Documentar o **protótipo de interface** dos três projetos e o **protocolo de usabilidade** (V09) num doc 11 |
| GD-20 | Suíte candidata | **P-256** como curva candidata até a escolha do HSM (recomendação do Fable, aceita no pacote "sim, registre") |

## Decisões após a revisão do Fable (01g), 12/11/2026

O responsável aceitou as cinco recomendações do `01g`:

| # | Tema | Decisão |
|---|---|---|
| GD-21 | Quórum na investigação sigilosa (01g-G01) | **Toda sessão** da investigação sigilosa, em qualquer nível, exige quórum presencial de detentores **não impedidos**. Os níveis diferem só pela base (CI validada ou ordem judicial) |
| GD-22 | Controle externo do rito (01g-G02, D-R37, D-R38) | Quando a suspeita alcança os detentores, a segunda autorização vem **de fora** dessa classe; no máximo **duas renovações**, cada uma com quórum; o revisor não é o iniciador nem a segunda autorização; **aviso ao investigado no arquivamento** como padrão; registro guardado até a revisão final |
| GD-23 | GD-17 confirmada, leitura de R3-2 | Na preservação, "expurgo suspenso" é a eliminação da cópia de custódia, não o expurgo nos aparelhos. Metadados de sala comum podem ser preservados, **nunca o cifrado MLS**, com o retrato selado para quórum e aberto pela regra 14; a transparência aos membros diz que a guarda pode se alongar |
| GD-24 | Gatilhos de destruição no perfil missão (R3-8, D-R40, 01g-G11, G12) | Número de PINs errados só depois de medir (proposta: um PIN errado, se a medição permitir). "Outra pessoa" exige **duas capturas**. Falha de vivacidade não conta para destruir até ser medida à parte. Medir o acionamento acidental do PIN de coação antes de ligá-lo ao nível 4 em campo |
| GD-25 | Regra 12 (R3-3) | Escrever na regra: o gabarito facial fica só no aparelho, cifrado por chave de hardware, nunca é enviado; a câmera só é usada em primeiro plano, na tela de conteúdo e no desbloqueio |
