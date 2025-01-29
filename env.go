package main

import "os"

// GetString - fetch the String environment variable.
// Take the name of env variable, parse it and return it or a fallback.
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}
