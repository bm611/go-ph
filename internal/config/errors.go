package config

import "errors"

var (
	// ErrMissingGeminiAPIKey is returned when the Gemini API key is not provided
	ErrMissingGeminiAPIKey = errors.New("GEMINI_API_KEY environment variable is required")

	// ErrMissingJinaAPIKey is returned when the Jina API key is not provided
	ErrMissingJinaAPIKey = errors.New("JINA_API_KEY environment variable is required")

	// ErrInvalidConfiguration is returned when the configuration is invalid
	ErrInvalidConfiguration = errors.New("invalid configuration provided")

	// ErrInvalidTemperature is returned when temperature is outside valid range (0.0-1.0)
	ErrInvalidTemperature = errors.New("temperature must be between 0.0 and 1.0")

	// ErrInvalidMaxProducts is returned when max products is less than 1
	ErrInvalidMaxProducts = errors.New("max products must be greater than 0")
)
