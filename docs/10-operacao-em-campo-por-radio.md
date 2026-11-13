# 10 — Operação em Campo por Rádio

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** chefia, área de segurança, equipe de comunicações, equipe de desenvolvimento

> **Decisão do responsável pelo produto:** o MIGDOL terá operação em campo por rádio, para a polícia e para o Exército (D-R22, opção B do doc 09). Entra como **Fase 8**, depois da v1 (doc 03 seção 5): não se adianta fase. O desenho exige a **ADR-0004** antes do código, porque muda o lease de 12 h (ADR-05), e depende da ADR-0005 e da ADR-0006. Nada aqui foi implementado ou medido; bandas e tempos são hipóteses.
>
> **Revisão de desenho (etapa 1b da Fase 0):** o nó de campo deixou de fazer commit (D13); a central fica congelada durante a missão (D14, GD-11); remoção em banda estreita passou a ser permitida (D15); o lease de campo virou criptográfico, com segredos de missão (D22).
>
> **Terceira rodada (`docs/fase0/01f-correcoes-rodada3.md`):** nó capturado é destruído de imediato, aceitando perder o que não foi sincronizado, como **exceção escrita à GD-13** só para o nó (GD-18; ADR-0004 item 13); passagem do lease entre central e nó descrita (01e-E12); no aparelho, escada por perfil com destruição imediata em missão (doc 07 seção 7.1a).

---

## 1. Princípio

O aplicativo **não fala com o rádio**. O rádio entrega uma **rede IP local**; o MIGDOL roda sobre ela com um **nó tático de campo** que faz o papel do serviço de entrega. A criptografia ponta a ponta (MLS) não muda: o rádio só transporta cifrado.

```
[Celular] --Wi-Fi/Ethernet--> [Rádio tático IP] ~~~ rádio ~~~ [Rádio] --> [Nó de campo]
                                                                            |
                                                         ao voltar: sincroniza com a instância central
```

---

## 2. Meios de rádio

| Meio | Banda (hipótese) | Modo do MIGDOL | Situação |
|---|---|---|---|
| Rádio tático IP em malha (MANET) | Mbps | **Completo de campo:** texto, foto, anexos pequenos | Alvo principal |
| Wi-Fi do kit de campo | Mbps, alcance curto | Completo de campo | Base avançada |
| VHF/UHF de banda estreita com IP | Poucos kbps | **Banda estreita:** só texto curto, comprimido, quadros de tamanho fixo | Segunda etapa |
| HF | Muito baixa | Fora do escopo | Meios da Força |
| Satélite sob controle do Estado (ex.: SGDC) **[conferir]** | Varia | Enlace de volta à central | A avaliar |
| Satélite comercial estrangeiro, LoRa, malhas amadoras | — | **Não usar** (regra 10; segurança insuficiente) | Vetado |

**Primeiro passo:** levantar **quais rádios** a polícia e o Exército já usam, se entregam IP, a banda real e a cifra do enlace (D-R25).

---

## 3. Componentes

### 3.1 Nó tático de campo
- Servidor portátil endurecido, com o **mesmo binário** do servidor central em perfil de campo: gateway, entrega e diretório local da missão. Sem nó de IA, sem abertura de custódia, sem painel de gestão completo.
- **Não é membro** de nenhum grupo e **não tem chave que altere membros**. Só ordena (um commit por época, ADR-0005) e retransmite cifrado (D13).
- Disco cifrado, chave liberada por PIN do responsável e por hardware. **Capturado: destruição imediata** da chave do disco (GD-18; ADR-0004 item 13), por botão físico, ordem assinada do administrador da missão ou do substituto ou violação do gabinete; perda de enlace, falta de energia e cabo não acionam. **Falha do boot verificado não destrói:** só deixa de liberar a chave do disco, o que é reversível na base, para que atualização corrompida ou defeito não destrua custódia e trilha não sincronizadas (01g-G20; gatilhos: D-R39). É **exceção escrita à GD-13**: um ato só, ou automático, sem duas autorizações. Perde-se o que não foi sincronizado (R118).
- **Nenhuma chave privada de custódia** (ADR-0003) nem cartão de comandante no campo.
- Guarda só o **segredo de missão vigente** de cada aparelho (ADR-0006, D22).
- Retenção curta: guarda só o necessário até sincronizar; depois do sincronismo, apaga.
- Credencial de campo emitida pela AC da instituição, com **validade curta, igual à janela da missão**, verificada pelos aparelhos a cada renovação (D22).

