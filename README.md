# go-ph

`go-ph` is a command-line interface (CLI) tool that fetches and displays products launched today on [Product Hunt](https://producthunt.com).

## Description

go-ph allows you to quickly see the latest product launches without having to visit the Product Hunt website, helping you stay updated on the newest tech products and startups. It presents the information in a clean, easy-to-read table format directly in your terminal.

## Features

- Fetches today's product launches from Product Hunt
- Uses Gemini AI to extract structured data from the page content
- Presents products in a beautifully formatted terminal table
- Shows product rank, name, and description

## Installation

### Prerequisites

- Go 1.23.2 or higher
- Gemini API key
- Jina AI API key

### Install from source

```bash
# Clone the repository
git clone https://github.com/bm611/go-ph.git
cd go-ph

# Build the application
go build -o go-ph

# Move the binary to your PATH (optional)
mv go-ph /usr/local/bin/
```

## Configuration

Before using go-ph, you need to set up your API keys as environment variables:

```bash
export GEMINI_API_KEY="your-gemini-api-key"
export JINA_API_KEY="your-jina-api-key"
```

You can add these to your shell profile file (`~/.bashrc`, `~/.zshrc`, etc.) for persistence.

## Usage

Just run the command in your terminal:

```bash
go-ph
```

The tool will fetch the latest products from Product Hunt and display them in a nice table.

## Example Output

```
Extracting content from producthunt.com.....
✓ Content extracted successfully
Formatting structured response with Gemini...
✓ Response formatted successfully
```
| RANK | PRODUCT | DESCRIPTION |
|------|---------|-------------|
| #1 | Pixel Pro | AI-powered design tool that transforms sketches into UI |
| #2 | DevStack | Complete development environment setup in one click |
| #3 | EcoTracker | Carbon footprint calculator for your digital products |
| #4 | SoundScape | Generate ambient background music with AI |
| #5 | LaunchMetrics | Analytics platform for product launches and marketing |
```
