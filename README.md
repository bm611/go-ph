# go-ph

`go-ph` is a command-line interface (CLI) tool that fetches and displays products launched today on [Product Hunt](https://producthunt.com).

[DeepWiki](https://deepwiki.com/bm611/go-ph)

## Description

go-ph allows you to quickly see the latest product launches without having to visit the Product Hunt website, helping you stay updated on the newest tech products and startups. It presents the information in a clean, easy-to-read table format directly in your terminal.

## Features

- 🚀 Fetches today's product launches from Product Hunt
- 🤖 Uses Gemini AI to extract structured data from the page content
- 📊 Presents products in a beautifully formatted terminal table
- 🎯 Shows product rank, name, and description
- ⚙️ Fully configurable via environment variables or CLI flags
- 🔧 Built-in configuration validation and debugging tools
- 📱 Multiple output formats and customization options
- 🛠️ Modular architecture with clean separation of concerns

## Installation

### Prerequisites

- Go 1.23.2 or higher
- [Gemini API key](https://aistudio.google.com/app/apikey)
- [Jina AI API key](https://jina.ai/)

### Quick Install

```bash
go install github.com/bm611/go-ph@latest
```

### Install from source

```bash
# Clone the repository
git clone https://github.com/bm611/go-ph.git
cd go-ph

# Build the application
go build -o go-ph

# Move the binary to your PATH (optional)
sudo mv go-ph /usr/local/bin/
```

## Quick Setup

1. **Get API Keys:**
   - [Gemini API Key](https://aistudio.google.com/app/apikey) (Google AI Studio)
   - [Jina AI API Key](https://jina.ai/) (Jina AI Dashboard)

2. **Set Environment Variables:**

   ```bash
   export GEMINI_API_KEY="your-gemini-api-key"
   export JINA_API_KEY="your-jina-api-key"
   ```

3. **Verify Setup:**

   ```bash
   go-ph config
   ```

4. **Run:**
   ```bash
   go-ph
   ```

For detailed setup instructions, see [SETUP.md](SETUP.md).

## Usage

### Basic Usage

```bash
# Fetch today's Product Hunt launches
go-ph

# Show detailed progress
go-ph --verbose

# Check configuration
go-ph config

# Show version
go-ph version
```

### Advanced Usage

```bash
# Customize number of products
go-ph --max-products 5

# Use different AI creativity level
go-ph --temperature 0.8

# Use different Gemini model
go-ph --model gemini-1.5-pro

# Show help
go-ph --help
```

## Example Output

```bash
$ go-ph
Extracting content from producthunt.com.....
✓ Content extracted successfully
Formatting structured response with Gemini...
✓ Response formatted successfully
┌──────┬───────────────────────────┬─────────────────────────────────────────────────────────────┐
│ RANK │           NAME            │                         DESCRIPTION                         │
├──────┼───────────────────────────┼─────────────────────────────────────────────────────────────┤
│ 1    │ nFactorial AI             │ Video calls with world's best minds as your personal tutors │
├──────┼───────────────────────────┼─────────────────────────────────────────────────────────────┤
│ 2    │ Hyprnote                  │ AI Notepad for Private Meetings — fully on your device      │
├──────┼───────────────────────────┼─────────────────────────────────────────────────────────────┤
│ 3    │ My Juno Health: AI Doctor │ Smarter Health. Sharper Mind. Reach Your Peak Productivity  │
├──────┼───────────────────────────┼─────────────────────────────────────────────────────────────┤
│ 4    │ Dad Reply                 │ Auto-respond with a 👍 - Minimal effort - Maximum ambiguity │
├──────┼───────────────────────────┼─────────────────────────────────────────────────────────────┤
│ 5    │ SuperCraft                │ Figma for designing physical products                       │
└──────┴───────────────────────────┴─────────────────────────────────────────────────────────────┘
```

### With Verbose Output

```bash
$ go-ph --verbose
Configuration:
- Max Products: 10
- Temperature: 0.10
- Gemini Model: gemini-2.5-flash
- Product Hunt URL: https://producthunt.com
- Jina Base URL: https://r.jina.ai/
- Response MIME: application/json

API Keys Status:
- Gemini API Key: ✓ Set
- Jina API Key: ✓ Set

Extracting content from producthunt.com.....
✓ Content extracted successfully
Formatting structured response with Gemini...
✓ Response formatted successfully
[Table output...]
```

## Configuration

All settings can be customized via environment variables:

| Variable             | Default                   | Description                |
| -------------------- | ------------------------- | -------------------------- |
| `GEMINI_API_KEY`     | _required_                | Your Google Gemini API key |
| `JINA_API_KEY`       | _required_                | Your Jina AI API key       |
| `PRODUCT_HUNT_URL`   | `https://producthunt.com` | URL to scrape              |
| `GEMINI_MODEL`       | `gemini-2.5-flash`        | AI model to use            |
| `GEMINI_TEMPERATURE` | `0.1`                     | Creativity level (0.0-1.0) |
| `MAX_PRODUCTS`       | `10`                      | Maximum products to fetch  |

See [SETUP.md](SETUP.md) for complete configuration options.

## Commands

- `go-ph` - Fetch and display today's products
- `go-ph config` - Show current configuration and validate setup
- `go-ph version` - Display version information
- `go-ph --help` - Show all available options

## Project Structure

```
go-ph/
├── cmd/                    # CLI commands
│   ├── root.go            # Main command
│   ├── config.go          # Configuration command
│   └── version.go         # Version command
├── internal/
│   ├── config/            # Configuration management
│   ├── service/           # Business logic layer
│   ├── llm/              # AI integration
│   ├── scraper/          # Web scraping
│   └── ui/               # Table rendering
├── SETUP.md              # Detailed setup guide
└── .env.example          # Configuration template
```
