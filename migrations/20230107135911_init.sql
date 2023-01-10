-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS archives
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    size       BIGINT                NOT NULL,
    name       VARCHAR(64)           NOT NULL,
    original   VARCHAR(128)          NOT NULL,
    unarchived BOOLEAN DEFAULT FALSE NOT NULL,
    created    TIMESTAMP             NOT NULL
);

CREATE TABLE IF NOT EXISTS chunks
(
    archive         BIGINT REFERENCES archives (id),
    number          INT                   NOT NULL,
    start_byte      BIGINT                NOT NULL,
    next_chunk_byte BIGINT                NOT NULL,
    uploaded        BOOLEAN DEFAULT FALSE NOT NULL,
    uploaded_time   TIMESTAMP             NULL,

    PRIMARY KEY (archive, number)
);
CREATE INDEX IF NOT EXISTS idx_chunks_archive_uploaded ON chunks (archive, uploaded);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chunks CASCADE;
DROP TABLE IF EXISTS archives CASCADE;
-- +goose StatementEnd
