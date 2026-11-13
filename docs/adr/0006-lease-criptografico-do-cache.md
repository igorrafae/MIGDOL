# ADR-0006 — Lease criptográfico do cache local

**Status:** proposta (etapa 1b da Fase 0; achados D21, D22, D23, D26; **segunda rodada**: 01c-C01, C12, C13, C14 e decisão GD-13; **terceira rodada**: 01e-E05, E06, E08, E12, E13, E21c, E21d) · **Contraria** ADR-05 (lease de 12 h como validade lógica, confiada ao app e ao relógio) · **Complementa** doc 02 seção 4.3, ADR-0003 seção 4 (classe "segredo de lease guardado") e ADR-0004 · **Fase 2** (e Fase 8 no campo)

## Contexto
Hoje `chave_cache = HKDF(segredo_hardware, segredo_servidor, "cache-v1")`, e o `segredo_servidor` é sempre o mesmo. A "validade" depende de o app esquecer o segredo e confiar no relógio. Quem capturou o segredo uma vez abre o cache para sempre, junto com o hardware (D21). E o app apaga o cache ao vencer o lease, o que pode ser destruição de prova numa apreensão legal (D23).

A primeira versão desta ADR fazia a **revogação apagar** o segredo vigente, salvo em apreensão declarada. A revisão (01c-C01) mostrou que isso é destruição irreversível do cache feita por uma pessoa, e que aparelho dado como perdido, mas na verdade apreendido, teria o cache destruído. Decisão do responsável (GD-13): **revogar e destruir são atos separados**.

## Decisão (proposta)
1. **Chave do banco aleatória:** a chave do SQLCipher é gerada no core e **não** é derivada direto dos segredos.
2. **Embrulho por lease:** essa chave fica gravada embrulhada (AEAD da pilha do core, regra 1) por `kek_n = HKDF(segredo_hardware, segredo_servidor_n, "cache-lease-v2")`.
3. **Renovação com geração (01c-C12):**
   - cada segredo tem **número de geração** `n`, gravado junto do embrulho;
   - o servidor entrega `segredo_servidor_{n+1}` e guarda **`n` e `n+1`** até a confirmação; pedido repetido de renovação devolve o **mesmo** `n+1` (idempotente);
   - o core desembrulha com `kek_n`, reembrulha com `kek_{n+1}`, grava de forma **atômica** (o embrulho antigo só sai depois de o novo estar gravado), confirma e zera `segredo_servidor_n`;
   - o servidor apaga `n` ao receber a confirmação e volta a guardar **só o vigente**;
   - queda no meio: aparelho e servidor retomam pela geração gravada, sem perder o cache.
