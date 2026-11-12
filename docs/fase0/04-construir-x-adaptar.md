# Construir × adaptar — Fase 0, etapa 4

**Modelo:** Opus 5.5 (pipeline da Fase 0, etapa 4). Fatos externos conferidos na web em 11/11/2026 (fontes no fim); o que não foi conferido está marcado **[verificar]**. Não é parecer jurídico sobre licenças.
**Conferência do orquestrador:** a recomendação coincide com o doc 01 §3.4 (construir com biblioteca MLS auditada, a confirmar no spike).

## Pergunta

Construir o MIGDOL na pilha planejada (core Rust com OpenMLS, monólito modular Go, UI Capacitor/React, plugins Kotlin) ou adaptar um mensageiro seguro de código aberto? Pesam as regras do `CLAUDE.md` e a revisão de desenho: custódia por HPKE para chave pública em HSM (D01–D03), nó de campo que só ordena (D13–D14), um commit por época com detecção de bifurcação (D16) e lease criptográfico (D21) mexem no protocolo, e nenhum produto pronto oferece isso.

## Comparação

| Critério | **A. Construir como planejado** | **B. Matrix** (Synapse/Dendrite ou Conduit + Element X) | **C. Wire** (wire-server + Android + core-crypto) | **D. Signal** (libsignal + app + servidor) | **E. Só bibliotecas** (OpenMLS ou mls-rs) |
|---|---|---|---|---|---|
| Licença on-prem | Código próprio; dependências MIT/Apache | Synapse e Dendrite em **AGPLv3** desde dez/2023, com licença comercial e CLA da Element; Conduit Apache-2.0 **[verificar]**; Element X AGPL **[verificar]** | Servidor **AGPL-3.0**; core-crypto **GPL-3.0** (contamina o app) | **AGPLv3**; uso fora do Signal "sem suporte" | OpenMLS **MIT**; mls-rs **Apache-2.0/MIT** |
| Protocolo de grupo | MLS (RFC 9420) via OpenMLS | Megolm em produção; MLS descentralizado ainda prova de conceito | MLS em produção **[verificar]**; core-crypto sobre OpenMLS | Double Ratchet e grupos v2; **sem MLS** (contraria D4, ADR-01) | MLS; OpenMLS auditada pela SRLabs (achados corrigidos em 0.7.3/0.8.1) |
| Servidor sem texto em claro (regra 3) | Por desenho | **Estado da sala em claro** e replicado na federação | Conteúdo cifrado; diretório e conversas visíveis | Conteúdo cifrado, metadados mínimos | Por desenho |
| Chaves só no core (regra 2) | Por desenho (com D30) | matrix-rust-sdk; casca do app da Element **[verificar]** | core-crypto em Rust | libsignal; app Java/Kotlin manipula estado | Por desenho |
| Sem push/telemetria/Google (regra 10) | WebSocket próprio | FCM padrão; UnifiedPush alternativa **[verificar]** | FCM padrão **[verificar]** | FCM com reserva WebSocket; servidor ligado a nuvem **[verificar]** | Por desenho |
| Atestação sem Play Integrity | Spike próprio | Não existe | Não existe | Não existe | Spike próprio |
| Expurgo por destruição de chave | Por desenho | Contra o modelo (histórico replicado, backup de chaves) | Parcial | Não há ordem institucional | Por desenho |
| Custódia com quórum (ADR-0003 corrigida) | Encaixa no envelope | Fork de cliente e servidor | Fork do core-crypto (GPL) e backend Haskell | Contra o modelo | Encaixa |
| Instância dedicada (ADR-13) | Por desenho | Possível sem federação | Possível; Kubernetes recomendado (conflita com ADR-15) | Pouco documentada **[verificar]** | Por desenho |
| Nó de campo que só ordena | Perfil de campo | Homeserver de campo federado bifurca estado | Sem modelo | Não existe | Encaixa |
| Metadados expostos | Sala opaca, horário, tamanho | **Altos** (grafo social, estado da sala) | Médios | Baixos | Iguais a A |
| Superfície de ataque | Pequena | Muito grande | Grande (Haskell, vários serviços) | Grande, mudanças sem aviso | Mínima |
| Esforço para tirar o que sobra | Nenhum | Muito alto; fork diverge | Alto | Alto; protocolo não serve | Nenhum |
| Fornecedor estrangeiro | Só bibliotecas | Element (Reino Unido) | Wire (Alemanha/Suíça) | Signal (EUA) | Mantenedores OpenMLS ou AWS (mls-rs) |
| Manutenção | Equipe própria sustenta | Seguir upstream num fork pesado | Seguir upstream em Haskell | API instável | Atualizar uma biblioteca |

## Recomendação

**Construir como planejado (A), reaproveitando só bibliotecas (E). Não adaptar Matrix, Wire nem Signal.**

