
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
CREATE INDEX users_username_idx ON users(username);

-- Create transactions table

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    recipient_user_id SERIAL NOT NULL,
    sender_user_id SERIAL,
    is_mined BOOLEAN NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    tx_description TEXT,
    CONSTRAINT fk_user_id
        FOREIGN KEY (recipient_user_id)
            REFERENCES users (id)
);

CREATE INDEX tx_id_idx ON transactions(id);
CREATE INDEX tx_recipient_user_id_idx ON transactions(recipient_user_id);
CREATE INDEX tx_sender_user_id_idx ON transactions(sender_user_id);
