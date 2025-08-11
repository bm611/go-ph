package llm

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bm611/go-ph/internal/config"
	"github.com/bm611/go-ph/internal/spinner"
	"google.golang.org/genai"
)

func GetGeminiResponse(prompt string, cfg *config.Config) ([]ProductRespType, error) {
	// Start animated spinner for AI processing
	s := spinner.New(spinner.PulseStyle, "Formatting structured response with Gemini...")
	s.Start()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.GeminiAPIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		s.StopWithError("✗ Failed to initialize Gemini client")
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		cfg.GeminiModel,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:      genai.Ptr[float32](cfg.Temperature),
			ResponseMIMEType: cfg.ResponseMIME,
		},
	)
	if err != nil {
		s.StopWithError("✗ Failed to generate content with Gemini")
		log.Fatal(err)
	}
	var products []ProductRespType
	err = json.Unmarshal([]byte(result.Text()), &products)
	if err != nil {
		s.StopWithError("✗ Failed to parse Gemini response")
		log.Fatal(err)
	}

	// Stop spinner with success message
	s.StopWithMessage("✓ Response formatted successfully")

	return products, nil
}
