# Correções da revisão de desenho — Fase 0, etapa 1b

**Modelo:** Claude Opus 5.5 (pipeline da Fase 0, etapa 1b). **Revisão seguinte:** Fable 5.2.
**Entradas:** `01-revisao-desenho.md` (achados D01–D34, com a nota de conferência do orquestrador) e `GD-respostas.md` (decisões do responsável, que **prevalecem** sobre as propostas da revisão).
**Limites:** nenhum código, commit ou push. Nada enviado a serviço externo. Fato jurídico não conferido está marcado **[conferir]**. Conteúdo substituído, não apagado sem substituto. Este relatório não aceita risco em nome de ninguém: "aceito" só onde o responsável aceitou no GD.

## Resumo

| Severidade | Total | Fechado | Aceito | Adiado |
|---|---|---|---|---|
| Crítica | 5 | 5 | 0 | 0 |
| Alta | 15 | 12 | 0 | 3 (D08, D28, D30) |
| Média | 13 | 11 | 0 | 2 (D26, D31) |
| Baixa | 1 | 1 | 0 | 0 |
| **Total** | **34** | **29** | **0** | **5** |

- A revisão falava em "16 altos"; a tabela dela tem **15** achados altos. Conferido achado a achado.
- "Fechado" quer dizer **fechado no texto**. Vários dependem de prova no spike MLS (etapa 3) ou de medição (marcados na nota). O risco residual de D01 (HSM ou quórum comprometido expõe o período) foi **aceito pelo responsável** no GD-1, só em sala sob custódia.
- Nenhum crítico ficou aberto. Altos não fechados: D08, D28, D30, com justificativa na tabela.

## 1. Achados D01–D34

