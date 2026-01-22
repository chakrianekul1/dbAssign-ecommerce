CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone TEXT,
    address TEXT
);
CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT DEFAULT 'active'
);
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    store_id INTEGER REFERENCES stores(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    availability BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    status TEXT DEFAULT 'created',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id),
    quantity INTEGER NOT NULL DEFAULT 1
);
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    amount NUMERIC NOT NULL,
    status TEXT DEFAULT 'pending',
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    product_id INTEGER REFERENCES products(id),
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP,
    status TEXT DEFAULT 'active'
);