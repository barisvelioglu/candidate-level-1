# Sensor Simulator API

This is a simple sensor simulator application that generates realistic sensor data for temperature, humidity, and vibration measurements. The application provides REST API endpoints to access the simulated sensor data.

## Features

- Simulates three different types of sensor data:
  - Temperature (18-30°C)
  - Humidity (30-70%)
  - Vibration (0-5 mm/s)
- Updates sensor values every second
- Provides REST API endpoints for each sensor
- Includes health check endpoint
- Configurable simulator ID and port via environment variables
- Support for running multiple simulator instances
- Simulates realistic sensor failures:
  - 10% chance of individual sensor failure (null values)
  - 5% chance of system-wide failure (500 errors)
- Docker support for easy deployment

## Environment Variables

- `SIMULATOR_ID`: Unique identifier for the simulator instance (default: "simulator-1")
- `PORT`: Port number for the simulator API (default: "8080")

## API Endpoints

### Simulator Information
```
GET /info
```
Response example:
```json
{
  "simulator_id": "simulator-1",
  "port": "8080",
  "sensor_types": ["temperature", "humidity", "vibration"],
  "start_time": "2024-03-14T12:34:56Z"
}
```

### Sensor Data Endpoints

1. **Temperature**
   ```
   POST /temperature
   ```
   Success response example:
   ```json
   {
     "value": "23.45",
     "unit": "°C"
   }
   ```
   Sensor failure response example:
   ```json
   {
     "value": null,
     "unit": "°C"
   }
   ```

2. **Humidity**
   ```
   POST /humidity
   ```
   Success response example:
   ```json
   {
     "value": "45.67",
     "unit": "%"
   }
   ```
   Sensor failure response example:
   ```json
   {
     "value": null,
     "unit": "%"
   }
   ```

3. **Vibration**
   ```
   POST /vibration
   ```
   Success response example:
   ```json
   {
     "value": "2.34",
     "unit": "mm/s"
   }
   ```
   Sensor failure response example:
   ```json
   {
     "value": null,
     "unit": "mm/s"
   }
   ```

### Health Check

```
GET /health
```
Success response example:
```json
{
  "status": "healthy",
  "timestamp": "2024-03-14T12:34:56Z"
}
```

Note: All endpoints may return a 500 Internal Server Error during system-wide failures (5% probability).

## Running the Application

### Using Docker Compose (Recommended)

1. Build and start the containers (this will start two simulator instances):
   ```bash
   docker-compose up --build
   ```

2. The APIs will be available at:
   - Simulator 1: `http://localhost:8080`
   - Simulator 2: `http://localhost:8081`

### Running Locally

1. Make sure you have Go 1.21 or later installed
2. Set environment variables (optional):
   ```bash
   export SIMULATOR_ID=my-simulator
   export PORT=8082
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Testing the API

You can use curl to test the endpoints:

```bash
# Get simulator information
curl http://localhost:8080/info

# Get temperature data (may return null or 500 error)
curl -X POST http://localhost:8080/temperature

# Get humidity data (may return null or 500 error)
curl -X POST http://localhost:8080/humidity

# Get vibration data (may return null or 500 error)
curl -X POST http://localhost:8080/vibration

# Check health status (may return 500 error)
curl http://localhost:8080/health

# Test second simulator instance
curl http://localhost:8081/info
``` 