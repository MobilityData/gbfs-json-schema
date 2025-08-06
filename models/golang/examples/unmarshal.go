package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/gbfs"
)

func ExampleUnmarshal() {
	endpoint := "https://gbfs.api.ridedott.com/public/v2/mannheim/gbfs.json"

	fmt.Println("=== GBFS Codec Example ===")

	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("Failed to fetch GBFS:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch GBFS: %s", resp.Status)
	}

	var gbfsFeed gbfs.Gbfs
	if err := json.NewDecoder(resp.Body).Decode(&gbfsFeed); err != nil {
		log.Fatal("Failed to unmarshal GBFS:", err)
	}

	fmt.Printf("Successfully unmarshaled GBFS feed (version: %s)\n", gbfsFeed.Version)
}

func main() {
	ExampleUnmarshal()
}
