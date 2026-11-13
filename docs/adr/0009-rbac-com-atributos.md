# ADR-0009 — Controle de acesso: papéis com atributos (RBAC + ABAC)

**Status:** proposta (decisão do responsável pelo produto, 12/11/2026; integrada aos documentos na segunda rodada da etapa 1b, `docs/fase0/01d-correcoes-rodada2.md`; **terceira rodada**, `docs/fase0/01f-correcoes-rodada3.md`: raiz de confiança fora do diretório para papéis privilegiados, 01e-E10; papéis que faltavam, 01e-E18; concessão de escopo, GD-15; papéis da ADR-0010) · **Decidir antes da Fase 1** (RF89 e RF90 são Essencial, Fase 1; 01e-E19c) · **Complementa** regra 5 (autorização por sala), regra 14, doc 01 seção 4 (atores), doc 04 seção 6.2 (separação de funções), doc 08 seção 5 (credencial por grau e compartimento), ADR-0005 (estado do grupo), ADR-0008

## Contexto
O acesso no MIGDOL depende de **papel** (o que a pessoa faz) e também de **ser membro da sala, do grau, do compartimento e do prazo**. Só RBAC não expressa isso; só ABAC fica difícil de auditar.

## Decisão (proposta)

### 1. Papéis
| Papel | Ator no doc 01 seção 4 | Pode | Nunca pode |
|---|---|---|---|
| Agente | Operador | Usar as salas das quais é membro | Ver indicadores; gerir sala |
| Administrador de operação (gestor de sala) | Administrador de operação | Incluir e remover membros da própria sala; ordens da sala | Ler ou gerir outras salas |
| Comandante e subcomandante de missão | Administrador de operação em missão de campo (GD-11) | Administrar o grupo de missão; ver posição na missão (MLS); assinar conclusão | Abrir custódia sozinho |
| Equipe de inteligência | Equipe de inteligência | Marcar a sala como ativa (RN18) | Concluir sozinha |
| Solicitante e, quando ele é comandante ou subcomandante, o superior imediato dele | Solicitante | Assinar a conclusão (RN18) | Assinar duas vezes a mesma conclusão |
| Chefe de inteligência e superior da área | Chefe da inteligência e superior da área | Aprovar revinculação, sempre os dois | Decifrar dado da instituição |
| Analista de inteligência | (novo; ADR-0008) | Ver indicadores nível 2 (agregados e nominais) só de salas com **concessão vigente** da sala de origem (GD-15), como membro da sala de análise; perguntar à IA (ADR-0008) | Ler conversas das salas do escopo; ver sala sem concessão |
| Gestor (painel) | Gestor do sistema | Ver indicadores nível 1a; cadastrar (com a identificação do agente, RN05) e revogar usuários e aparelhos | Ver, **no painel de indicadores**, nomes, salas, posição ou conteúdo (01e-E18a) |
| Detentor de quórum | Comandantes detentores de parte | Participar da cerimônia presencial no HSM offline | Agir sozinho ou remotamente |
| Administrador de sistema | (doc 04: administrador de infraestrutura) | Infraestrutura, instalação, atualização | Conteúdo, indicador ou chave |
| Corregedoria e auditoria | Auditor | Consultar a trilha de metadados (RF53); **pedir** abertura de registro de acessos ou de visualização pelo rito; **iniciar** investigação sigilosa (ADR-0010) | Abrir qualquer registro cifrado sem quórum presencial (regra 14) |
| Segunda assinatura da sala comum com a flag ligada | (GD-14; quem: D-R36) | Assinar a exceção "comum" na criação da sala | Ser o criador ou membro da sala |
| Autorizadores do apagamento remoto, da revogação em lote e da destruição do segredo selado | (doc 07 seção 7.4; ADR-0006 item 8) | Dar a segunda autorização, sempre duas pessoas distintas | Autorizar sozinho; ser o agente do aparelho |
| Criador da sala de análise | Chefia (D-R32) | Criar a sala de análise com escopo e prazo, com aprovação dupla (ADR-0008) | Incluir sala sem concessão da origem |
| Plantão da sala de resposta | Segurança de plantão e gestor de serviço (pergunta 8 do 01c) | Receber posição em perda e alerta de coação pela sala de resposta (MLS) | Ler outras salas do agente |
| Iniciador e segunda autorização da investigação sigilosa | Detentor de quórum ou corregedoria; segunda autorização independente (ADR-0010; D-R37) | Abrir o pedido selado, com CI, na estação isolada | Estar impedido; ser a mesma pessoa; aprovar remotamente |
| Revisor do relatório de rastro | Duas pessoas não impedidas (ADR-0010 seção 10) | Revisar e assinar o relatório antes do envio | Ler conversa; estar impedido |

