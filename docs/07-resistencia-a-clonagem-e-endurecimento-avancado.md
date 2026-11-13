# 07 — Resistência a Clonagem, Aparelho e Local Autorizados, Vigilância por Câmera e Endurecimento Avançado

**Projeto:** MIGDOL
**Versão do documento:** 0.1 — rascunho para validação
**Público:** área de segurança, jurídico, encarregado de dados, equipe de desenvolvimento

> Este documento fecha as lacunas listadas na análise posterior ao documento 06. Os mecanismos que observam o usuário (câmera, local) entram no RIPD e seguem o teste do documento 06 seção 11 como **registro da decisão da instituição**, não como barreira. Nada aqui foi implementado ou medido: as taxas de acerto citadas são hipóteses a testar na Fase 0 e na Fase 2.

---

## 1. Princípios deste documento

1. **Várias camadas independentes, nenhuma confiável sozinha.** Câmera, local e série reduzem risco; nenhuma impede um atacante dedicado.
2. **Falha segura.** Sinal ausente, câmera coberta, localização indisponível ou atestação vencida significam **ocultar o conteúdo**, não "assumir que está tudo bem".
3. **Coletar o necessário e proteger o que se coleta.** A instituição assume a responsabilidade pela privacidade dos usuários (doc 06 seção 0). O limite deste documento é **segurança operacional**: imagens e posições de agentes são alvo de alto valor, então ficam cifradas, com acesso restrito e prazo curto. Por padrão, o que sobe do aparelho é um evento; imagem só sobe em alerta de câmera apontada (6.3).
4. **Alerta abre apuração; não pune.** Vale a regra RN14.
5. **O que não dá para fazer é declarado.** Câmera não detecta transcrição à mão nem memorização, e não detecta coação; detecta **presença** e **objetos**.

---

## 2. Antes de instalar: aparelho limpo e triagem de spyware

### 2.1 O que a triagem pode e não pode fazer
- **Pode:** encontrar spyware **conhecido** (por indicadores públicos), apps e configurações suspeitas, certificados e administradores indevidos, estado de boot e nível de patch fora do padrão.
- **Não pode:** provar que o aparelho está limpo. Spyware estatal sofisticado evita detecção por indicadores. **Ausência de achado não é atestado de limpeza.**

Por isso a base da segurança é **não confiar em aparelho de origem desconhecida**, e a triagem é um complemento.

### 2.2 Regra principal: aparelhos da instituição, preparados em estação controlada

| Etapa | Descrição |
|---|---|
| **Aquisição** | Compra pela instituição, lacre de fábrica conferido, número de série registrado no inventário e no MDM |
| **Estação de preparação** | Local físico controlado, rede isolada, equipe de duas pessoas, registro de quem preparou |
| **Estado conhecido** | Restauração de fábrica; firmware e sistema conferidos contra imagem oficial; bootloader travado; versão e patch no mínimo exigido |
| **Verificação de hardware** | Atestação de chave com boot verificado e estado "verde"; registro dos identificadores atestados (série, modelo) |
| **Provisionamento** | Modo de propriedade total pelo MDM (aparelho da organização), sem conta pessoal, sem serviços de terceiros desnecessários |
| **Triagem** | Ferramentas de verificação de indicadores executadas na estação (ver 2.3); relatório assinado anexado ao cadastro do aparelho |
| **Entrega** | Lacre, entrega presencial, cadastro por QR (RN05) |

**Aparelho pessoal ou de origem desconhecida não é cadastrado.** Quem quiser usar aparelho próprio precisa passar pela mesma preparação, o que na prática o torna da instituição.

### 2.3 Ferramentas de triagem
- **Open source de verificação de indicadores** (como o MVT, da Amnesty, e coletores forenses associados): funcionam por conexão de depuração na estação, o que é aceitável **na estação**, com a depuração desligada de novo na etapa final.
- **Defesa de ameaças móveis comerciais:** só se rodarem **dentro da instituição e sem telemetria para fora** (regra 10 do `CLAUDE.md`). Se o fornecedor exigir nuvem, **não adotar**.
- Lista de indicadores atualizada por processo interno, com origem registrada e validação antes de uso.
- Resultado: **aprovado**, **reprovado** (reflash e nova triagem; se repetir, aparelho descartado) ou **inconclusivo** (tratado como reprovado).

### 2.4 Triagens recorrentes
| Gatilho | Ação |
|---|---|
| A cada 90 dias (proposta) | Triagem leve de configuração e indicadores na estação |
| Aparelho perdido e recuperado, apreendido, deixado com terceiros, enviado a conserto | Reflash completo + triagem; até lá, suspenso |
| Viagem a local de risco (se a unidade definir) | Aparelho dedicado ou reflash ao retornar |
| Alerta de integridade repetido | Triagem profunda |
| Achado positivo | **SEV1**: considerar expostas todas as salas em que o aparelho esteve (playbook P2 + P5 adaptado) |

### 2.5 Verificações contínuas (sem depender de ferramenta externa)
O app e o MDM já cobrem apps instalados, acessibilidade, depuração, root e sobreposição (doc 02 seção 8.2). Acrescentar:
- **Autoridades certificadoras instaladas pelo usuário** (permitem interceptar tráfego): bloquear por política e detectar.
- **Perfis de VPN ou proxy** não previstos.
- **Administradores de dispositivo** e **serviços de notificação** fora da lista.
- **Alteração de data/hora fora do padrão**, **uso de depuração do Android** (já), **gravação de áudio ou vídeo por outro app durante o uso**.

---

## 3. Aparelho autorizado: série, atestação de identificadores e inventário

### 3.1 O que se usa
Aparelho de propriedade da instituição, com **atestação de identificadores** feita pelo hardware (número de série e modelo assinados na cadeia de atestação), que é muito mais difícil de falsificar que ler o número de série no software.

Detalhes e limites (validar na Fase 0):
- A leitura direta do número de série por app comum é restrita em versões modernas do Android. Em aparelho da organização, a atestação de identificadores pelo hardware e o identificador de inscrição do MDM são os caminhos corretos.
- Suporte à atestação de identificadores varia por fabricante e modelo. **Só entram na lista de modelos homologados os que suportam.**

### 3.2 Regra
O servidor só aceita aparelho se: (a) a cadeia de atestação é válida, (b) o **número de série atestado** está no inventário da instituição, (c) o modelo e o patch atendem ao mínimo, (d) o estado do boot é verificado. Série divergente do cadastro = recusa e incidente.

### 3.3 Inventário
- Fonte da verdade: inventário da instituição, cruzado com MDM e diretório **a cada dia** (R28).
- Aparelho no MDM e fora do diretório, ou o inverso, gera alerta.

---

## 4. Local autorizado: geocerca em camadas

### 4.1 O problema
Coordenadas de GPS e rede são **falsificáveis** em aparelho comprometido; além disso, **localização de agente é dado operacionalmente sensível**: quem a enxerga sabe onde a equipe está.

### 4.2 Desenho em camadas

| Camada | Como | Força | Observação |
|---|---|---|---|
| **1. Rede de origem** | O servidor só aceita a sala sensível de conexões vindas de redes autorizadas (APN privada, Wi-Fi corporativo com certificado, faixa de IP) | Alta | Atesta *onde a conexão entra*, não coordenadas |
| **2. Marcadores de proximidade** | Sinais de curto alcance no local (balizas Bluetooth com identificadores rotativos e desafio-resposta) | Média-alta em ambiente interno | Prova presença na sala, não só na região; exige instalar e manter balizas |
| **3. Verificação no aparelho** | O aparelho confere GPS, Wi-Fi e células dentro da zona e envia **apenas** um resultado atestado: "dentro da zona Z" | Baixa-média | Localização simulada bloqueada por política; sinal enganável em aparelho comprometido |
| **4. Atestação e integridade** | Aparelho íntegro (doc 02 8.2) | Base | Sem ela, as camadas acima valem pouco |

Recomendação: **camada 1 como obrigatória** para salas em perfil "local autorizado"; **camada 2 como opcional** para salas de análise; **camada 3 apenas complementar**.

### 4.3 Perfis de local por sala
| Perfil | Regra |
|---|---|
| **Livre** | Sem restrição de local (campo) |
| **Rede autorizada** | Só conecta por redes da instituição |
| **Local autorizado** | Rede autorizada + proximidade (camada 2); ao sair da zona, o conteúdo some e o cache da sala é fechado |
| **Compartimento** | Local autorizado + janela de leitura + sem exportação |

O administrador escolhe o perfil ao criar a sala ou canal. **Operações de campo usam perfil Livre**; o restante da operação (análise, relatórios) pode usar perfis mais restritos. Esse equilíbrio é decisão de negócio (D-R7).

### 4.4 Privacidade
- Coordenadas **não** são guardadas no servidor nem na trilha; só eventos "entrada/saída de zona" e "acesso negado por local".
- Registro de evento com zona, não com posição.
- Entra no inventário do documento 06 como dado pessoal de usuário, com retenção curta.

---

## 5. Resistência a clonagem

O desenho já impede copiar dados de um aparelho para outro (chaves em hardware, cache com chave dividida, backup desativado). O que falta é **detectar e conter** duplicação de identidade.

| Controle | Descrição |
|---|---|
| **Uma sessão por aparelho** | Nova conexão da mesma identidade **substitui** a anterior se trouxer contador de sessão assinado maior (troca de rede, rádio instável). SEV1 só com uso **simultâneo** comprovado: duas sessões ativas com contadores conflitantes (D18) |
| **Detecção de bifurcação de época** | Um commit por época no servidor; toda mensagem leva época e hash do transcript confirmado nos dados autenticados em claro; clientes e nó de IA comparam **sem decifrar** e alertam com mensagem da mesma época que diverge ou não decifra, ou com commit divergente (ADR-0005 item 2, D16, 01c-C02). Sinais classificados: SEV1 só com **bifurcação provada** (commit conflitante de administrador verificado ou importação de cadeia); pacotes que não decifram, em qualquer quantidade, são anomalia não autenticada, sem SEV1 (01g-G13); commit retido aparece pela confirmação de época **[proposta]** (01e-E02, E03). Partição total só se detecta na reunificação |
| **Chave usada em cada conexão** | A chave de identidade assina o handshake mTLS **de toda conexão**, não só do cadastro |
| **Reatestação periódica** | A cada sessão e a cada N horas, com desafio novo do servidor |
| **Chaves com autenticação por uso** | Chave de administrador exige biometria a cada uso; chave de identidade exige aparelho desbloqueado |
| **Invalidação por mudança de biometria** | Nova digital ou rosto cadastrado invalida as chaves; reativação pelo gestor |
| **Credencial com vínculo de hardware** | A credencial MLS do membro inclui o resumo da atestação; outros clientes recusam credencial sem hardware. Exceção: o nó de IA usa **credencial de serviço** emitida em cerimônia, com chave em TPM ou HSM, marcada como serviço (D19) |
| **Atualização periódica de chaves** | Commit de atualização MLS a cada 24 h (proposta) para limitar janela de comprometimento, feito pelo administrador da sala, um commit por época (ADR-0005). Em campo, serializado pelo nó e fora da banda estreita (ADR-0004, D15) |
| **Detecção de contêineres e clonadores** | Execução fora do perfil de trabalho esperado, UID e caminho de dados anômalos; sem múltiplos usuários no aparelho |
| **Ferramentas de migração do fabricante** | Bloqueadas por política; teste explícito de que nada do app migra |
| **Alerta de uso anormal** | Padrão de horário, volume ou local incompatível com o histórico do aparelho |