4. **Expiração pelo servidor:** passado o prazo sem renovação, pelo **relógio do servidor**, o servidor **para de entregar** o segredo vigente sem nova verificação (sessão autenticada e atestação). Vencer o prazo, sozinho, não destrói nada (D23).
5. **Suspensão (nível 3 da escada, doc 07 seção 7.1):** o servidor **para de entregar** o segredo, mas o guarda. É reversível. Revogar é ato humano registrado (D25).
6. **No aparelho:** o segredo vigente fica só em memória. O aparelho deixa de usá-lo no prazo medido pelo **relógio monotônico** desde o boot (não ajustável pelo usuário). Reinício apaga a memória.
7. **Lease vencido não apaga o cache:** só o deixa ilegível (D23). Apagar é ato separado: ordem assinada ou nível 4 da escada.
8. **Revogar não destrói (GD-13, 01c-C01).** São dois atos:
   - **Revogar a credencial** (certificado, sessão, lease): imediato, uma pessoa, registrado. Em **toda** revogação, qualquer que seja o motivo declarado (perda, roubo, comprometimento, apreensão, desligamento), o servidor **sela todas as gerações que guarda** para o aparelho (`n` e, se houver renovação pendente, `n+1`), com o número da geração no registro (01e-E05): cifra-as para a **chave de quórum do segredo de lease** (ADR-0003 seção 4) e apaga as cópias em claro. Daí em diante o servidor não consegue entregá-lo, e o cache fica ilegível também para aparelho offline **depois que ele esquece o segredo** (relógio monotônico ou reinício; limite declarado abaixo) (01e-E21c).
   - **Revogação em lote** (mais de um aparelho por ato) exige **duas autorizações**, como o apagamento em lote (R71). Em caso de dúvida, o padrão é **suspender** (nível 3, reversível), não revogar (01e-E08a).
   - **Rastro do selado (01e-E08b):** no ato da selagem, o hash do registro selado vai à trilha de auditoria. O Auditor confere periodicamente que todo hash da trilha tem registro no `custody` (ou no `directory`, até a Fase 4; item 11): exclusão por administrador de banco fica visível. Backup com RPO 0 (item 10).
   - **Destruir o segredo selado:** ato separado, com **duas autorizações** de pessoas distintas e registro, como o apagamento remoto (doc 07 seção 7.4). Com indício de apreensão por autoridade, só depois da decisão D-R14.
   - **Prazo (01e-E06):** o segredo selado fica guardado pelo **prazo da apuração** do evento que motivou a revogação, fixado no registro da revogação. O `instituicao.yaml` traz o **prazo padrão** e um **piso** (D-R35; proposta do 01e: padrão de 5 anos, piso de 90 dias **[conferir]**); prazo abaixo do padrão só com as duas autorizações da destruição, para que quem revoga não transforme um prazo curto em destruição por uma pessoa. O selado vai para a chave do **período de vencimento** (mês em que o prazo termina), e a eliminação por período fica exata (ADR-0003 seções 4 e 7). **Prorrogação** é ordem de preservação sobre a chave do período (RF58; ADR-0010 seção 8), declarada como tal.
   - **O que ele abre:** sozinho, nada. Abrir o cache exige o segredo selado **e** o aparelho físico **e** o desbloqueio do agente (segredo de hardware liberado pelo Keystore). Por isso o segredo selado não cria custódia de sala comum: a instituição só lê o cache de um aparelho que tenha nas mãos, por quórum presencial com base registrada (regra 14).
   - **Volta:** aparelho achado ou revogado por engano volta ao uso por **revinculação** e `history_share` (doc 07 seção 7.4), com nova credencial, e não por cerimônia; o cache antigo continua selado (01e-E08a). **Reabrir o cache antigo** (apuração) é só por cerimônia de quórum com base registrada: o segredo aberto é entregue ao core do aparelho presente, sem passar pelo `server/` em claro (meio a definir na Fase 2 **[proposta]**).
   - **Painel e recibos:** "aparelho revogado: cache ilegível, segredo guardado sob quórum até DD/MM/AAAA". Nunca "destruído" enquanto o segredo selado existir (D27). O painel distingue "registro excluído" (destruição antecipada, que não alcança backups) de "chave do período eliminada" (destruição criptográfica) (01e-E06d).
   - **Reconexão de aparelho revogado:** recebe primeiro as ordens pendentes (expurgo, apagamento) e nunca o segredo.
9. **Em campo:** o nó de campo faz o papel do servidor, com segredos de missão distintos e teto pelo relógio monotônico (ADR-0004). Na remoção em campo vale o item 8: o nó sela o segredo de missão do aparelho removido com a chave **pública** vigente da classe (do lote assinado em cerimônia, levado da base) e apaga a cópia em claro. O nó nunca tem chave privada de quórum.
   **Passagem do lease entre central e nó (01e-E12):** usa o mesmo protocolo de duas fases do item 3, com o emissor trocado.
   - **Saída:** ainda na base, com contato com a central e com o nó, o nó emite a geração `n+1` (segredo de missão); o aparelho reembrulha, grava de forma atômica e confirma ao nó, que registra a geração e repassa a confirmação à central; a central apaga `n`, como numa renovação comum (item 3), porque o embrulho antigo já saiu do aparelho. Queda no meio: o aparelho retoma pela geração gravada, e a central guarda `n` até a confirmação.
   - **Retorno:** a central emite `n+k+1`; o aparelho, ainda com o segredo de missão em memória, reembrulha e confirma à central; o nó apaga o segredo dele só depois de a confirmação chegar pelo sincronismo (doc 10 seção 5).
   - **Aparelho que reiniciou ou passou do teto em campo:** fica ilegível até nova entrega pelo nó; com o nó destruído, até a reabertura por cerimônia, se houver selado, ou para sempre (perda aceita pela GD-18; ADR-0004 item 13).
