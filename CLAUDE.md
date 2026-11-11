# Regras do projeto

## Leitura obrigatória
- docs/01-negocio-requisitos.md
- docs/02-arquitetura-engenharia.md
- docs/03-seguranca-homologacao-roadmap.md

## Regras que não podem ser quebradas
1. Nenhuma primitiva ou protocolo criptográfico é implementado à mão.
   Use as bibliotecas definidas em docs/02. Se faltar algo, pare e pergunte.
2. Chaves privadas e chaves de sessão existem apenas em core/ (Rust).
   Nunca atravessam a ponte para Kotlin ou TypeScript.
3. O servidor (server/) nunca recebe, registra ou grava conteúdo em claro.
   Nenhum log contém corpo de mensagem ou room_id em claro.
4. app/src não usa localStorage, sessionStorage, IndexedDB nem fetch direto.
   Todo acesso a dados e rede passa pelo plugin que fala com o core.
5. Todo endpoint e todo tratador de mensagem verifica autorização por sala.
6. Verificação de assinatura vem antes de qualquer gravação de estado.
7. ai-node/ não grava arquivo em claro e não abre conexão de saída.
8. Dependência nova exige justificativa no PR e versão fixada com hash.

## Como trabalhar
- Uma fase do roteiro por vez (docs/03, seção 5). Não adiantar fases.
- Cada entrega vem com teste automatizado ligado ao critério de aceite.
- Cenários de docs/03 seção 3 viram testes em tests/security quando automatizáveis.
- Mudança que contrarie um ADR de docs/02 seção 11 exige novo ADR antes do código.
- Em dúvida sobre segurança, escolher a opção mais restritiva e registrar a dúvida.
