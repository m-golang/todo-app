package config

import (
	"fmt"
	"os"
)

// LoadEnvVars retrieves the values of environment variables specified in the 'keys' slice.
// It returns a slice of values corresponding to the environment variable keys.
// If any key is missing, it returns an error indicating which variable is missing.
func LoadEnvVars(keys []string) ([]string, error) {
	var values []string

	// Iterate over the list of keys and fetch the corresponding environment variable value
	for _, key := range keys {
		value := os.Getenv(key)
		// If the environment variable is not found or is empty, return an error
		if value == "" {
			return nil, fmt.Errorf("config: missing environment variable: %s", key)
		}
		values = append(values, value) // Append the retrieved value to the 'values' slice
	}

	// Return the list of environment variable values
	return values, nil
}
