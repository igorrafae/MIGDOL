# ADR-0004 — Nó tático de campo e lease de campo

**Status:** proposta · **Contraria** ADR-05 (lease de 12 h emitido só pela central) · **Complementa** ADR-0001 (mesmo binário, perfil de campo) e ADR-0003 (custódia) · **Fase 8** (doc 10)

## Contexto
O responsável pelo produto decidiu que o MIGDOL terá operação em campo por rádio, para a polícia e para o Exército (D-R22). Em campo não há conexão com a central por horas ou dias. Com a ADR-05, o aparelho perde o acesso em 12 h sem a central.

## Decisão (proposta)
1. **Nó tático de campo:** o mesmo binário do servidor (`cmd/server`) em perfil de campo, com gateway, entrega e diretório da missão. Sem nó de IA, sem custódia aberta.
2. **Lease de campo:** durante a missão, o nó de campo emite o `segredo_servidor` com validade de 12 h renováveis enquanto houver contato, e **teto de 72 h** sem contato (D-R26).
3. **Grupo de missão:** montado na base; em campo, só o nó de campo faz commits MLS que alteram membros.
4. **Sincronização:** ao reconectar, o nó envia mensagens cifradas, cópias de custódia e auditoria; a central prevalece em conflito de grupo; o nó apaga o sincronizado.
5. **Credencial de campo** do nó, emitida pela AC da instituição, válida só durante a missão e revogável.
6. **O rádio é só transporte IP.** Nenhum protocolo de rádio é implementado no projeto.

## Alternativas rejeitadas
- **Lease longo sem nó de campo:** aparelho perdido legível por dias, sem ninguém para revogar.
- **Sem operação em campo:** contraria a decisão D-R22.
- **Cada aparelho como servidor (malha sem nó):** bifurcação de época MLS e diretório distribuído sem dono.
- **Satélite comercial estrangeiro:** regra 10.

## Consequências
- Novo componente a endurecer e a perder (R89).
- Janela maior para aparelho perdido em campo (R92).
- Reconciliação de grupo entre campo e central precisa de teste (F-R3).
- O servidor ganha um **perfil de campo**; as fronteiras de módulos (ADR-0001) não mudam.

## Revisar quando
Os rádios forem levantados (D-R25) e houver medição de banda e bateria (F-R1).
