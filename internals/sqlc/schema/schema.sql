CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,  -- Changed to BIGSERIAL for PostgreSQL auto-increment
  name TEXT      NOT NULL,
  bio  TEXT
);
