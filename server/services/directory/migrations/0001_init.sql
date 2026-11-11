-- Dono: directory. Sem conteúdo em claro. Sem FK para tabelas de outros serviços.
CREATE TABLE users (
    user_id      uuid PRIMARY KEY,
    registration text NOT NULL UNIQUE,
    codename     text NOT NULL UNIQUE,
    role         text NOT NULL,             -- gestor | operador | auditor
    created_at   timestamptz NOT NULL DEFAULT now(),
    revoked_at   timestamptz
);
CREATE TABLE devices (
    device_id    uuid PRIMARY KEY,
    user_id      uuid NOT NULL REFERENCES users,
    cert_fpr     bytea NOT NULL UNIQUE,
    attestation  jsonb NOT NULL,
    lease_until  timestamptz,
    status       text NOT NULL,             -- active | suspended | revoked
    last_seen_at timestamptz
);
CREATE TABLE rooms (
    room_id        uuid PRIMARY KEY,        -- aleatório, sem significado
    admin_keys     bytea[] NOT NULL,
    retention_secs integer NOT NULL,
    expunge_seq    bigint NOT NULL DEFAULT 0,
    dual_custody   boolean NOT NULL DEFAULT false,
    created_at     timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE memberships (
    room_id   uuid REFERENCES rooms ON DELETE CASCADE,
    device_id uuid REFERENCES devices,
    PRIMARY KEY (room_id, device_id)
);
CREATE TABLE key_packages (
    id         uuid PRIMARY KEY,
    device_id  uuid NOT NULL REFERENCES devices,
    package    bytea NOT NULL,
    claimed_at timestamptz
);