**Limite:** aparelho com root e desvio do Keystore consegue **usar** a chave sem exportá-la. A defesa é atestação do estado do boot a cada sessão, autenticação por uso e detecção de anomalia.

---

## 6. Vigilância por câmera durante o uso

### 6.1 O que é viável
A câmera frontal, **somente com o app em primeiro plano e conteúdo na tela**, processada **no aparelho**, pode detectar:

| Detecção | Viabilidade | Observação |
|---|---|---|
| **Usuário ausente** | Boa | Nenhum rosto no quadro por N segundos → oculta o conteúdo |
| **Mais de uma pessoa olhando** | Boa | Contagem de rostos, sem identificar ninguém |
| **Câmera ou celular apontado para a tela** | **Média, a medir** | Detecção de objeto (aparelho com lente voltada à tela) e de postura (braço levantado). Funciona melhor com o aparelho visível de frente; falha se a foto vem de trás, de ângulo fechado, de longe, com câmera pequena ou oculta |
| **Câmera coberta ou bloqueada** | Boa | Quadro escuro ou obstruído → tratar como estado inseguro |
| **Pessoa diferente do dono** | Possível, mas **sensível** | Exige reconhecimento facial (dado biométrico, doc 06). **Adotado só no desbloqueio** (GD-18; seção 7.2), não durante a leitura |
| **Coação física** | **Não detectável de forma confiável** | A câmera vê rostos e objetos, não intenção. Ver seção 7 |
| **Transcrição à mão, memorização** | **Não detectável** | Papel e caneta podem estar fora do quadro; ler e decorar não deixa sinal |

### 6.2 O que acontece ao detectar
| Evento | Ação padrão |
|---|---|
| Usuário ausente | Oculta; pede desbloqueio |
| Terceiro olhando | Oculta; aviso discreto; evento |
| Câmera ou celular apontado | Oculta; evento **de alta prioridade**; aviso ao usuário e ao Auditor; mensagem "ambiente não autorizado para leitura" |
| Câmera coberta ou indisponível | Oculta; evento |
| Falhas repetidas | Aparelho entra em revisão |

Todas as ações são **ocultar e registrar**. Nenhuma acusa nem pune sozinha (RN14).

### 6.3 Imagem, evidência e proteção do que é coletado
- **Fluxo contínuo:** nenhum quadro é gravado. O modelo roda em memória; só o resultado (evento) sobe. Gravar vídeo contínuo geraria um acervo enorme de alto valor para o atacante e não melhora a detecção.
- **Evidência de imagem em alerta (decisão D-R6: adotada).** No momento de um alerta de "câmera apontada" ou "terceiro olhando", o aparelho captura **uma ou poucas imagens** e as envia:
  - cifradas **para a chave de quórum (ADR-0003: mínimo 3 comandantes)** (nem o servidor, nem o administrador da operação, nem o Auditor sozinho leem);
  - com hora, aparelho e resultado do detector;
  - retenção definida pela instituição (proposta: 90 dias, ou até o fim do procedimento aberto);
  - consulta somente pelo rito do documento 05 seção 3.4, com registro.
- A decisão de adotar vale como decisão **da instituição**; o jurídico ainda confere a base legal e o texto de ciência (doc 06 seção 7), porque isso protege a validade da prova.
- O indicador de câmera do sistema acende; o usuário **sabe** que a câmera está ativa (RN15).
- **Cuidados vindos da verificação jurídica** (docs/anexos/validacao-juridica-preliminar.md seção 3.5): a câmera só funciona com conteúdo sensível na tela; a imagem de evidência é **recortada** à região relevante e de baixa resolução; nada é capturado fora de alerta; sala sensível não deve ser lida em ambiente privado (perfil de local); a câmera não é usada em tela de configuração, chamada ou fora do app. Não encontrei decisão sobre câmera de aparelho funcional: **parecer do jurídico** antes da Fase 2.
- O indicador de câmera do sistema acende; o usuário **sabe** que a câmera está ativa (transparência, RN15).
- **Fora da tela de conteúdo, a câmera não é usada**, exceto na tela de desbloqueio do app, para a verificação facial de quem acessa (seção 7.2; GD-18). Esse uso é novo em relação às cautelas da validação jurídica acima e depende de parecer **[conferir]**.
- Custo: bateria e aquecimento. Medir na Fase 2 (taxa de quadros baixa, ex.: 1–2 por segundo, e modelo leve).

### 6.4 Reconhecimento facial: adotado no desbloqueio (GD-18; antes D-R8)
O responsável adotou a verificação facial de quem acessa (GD-18; seção 7.2). Os cuidados abaixo **não são de privacidade**, e sim de segurança e de eficácia, e continuam valendo:
- **Modelos próprios têm erro.** Falso positivo trava o agente; falso negativo deixa passar. Medir antes (A42).
- **Modelo facial é alvo de falsificação** (foto, vídeo, máscara). Exige detecção de vivacidade e testes.
- **O gabarito facial é dado de altíssimo valor** para o adversário. Se adotado: gabarito **só no aparelho**, cifrado em hardware, nunca no servidor.
- **A biometria do sistema** (digital ou rosto de classe forte) já autentica o dono do aparelho com mais maturidade; o reconhecimento próprio só se justifica para **verificar continuamente** quem está lendo.
A recomendação anterior (v1 sem reconhecimento próprio) foi substituída pela GD-18. O recurso só entra no gatilho de destruição depois de medido (A42, A51, A76–A79), nas condições da GD-24 (seção 7.1a) e dentro dos limites da D-R40; acima deles, fica como sinal auxiliar. A ANPD abriu tomada de subsídios sobre biometria em 2025 e pretende concluir regras em 2026; o parecer jurídico acompanha **[conferir]**.

### 6.5 Honestidade sobre eficácia
- A detecção de câmera apontada **reduz** o vazamento por foto, mas não o elimina: um atacante informado fotografa de ângulo cego, usa câmera oculta, ou cobre e descobre que o app oculta o conteúdo e tenta outro modo.
- Valor real: **dissuasão**, **redução de oportunidade** e **evidência de intenção** quando o alerta ocorre.
- Taxa de acerto e de falso positivo são medidas por cenário (A38–A42). Meta de partida: ≥ 80% de detecção em foto frontal a até 1 m e ≤ 2% de falso positivo por hora de uso; **hipótese a validar**.

---

## 7. Modo de coação

Não há como **detectar** coação por sensor com confiança. O que se faz é dar ao usuário **meios discretos de acionar** uma resposta, e usar a presença de terceiro como **sinal auxiliar**.

| Recurso | Funcionamento |
|---|---|
| **PIN de coação** | Segundo PIN que abre uma conta **isca** (conversas e salas genéricas, sem dado real), envia alerta silencioso e marca a sessão. Nunca revela que é isca. **Por perfil (GD-18; seção 7.1a):** no uso urbano, também suspende (nível 3); no perfil missão de campo, aciona o nível 4 por baixo da isca **só depois de medido o acionamento acidental** (GD-24); até lá, nível 3, como no urbano |
| **Gesto ou botão de pânico** | Ação discreta que oculta tudo, envia alerta pela sala de resposta (seção 7.4) e aciona os níveis 2 e 3 da escada (seção 7.1). O **nível 4** (apagar) por pânico só se o jurídico aprovar por escrito (D-R14, D23) |
| **Trava por tentativas** | Depois de N erros de desbloqueio, aciona o **nível 3** (seção 7.1): trava sem apagar. Apagar por tentativas (nível 4) só se o jurídico aprovar por escrito (D-R14, D23); no perfil missão de campo, D-R40. Rosto de outra pessoa com PIN errado segue a seção 7.1a |
| **Presença de terceiro** | Contagem de rostos como sinal auxiliar de risco; aciona oculte e alerta, não conclusões |
| **Prazo de check-in** (opcional) | Para operações de risco: se o aparelho não renovar a presença no prazo definido, o servidor trata como possível apreensão (playbook P2) |

Cuidados:
- **PIN de coação e pânico exigem treinamento e decisão de unidade.** O erro mais provável é o usuário acionar sem querer: confirmação curta onde não comprometa o disfarce, e procedimento de retorno.
- A conta isca precisa parecer real o suficiente; não é pelo código que ela se sustenta, é pelo procedimento.
- O alerta precisa chegar a alguém que **age**: define-se na Fase 0 quem recebe e em quanto tempo responde (doc 05 seção 9).
- Isso **não resolve** coação física; reduz o dano e dá tempo à equipe.

### 7.1 Escada de resposta do aparelho (revisão: o que apagar, quando e por quê)

Princípio: **reversível antes de irreversível**. Destruição só quando há adversário ativo confirmado. No **perfil missão de campo** (GD-18), rosto de outra pessoa (duas capturas) com PIN errado, ou PIN de coação, contam como adversário ativo, nas condições de ativação da GD-24 (seção 7.1a). Nenhum nível aplica consequência contra a **pessoa** (RN14, regra 11 do projeto): o aparelho se protege; quem decide sobre o agente é um humano.

