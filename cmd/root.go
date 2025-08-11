package cmd

import (
	"fmt"
	"os"

	"github.com/bm611/go-ph/internal/config"
	"github.com/bm611/go-ph/internal/service"
	"github.com/bm611/go-ph/internal/spinner"
	"github.com/bm611/go-ph/internal/ui"
	"github.com/spf13/cobra"
)

// Command line flags
var (
	maxProducts    int
	temperature    float64
	geminiModel    string
	productHuntURL string
	verbose        bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ph",
	Short: "Display products launched today on Product Hunt",
	Long: `go-ph is a CLI tool that fetches and displays products
that were launched today on producthunt.com.

It allows you to quickly see the latest product launches without
having to visit the website, helping you stay updated on the
newest tech products and startups.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg := config.Load()

		// Override with CLI flags if provided
		if maxProducts > 0 {
			cfg.MaxProducts = maxProducts
		}
		if temperature >= 0 {
			cfg.Temperature = float32(temperature)
		}
		if geminiModel != "" {
			cfg.GeminiModel = geminiModel
		}
		if productHuntURL != "" {
			cfg.ProductHuntURL = productHuntURL
		}

		// Validate configuration
		if err := cfg.Validate(); err != nil {
			fmt.Printf("Configuration error: %v\n", err)
			fmt.Println("\nPlease ensure you have set the required environment variables:")
			fmt.Println("- GEMINI_API_KEY: Your Google Gemini API key")
			fmt.Println("- JINA_API_KEY: Your Jina AI API key")
			os.Exit(1)
		}

		if verbose {
			fmt.Println(cfg.String())
			fmt.Println()
			fmt.Println(cfg.GetAPIKeysStatus())
			fmt.Println()
		}

		// Start overall progress spinner
		overallSpinner := spinner.New(spinner.BouncingBar, "Fetching and processing Product Hunt data...")
		overallSpinner.Start()

		// Create service
		productHuntService := service.NewProductHuntService(cfg)

		// Fetch today's products
		products, err := productHuntService.GetTodaysProducts()
		if err != nil {
			overallSpinner.StopWithError("✗ Failed to fetch products")
			fmt.Printf("Error fetching products: %v\n", err)
			os.Exit(1)
		}

		// Stop the overall spinner
		overallSpinner.StopWithMessage("✓ All data processed successfully!")
		fmt.Println() // Add a blank line for better formatting

		// Display results
		if len(products) == 0 {
			fmt.Println("No products found for today.")
			return
		}

		tableString := ui.RenderTable(products)
		fmt.Println(tableString)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Command line flags
	rootCmd.Flags().IntVarP(&maxProducts, "max-products", "m", 0, "Maximum number of products to fetch (overrides env)")
	rootCmd.Flags().Float64VarP(&temperature, "temperature", "T", -1, "AI model temperature (0.0-1.0, overrides env)")
	rootCmd.Flags().StringVar(&geminiModel, "model", "", "Gemini model to use (overrides env)")
	rootCmd.Flags().StringVarP(&productHuntURL, "url", "u", "", "Product Hunt URL to scrape (overrides env)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed configuration and progress")
}
