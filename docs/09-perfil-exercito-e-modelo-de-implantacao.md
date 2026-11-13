# 09 — Perfil Exército (Inteligência Militar) e Modelo de Implantação

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** chefia, área de segurança, assessoria jurídica, equipe de desenvolvimento

> O MIGDOL foi desenhado para uma unidade de inteligência policial. Este documento registra o que muda para atender a inteligência do Exército e decide o **modelo de implantação** para qualquer instituição. As normas militares citadas **não foram lidas**: tudo marcado **[conferir]** depende da assessoria jurídica da Força. Não é parecer jurídico.

---

## 1. Resposta curta

| Pergunta | Resposta |
|---|---|
| Está pronto para o Exército? | **Não.** Nada implementado; sem algoritmo de Estado (D-R18), sem homologação, sem credenciamento; normas militares não analisadas |
| É ponta a ponta? | No uso normal, **sim por desenho**: o servidor só vê cifrado. Com duas ressalvas declaradas: o **nó de IA** é uma ponta (membro declarado), e a **custódia** (ADR-0003) existe só em **sala sob custódia**. No perfil Exército a flag vem **ativa**, com o motivo investigação militar; a instituição decifra só por quórum presencial em HSM offline, com base registrada. Nome correto: **ponta a ponta, com custódia institucional sob quórum nas salas marcadas** |
| Tenant ou local? | **Instância dedicada, hospedada na infraestrutura da própria Força.** Nunca tenant em ambiente compartilhado (seção 2) |

---

## 2. Modelo de implantação: instância dedicada por instituição (decisão D-R21)

### 2.1 Decisão proposta
- **Uma base de código, instâncias isoladas.** Cada instituição (polícia, Exército, outra) recebe a sua **própria instalação**, na **sua própria rede e datacenter**, operada pelo **seu** pessoal credenciado.
- **Nada é compartilhado entre instituições:** servidores, banco, HSM, autoridade certificadora, cartões de comandantes, MDM, nó de IA, chaves de custódia, logs.
- O desenvolvimento entrega **binários assinados e reproduzíveis** (doc 08 seção 4.5); a instituição verifica e instala. O fornecedor **não tem acesso** à produção.
- Já é compatível com o doc 02 seção 10 (rede da instituição, sem nuvem) e com a regra 10 (sem terceiros).

### 2.2 Por que não multi-tenant

| Motivo | Detalhe |
|---|---|
| Decreto 7.845, art. 39 | Equipamento que produz informação classificada fica isolado ou em canal **física ou logicamente separado** |
| Raízes de confiança próprias | Cada Força precisa da sua AC, do seu HSM e do seu quórum. Compartilhar significa que o operador do ambiente vira parte da cadeia de confiança de todos |
| Metadados | Num ambiente comum, quem opera vê **quem fala com quem** em todas as instituições: alvo de altíssimo valor |
| Raio de dano | Um incidente ou uma falha de isolamento entre tenants expõe todas as instituições de uma vez |
| Regimes jurídicos diferentes | Polícia estadual, Polícia Federal e Forças Armadas têm normas, corregedorias e justiça diferentes; ordens judiciais e prazos não se misturam |
| Credenciamento | Administrador de um ambiente comum precisaria de credencial em todas as instituições |
| Complexidade | Multi-tenant exige isolamento por tenant em cada serviço, consulta e chave; é uma classe inteira de falhas que a instância dedicada elimina |

### 2.3 Variações aceitas dentro da Força
- **Instância central** da Força (comando de inteligência), com **salas por grau e por unidade** e credencial por grau (doc 08 seção 5).
- **Instância separada por compartimento**, quando a chefia exigir que nem os metadados de um compartimento fiquem no mesmo servidor (decisão D-R21b).
- **Nó tático de campo** (seção 4), subordinado à instância central.

### 2.4 Comunicação entre instituições
- **Fora do escopo da v1.** Se for necessária (ex.: operação conjunta polícia–Exército), exige ADR própria com federação controlada por gateway, sem compartilhar chaves nem diretório. Até lá: salas de operação conjunta ficam numa **única** instância, com convidados credenciados por aquela instituição.

---

## 3. Normas e ritos que mudam [conferir]

| Tema | Polícia (desenho atual) | Exército (a confirmar) |
|---|---|---|
| Informação classificada | Decreto 7.845 ou norma estadual | **Decreto 7.845 aplica-se diretamente** (Executivo federal) e normas do Ministério da Defesa e do Exército sobre salvaguarda de assuntos sigilosos |
| Sistema de inteligência | Unidade de inteligência policial | Sistema de inteligência do Exército e SISBIN (Lei 9.883/1999) |
| LGPD | Art. 4º, III, segurança pública | Art. 4º, III, **defesa nacional e segurança do Estado** |
| Disciplina | Corregedoria, procedimento administrativo | Regulamento Disciplinar do Exército e sindicância militar |
| Crimes por vazamento | Código Penal, arts. 325 e 314 | **Código Penal Militar** (Decreto-Lei 1.001/1969), crimes de revelação de segredo e violação de sigilo (artigos a conferir) |
| Ordem judicial e custódia | Justiça comum. Custódia **desligada por padrão**; pode ligar com motivo corregedoria ou inquérito, quando o edital exigir (GD-2) | **Justiça Militar da União** quando couber. Custódia **ativa**, motivo investigação militar. Rito da ADR-0003 igual: quórum presencial, CI e base (ordem judicial, cumprida; ou processo administrativo com validação jurídica). Com a flag ativa, custódia é o padrão; sala comum só com motivo e segunda assinatura (GD-14) |
| Termo de ciência | Assinatura via gov.br, contingência em três vias | Rito militar: ciência registrada em sistema da Força ou assinatura ICP-Brasil; gov.br só se a Força adotar (D-R23) |
| Quórum de custódia | Comandantes da companhia | Cadeia de comando da OM ou da área de inteligência (definir quem são os N detentores) |