| ID | Achado (resumo) | Sev. | Status | Onde foi corrigido | Nota |
|---|---|---|---|---|---|
| D01 | Chave simétrica de custódia com os membros: um aparelho comprometido abre o período inteiro | Crítica | Fechado | ADR-0003 §2; doc 07 §7.3; doc 01 RNF-A4, RN17; doc 02 §5.3; `CLAUDE.md` regra 14 | HPKE no remetente para **chave pública** do período; privada só no HSM offline; membros não decifram. Residual (HSM ou quórum) aceito pelo GD-1, só em sala sob custódia (R68) |
| D02 | Troca de chave na remoção só como exemplo; chave errada embrulhada sem ninguém notar | Crítica | Fechado | ADR-0003 §1–2; doc 02 §5.2 | Período muda a cada commit de membros (e no máximo 30 dias, **[proposta]**); chave pública no estado autenticado do grupo, assinada pela chave de cerimônia; lote gerado no HSM offline |
| D03 | Política do HSM online depende do diretório hostil; registro das chaves de aprovação indefinido | Crítica | Fechado | ADR-0003 §4 e Alternativas; doc 07 §7.4; regra 14 | **Removido da v1** (GD-4): sem HSM online nem aprovação remota. Se voltar, exige nova ADR com D03 (a)–(e) e D05 |
| D04 | Mesma chave embrulhada para HSM online e offline: vale o mais fraco | Alta | Fechado | ADR-0003 | Não há chave de recuperação nem HSM online de custódia |
| D05 | Aprovação remota sem tela confiável, falha comum, biometria e PIN, coação | Alta | Fechado | ADR-0003; doc 07 RF77 (retirado), A61–A64 (não se aplicam), R81–R83 (não se aplicam) | Removido da v1 (GD-4) |
| D06 | Cópia de custódia não verificável nem autenticada | Alta | Fechado | ADR-0003 §3; doc 02 §5.3; doc 07 A65, A68, R97 | Hash da cópia dentro da mensagem MLS; assinatura do remetente **dentro** da cópia (não revela remetente ao servidor); servidor recusa envelope sem cópia. Residual R97 (remetente cifra lixo na própria cópia): **aceite pendente do responsável** |
| D07 | "Destruir chave embrulhada" é apagar blob (WAL, réplicas, backups) | Alta | Fechado | ADR-0003 §7; doc 02 §6.3; doc 04 §6.4 | Chaves privadas só no HSM; eliminação destrói no HSM **e nos backups do HSM**, com recibo; inventário de backups do HSM; backup do acervo cifrado deixa de importar após a destruição. Apagar chave dentro de backup: **[conferir no modelo de HSM]** (pergunta 14) |
| D08 | Acervo de 5 anos sob HPKE clássico; sem segunda camada de Estado | Alta | **Adiado** | ADR-0003 §2 e Consequências; doc 08 §4.1 | KEM híbrido na custódia quando a pilha HPKE tiver suíte estável (spike e Fase 1, D-R12); X-Wing ainda experimental (`04-construir-x-adaptar.md`). Segunda camada de Estado e grau classificado só depois de G3 (GD-9). R78 segue aberto |
| D09 | Matrícula ↔ codinome ↔ aparelho ↔ sala em claro no diretório | Alta | Fechado | doc 02 §4.2, §7.5; doc 05 §3.4; doc 06 §3; doc 04 R40; ADR-0003 §4 | Matrícula cifrada para a chave de quórum de pseudonimato; índice de consulta com chave fora do `server/`. Custo para o gestor (pergunta 10) **não respondido**: confirmar |
| D10 | Posição em perda e imagem de coação sob chave judicial: inviável em urgência | Média | Fechado | doc 07 §7 (pânico), §7.4 (sala de resposta); doc 05 §9 | A proposta da revisão (chave operacional de dupla custódia) **contrariaria a regra 14**; adotada alternativa dentro da regra: **sala de resposta** MLS por aparelho, dormente até perda ou coação; posição guardada só para quórum. Quem e quantos (pergunta 13): pendente |
| D11 | Nenhum serviço é dono da custódia, recibos, evidências, pedidos | Média | Fechado | doc 02 §7.1, §7.2, §7.5, §6.3; doc 01 RN02 | Serviço `custody` planejado (depende de directory e audit; delivery depende dele). Script `check-boundaries.sh` e `server/CLAUDE.md` mudam **com o código**, na Fase 4 (regra: não implementar agora). ADR-0001 não lista serviços: não precisa mudar |
| D12 | Estação isolada é segundo lugar com texto em claro | Média | Fechado | ADR-0003 §6; doc 02 §1 (princípio 2); doc 03 §1.2, §1.3 | Requisitos da estação e inclusão no modelo de ameaças |
| D13 | Nó de campo como único committer: membro ou autoridade sobre membros | Crítica | Fechado | ADR-0004 §1–3; doc 10 §3.1, §3.3; RF79, RNF-R2; F-R8 | Nó só ordena (um commit por época) e retransmite; commits só do administrador da missão (comandante) ou do substituto (subcomandante), GD-11 |
| D14 | "Central prevalece" readmite removido; capturado fala com a central | Crítica | Fechado | ADR-0004 §4–6; doc 10 §5; F-R3, F-R10; R90 | Central congelada (GD-11); recusa sessão direta dos aparelhos em missão; importação de cadeia verificada; lista de revogação de campo pelo primeiro enlace |
| D15 | Sem commits de atualização nem remoção em banda estreita | Alta | Fechado | ADR-0004 §9; doc 10 §6; doc 07 §5; F-R9 | Remoção permitida com orçamento de banda; atualização espera banda larga. Medir em F-R1 e F-R9 |
| D16 | Bifurcação sem detecção definida; corridas; diretório diverge do MLS | Alta | Fechado | **ADR-0005 (nova)**; doc 02 §5.2, §5.3, §5.5, §7.5; doc 07 §5; doc 03 R11; doc 04 R98 | Um commit por época; resumo autenticado do estado em toda mensagem; chaves de administrador no estado do grupo; declaração de membros verificada antes de gravar. **A provar no spike MLS** |
| D17 | Aparelho que zerou chaves não se remove; ninguém designado | Alta | Fechado | doc 07 §7.1 (níveis 3 e 4); A52 | Administrador ou substituto remove em até 1 h **[proposta]**, com alerta; aviso assinado de autoexclusão antes de zerar (suporte da biblioteca **[verificar no spike]**); nível 3 remove e é reversível por reinclusão |
| D18 | Segunda conexão derruba ambas e abre SEV1: falsos positivos | Média | Fechado | doc 07 §5; A25 | Substituição de sessão com contador assinado; SEV1 só com uso simultâneo |
| D19 | Nó de IA sem atestação Android; chave de administrador aprendida pelo diretório | Média | Fechado | doc 07 §5; doc 02 §4.1, §5.2; ADR-0005 §3 | Credencial de serviço em cerimônia (TPM/HSM); chaves de administrador e de papéis no estado do grupo. **Resta:** salas "compartimento" sem IA por padrão (pergunta 15) |
| D20 | Diretório comprometido emite credencial fantasma | Média | Fechado | doc 02 §4.2; doc 07 §8, §9 | Estação do gestor confere a atestação e um código curto; transparência de chaves antes do piloto. Implementação na Fase 1 |
| D21 | `segredo_servidor` fixo: captura única abre o cache para sempre | Alta | Fechado | **ADR-0006 (nova)**; doc 02 §4.3, §11; doc 03 Fase 2; doc 04 R99 | Chave do banco reembrulhada a cada renovação; servidor guarda só o vigente; prazo pelo relógio do servidor e, no aparelho, pelo relógio monotônico |
| D22 | Nó de campo guarda leases de todos; revogação não chega; teto por relógio | Alta | Fechado | ADR-0004 §7–8; doc 10 §3.1, §4; R92; F-R11 | Segredos de missão por aparelho; credencial curta do nó verificada pelos aparelhos; teto monotônico; destruição do nó apaga primeiro os segredos |
| D23 | Três destruições automáticas durante apreensão legal | Alta | Fechado | doc 02 §4.3; doc 07 §7, §7.1, §7.4; ADR-0006 itens 4, 7, 8; doc 04 D-R14, R63; doc 07 R66, A72 | Lease vencido não apaga (nem o servidor apaga o segredo por prazo); tentativas e pânico só travam até decisão escrita do jurídico; apagamento remoto com duas autorizações; em apreensão, segredo de lease guardado para quórum. **D-R14 continua com o jurídico** |
| D24 | Recibo antes de zerar não prova; MDM fora da ordem; "apagar tudo" × "rastrear" | Média | Fechado | doc 07 §7.4; doc 02 §6.2, §8.2 | Recibo de intenção; ordem verificada no core; MDM com duas pessoas e só a partir da ordem; apagamento em etapas |
| D25 | Nível 3 automático revoga; §7.2 diz que revogar é humano | Média | Fechado | doc 07 §7.1; ADR-0006 item 5 | Automático só suspende; revogar é ato humano registrado |
| D26 | Android mata o processo: segredo só em memória perde o cache em campo | Média | **Adiado** | doc 02 §4.3; ADR-0006 Consequências | Decidir no protótipo de D-R1 (serviço em primeiro plano) e medir; segredo embrulhado com prazo no Keystore foi rejeitado (depende do relógio do aparelho) |
| D27 | Com custódia, expurgo afirma destruição que não ocorreu | Alta | Fechado | ADR-0003 §7; doc 01 RN03, RN04, RF52, §3.1, §8.6; doc 06 §4, §7, §8.2; doc 02 §6.3; anexo termo item 2 | Conclusão expurga a cópia de trabalho; eliminação da custódia no fim do prazo, com recibo; painel e recibos nunca dizem "destruído" para o que segue guardado; marca de custódia no estado do grupo. Exceção de compartimento virou D-R29 (até lá, não existe) |
| D28 | Prazos conflitantes de registro de acessos e visualização; mesma chave 5 × 15 anos; negações; `actor` em claro | Alta | **Adiado** | ADR-0003 §4 (parte); doc 08 §2.3; doc 04 §5.1 | Feito: uma chave de quórum por classe de dado, com prazo próprio. Adiado para **antes da Fase 4**, com D-R4 e D-R5: hierarquia por grau, registro de acesso gerado no servidor e cifrado para quórum, negações, `audit_log.actor`. Justificativa: sem grau classificado até G3 (GD-9), o prazo de 15/25 anos não se aplica ainda; R60 já bloqueia a Fase 4 |
| D29 | `seq` sem dono; janela de tempo offline; revinculado perde `seq`; ordem sem alvo | Média | Fechado | doc 02 §6.2; ADR-0004 §11; doc 07 §7.4 | Único escritor; aparelho usa só assinatura e contador; `seq` no estado do grupo; formato `MT-WIPE-v1` com aparelho-alvo |
| D30 | Regra 2 literal não vale para StrongBox, HSM, CA, nó de IA; Keystore passa por Java | Alta | **Adiado** | doc 02 §4.1 (escopo proposto); doc 03 Fase 0 (critério do spike) | A tarefa autorizou mudar só a regra 14. Texto proposto para a regra 2 na seção 3 abaixo, **não aplicado**: depende do responsável. O spike MLS (etapa 3) prova o desembrulho pelo Keystore chamado do core. A aprovação do comandante no StrongBox, um dos casos citados, deixou de existir |
| D31 | "ai-node não abre conexão de saída" é impossível | Média | **Adiado** | doc 02 §9.1 (proposta) | Texto proposto para a regra 7 na seção 3, **não aplicado**: depende do responsável |
| D32 | R52 diz "Auditor"; cópia das regras no doc 03 §6 desatualizada | Média | Fechado | doc 07 R52; doc 03 §6 | Conforme a nota do orquestrador, a regra 12 do `CLAUDE.md` já dizia "chave de quórum". Doc 03 §6 agora aponta para o `CLAUDE.md` em vez de copiar |
| D33 | Ordem D-R17/D-R18 × spikes; spike MLS não cobre custódia, commits, remoção offline, Keystore | Média | Fechado | doc 08 §1, §2.1, §9; doc 03 §4.11 e Fase 0; doc 04 §4, §6.1 | Ordem decidida pelo GD-9 (G3, depois dos spikes; nenhum grau classificado até lá). Critério do spike ampliado no doc 03 (o `PIPELINE.md` já tinha a nota) |
| D34 | Numeração de ADRs colide; ADR-0002 sem arquivo; doc 08 §1 desatualizado | Baixa | Fechado | `docs/adr/README.md`; doc 02 §11; doc 08 §1 | Tabela de correspondência; decisões novas usam quatro dígitos. Arquivo da ADR-0002 fica para **antes do spike de atestação** (etapa 5), como a revisão propôs |

