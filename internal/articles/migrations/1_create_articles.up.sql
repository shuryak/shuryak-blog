CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    custom_id TEXT NOT NULL,
    author_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    thumbnail TEXT NOT NULL DEFAULT '',
    content JSON NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
)