Se o administrador de toda operação é o comandante, ou só na missão de campo, é pergunta aberta (pergunta 9 do 01c; 01c-C25). Quem ocupa cada papel por instituição: D-R33 (01e-E18b, em parte).

### 2. Atributos verificados em toda ação (regra 5)
- **Membro da sala:** no `server/`, pela pertença verificada a partir da declaração assinada (ADR-0005 item 4); nos clientes, pelo estado MLS.
- **Credencial ≥ grau da sala** (doc 08 seção 5).
- **Compartimento e escopo de análise** (doc 08 seção 5.1).
- **Validade do papel.**
- **Instituição** (instância dedicada, doc 09).

### 3. Regras do modelo
- **Nenhum papel dá leitura de conversa.** Só ser membro MLS da sala dá.
- **Separação de funções imposta pelo sistema:** as combinações proibidas do doc 04 seção 6.2 (entre elas administrador de sistema × detentor de quórum × analista, e analista × gestor do sistema, 01e-E18c) são recusadas na atribuição.
- **Atribuição por duas pessoas, só para papéis privilegiados** (administrador de operação, comandante e subcomandante, aprovadores, segunda assinatura, autorizadores, analista, criador da sala de análise, detentor de quórum, Auditor), registrada, com prazo de validade e motivo; renovação explícita. Quem são as duas por instituição: D-R33. O cadastro do agente segue a RN05 (um gestor, presencial) e o D20 (01e-E10d).
- **Raiz de confiança fora do diretório (01e-E10):** o diretório é hostil. A atribuição de papel privilegiado é um **objeto assinado** pelas duas pessoas, com chaves que constam de uma **lista assinada em cerimônia** (a mesma dos aprovadores da revinculação, doc 07 seção 7.4). Clientes, nó de IA e estação isolada conferem o objeto contra essa lista, nunca contra o registro do diretório. A segunda assinatura da sala comum (ADR-0003 seção 1), a concessão de escopo (ADR-0008) e os atos da ADR-0010 são conferidos da mesma forma. **Vínculo com o diretório (01g-G05):** a lista leva, para cada chave, o **identificador opaco do usuário no diretório**, para que a estação isolada confira o impedimento por posição da ADR-0010 sem quebrar o pseudonimato de terceiros; participante sem vínculo conferível é tratado como impedido (falha fechada).
- **Detentores de cartão:** as combinações proibidas com detentor de quórum são conferidas na **cerimônia de emissão** dos cartões, contra a lista assinada de papéis. A identidade dos detentores não fica no diretório em claro (fica na lista da cerimônia, ou selada), para não criar a lista de quem coagir (R80; 01e-E10c).
- **Papéis e combinações proibidas** definidos no `instituicao.yaml` assinado (RNF-T3); mudança por cerimônia registrada.
- **Negação por padrão:** ação sem regra explícita é recusada e registrada.
- A decisão de autorização dos **endpoints** fica num ponto só do `server/` (módulo de diretório, chamado por `api/`; regra 9), com teste de tabela por papel e atributo em `tests/security`.
- **Ações dentro do grupo** (commit, ordens assinadas, marca, conclusão, aprovação de revinculação, concessão de escopo) são conferidas também pelos **clientes**, contra as chaves e papéis do estado autenticado do grupo ou da lista assinada em cerimônia (ADR-0005). O ponto único do `server/` não substitui essa conferência: servidor hostil não concede papel.

## Alternativas rejeitadas
- **Só RBAC:** não expressa sala, grau e compartimento.
- **Só ABAC com política livre:** difícil de auditar e de homologar.
- **Papel de "super administrador":** contraria a regra 14 e a separação de funções.

## Consequências
- O diretório guarda papéis com prazo e o motivo da atribuição; a coluna única `users.role` do doc 02 seção 7.5 vira atribuições com prazo (Fase 1).
- Riscos: R108 (acúmulo de papéis), R109 (papel vencido não revogado). Requisitos: RF89, RF90.

## Revisar quando
O jurídico e a chefia definirem os papéis por instituição (D-R24, D-R33, D-R36, D-R37), a pergunta 9 do 01c for respondida, ou na Fase 1, quando o diretório for implementado. Precisa estar decidida antes da Fase 1.
