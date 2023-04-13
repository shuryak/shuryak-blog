DROP TABLE IF EXISTS user_sessions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_sessions (
id            UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
user_id       BIGINT REFERENCES users (id),
expires_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL,
updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);