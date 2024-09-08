-- Table definition with an auto-incrementing integer primary key
CREATE TABLE authors (
  id SERIAL PRIMARY KEY,   -- Use SERIAL for auto-incrementing integer primary key
  name TEXT NOT NULL,
  bio TEXT,
  email TEXT UNIQUE NOT NULL,
  date_of_birth DATE
);
