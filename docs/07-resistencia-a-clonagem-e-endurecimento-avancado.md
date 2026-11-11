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
| **Uma sessão por aparelho** | Segunda conexão com a mesma identidade derruba ambas e abre SEV1 |
| **Detecção de bifurcação de época** | Clientes e nó de IA detectam dois estados de grupo divergentes do mesmo membro |
| **Chave usada em cada conexão** | A chave de identidade assina o handshake mTLS **de toda conexão**, não só do cadastro |
| **Reatestação periódica** | A cada sessão e a cada N horas, com desafio novo do servidor |
| **Chaves com autenticação por uso** | Chave de administrador exige biometria a cada uso; chave de identidade exige aparelho desbloqueado |
| **Invalidação por mudança de biometria** | Nova digital ou rosto cadastrado invalida as chaves; reativação pelo gestor |
| **Credencial com vínculo de hardware** | A credencial MLS do membro inclui o resumo da atestação; outros clientes recusam credencial sem hardware |
| **Atualização periódica de chaves** | Commit de atualização MLS a cada 24 h (proposta) para limitar janela de comprometimento |
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
| **Pessoa diferente do dono** | Possível, mas **sensível** | Exige reconhecimento facial (dado biométrico, doc 06). Não recomendado na v1 (ver 6.4) |
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
  - cifradas **para a chave do Auditor** (nem o servidor nem o administrador da operação leem);
  - com hora, aparelho e resultado do detector;
  - retenção definida pela instituição (proposta: 90 dias, ou até o fim do procedimento aberto);
  - consulta somente pelo rito do documento 05 seção 3.4, com registro.
- A decisão de adotar vale como decisão **da instituição**; o jurídico ainda confere a base legal e o texto de ciência (doc 06 seção 7), porque isso protege a validade da prova.
- O indicador de câmera do sistema acende; o usuário **sabe** que a câmera está ativa (RN15).
- **Cuidados vindos da verificação jurídica** (docs/anexos/validacao-juridica-preliminar.md seção 3.5): a câmera só funciona com conteúdo sensível na tela; a imagem de evidência é **recortada** à região relevante e de baixa resolução; nada é capturado fora de alerta; sala sensível não deve ser lida em ambiente privado (perfil de local); a câmera não é usada em tela de configuração, chamada ou fora do app. Não encontrei decisão sobre câmera de aparelho funcional: **parecer do jurídico** antes da Fase 2.
- O indicador de câmera do sistema acende; o usuário **sabe** que a câmera está ativa (transparência, RN15).
- **Fora da tela de conteúdo, a câmera não é usada.**
- Custo: bateria e aquecimento. Medir na Fase 2 (taxa de quadros baixa, ex.: 1–2 por segundo, e modelo leve).

### 6.4 Reconhecimento facial: opcional, decisão D-R8
A instituição pode decidir adotá-lo. Cuidados que **não são de privacidade**, e sim de segurança e de eficácia:
- **Modelos próprios têm erro.** Falso positivo trava o agente; falso negativo deixa passar. Medir antes (A42).
- **Modelo facial é alvo de falsificação** (foto, vídeo, máscara). Exige detecção de vivacidade e testes.
- **O gabarito facial é dado de altíssimo valor** para o adversário. Se adotado: gabarito **só no aparelho**, cifrado em hardware, nunca no servidor.
- **A biometria do sistema** (digital ou rosto de classe forte) já autentica o dono do aparelho com mais maturidade; o reconhecimento próprio só se justifica para **verificar continuamente** quem está lendo.
Recomendação mantida: v1 sem reconhecimento próprio; decidir na Fase 2 com os números de A38–A42. A ANPD abriu tomada de subsídios sobre biometria em 2025 e pretende concluir regras em 2026; acompanhar antes de decidir.

### 6.5 Honestidade sobre eficácia
- A detecção de câmera apontada **reduz** o vazamento por foto, mas não o elimina: um atacante informado fotografa de ângulo cego, usa câmera oculta, ou cobre e descobre que o app oculta o conteúdo e tenta outro modo.
- Valor real: **dissuasão**, **redução de oportunidade** e **evidência de intenção** quando o alerta ocorre.
- Taxa de acerto e de falso positivo são medidas por cenário (A38–A42). Meta de partida: ≥ 80% de detecção em foto frontal a até 1 m e ≤ 2% de falso positivo por hora de uso; **hipótese a validar**.