## 2. Contradições entre documentos

| # | Contradição (revisão) | Situação | Onde |
|---|---|---|---|
| 1 | Chave de custódia: doc 07 §7.3 (pública) × ADR-0003 (com os membros) | **Resolvida**: os dois dizem HPKE para chave pública, privada no HSM offline | ADR-0003 §2; doc 07 §7.3 |
| 2 | Sigilo futuro: RNF-A4 e doc 03 cenário D02 × ADR-0003 | **Resolvida**: RNF-A4 declara a exceção da sala sob custódia; os membros mantêm sigilo futuro | doc 01 RNF-A4; ADR-0003 Consequências |
| 3 | ADR-0003 rejeita "três celulares" e adota aprovação por três celulares | **Resolvida**: sem aprovação remota na v1 | ADR-0003; regra 14 |
| 4 | Nó de campo faz commits × diretório fora dos grupos, princípio 1, regra 3 | **Resolvida**: o nó só ordena | ADR-0004; doc 10 §3 |
| 5 | Commit a cada 24 h × "aparelhos só enviam" em campo | **Resolvida**: administrador da missão faz commit; banda estreita só remoção | doc 07 §5; doc 10 §6 |
| 6 | Lease vencido apaga o cache × nível 3 "nada é destruído" e D-R14 | **Resolvida**: lease vencido não apaga; o servidor também não apaga o segredo por prazo | doc 02 §4.3; ADR-0006 |
| 7 | Revogação automática × decisão humana | **Resolvida**: automático só suspende | doc 07 §7.1 |
| 8 | Posição em perda para "Auditor e gestor" × quórum; R52 "Auditor" | **Resolvida**: tempo real pela sala de resposta (MLS); guardada só para quórum; R52 corrigido | doc 07 §7.4, R52 |
| 9 | Retenção dos registros de visualização e de acessos | **Em parte**: chaves separadas por classe; prazos adiados (D28) para antes da Fase 4 | ADR-0003 §4; doc 08 §2.3 |
| 10 | O que resta após o expurgo: RN02, RN04, doc 02 §6.3 × custódia | **Resolvida** | doc 01 RN02, RN04; doc 02 §6.3 |
| 11 | Pseudonimato "só por quórum" × matrícula em claro | **Resolvida** no texto; custo para o gestor a confirmar (pergunta 10) | doc 02 §7.5 |
| 12 | Texto em claro só no nó de IA × estação isolada | **Resolvida** | doc 02 §1 |
| 13 | Regra 7 × doc 02; regra 2 × doc 02 e ADR-0003 | **Não resolvida no `CLAUDE.md`**: textos propostos (seção 3), dependem do responsável. A parte da regra 2 ligada à ADR-0003 (chave do comandante no StrongBox) sumiu | doc 02 §4.1, §9.1 |
| 14 | Apagar perfil × aparelho inteiro; "apagar tudo" × "rastrear" | **Resolvida**: aparelho inteiro, apagamento em etapas | doc 07 §7.4; doc 02 §8.2 |
| 15 | Ordem da Fase 0 (doc 08 × `PIPELINE.md`); numeração de ADRs | **Resolvida** | doc 08 §9; `docs/adr/README.md` |
| 16 | (citada na nota do orquestrador) | **Inválida**: a regra 14 existe no `CLAUDE.md`; a lista da revisão já não a traz | — |

