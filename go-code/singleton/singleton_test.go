package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("GetInstance is not returning a singleton instance")
	}
}
func TestSingletonProperties(t *testing.T) {
	instance := GetInstance()

	// Test the Name property
	expectedName := "SingletonInstance"
	if instance.Name != expectedName {
		t.Errorf("Expected Name to be %s, but got %s", expectedName, instance.Name)
	}

	// Test the ID property
	expectedID := 12345
	if instance.ID != expectedID {
		t.Errorf("Expected ID to be %d, but got %d", expectedID, instance.ID)
	}
}

func TestConcurrentAccess(t *testing.T) {
	// Create multiple goroutines to access the singleton instance concurrently
	numGoroutines := 100
	done := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		go func() {
			instance := GetInstance()
			// verify the content of the singleton instance
			assert.Equal(t, "SingletonInstance", instance.Name)
			done <- struct{}{}
		}()
	}

	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}