1. **O diferencial está no protocolo, não na tela:** custódia HPKE com cópia assinada, expurgo por chave com recibo, ordem binária com contador, lease criptográfico, nó de campo sem commit, um commit por época e atestação sem Google não existem em nenhum dos três; exigiriam fork de cliente e servidor.
2. **Signal sai:** sem MLS, AGPL, sem suporte fora do Signal.
3. **Matrix contraria o modelo de ameaças:** estado da sala em claro, federação e histórico replicado contra o expurgo, MLS ainda em pesquisa; metadados são o alvo de adversário estatal (doc 08 §4.2).
4. **Wire é o mais próximo, mas não compensa:** GPL-3.0 do core-crypto, backend em Haskell, e o core-crypto é camada sobre o OpenMLS: usar o OpenMLS direto dá o mesmo protocolo com licença MIT e menos código.
5. **Homologação:** auditar o que se escreve sobre biblioteca auditada custa muito menos que auditar um fork grande e móvel.
6. **Licença:** MIT/Apache sem obrigação de entregar código aos agentes nem CLA de fornecedor.

## O que reaproveitar como biblioteca

Cada dependência com justificativa no PR, versão fixada com hash (regra 8), revisão de dependências críticas e SBOM.

| Biblioteca | Uso | Nota |
|---|---|---|
| **OpenMLS** (MIT) | Protocolo de grupo no core e no cliente do nó de IA | Principal; versão com as correções da auditoria SRLabs (≥ 0.7.3 / 0.8.1) **[verificar a vigente]** |
| **mls-rs** (Apache/MIT) | Reserva e teste de interoperabilidade no spike | Auditoria externa **[verificar]**; medir persistência própria (SQLCipher) |
| Backend criptográfico (RustCrypto ou libcrux) | Primitivas | libcrux pré-1.0; preferir o auditado **[verificar]** |
| HPKE (RFC 9180) da mesma pilha | Cópia de custódia (D01) | Nunca à mão (regra 1) |
| X-Wing (ML-KEM-768 + X25519) | Híbrido pós-quântico (D-R12, D08) | **Experimental**; não ser a única proteção em grau classificado |
| UniFFI, SQLCipher, `pgx/v5` | Ponte, cache, banco | Já previstos (doc 02 §3) |
| Código Matrix, Wire, Signal | Só **referência de leitura** | Não copiar código AGPL/GPL |

## Riscos da recomendação

| Risco | Mitigação |
|---|---|
| Mais esforço em UI, entrega e anexos | Escopo mínimo por fase; a tela é a parte barata |
| Erro de orquestração do MLS (validação de commit, bifurcação, persistência) | Spike com critério ampliado (D33); revisão 3b; testes de bifurcação; auditoria externa antes do piloto |
| Uma biblioteca MLS, poucos mantenedores | Camada fina de abstração; mls-rs como reserva testada; acompanhar avisos (GHSA) |
| Pós-quântico imaturo | Segunda camada de Estado (D-R18); sem grau classificado antes disso |
| Equipe própria sustenta o protocolo por anos | ADRs; build reproduzível; responsável pela pilha criptográfica |

## Pontos em aberto

1. OpenMLS **ou** mls-rs, decidido no spike (etapa 3): persistência no SQLCipher, HPKE de custódia, um commit por época, remoção offline, Keystore via core (D30, D33).
2. Versão exata e hash, com as correções da auditoria **[verificar]**.
3. Auditoria externa publicada da mls-rs **[verificar]**.
4. X-Wing ou ML-KEM na biblioteca e caminho para a suíte oficial da IETF **[verificar]**.
5. Jurídico: risco de obra derivada ao ler código AGPL/GPL; distribuição interna do app conta como "distribuição"? (Só importa se entrar dependência copyleft.)
6. `ai-node/client` reutiliza o core Rust (D30)?
7. Rever se o Matrix publicar MLS descentralizado estável **com** estado de sala cifrado.

## Fontes

- Element, Synapse e Dendrite para AGPLv3: https://element.io/blog/element-to-adopt-agplv3/ · https://github.com/element-hq/synapse
- Matrix e MLS: https://arewemlsyet.com/ · https://www.matrix.org/blog/2023/07/a-giant-leap-with-mls/ · https://www.uhoreg.ca/documents/fosdem-2026/__/dmls_vs_dmls.html · https://cfp.2026.matrix.org/matrix-conference-2026/talk/7GYBAP/
- Wire: https://github.com/wireapp/wire-server · https://github.com/wireapp/core-crypto · https://wireapp.github.io/core-crypto/main/rust/core_crypto/
- Signal: https://github.com/signalapp/libsignal · https://awesome.ecosyste.ms/projects/github.com%2Fsignalapp%2FSignal-Server
- OpenMLS: https://github.com/openmls/openmls · https://blog.openmls.tech · https://advisories.gitlab.com/cargo/openmls/GHSA-qr9h-x63w-vqfm/
- mls-rs: https://github.com/awslabs/mls-rs · https://docs.rs/crate/mls-rs/0.38.0/source/README.md
- Pós-quântico: https://cryspen.com/post/pq-openmls/ · https://pub.dev/documentation/openmls/latest/ · https://www.ietf.org/archive/id/draft-mahy-mls-pq-00.html