## 3. Mudanças em `CLAUDE.md`

### Aplicada: regra 14

Antes:
> 14. Membros leem as salas pelo MLS. Todo acesso institucional a dado cifrado (custódia, evidência, registros, localização, pseudonimato) exige quórum de no mínimo 3 comandantes, via HSM (docs/adr/0003). Nenhuma pessoa decifra sozinha. Divisão de segredo nunca é feita à mão. Só a revinculação usa quórum remoto (app do comandante, biometria do sistema + PIN); abrir dado para outra pessoa é sempre presencial, com cartões, em HSM offline.

Depois:
> 14. Membros leem as salas pelo MLS e nunca decifram cópia de custódia. Cópia de custódia só existe em sala marcada sob custódia na criação (flag assinada no instituicao.yaml; docs/adr/0003): HPKE para chave pública, privada só no HSM offline. Sala comum é MLS puro e é expurgada ao concluir. Todo acesso institucional a dado cifrado (custódia, evidência, registros, termo, localização, pseudonimato) exige quórum presencial de no mínimo 3 comandantes, com cartões, em HSM offline, e base registrada (ordem judicial, ou CI em processo administrativo com validação jurídica). Nenhuma pessoa decifra sozinha. Divisão de segredo nunca é feita à mão. Nenhuma chave de quórum fica em HSM online, e não há aprovação remota por quórum. A revinculação não decifra dado da instituição: aprovação dupla (chefe da inteligência + superior da área) e histórico enviado por outro membro pelo MLS.

