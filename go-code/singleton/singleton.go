package singleton

import "sync"

type singleton struct {
	Name string // Name is the name of the singleton instance
	ID   int    // ID is the unique identifier of the singleton instance
}

var instance *singleton
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			Name: "SingletonInstance", // Initialize the Name field here
			ID:   12345,               // Initialize the ID field here
		}
	})
	return instance
}
