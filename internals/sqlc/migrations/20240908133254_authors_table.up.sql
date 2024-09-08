-- up.sql

-- Create authors table
CREATE TABLE authors (
  id SERIAL PRIMARY KEY,  -- Changed to SERIAL for auto-incrementing INTEGER
  name TEXT NOT NULL,
  bio TEXT,
  email TEXT UNIQUE NOT NULL,
  date_of_birth DATE
);
