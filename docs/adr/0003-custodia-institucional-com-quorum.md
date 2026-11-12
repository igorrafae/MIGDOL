# ADR-0003 — Custódia institucional de dados com quórum de comandantes

**Status:** proposta · **Contraria** o objetivo "sem histórico permanente" (doc 01) e a retenção só por operação (RN03); **complementa** ADR-04 (expurgo por destruição de chave) e ADR-10 (anexos cifrados no cliente) e doc 04 seção 6.4 (chaves da instituição) · **Depende de** parecer do jurídico (D-R5, D-R15)

## Contexto
O responsável pelo produto definiu:
1. Todo dado (texto, imagem, anexo, conversa) é cifrado.
2. A instituição só decifra com um **quórum mínimo de 3 comandantes da companhia**, cada um com uma parte da chave.
3. Mensagens apagadas no aparelho continuam cifradas no servidor; o aparelho revinculado recebe de volta as suas mensagens (doc 07 seção 7.4).
4. Entrega a autoridade externa só por ordem judicial, com prazo máximo de guarda (proposta de 5 anos; doc 07 seção 7.3).
5. Para a revinculação não travar, os comandantes aprovam **remotamente pelo próprio app**, com reconhecimento facial.

Hoje não existe cópia institucional: o servidor guarda cifrado que só os aparelhos da sala abrem (MLS), e um aparelho novo não lê épocas anteriores.

## Decisão (proposta)
**Leitura normal não muda:** membros da sala leem pelo MLS no próprio aparelho. O quórum vale para **todo acesso institucional** a dado cifrado.

1. **Chave de custódia por sala e período.** Para cada sala e período (ex.: mês, ou a cada troca de membros), os membros têm uma chave de dados de custódia. Cada mensagem e anexo leva uma cópia cifrada com ela, feita no cliente (core). O servidor guarda só cifrado (regra 3).
2. **Duas chaves de custódia, a mesma cópia.** A chave de cada sala e período é embrulhada (envelope) **duas vezes**:
   - para a **chave de recuperação**, em HSM online, com uma única operação permitida (seção "Quórum remoto"); e
   - para a **chave judicial**, em HSM **offline**, ativável só por **M de N** cartões de comandantes em cerimônia presencial.
   Nos dois, **M ≥ 3** (proposta: 3 de 5, para não travar quando um comandante falta).
3. **O que exige quórum:**
   - reenvio de histórico a aparelho revinculado (doc 07 seção 7.4): **quórum remoto**, chave de recuperação;
   - todo o resto, **quórum presencial** com cartões, chave judicial:
   - entrega por ordem judicial (doc 07 seção 7.3);
   - imagens de evidência de câmera (doc 07 seção 6.3);
   - registro de visualização (doc 05 seção 3.3);
   - localização em modo perda (doc 07 seção 7.4);
   - registro de acessos (doc 08 seção 2.3);
   - quebra de pseudonimato (doc 05 seção 3.4).
4. **Abertura com escopo.** O quórum libera **só** as chaves da sala e do período pedidos, nunca a chave-mestra. A decifração ocorre dentro do HSM ou em estação isolada sem rede. A saída é um pacote de evidência assinado (RF67) ou o reenvio cifrado para a chave do aparelho novo.
5. **Pedido antes do quórum.** Cada abertura tem pedido registrado com número de procedimento ou ordem judicial, finalidade, escopo e solicitante. Para autoridade externa, a ordem judicial é obrigatória.
6. **Registro.** Pedido, comandantes presentes, escopo e resultado vão para a trilha de auditoria (cadeia de hash), com aviso ao Auditor.
7. **Prazo.** Ao fim do prazo de guarda, destrói-se a chave embrulhada daquele período (crypto-shredding, ADR-04). Ordem de preservação (RF58) suspende a destruição.
8. **Primitivas.** Envelope com HPKE (RFC 9180) e quórum nativo do HSM. **Nenhuma divisão de segredo implementada à mão** (regra 1). A biblioteca HPKE do core será escolhida e fixada com hash na Fase 1; se a pilha do doc 02 não tiver, parar e perguntar.

## Quórum remoto (somente revinculação)
O risco da revinculação é baixo: o histórico volta para o **mesmo agente**, que já tinha acesso. Por isso ela pode usar aprovação remota. Abrir dado para **outra pessoa** (os demais itens) continua presencial.

1. **HSM de recuperação** online, com política que permite **só** recifrar chaves de sala e período para a chave de um aparelho que:
   - passou na atestação de hardware e tem série no inventário;
   - pertence ao **mesmo agente** do aparelho revogado;
   - é membro atual das salas pedidas;
   e só dentro da retenção de cada sala. Não exporta chave, não decifra conteúdo, não atende outro destino.
