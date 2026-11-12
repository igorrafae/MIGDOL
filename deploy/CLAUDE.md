# deploy/ — Implantação

Contexto: docs/02 seção 10. Zona de serviço e zona de IA separadas; zona de IA sem internet.
Banco da IA fora de backup. Autoridade certificadora com raiz offline. Imagens fixadas por digest.

## Pacote de instalação (docs/02 seção 10.1)
- Tudo sai como pacote offline assinado: imagens por digest, manifestos YAML por perfil
  (central, zona-ia, campo), instituicao.exemplo.yaml, APK, SBOM, hashes, runbooks, verificador.
- A instalação não baixa nada. O verificador roda antes de qualquer passo e recusa pacote que não confere.
- Nenhum segredo em YAML ou no pacote: chaves no HSM (cerimônia), senhas geradas na instalação.
- Um pacote para todas as instituições; diferenças só por perfil e instituicao.yaml.
- Orquestração: Compose ou systemd (D-R28); Kubernetes só por decisão da instituição.
- Fase: o pacote real começa na Fase 1. Até lá, nada de manifestos de produção aqui.
