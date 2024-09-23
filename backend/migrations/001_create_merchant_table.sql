--SQL script for creating merchants table

CREATE TABLE merchants (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  wallet_address VARCHAR(100) NOT NULL,
  preferred_token VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