---

## 7. Modo de coação

Não há como **detectar** coação por sensor com confiança. O que se faz é dar ao usuário **meios discretos de acionar** uma resposta, e usar a presença de terceiro como **sinal auxiliar**.

| Recurso | Funcionamento |
|---|---|
| **PIN de coação** | Segundo PIN que abre uma conta **isca** (conversas e salas genéricas, sem dado real), envia alerta silencioso e marca a sessão. Nunca revela que é isca |
| **Gesto ou botão de pânico** | Ação discreta que oculta tudo, envia alerta e, se configurado, **apaga a chave local** |
| **Apagamento por tentativas** | Depois de N erros de desbloqueio, apaga o cache local e invalida a sessão |
| **Presença de terceiro** | Contagem de rostos como sinal auxiliar de risco; aciona oculte e alerta, não conclusões |
| **Prazo de check-in** (opcional) | Para operações de risco: se o aparelho não renovar a presença no prazo definido, o servidor trata como possível apreensão (playbook P2) |

Cuidados:
- **PIN de coação e pânico exigem treinamento e decisão de unidade.** O erro mais provável é o usuário acionar sem querer: confirmação curta onde não comprometa o disfarce, e procedimento de retorno.
- A conta isca precisa parecer real o suficiente; não é pelo código que ela se sustenta, é pelo procedimento.
- O alerta precisa chegar a alguém que **age**: define-se na Fase 0 quem recebe e em quanto tempo responde (doc 05 seção 9).
- Isso **não resolve** coação física; reduz o dano e dá tempo à equipe.

---

## 8. Atestação sem Play Integrity (ADR proposta)

**Conflito:** o documento 02 prevê Play Integrity, que depende de serviço do Google e contraria a regra 10.

**Proposta (ADR-0002):**
1. **Atestação de chave do hardware**, validada no servidor da instituição contra as raízes de confiança dos fabricantes, mantidas internamente e atualizadas por processo próprio.
2. Verificar: estado do boot verificado, nível de patch, bloqueio do bootloader, chave em hardware seguro, identificadores (seção 3).
3. **Não usar Play Integrity** nem serviços do Google Play no aparelho.
4. Aparelho sem esses requisitos não entra na lista de modelos homologados.

Consequências: o servidor passa a **manter a lista de raízes e a política de aceite**; o spike de atestação da Fase 0 precisa cobrir esse processo.

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
| **Transparência de chaves** | Registro verificável das chaves do diretório, para detectar servidor que entrega pacote falso |
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
| R52 | Acervo de imagens de evidência (rostos de agentes e terceiros) é exfiltrado | Baixo | Crítico | Cifra para o Auditor; retenção curta; acesso pelo rito; só imagens de alerta, nunca vídeo contínuo; isolamento do armazenamento | Médio |
| R53 | Dado de localização de agentes exposto | Baixo | Crítico | Sem coordenadas no servidor; eventos por zona; retenção curta | Baixo |
| R54 | MDM comprometido empurra app ou política maliciosa | Baixo | Crítico | Verificação de assinatura independente; rede administrativa separada; auditoria | Médio |
| R55 | PIN de coação acionado por engano, ou isca percebida como falsa | Médio | Médio | Treinamento; procedimento de retorno; isca revisada | Médio |
| R56 | Exploração de parser de anexo (zero-click) | Médio | Crítico | Reencode; isolamento; fuzzing | Médio |

---

## 11. Cenários de teste (continuação do documento 03)

### Clonagem e identidade (A25–A30)
| # | Cenário | Esperado |
|---|---|---|
| A25 | Duas conexões com a mesma identidade | Ambas derrubadas; SEV1 |
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
| A43 | Quadro da câmera em disco, memória persistente, rede e logs, **sem alerta** | Nada gravado ou enviado. **Com alerta:** só as imagens de evidência, cifradas ao Auditor |
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

Critério: A25–A37 e A43–A50 aprovados; **A38, A39 e A42 documentados como limites conhecidos**, mesmo sem correção possível.

---

## 12. Requisitos novos

