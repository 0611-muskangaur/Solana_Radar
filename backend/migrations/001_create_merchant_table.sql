CREATE TABLE merchants (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  wallet_address VARCHAR(100) NOT NULL UNIQUE,  -- Ensure wallet address is unique
  password VARCHAR(255) NOT NULL UNIQUE,         -- Ensure passwords are unique (less common)
  preferred_token VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
