package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	brands := []string{"Logitech", "Razer", "Apple", "Samsung", "Sony", "Dell"}
	items := []string{"Keyboard", "Mouse", "Monitor", "Headset", "Webcam", "Laptop"}

	fmt.Println("Starting seeding process...")

	for i := 1; i <= 50; i++ {
		brand := brands[rand.Intn(len(brands))]
		item := items[rand.Intn(len(items))]
		
		p := Product{
			Name:         fmt.Sprintf("%s %s %d", brand, item, i),
			Price:        float64(rand.Intn(900) + 50) + 0.99,
			Availability: true,
		}

		jsonData, _ := json.Marshal(p)
		resp, err := http.Post("http://localhost:8000/stores/1/products", "application/json", bytes.NewBuffer(jsonData))
		
		if err != nil {
			fmt.Printf("Error seeding product %d: %v\n", i, err)
			continue
		}
		
		if resp.StatusCode == http.StatusCreated {
			fmt.Printf("Successfully seeded: %s\n", p.Name)
		}
		resp.Body.Close()
	}

	fmt.Println("\nSeeding complete! 50 products added.")
	fmt.Println("Redis cache 'products:all' has been invalidated and will refresh on next GET.")
}