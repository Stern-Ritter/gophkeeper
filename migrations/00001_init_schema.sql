-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS gophkeeper;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS gophkeeper.users
(
    id       UUID DEFAULT uuid_generate_v4(),
    login    VARCHAR(30)  NOT NULL,
    password VARCHAR(256) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY (id),
    CONSTRAINT users_login_unique UNIQUE (login)
);

CREATE TYPE gophkeeper.data_type AS ENUM ('ACCOUNT','TEXT', 'CARD');

CREATE TABLE IF NOT EXISTS gophkeeper.data
(
    id      UUID          DEFAULT uuid_generate_v4(),
    user_id UUID                 NOT NULL,
    data    BYTEA                NOT NULL,
    type    gophkeeper.DATA_TYPE NOT NULL,
    comment VARCHAR(1024) DEFAULT '',
    CONSTRAINT pk_data PRIMARY KEY (id),
    CONSTRAINT data_to_users_fk
        FOREIGN KEY (user_id) REFERENCES gophkeeper.users (id)
);

CREATE TABLE IF NOT EXISTS gophkeeper.files
(
    id      UUID          DEFAULT uuid_generate_v4(),
    user_id UUID          NOT NULL,
    name    VARCHAR(256)  NOT NULL,
    size    BIGINT        NOT NULL,
    path    VARCHAR(1024) NOT NULL,
    comment VARCHAR(1024) DEFAULT '',
    CONSTRAINT pk_files PRIMARY KEY (id),
    CONSTRAINT files_to_users_fk
        FOREIGN KEY (user_id) REFERENCES gophkeeper.users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gophkeeper.files;
DROP TABLE IF EXISTS gophkeeper.data;
DROP TYPE IF EXISTS gophkeeper.data_type;
DROP TABLE IF EXISTS gophkeeper.users;
-- +goose StatementEnd