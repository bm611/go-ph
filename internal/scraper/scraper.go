package scraper

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func GetPageContent(site string) (string, error) {
	loadingStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("105")).
		Bold(true)
	fmt.Println(loadingStyle.Render("Extracting content from producthunt.com....."))
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

	// Create a styled success message with a checkmark
	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true)

	fmt.Println(successStyle.Render("✓ Content extracted successfully"))

	return string(body), nil
}
