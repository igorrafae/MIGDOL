# Revisão de desenho — Fase 0, etapa 1

**Modelo:** Claude Fable 5.2 (pipeline da Fase 0, etapa 1). Leitura dos documentos; nada foi alterado pela revisão.
**Conferência do orquestrador:** D01, D09 e D23 conferidos no texto (ADR-0003 §1, doc 02 §7.5 `registration`/`codename`, doc 02 §4.3). **Correções ao relatório:** D32 está certo quanto ao R52 (doc 07), mas a regra 12 do `CLAUDE.md` já diz "chave de quórum"; e a contradição 16 não procede: a regra 14 existe no `CLAUDE.md` (a cópia no doc 03 §6 é que está desatualizada).
Afirmações sobre Android Keystore, biometria e HSM vêm de conhecimento geral do revisor, não de teste: confirmar nos spikes.

## Veredito

A base (MLS por biblioteca, servidor só com cifrado, expurgo por chave, nó de IA como membro declarado) continua de pé. **A ADR-0003 e a ADR-0004, como escritas, não devem ser aceitas.**

- **ADR-0003:** a chave de custódia simétrica, conhecida por todos os membros e válida por um período, anula sigilo futuro e recuperação pós-comprometimento para tudo que está no arquivo. A política do HSM de recuperação online depende de dados do diretório, que o modelo de ameaças trata como hostil.
- **ADR-0004:** "nó de campo como único committer" exige que um binário de servidor tenha segredos de grupo ou autoridade sobre membros. "Central prevalece" pode devolver à sala um aparelho removido em campo.
- **Lease:** como descrito, é obrigação do cliente e não expiração criptográfica.
- **D-R14:** o padrão "suspender sem destruir" é contrariado por três caminhos automáticos de destruição.

São 5 achados críticos e 16 altos. Os críticos têm correção de desenho conhecida e nenhum exige criptografia à mão. Resolver D01–D03, D13 e D14 em revisão das duas ADRs antes do spike MLS (etapa 3), porque mudam o que o spike precisa provar.

## Achados

