package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type SensorData struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

var (
	temperature float64
	humidity    float64
	vibration   float64
)

func updateSensorData() {
	for {
		// Temperature: 18-30°C
		temperature = 18 + rand.Float64()*12
		// Humidity: 30-70%
		humidity = 30 + rand.Float64()*40
		// Vibration: 0-5 mm/s
		vibration = rand.Float64() * 5

		time.Sleep(time.Second)
	}
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	data := SensorData{
		Value: fmt.Sprintf("%.2f", temperature),
		Unit:  "°C",
	}
	json.NewEncoder(w).Encode(data)
}

func humidityHandler(w http.ResponseWriter, r *http.Request) {
	data := SensorData{
		Value: fmt.Sprintf("%.2f", humidity),
		Unit:  "%",
	}
	json.NewEncoder(w).Encode(data)
}

func vibrationHandler(w http.ResponseWriter, r *http.Request) {
	data := SensorData{
		Value: fmt.Sprintf("%.2f", vibration),
		Unit:  "mm/s",
	}
	json.NewEncoder(w).Encode(data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Start sensor data simulation in background
	go updateSensorData()

	// Set up API endpoints
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/humidity", humidityHandler)
	http.HandleFunc("/vibration", vibrationHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
