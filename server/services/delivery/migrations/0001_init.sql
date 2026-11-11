-- Dono: delivery. room_id e device_id são uuids SEM chave estrangeira:
-- pertencem ao diretório, e FK entre serviços impediria a separação de bancos.
-- A limpeza por expurgo vem por chamada/evento, não por ON DELETE CASCADE.
CREATE TABLE outbox (
    msg_id     uuid PRIMARY KEY,
    room_id    uuid NOT NULL,
    envelope   bytea NOT NULL,
    expires_at timestamptz NOT NULL
);
CREATE INDEX outbox_room_idx ON outbox (room_id);
CREATE TABLE outbox_pending (
    msg_id    uuid REFERENCES outbox ON DELETE CASCADE,
    device_id uuid,
    PRIMARY KEY (msg_id, device_id)
);
