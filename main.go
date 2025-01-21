package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
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

type SimulatorInfo struct {
	ID          string   `json:"simulator_id"`
	Port        string   `json:"port"`
	SensorTypes []string `json:"sensor_types"`
	StartTime   string   `json:"start_time"`
}

var (
	temperature float64
	humidity    float64
	vibration   float64
	simulatorID string
	startTime   time.Time
)

func init() {
	// Initialize simulator ID from environment variable or use default
	simulatorID = os.Getenv("SIMULATOR_ID")
	if simulatorID == "" {
		simulatorID = "simulator-1"
	}
	startTime = time.Now()
}

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

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	info := SimulatorInfo{
		ID:          simulatorID,
		Port:        getPort(),
		SensorTypes: []string{"temperature", "humidity", "vibration"},
		StartTime:   startTime.Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(info)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data := SensorData{
		Value: fmt.Sprintf("%.2f", temperature),
		Unit:  "°C",
	}
	json.NewEncoder(w).Encode(data)
}

func humidityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data := SensorData{
		Value: fmt.Sprintf("%.2f", humidity),
		Unit:  "%",
	}
	json.NewEncoder(w).Encode(data)
}

func vibrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data := SensorData{
		Value: fmt.Sprintf("%.2f", vibration),
		Unit:  "mm/s",
	}
	json.NewEncoder(w).Encode(data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	go updateSensorData()

	// Set up API endpoints
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/humidity", humidityHandler)
	http.HandleFunc("/vibration", vibrationHandler)

	port := getPort()
	log.Printf("Server starting on port %s with ID: %s\n", port, simulatorID)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