| Nível | Resposta | Gatilhos | Reversível? |
|---|---|---|---|
| **1. Travar** | Oculta o conteúdo, zera da memória o `segredo_servidor` (doc 02 seção 4.3) e exige nova sessão autenticada | Cabo de dados conectado com conteúdo na tela, ADB, projeção, terceiro olhando | Sim |
| **2. Isca e alerta silencioso** | Abre conta isca; alerta ao Auditor e à segurança; imagem de evidência cifrada para a chave de quórum (seção 6.3) | PIN de coação; PIN correto com rosto de outra pessoa ou falsificação (GD-18); pânico sem apagamento | Sim |
| **3. Trava criptográfica** | O servidor **para de entregar** o segredo de lease (ADR-0006): o cache fica ilegível. O aparelho é **suspenso** automaticamente; **revogar é ato humano** registrado (D25). O administrador da sala, ou o substituto, remove o aparelho dos grupos em até 1 h **[proposta]**, com alerta se não remover (D17); a remoção se desfaz por reinclusão depois da perícia. **Nada é destruído** | Sinais de ferramenta forense via cabo; apreensão declarada pelo gestor; check-in vencido; spyware suspeito; N erros de desbloqueio; pânico (até D-R14). **Perfil urbano (GD-18):** rosto de outra pessoa com PIN errado; PIN de coação (junto com o nível 2) | Sim |
| **4. Apagamento local** | Zera chaves e cache **deste aparelho**. Ordem dos atos (01g-G12b): enfileira o alerta, envia o alerta e o **aviso assinado de autoexclusão** se houver enlace, e só então zera; sem enlace, o alerta se perde com o aparelho (limite declarado). O aviso de autoexclusão é uma proposta de remoção de si mesmo (suporte da biblioteca a verificar no spike). O administrador da sala, ou o substituto, faz o commit de remoção em até 1 h **[proposta]**, com alerta se não fizer (D17). Só volta por revinculação (seção 7.4) | Instrumentação ativa confirmada (hook, root ou falha de atestação com sessão aberta). N erros de desbloqueio e pânico com apagamento **só** com decisão escrita do jurídico (D-R14, D23). **Perfil missão de campo (GD-18), nas condições da GD-24 (seção 7.1a):** rosto de outra pessoa em duas capturas, com PIN errado; PIN de coação, por baixo da isca, depois de medido o acionamento acidental. Falsificação detectada (vivacidade) não aciona até ser medida à parte | Não, para este aparelho |

**O que o nível 4 não apaga:** cópia no servidor (cifrada), cópias nos outros membros, cópia de custódia (sala sob custódia), trilha de auditoria. Só o aparelho perde o acesso. O evento é **registrado e assinado**, com motivo e sinal que o disparou, e o agente é informado na próxima autenticação legítima (RN15), não no momento do ataque.

**Lease vencido não apaga:** só deixa o cache ilegível (ADR-0006, D23).

**Revogar não destrói (GD-13; ADR-0006 item 8; 01c-C01):** revogar a credencial e destruir o segredo de lease são **atos separados**.
- **Revogar** (certificado, sessão, lease): imediato, uma pessoa, registrado. Em **toda** revogação, qualquer que seja o motivo declarado, o servidor **sela** o segredo de lease vigente para a chave de quórum do segredo de lease e apaga a cópia em claro. O cache fica ilegível (o servidor não consegue mais entregar o segredo), mas nada é destruído. Aparelho dado como perdido que na verdade foi apreendido não perde o cache por engano.
- **Destruir o segredo selado:** duas autorizações de pessoas distintas e registro, como o apagamento remoto; com indício de apreensão por autoridade, só depois da decisão D-R14.
- **Prazo:** o da apuração do evento, fixado na revogação (valor padrão: D-R35). Abrir exige quórum presencial com base registrada (regra 14), e o segredo sozinho não abre nada sem o aparelho e o desbloqueio do agente.
- **Painel:** "aparelho revogado: cache ilegível, segredo guardado sob quórum até DD/MM/AAAA", nunca "destruído" (D27). Aparelho revogado que reconecta recebe primeiro as ordens pendentes e nunca o segredo.

**Sobre a ideia de "apagar tudo ao detectar captura por cabo ou spyware":**
- **Spyware ativo confirmado** (instrumentação, hook, root com sessão aberta): níveis 3 e 4 se justificam. O adversário é a ameaça e o cache já pode estar exposto.
- **Cabo de dados:** níveis 1 e 3, **não 4**. A extração forense policial e o ataque de um adversário usam o mesmo cabo e **o aparelho não distingue um do outro**. Com o `segredo_servidor` zerado, a extração já não lê o cache. Apagar nesse caso, depois de apreensão legal, pode ser lido como destruição de prova (R63, R66; D-R14) e perde o registro institucional. Carregar em computador sem dados não deve disparar nada (política de portas, seção 9).
- **Falso positivo apaga o agente legítimo:** por isso o nível 4 exige sinal forte ou várias tentativas, nunca um sinal isolado. No perfil missão, os dois sinais são o rosto de outra pessoa, em duas capturas independentes, e o PIN errado; o rosto sozinho nunca basta.

### 7.1a Escada por perfil (GD-18)

O `instituicao.yaml` define o perfil da instituição. O aparelho entra no **perfil missão de campo** só enquanto é membro de um grupo de missão em andamento (da saída ao retorno, registrados pelo administrador da missão no estado autenticado do grupo; ADR-0004). Fora disso, vale o perfil comum. O servidor sozinho não liga o perfil missão (A82).

| Situação no desbloqueio (seção 7.2) | Perfil policial urbano (e fora de missão) | Perfil missão de campo (Exército, captor adversário) |
|---|---|---|
| Rosto do dono e PIN correto | Abre | Abre |
| Rosto do dono, ou inconclusivo, e PIN errado | Conta tentativa; N erros → nível 3 | Conta tentativa; N erros → nível 3 (nível 4 por tentativas: D-R40) |
| Rosto inconclusivo e PIN correto | Abre, com evento; nunca destrói | Idem |
| Rosto de outra pessoa (duas capturas), ou falsificação detectada, e PIN correto | Nível 2 (isca e alerta silencioso) e imagem de evidência | Idem |
| Rosto de outra pessoa em duas capturas independentes, e PIN errado | **Nível 3**: suspende; revogar é ato humano; nada se destrói até D-R14; imagem de evidência | **Nível 4** (apagamento local), com alerta e aviso de autoexclusão antes de zerar, se houver enlace; imagem de evidência. **Só depois das medições da GD-24**; até lá, como no perfil urbano |
| Falsificação detectada (falha de vivacidade), ou "outra pessoa" numa só captura, e PIN errado | Conta tentativa; N erros → nível 3; imagem de evidência | Idem: **não** aciona o nível 4 até a vivacidade ser medida à parte (GD-24) |
| PIN de coação, com qualquer rosto | Nível 2 (isca) e **nível 3** | Isca e **nível 4** por baixo dela, sem sinal na tela, **só depois de medido o acionamento acidental do PIN de coação** (GD-24); até lá, isca e nível 3 |

- A face **nunca decide sozinha**: nenhuma linha destrói só pelo rosto, e "inconclusivo" (sem rosto, rosto coberto, pouca luz, baixa qualidade) nunca conta como "outra pessoa".
- **Condições de ativação do nível 4 no perfil missão (GD-24; 01g-G11, G12).** Até cada medição ser feita e aprovada, o gatilho correspondente segue a coluna do perfil urbano (nível 3, sem destruição):
  - **"Outra pessoa" exige duas capturas** independentes, cada uma com rosto vivo, de boa qualidade, que não corresponde ao gabarito (seção 7.2). Uma captura só vale como inconclusiva para o nível 4.
  - **Falha de vivacidade** (falsificação detectada) **não conta para destruir** até ser medida à parte (A76); vale como inconclusiva para o nível 4: com PIN correto, nível 2; com PIN errado, conta tentativa (N erros → nível 3).
  - **Número de PINs errados** com "outra pessoa": fixado só **depois de medir** o falso negativo em condição de campo (A78, com muitas tentativas; o V09 mede só aceitação e tempo). Proposta: **um** PIN errado, se a medição permitir (D-R40).
  - **PIN de coação → nível 4:** só depois de medir o **acionamento acidental** do PIN de coação sob estresse (pressa, luva, escuro; V09, tarefa U05 do doc 11, com repetição depois de semanas sem uso). Até lá, no perfil missão, o PIN de coação abre a isca e aciona o nível 3. O resultado entra na D-R40 (R55).
  - **Isca no nível 4 (01g-G12c):** a conta isca fica fora do cache destruído, e o apagamento por baixo dela não muda tempo de resposta, aquecimento nem consumo perceptível (PI-4 do doc 11; A79).
- Quantos PINs errados com rosto de outra pessoa, nível 4 por tentativas no perfil missão e se a polícia liga o perfil missão nas missões de campo dela: D-R40, com os valores medidos acima.
- O nível 4 continua alcançando só o aparelho: cópia no servidor, nos outros membros, custódia e trilha ficam. A destruição no perfil missão é aceita pelo responsável para captor adversário (GD-18); no uso policial urbano vale a D-R14 (R63, R66).
- Nó de campo capturado: destruição imediata própria (ADR-0004 item 13).

### 7.2 Verificação facial de quem acessa (GD-18)

Há duas coisas diferentes:
- **Foto do intruso em alerta:** já adotada (seção 6.3). Cifrada para a chave de quórum, só com alerta.
- **Verificação facial de quem acessa:** **adotada pelo responsável (GD-18)**, que fecha a D-R8 nesse sentido. Falta o parecer jurídico sobre dado biométrico (LGPD art. 11; regras da ANPD em elaboração) **[conferir]**.

