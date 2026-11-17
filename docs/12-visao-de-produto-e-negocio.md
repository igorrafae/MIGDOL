# 12 — Visão de produto e de negócio

**Projeto:** MIGDOL
**Versão do documento:** 0.1, rascunho
**Público:** responsável pelo produto, chefia, assessoria jurídica

> **Fora do escopo das Fases 0 a 8.** Este documento registra a visão de longo prazo do responsável pelo produto, ou seja, o que vem depois da v1 e da Fase 8. Nada daqui muda o roteiro do doc 03 nem permite adiantar fase: a v1 roda em aparelho de linha corporativa com Android Enterprise (ADR-06, ADR-0007). Valores e regras de contratação são hipóteses **[conferir]**. Não é parecer jurídico nem recomendação financeira.

---

## 1. Visão do responsável pelo produto

1. **Plataforma própria completa:** hardware montado no Brasil, a partir de componentes comprados separadamente, com **sistema operacional próprio** e **aplicativo próprio** (ADR-0011, proposta).
2. **Duas linhas no mesmo aparelho:**
   - **MIGDOL institucional:** exclusivo para militar e para polícia;
   - **produto civil:** caçadores, pescadores e expedições em área sem sinal.
3. **Caminho de caixa:** primeiro contrato de software com uma unidade especializada; depois, um contrato de maior porte que financie a célula de montagem e provisionamento (proposta: Cáceres, MT); a assinatura civil como receita recorrente.

## 2. Etapas propostas

| Etapa | O que se vende | Hardware | Condição para entrar |
|---|---|---|---|
| **A. Software** (v1 até piloto) | Licença ou assinatura por instância e por usuário; implantação; condução das cerimônias; suporte com releases assinados; treinamento; lista de aparelhos homologados | Aparelho de linha corporativa comprado pela instituição (Android Enterprise, StrongBox, atestação) | Fase 7 aprovada (V01–V10) |
| **B. Kit de campo** (Fase 8) | Nó tático montado e provisionado, mais o software | Kit com rádio IP, bateria e servidor portátil | ADR-0004 aceita; levantamento dos rádios (D-R25) |
| **C. Plataforma própria** (depois da v1) | Aparelho próprio, SO próprio e app | Montagem a partir de componentes (ADR-0011) | Estudo de viabilidade aprovado; equipe de SO; contrato que pague o investimento (seção 5) |

Em todas as etapas, **o fornecedor não opera a produção nem guarda chaves da instituição** (RNF-T2, ADR-13): HSM, cartões do quórum, AC e `instituicao.yaml` pertencem à instituição. Na etapa C, a chave de assinatura do SO e do boot de cada instituição é gerada em cerimônia da instituição; o fornecedor assina só o build de referência, que a instituição verifica e reassina.

## 3. Linha institucional e linha civil

**Regra de separação (proposta D-R50):** o mesmo **hardware** pode servir às duas linhas, como um modelo de mercado serve a muitos clientes. A separação fica no software e nas chaves:

| Item | Institucional | Civil |
|---|---|---|
| Nome e marca | MIGDOL | Outro nome |
| Repositório | Este | Separado |
| SO | Build institucional, boot travado com chave da instituição | Build civil, chave civil própria |
| App e protocolo configurado | Core MIGDOL, regras 1 a 14 | App próprio; pode usar bibliotecas abertas genéricas (mapa offline, NFC), nunca o core configurado |
| Servidores | Instância dedicada na instituição | Infraestrutura civil própria |
| Exportação | Proibida fora do rito (regras 3, 10, 14) | Pode exportar diário e conversas para mensageiro comercial por escolha do usuário |
| Atestação | Raiz da instituição; o aparelho civil **não** passa | Raiz civil |

Por quê: o aparelho civil vai parar em qualquer mão. Ele pode ser desmontado e estudado, mas não pode ensinar nada que abra o institucional. Com a mesma placa e boot, SO, chaves e servidores diferentes, o que se aprende no civil não vale no institucional. Um incidente no produto civil também não pode contaminar a confiança da homologação.

## 4. Contratação pública [conferir]

