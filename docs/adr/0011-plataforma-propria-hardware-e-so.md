# ADR-0011 — Plataforma própria: hardware montado, SO e app próprios

**Status:** proposta para **depois da v1** (visão do responsável pelo produto, 17/11/2026; doc 12) · **Não altera** ADR-06 nem ADR-0007 na v1 · **Complementa** ADR-0002 (atestação), ADR-13 (instância dedicada), ADR-15 (pacote assinado), doc 10 (rádio)

## Contexto
O responsável quer, no longo prazo, comprar componentes separados, montar o aparelho no Brasil e usar **SO e app próprios**, com duas linhas no mesmo hardware: institucional (militar e polícia) e civil. Hoje o desenho depende de StrongBox, atestação de hardware, boot travado e correções de segurança do fabricante (ADR-0002, ADR-0006, regras 2 e 12).

## Decisão (proposta)
1. **Na v1, nada muda:** aparelho de linha corporativa, Android Enterprise, sem SO próprio nem edição de SO (D-R48).
2. **Etapa C, depois da v1:** plataforma própria, só se um **estudo de viabilidade** aprovar os requisitos do doc 12 seção 6 (elemento seguro, boot com chave da instituição, SoC com suporte longo, SO derivado do AOSP com correções mensais, rádio com interface aberta, cadeia de custódia física, homologação).
3. **Chaves de boot e de SO por instituição**, geradas em cerimônia da instituição. O fornecedor assina só o build de referência, que a instituição verifica e reassina (RNF-T2).
4. **Atestação para a raiz da instituição**, sem terceiros (regra 12). O aparelho civil não passa na atestação institucional.
5. **Linha civil separada** (D-R50): mesmo hardware; nome, repositório, SO, chaves, app e servidores diferentes; nunca o core MIGDOL configurado.

## Alternativas rejeitadas
- **Editar o SO de um aparelho comercial na v1:** depende do kernel e dos drivers do fabricante, de travar o bootloader com chave própria e de correções que o fabricante pode não entregar. Vira uma frente maior que o próprio MIGDOL.
- **Mesmo SO e mesmo app para civil e institucional:** o aparelho civil ensinaria o adversário a atacar o institucional (R133).
- **Fornecedor guardando as chaves de boot:** coloca o fornecedor na cadeia de confiança (R132).

## Consequências
- Equipe de SO permanente, com correções mensais pelo prazo de vida do aparelho (R130).
- Laboratório de montagem com cadeia de custódia física (R134).
- Homologação na Anatel e no órgão da instituição.
- Riscos R129 a R134 (doc 04 seção 5.6).

## Revisar quando
O estudo de viabilidade terminar, ou houver contrato que pague o investimento da etapa C (D-R51).
