package cmd

import (
	"fmt"

	"github.com/bm611/go-ph/internal/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current configuration and validate settings",
	Long: `The config command displays the current configuration settings
and validates that all required environment variables and settings are properly configured.

This is useful for troubleshooting configuration issues and verifying your setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg := config.Load()

		fmt.Println("=== go-ph Configuration ===")
		fmt.Println()

		// Display API Keys Status
		fmt.Println(cfg.GetAPIKeysStatus())
		fmt.Println()

		// Display full configuration
		fmt.Println(cfg.String())
		fmt.Println()

		// Validate configuration
		fmt.Println("=== Configuration Validation ===")
		if err := cfg.Validate(); err != nil {
			fmt.Printf("❌ Configuration validation failed: %v\n", err)
			fmt.Println()
			fmt.Println("Required environment variables:")
			fmt.Println("- GEMINI_API_KEY: Your Google Gemini API key")
			fmt.Println("- JINA_API_KEY: Your Jina AI API key")
			fmt.Println()
			fmt.Println("Optional environment variables:")
			fmt.Println("- PRODUCT_HUNT_URL: Product Hunt URL (default: https://producthunt.com)")
			fmt.Println("- JINA_BASE_URL: Jina AI base URL (default: https://r.jina.ai/)")
			fmt.Println("- GEMINI_MODEL: Gemini model name (default: gemini-2.5-flash)")
			fmt.Println("- GEMINI_TEMPERATURE: AI temperature 0.0-1.0 (default: 0.1)")
			fmt.Println("- MAX_PRODUCTS: Maximum products to fetch (default: 10)")
			fmt.Println("- RESPONSE_MIME_TYPE: Response format (default: application/json)")
			return
		}

		fmt.Println("✅ Configuration is valid and ready to use!")
		fmt.Println()
		fmt.Println("You can now run 'go-ph' to fetch today's Product Hunt launches.")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
