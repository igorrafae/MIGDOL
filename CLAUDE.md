# Regras do projeto

## Leitura obrigatória
- docs/01-negocio-requisitos.md
- docs/02-arquitetura-engenharia.md
- docs/03-seguranca-homologacao-roadmap.md
- docs/04-gestao-de-riscos.md
- docs/05-vazamentos-incidentes-e-atribuicao.md
- docs/06-lgpd-e-conformidade.md
- docs/07-resistencia-a-clonagem-e-endurecimento-avancado.md
- docs/08-informacao-classificada-e-nivel-inteligencia.md
- docs/09-perfil-exercito-e-modelo-de-implantacao.md
- docs/10-operacao-em-campo-por-radio.md

Cada componente (app/, core/, server/, ai-node/, admin-panel/, deploy/, tests/) tem um CLAUDE.md local com fronteiras e fase. Leia o do componente antes de mexer nele.

## Regras que não podem ser quebradas
1. Nenhuma primitiva ou protocolo criptográfico é implementado à mão.
   Use as bibliotecas definidas em docs/02. Se faltar algo, pare e pergunte.
2. Chaves privadas e de sessão de membro (MLS, assinatura, cache) existem apenas em core/ (Rust)
   e nunca atravessam a ponte para Kotlin, Swift ou TypeScript. Chaves de hardware (StrongBox/TEE,
   Secure Enclave) não saem do hardware: Kotlin e Swift só seguram a referência, e o uso é chamado
   a partir do core; o que o hardware devolve (ex.: desembrulho pelo Keystore) transita só no quadro
   da chamada iniciada pelo core, é zerado em seguida e sozinho não abre nada. Chaves da instituição
   (AC, quórum, cerimônia) ficam em HSM, fora do app.
3. O servidor (server/) nunca recebe, registra ou grava conteúdo em claro.
   Nenhum log contém corpo de mensagem ou room_id em claro.
4. A interface (Compose, SwiftUI e, enquanto existir, app/src) não grava dado fora do core (sem
   localStorage, sessionStorage, IndexedDB, SharedPreferences, DataStore, arquivos próprios,
   UserDefaults nem Core Data) e não abre rede própria (sem fetch, OkHttp nem URLSession).
   Todo acesso a dados e rede passa pelo core.
5. Todo endpoint e todo tratador de mensagem verifica autorização por sala.
6. Verificação de assinatura vem antes de qualquer gravação de estado.
7. ai-node/ não grava arquivo em claro e não abre conexão de saída.
8. Dependência nova exige justificativa no PR e versão fixada com hash.
9. server/ é monólito modular: serviços só se falam por api/, sem importar internal/
   de outro serviço, sem consultar tabela alheia e sem FK entre serviços.
   Rode scripts/check-boundaries.sh antes de commitar mudanças no server/.
10. Nenhum dado de produção em desenvolvimento ou teste. Nenhuma telemetria, serviço de
    erro, push ou IA de terceiros. Dado pessoal de usuário só por rito documentado (docs/06).
11. Nenhuma decisão contra pessoa é tomada automaticamente a partir de IA, marca d'água
    ou integridade (docs/05). Mudança que altere o perfil de risco atualiza docs/04.
12. Câmera: só com o app em primeiro plano, na tela de conteúdo e no desbloqueio. Sem alerta, nenhum
    quadro é gravado nem enviado; com alerta, só imagens de evidência cifradas para a chave de quórum
    (docs/07). O gabarito facial fica só no aparelho, cifrado por chave de hardware, e nunca é enviado.
    Atestação sem Play Integrity nem serviço de terceiro.
13. A instituição assume a privacidade dos usuários (docs/06 seção 0). Todo dado coletado sobre
    agentes é cifrado, com acesso por rito e prazo, por segurança operacional.
14. Membros leem as salas pelo MLS e nunca decifram cópia de custódia. Cópia de custódia só existe
    em sala marcada sob custódia na criação (flag assinada no instituicao.yaml; docs/adr/0003):
    HPKE para chave pública, privada só no HSM offline. Sala comum é MLS puro e é expurgada ao concluir.
    Todo acesso institucional a dado cifrado (custódia, evidência, registros, termo, localização,
    pseudonimato, segredo de lease guardado) exige quórum presencial de no mínimo 3 comandantes,
    com cartões, em HSM offline, e base registrada (ordem judicial, ou CI em processo administrativo
    com validação jurídica).
    Nenhuma pessoa decifra sozinha. Divisão de segredo nunca é feita à mão. Nenhuma chave de quórum
    fica em HSM online, e não há aprovação remota por quórum. A revinculação não decifra dado da
    instituição: aprovação dupla (chefe da inteligência + superior da área) e histórico enviado por
    outro membro pelo MLS.

## Como trabalhar
- Uma fase do roteiro por vez (docs/03, seção 5). Não adiantar fases.
- Cada entrega vem com teste automatizado ligado ao critério de aceite.
- Cenários de docs/03 seção 3 viram testes em tests/security quando automatizáveis.
- Mudança que contrarie um ADR de docs/02 seção 11 ou de docs/adr exige novo ADR antes do código.
- Em dúvida sobre segurança, escolher a opção mais restritiva e registrar a dúvida.
- Modelos (docs/03 seção 6.1): Fable 5.2 revisa desenho, ADRs e ameaças; Opus 5.5 implementa e revisa
  o que é de segurança (core/, cripto, autorização, chaves, custódia, expurgo, atestação); Sonnet 5.5
  faz o volume (UI, apoio, testes comuns, docs, scripts). Na dúvida, Opus. Nunca enviar dado real.

## Arquitetura do backend
O server/ é um monólito modular: um módulo Go por serviço em server/services/, falando só por api/.
Leia server/CLAUDE.md, docs/adr/0001-monolito-modular.md e docs/MIGRACAO-MICROSSERVICOS.md.
Rode scripts/check-boundaries.sh antes de commitar mudanças no server/.
