-- Dono: audit. O papel da aplicação não tem UPDATE nem DELETE nesta tabela.
CREATE TABLE audit_log (
    seq       bigserial PRIMARY KEY,
    at        timestamptz NOT NULL DEFAULT now(),
    actor     uuid,
    event     text NOT NULL,
    room_ref  bytea,                        -- hash do room_id
    detail    jsonb NOT NULL,
    prev_hash bytea NOT NULL,
    this_hash bytea NOT NULL
);
