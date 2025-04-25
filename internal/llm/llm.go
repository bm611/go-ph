package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func GetGeminiResponse(prompt string) {
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

	for _, product := range products {
		fmt.Println(product.Name)
	}
}
