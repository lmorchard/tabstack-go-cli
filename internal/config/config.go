package config

// Config holds application configuration.
type Config struct {
	// Core settings
	Verbose bool
	Debug   bool
	LogJSON bool

	// Tabstack API client settings
	APIKey  string
	BaseURL string
}