| ID | Severidade | Onde | Problema | Proposta |
|---|---|---|---|---|
| D01 | Crítica | ADR-0003 §1–2; doc 01 RNF-A4, doc 03 D02, doc 07 §7.3 | Membros têm uma chave simétrica de custódia por sala e período. Comprometer **um** aparelho-membro (ou o nó de IA) e obter o acervo do servidor dá leitura de **todas** as mensagens do período, inclusive as já expiradas do cache, por até 5 anos. Quebra sigilo futuro e recuperação pós-comprometimento | Membros não decifram custódia: cada cópia cifrada no remetente com HPKE para a **chave pública** de custódia do período, efêmera descartada; privada só em HSM. Registrar que comprometer o HSM expõe o período |
| D02 | Crítica | ADR-0003 §1; doc 01 RN12 | Troca da chave na remoção de membro é só exemplo; removido decifra até o fim do período; membro malicioso pode embrulhar chave errada sem ninguém notar | Período cortado a cada mudança de membros; chave pública do período no estado autenticado do grupo, assinada em cerimônia |
| D03 | Crítica | ADR-0003 "Quórum remoto"; doc 07 R81, R82 | A política do HSM online usa fatos do **diretório** (hostil). Diretório comprometido ou cadastro falso apresenta chave do atacante como "aparelho novo"; não está dito quem registra no HSM as chaves de aprovação e raízes de atestação; "quórum nativo" do HSM costuma ser de cartões, não de assinaturas de celular | (a) Chaves de aprovação e raízes de atestação no HSM só em cerimônia; (b) HSM verifica a atestação e um vínculo assinado fora do `server/`; (c) nonce e contador no HSM; (d) escopo só nos intervalos de pertença; (e) limite de volume. Se o HSM não tiver política programável assinada: v1 sem histórico (opção C) ou quórum presencial |
| D04 | Alta | ADR-0003 §2, §7 | A mesma chave é embrulhada para o HSM offline e para o online: vale o caminho mais fraco | Destruir o embrulho de recuperação no fim da retenção da sala; depois só resta o judicial |
| D05 | Alta | ADR-0003 "Quórum remoto"; doc 07 R83 | Aprovação remota: sem tela confiável; falha comum dos três aparelhos (mesmo MDM, R54); Keystore liga chave a biometria **ou** credencial, não às duas; rosto classe 3 raro; PIN de coação que recusa é percebido pelo coator; tela vaza agente ↔ salas | Confirmação protegida por hardware ou aparelho dedicado de função única; duas chaves se a conjunção for exigida; sob coação parecer aprovado e invalidar no HSM; rótulo opaco de sala; confirmação fora de banda antes da terceira assinatura |
| D06 | Alta | ADR-0003 §1, §4; doc 01 RF67 | Cópia de custódia não é verificável nem autenticada: remetente grava lixo ou texto diferente; com chave comum, qualquer membro forja | Assinatura MLS do remetente na cópia; hash do cifrado de custódia dentro da mensagem MLS; servidor recusa envelope sem cópia; auditoria amostral |
| D07 | Alta | ADR-0003 §7; doc 02 §12, §10 | "Destrói-se a chave embrulhada" é apagar blob no banco (WAL, réplicas, backups) | KEK por período dentro do HSM, destruída com registro; inventário de backups do HSM; definir backup do acervo |
| D08 | Alta | ADR-0003 §8; doc 08 | Acervo de 5 anos sob HPKE clássico: alvo de "colher agora, decifrar depois"; sem segunda camada de Estado | KEM híbrido e segunda camada de Estado também na custódia; sem grau classificado antes disso |
| D09 | Alta | doc 05 §3.4; doc 02 §7.5; doc 01 RF56 | Matrícula ↔ codinome ↔ aparelho ↔ sala em claro no diretório: administrador de infraestrutura quebra o pseudonimato com um SELECT | Matrícula cifrada para chave de quórum ou fora do `server/`; diretório só com identificador opaco e codinome |
| D10 | Média | ADR-0003 §3; doc 07 §7.4, §7.1, §6.3 | Posição em perda e imagem de coação sob a chave judicial (cerimônia): inviável para uso urgente | Separar dado de resposta imediata (chave operacional de dupla custódia, prazo curto) de dado de apuração |
| D11 | Média | ADR-0003; doc 02 §7.1, §7.5, §6.3; ADR-0001; doc 01 RN02 | Nenhum serviço é dono da custódia, recibos, evidências, registro de acessos, pedidos de quórum; RN02 e doc 02 §6.3 ficam falsos | Novo serviço `custody` com `api/` e migrações; atualizar grafo, `check-boundaries.sh`, ADR-0001, RN02 e §6.3 |
| D12 | Média | ADR-0003 §4; doc 02 §1, §6.4 | Estação isolada é um segundo lugar com texto em claro | Atualizar princípio 2; especificar a estação (boot verificado, sem disco persistente, duas pessoas) e incluí-la no modelo de ameaças |
| D13 | Crítica | ADR-0004 §3; doc 10 §3.1, §3.3; doc 02 §7.1; regras 2 e 3 | Para fazer commit, o nó de campo teria de ser membro (lê tudo) ou ter autoridade aceita pelos clientes (nó capturado ligado inclui aparelho do adversário) | Nó só **ordena** e retransmite; commit pelo aparelho do administrador da missão (com substituto); nó sem chave que altere membros |
| D14 | Crítica | ADR-0004 §4; doc 10 §3.3, §5, F-R3 | "Central prevalece": aparelho removido em campo volta a ser membro; aparelho capturado fala com a central durante a missão; mensagens do ramo de campo ficam indecifráveis | Central **congela** os grupos de missão e recusa sessão direta dos aparelhos em missão; campo é o único escritor; ao voltar, central importa a cadeia; lista de revogação de campo pelo primeiro enlace |
| D15 | Alta | doc 10 §3.3, §6; doc 07 §5 | Em campo não há commits de atualização nem remoção em banda estreita | Permitir atualização e remoção serializadas pelo nó; orçamento de banda para remoção; medir no F-R1 |
| D16 | Alta | doc 07 §5, RF65; doc 02 §5.2, §5.3, §7.1, §7.4 | Bifurcação por servidor hostil sem detecção definida; corridas de commit; diretório não acompanha as mudanças do MLS (furo de autorização) | Um commit por época; resumo autenticado do estado em toda mensagem; regra de vencedor; membros verificáveis pelo servidor ou declaração assinada; ADR |
| D17 | Alta | doc 07 §7.1 níveis 3 e 4; §7.4 | Aparelho que zerou chaves não se remove; ninguém designado para remover; nível 3 sem remoção mantém segredos de época | Remoção por terceiros com prazo e alerta; nível 3 remove e é reversível por reinclusão; aviso assinado de autoexclusão antes de zerar |
| D18 | Média | doc 07 §5 | "Segunda conexão derruba ambas e abre SEV1": troca de rede e rádio instável geram falsos | Substituição de sessão com contador assinado; SEV1 só com uso simultâneo comprovado |
| D19 | Média | doc 07 §5; doc 02 §5.2, §9.1; doc 04 §6.4 | Nó de IA sem atestação Android seria recusado; nó de IA é ponto único; chave de administrador aprendida pelo diretório | Credencial de serviço em cerimônia com TPM/HSM; salas "compartimento" sem IA por padrão; chaves de administrador no estado autenticado do grupo |
| D20 | Média | doc 02 §4.2, §7.1; doc 04 §6.4; doc 07 §9 | Diretório valida atestação e emite certificado: comprometido, emite credencial fantasma | Cliente do administrador verifica a atestação no core; código conferido no cadastro; transparência de chaves antes do piloto |
| D21 | Alta | doc 02 §4.3; doc 01 RN10; doc 07 §7.1 | `segredo_servidor` é sempre o mesmo valor: a "validade" é o app esquecer e confiar no relógio; quem o capturou uma vez abre o cache para sempre | Lease criptográfico: chave do banco reembrulhada a cada renovação com o segredo da vez; servidor guarda só o vigente; tempo por contador do servidor; nova ADR sobre a ADR-05 |
| D22 | Alta | ADR-0004 §2, §5; doc 10 §3.1, §4, R92 | Nó de campo guarda os segredos de lease de todos; revogação não chega aos aparelhos; teto pelo relógio | Segredos de missão distintos; credencial do nó com validade curta verificada pelos aparelhos; teto por contador; destruição do nó apaga primeiro os segredos |
| D23 | Alta | doc 02 §4.3, §6.4; doc 07 §7, §7.1, §7.4; doc 04 D-R14 | Três destruições automáticas durante apreensão legal: cache apagado ao vencer o lease; nível 4 por erros de PIN; pânico com apagamento; motivo declarado por um só administrador | Lease vencido não apaga, só deixa ilegível; jurídico decide por escrito os outros dois; segunda autorização quando houver indício de apreensão; corrigir doc 02 §4.3 |
| D24 | Média | doc 07 §7.4, §7.1, R71; doc 02 §6.2, §8.2 | Recibo assinado antes de zerar não prova execução; apagamento pelo MDM fora da ordem assinada; "apagar tudo" e "rastrear" incompatíveis | Recibo de "intenção"; ordem verificada pelo app; MDM com duas pessoas; apagar em etapas, mantendo chave de rastreio curta |
| D25 | Média | doc 07 §7.1, §7.2; regra 11 | Nível 3 automático revoga; §7.2 diz que revogar é humano | Automático só suspende; revogar é ato humano registrado |
| D26 | Média | doc 02 §4.3; doc 10 §4 | Android mata processos: segredo só em memória perde o cache em campo; tendência a persistir | Decidir: serviço em primeiro plano ou segredo embrulhado em hardware com prazo; medir em D-R1 |
| D27 | Alta | ADR-0003 consequências; doc 01 RN03, RN04, RF52; doc 02 §6.2–6.3; doc 07 §7.3, §7.4 | Com custódia, expurgo destrói só a cópia de trabalho: RN04, painel e recibos afirmam destruição que não ocorreu; exceção "compartimento" não verificada | Dois atos: **encerramento** (cópia de trabalho e embrulho de recuperação) e **eliminação** (chave judicial, com autorização arquivística); flag "sem custódia" no estado do grupo; corrigir RN03, RN04, interface |
| D28 | Alta | doc 08 §2.3, §3, RF76, R73; ADR-0003 §3, §7; doc 05 §3.3; doc 06 §8.2; doc 02 §7.5 | Prazos conflitantes do registro de acessos e de visualização; mesma chave morre em 5 anos e deve viver 15; recibo do cliente não registra negações; `audit_log.actor` em claro | Hierarquia de chaves separada por grau; registro de acesso gerado no servidor e cifrado para quórum; identificador de sala retido; resolver com D-R4/D-R5 antes da Fase 4 |
| D29 | Média | doc 02 §6.2; doc 07 §7.4; ADR-0004 §4 | `seq` por sala sem dono; janela de tempo inviável offline; revinculado perde o `seq`; ordem sem alvo de aparelho | Único escritor por sala; aparelho usa só assinatura e contador; contador entregue no reingresso; formato próprio da ordem de apagamento |
| D30 | Alta | regra 2; doc 02 §4.1, §4.3, §8.1; `core/CLAUDE.md`; ADR-0003 | Regra 2 literal não vale para StrongBox, aprovação do comandante, HSM, CA, nó de IA; Keystore passa por Java | Reescrever a regra 2 por escopo; spike prova desembrulho via JNI a partir do core; decidir se `ai-node/client` reutiliza o core |
| D31 | Média | regra 7; doc 02 §2, §9.1, §9.3 | "ai-node não abre conexão de saída" é impossível: conecta ao gateway, anexos e banco | Regra 7 com lista fixa de destinos; teste em `tests/security` |
| D32 | Média | doc 07 R52; doc 03 §6 | R52 ainda diz "cifra para o Auditor"; a cópia das regras no doc 03 §6 está desatualizada | Corrigir R52; doc 03 §6 aponta para o `CLAUDE.md` em vez de copiar |
| D33 | Média | doc 08 §9; `PIPELINE.md`; doc 03 Fase 0 | Doc 08 manda decidir D-R17/D-R18 antes dos spikes; pipeline põe em G3; spike MLS não cobre custódia, commits, remoção offline nem Keystore | Escolher a ordem; ampliar o critério da etapa 3 (HPKE de custódia, um commit por época, remoção offline, Keystore) |
| D34 | Baixa | doc 02 §11; `docs/adr/README.md`; doc 08 §1 | Numerações colidem (ADR-03 × ADR-0003, ADR-04 × ADR-0004); ADR-0002 sem arquivo; doc 08 §1 desatualizado | Tabela de correspondência; criar 0002 antes do spike de atestação; atualizar doc 08 §1 |

