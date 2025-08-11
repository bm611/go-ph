package service

import (
	"fmt"

	"github.com/bm611/go-ph/internal/config"
	"github.com/bm611/go-ph/internal/llm"
	"github.com/bm611/go-ph/internal/scraper"
)

// ProductHuntService handles the core business logic for fetching Product Hunt data
type ProductHuntService interface {
	GetTodaysProducts() ([]llm.ProductRespType, error)
}

// productHuntService implements the ProductHuntService interface
type productHuntService struct {
	config *config.Config
}

// NewProductHuntService creates a new instance of ProductHuntService
func NewProductHuntService(cfg *config.Config) ProductHuntService {
	return &productHuntService{
		config: cfg,
	}
}

// GetTodaysProducts fetches and processes today's Product Hunt launches
func (s *productHuntService) GetTodaysProducts() ([]llm.ProductRespType, error) {
	// Fetch raw content from Product Hunt via Jina AI
	content, err := scraper.GetPageContent(s.config.ProductHuntURL, s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page content: %w", err)
	}

	// Generate the prompt with the fetched content
	prompt := fmt.Sprintf(s.config.ExtractionPrompt, s.config.MaxProducts, content)

	// Use Gemini to extract structured data
	products, err := llm.GetGeminiResponse(prompt, s.config)
	if err != nil {
		return nil, fmt.Errorf("failed to process content with Gemini: %w", err)
	}

	// Validate and limit results
	if len(products) > s.config.MaxProducts {
		products = products[:s.config.MaxProducts]
	}

	return products, nil
}

// ValidateService checks if the service is properly configured
func (s *productHuntService) ValidateService() error {
	return s.config.Validate()
}
