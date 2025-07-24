CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    image_url VARCHAR(100),
    purchase_price DECIMAL(12, 2) NOT NULL,
    selling_price DECIMAL(12, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0
)