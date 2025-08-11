# go-ph Setup Guide

This guide will help you set up and configure `go-ph` to fetch Product Hunt launches.

## Prerequisites

- Go 1.23.2 or higher
- Internet connection
- API keys for Gemini AI and Jina AI

## Installation

### Option 1: Install from Source

```bash
# Clone the repository
git clone https://github.com/bm611/go-ph.git
cd go-ph

# Build the application
go build -o go-ph

# Move to your PATH (optional)
sudo mv go-ph /usr/local/bin/
```

### Option 2: Direct Go Install

```bash
go install github.com/bm611/go-ph@latest
```

## API Key Setup

### 1. Get a Gemini API Key

1. Visit [Google AI Studio](https://aistudio.google.com/app/apikey)
2. Sign in with your Google account
3. Click "Create API Key"
4. Copy the generated API key

### 2. Get a Jina AI API Key

1. Visit [Jina AI](https://jina.ai/) and sign up for a free account
2. Go to your dashboard
3. Navigate to API keys section
4. Create a new API key
5. Copy the generated API key

## Configuration

### Method 1: Environment Variables (Recommended)

Set the required environment variables in your shell:

```bash
export GEMINI_API_KEY="your-gemini-api-key-here"
export JINA_API_KEY="your-jina-api-key-here"
```

To make these permanent, add them to your shell profile:

**For Bash (.bashrc or .bash_profile):**
```bash
echo 'export GEMINI_API_KEY="your-gemini-api-key-here"' >> ~/.bashrc
echo 'export JINA_API_KEY="your-jina-api-key-here"' >> ~/.bashrc
source ~/.bashrc
```

**For Zsh (.zshrc):**
```bash
echo 'export GEMINI_API_KEY="your-gemini-api-key-here"' >> ~/.zshrc
echo 'export JINA_API_KEY="your-jina-api-key-here"' >> ~/.zshrc
source ~/.zshrc
```

### Method 2: Using .env File

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and add your API keys:
   ```bash
   nano .env  # or use your preferred editor
   ```

3. Load the environment variables:
   ```bash
   source .env
   ```

## Configuration Options

All configuration can be set via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `GEMINI_API_KEY` | *required* | Your Google Gemini API key |
| `JINA_API_KEY` | *required* | Your Jina AI API key |
| `PRODUCT_HUNT_URL` | `https://producthunt.com` | Product Hunt URL to scrape |
| `JINA_BASE_URL` | `https://r.jina.ai/` | Jina AI reader service URL |
| `GEMINI_MODEL` | `gemini-2.5-flash` | Gemini model to use |
| `GEMINI_TEMPERATURE` | `0.1` | AI creativity level (0.0-1.0) |
| `MAX_PRODUCTS` | `10` | Maximum products to fetch |
| `RESPONSE_MIME_TYPE` | `application/json` | Response format |

## Verification

### 1. Check Configuration

```bash
go-ph config
```

This will show your current configuration and validate that everything is set up correctly.

### 2. Test the Application

```bash
go-ph --verbose
```

This will run the application with detailed output, showing each step of the process.

## Usage Examples

### Basic Usage

```bash
# Fetch today's Product Hunt launches
go-ph
```

### With Custom Options

```bash
# Fetch only 5 products with higher creativity
go-ph --max-products 5 --temperature 0.5 --verbose
```

### Using Different Model

```bash
# Use a different Gemini model
go-ph --model gemini-1.5-pro
```

## Command Reference

### Available Commands

- `go-ph` - Fetch and display today's products
- `go-ph config` - Show configuration and validate settings
- `go-ph version` - Display version information
- `go-ph --help` - Show help and available options

### Available Flags

- `-m, --max-products int` - Maximum number of products to fetch
- `-T, --temperature float` - AI model temperature (0.0-1.0)
- `--model string` - Gemini model to use
- `-u, --url string` - Product Hunt URL to scrape
- `-v, --verbose` - Show detailed progress information

## Troubleshooting

### Common Issues

#### 1. "Configuration error: GEMINI_API_KEY environment variable is required"

**Solution:** Make sure you have set your Gemini API key:
```bash
export GEMINI_API_KEY="your-actual-api-key"
```

#### 2. "Configuration error: JINA_API_KEY environment variable is required"

**Solution:** Make sure you have set your Jina API key:
```bash
export JINA_API_KEY="your-actual-api-key"
```

#### 3. "Error fetching products: failed to fetch page content"

**Possible causes:**
- Invalid Jina API key
- Network connectivity issues
- Jina AI service is down

**Solution:**
1. Verify your Jina API key is correct
2. Check your internet connection
3. Try again later if the service is down

#### 4. "Error fetching products: failed to process content with Gemini"

**Possible causes:**
- Invalid Gemini API key
- API quota exceeded
- Invalid temperature or model settings

**Solution:**
1. Verify your Gemini API key is correct
2. Check your API usage limits
3. Try with default temperature (0.1)

### Debug Mode

Run with verbose flag to see detailed execution:

```bash
go-ph --verbose
```

### Configuration Validation

Check your configuration anytime:

```bash
go-ph config
```

## Advanced Configuration

### Custom Prompt

You can customize the AI prompt by setting the `EXTRACTION_PROMPT` environment variable. The prompt must include placeholders for the number of products (`%d`) and content (`%s`):

```bash
export EXTRACTION_PROMPT="Extract exactly %d products from this content: %s"
```

### Different Product Hunt Pages

You can scrape different Product Hunt pages:

```bash
go-ph --url "https://producthunt.com/topics/artificial-intelligence"
```

## Support

If you encounter any issues:

1. Check this setup guide
2. Run `go-ph config` to validate your configuration
3. Try running with `--verbose` flag for detailed output
4. Check the [GitHub repository](https://github.com/bm611/go-ph) for issues and updates

## Security Notes

- Never commit your API keys to version control
- Keep your `.env` file private
- Regularly rotate your API keys for security
- Monitor your API usage to prevent unexpected charges

## Updates

To update go-ph to the latest version:

```bash
# If installed from source
cd go-ph
git pull origin main
go build -o go-ph

# If installed via go install
go install github.com/bm611/go-ph@latest
```
