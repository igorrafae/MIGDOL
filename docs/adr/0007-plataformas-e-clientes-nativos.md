# ADR-0007 — Plataformas, clientes nativos e núcleo único

**Status:** **aceita** (GD-19, 12/11/2026; integrada na segunda rodada da etapa 1b, `docs/fase0/01d-correcoes-rodada2.md`, e aceita na terceira, `docs/fase0/01f-correcoes-rodada3.md`) · **Substitui** doc 02 seção 3 (casca Capacitor com UI React/TypeScript) e doc 01 decisão D1, a partir da Fase 1 · **Regra 4 do `CLAUDE.md` alterada** com o texto do item 5 (terceira rodada). A regra 2 já cita Swift desde a GD-12 · Diretórios **não** são renomeados antes da Fase 1 · **Mantém** ADR-06 (Android gerenciado na v1) e ADR-01 (MLS em Rust) · **Complementa** ADR-0008 e ADR-0009

## Contexto
O responsável pensou o produto em três projetos: Android, Apple e Linux (painel web de gestão e servidor). O macOS foi **descartado**: desktop não tem o controle de MDM, câmera e coação de um celular gerenciado (R65). O iPhone traz dependência da Apple: o App Attest consulta servidores da Apple (regra 12), a entrega em segundo plano depende do APNs (regra 10, D-R1) e o MDM da Apple também usa o APNs.

## Decisão (proposta)
1. **Plataformas:**
   - **Android:** v1, como hoje (ADR-06).
   - **iPhone (iOS):** fase posterior, só depois de um **estudo de viabilidade** na Fase 0 com aparelho dedicado e de uma ADR própria sobre APNs, App Attest e MDM da Apple (D-R31). Até essa ADR, o iOS não entra em produção. O estudo não escreve código de produto (uma fase por vez). Etapa própria no `PIPELINE.md` (5b, depois de G1), com escopo: APNs, App Attest, MDM da Apple, bloqueio de captura de tela (doc 01 D2), série atestada (doc 07 seção 3) e conexão em segundo plano (D-R1); usa o protótipo do doc 11 (01e-E19a).
   - **macOS:** fora do escopo.
   - **Linux:** servidor (`server/`), nó de IA (`ai-node/`) e painel web de gestão (`admin-panel/`), só na rede interna (ADR-0008).
2. **Núcleo único em Rust** (`core/`): MLS, chaves, cache, custódia, lease e expurgo existem **uma vez só**, gerados para Kotlin e Swift pelo UniFFI. Nenhum cliente reimplementa criptografia (regras 1 e 2).
3. **Telas nativas:** Jetpack Compose no Android e SwiftUI no iPhone, no lugar do Capacitor/WebView. Sem WebView, some o armazenamento web, mas **não** a disciplina da regra 4: Kotlin e Swift têm armazenamento e rede próprios (SharedPreferences, DataStore, arquivos, UserDefaults, Core Data, OkHttp, URLSession). A regra 4 precisa passar a cobrir a interface nativa (item 5); não vale "por construção".
4. **Um repositório** (monorepo): `core/`, `server/`, `android/`, `apple/`, `admin-panel/`, `ai-node/`, `deploy/`, `tests/`. Uma versão, um build reproduzível, um pacote assinado (ADR-15).
5. **Redação das regras:**
   - **Regra 2:** já atualizada pela GD-12 (Kotlin, Swift ou TypeScript; chaves de hardware no Keystore e no Secure Enclave).
   - **Regra 4:** **aplicada** ao `CLAUDE.md` na terceira rodada (GD-19), com este texto: "A interface (Compose, SwiftUI e, enquanto existir, app/src) não grava dado fora do core (sem localStorage, sessionStorage, IndexedDB, SharedPreferences, DataStore, arquivos próprios, UserDefaults nem Core Data) e não abre rede própria (sem fetch, OkHttp nem URLSession). Todo acesso a dados e rede passa pelo core." Teste estático em `tests/security` que procura essas APIs fora do core. O Fable recomendou (01e, seção 5.1) a regra em termos gerais e a lista de APIs só no teste, porque a lista envelhece; fica como sugestão ao responsável, sem mudar o texto aplicado.

## Alternativas rejeitadas
- **Três repositórios independentes:** difícil fixar a versão do core, três homologações, risco de cripto duplicada.
- **Capacitor para Android e iOS:** WebView aumenta a superfície e exige a regra 4 como disciplina.
- **macOS como cliente restrito:** descartado pelo responsável; sem MDM e coação equivalentes.
- **iOS na v1:** dependência da Apple ainda não avaliada.

## Consequências
- A interface é feita duas vezes (Compose e SwiftUI) quando o iPhone entrar.
- `app/` (Capacitor) é renomeado para `android/` **no início da Fase 1** (aceita pela GD-19; renomear é trabalho da Fase 1, não da Fase 0); o `CLAUDE.md` do componente muda junto. Até lá, `app/` existe como está, sob a regra 4 nova.
- O spike MLS (etapa 3) não usa React: núcleo Rust com cliente de linha de comando e ponte Kotlin (doc 03, Fase 0; 01e-E19b).
- Os plugins nativos de hoje (segurança, câmera, visualizador) passam a ser código Kotlin da interface; o que tocar conteúdo, chave ou quadro de câmera continua indo direto ao core (regras 2 e 12).
- Doc 02 seções 1, 2, 3, 7.2 e 8 ainda descrevem a pilha anterior, com nota desta ADR aceita; o texto muda com o código, na Fase 1.
- Protótipo de interface dos três projetos e protocolo de usabilidade (V09): doc 11 (GD-19).
- Equipamento de teste: um Android (portão G1) e um iPhone (estudo de viabilidade).
- Riscos: R101 (dependência da Apple) e R102 (divergência entre clientes).

## Revisar quando
O estudo de viabilidade do iOS terminar, ou se a equipe não tiver capacidade para duas interfaces nativas.
