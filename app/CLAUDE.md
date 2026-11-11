# app/ — React 18 + TypeScript + Capacitor (Android)

Contexto: docs/02 seção 8; regra 4 do CLAUDE.md da raiz.

- app/src não usa localStorage, sessionStorage, IndexedDB nem fetch direto. Dados e rede passam pelos plugins que falam com o core.
- A WebView só recebe texto em claro para exibir; sem depuração remota em produção, sem acesso a arquivos.
- Marca d'água desenhada em camada nativa, nunca no DOM.
- android/plugins/security: integridade do aparelho (contrato DeviceSecurity, docs/02 8.2).
- android/plugins/camera: câmera interna, sem passar pela galeria.
- android/plugins/viewer: visualizador de anexos com FLAG_SECURE.

Fase: 1 (conversa); 2 (plugin de segurança); 3 (câmera e visualizador).
