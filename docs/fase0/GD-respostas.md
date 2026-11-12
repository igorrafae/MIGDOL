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
