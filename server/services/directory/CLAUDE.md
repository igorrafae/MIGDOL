# services/directory

## Responsabilidade
Cadastro, atestação, pacotes de chave MLS, membros por sala, leases, administradores da sala.

## Não faz
Não participa dos grupos MLS e não vê conteúdo.

## Dono dos dados
users, devices, rooms, memberships, key_packages, terms, term_acceptances.

## Depende de (somente via api/)
audit/api

## Contrato exposto
api.Directory (IsMember, DeviceActive)

## Regras do módulo
- Código em `internal/` é privado; outros módulos só importam `api/`.
- Nunca ler nem escrever tabela de outro serviço. Nunca criar FK para tabela de outro serviço.
- Nunca importar tipo de implementação de outro serviço, só o contrato.
- Regras globais: CLAUDE.md da raiz (autorização por sala, assinatura antes de gravar estado, logs sem conteúdo).

## Plano de extração
É consultado em toda publicação: ao separar, exige cache local com TTL curto no cliente remoto e política de negar por padrão se o diretório estiver indisponível. Banco próprio.

Fase: 1
