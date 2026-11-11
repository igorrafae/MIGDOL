-- Dono: blobs. room_id sem FK (pertence ao diretório).
CREATE TABLE blobs (
    object_key text PRIMARY KEY,            -- aleatório
    room_id    uuid NOT NULL,
    size_bytes bigint NOT NULL,
    expires_at timestamptz NOT NULL
);
CREATE INDEX blobs_room_idx ON blobs (room_id);