Mantido, como pedido: nenhuma pessoa decifra sozinha; acesso institucional só por quórum ≥ 3 presencial em HSM offline; divisão de segredo nunca à mão. A redação diz "nenhuma chave de quórum em HSM online" (e não "nenhum HSM online") porque a CA intermediária já usa HSM online (doc 04 §6.4).

### Propostas, não aplicadas (dependem de autorização do responsável)

- **Regra 2 (D30):** "Chaves privadas e de sessão de membro (MLS, assinatura, cache) existem apenas em core/ (Rust) e nunca atravessam a ponte para Kotlin ou TypeScript. Chaves de hardware (StrongBox/TEE) não saem do hardware: o Kotlin só manipula a referência, e o uso é chamado a partir do core. Chaves da instituição (AC, quórum, cerimônia) ficam em HSM, fora do app."
- **Regra 7 (D31):** "ai-node/ não grava arquivo em claro e só abre conexão para destinos fixos: gateway (mTLS), serviço de anexos da instância e PostgreSQL da IA. Qualquer outro destino é bloqueado e testado em tests/security (E01)."

### Dúvidas registradas sobre regras que não mudam (D-R30)

O GD-4 pede termo de revinculação **com foto**, exportado ao sistema corporativo. Duas tensões, tratadas pela opção mais restritiva até o responsável decidir:
- **Regra 13** (dado de agente cifrado, acesso por rito): o documento exportado leva o **hash da foto**, não a foto; a foto fica só no MIGDOL, cifrada para a chave de quórum.
- **Regra 12** (câmera: sem alerta, nenhum quadro gravado nem enviado): a captura da foto pelo próprio agente foi lida como fora da câmera de vigilância, como a foto de anexo (RF13). Pede confirmação; se a leitura não valer, a foto do termo precisa de outro meio (ex.: estação do gestor).

