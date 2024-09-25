CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  merchant_id INT REFERENCES merchants(id),
  amount DECIMAL(18, 8) NOT NULL,
  currency VARCHAR(10),
  payment_status VARCHAR(50) DEFAULT 'pending',
  transaction_hash VARCHAR(255) UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  payment_request_id INT REFERENCES payments(id),
  amount DECIMAL(18, 8) NOT NULL,
  status VARCHAR(50) DEFAULT 'pending',
  blockchain VARCHAR(50) NOT NULL,
  date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  transaction_hash VARCHAR(255) UNIQUE
);
