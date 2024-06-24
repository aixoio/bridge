
-- PostgreSQL database init script

CREATE DATABASE bridge;

\c bridge

-- Create users table

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    password_hash TEXT NOT NULL,
    balance BIGINT DEFAULT 0 NOT NULL,
    public_account BOOLEAN DEFAULT TRUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Add indexes to the users table

CREATE INDEX users_id_idx ON users(id);
CREATE INDEX user_username_idx ON users(username);
