# core/ — Núcleo criptográfico (Rust)

Contexto: docs/02 seções 4, 5, 6 e ADR-01; regras 1 e 2 do CLAUDE.md da raiz.

- Único lugar com chaves privadas e de sessão. Nada atravessa a ponte para Kotlin/TypeScript.
- MLS por biblioteca auditada (OpenMLS candidata; confirmar no spike da Fase 0). Nunca implementar primitiva à mão.
- Responsável por: MLS, cifra de anexos (AES-256-GCM em blocos), cache SQLCipher, rede (WebSocket + mTLS).
- Chaves em memória travada e zeradas explicitamente.
- Exposição ao Kotlin por UniFFI/JNI.
- Dependências com versão fixada e hash; nenhuma versão está decidida ainda (Fase 0).

Fase: 0 (spike MLS); 1 em diante.
