CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    users_id INT NOT NULL,
    type VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(12, 2) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (users_id) REFERENCES users (id) ON DELETE CASCADE
);