## 4. Pendências para a revisão do Fable

**Leituras do GD feitas nesta etapa (confirmar):**
1. Flag de custódia ligada ⇒ **toda** sala nova nasce sob custódia; sem exceção até D-R29 (ADR-0003 §1).
2. Foto do termo fora da exportação e leitura da regra 12 (D-R30, seção 3).
3. "Três pessoas distintas" na conclusão: quem assina quando o solicitante é o comandante ou o subcomandante (doc 01 §11, item 12).

**Desenho novo a revisar:**
4. ADR-0005 inteira (um commit por época, resumo de estado, declaração de membros × commit público).
5. ADR-0006 inteira, em especial: vencimento só para de entregar (não apaga); segredo de lease guardado para quórum em apreensão; teto pelo relógio monotônico.
6. ADR-0003 §2: lote de chaves públicas pré-geradas no HSM offline, assinadas pela chave de cerimônia (esgotamento, reuso entre salas, auditoria da atribuição); assinatura do remetente por dentro da cópia.
7. ADR-0003 §5: HSM faz a decapsulação HPKE internamente, ou a chave do escopo vai à estação isolada só durante a cerimônia (pergunta 14).
8. ADR-0004 §4: recusar sessão direta dos aparelhos em missão **para qualquer sala** (custo operacional para quem ficou na base).
9. Sala de resposta (D10, doc 07 §7.4): MLS por aparelho, dormente; alternativa à "chave operacional" que quebraria a regra 14.
10. Histórico da revinculação: guardar o conteúdo autenticado para manter a assinatura original; suporte a autoexclusão (SelfRemove) na biblioteca. Ambos para o spike.

**Riscos alterados sem aceite formal (doc 04 §5.1):**
11. R66 baixado de Alto para Médio até D-R14 (aceite do comitê).
12. R68 Médio: aceito pelo responsável no GD-1; formalizar pela seção 3.2 do doc 04.
13. R97 (cópia inválida do remetente) e R100 (aparelho do administrador da missão capturado): aceite pendente.

**Adiados com dono:**
14. D08 (KEM híbrido na custódia): spike e Fase 1, D-R12. D28: antes da Fase 4, com D-R4 e D-R5. D26: protótipo D-R1. D30 e D31: responsável (regras 2 e 7). ADR-0002: antes da etapa 5. `check-boundaries.sh` e `server/CLAUDE.md`: com o código do `custody`, Fase 4.

