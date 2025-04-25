package cmd

import (
	"fmt"
	"os"

	"github.com/bm611/go-ph/internal/llm"
	"github.com/bm611/go-ph/internal/scraper"
	"github.com/bm611/go-ph/internal/ui"
	"github.com/spf13/cobra"
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		res, err := scraper.GetPageContent("https://producthunt.com")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		Prompt := `extract the top 20 products launched today from this text extract in json format
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
									` + "`" + res + "`"
		products, err := llm.GetGeminiResponse(Prompt)
		if err != nil {
			fmt.Println("Error:", err)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ph.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
