
-- PostgreSQL database init script

CREATE DATABASE bridge;

\c bridge

-- Create users table

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    password_hash TEXT NOT NULL,
    balance BIGINT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

