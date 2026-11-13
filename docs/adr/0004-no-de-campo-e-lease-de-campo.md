# ADR-0004 — Nó tático de campo e lease de campo

**Status:** proposta, **revisada na etapa 1b da Fase 0** (achados D13–D15, D22, D29 de `docs/fase0/01-revisao-desenho.md`; decisão GD-11; segunda rodada: itens 3, 6 e 7 ajustados a 01c-C02 e GD-13; **terceira rodada**: item 13, destruição imediata do nó capturado como **exceção escrita à GD-13** (GD-18; 01e-E07), e passagem do lease entre central e nó (01e-E12); 01c-C17 pendente) · **Contraria** ADR-05 (lease de 12 h emitido só pela central) · **Depende de** ADR-0005 (um commit por época) e ADR-0006 (lease criptográfico) · **Complementa** ADR-0001 (mesmo binário, perfil de campo) e ADR-0003 (custódia) · **Fase 8** (doc 10)

## Contexto
O responsável decidiu operação em campo por rádio, para polícia e Exército (D-R22). Em campo não há conexão com a central por horas ou dias. Com a ADR-05, o aparelho perde o acesso em 12 h sem a central.

A versão anterior desta ADR tinha dois erros críticos apontados pela revisão:
- **D13:** "só o nó de campo faz commits" exige que o nó seja membro (lê tudo) ou tenha autoridade aceita pelos clientes (nó capturado ligado inclui aparelho do adversário).
- **D14:** "a central prevalece" devolve à sala um aparelho removido em campo, deixa o aparelho capturado falar com a central e torna indecifráveis as mensagens do ramo de campo.

Decisão do responsável (GD-11): durante a missão a **central fica congelada** e só o nó de campo **ordena**; o **administrador da missão** é o **comandante da missão**, e o **substituto** é o **subcomandante**.

## Decisão (proposta)
1. **Nó tático de campo:** o mesmo binário do servidor (`cmd/server`) em perfil de campo: gateway, entrega e diretório da missão. Sem nó de IA, sem chave privada de custódia, sem cartão de comandante. **Não é membro de nenhum grupo** e **não tem chave que altere membros**. Só ordena e retransmite cifrado (D13).
2. **Quem altera o grupo em campo:** só o aparelho do **administrador da missão** (comandante da missão) e o do **substituto** (subcomandante). Os dois ficam registrados como administradores no **estado autenticado do grupo** antes de sair (ADR-0005). Os clientes só aceitam commit que altere membros assinado por um deles.
3. **Ordenação:** o nó aceita **um commit por época**, na ordem de chegada, e entrega nessa ordem (ADR-0005). Se os dois administradores fizerem commit na mesma época, vale o primeiro; o outro refaz sobre a época nova. Toda mensagem leva a época e o hash do transcript nos dados autenticados em claro (ADR-0005 item 2): nó que bifurque commits entre aparelhos que ainda trocam mensagens é detectado por eles; partição total só se detecta no retorno, na importação da cadeia (item 5). Nó que **retenha** um commit atrasa a mudança; quem percebe é o administrador, pelo alerta de remoção pendente.
4. **Central congelada (D14, GD-11):** ao sair a missão, a central marca os grupos de missão como **congelados**. Não aceita commit, não entrega mensagem para esses grupos e **recusa sessão direta dos aparelhos em missão**. O campo é o único escritor até o retorno.
5. **Retorno:** a central importa a **cadeia de commits** do campo, verifica cada assinatura e a continuidade das épocas antes de gravar (regra 6), e só então descongela. Não há "central prevalece": durante a missão não existe ramo central. Cadeia inválida: o grupo segue congelado e abre SEV1.
6. **Lista de revogação de campo:** cada aparelho removido em campo entra numa lista assinada pelo administrador da missão. O nó a envia à central **pelo primeiro enlace** que houver (rádio, satélite autorizado ou retorno), antes de qualquer outro dado. A central revoga certificados e leases desses aparelhos; o segredo de lease de cada um fica guardado para a chave de quórum, sem destruição (ADR-0006 item 8; GD-13).
7. **Lease de campo (ADR-0006, D22):**
   - o nó entrega a cada aparelho um **segredo de missão** próprio, distinto dos da central, trocado a cada renovação (proposta: 12 h);
   - o nó guarda **só o segredo vigente** de cada aparelho;
   - **teto de 72 h** sem contato (D-R26), medido no aparelho pelo **relógio monotônico** desde o boot, que o usuário não ajusta; reinício apaga a memória e exige nova entrega;
   - na remoção em campo, o nó **sela** o segredo de missão do aparelho removido com a chave pública da classe "segredo de lease guardado" e apaga a cópia em claro (ADR-0006 itens 8 e 9);
   - **passagem do lease** na saída e no retorno pelo protocolo de duas fases, com o nó como emissor da geração de campo e a central como emissora da geração do retorno (ADR-0006 item 9; 01e-E12);
   - nó capturado: destruição imediata (item 13).
