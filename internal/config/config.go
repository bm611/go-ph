package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	// API Configuration
	GeminiAPIKey string
	JinaAPIKey   string

	// Service URLs
	ProductHuntURL string
	JinaBaseURL    string

	// LLM Configuration
	GeminiModel  string
	Temperature  float32
	MaxProducts  int
	ResponseMIME string

	// Prompts
	ExtractionPrompt string
}

// Load creates a new Config with values from environment variables or defaults
func Load() *Config {
	return &Config{
		// API Keys
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
		JinaAPIKey:   getEnv("JINA_API_KEY", ""),

		// URLs
		ProductHuntURL: getEnv("PRODUCT_HUNT_URL", "https://producthunt.com"),
		JinaBaseURL:    getEnv("JINA_BASE_URL", "https://r.jina.ai/"),

		// LLM Settings
		GeminiModel:  getEnv("GEMINI_MODEL", "gemini-2.5-flash"),
		Temperature:  getEnvFloat("GEMINI_TEMPERATURE", 0.1),
		MaxProducts:  getEnvInt("MAX_PRODUCTS", 10),
		ResponseMIME: getEnv("RESPONSE_MIME_TYPE", "application/json"),

		// Prompt Template
		ExtractionPrompt: getEnv("EXTRACTION_PROMPT", getDefaultPrompt()),
	}
}

// Validate checks if all required configuration is present and valid
func (c *Config) Validate() error {
	if c.GeminiAPIKey == "" {
		return ErrMissingGeminiAPIKey
	}
	if c.JinaAPIKey == "" {
		return ErrMissingJinaAPIKey
	}

	// Validate temperature range
	if c.Temperature < 0.0 || c.Temperature > 1.0 {
		return ErrInvalidTemperature
	}

	// Validate max products
	if c.MaxProducts < 1 {
		return ErrInvalidMaxProducts
	}

	// Validate URLs are not empty
	if c.ProductHuntURL == "" {
		return errors.New("Product Hunt URL cannot be empty")
	}
	if c.JinaBaseURL == "" {
		return errors.New("Jina base URL cannot be empty")
	}

	// Validate model name
	if c.GeminiModel == "" {
		return errors.New("Gemini model name cannot be empty")
	}

	return nil
}

// String returns a formatted string representation of the configuration
func (c *Config) String() string {
	return fmt.Sprintf(`Configuration:
- Max Products: %d
- Temperature: %.2f
- Gemini Model: %s
- Product Hunt URL: %s
- Jina Base URL: %s
- Response MIME: %s`,
		c.MaxProducts,
		c.Temperature,
		c.GeminiModel,
		c.ProductHuntURL,
		c.JinaBaseURL,
		c.ResponseMIME,
	)
}

// GetAPIKeysStatus returns the status of API key configuration
func (c *Config) GetAPIKeysStatus() string {
	geminiStatus := "✓ Set"
	jinaStatus := "✓ Set"

	if c.GeminiAPIKey == "" {
		geminiStatus = "✗ Missing"
	}
	if c.JinaAPIKey == "" {
		jinaStatus = "✗ Missing"
	}

	return fmt.Sprintf(`API Keys Status:
- Gemini API Key: %s
- Jina API Key: %s`,
		geminiStatus,
		jinaStatus,
	)
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvFloat(key string, defaultValue float32) float32 {
	if value := os.Getenv(key); value != "" {
		if floatValue, err := strconv.ParseFloat(value, 32); err == nil {
			return float32(floatValue)
		}
	}
	return defaultValue
}

func getDefaultPrompt() string {
	return `extract the top %d products launched today from this text extract in json format
format:
{
	"rank": 1,
	"name": "Peek",
	"description": "AI personal finance coach that guides you through decisions",
	"product_url": "https://www.producthunt.com/posts/peek-1081",
	"image_url": "https://ph-files.imgix.net/0dcafea3-a3bd-40f6-bb99-49392faede45.png?auto=compress&codec=mozjpeg&cs=strip&auto=format&w=48&h=48&fit=crop&frame=1",
	"categories": [
		"Productivity",
		"Lifestyle",
		"Personal Finance"
	]
},

Here is the text extract:
%s`
}
