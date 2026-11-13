# 08 — Informação Classificada e Requisitos de Nível de Inteligência

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** chefia, área de segurança, jurídico, equipe de desenvolvimento

> "Nível CIA" não é certificação. O que existe são **regimes de informação classificada**. No Brasil, o regime federal é o da Lei 12.527/2011 (graus de sigilo) e do Decreto 7.845/2012 (credenciamento e tratamento). Este documento enquadra o MIGDOL nesse regime e lista o que falta para um padrão de serviço de inteligência. Nada aqui foi implementado. Não é parecer jurídico.

---

## 1. Situação atual (honesta)

| Item | Estado |
|---|---|
| Código | Esqueleto (serviços Go com negação por padrão); sem criptografia implementada, sem teste real |
| Decisões | D-R1 a D-R41 com parte em aberto. ADR-0003 (custódia) e ADR-0004 (campo) revisadas na etapa 1b da Fase 0 e nas rodadas seguintes; ADR-0005 (commits) e ADR-0006 (lease) propostas; ADR-0007 (plataformas) **aceita** (GD-19); ADR-0008 (painel e IA de análise, redesenhada pela GD-15), ADR-0009 (RBAC + ABAC) e ADR-0010 (investigação sigilosa) propostas; ADR-0002 (atestação) proposta no doc 07 seção 8, arquivo como entrada do spike de atestação. D-R17 e D-R18 ficam para o portão G3 (GD-9) |
| Jurídico | Cerca de 28 pontos marcados "validar" ou "conferir" |
| Avaliação externa | Nenhuma; pentest previsto em três rodadas (doc 03) |

O MIGDOL é hoje um **projeto documentado**, não um aplicativo seguro. A segurança só existe depois de implementada, testada e avaliada por terceiros.

---

## 2. Enquadramento no Decreto 7.845/2012