2. **Pedido** nasce de um cadastro **presencial** do aparelho novo (QR, atestação, termo vigente, triagem). Contém agente, aparelho novo, aparelho revogado, salas, períodos e o gestor que cadastrou.
3. **Aprovação no app do comandante:** o pedido aparece com esses dados; o comandante assina a aprovação daquele pedido (hash do pedido) com uma **chave de aprovação** gerada no StrongBox do aparelho dele, utilizável só após **biometria forte do sistema** (classe 3, com prova de vida) **e** PIN.
4. **Biometria do sistema, não modelo facial próprio:** o rosto fica no hardware seguro do aparelho e nunca vai ao servidor; não há gabarito do projeto (D-R8 continua recomendando não ter modelo próprio). Se o aparelho do comandante não tiver rosto classe 3, usa digital.
5. **Quórum:** o HSM só recifra com **3 assinaturas válidas de comandantes distintos**, dentro de uma janela (proposta: 24 h). Não contam o próprio agente nem o gestor que fez o cadastro.
6. **O comandante não vê conteúdo:** o histórico vai do HSM direto para o aparelho novo, cifrado para a chave dele.
7. **PIN de coação na aprovação:** recusa em silêncio e alerta a segurança.
8. **Aviso ao agente** pelos canais anteriores e ao Auditor; tudo na trilha de auditoria.
9. A chave de aprovação **não é parte da chave de custódia**: roubar o aparelho de um comandante não abre nada; revoga-se a chave de aprovação dele.

## Comandantes (detentores de parte)
- Cartão ou token próprio com PIN para a chave judicial, guardado em local separado; nunca em aparelho conectado.
- Aparelho de comandante com o mesmo endurecimento dos agentes, de preferência dedicado, com a chave de aprovação da revinculação.
- Credencial de segurança do grau mais alto tratado (doc 08 seção 5).
- Saída de um comandante: nova divisão no HSM (reemissão de cartões) e revogação do cartão antigo, registradas.
- Comandantes não precisam ser membros das salas; não leem nada no dia a dia.
- Cerimônia de ativação documentada, com duas testemunhas e gravação, como a da AC raiz (doc 04 seção 6.4).

## Alternativas rejeitadas
- **Sem custódia (MLS puro):** não atende os itens 2 a 4 do contexto.
- **Chave do Auditor ou de um administrador sozinho:** ponto único de abuso e de coação.
- **Divisão de segredo implementada no projeto:** viola a regra 1; usar o quórum do HSM.
- **Reconhecimento facial próprio para aprovar:** erro, falsificação e gabarito biométrico a guardar; usa-se a biometria do sistema.
- **Quórum remoto para todos os atos:** abrir dado para outra pessoa com aparelhos conectados aumenta demais a superfície.
- **Parte da chave no celular do comandante:** três celulares comprometidos abririam tudo.
- **Cópia em claro no servidor:** viola a regra 3.

## Consequências
- O expurgo por operação passa a destruir a **cópia de trabalho** (aparelhos e salas); a cópia de custódia segue o prazo de guarda. Atualizar RN03 e doc 06 seção 8 depois do parecer.
- Ajuda a resolver o conflito com a Lei 8.159 (R60, D-R5) e o art. 38 §4º do Decreto 7.845 (doc 08).
- **Novo alvo:** HSM + banco = todo o histórico do prazo. Mitigação: quórum, chaves por período, HSM judicial sem rede, HSM de recuperação com operação única, registro e auditoria externa (R68).
- **HSM de recuperação online** é mais exposto; o que ele pode fazer se limita a recifrar para aparelho atestado do mesmo agente (R81).
- **Engenharia social** com cadastro falso aprovado de boa-fé por três comandantes (R82); **aparelhos de comandante** viram alvo (R83).
- **Conluio de 3 comandantes** é o risco residual; mitigação: registro imutável, aviso ao Auditor, ordem judicial para uso externo, revisão por órgão de controle (R79).
- **Coação de comandantes:** M ≥ 3 e cartões em locais distintos elevam o custo; procedimento de coação para detentor (R80).
- Revinculação depende de quórum, mas remoto: minutos a horas, sem deslocamento.
- Custo no aparelho: uma cifragem simétrica a mais por mensagem.

## Revisar quando
O jurídico concluir D-R5 e D-R15, ou o grau máximo (D-R17) e o algoritmo de Estado (D-R18) forem definidos.
