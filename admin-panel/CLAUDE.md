# admin-panel/ — Painel do gestor

Rede interna, estação gerenciada. O gestor cadastra e revoga usuários e aparelhos; não lê conteúdo.
Fase: 1 (painel mínimo). Pilha ainda não decidida: registrar em ADR antes de iniciar.

Nota (ADR-0008 e ADR-0009, propostas; depois da v1): o painel pode mostrar só indicadores de nível 1a,
agregados dos metadados que o servidor já tem (sem nome, sala, operação identificável nem posição;
contagens arredondadas ou em faixas; lista fechada no instituicao.yaml; datas por semana ou mês). Nível 1b
não existe (GD-15: a regra 3 não muda). Agregados derivados e nominais (nível 2) nunca aparecem no painel:
só no app do analista, por MLS, e só de salas que concederam entrada no escopo. Acesso por papel com atributos e negação por padrão (ADR-0009).

No navegador (01g-G24): nenhum dado de cadastro em armazenamento do navegador (por exemplo, localStorage,
sessionStorage, IndexedDB, cache de service worker ou cookie além da sessão) e nenhum recurso externo
(script, fonte, folha de estilo ou imagem de fora da rede interna). Telas de protótipo: PW01 a PW12 (doc 11).