## Contradições entre documentos

1. Chave de custódia: doc 07 §7.3 (chave pública, privada em HSM) × ADR-0003 §1–2 (chave que os membros têm) (D01).
2. Sigilo futuro: doc 01 RNF-A4 e doc 03 D02 × ADR-0003, que não reconhece o conflito.
3. ADR-0003 rejeita "três celulares abririam tudo" e adota aprovação por três celulares (D03).
4. Nó de campo faz commits (ADR-0004 §3, doc 10 §3.3) × "diretório não participa dos grupos", princípio 1 e regra 3 (D13).
5. Commit a cada 24 h (doc 07 §5) × "aparelhos só enviam mensagens" em campo (D15).
6. Lease vencido apaga o cache (doc 02 §4.3) × nível 3 "nada é destruído" e D-R14 (D23).
7. Revogação automática (doc 07 §7.1) × decisão humana (§7.2) (D25).
8. Destinatário da posição em perda: "Auditor e gestor" (doc 07 §7.4) × quórum judicial (ADR-0003 §3); R52 ainda "Auditor" (D10, D32).
9. Retenção dos registros de visualização e de acessos (doc 06 §8.2, doc 05 §3.3, doc 08 §3, ADR-0003 §7) (D28).
10. O que resta após o expurgo: RN02, RN04, doc 02 §6.3 × custódia por anos (D11, D27).
11. Quebra de pseudonimato "só por quórum" × matrícula em claro no diretório (D09).
12. Lugares com texto em claro: só o nó de IA (doc 02 §1) × estação isolada (ADR-0003 §4) (D12).
13. Regra 7 × doc 02 §2, §9.1 (D31); regra 2 × doc 02 §4.1 e ADR-0003 (D30).
14. Apagar aparelho: perfil × aparelho inteiro; "apagar tudo" × "rastrear" (D24).
15. Ordem da Fase 0: doc 08 §9 × `PIPELINE.md` (D33); numeração de ADRs (D34).

