# 10 — Operação em Campo por Rádio

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** chefia, área de segurança, equipe de comunicações, equipe de desenvolvimento

> **Decisão do responsável pelo produto:** o MIGDOL terá operação em campo por rádio, para a polícia e para o Exército (D-R22, opção B do doc 09). Entra como **Fase 8**, depois da v1 (doc 03 seção 5): não se adianta fase. O desenho exige a **ADR-0004** antes do código, porque muda o lease de 12 h (ADR-05). Nada aqui foi implementado ou medido; bandas e tempos são hipóteses.

---

## 1. Princípio

O aplicativo **não fala com o rádio**. O rádio entrega uma **rede IP local**; o MIGDOL roda sobre ela com um **nó tático de campo** que faz o papel do serviço de entrega. A criptografia ponta a ponta (MLS) não muda: o rádio só transporta cifrado.

```
[Celular] --Wi-Fi/Ethernet--> [Rádio tático IP] ~~~ rádio ~~~ [Rádio] --> [Nó de campo]
                                                                            |
                                                         ao voltar: sincroniza com a instância central
```

---

## 2. Meios de rádio

| Meio | Banda (hipótese) | Modo do MIGDOL | Situação |
|---|---|---|---|
| Rádio tático IP em malha (MANET) | Mbps | **Completo de campo:** texto, foto, anexos pequenos | Alvo principal |
| Wi-Fi do kit de campo | Mbps, alcance curto | Completo de campo | Base avançada |
| VHF/UHF de banda estreita com IP | Poucos kbps | **Banda estreita:** só texto curto, comprimido, quadros de tamanho fixo | Segunda etapa |
| HF | Muito baixa | Fora do escopo | Meios da Força |
| Satélite sob controle do Estado (ex.: SGDC) **[conferir]** | Varia | Enlace de volta à central | A avaliar |
| Satélite comercial estrangeiro, LoRa, malhas amadoras | — | **Não usar** (regra 10; segurança insuficiente) | Vetado |

**Primeiro passo:** levantar **quais rádios** a polícia e o Exército já usam, se entregam IP, a banda real e a cifra do enlace (D-R25).

---

## 3. Componentes

### 3.1 Nó tático de campo
- Servidor portátil endurecido, com o **mesmo binário** do servidor central em perfil de campo: gateway, entrega e diretório local da missão. Sem nó de IA, sem custódia aberta, sem painel de gestão completo.
- Disco cifrado, chave liberada por PIN do responsável e por hardware; **apagamento próprio** (escada do doc 07 seção 7.1) e botão físico de destruição de chave.
- **Nenhuma chave de custódia** (ADR-0003) nem cartão de comandante no campo.
- Retenção curta: guarda só o necessário até sincronizar; depois do sincronismo, apaga.
- Credencial de campo emitida pela AC da instituição, válida só durante a missão.

### 3.2 Aparelho
- Liga ao rádio por **Wi-Fi ou Ethernet**, com o rádio na **lista permitida do MDM**. **Não por cabo USB de dados**: o cabo dispara a escada de resposta (doc 07 seção 7.1).
- Perfil de campo obrigatório: câmera, PIN de coação, escada de resposta, sem IA.

### 3.3 Grupo de missão
- **Montado na base**, antes de sair: membros, salas e chaves definidos com conexão à central.
- **Durante a missão, só o nó de campo altera o grupo** (entrada e saída de membros). Os aparelhos só enviam mensagens. Isso evita bifurcação de época MLS quando a rede se parte.
- Remoção de membro em campo (aparelho perdido, capturado) é feita pelo nó de campo e repetida na central ao voltar.

---

## 4. Lease de campo (muda a ADR-05)
- Em campo, o **nó de campo emite o `segredo_servidor`** (doc 02 seção 4.3) com validade de missão (proposta: 12 h renováveis enquanto houver contato com o nó; teto de 72 h sem contato).
- Aparelho fora do alcance do nó além do teto: cache ilegível, como hoje.
- Captura do nó de campo: os aparelhos param de renovar; a central revoga a credencial de campo ao saber.

---

## 5. Sincronização com a central
- Ao reconectar (base, satélite autorizado ou rede da instituição), o nó de campo:
  1. envia as mensagens cifradas e as **cópias de custódia** já feitas nos aparelhos (ADR-0003);
  2. envia os eventos e a trilha de auditoria da missão;
  3. entrega as mudanças de grupo à central, que as confirma;
  4. apaga o que já foi sincronizado.
- **Conflito de grupo** (central e campo alteraram o mesmo grupo): a central prevalece e o grupo de campo é reconstruído; a regra precisa de teste (F-R3).

---

## 6. Modo de banda estreita
- Só texto, mensagens curtas, compressão antes de cifrar **[avaliar: compressão antes de cifrar vaza tamanho; usar quadros de tamanho fixo]**.
- Quadros de tamanho fixo (padding) para não revelar o conteúdo pelo tamanho.
- Sem anexos, sem recibos de visualização em tempo real (enviados depois do sincronismo), sem mudanças de grupo.
- Fila com prioridade: alerta e coação primeiro.

