
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Ensure UUID extension is available
    
CREATE TABLE authors (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT      NOT NULL,
  bio  TEXT
);
