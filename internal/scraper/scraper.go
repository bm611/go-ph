package scraper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetPageContent(site string) (string, error) {
	BaseURL := "https://r.jina.ai/"
	fullURL := BaseURL + site
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	apiKey := os.Getenv("JINA_API_KEY")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return string(body), nil
}