Como funciona:
- **Onde e quando:** na tela de desbloqueio do app, em primeiro plano, a cada abertura e após inatividade (RF02), com as capturas feitas **durante a digitação do PIN**, sem tela própria que dependa do PIN e igual com o PIN de coação (doc 11, fluxo 4.8; 01g-G15). Não é verificação contínua durante a leitura: a vigilância da seção 6 continua só contando rostos.
- **Modelo no aparelho:** de origem conhecida, com hash fixado (regra 8), sem rede e sem serviço de terceiro (regras 10 e 12). Detecção de **vivacidade** passiva no aparelho, contra foto, vídeo e máscara. Os quadros são processados **em memória**, pelo core, e nunca gravados nem enviados (regra 12; A81).
- **Gabarito:** gerado no cadastro presencial, com vivacidade; guardado **cifrado por chave não exportável do hardware seguro** (StrongBox/TEE; Secure Enclave no iPhone); nunca no servidor, em backup, em log ou em exportação (A80). Para comparar (01g-G22a): o gabarito é cifrado pelo core, com a biblioteca da pilha, sob uma chave de dados que a chave do hardware embrulha; o core pede ao hardware o desembrulho **dessa chave**, que transita só no quadro da chamada, num buffer direto alocado pelo core (doc 02 seção 4.1), e é zerada em seguida (regra 2), e decifra o gabarito **em Rust**. O gabarito em claro nunca passa pela JVM nem pelo Swift. **Dúvida registrada:** essa chave abre o gabarito guardado no próprio aparelho, sem outro segredo, porque a matriz precisa do resultado facial também com PIN errado; conferir na Fase 2 se isso cabe no "sozinho não abre nada" da regra 2 **[verificar]**. Zerado no nível 4.
- **Limite declarado:** a comparação roda no core, fora do hardware seguro. Fazê-la dentro do TEE exigiria aplicação confiável instalada no TEE, o que um app não faz no Android. Aparelho com root e instrumentação pode observar a comparação (pergunta R3-4 do `docs/fase0/01f-correcoes-rodada3.md`; verificar na Fase 2). Quem cobre o rosto cai em "inconclusivo" e entra só com o PIN correto, como se não houvesse reconhecimento; contra PIN obtido por coação, a defesa continua sendo o PIN de coação. **A função do reconhecimento é detectar, não barrar (01g-G11b):** contra adversário informado ele não reforça o acesso; serve contra o desavisado, e o custo de erro recai no dono, por isso o limite de falso negativo para o dono fica perto de zero, medido em condição de campo (pintura, capacete, ferimento, noite; D-R40).
- **Resultado:** "dono", "outra pessoa" (rosto vivo, de boa qualidade, que não corresponde ao gabarito, em **duas capturas independentes**; GD-24), "falsificação detectada" (falha de vivacidade, que fica fora do gatilho de destruição até ser medida à parte; GD-24) ou "inconclusivo" (inclui uma só captura de outra pessoa).
- **Perfil missão lido antes do desbloqueio (01g-G22b):** a marca de missão em andamento é gravada na saída e apagada no retorno, protegida por chave do hardware, fora do cache cifrado, para que o aparelho saiba qual coluna da seção 7.1a vale antes de abrir o cache. O servidor sozinho não a liga (A82).
- **A face nunca decide sozinha (GD-18):** destrói só com "outra pessoa" **e** PIN errado, ou com PIN de coação, só no perfil missão de campo e só depois das medições da GD-24 (seção 7.1a). Rosto diferente com PIN correto indica PIN roubado ou coação: a resposta é o **nível 2** (isca e alerta silencioso), não derrubar o usuário.
- **Imagem de evidência:** em "outra pessoa" ou falsificação, uma ou poucas imagens, cifradas para a chave de quórum de evidência (seção 6.3). É o alerta que a regra 12 admite.
- **Medição antes de ativar:** falso negativo (dono não reconhecido) e falso positivo ou falsificação aceita (impostor reconhecido) medidos em A42, A51 e A76–A79. Acima do limite (D-R40; doc 11), o recurso fica só como sinal auxiliar e sai do gatilho de destruição.
- **"Dropar tudo desse usuário" em tempo real** (revogar, apagar dados dele no servidor) **continua fora**: seria decisão automática contra pessoa a partir de IA (regra 11, RN14). A revogação é decisão humana do Auditor e do gestor. O nível 4 do perfil missão protege o **aparelho**; não atinge a pessoa nem o servidor.
- "Sem falar nada" vale **para o intruso**. Para o agente legítimo, o termo de ciência prevê o comportamento, e ele é informado depois.

**Regra 12 (GD-25; alterada na quarta rodada, `01h`):** a regra passou a dizer que a câmera só é usada com o app em primeiro plano, na tela de conteúdo e no desbloqueio, e que o gabarito facial fica só no aparelho, cifrado por chave de hardware, e nunca é enviado. Processar quadro em memória para reconhecer o rosto, sem gravar nem enviar, já cabia na letra. Continua pendente a frase de escopo da D-R30 (foto do termo de revinculação e câmera do app para foto, fora de alerta; pergunta 2 do `01c`).

### 7.3 Custódia e entrega de mensagens (D-R15, ADR-0003 revisada)

Proposta original do usuário: a corporação entrega mensagens armazenadas por no máximo 5 anos **somente com ação judicial**. Decisão do responsável (GD-1, GD-2, GD-6, GD-14), na forma da ADR-0003 revisada na etapa 1b e na segunda rodada:

| Ponto | Decisão |
|---|---|
| Onde há custódia | **Só em sala sob custódia.** Flag por instituição no `instituicao.yaml` assinado: Exército **ativa** (investigação militar); polícia **desligada por padrão**, pode ligar (corregedoria ou inquérito). Com a flag ligada, a custódia é o **padrão** de toda sala nova (GD-14). Marca na criação, assinada, visível aos membros, irreversível. Mudar a flag é cerimônia e **não retroage** |
| Sala comum | MLS puro. Na conclusão, sai dos aparelhos e é **expurgada**; não há cópia. Ordem externa recebe só o que existe (cifrado ilegível, metadados, auditoria): impossibilidade **por desenho** |
| Como é a cópia | No remetente, HPKE para a **chave pública** de custódia do período; a privada fica só no HSM offline. Membros não decifram (D01). Período só por tempo, como hipótese do spike (ADR-0003 seção 2; 01c-C06, 01e-E21e). Toda mensagem da sala leva cópia, sem exceção por tipo; apagar mensagem ou autodestruição não tiram a cópia (01c-C09) |
| Prazo | Até **5 anos a contar da conclusão** (ou do expurgo, se a sala foi expurgada sem conclusão; 01c-C10); depois, eliminação no HSM com recibo. Sala ativa além do teto (proposta: 12 meses) gera alerta. O número ainda deve vir da tabela de temporalidade e do jurídico (D-R5) **[conferir]** |
| Quem abre | **Quórum presencial** de no mínimo 3 comandantes, com cartões, no HSM offline. Ninguém sozinho; nenhum administrador do aplicativo |
| Com que base | **Ordem judicial** (é cumprida) ou **processo administrativo com validação jurídica** **[conferir]**. A **CI** registra e tramita o pedido **em todos os casos**, com número obrigatório; **não substitui** ordem judicial (01c-C23). Para autoridade externa, ordem judicial obrigatória |
| Saída | Pacote de evidência assinado (RF67), com registro na trilha |
| Corregedoria e controle interno | Se dependem de ordem judicial ou basta processo administrativo com validação jurídica (com CI): jurídico **[conferir]** |
| Sala comum com a flag ligada | **Decidido (GD-14, D-R29 opção B):** só com **motivo registrado** e **segunda assinatura** de pessoa distinta do criador (quem: D-R36). Vale também para sala de compartimento. **Salas de sistema** (sala de resposta, seção 7.4) são sempre comuns |

**Riscos:** o acervo das salas sob custódia é alvo de alto valor (R68). Comprometer o HSM offline ou o quórum expõe os períodos guardados por até 5 anos. Risco admitido pelo responsável pelo produto **só** em sala sob custódia (GD-1); aceite formal pendente (doc 04 seção 3.2; 01c-C18). Com a flag ligada a custódia é o padrão (GD-14), e o acervo tende a cobrir quase todas as salas. Contradiz o objetivo "sem histórico permanente" do documento 01 apenas nessas salas. Sala comum criada para fugir da custódia: R111.

**Isto é segurança jurídica correta?** A direção (base registrada, com ordem judicial obrigatória para terceiros; controle múltiplo presencial; prazo máximo; registro) **aumenta** a segurança jurídica. O apagamento silencioso e automático **em apreensão** a **diminui**. Nada disso é parecer jurídico: itens em **[conferir]** vão ao jurídico.

### 7.4 Perda, roubo, apagamento remoto e revinculação

Premissas confirmadas pelo responsável do produto (a segunda revista no GD-4):
1. As mensagens locais se apagam; o dado continua **cifrado no servidor**.
2. A revinculação vale só para operação **ativa**, com **aprovação dupla** (chefe da inteligência e superior da área), termo de ciência e aceite, e histórico enviado pelo aparelho de **outro membro** pelo MLS. Sem HSM online e sem aprovação remota por quórum.
3. A administração apaga o aparelho remotamente, **sem deixar rastros no aparelho**.
4. Aparelho perdido: apagar e rastrear, mesmo que tentem redefini-lo, em etapas (abaixo).

**Fluxo de perda ou roubo (administrador, em tela):**
- O administrador declara o **motivo**: perdido ou roubado; suspeita de comprometimento; ou **apreendido por autoridade**.
- Perdido, roubado e comprometido levam, na hora, à **trava** (nível 3: o servidor para de entregar o segredo de lease, e o cache fica ilegível) e ao modo rastreio. O **apagamento** exige **segunda autorização** de pessoa distinta (D23). Ela pode vir depois, porque a trava já protege o conteúdo.
- **Apreensão por autoridade** leva a suspender e revogar, sem apagar, até a decisão D-R14. Como em toda revogação, o segredo de lease vigente fica selado para a chave de quórum (seção 7.1; ADR-0006 item 8); com indício de apreensão, destruí-lo só depois de D-R14.
- **Ordem de apagamento:** formato binário próprio (`MT-WIPE-v1`, doc 02 seção 6.2), com **aparelho-alvo**, motivo, autor e **duas assinaturas**. O contador é por aparelho; o aparelho confere só assinaturas e contador, sem janela de tempo (D29). O app verifica a ordem no core antes de agir. Se o aparelho está offline, o servidor guarda a ordem e a entrega na próxima conexão. Apagar em lote também exige duas autorizações (R71).
- **Escopo:** aparelho totalmente gerenciado da instituição (Android Enterprise): apaga o aparelho inteiro pelo MDM, e não só o perfil de trabalho (doc 02 seção 8.2 hoje fala em perfil). O MDM só apaga a partir da ordem assinada, com duas pessoas no console do MDM (R54, R71, D24).
- **Recibo:** antes de zerar, o aparelho envia um recibo assinado **de intenção** ("vou apagar"). Ele não prova a execução. A execução se confirma pela ausência de sessão e de atestação depois (D24).
- **Apagar em etapas (D24):** primeiro chaves de sala, cache e identidade; a chave de rastreio, curta, fica só até o fim do prazo de rastreio; depois, o resto. "Apagar tudo" e "rastrear" deixam de ser incompatíveis.
- **"Sem rastros no aparelho"** significa apagamento criptográfico, sem log, aviso ou resíduo local. **O rastro existe no servidor**, assinado, com autor e motivo. É ele que defende a instituição de acusação de destruição de prova (R66) e cumpre a regra 13.
- **Aparelho offline:** a ordem não executa. A proteção é o lease criptográfico (ADR-0006): ao revogar, o servidor sela o segredo vigente para a chave de quórum e não consegue mais entregá-lo, e o cache fica ilegível mesmo sem a ordem. Nada é destruído sem duas autorizações (seção 7.1).

