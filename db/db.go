package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func InitDB() {
	connStr := os.Getenv("DB_URL")
	var err error
	
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err == nil && DB.Ping() == nil {
			break
		}
		log.Printf("DB not ready, retrying... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	TableSQL := `
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
        price DECIMAL NOT NULL,
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
        quantity INTEGER NOT NULL
    );

    CREATE TABLE IF NOT EXISTS payments (
        id SERIAL PRIMARY KEY,
        order_id INTEGER REFERENCES orders(id),
        user_id INTEGER REFERENCES users(id),
        amount DECIMAL NOT NULL,
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
    `
	_, err = DB.Exec(TableSQL)
	if err != nil {
		log.Fatal("Failed to build tables: ", err)
	}
	log.Println("DB and Redis connected.")
}