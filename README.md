# Archiver

A powerful web archive scanner tool that helps you discover and verify live URLs from the Wayback Machine archive.

## Features

- 🔍 Scan web archives for specific domain patterns
- ⚡ Concurrent scanning with configurable workers
- 📝 Save results to file
- 🎯 Custom URL pattern matching
- 📊 Status code reporting
- 🚀 High-performance goroutine implementation

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/archiver.git

# Navigate to the directory
cd archiver

# Build the project
go build -o archiver
```

## Usage

```bash
# Basic usage with default settings
./archiver

# Custom number of workers and pattern
./archiver -w 20 -p "*.example.com/*"

# Save results to file
./archiver -o results.txt

# Full options
./archiver -w 20 -p "*.example.com/*" -o results.txt
```

### Command Line Options

- `-w`: Number of concurrent workers (default: 10)
- `-p`: URL pattern to search for (default: "*.example.com/*")
- `-o`: Output file (optional)

## Example Output

```
    █████╗ ██████╗  ██████╗██╗  ██╗██╗██╗   ██╗███████╗██████╗ 
   ██╔══██╗██╔══██╗██╔════╝██║  ██║██║██║   ██║██╔════╝██╔══██╗
   ███████║██████╔╝██║     ███████║██║██║   ██║█████╗  ██████╔╝
   ██╔══██║██╔══██╗██║     ██╔══██║██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
   ██║  ██║██║  ██║╚██████╗██║  ██║██║ ╚████╔╝ ███████╗██║  ██║
   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
                                                                  
   Web Archive Scanner v1.0
   Author: MIRI

[*] Starting scan with 10 workers
[*] Pattern: *.example.com/*
[*] Time: 2024-03-14 12:34:56

[*] Found 150 URLs to check

[+] https://example.com --> Alive (Status: 200)
[+] https://sub.example.com --> Alive (Status: 200)
```

## Requirements

- Go 1.16 or higher
- Internet connection for accessing web archives

## License

MIT License

## Author

MIRI 