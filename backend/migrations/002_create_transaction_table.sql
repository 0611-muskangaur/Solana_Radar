-- SQL script for creating transactions table

CREATE TABLE payments (    --specifies that a new table called payments will be created.
  id SERIAL PRIMARY KEY,
  merchant_id INT REFERENCES merchants(id),
  amount DECIMAL(18, 8) NOT NULL,
  currency VARCHAR(10),
  payment_status VARCHAR(50) DEFAULT 'pending',
  transaction_hash VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
