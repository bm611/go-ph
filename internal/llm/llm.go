package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"google.golang.org/genai"
)

func GetGeminiResponse(prompt string) ([]ProductRespType, error) {
	s := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))
	fmt.Println(s.Render("Formatting structured response with Gemini..."))
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:      genai.Ptr[float32](0.1),
			ResponseMIMEType: "application/json",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	var products []ProductRespType
	err = json.Unmarshal([]byte(result.Text()), &products)
	if err != nil {
		log.Fatal(err)
	}

	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true)

	fmt.Println(successStyle.Render("✓ Response formatted successfully"))

	return products, nil
}