Texto lido no [Planalto](https://www.planalto.gov.br/ccivil_03/_ato2011-2014/2012/decreto/d7845.htm). O decreto vale para o **Poder Executivo federal** (art. 1º) e para entidades com vínculo que tratem informação classificada (art. 17, parágrafo único). Para polícia **estadual**, conferir a norma do estado **[conferir]**.

### 2.1 Graus de sigilo
- Os graus legais são **reservado, secreto e ultrassecreto** (Lei 12.527/2011). "Restrito" não é grau legal; foi corrigido no doc 04 seção 6.1.
- **Decisão D-R17: grau máximo que o MIGDOL trata.** Proposta: **reservado e secreto**. Ultrassecreto, pelo que se conhece de outros países, não trafega em celular comercial com aplicativo; exige equipamento homologado e ambiente controlado **[conferir]**.
- Cada sala tem grau definido na criação; relatório de IA herda o grau mais alto das salas analisadas (já no doc 04).
- **Decisão do responsável (GD-9):** D-R17 e D-R18 ficam para **depois dos spikes**, no portão G3. **Até lá, nenhum grau classificado é tratado.**

### 2.2 Exigências do decreto e lacunas

| Artigo | Exigência (resumo) | Situação no MIGDOL | Ação |
|---|---|---|---|
| Art. 40 | Cifração de informação classificada com **recurso criptográfico baseado em algoritmo de Estado**; parâmetros pelo GSI | MLS com algoritmos comerciais **não atende sozinho** | Segunda camada com biblioteca fornecida pelo Estado (D-R18); contato com GSI e CEPESC |
| Art. 41 | Inspeções, inventário de material criptográfico, comunicação de anomalias | Inventário de chaves existe em parte (doc 04 seção 6.4) | Inventário de material criptográfico e rotina de inspeção |
| Art. 38, §1º | Transmissão por **canal seguro** | mTLS 1.3 + MLS | Confirmar que atende o padrão do GSI |
| Art. 38, §2º | Identidade do usuário garantida **no mínimo por certificado digital** | Certificado de aparelho emitido pela AC interna | Confirmar se a AC interna basta ou se exige ICP-Brasil **[conferir]** |
| Art. 38, §3º | Níveis de controle de acesso e criptografia adequados a cada grau | Autorização por sala; sem diferença por grau | Política por grau (seção 3) |
| Art. 38, §4º | **Registro de acessos** autorizados e não autorizados, guardado **no mínimo enquanto durar a restrição** | Retenção curta da trilha (D-R4) e expurgo | **Conflito** (seção 2.3) |
| Art. 39 | Equipamento que produz documento classificado isolado ou em canal **física ou logicamente separado** | App em internet móvel comum | APN privada ou VPN dedicada (seção 4.2) |
| Art. 12 | **Credencial de segurança** para quem acessa | Não tratado | Credenciamento de usuários, desenvolvedores, administradores e fornecedores (seção 5) |
| Arts. 27, 28, 31 | Transmissão e meio eletrônico com criptografia compatível com o grau | Ver art. 40 | Idem |

### 2.3 Conflito: registro de acessos × expurgo
- O art. 38, §4º manda guardar o registro de acessos enquanto durar a restrição (até 25 anos no ultrassecreto, 15 no secreto, 5 no reservado, pela Lei 12.527).
- O projeto prevê expurgo de conteúdo (RN03) e retenção curta da trilha (D-R4).
- **Leitura proposta:** o expurgo destrói **conteúdo**; o **registro de acesso** (quem acessou qual sala e quando, sem conteúdo) é guardado pelo prazo do grau. Isso também ajuda no conflito com a Lei 8.159 (R60, D-R5).
- O registro de acesso é dado pessoal do agente: cifrado, com acesso por rito (regra 13).
- **[conferir]** com o jurídico.
- **Revisão de desenho (D28), adiado para antes da Fase 4, junto com D-R4 e D-R5:** a ADR-0003 já separa uma chave de quórum por classe de dado, o que permite prazos diferentes. Faltam: hierarquia de chaves por grau (a cópia de 5 anos e o registro de 15 anos não podem usar a mesma chave); registro de acesso gerado no servidor, cifrado para quórum, incluindo **negações**; identificador de sala retido para o registro; `audit_log.actor` sem identificador em claro. Sem grau classificado até G3, o conflito com 15 e 25 anos não se aplica ainda.

---

## 3. Política por grau (proposta)

| Grau | Perfil de local mínimo (doc 07) | Camadas de cifragem | Câmera | IA | Retenção do registro de acesso |
|---|---|---|---|---|---|
| Reservado | Livre ou rede autorizada | MLS + algoritmo de Estado | Opcional | Permitida | 5 anos |
| Secreto | Local autorizado | MLS + algoritmo de Estado + canal separado | Obrigatória | Por decisão da chefia | 15 anos |
| Ultrassecreto | Fora do escopo (D-R17) | — | — | — | — |

---

## 4. O que um padrão de serviço de inteligência exige além do decreto

### 4.1 Criptografia
- **Duas camadas independentes:** implementações, chaves e fornecedores diferentes, de modo que a falha de uma não exponha o conteúdo. Uma camada é MLS; a outra, o algoritmo de Estado.
- **Pós-quântico (D-R12):** segredos de 15 a 25 anos estão expostos ao risco de "colher agora, decifrar depois". Adotar troca de chaves híbrida (clássica + pós-quântica) assim que houver suíte madura na biblioteca MLS **[verificar suporte]**. Vale também para a **cópia de custódia** (HPKE com KEM híbrido), que guarda até 5 anos (D08; ADR-0003). A segunda camada de Estado, quando houver, também cobre a custódia.
- **Módulos criptográficos validados ou homologados:** registrar quais bibliotecas, versões e validações; regra 1 continua (nada à mão).

### 4.2 Rede e metadados
- **Metadados são o alvo de um adversário estatal:** quem fala com quem, quando e de onde. Hoje o servidor vê sala opaca, horário e tamanho.
  - Tamanho padronizado (padding) de mensagens e anexos.
  - Tráfego de cobertura em intervalos regulares, para esconder horário de atividade.
  - Ocultação do remetente para o servidor, quando o protocolo permitir **[avaliar]**.
  - Retenção mínima de metadados no servidor.
- **Canal separado (art. 39):** APN privada da operadora ou VPN dedicada até a rede da instituição.
- **Rede celular:** o chip expõe a localização do agente à operadora; ERB falsa (IMSI catcher) e baseband são vetores. Medidas: chip só de dados, sem número publicado; política de rede (4G/5G, sem 2G); avaliar aparelhos com isolamento de baseband.

### 4.3 Aparelho e sistema
- **Decisão D-R19: sistema e aparelho.** Sem serviços do Google (regra 10). Avaliar Android endurecido com bootloader travado com chave da instituição e compatibilidade com o MDM.
- **Compra controlada:** cadeia de custódia do aparelho desde o fornecedor (risco de interceptação e implante na entrega); triagem (doc 07 seção 2).

### 4.4 Pessoas
- **Credenciamento (art. 12)** de usuários, desenvolvedores, administradores de infraestrutura e fornecedores.
- **Programa de ameaça interna:** indicadores, canal de reporte, revisão periódica de acessos, separação de funções (doc 04 seção 6.2).
- **Duas pessoas para atos críticos:** cerimônia de chaves, release, apagamento em lote, abertura de custódia (já em parte).
- **Treinamento** com exercícios de engenharia social e de coação (doc 05 seção 7).

### 4.5 Cadeia de software
- **Build reproduzível:** dois ambientes independentes geram o mesmo binário.
- **Revisão de dependências críticas** (MLS, SQLCipher, runtime) e versão fixada com hash (regra 8).
- SBOM e assinatura em HSM (já no doc 03).

### 4.6 Ambiente físico
- Sala de servidores e HSM com controle de acesso, registro e vídeo; servidores sem acesso remoto direto.
- **Emissões eletromagnéticas:** avaliar para o nó de IA e as estações de administração **[avaliar]**.
- Estação de preparação de aparelhos isolada (doc 07).

### 4.7 Avaliação e homologação
- **Auditoria externa** do desenho criptográfico e do código, por equipe sem vínculo com o desenvolvimento.
- **Red team** com cenários de adversário estatal (metadados, cadeia de suprimentos, insider).
- **Homologação** junto ao órgão competente (GSI ou equivalente estadual) **[conferir rito]**.
- Repetição anual e a cada mudança relevante.

---

## 5. Credenciamento e acesso por grau

- Só acessa sala de grau X quem tem credencial de grau X **e** necessidade de saber.
- O diretório (server/services/directory) passa a guardar o **grau da credencial** do usuário e o **grau da sala**; a autorização por sala (regra 5) verifica os dois.
- Credencial vencida ou cassada suspende o acesso às salas daquele grau.

### 5.1 Compartimento de análise (ADR-0008 e ADR-0009, propostas; depois da v1)
- **Compartimento** aqui é um conjunto de salas com necessidade de saber própria, atributo do controle de acesso (ADR-0009). Não confundir com o **perfil de local** "compartimento" do doc 07 seção 4.3.
- A **sala de análise** é um compartimento: escopo (lista de operações), prazo e revisão definidos pela chefia (D-R32).
- **Concessão da sala de origem (GD-15):** uma sala só entra no escopo se o administrador dela conceder, por mensagem de controle assinada no próprio grupo, com CI e/ou autorização do comandante, visível aos membros e registrada. Sem concessão, a sala fica fora; a instituição só chega ao conteúdo pelo rito da regra 14. Sala de análise sob custódia só inclui salas sob custódia (ADR-0008 seção 2).
- **Credencial em todo o escopo:** só entra quem tem credencial ≥ ao maior grau das salas do escopo **e** necessidade de saber registrada para cada uma. O escopo não pode crescer sem nova autorização registrada. Credencial, papel de analista e concessão são conferidos pelos clientes e pelo nó de IA contra a lista assinada de papéis e o estado autenticado do grupo, não contra o diretório, que é hostil (ADR-0009 seção 3; 01e-E10b).
- O relatório e os indícios herdam o **maior grau** das salas do escopo (doc 04 seção 6.1). Até o portão G3, nenhuma sala de análise trata grau classificado (GD-9).
- O nó de IA só cruza salas do escopo de que é membro e com concessão vigente, com filtro antes da busca; o expurgo da origem apaga os indícios dela na sala de análise (ADR-0008 seção 2). Pergunta aberta: IA em sala secreta ou de compartimento (pergunta 15 da revisão).
- Toda pergunta à IA gera registro de acesso cifrado para a chave de quórum de registros (seção 2.3).
- Os membros das salas do escopo veem a concessão na própria sala e sabem o que a sala de análise devolve (doc 06 seção 7).

---

## 6. Riscos novos (continuação da matriz)

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R72 | Uso do sistema para informação classificada **sem algoritmo de Estado** (descumprimento do art. 40) | Alto | Alto | D-R18; não tratar grau classificado antes da segunda camada | **Alto até D-R18** |
| R73 | Registro de acessos apagado antes do prazo do grau (art. 38, §4º) | Médio | Alto | Separar registro de acesso do conteúdo; retenção por grau | Médio |
| R74 | Análise de tráfego e metadados por adversário estatal ou operadora | Alto | Alto | Padding, tráfego de cobertura, canal separado, retenção mínima | Médio |
| R75 | Rastreio do agente pela rede celular (chip, ERB falsa, baseband) | Médio | Alto | Chip só de dados, política de rede, aparelho avaliado | Médio |
| R76 | Aparelho interceptado na cadeia de compra | Baixo | Crítico | Compra controlada, triagem, reflash | Médio |
| R77 | Acesso de pessoa sem credencial do grau | Médio | Alto | Credencial no diretório; autorização por grau | Baixo |
| R78 | Segredo de longo prazo decifrado no futuro por computação quântica | Médio | Alto | Troca de chaves híbrida (D-R12) | Médio |

---

## 7. Decisões novas

| ID | Decisão | Quem decide |
|---|---|---|
| D-R17 | Grau máximo tratado pelo MIGDOL (proposta: reservado e secreto) | Chefia + jurídico |
| D-R18 | Algoritmo de Estado como segunda camada: fornecedor, biblioteca, homologação | Chefia + segurança, com GSI/CEPESC |
| D-R19 | Sistema operacional e modelo de aparelho | Segurança + desenvolvimento |
| D-R20 | Proteção de metadados (padding, tráfego de cobertura, ocultação do remetente) e custo de bateria | Segurança + desenvolvimento |

---

## 8. Requisitos novos

| ID | Requisito | Prioridade |
|---|---|---|
| RF74 | Grau de sigilo por sala e grau de credencial por usuário; autorização verifica os dois | Essencial |
| RF75 | Segunda camada de cifragem com algoritmo de Estado, independente do MLS | Essencial para grau classificado |
| RF76 | Registro de acessos separado do conteúdo, retido pelo prazo do grau | Essencial |
| RNF-I1 | Mensagens e anexos com tamanho padronizado; tráfego de cobertura configurável | Importante |
| RNF-I2 | Transmissão por canal física ou logicamente separado (APN privada ou VPN) | Essencial para grau classificado |
| RNF-I3 | Build reproduzível verificado em dois ambientes | Importante |

---

## 9. Ordem decidida (GD-9)

A ordem anterior punha D-R17 e D-R18 antes dos spikes; o pipeline da Fase 0 os punha no portão G3 (D33). O responsável decidiu:

1. Spikes da Fase 0 (MLS, atestação, inferência; doc 03 seção 5).
2. **Portão G3:** D-R17 (grau máximo e enquadramento no Decreto 7.845 ou na norma do ente) e D-R18 (algoritmo de Estado com GSI ou CEPESC, e homologação).
3. Resolver o conflito do registro de acessos com o expurgo, junto com D-R4, D-R5 e R60 (D28), antes da Fase 4.
4. Documentar e testar a proteção de metadados e a rede celular.

**Até o portão G3, nenhum grau classificado é tratado** pelo MIGDOL.

## 10. Pontos para validação

1. Norma do ente (federal ou estadual) sobre informação classificada.
2. Normas complementares do GSI sobre recurso criptográfico e homologação (não lidas).
3. Se a AC interna atende o art. 38, §2º.
4. Prazo de guarda do registro de acessos e relação com a LGPD.
5. Rito de credenciamento aplicável a fornecedores de desenvolvimento.