| ID | Requisito | Prioridade |
|---|---|---|
| RF59 | Preparação de aparelho em estação controlada, com triagem registrada no cadastro | Essencial |
| RF60 | Validação do número de série atestado contra o inventário em todo cadastro e sessão | Essencial |
| RF61 | Perfis de local por sala (livre, rede autorizada, local autorizado, compartimento) | Importante |
| RF62 | Detecção de usuário ausente, terceiro e câmera coberta, com ocultação do conteúdo | Importante |
| RF63 | Detecção de câmera ou celular apontado para a tela, com evento de alta prioridade | Importante |
| RF64 | PIN de coação, botão de pânico e apagamento por tentativas | Importante |
| RF65 | Sessão única por aparelho e detecção de bifurcação de época | Essencial |
| RF66 | Reatestação por sessão e periódica; atestação sem serviços de terceiros | Essencial |
| RNF-G1 | Sem alerta, quadros da câmera não são gravados nem enviados. Com alerta, somente as imagens de evidência, cifradas para o Auditor, com retenção definida | Essencial |
| RNF-G2 | Câmera usada apenas com conteúdo na tela e app em primeiro plano | Essencial |

---

## 13. Decisões pendentes

| ID | Decisão | Quem decide |
|---|---|---|
| D-R6 | ~~Guardar evidência de imagem em alerta de câmera apontada~~ **Decidido: adotada** (seção 6.3). Resta o jurídico conferir base legal e texto de ciência, e definir a retenção | Jurídico + comitê |
| D-R7 | Quais salas usam cada perfil de local; como conciliar com operação de campo | Chefia da unidade |
| D-R8 | Reconhecimento facial próprio (recomendado não na v1) | Jurídico + encarregado |
| D-R9 | Modo de coação: uso, treinamento, quem recebe o alerta | Chefia + segurança |
| D-R10 | Ferramenta de triagem de spyware e origem dos indicadores | Segurança |
| D-R11 | Atestação sem Play Integrity (ADR-0002) | Segurança + desenvolvimento |
| D-R12 | Criptografia pós-quântica híbrida | Segurança + desenvolvimento |
| D-R13 | Check-in periódico em operações de risco | Chefia |

---

## 14. Entradas no roteiro

| Fase | Acréscimo |
|---|---|
| **0** | ADR-0002 (atestação sem Play Integrity); spike de atestação de identificadores; decisões D-R10 a D-R12; processo da estação de preparação; protótipo da câmera (taxa de acerto, bateria) |
| **1** | Sessão única; reatestação; credencial ligada ao hardware (A25, A29, A30) |
| **2** | Detecção por câmera (A38–A43); perfil de rede autorizada (A34); série atestada (A31–A32); PIN de coação e pânico (A44) |
| **3** | Reencode e isolamento de parsers (A48) |
| **4** | Perfis de local e proximidade (A33–A37) |
| **5** | Isolamento adicional do nó de IA; avaliação de computação confidencial |
| **7** | Triagem recorrente testada (A45–A47); relatório de limites A38, A39, A42; equipe vermelha |

---

## 15. Impacto em LGPD (documento 06)

Acrescentar ao inventário da seção 3:

| Dado | Titular | Finalidade | Tratamento |
|---|---|---|---|
| Imagem da câmera frontal | Usuário e terceiros no quadro | Detectar presença e câmera apontada | Fluxo contínuo processado em memória, sem gravação; imagens só em alerta |
| Eventos de câmera (terceiro, câmera apontada, câmera coberta) | Usuário | Segurança e apuração | Trilha; retenção a definir |
| Evidência de imagem (alerta) | Usuário e terceiros | Atribuição | **Ativa**; cifrada para o Auditor; retenção definida pela instituição; acesso pelo rito |
| Eventos de zona (entrada, saída, acesso negado) | Usuário | Controle de local | Sem coordenadas; retenção curta |
| Número de série e identificadores atestados | Aparelho (vinculável ao usuário) | Controle de acesso | Diretório; acesso restrito |
| Resultado de triagem de spyware | Aparelho e usuário | Segurança | Cadastro do aparelho; retenção definida |

Itens do RIPD: monitoramento por câmera, localização, risco de falso positivo, decisão sobre reconhecimento facial, transparência ao usuário (RN15 passa a listar câmera e local). A responsabilidade pelas escolhas é da instituição (doc 06 seção 0).