**Redefinição e rastreio, com limites honestos:**
- A política do MDM bloqueia o reset pelas Configurações. O reset pelo **modo de recuperação** e a **persistência de matrícula depois do reset** (proteção de redefinição, Zero-touch ou Knox) dependem do modelo e de recursos do fabricante ou do Google, que são serviços de terceiros (regra 10). Decisão **D-R16** e teste no spike do MDM na Fase 0.
- Mesmo depois de reset: o aparelho não tem mais a chave de identidade (destruída), nem cache, e **não consegue se cadastrar** sem QR presencial (R21). O ladrão fica com um aparelho vazio. O que se protege é o dado, não o hardware.
- **Rastreio por localização** só depois da declaração de perda, enquanto o app estiver presente e conectado, sem serviço de localização de terceiro. Base legal e proporcionalidade: jurídico (anexo seção 4, item 4).
  - **Sala de resposta (proposta, D10):** cada aparelho tem, desde o cadastro, uma sala MLS dormente com quem precisa agir numa perda ou coação (segurança de plantão e gestor; quem e quantos: pergunta 13 da revisão; pergunta 8 do 01c). Declarada a perda, a posição em tempo real vai **pelo MLS** para essa sala, que os membros leem como qualquer sala (regra 14). Retenção curta; expurgo ao encerrar a perda. O alerta de coação usa a mesma sala. É **sala de sistema**: sempre comum, também com a flag de custódia ligada (GD-14). Pendentes (01c-C03, C04): estado mínimo fora do lease para funcionar no nível 3, credencial de rastreio que sobrevive à primeira etapa do apagamento, membros por aparelho de função, batimento e operação em missão.
  - Posição **guardada** para apuração vai só cifrada para a chave de quórum de localização e se abre por quórum presencial (ADR-0003). Isso substitui o texto anterior "posição cifrada para Auditor e gestor", que contrariava a regra 14.
  - Fora da perda, continua valendo a seção 4.4 (sem coordenada no servidor).
- Aparelho desligado, sem rede ou em bolsa de bloqueio de sinal **não recebe a ordem e não é rastreado**. Declarar isso na norma de uso.

**Revinculação (reposição do aparelho), v1 (GD-4, GD-5):**
- Vale **só** para salas de operação **ativa** (doc 01 RN18). Sala concluída não é reaberta para o aparelho novo.
- **Sem HSM online e sem aprovação remota por quórum.** A revinculação **não decifra** nada da instituição; a custódia não é usada.
- **Pré-requisitos:** aparelho antigo revogado; aparelho novo preparado e triado (RF59), série no inventário (RF60); cadastro **presencial** com identificação do agente; Termo de Ciência e Autorização vigente (RN16); QR novo.
- **Aprovação dupla:** o **chefe da inteligência** e o **superior da área** assinam no app o pedido (hash de: agente, aparelho novo, aparelho revogado, salas e, por sala, os **intervalos de pertença** do agente, tirados do estado do grupo pelo administrador; 01c-C20). Duas pessoas distintas entre si, do agente e do gestor que fez o cadastro. As chaves de aprovador são registradas em cerimônia e conferidas no core; não vêm do diretório (D03, D20).
- **Termo de ciência e aceite da revinculação:** o app gera o termo; o proprietário assina **no próprio app** (assinatura na tela e assinatura com a chave do aparelho novo, no core) e tira **foto** com a câmera interna.
  - O termo, com a foto, é cifrado no aparelho para a **chave de quórum do termo** e guardado no MIGDOL, com prazo definido (D-R30). O hash vai à trilha de auditoria. Abre-se só por quórum presencial, para responder a uma CI futura (GD-5, regra 14).
  - O app **exporta** o termo para anexar no sistema da corporação, pelo fluxo de exportação autorizada (RN07).
  - **Dúvida registrada (D-R30):** a foto é dado biométrico sensível (LGPD art. 11) e a regra 13 exige dado de agente cifrado, com acesso por rito. Até o responsável decidir, o documento exportado leva o **hash da foto**, não a foto (opção mais restritiva). Também falta confirmar que a regra 12 trata só da câmera de vigilância durante a leitura, como já ocorre com a foto de anexo (RF13).
- **Inclusão:** o administrador da sala, ou o substituto, inclui o aparelho novo por commit que referencia a aprovação dupla. Os clientes conferem as duas assinaturas antes de aplicar (ADR-0005).
- **Histórico:** vem do **aparelho de outro membro** da sala, pelo MLS. O membro indicado pelo administrador envia ao aparelho novo, em mensagem cifrada só para ele (como as respostas individuais, doc 02 seção 9.2), as mensagens que tem em cache **dentro da retenção** e **dos intervalos em que o agente era membro**. Leva a assinatura original do remetente se a biblioteca permitir guardar o conteúdo autenticado; senão, vai marcado "reenviado por <codinome>" e assinado por quem reenviou **[verificar no spike]**.
- O aparelho novo recusa, do pacote, o que estiver fora dos intervalos de pertença do pedido assinado ou da retenção (A70).
- O que nenhum membro tem em cache (vencido, expurgado) não volta.
- **Reinclusão do mesmo agente depois da perícia** (nível 3 da seção 7.1): mesmo rito da revinculação, com aprovação dupla e histórico por `history_share` **[proposta; 01c-C20]**.
- **Registro:** pedido, aprovadores, hash do termo, inclusão e hash do pacote reenviado na trilha. Aviso ao agente pelos canais anteriores, ao Auditor e aos membros da sala.
- **Opções avaliadas para o histórico:** A (chave de recuperação sob custódia) foi **retirada da v1** com a revisão da ADR-0003 (D01–D05); **B** (outro membro reenvia) é a **adotada**, na forma acima; C (sem reenvio) vale para o que nenhum membro tiver.

---

## 8. Atestação sem Play Integrity (ADR proposta)

**Conflito:** o documento 02 prevê Play Integrity, que depende de serviço do Google e contraria a regra 10.

**Proposta (ADR-0002):**
1. **Atestação de chave do hardware**, validada no servidor da instituição contra as raízes de confiança dos fabricantes, mantidas internamente e atualizadas por processo próprio.
2. Verificar: estado do boot verificado, nível de patch, bloqueio do bootloader, chave em hardware seguro, identificadores (seção 3).
3. **Não usar Play Integrity** nem serviços do Google Play no aparelho.
4. Aparelho sem esses requisitos não entra na lista de modelos homologados.

Consequências: o servidor passa a **manter a lista de raízes e a política de aceite**; o spike de atestação da Fase 0 precisa cobrir esse processo. O arquivo `docs/adr/0002` deve ser criado antes do spike de atestação (etapa 5; D34). A estação do gestor confere também a atestação, para que o diretório sozinho não ative aparelho fantasma (D20; doc 02 seção 4.2).

---

## 9. Outro endurecimento necessário

| Tema | Ação |
|---|---|
| **Parsers de anexo** | Reencodar imagens e documentos no remetente e no destinatário; parsers em código seguro quanto à memória ou em processo isolado; fuzzing; limites de tamanho e formato |
| **Nó de IA** | Processar imagens em isolamento adicional; pesos de modelo só em formato seguro (sem execução de código na carga); computação confidencial avaliada para reduzir R15 |
| **MDM como adversário** | Assinatura do app verificada em tempo de execução contra chave que o MDM não controla; MDM em rede administrativa separada; alteração de política auditada |
| **Assistentes do sistema e resumos de tela** | Bloqueados por política e testados (A43) |
| **Wearables, USB, NFC, Bluetooth, autopreenchimento** | Política de rádios e portas no perfil de trabalho |
| **Criptografia pós-quântica** | Avaliar suíte híbrida MLS e TLS; decisão e ADR até a Fase 1 |
| **Transparência de chaves** | Registro verificável das chaves do diretório, para detectar servidor que entrega pacote falso; **antes do piloto** (D20) |
| **Modelo de ameaças por componente** | STRIDE por fronteira, versionado em `docs/` |
| **Verificação formal do expurgo** | Modelar o protocolo (ADR-08) em ferramenta de verificação |
| **Análise de tráfego** | Decidir preenchimento e envio em lotes (B05, R35) |
| **Segurança física** | Salas de servidor, HSM e estação de preparação com acesso controlado e registro |
| **Equipe vermelha contínua e SDL** | Revisão de segurança por funcionalidade, análise estática e dinâmica, fuzzing, SLSA medido |
| **Detecção de uso anômalo** | Regras de comportamento (doc 05 seção 4) |

---

## 10. Riscos novos (continuação da matriz)

