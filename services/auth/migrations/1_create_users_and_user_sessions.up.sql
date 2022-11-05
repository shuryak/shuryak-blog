CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY NOT NULL,
    username        TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_sessions (
    id            BIGSERIAL PRIMARY KEY NOT NULL,
    user_id       BIGINT REFERENCES users (id),
    refresh_token TEXT NOT NULL,
    expires_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
)
