package scraper

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bm611/go-ph/internal/config"
	"github.com/bm611/go-ph/internal/spinner"
)

func GetPageContent(site string, cfg *config.Config) (string, error) {
	// Start animated spinner
	s := spinner.New(spinner.DotsStyle, "Extracting content from producthunt.com...")
	s.Start()

	fullURL := cfg.JinaBaseURL + site
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		s.StopWithError("✗ Failed to create request")
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+cfg.JinaAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.StopWithError("✗ Failed to fetch content")
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.StopWithError("✗ Failed to read response")
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Stop spinner with success message
	s.StopWithMessage("✓ Content extracted successfully")

	return string(body), nil
}