10. **Continuidade (01c-C12):** os segredos de lease vigentes e selados seguem **RPO 0**, como a auditoria (doc 04 seção 8) **[proposta]**. Se RPO 0 não for viável, alternativa: reter `n−1` por janela maior que o RPO do diretório, com o custo declarado (segredo antigo vale mais tempo).
11. **Fase e dono dos selados (01e-E13) [proposta; pergunta Q5 do 01e]:** a RF91 e o A74 são da Fase 2, mas o serviço `custody` e a cerimônia do HSM são da Fase 4. Na **Fase 2**, o `directory` sela com a chave **pública** da classe lida do pacote assinado (chave pública não é segredo; chave de teste até a cerimônia) e guarda o selado em **tabela própria**. Na **Fase 4**, o `custody` **busca** os selados no `directory` pelo `api/` (sentido permitido do grafo: `custody` → `directory`), ou a tabela continua no `directory`. Nenhum ciclo `directory` → `custody` (regra 9; doc 02 seção 7.1).

## Limites declarados (01c-C13)
- **Chave do banco fixa:** quem copiou a chave do banco da memória do aparelho, ou guarda um embrulho antigo junto com o segredo daquela geração e o hardware, abre o banco também depois das renovações. A troca periódica da chave do banco (rekey do SQLCipher) será medida na Fase 2.
- **Offline, o prazo continua sendo o app esquecer:** o prazo do servidor só força nova atestação. Sem contato, quem esquece o segredo é o app, pelo relógio monotônico; root ou instrumentação contornam isso. É inerente ao uso offline; a defesa é a atestação a cada sessão e a escada de resposta.
- **Backup do banco do diretório:** segredos vigentes antes de uma revogação existem em claro nos backups feitos antes dela. A selagem protege contra o servidor vivo, não contra backup antigo. O backup do diretório é isolado, cifrado e de acesso restrito (doc 04 seção 8, R32).

## Alternativas rejeitadas
- **Segredo fixo com validade lógica (ADR-05):** uma captura vale para sempre.
- **Validade pela data de fim da chave no Keystore:** depende do relógio do aparelho, que o adversário controla.
- **Apagar o cache ao vencer o lease:** destruição automática durante apreensão (D23).
- **Revogar apaga o segredo, salvo apreensão declarada** (primeira versão): destruição irreversível por uma pessoa, dependente do motivo declarado (01c-C01).

## Consequências
- Uma gravação a mais por renovação (reembrulho), com geração e gravação atômica.
- Servidor indisponível além do prazo deixa o cache ilegível até voltar. Perder o segredo vigente no servidor torna o cache ilegível; por isso RPO 0 (item 10).
- Processo encerrado pelo Android perde o segredo da memória e só reabre com contato. Em campo, isso pesa (D26): decidir no protótipo de D-R1 se há serviço em primeiro plano.
- Nova classe sob quórum: **segredo de lease guardado** (ADR-0003 seção 4; regra 14).
- Testes: A21, A22 (doc 03), A73 e A74 (doc 07); queda no meio da renovação e restauração do banco (Fase 2; V05).
- Riscos: R99 e **R110** (segredo guardado mais aparelho e desbloqueio abrem o cache; acervo de selados é alvo) (01e-E21d).

## Revisar quando
Houver medição de D-R1 e D26 no protótipo, o spike de atestação mostrar como o core chama o Keystore, ou o jurídico decidir D-R14.