| ID | Risco | P | I | Controles | Residual |
|---|---|---|---|---|---|
| R45 | Aparelho clonado ou identidade duplicada em segundo aparelho | Baixo | Crítico | Seção 5; sessão única; reatestação; credencial ligada ao hardware | Baixo |
| R46 | Uso da chave de identidade por aparelho com root sem exportá-la | Baixo | Crítico | Atestação do boot a cada sessão; autenticação por uso; anomalia | Médio |
| R47 | Aparelho comprometido por spyware antes da instalação | Médio | Crítico | Preparação em estação; triagem; aparelho da instituição | Médio |
| R48 | Spyware avançado não detectado pela triagem | Baixo | Crítico | Estado conhecido, reflash, triagem periódica, limites declarados | **Médio — não eliminável** |
| R49 | Falsificação de localização para burlar o perfil de local | Médio | Alto | Camada 1 de rede e camada 2 de proximidade; localização simulada bloqueada | Baixo |
| R50 | Câmera frontal contornada (coberta, desativada, ângulo cego) | Alto | Médio | Câmera obstruída = ocultar; combinação com outras camadas | Médio |
| R51 | Falso positivo da câmera atrapalha a operação | Médio | Médio | Ajuste de limiares; revisão mensal; ocultar sem punir | Baixo |
| R52 | Acervo de imagens de evidência (rostos de agentes e terceiros) é exfiltrado | Baixo | Crítico | Cifra para a chave de quórum de evidência (ADR-0003; abre só por quórum presencial em HSM offline); retenção curta; acesso pelo rito; só imagens de alerta, nunca vídeo contínuo; isolamento do armazenamento | Médio |
| R53 | Dado de localização de agentes exposto | Baixo | Crítico | Sem coordenadas no servidor; eventos por zona; retenção curta | Baixo |
| R54 | MDM comprometido empurra app ou política maliciosa | Baixo | Crítico | Verificação de assinatura independente; rede administrativa separada; auditoria | Médio |
| R55 | PIN de coação ou pânico acionados por engano (pressa, luva, escuro), PIN de coação confundido com o normal (na coação real, abre a conta real), ou isca percebida como falsa. **No perfil missão**, o PIN de coação por engano destruiria o aparelho, se ligado ao nível 4 (01g-G12a; doc 11, antes R-U2 e R-U3) | Médio | **Alto** no perfil missão; Médio no urbano | Treinamento e repetição depois de semanas sem uso; procedimento de retorno; isca revisada; PIN de coação diferente do normal por regra (sem sequência trivial); confirmação curta do pânico que não compromete o disfarce; acionamento acidental medido no V09 (doc 11, U05 e U06); **nível 4 pelo PIN de coação só depois dessa medição** (GD-24); sem consequência contra pessoa (regra 11) | Médio |
| R66 | Apagamento automático após cabo ou apreensão lido como destruição ou supressão de prova (ver R63) | Médio | Alto | Cabo: níveis 1 e 3, não 4; tentativas e pânico sem apagar até D-R14; lease vencido não apaga (ADR-0006); revogar não destrói: segredo de lease selado para quórum em toda revogação (GD-13); apagamento remoto e destruição do segredo com duas autorizações; evento assinado com motivo. **Perfil missão de campo (GD-18):** destruição do aparelho com rosto de outra pessoa e PIN errado, ou com PIN de coação, e destruição do nó capturado, aceitas pelo responsável para captor adversário; no uso urbano nada se destrói até D-R14 (seção 7.1a) | **Alto até D-R14.** Reavaliação proposta: Médio, depois de 01c-C01 fechado no texto, A72, A73 e A74 aprovados e aprovação do comitê (doc 04 seção 3.2) |
| R67 | **Falso negativo** da verificação facial (dono não reconhecido: pouca luz, ferimento, óculos, luva, pintura, capacete) leva a isca, a trava ou, no perfil missão de campo com PIN também errado, à destruição do aparelho do agente legítimo; o agente tenta atalho ou pede para desligar o recurso; ou um colega destrói de propósito o aparelho de outro, com o próprio rosto e um PIN errado (01g-G11c; doc 11, antes R-U4 e parte do R-U5) | Médio | Alto | Face nunca decide sozinha (GD-18); "inconclusivo" nunca conta como outra pessoa; "outra pessoa" só com duas capturas; falha de vivacidade fora do gatilho de destruição; número de PINs errados só depois de medir (GD-24); destruição só no perfil missão e com PIN errado; o nível 4 só atinge o aparelho; volta por revinculação; caminho de volta do dono que caiu na isca, pelo gestor e sem destruição (doc 11, fluxo 4.8); nunca revogação automática; revisão humana; medir (A51, A78); limite de falso negativo perto de zero em condição de campo (D-R40) | Médio |
| R116 | **Falso positivo ou falsificação** (foto, vídeo, máscara, semelhança) faz o reconhecimento aceitar quem não é o dono; com PIN obtido por coação ou observação, o adversário entra sem disparar a escada (doc 11, parte do antigo R-U5) | Médio | Alto | Vivacidade passiva no aparelho; PIN sempre exigido (a face não abre sozinha); PIN de coação; imagem de evidência; medir (A76). Limite declarado: câmera frontal comum, sem sensor de profundidade; o reconhecimento detecta, não barra (seção 7.2) | Médio |
| R117 | Gabarito facial do agente extraído do aparelho (dado biométrico sensível; ajuda a identificar o agente) | Baixo | Alto | Cifrado sob chave de dados embrulhada por chave não exportável do hardware seguro; decifrado e comparado só no core, em Rust, nunca na JVM nem no Swift (01g-G22a); regra 12 (GD-25); nunca no servidor, backup, log ou exportação; zerado no nível 4 (A80). Limite: aparelho com root observa a comparação | Baixo |
| R68 | Arquivo custodiado (D-R15) vira acervo único alvo de ataque; comprometer o HSM offline ou o quórum expõe os períodos guardados por até 5 anos | Baixo | Crítico | Só em sala sob custódia (GD-1); HPKE para chave pública, privada só no HSM offline; chave por sala, período e classe; quórum presencial M ≥ 3 com base registrada; eliminação no HSM com recibo; inventário de backups do HSM (ADR-0003) | **Médio** — proposto para aceite do responsável pela homologação, validade de 6 meses (doc 04 seção 3.2). O GD-1 limita o risco à sala sob custódia; com a flag ligada, a custódia é o padrão (GD-14) |
| R69 | **Revinculação indevida**: engenharia social ou conluio entrega o histórico a um aparelho de quem não deve; ou o aprovador assina sem ler (doc 11, antes R-U8) | Médio | Crítico | Só operação ativa; cadastro presencial, identificação do agente, série no inventário, termo vigente; **aprovação dupla** (chefe da inteligência + superior da área), distintos do agente e do gestor; resumo legível e curto na tela de aprovação, com hash visível e confirmação com biometria, testado no V09 (doc 11, T22 e U14); termo de ciência e aceite com foto; histórico só do que outro membro tem em cache; aviso ao agente, ao Auditor e aos membros; antigo aparelho revogado antes | Médio |
| R70 | Rastreio de localização do agente sem base legal ou além da perda | Baixo | Alto | Só após declaração de perda; posição em tempo real só pela sala de resposta (MLS); cópia guardada só para a chave de quórum; prazo curto; registro; parecer do jurídico | Médio |
| R79 | Conluio de 3 ou mais comandantes para abrir dados sem pedido legítimo | Baixo | Crítico | Pedido registrado antes do quórum; trilha imutável; aviso ao Auditor; ordem judicial para uso externo; revisão por órgão de controle | Médio |
| R80 | Coação ou roubo de cartões de comandantes | Baixo | Crítico | M ≥ 3; cartões em locais separados com PIN; quórum só em estação isolada; reemissão ao suspeitar | Baixo |
| R81 | HSM de recuperação online explorado | — | — | **Não se aplica na v1**: sem HSM online (ADR-0003 revisada, GD-4) | Retirado na v1 |
| R82 | Cadastro falso de aparelho aprovado de boa-fé por três comandantes | — | — | **Não se aplica na v1**: sem quórum remoto. O cadastro falso aprovado pelos dois aprovadores da revinculação está no R69 | Substituído pelo R69 |
| R83 | Aparelho de comandante comprometido ou coagido | — | — | **Não se aplica na v1**: comandante não aprova pelo celular; cartão fica fora de aparelho conectado (R80) | Retirado na v1 |
| R96 | Termo de revinculação com foto (dado biométrico, LGPD art. 11) exposto, no acervo do MIGDOL ou na cópia exportada ao sistema corporativo | Baixo | Alto | Cifrado no aparelho para a chave de quórum do termo; prazo (D-R30); hash na trilha; exportação só pelo fluxo autorizado e, até D-R30, sem a foto | Baixo |
| R97 | Remetente grava cópia de custódia inválida (lixo ou texto diferente) em sala sob custódia | Baixo | Médio | Hash da cópia dentro da mensagem MLS, conferido por quem recebe; assinatura do remetente dentro da cópia; servidor recusa envelope sem cópia; auditoria amostral nas cerimônias (ADR-0003) | Baixo — proposto para aceite do dono do ativo (`custody`), revisão anual (doc 04 seção 3.2); equivale a vazar por outro meio |
| R71 | Ordem de apagamento usada com abuso, ou MDM comprometido apaga a frota | Baixo | Alto | Ordem assinada com contador por aparelho; **duas autorizações em todo apagamento remoto**, inclusive em lote; MDM só a partir da ordem, com duas pessoas; auditoria; separação de funções (R54) | Médio |
| R56 | Exploração de parser de anexo (zero-click) | Médio | Crítico | Reencode; isolamento; fuzzing | Médio |

---

## 11. Cenários de teste (continuação do documento 03)

### Clonagem e identidade (A25–A30)
| # | Cenário | Esperado |
|---|---|---|
| A25 | Duas conexões com a mesma identidade | Nova conexão com contador de sessão maior substitui a anterior, sem SEV1. Duas sessões simultâneas com contadores conflitantes: ambas derrubadas e SEV1 (D18) |
| A26 | Restaurar dados do app em outro aparelho (por ferramentas do fabricante e por cópia manual) | Dados inúteis |
| A27 | Cadastrar nova biometria | Chaves invalidadas |
| A28 | App rodando em clonador ou segundo usuário | Recusado |
| A29 | Reatestação com desafio reutilizado | Recusada |
| A30 | Credencial MLS sem resumo de atestação | Clientes recusam |

### Série e local (A31–A37)
| # | Cenário | Esperado |
|---|---|---|
| A31 | Aparelho com série fora do inventário | Recusado |
| A32 | Série divergente da atestada | Recusado; incidente |
| A33 | Localização simulada | Bloqueada ou ignorada |
| A34 | Conexão por rede não autorizada a sala com perfil de local | Negada |
| A35 | Sair da zona com a sala aberta | Conteúdo some; cache da sala fechado |
| A36 | Baliza de proximidade clonada ou repetida | Rejeitada pelo desafio-resposta |
| A37 | Coordenadas e histórico de posição no servidor e nos logs | Inexistentes |

### Câmera e coação (A38–A44)
| # | Cenário | Esperado |
|---|---|---|
| A38 | Celular apontado para a tela, de frente, vários ângulos e distâncias | Medir taxa de detecção |
| A39 | Câmera pequena ou oculta, ângulo cego | Registrar limite |
| A40 | Câmera frontal coberta | Oculta |
| A41 | Terceiro olhando por cima do ombro | Oculta |
| A42 | Falso positivo em uso normal por 1 hora | Medir taxa |
| A43 | Quadro da câmera em disco, memória persistente, rede e logs, **sem alerta** | Nada gravado ou enviado. **Com alerta:** só as imagens de evidência, cifradas para a chave de quórum |
| A44 | PIN de coação e botão de pânico | Conta isca; alerta chega; nada vaza ao observador |

### Triagem e aparelho (A45–A47)
| # | Cenário | Esperado |
|---|---|---|
| A45 | Aparelho com app espião conhecido na triagem | Reprovado |
| A46 | Aparelho com certificado de usuário instalado | Detectado e bloqueado |
| A47 | Aparelho recuperado após perda | Suspenso até reflash e triagem |