8. **Credencial de campo do nó:** emitida pela AC da instituição, com **validade curta, igual à janela da missão**, verificada pelos aparelhos a cada renovação. Vencida, os aparelhos não renovam com o nó. A revogação feita na central não chega ao campo: o limite é a validade curta. **Pendente (01c-C17, pergunta 10 do 01c):** retorno com o nó destruído ou capturado (cadeia importável também do aparelho do administrador da missão, com a mesma verificação; quem descongela), renovação diária da credencial do nó quando houver enlace, e alcance da recusa de sessão direta (item 4) com cobertura celular parcial.
9. **Remoção e atualização em campo (D15):** remoção de membro e commit de atualização são permitidos, feitos pelo administrador da missão e serializados pelo nó. Em banda estreita, a remoção tem prioridade e orçamento de banda reservado; a atualização periódica espera banda larga ou o retorno. Medir em F-R1 e F-R9 (doc 10).
10. **Custódia em campo:** sala de missão sob custódia continua gerando cópia para a chave pública do período. O administrador da missão leva, da base, chaves públicas de período assinadas em número suficiente para a missão, e a pública vigente da classe "segredo de lease guardado" vai para o nó (ADR-0003 seções 2 e 4). O nó guarda e sincroniza as cópias; não as abre.
11. **Ordens assinadas (D29):** em campo, o administrador da missão é o **único escritor** das ordens da sala (expurgo); o contador continua o da sala, e o substituto parte do último valor visto no estado do grupo.
12. **O rádio é só transporte IP.** Nenhum protocolo de rádio é implementado no projeto.
13. **Nó capturado: destruição imediata (GD-18).** **Exceção escrita à GD-13, só para o nó de campo.** A GD-13 exige duas autorizações para destruir segredo selado; aqui a destruição é um ato só, ou automática, porque há captor adversário.
   - **O que se destrói:** a chave do disco do nó (o disco inteiro fica ilegível) e, antes, se der tempo, a memória: segredos de missão vigentes, segredos selados de removidos ainda não sincronizados, mensagens e cópias de custódia não sincronizadas, trilha de auditoria da missão não sincronizada, lista de revogação de campo não enviada e credencial de campo.
   - **Gatilhos [proposta; D-R39]:** botão físico; ordem de destruição assinada pelo administrador da missão ou pelo substituto (chave de administrador do estado do grupo, com biometria); violação do gabinete. **Não acionam:** perda de enlace, falta de energia, cabo conectado, queda do equipamento sem violação (para não destruir por acidente). **Falha do boot verificado** não destrói: o nó só deixa de liberar a chave do disco, o que é reversível na base; assim atualização corrompida ou defeito não destroem custódia e trilha não sincronizadas (01g-G20).
   - **Perda aceita (GD-18), registrada por escrito:** (1) o cache de missão de aparelhos não apreendidos que reiniciarem ou passarem do teto antes de nova entrega (quem volta à base com o segredo em memória reembrulha e não perde); (2) cópias de custódia e trilha de auditoria da missão ainda não sincronizadas, o que contraria, só nesse caso, o RPO 0 da auditoria (doc 04 seção 8); (3) segredos já selados de aparelhos removidos ainda não sincronizados (o cache deles fica ilegível para sempre) e a lista de revogação de campo ainda não enviada. Riscos R89 e R118 (doc 10 seção 8).
   - **Mitigações que não mudam a decisão:** sincronizar cedo e a cada enlace, na ordem lista de revogação, trilha, cópias de custódia; a lista de revogação e a cadeia de commits ficam também no aparelho do administrador da missão, que as assinou, para entrega no retorno (01c-C17, pendente); **[proposta, não decidida]** cada aparelho guarda a cópia de custódia que gerou até a **central** confirmar o recebimento (ele não a lê: é cifrada para a chave pública de quórum).
   - **Alternativa recomendada pelo Fable e não adotada (01e-E07):** "selar antes de destruir" (o nó sela cada segredo de missão na emissão e entrega o selado ao aparelho; a chave do disco tem cópia selada para o quórum). O responsável preferiu a destruição imediata; a alternativa fica como mitigação a reavaliar (pergunta R3-6 do `01f`).
   - **Registro:** a ordem manual fica no aparelho de quem a assinou e vai à central no retorno; a destruição automática é registrada no retorno, pela ausência do nó e pelo relato do administrador da missão. Painel: "nó de campo destruído em DD/MM/AAAA; dados não sincronizados perdidos", nunca "guardado".

