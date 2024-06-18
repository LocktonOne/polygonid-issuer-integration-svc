-- +migrate Up

CREATE TABLE dids (
   did text primary key,
   address text not null unique
);

-- +migrate Down
