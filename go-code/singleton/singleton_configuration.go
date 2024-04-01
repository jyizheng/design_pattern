package main

import (
	"fmt"
	"sync"
)

// Config represents application configuration settings
type Config struct {
	DatabaseURL  string
	Port         int
	FeatureFlags map[string]bool
}

// singleton instance variable
var instance *Config

// mutex for making instance creation thread-safe
var once sync.Once

// GetConfig returns the singleton instance of the application configuration
func GetConfig() *Config {
	once.Do(func() {
		// Imagine these settings come from a file, environment variables, or another source
		instance = &Config{
			DatabaseURL: "localhost:5432/mydb",
			Port:        8080,
			FeatureFlags: map[string]bool{
				"newFeatureEnabled": true,
			},
		}
	})
	return instance
}

func main() {
	// Accessing the singleton Config instance from different parts of the application
	config := GetConfig()
	fmt.Println("Database URL:", config.DatabaseURL)
	fmt.Println("Application Port:", config.Port)
	fmt.Println("Is new feature enabled?", config.FeatureFlags["newFeatureEnabled"])

	// Any subsequent calls to GetConfig will return the same instance
	anotherConfigReference := GetConfig()
	if config == anotherConfigReference {
		fmt.Println("config and anotherConfigReference point to the same instance.")
	}
}