### 3.2 Aparelho
- Liga ao rádio por **Wi-Fi ou Ethernet**, com o rádio na **lista permitida do MDM**. **Não por cabo USB de dados**: o cabo dispara a escada de resposta (doc 07 seção 7.1).
- Perfil de campo obrigatório: câmera, PIN de coação, escada de resposta, sem IA. Durante a missão em andamento, e conforme o perfil da instituição (D-R40; no perfil polícia, desligado por padrão), o aparelho fica no **perfil missão de campo** da escada: rosto de outra pessoa em duas capturas com PIN errado, ou PIN de coação, levam ao nível 4, cada gatilho só depois das medições da GD-24; até lá, nível 3 (GD-18; doc 07 seção 7.1a; 01g-G21).

### 3.3 Grupo de missão
- **Montado na base**, antes de sair: membros, salas e chaves definidos com conexão à central. O **administrador da missão** (comandante da missão) e o **substituto** (subcomandante) ficam registrados como administradores no estado autenticado do grupo (GD-11). Em sala sob custódia, o administrador leva chaves públicas de período assinadas suficientes para a missão (ADR-0003).
- **Durante a missão, só o aparelho do administrador da missão, ou do substituto, altera o grupo** (entrada, saída, atualização). O nó só ordena: um commit por época, na ordem de chegada (ADR-0005). Os aparelhos recusam commit sem assinatura de administrador. Isso evita bifurcação de época MLS quando a rede se parte, sem dar ao nó autoridade sobre membros (D13).
- Remoção de membro em campo (aparelho perdido, capturado) é commit do administrador da missão, serializado pelo nó, e entra na **lista de revogação de campo**, assinada, enviada à central pelo primeiro enlace (D14).

---

## 4. Lease de campo (muda a ADR-05)
- Em campo, o **nó de campo faz o papel do servidor no lease criptográfico** (ADR-0006): entrega a cada aparelho um **segredo de missão** próprio, distinto dos da central, trocado a cada renovação (proposta: 12 h enquanto houver contato com o nó). O nó guarda só o vigente.
- **Teto de 72 h** sem contato (D-R26), medido no aparelho pelo **relógio monotônico** desde o boot, que o usuário não ajusta. Reinício apaga a memória e exige nova entrega. Mudar a hora do aparelho não estende nada (D22).
- Aparelho fora do alcance do nó além do teto: cache ilegível, **não apagado**.
- **Passagem do lease (01e-E12):** na saída, o nó emite a geração de campo pelo protocolo de duas fases da ADR-0006 (item 3), e a central apaga a geração dela depois da confirmação; no retorno, a central emite a geração seguinte e o nó apaga o segredo de missão só depois da confirmação, pelo sincronismo (ADR-0006 item 9).
- Captura do nó de campo: destruição imediata (seção 3.1). Se não houver tempo, os aparelhos deixam de renovar quando a credencial curta do nó vence; a central revoga a credencial ao saber. Aparelho que reiniciar ou passar do teto com o nó destruído fica com o cache ilegível (perda aceita, GD-18).

---

## 5. Sincronização com a central
- **Central congelada (GD-11, D14):** ao sair a missão, a central marca os grupos de missão como congelados. Não aceita commit, não entrega mensagem para esses grupos e **recusa sessão direta dos aparelhos em missão**. O campo é o único escritor até o retorno.
- Ao reconectar (base, satélite autorizado ou rede da instituição), o nó de campo:
  1. envia **primeiro** a lista de revogação de campo (mesmo por enlace de banda estreita); a central revoga certificados e leases, sem destruir: os segredos de missão dos removidos chegam já selados para a chave de quórum (ADR-0006 itens 8 e 9; GD-13);
  2. envia as mensagens cifradas e as **cópias de custódia** já feitas nos aparelhos (ADR-0003);
  3. envia os eventos e a trilha de auditoria da missão;
  4. entrega a **cadeia de commits** da missão; a central verifica cada assinatura e a continuidade das épocas antes de gravar e só então descongela;
  5. apaga o que já foi sincronizado.
- **Sincronizar cedo:** a cada enlace disponível, na ordem acima, para reduzir o que se perde se o nó for destruído (R118). A lista de revogação e a cadeia de commits ficam também no aparelho do administrador da missão, que as assinou (01c-C17, pendente).
- **Não há conflito de grupo:** durante a missão não existe ramo central. A regra anterior ("a central prevalece") foi retirada porque readmitia aparelho removido em campo (D14). Cadeia inválida: o grupo segue congelado e abre SEV1 (F-R3, F-R10).

---