**Perguntas da revisão ainda abertas:** 4 (D-R24, detentores; comandante membro ou gestor), 5 (D-R14), 6 (quem autoriza a eliminação definitiva, D-R5), 7 (D-R29), 8 (D28), 10 (custo do D09), 12 (nó capturado ligado), 13 (sala de resposta: quem e quantos), 14 (modelo de HSM), 15 (IA em sala secreta ou compartimento). A pergunta 3 foi respondida pelo GD-4: volta só o que outro membro tem em cache, dentro da retenção.

**Valores marcados como proposta:** período de custódia de no máximo 30 dias; remoção em até 1 h nos níveis 3 e 4; 3 de 5 cartões; renovação de 12 h e teto de 72 h em campo.

**Jurídico [conferir]:** CI em processo administrativo como base para abrir custódia; corregedoria sem ordem judicial; prazo de 5 anos; autorização arquivística para eliminar (Lei 8.159); base para foto no termo (LGPD art. 11) e nível da assinatura no app (Lei 14.063); textos novos do termo de ciência (anexo, itens 2, 9, 10 e 12).

## 5. Arquivos alterados

| Arquivo | O que mudou |
|---|---|
| `CLAUDE.md` | Regra 14 reescrita |
| `docs/adr/0003-custodia-institucional-com-quorum.md` | Reescrita (D01–D07, D12, D27; GD-1 a GD-6) |
| `docs/adr/0004-no-de-campo-e-lease-de-campo.md` | Reescrita (D13–D15, D22, D29; GD-11) |
| `docs/adr/0005-um-commit-por-epoca-e-bifurcacao.md` | **Nova** (D16) |
| `docs/adr/0006-lease-criptografico-do-cache.md` | **Nova** (D21, D23) |
| `docs/adr/README.md` | Lista e correspondência de numeração (D34) |
| `docs/01-negocio-requisitos.md` | Objetivo, atores, RN02–RN04, RN17–RN20, RF52, RF56, RF82–RF84, RNF-A4, §8.6, §9, §11 |
| `docs/02-arquitetura-engenharia.md` | §1, §4.1, §4.2, §4.3, §5.2–§5.5, §6.2–§6.4, §7.1, §7.2, §7.5, §8.2, §9.1, §10.1, §11 |
| `docs/03-seguranca-homologacao-roadmap.md` | §1, R11, §4.4, §4.9, §4.11, §5 (Fases 0, 2, 4, 8), §6 |
| `docs/04-gestao-de-riscos.md` | §4, R40, R63, R98, R99, §5.1 (novo), §6.1, §6.2, §6.4, §10 |
| `docs/05-vazamentos-incidentes-e-atribuicao.md` | §3.3, §3.4, P2, P11, §9 |
| `docs/06-lgpd-e-conformidade.md` | §3, §4, §6.3, §7, §8.2, §8.3, RL11, RL12, §14 |
| `docs/07-resistencia-a-clonagem-e-endurecimento-avancado.md` | §5, §7, §7.1, §7.3, §7.4, §8, §9, riscos, A25, A52–A73, requisitos, D-R15, D-R29, D-R30, §14, §15 |
| `docs/08-informacao-classificada-e-nivel-inteligencia.md` | §1, §2.1, §2.3, §4.1, §9 |
| `docs/09-perfil-exercito-e-modelo-de-implantacao.md` | §1, §3, §4 |
| `docs/10-operacao-em-campo-por-radio.md` | Nota, §3.1, §3.3, §4, §5, §6, §7, riscos, D-R22, requisitos, F-R2–F-R11, §12 |
| `docs/anexos/termo-de-ciencia-e-autorizacao.md` | Itens 2, 9, 10, 12 e seção 4 |
| `docs/anexos/validacao-juridica-preliminar.md` | §3.7, linha de ordem para entregar conteúdo |

**IDs novos** (conferidos sem colisão): RN17–RN20, RF82–RF84, RL11–RL12, R96–R100, D-R29–D-R30, A65–A73, F-R8–F-R11, ADR-0005, ADR-0006. **IDs mantidos e marcados**: RF77 (retirado na v1), R81–R83 (não se aplicam na v1), A61–A64 (não se aplicam na v1).
