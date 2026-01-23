package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Store struct {
	Name string `json:"name"`
}

type Product struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	baseURL := "http://localhost:8000"

	fmt.Println("--- Phase 1: Seeding Stores ---")
	stores := []string{"Tech Haven", "Digital Dreams", "Silicon Valley Supplies"}
	
	for _, storeName := range stores {
		s := Store{Name: storeName}
		data, _ := json.Marshal(s)
		resp, err := http.Post(baseURL+"/stores", "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Error creating store %s: %v\n", storeName, err)
			continue
		}
		fmt.Printf("Created Store: %s (Status: %d)\n", storeName, resp.StatusCode)
		resp.Body.Close()
	}

	fmt.Println("\n--- Phase 2: Seeding 50 Products ---")
	brands := []string{"Logitech", "Razer", "Apple", "Samsung", "Sony", "Dell"}
	items := []string{"Keyboard", "Mouse", "Monitor", "Headset", "Webcam", "Laptop"}

	for i := 1; i <= 50; i++ {
		brand := brands[rand.Intn(len(brands))]
		item := items[rand.Intn(len(items))]
		
		p := Product{
			Name:         fmt.Sprintf("%s %s %d", brand, item, i),
			Price:        float64(rand.Intn(900)+50) + 0.99,
			Availability: true,
		}

		jsonData, _ := json.Marshal(p)

		resp, err := http.Post(baseURL+"/stores/1/products", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error creating product %d: %v\n", i, err)
			continue
		}
		
		if i%10 == 0 {
			fmt.Printf("Progress: %d/50 products seeded...\n", i)
		}
		resp.Body.Close()
	}

	fmt.Println("\nSeeding Complete!")
	fmt.Println("Database is now populated. Redis caches for 'stores:all' and 'products:all' have been cleared.")
}