## 6. Modo de banda estreita
- Só texto, mensagens curtas, compressão antes de cifrar **[avaliar: compressão antes de cifrar vaza tamanho; usar quadros de tamanho fixo]**.
- Quadros de tamanho fixo (padding) para não revelar o conteúdo pelo tamanho.
- Sem anexos e sem recibos de visualização em tempo real (enviados depois do sincronismo).
- **Remoção de membro permitida**, com orçamento de banda reservado; commit de atualização periódica espera banda larga ou o retorno (D15). Medir em F-R1 e F-R9.
- Fila com prioridade: alerta, coação e remoção primeiro.

---

## 7. Segurança própria do rádio

| Ameaça | Quem protege | Observação |
|---|---|---|
| **Radiogoniometria** (localizar o emissor) | Rádio: salto de frequência, baixa probabilidade de detecção, potência mínima | Norma de uso: transmitir só o necessário |
| **Interferência** | Rádio e procedimento | App tolera perda, atraso e reenvio |
| **Análise de tráfego** (horário e volume) | App: quadros fixos, tráfego de cobertura opcional | Custa banda e bateria (D-R20) |
| **Interceptação do enlace** | Duas camadas: MLS do MIGDOL + cifra do enlace do rádio | Se a cifra do rádio for algoritmo de Estado, ela **pode ser a segunda camada** do doc 08 (D-R18) **[conferir]** |
| **Nó de campo capturado** | Nó: sem chave de grupo nem de membros; disco cifrado; **destruição imediata** (GD-18); credencial curta | Aparelhos deixam de renovar quando a credencial vence; revogação na central; perde-se o não sincronizado (R118) |
| **Rádio capturado** | Rádio: zeroização de chaves da Força | Fora do MIGDOL |

---

## 8. Riscos novos

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R89 | Nó de campo capturado com mensagens da missão | Médio | Alto | Nó sem chave de grupo nem de membros (D13); disco cifrado; **destruição imediata** quando capturado (GD-18; ADR-0004 item 13); retenção curta; sem custódia privada; credencial curta; revogação | Médio |
| R118 | Destruição do nó capturado, ou por gatilho falso, perde o que não foi sincronizado: cópias de custódia e trilha da missão (contra o RPO 0 da auditoria), segredos selados de removidos, lista de revogação de campo e, para aparelhos que reiniciarem, o cache de missão | Médio | Alto | Sincronizar cedo e a cada enlace, com prioridade; lista de revogação e cadeia também no aparelho do administrador da missão (01c-C17); gatilhos sem perda de enlace, energia ou cabo (D-R39); **[proposta]** aparelho guarda a cópia de custódia que gerou até a central confirmar | **Médio [proposta, com as mitigações]**. Perda admitida pelo responsável pelo produto (GD-18), que não é aceite formal; aceite pelo responsável pela homologação pendente (doc 04 seção 3.2) |
| R90 | Bifurcação de época MLS com rede partida, ou aparelho removido em campo readmitido pela central | Médio | Alto | Um commit por época no nó (ADR-0005); hash do transcript nos dados autenticados, comparado sem decifrar (detecta bifurcação entre aparelhos que ainda trocam mensagens; partição total só no retorno); só o administrador da missão ou o substituto altera o grupo; central congelada; importação de cadeia verificada; lista de revogação de campo (D14) | Baixo |
| R91 | Emissão de rádio revela posição da equipe | Alto | Alto | Rádio com baixa probabilidade de detecção; norma de transmissão mínima | Médio |
| R92 | Lease de campo prolongado aumenta a janela do aparelho perdido | Médio | Alto | Teto de 72 h pelo relógio monotônico; segredos de missão por aparelho; credencial curta do nó (ADR-0006, D22); perfil de campo com controles obrigatórios | Médio |
| R93 | Compressão e tamanho das mensagens vazam informação em banda estreita | Médio | Médio | Quadros de tamanho fixo | Baixo |
| R100 | Aparelho do administrador da missão (ou do substituto) capturado desbloqueado inclui membro do adversário no grupo de missão | Baixo | Crítico | Chave de administrador com biometria a cada uso; aviso de inclusão a todos os membros; o outro administrador remove; PIN de coação; escada de resposta | Médio |

---

## 9. Decisões

| ID | Decisão | Quem decide |
|---|---|---|
| D-R22 | **Decidido: operação em campo por rádio (opção B), Fase 8, para polícia e Exército.** Na missão: central congelada; nó só ordena; administrador da missão = comandante da missão; substituto = subcomandante (GD-11) | Responsável pelo produto |
| D-R25 | Quais rádios e redes (MANET, VHF/UHF, satélite do Estado) e se entregam IP | Comunicações de cada instituição |
| D-R26 | Teto do lease de campo (proposta: 72 h) | Segurança + chefia |
| D-R27 | Cifra do rádio como segunda camada (D-R18) | Segurança, com o órgão que homologa |
| D-R39 | Gatilhos da destruição imediata do nó capturado (GD-18): botão físico, ordem assinada do administrador da missão ou do substituto, violação do gabinete (proposta); o que não aciona (perda de enlace, energia, cabo, queda sem violação, **falha do boot verificado**, que só deixa de liberar a chave do disco; 01g-G20) | Segurança + comando |

