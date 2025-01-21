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
	Value *string `json:"value"`
	Unit  string  `json:"unit"`
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
	temperature   *float64
	humidity      *float64
	vibration     *float64
	simulatorID   string
	startTime     time.Time
	isSystemError bool
)

const (
	sensorFailureProb = 0.1  // 10% chance of sensor failure
	systemFailureProb = 0.05 // 5% chance of system failure
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
		// Simulate system-wide failure
		isSystemError = rand.Float64() < systemFailureProb

		// Update temperature with possible failure
		if rand.Float64() >= sensorFailureProb {
			temp := 18 + rand.Float64()*12
			temperature = &temp
		} else {
			temperature = nil
		}

		// Update humidity with possible failure
		if rand.Float64() >= sensorFailureProb {
			hum := 30 + rand.Float64()*40
			humidity = &hum
		} else {
			humidity = nil
		}

		// Update vibration with possible failure
		if rand.Float64() >= sensorFailureProb {
			vib := rand.Float64() * 5
			vibration = &vib
		} else {
			vibration = nil
		}

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

func handleSystemError(w http.ResponseWriter) bool {
	if isSystemError {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}
	return false
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if handleSystemError(w) {
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

	if handleSystemError(w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var value *string
	if temperature != nil {
		strVal := fmt.Sprintf("%.2f", *temperature)
		value = &strVal
	}
	data := SensorData{
		Value: value,
		Unit:  "°C",
	}
	json.NewEncoder(w).Encode(data)
}

func humidityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if handleSystemError(w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var value *string
	if humidity != nil {
		strVal := fmt.Sprintf("%.2f", *humidity)
		value = &strVal
	}
	data := SensorData{
		Value: value,
		Unit:  "%",
	}
	json.NewEncoder(w).Encode(data)
}

func vibrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if handleSystemError(w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var value *string
	if vibration != nil {
		strVal := fmt.Sprintf("%.2f", *vibration)
		value = &strVal
	}
	data := SensorData{
		Value: value,
		Unit:  "mm/s",
	}
	json.NewEncoder(w).Encode(data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if handleSystemError(w) {
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
