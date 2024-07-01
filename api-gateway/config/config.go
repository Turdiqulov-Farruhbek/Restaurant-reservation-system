package config

import (
    "os"
)

var (
    AuthServiceURL       = getEnv("AUTH_SERVICE_URL", "http://localhost:8001")
    ReservationServiceURL = getEnv("RESERVATION_SERVICE_URL", "http://localhost:8002")
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
