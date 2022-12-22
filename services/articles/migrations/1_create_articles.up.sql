CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    custom_id TEXT NOT NULL,
    author_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    content JSON NOT NULL
);
