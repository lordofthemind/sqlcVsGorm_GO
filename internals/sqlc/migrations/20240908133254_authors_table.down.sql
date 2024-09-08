-- down.sql

-- Drop authors table
DROP TABLE IF EXISTS authors;

-- Optionally, you can drop the UUID extension if no other tables or processes depend on it
-- But typically, UUID extension is used in multiple tables, so you might want to keep it
-- DROP EXTENSION IF EXISTS "uuid-ossp";
