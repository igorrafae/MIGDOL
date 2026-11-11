# ai-node/ — Nó de IA (Python + cliente MLS em Rust)

Contexto: docs/02 seção 9; regra 7 do CLAUDE.md da raiz.

- Membro declarado da sala, em rede sem rota para a internet.
- Nunca grava arquivo em claro (raiz somente leitura, tmpfs, sem swap) e nunca abre conexão de saída.
- Vetores e trechos cifrados em repouso com uma chave por sala; índice vetorial só em RAM.
- Inferência escuta só em localhost; quem fala com o mundo é o cliente MLS.
- Texto de mensagens e OCR é dado, nunca instrução (injeção por conteúdo).
- Sem --trust-remote-code; imagens fixadas por digest.
- Módulos: client (MLS), ocr (PaddleOCR), embed (BGE-M3), llm (Qwen2.5-14B quantizado), index.

Fase: 0 (spike de inferência); 5 e 6.