---

## 10. Requisitos

| ID | Requisito | Prioridade |
|---|---|---|
| RF78 | Nó tático de campo com entrega local, diretório da missão e sincronização com a central | Fase 8 |
| RF79 | Grupo de missão montado na base; em campo, só o administrador da missão (comandante) ou o substituto (subcomandante) altera o grupo, com o nó só ordenando; central congelada durante a missão; lista de revogação de campo | Fase 8 |
| RF80 | Lease de campo criptográfico: segredos de missão por aparelho, teto pelo relógio monotônico, credencial curta do nó (ADR-0006) | Fase 8 |
| RF81 | Modo de banda estreita: texto curto, quadros de tamanho fixo, fila com prioridade, remoção de membro com orçamento de banda | Fase 8, segunda etapa |
| RNF-R1 | Ligação aparelho–rádio só por Wi-Fi ou Ethernet de rádio permitido pelo MDM | Fase 8 |
| RF97 | Destruição imediata do nó de campo capturado (GD-18; ADR-0004 item 13): gatilhos D-R39; perda do não sincronizado registrada e exibida no painel como perda, nunca como guarda | Fase 8 |
| RNF-R2 | Nó de campo sem chave privada de custódia, sem chave de grupo e sem chave que altere membros; destruição imediata da chave do disco quando capturado | Fase 8 |

---

## 11. Testes (F-R)

| # | Cenário | Esperado |
|---|---|---|
| F-R1 | Missão de 24 h em MANET com 10 aparelhos, texto e foto | Entrega com perda e atraso tolerados; medir banda e bateria |
| F-R2 | Aparelho fora do alcance do nó além do teto | Cache ilegível, não apagado |
| F-R3 | Aparelho removido em campo: a central tenta readmiti-lo, ou um aparelho em missão tenta sessão direta com a central | Recusado (central congelada); no retorno, cadeia importada sem membro indevido |
| F-R4 | Nó de campo capturado e desligado | Disco ilegível; revogação na central; aparelhos param de renovar quando a credencial curta vence |
| F-R5 | Captura de tráfego de rádio | Só cifrado; tamanhos fixos em banda estreita |
| F-R6 | Celular ligado ao rádio por cabo USB de dados | Escada de resposta dispara |
| F-R7 | Sincronização ao voltar | Lista de revogação primeiro; mensagens, custódia e auditoria na central; nó apaga o sincronizado |
| F-R8 | Nó de campo tenta incluir aparelho no grupo de missão | Aparelhos recusam: o nó não tem chave de administrador |
| F-R9 | Remoção de membro em banda estreita | Chega a todos dentro do orçamento de banda; medir |
| F-R10 | Cadeia de commits do campo adulterada no retorno | Central recusa; grupo segue congelado; SEV1 |
| F-R11 | Relógio do aparelho alterado em campo | Teto de 72 h inalterado (relógio monotônico) |
| F-R12 | Nó de campo entrega commits diferentes da mesma época a dois grupos de aparelhos que continuam trocando mensagens | Sinal de bifurcação nos dois lados (ADR-0005 item 2); SEV1 ao administrador da missão |
| F-R13 | Destruição do nó por botão, por ordem assinada e por violação do gabinete | Chave do disco destruída; disco ilegível; no retorno, a central registra "nó destruído" e a perda; lista de revogação e cadeia entregues pelo aparelho do administrador (se 01c-C17 for decidido assim) |
| F-R14 | Perda de enlace, falta de energia, cabo conectado e queda do nó sem violação | Nada é destruído |
| F-R15 | Passagem do lease na saída e no retorno, com queda no meio | O aparelho retoma pela geração gravada; o cache não se perde (ADR-0006 item 9) |

---

## 12. Fase 8 — Operação em campo por rádio (depois da v1)

| Entrega | Critério de aceite |
|---|---|
| ADR-0004, ADR-0005 e ADR-0006 aprovadas | Lease de campo, nó de campo e ordenação de commits decididos |
| Levantamento dos rádios (D-R25) | Banda, IP e cifra de enlace medidos |
| Nó de campo em MANET (etapa 1) | F-R1 a F-R4, F-R6 a F-R8 e F-R10 a F-R15 aprovados |
| Banda estreita (etapa 2) | F-R5 e F-R9 aprovados; modo de banda estreita medido |
| Exercício de campo com a polícia e com o Exército | Relatórios anexados |