### Anexos e MDM (A48–A50)
| # | Cenário | Esperado |
|---|---|---|
| A48 | Imagem e PDF malformados (fuzzing) | Sem falha de memória; recusa |
| A49 | App com assinatura trocada empurrado pelo MDM | Recusado em tempo de execução |
| A50 | Assistente do sistema lendo a tela | Sem conteúdo |
| A51 | Cabo de dados com conteúdo na tela; PIN correto com rosto diferente (GD-18) | Nível 1 e 3, sem apagar; nível 2 no caso do rosto; medir erro |
| A52 | Instrumentação ativa com sessão aberta | Nível 4; aviso de autoexclusão; evento assinado; administrador remove o aparelho no prazo ou recebe alerta; agente informado na próxima autenticação |
| A53 | Depois do nível 4, o servidor e os outros membros mantêm o conteúdo | Só este aparelho perde o acesso |
| A54 | Abertura de custódia sem número de CI, sem base registrada (ordem judicial, ou processo administrativo com validação jurídica) ou com menos de 3 cartões | Recusada |
| A55 | Reset de fábrica pelas Configurações e pelo modo de recuperação, em aparelho gerenciado | Primeiro bloqueado; o segundo medido por modelo; aparelho resetado não lê nada |
| A56 | Aparelho resetado tenta se cadastrar sem QR presencial | Recusado |
| A57 | Ordem de apagamento repetida, ou sem assinatura | Rejeitada pelo contador e pela verificação |
| A58 | Revinculação com uma só aprovação, ou em sala concluída | Recusada |
| A59 | Aparelho offline recebe ordem de apagamento | Ilegível pelo lease criptográfico (servidor selou o segredo para quórum e não o entrega); apaga ao reconectar se a ordem tiver as duas assinaturas; evento registrado no servidor |
| A60 | Revinculação aprovada pelo próprio agente, pelo gestor que cadastrou, ou duas vezes pela mesma pessoa | Recusada |
| A61 | HSM de recuperação recebe pedido para aparelho de outro agente | **Não se aplica na v1** (sem HSM online) |
| A62 | Aprovação com foto ou vídeo do rosto do comandante | **Não se aplica na v1** (sem aprovação remota por comandante) |
| A63 | Aprovação com PIN de coação | **Não se aplica na v1** |
| A64 | Comandante tenta ver o conteúdo durante a aprovação | **Não se aplica na v1** |
| A65 | Envelope de sala sob custódia sem cópia de custódia | Servidor recusa |
| A66 | Commit que altera ou remove a marca de custódia | Clientes recusam |
| A67 | Membro tenta decifrar a cópia de custódia com o próprio estado de grupo | Impossível: a chave privada está só no HSM offline |
| A68 | Servidor troca a cópia de custódia de uma mensagem | Quem recebe detecta pelo hash dentro da mensagem MLS; alerta |
| A69 | Conclusão com duas assinaturas, com o mesmo signatário duas vezes, sem número e hash do relatório, ou com solicitante que é comandante ou subcomandante sem a assinatura do superior imediato dele | Recusada |
| A70 | Pacote de histórico da revinculação com mensagens de fora dos intervalos de pertença (do pedido assinado pelos aprovadores) ou da retenção | Aparelho novo recusa o excedente |
| A71 | Foto do termo de revinculação em disco em claro, no servidor ou na exportação (até D-R30) | Inexistente; só cifrada para a chave de quórum |
| A72 | N erros de desbloqueio e pânico, sem decisão D-R14 | Nível 3; nada apagado |
| A73 | Lease vencido; segredo antigo reapresentado depois da renovação; queda no meio da renovação | Cache ilegível e não apagado; segredo antigo não abre; queda retoma pela geração sem perder o cache (ADR-0006) |
| A74 | Revogação de aparelho por qualquer motivo; destruição do segredo com uma só autorização | Segredo de lease selado para quórum e sem cópia em claro no servidor; nada destruído; destruição com uma autorização recusada (GD-13) |
| A75 | Sala comum criada com a flag ligada sem motivo ou com uma só assinatura; sala de sistema com cópia de custódia; sala de usuário marcada como sistema | Recusada pelos clientes (GD-14) |

### Verificação facial e escada por perfil (A76–A83; GD-18)
| # | Cenário | Esperado |
|---|---|---|
| A76 | Foto, vídeo em tela e máscara do rosto do dono, com PIN correto e com PIN errado | Falsificação detectada → nível 2 (PIN correto) ou nível 3 (PIN errado), nunca nível 4 até a vivacidade ser medida à parte (GD-24); a conta real não abre; medir taxa de falsificação aceita e de falha de vivacidade do dono |
| A77 | Pessoa diferente com PIN errado, em duas capturas e numa só | Duas capturas: perfil missão, nível 4 e imagem de evidência, só depois das medições da GD-24 (antes, nível 3); perfil urbano, nível 3, nada destruído. Uma captura só: nunca nível 4 |
| A78 | Dono com pouca luz, rosto coberto, óculos, luva ou ferimento, com PIN errado | "Inconclusivo" ou falso negativo nunca leva ao nível 4; medir taxa de falso negativo |
| A79 | PIN de coação | Perfil missão: isca exibida e nível 4 por baixo, só depois de medido o acionamento acidental (GD-24; antes, isca e nível 3), sem diferença perceptível na tela, no tempo de resposta, no aquecimento nem no consumo; alerta enfileirado e enviado antes de zerar, se houver enlace (01g-G12b, c). Perfil urbano: isca e nível 3 |
| A80 | Gabarito facial em disco em claro, no servidor, em backup, em log ou na exportação | Inexistente; só cifrado pela chave do hardware seguro |
| A81 | Quadro da verificação facial em disco, memória persistente, rede e logs | Nada gravado nem enviado; com "outra pessoa", só a imagem de evidência cifrada para quórum |
| A82 | Perfil missão fora de missão em andamento, ou ligado só pelo servidor | Vale a escada comum; sem destruição |
| A83 | Rosto de outra pessoa sem digitar PIN, ou com PIN correto | Nunca nível 4 (a face não decide sozinha) |

### Investigação sigilosa (A84–A95; ADR-0010)
| # | Cenário | Esperado |
|---|---|---|
| A84 | Pedido sem CI validada, sem segunda autorização, com a segunda dada pelo próprio iniciador, com a segunda dada por detentor quando a suspeita alcança os detentores (GD-22), ou assinado fora da estação isolada | Recusado |
| A85 | Segunda autorização, cartão de detentor, operador, testemunha ou revisor que consta da lista de impedidos, ou que é membro ou administrador de sala do escopo no período; iniciador nessa posição tentando operar a estação, revisar ou receber o relatório; participante cuja chave não tem identificador opaco conferível na lista assinada de papéis (01g-G03, G05) | Recusado pela estação (falha fechada); o ato segue só com pessoa não impedida |
| A86 | Abertura de cópia de custódia sem ordem judicial registrada; qualquer sessão, de qualquer nível, sem CI validada nem ordem judicial; qualquer sessão com menos de M cartões não impedidos (GD-21); terceira renovação (GD-22) | Recusada |
| A87 | Membros das salas do escopo, nó de IA e servidor durante a investigação | Nada observável muda (membros, marca, comportamento do cliente, notificações); a entrada na trilha tem tamanho e formato iguais aos de outros registros selados |
| A88 | Incluir membro sem aviso; mudar a marca de sala comum para custódia; ligar cópia de custódia em sala comum | Clientes recusam (sem membro oculto; GD-17) |
| A89 | Cerimônia de eliminação tenta eliminar chave de período que consta de lista de bloqueio vigente, ou com objeto faltando para hash de registro selado da trilha; `custody` tenta apagar cifrado sem recibo de eliminação | Estação recusa sem abrir o registro selado nem usar a chave da série de investigação, mostrando só a contagem (01g-G01c); `custody` não apaga |
| A90 | Estação de análise com rede, disco persistente, ou saída que não seja o pacote assinado | Recusada no ensaio; nada em claro sai |
| A91 | Relatório de rastro com trecho de conversa (sobreposição com a fonte acima do limite) ou fora do formato estruturado | Bloqueado pela validação da saída |
| A92 | Relatório sem a revisão e a assinatura de duas pessoas não impedidas, ou revisado pelo iniciador ou pela segunda autorização (GD-22) | Pacote de evidência não é gerado |
| A93 | Mensagem plantada numa sala do escopo para induzir a análise (injeção) | Não muda o rastro além de indício rotulado; comportamento registrado |
| A94 | Registro selado apagado do `custody` | Detectado pela conferência do hash na trilha |
| A95 | Pedido de análise ao nó de IA vindo de fora da sala, ou para a investigação | Impossível: o nó só responde a pedido feito dentro da sala |

### Análise entre operações (A96–A99; ADR-0008)
| # | Cenário | Esperado |
|---|---|---|
| A96 | Nó de IA cruza sala sem concessão vigente vista no grupo, ou com concessão assinada por quem não é administrador da sala; indicador de mensagem anterior à concessão (não retroage, 01g-G23b) | Recusa; a sala, ou a parte anterior à concessão, fica fora do escopo |
| A97 | Sala de análise sob custódia com sala comum no escopo | Recusada (regra de tipo) |
| A98 | Conclusão, expurgo ou retirada da concessão na sala de origem | Indicadores daquela origem somem da sala de análise e do índice do nó (expurgo em cascata) |
| A99 | Resposta com trecho de conversa ou fora do formato estruturado; pergunta de um escopo tentando obter dado de outro; injeção dentro do escopo | Bloqueada ou registrada; nada vaza entre escopos (R104) |

Critério: A25–A37 e A43–A50 aprovados; A65–A99 aprovados na fase de cada recurso (doc 03 seção 5); **A38, A39 e A42 documentados como limites conhecidos**, mesmo sem correção possível; taxas de A76 e A78 documentadas antes de ligar o reconhecimento ao gatilho de destruição.

---

## 12. Requisitos novos