- **Primeiro contrato por dispensa de valor** (Lei 14.133/2021, art. 75, II): conferir o limite vigente e a vedação a fracionamento. Hardware no mesmo contrato tende a passar do limite e muda a modalidade.
- **Valor do primeiro contrato** (hipótese: cerca de R$ 45 mil por ano): cobre software e suporte. Não cobre HSM, servidores nem MDM da instituição, que são compra da própria instituição (ADR-13).
- **Conflito de interesse:** se o fornecedor, um sócio ou alguém da equipe tiver vínculo funcional com o órgão contratante, a contratação pode ser vedada ou restrita (Lei 14.133 e estatuto do servidor; risco de improbidade). **Parecer jurídico antes de qualquer proposta** (R131).
- **Dispositivo de radiofrequência:** homologação na Anatel do aparelho próprio, e uso das faixas licenciadas da instituição.
- **Fabricação no Brasil:** avaliar certificações e incentivos aplicáveis à montagem local [conferir].

## 5. Aparelho candidato de referência (Blackview Xplore 1, versão com rádio)

Registrado como **candidato na avaliação de aparelhos (D-R19)** e como **referência de requisitos** para a etapa C. **Não é base da v1.**

| Ponto | Situação |
|---|---|
| Rádio VHF/UHF 5 W integrado, bateria de 20.000 mAh, IP68/IP69K, MIL-STD-810H, câmera noturna | Atende ao perfil de campo (doc 10) |
| Fabricante estrangeiro e SoC MediaTek | Cadeia de suprimentos sob adversário estatal (R88, R129) |
| StrongBox e atestação de hardware | **Verificar no aparelho.** Sem eles, caem ADR-0002, ADR-0006 e a proteção contra clonagem |
| Travar o bootloader com chave própria (AVB) | Verificar; sem isso, não há SO próprio confiável |
| Rádio integrado | É módulo de voz analógico ou DMR, controlado pelo app do fabricante, sem transporte IP. Só vale como segunda camada (D-R27) se houver interface aberta |
| Frequências e Anatel | Homologação e faixas da instituição |

Teste proposto, sem dado real: atestação e StrongBox, trava de boot com chave própria, interface do módulo de rádio, autonomia. **Sem editar o SO na v1** (D-R48).

## 6. Requisitos mínimos da plataforma própria (etapa C)

1. **Elemento seguro dedicado**, equivalente ao StrongBox, com chave que não sai do chip e atestação para a **raiz da instituição** (regra 12, sem terceiros).
2. **Boot verificado** com chave da instituição gerada em cerimônia, com proteção contra retorno a versão antiga.
3. **SoC com documentação e código de kernel e drivers** suficientes para manter o SO, e **fornecedor que entregue correções de segurança** pelo prazo de vida do aparelho.
4. **SO derivado do AOSP**, sem serviços do Google (regra 10), com correções de segurança mensais, MDM próprio e build reproduzível.
5. **Módulo de rádio com interface aberta** documentada, para o core usar como transporte (doc 10) sem depender do app do fabricante.
6. **Cadeia de custódia física:** compra rastreada dos componentes, inspeção, montagem e provisionamento em ambiente controlado, com registro por número de série.
7. **Homologação:** Anatel, órgão homologador da instituição (V11) e red team de hardware (V12).

## 7. Riscos (doc 04 seção 5.6)
R129 (cadeia de suprimentos do aparelho), R130 (SO próprio sem sustentação), R131 (conflito de interesse na contratação), R132 (fornecedor na cadeia de confiança), R133 (linha civil compartilhada revela o institucional), R134 (montagem própria sem cadeia de custódia física).

## 8. Decisões
| ID | Decisão | Estado |
|---|---|---|
| D-R48 | Sem SO próprio nem edição de SO na v1; aparelho de linha corporativa com Android Enterprise | Proposta (recomendada) |
| D-R49 | Plataforma própria (hardware montado a partir de componentes, SO e app próprios) como etapa C, depois da v1 (ADR-0011) | **Visão do responsável**; depende do estudo de viabilidade |
| D-R50 | Regra de separação entre linha institucional e civil (seção 3) | Proposta |
| D-R51 | Modelo de venda por etapa (seção 2): software primeiro, kit de campo na Fase 8, plataforma própria depois | Proposta |
| D-R52 | Parecer jurídico sobre conflito de interesse e modalidade de contratação antes da primeira proposta comercial | Obrigatória |

## 9. Pontos para validação
1. Limite vigente da dispensa e enquadramento do objeto [conferir].
2. Vínculo funcional da equipe com órgãos contratantes [conferir].
3. Viabilidade técnica da etapa C: SoC com suporte longo, elemento seguro e módulo de rádio aberto.
4. Custo da equipe de SO (correções mensais por anos) diante da receita esperada.
5. Homologação Anatel e incentivos à montagem local [conferir].
