-- Ensure UUID extension is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table definition with additional fields
CREATE TABLE authors (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT      NOT NULL,
  bio  TEXT,
  email TEXT     UNIQUE NOT NULL,
  date_of_birth DATE
);