## Alternativas rejeitadas
- **Nó de campo como único committer** (versão anterior): o nó teria segredos de grupo ou autoridade sobre membros (D13).
- **Central prevalece em conflito** (versão anterior): readmite removido; aparelho capturado fala com a central; ramo de campo indecifrável (D14).
- **Lease longo sem nó de campo:** aparelho perdido legível por dias, sem ninguém para revogar.
- **Sem operação em campo:** contraria D-R22.
- **Cada aparelho como servidor (malha sem nó):** bifurcação de época e diretório sem dono.
- **Satélite comercial estrangeiro:** regra 10.
- **Destruição do nó com duas autorizações, como na GD-13:** inviável com captor presente; o responsável decidiu a exceção (GD-18).
- **Selar antes de destruir como regra (01e-E07):** não adotada pela GD-18; mantida como mitigação a reavaliar.

## Consequências
- Novo componente a endurecer e a perder (R89). Capturado sem tempo de destruir, expõe metadados da missão, cifrados ainda não sincronizados e segredos de missão vigentes, que só abrem cache junto com o hardware do aparelho. Não tem segredo de grupo. Destruído (item 13), perde o que não foi sincronizado (R118), por decisão do responsável (GD-18).
- Aparelhos do comandante e do subcomandante viram alvo: capturado desbloqueado, inclui membro (R100). Mitigação: chave de administrador com biometria a cada uso, aviso de inclusão a todos os membros, remoção pelo outro administrador, PIN de coação.
- Janela maior para aparelho perdido em campo (R92), limitada pelo teto monotônico e pela validade da credencial do nó.
- Durante a missão a central não entrega aos grupos de missão: quem ficou na base fala com a missão por outra sala ou por um membro em campo.
- A reconciliação vira importação de cadeia verificada; testar em F-R3 e F-R10.
- O servidor ganha um **perfil de campo**; fronteiras de módulos (ADR-0001) não mudam.

## Revisar quando
Os rádios forem levantados (D-R25), houver medição de banda e bateria (F-R1), D-R39 for decidida, a pergunta 10 do 01c (01c-C17) for respondida, ou o spike MLS mostrar o custo real de commit em banda estreita.
