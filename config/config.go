package config

import (
    "os"
)

// GetEnv reads an environment variable or returns a fallback value.
func GetEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