## Perguntas que só a instituição responde

1. Com custódia, comprometer o HSM ou o quórum expõe todo o período por até 5 anos. Aceita? Se não: MLS puro sem reenvio (opção C).
2. Revinculação com histórico é requisito da v1? Se puder esperar, HSM online e aprovação remota saem da v1.
3. O histórico reenviado cobre só a retenção operacional da sala ou todo o prazo de guarda?
4. Quem são os N detentores (D-R24)? Haverá aparelho dedicado para aprovar? Comandante pode ser membro de sala ou gestor?
5. D-R14: apagamento por erros de PIN e pânico ficam ativos quando pode haver apreensão legal? A custódia basta para afastar "destruição de prova"?
6. Agentes e administradores serão informados de que "apagar a operação" não elimina a cópia institucional? Quem autoriza a eliminação definitiva (D-R5)?
7. Haverá salas "compartimento" sem custódia? Como fica o art. 38 §4º e a Lei 8.159?
8. Prazo do registro de acessos por grau, reclassificação, e norma federal ou estadual.
9. D-R17 e D-R18 antes ou depois dos spikes?
10. Tirar a matrícula do banco do diretório (D09), com custo operacional para o gestor?
11. Durante a missão, a central pode ficar sem alterar nem entregar para os grupos de missão (D14)? Quem é o administrador da missão e o substituto?
12. Nó de campo capturado ligado: procedimento e prazo máximo até os aparelhos deixarem de confiar nele.
13. Dados de resposta imediata (posição em perda, coação): quem precisa ver em minutos, e com quantas pessoas (D10)?
14. HSM: modelo, política programável assinada, quórum de cartões, limite de chaves, auditoria externa.
15. Nó de IA em sala secreta: decisão por sala na criação, visível aos membros?