---

## 7. Segurança própria do rádio

| Ameaça | Quem protege | Observação |
|---|---|---|
| **Radiogoniometria** (localizar o emissor) | Rádio: salto de frequência, baixa probabilidade de detecção, potência mínima | Norma de uso: transmitir só o necessário |
| **Interferência** | Rádio e procedimento | App tolera perda, atraso e reenvio |
| **Análise de tráfego** (horário e volume) | App: quadros fixos, tráfego de cobertura opcional | Custa banda e bateria (D-R20) |
| **Interceptação do enlace** | Duas camadas: MLS do MIGDOL + cifra do enlace do rádio | Se a cifra do rádio for algoritmo de Estado, ela **pode ser a segunda camada** do doc 08 (D-R18) **[conferir]** |
| **Nó de campo capturado** | Nó: disco cifrado, apagamento, sem custódia | Revogação na central |
| **Rádio capturado** | Rádio: zeroização de chaves da Força | Fora do MIGDOL |

---

## 8. Riscos novos

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R89 | Nó de campo capturado com mensagens da missão | Médio | Alto | Disco cifrado, apagamento, retenção curta, sem custódia, revogação | Médio |
| R90 | Bifurcação de época MLS com rede partida | Médio | Médio | Só o nó de campo altera o grupo; reconciliação com a central prevalecendo | Baixo |
| R91 | Emissão de rádio revela posição da equipe | Alto | Alto | Rádio com baixa probabilidade de detecção; norma de transmissão mínima | Médio |
| R92 | Lease de campo prolongado aumenta a janela do aparelho perdido | Médio | Alto | Teto de 72 h; perfil de campo com controles obrigatórios | Médio |
| R93 | Compressão e tamanho das mensagens vazam informação em banda estreita | Médio | Médio | Quadros de tamanho fixo | Baixo |

---

## 9. Decisões

| ID | Decisão | Quem decide |
|---|---|---|
| D-R22 | **Decidido: operação em campo por rádio (opção B), Fase 8, para polícia e Exército** | Responsável pelo produto |
| D-R25 | Quais rádios e redes (MANET, VHF/UHF, satélite do Estado) e se entregam IP | Comunicações de cada instituição |
| D-R26 | Teto do lease de campo (proposta: 72 h) | Segurança + chefia |
| D-R27 | Cifra do rádio como segunda camada (D-R18) | Segurança, com o órgão que homologa |

---

## 10. Requisitos

| ID | Requisito | Prioridade |
|---|---|---|
| RF78 | Nó tático de campo com entrega local, diretório da missão e sincronização com a central | Fase 8 |
| RF79 | Grupo de missão montado na base; só o nó de campo altera o grupo em campo | Fase 8 |
| RF80 | Lease emitido pelo nó de campo, com teto sem contato | Fase 8 |
| RF81 | Modo de banda estreita: texto curto, quadros de tamanho fixo, fila com prioridade | Fase 8, segunda etapa |
| RNF-R1 | Ligação aparelho–rádio só por Wi-Fi ou Ethernet de rádio permitido pelo MDM | Fase 8 |
| RNF-R2 | Nó de campo sem chave de custódia e com apagamento próprio | Fase 8 |

---

## 11. Testes (F-R)

| # | Cenário | Esperado |
|---|---|---|
| F-R1 | Missão de 24 h em MANET com 10 aparelhos, texto e foto | Entrega com perda e atraso tolerados; medir banda e bateria |
| F-R2 | Aparelho fora do alcance do nó além do teto | Cache ilegível |
| F-R3 | Central e campo alteram o mesmo grupo | Reconciliação com a central prevalecendo; nenhum membro indevido |
| F-R4 | Nó de campo capturado e desligado | Disco ilegível; revogação na central; aparelhos param de renovar |
| F-R5 | Captura de tráfego de rádio | Só cifrado; tamanhos fixos em banda estreita |
| F-R6 | Celular ligado ao rádio por cabo USB de dados | Escada de resposta dispara |
| F-R7 | Sincronização ao voltar | Mensagens, custódia e auditoria na central; nó apaga o sincronizado |

---

## 12. Fase 8 — Operação em campo por rádio (depois da v1)

| Entrega | Critério de aceite |
|---|---|
| ADR-0004 aprovada | Lease de campo e nó de campo decididos |
| Levantamento dos rádios (D-R25) | Banda, IP e cifra de enlace medidos |
| Nó de campo em MANET (etapa 1) | F-R1 a F-R4, F-R6 e F-R7 aprovados |
| Banda estreita (etapa 2) | F-R5 aprovado; modo de banda estreita medido |
| Exercício de campo com a polícia e com o Exército | Relatórios anexados |