---

## 4. Operação em campo, sem internet (decisão D-R22)

O desenho atual exige conexão periódica: o lease de 12 h protege o aparelho perdido, mas em campo o aparelho pode ficar dias sem rede.

| Opção | Como funciona | Custo |
|---|---|---|
| **A. Lease longo por perfil** | Salas de campo com lease maior (ex.: 72 h) e câmera, PIN de coação e escada de resposta obrigatórios | Aparelho perdido fica legível por mais tempo |
| **B. Nó tático de campo** | Servidor portátil (kit de campo) atua como serviço de entrega local para os grupos da missão, numa rede local (Wi-Fi ou rádio IP); sincroniza com a instância central ao voltar | Novo componente a endurecer; ordenação MLS entre nó de campo e central precisa de desenho |
| **C. Sem operação offline na v1** | Campo usa os meios táticos já existentes da Força; MIGDOL só com conexão | Limita o uso |

**Decidido (D-R22): opção B, operação em campo por rádio, para a polícia e para o Exército**, como **Fase 8** depois da v1 (na v1 vale a C). Desenho em **docs/10-operacao-em-campo-por-radio.md** e **docs/adr/0004** (proposta revisada). HF fica fora do escopo do aplicativo. Na missão, a central fica congelada; o nó só ordena; o administrador da missão é o **comandante da missão**, e o substituto, o **subcomandante** (GD-11).

---

## 5. Adversário e requisitos que passam a obrigatórios

Para a inteligência militar, o adversário de referência é **serviço de inteligência estrangeiro**. Os itens do doc 08 seção 4, hoje "importantes", passam a **obrigatórios antes do piloto**:
- algoritmo de Estado como segunda camada (D-R18) e troca de chaves pós-quântica híbrida (D-R12);
- proteção de metadados: padding, tráfego de cobertura, retenção mínima (D-R20);
- rede: APN privada ou rede da Força; política contra ERB falsa; chip só de dados;
- aparelho e sistema escolhidos e avaliados (D-R19); compra com cadeia de custódia;
- build reproduzível; auditoria externa e red team com cenário de adversário estatal;
- credenciamento de todo o pessoal com acesso (usuários, administradores, desenvolvedores);
- segurança física e de emissões nas salas de servidores e HSM;
- em missão de campo, escada de resposta no **perfil missão**, com destruição do aparelho diante de rosto de outra pessoa (duas capturas) com PIN errado ou de PIN de coação, cada gatilho ligado só depois das medições da GD-24, e destruição imediata do nó de campo capturado (GD-18; doc 07 seção 7.1a; ADR-0004 item 13).

---

## 6. Riscos novos

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R84 | Implantação compartilhada entre instituições expõe metadados e chaves de todas | Médio | Crítico | Instância dedicada por instituição (D-R21) | Baixo |
| R85 | Fornecedor do software com acesso à produção | Médio | Crítico | Binário assinado e reproduzível; instalação pela própria Força; sem acesso remoto do fornecedor | Baixo |
| R86 | Uso em campo sem conexão força lease longo e aumenta a janela do aparelho perdido | Médio | Alto | Nó de campo e lease de campo com teto (doc 10, ADR-0004); v1 sem offline | Médio |
| R87 | Rito do termo e da apuração incompatível com a norma militar invalida prova ou sanção | Médio | Alto | D-R23; assessoria jurídica da Força antes do piloto | Médio |
| R88 | Adversário estatal estrangeiro explora metadados, rede ou cadeia de suprimentos | Alto | Crítico | Seção 5 obrigatória antes do piloto | Médio |

---

## 7. Decisões novas

| ID | Decisão | Quem decide |
|---|---|---|
| D-R21 | Modelo de implantação: **instância dedicada por instituição**, hospedada na infraestrutura dela (recomendado); D-R21b: instância separada por compartimento | Chefia + segurança |
| D-R22 | **Decidido: B, por rádio, Fase 8** (doc 10) | Responsável pelo produto |
| D-R23 | Rito do termo de ciência e da apuração para militares | Assessoria jurídica da Força |
| D-R24 | Quem são os N detentores do quórum na cadeia de comando | Comando |

---

## 8. Requisitos novos

| ID | Requisito | Prioridade |
|---|---|---|
| RNF-T1 | Instância dedicada por instituição: nenhum servidor, banco, HSM, AC, MDM ou chave compartilhado entre instituições | Essencial |
| RNF-T2 | Entrega ao cliente só por binário assinado e reproduzível; fornecedor sem acesso à produção | Essencial |
| RNF-T3 | Configuração por instituição (normas, rito do termo, papéis, prazos) sem mudar o código | Importante |

---

## 9. Pontos para validação [conferir]
1. Normas do Ministério da Defesa e do Exército sobre salvaguarda de assuntos sigilosos e sistemas de comunicação.
2. Artigos do Código Penal Militar aplicáveis a vazamento e acesso indevido.
3. Regulamento disciplinar e rito de sindicância para uso dos indícios do sistema (doc 05).
4. Órgão que homologa recurso criptográfico e sistema para a Força.
5. Se a Força já tem algoritmo de Estado e infraestrutura de chaves próprias.