| ID | Requisito | Prioridade |
|---|---|---|
| RF59 | Preparação de aparelho em estação controlada, com triagem registrada no cadastro | Essencial |
| RF60 | Validação do número de série atestado contra o inventário em todo cadastro e sessão | Essencial |
| RF61 | Perfis de local por sala (livre, rede autorizada, local autorizado, compartimento) | Importante |
| RF62 | Detecção de usuário ausente, terceiro e câmera coberta, com ocultação do conteúdo | Importante |
| RF63 | Detecção de câmera ou celular apontado para a tela, com evento de alta prioridade | Importante |
| RF64 | PIN de coação, botão de pânico e trava por tentativas; apagamento por tentativas ou pânico só com decisão D-R14 | Importante |
| RF65 | Sessão única por aparelho (substituição com contador assinado) e detecção de bifurcação de época (ADR-0005) | Essencial |
| RF66 | Reatestação por sessão e periódica; atestação sem serviços de terceiros | Essencial |
| RF69 | Escada de resposta em quatro níveis (seção 7.1); apagamento local registrado e assinado; cabo de dados não aciona o nível 4; lease vencido não apaga; revogar não destrói (segredo de lease selado para quórum; doc 01 RF91) | Importante |
| RF70 | Custódia só em sala sob custódia (padrão com a flag ligada; sala comum com motivo e segunda assinatura; sala de sistema sempre comum, GD-14): cópia HPKE para chave pública de quórum; abertura só por quórum presencial com CI e base registrada (ordem judicial, ou processo administrativo com validação jurídica **[conferir]**); guarda até 5 anos da conclusão ou do expurgo; eliminação com recibo (seção 7.3; doc 01 RF82) | Proposto (ADR-0003 revisada) |
| RF71 | Declaração de perda com motivo; trava imediata; apagamento total do aparelho por ordem assinada com duas autorizações e registrada no servidor, sem resíduo no aparelho; apagamento em etapas | Importante |
| RF72 | Revinculação só em operação ativa, com aprovação dupla (chefe da inteligência + superior da área), termo de ciência e aceite (doc 01 RF84) e histórico enviado por outro membro pelo MLS (seção 7.4) | Proposto |
| RF73 | Modo rastreio por localização depois da declaração de perda: posição pela sala de resposta (MLS), cópia guardada só para a chave de quórum, com prazo | Proposto; depende do jurídico |
| RF77 | ~~Aprovação remota da revinculação pelo app do comandante~~ | **Retirado na v1** (GD-4; ADR-0003 revisada) |
| RF95 | Verificação facial de quem acessa (GD-18; seção 7.2): no desbloqueio do app, modelo no aparelho, sem terceiros, com vivacidade; gabarito cifrado por chave do hardware seguro; resultado "dono", "outra pessoa" ou "inconclusivo"; a face nunca decide sozinha | Importante (Fase 2), depois do parecer jurídico **[conferir]** |
| RF96 | Escada por perfil (seção 7.1a): perfil missão de campo com nível 4 para rosto de outra pessoa (duas capturas) e PIN errado e para PIN de coação (por baixo da isca), cada gatilho ligado só depois da medição exigida pela GD-24; perfil urbano suspende e revoga sem destruir (D-R14); perfil missão só durante missão em andamento, registrada no estado do grupo | Importante (Fase 2; perfil missão completo na Fase 8) |
| RNF-G1 | Sem alerta, quadros da câmera não são gravados nem enviados. Com alerta, somente as imagens de evidência, cifradas para a chave de quórum, com retenção definida | Essencial |
| RNF-G2 | Câmera usada apenas com o app em primeiro plano: na tela de conteúdo (seção 6) e na tela de desbloqueio, para a verificação facial (seção 7.2; GD-18; parecer jurídico **[conferir]**) | Essencial |
| RNF-G3 | Verificação facial com modelo no aparelho, hash fixado, sem rede nem serviço de terceiro; quadros só em memória; gabarito só no aparelho, cifrado por chave não exportável do hardware seguro, nunca no servidor, backup, log ou exportação | Essencial, se adotada |

---

## 13. Decisões pendentes

| ID | Decisão | Quem decide |
|---|---|---|
| D-R6 | ~~Guardar evidência de imagem em alerta de câmera apontada~~ **Decidido: adotada** (seção 6.3). Resta o jurídico conferir base legal e texto de ciência, e definir a retenção | Jurídico + comitê |
| D-R7 | Quais salas usam cada perfil de local; como conciliar com operação de campo | Chefia da unidade |
| D-R8 | ~~Reconhecimento facial próprio~~ **Decidido (GD-18):** verificação facial de quem acessa, no aparelho, com vivacidade; a face nunca decide sozinha. Falta o parecer jurídico (biometria, LGPD art. 11, ANPD) e os limites de erro (D-R40) | Jurídico + encarregado (parecer) |
| D-R9 | Modo de coação: uso, treinamento, quem recebe o alerta | Chefia + segurança |
| D-R10 | Ferramenta de triagem de spyware e origem dos indicadores | Segurança |
| D-R11 | Atestação sem Play Integrity (ADR-0002) | Segurança + desenvolvimento |
| D-R12 | Criptografia pós-quântica híbrida | Segurança + desenvolvimento |
| D-R13 | Check-in periódico em operações de risco | Chefia |
| D-R14 | Expurgo ou apagamento após apreensão por autoridade (ver doc 04) | Jurídico + comitê |
| D-R15 | **Definido pelo responsável do produto (GD-1 a GD-6; ADR-0003 revisada):** custódia só em sala sob custódia, por flag da instituição; cópia HPKE para chave pública; abertura por quórum presencial de no mínimo 3 comandantes em HSM offline, com CI (sempre) e ordem judicial ou processo administrativo com validação jurídica; guarda até 5 anos da conclusão; revinculação sem custódia; com a flag ligada, custódia é o padrão (GD-14). **Falta o jurídico:** processo administrativo como base, prazo, eliminação (D-R5) | Chefia + jurídico + segurança |
| D-R16 | Persistência de matrícula depois do reset e rastreio (proteção de redefinição, Zero-touch, Knox) frente à regra 10 | Segurança + chefia |
| D-R29 | ~~Com a flag de custódia ligada, alguma sala pode nascer comum?~~ **Decidido (GD-14, opção B):** custódia é o padrão; sala comum só com motivo registrado e segunda assinatura; salas de sistema sempre comuns (ADR-0003 seção 1). Resta D-R36 | Responsável pelo produto (decidido) |
| D-R35 | Prazo padrão de guarda do segredo de lease selado na revogação ("prazo da apuração", GD-13; ADR-0006 item 8) | Responsável pelo produto + jurídico |
| D-R36 | Quem dá a segunda assinatura da sala comum com a flag ligada, e se o motivo segue lista fechada no `instituicao.yaml` (GD-14; proposta: categoria fechada na trilha, texto livre só no grupo, 01e-E09b) | Responsável pelo produto + chefia |
| D-R40 | Escada por perfil (GD-18, **decidida em parte pela GD-24**). Decidido: "outra pessoa" exige duas capturas; falha de vivacidade fora do gatilho de destruição até ser medida à parte; número de PINs errados só depois de medir (proposta: um, se a medição permitir); PIN de coação no nível 4 só depois de medir o acionamento acidental. **Falta:** os valores medidos; nível 4 por tentativas no perfil missão; se a polícia liga o perfil missão nas missões de campo dela; limites de falso negativo (perto de zero para o dono, em condição de campo) e de falso positivo para o reconhecimento entrar no gatilho de destruição (inclui a parte facial da antiga U-3 do doc 11) | Responsável pelo produto + segurança + jurídico |
| D-R30 | Foto do termo de revinculação: pode ir ao sistema corporativo (regra 13)? A captura fora de alerta está fora da regra 12? Prazo de guarda do termo. Até a decisão: exporta-se o termo com o hash da foto, não a foto | Responsável pelo produto + encarregado + jurídico |

---

## 14. Entradas no roteiro

| Fase | Acréscimo |
|---|---|
| **0** | ADR-0002 (atestação sem Play Integrity); spike de atestação de identificadores; decisões D-R10 a D-R12; processo da estação de preparação; protótipo da câmera (taxa de acerto, bateria) |
| **1** | Sessão única; reatestação; credencial ligada ao hardware (A25, A29, A30) |
| **2** | Detecção por câmera (A38–A43); perfil de rede autorizada (A34); série atestada (A31–A32); PIN de coação e pânico (A44); lease e revogação sem destruição (A73, A74); verificação facial e escada por perfil (A76–A83; perfil missão completo na Fase 8) |
| **3** | Reencode e isolamento de parsers (A48) |
| **4** | Perfis de local e proximidade (A33–A37); ciclo da sala, marca e cópia de custódia, sala comum por exceção e sala de sistema, revinculação v1 (A54, A58–A60, A65–A71, A75); investigação sigilosa: rito, impedimento, registro selado, preservação, estação (A84–A90, A92, A94, A95; ADR-0010) |
| **5** | Isolamento adicional do nó de IA; avaliação de computação confidencial |
| **7** | Triagem recorrente testada (A45–A47); relatório de limites A38, A39, A42; equipe vermelha |
| **Depois da 6** | Análise automatizada da investigação sigilosa na estação (A91, A93); análise entre operações com concessão (A96–A99; ADR-0008) |

---

## 15. Impacto em LGPD (documento 06)

Acrescentar ao inventário da seção 3:

| Dado | Titular | Finalidade | Tratamento |
|---|---|---|---|
| Imagem da câmera frontal | Usuário e terceiros no quadro | Detectar presença e câmera apontada | Fluxo contínuo processado em memória, sem gravação; imagens só em alerta |
| Eventos de câmera (terceiro, câmera apontada, câmera coberta) | Usuário | Segurança e apuração | Trilha; retenção a definir |
| Evidência de imagem (alerta) | Usuário e terceiros | Atribuição | **Ativa**; cifrada para a chave de quórum; retenção definida pela instituição; acesso pelo rito |
| Foto e assinatura do termo de revinculação | Usuário | Ciência e aceite da reposição do aparelho | Dado biométrico sensível (art. 11); cifrado para a chave de quórum do termo; prazo D-R30; hash na trilha |
| Gabarito facial do dono (GD-18) | Usuário | Verificar quem acessa o aparelho | Dado biométrico sensível (art. 11); só no aparelho, cifrado por chave do hardware seguro; zerado no nível 4 e na devolução do aparelho |
| Posição em modo perda | Usuário | Recuperar aparelho e proteger o agente | Tempo real pela sala de resposta (MLS); cópia guardada só para a chave de quórum; prazo curto |
| Segredo de lease guardado de aparelho revogado | Usuário (indireto: abre o cache do aparelho dele) | Preservar o cache para apuração (GD-13) | Cifrado para a chave de quórum; prazo da apuração (D-R35); sozinho não abre nada sem o aparelho e o desbloqueio do agente |
| Eventos de zona (entrada, saída, acesso negado) | Usuário | Controle de local | Sem coordenadas; retenção curta |
| Número de série e identificadores atestados | Aparelho (vinculável ao usuário) | Controle de acesso | Diretório; acesso restrito |
| Resultado de triagem de spyware | Aparelho e usuário | Segurança | Cadastro do aparelho; retenção definida |

Itens do RIPD: monitoramento por câmera, localização, risco de falso positivo, reconhecimento facial adotado no desbloqueio (GD-18) e destruição por perfil, transparência ao usuário (RN15 passa a listar câmera, verificação facial e local). A responsabilidade pelas escolhas é da instituição (doc 06 seção 0).
