# Geo-ID

Indonesian Administrative Region API built with Go and [Fiber](https://gofiber.io) framework.

## Overview

Geo-ID provides a REST API for accessing Indonesian administrative region data including provinces (states), cities, districts, and villages. This is a Go port of the [lokasi-id](https://github.com/ikhsanfalakh/lokasi-id) project.

## Features

- 🇮🇩 Complete Indonesian administrative region data
- 🚀 Built with Go for high performance
- 🔧 Uses Fiber framework
- 📦 File-based data storage (JSON)
- � Swagger/OpenAPI documentation
- �🔄 Easy data updates via download script

## Prerequisites

- Go 1.21 or higher
- Python 3.x (for data extraction script)
- `swag` CLI (optional, for regenerating docs)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/ikhsanfalakh/geo-id.git
cd geo-id
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (optional):
```bash
cp .env.example .env
# Edit .env with your preferred configuration
```

4. Download the data:
```bash
chmod +x scripts/download_data.sh
./scripts/download_data.sh
```

## Running the Server

### Development Mode

```bash
```bash
go run main.go
```

### Production Build

```bash
go build -o geo-id
./geo-id
```

The server will start on `http://localhost:8080` by default.

## API Documentation

Interactive API documentation (Swagger/OpenAPI) is available at:

```
http://localhost:8080/apidocs/index.html
```

## API Endpoints

### States (Provinces)

- `GET /states` - Get all states/provinces
- `GET /states/:id` - Get specific state by code
- `GET /states/:id/cities` - Get all cities in a state

### Cities

- `GET /cities/:id` - Get specific city by code
- `GET /cities/:id/districts` - Get all districts in a city

### Districts

- `GET /districts/:id` - Get specific district by code
- `GET /districts/:id/villages` - Get all villages in a district

### Villages

- `GET /villages/:id` - Get specific village by code

## Example Usage

### Get all provinces
```bash
curl http://localhost:8080/states
```

Response:
```json
[
  {
    "code": "11",
    "value": "ACEH"
  },
  {
    "code": "12",
    "value": "SUMATERA UTARA"
  },
  ...
]
```

### Get specific province
```bash
curl http://localhost:8080/states/11
```

### Get cities in a province
```bash
curl http://localhost:8080/states/11/cities
```

## Configuration

The application can be configured using environment variables. You can set these in a `.env` file or export them in your shell.

### Environment Variables

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `PORT` | Server port | `8080` | `3000` |
| `APP_NAME` | Application name | `Geo-ID API` | `My Geo API` |
| `APP_VERSION` | Application version | `1.0` | `2.0` |
| `ENV` | Environment mode | `development` | `production` |
| `ENABLE_SWAGGER` | Enable/disable Swagger UI | `true` | `false` |
| `DATA_DIR` | Custom data directory path | `./data` | `/path/to/data` |

### Using .env File

1. Copy the example file:
```bash
cp .env.example .env
```

2. Edit `.env` with your configuration:
```bash
PORT=3000
ENV=production
ENABLE_SWAGGER=false
```

3. Run the application (it will automatically load `.env`):
```bash
./geo-id
```

### Using Environment Variables Directly

```bash
PORT=3000 ENV=production ENABLE_SWAGGER=false ./geo-id
```

### Configuration Examples

**Development Mode (default):**
```bash
# Uses defaults from .env.example
go run main.go
```

**Production Mode:**
```bash
# Create .env file
cat > .env << EOF
PORT=8080
ENV=production
ENABLE_SWAGGER=false
APP_NAME=Geo-ID API
APP_VERSION=1.0
EOF

# Run the application
./geo-id
```


## Project Structure

```
.
├── main.go                  # Application entry point
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
├── .gitignore               # Git ignore rules
├── README.md                # Project documentation
├── docs/                    # Swagger documentation
│   ├── assets/              # Static assets (logos, etc.)
│   ├── docs.go              # Generated Swagger Go code
│   ├── swagger.json         # Generated Swagger JSON
│   └── swagger.yaml         # Generated Swagger YAML
├── internal/                # Internal application code
│   ├── model/
│   │   ├── region.go        # Data models (Region struct)
│   │   └── error.go         # Error response model
│   ├── service/
│   │   └── location.go      # Business logic (data reading)
│   └── handler/
│       └── location.go      # HTTP handlers (API endpoints)
├── scripts/                 # Utility scripts
│   ├── download_data.sh     # Bash wrapper for extraction
│   └── extract_data.py      # Python script to extract SQL to JSON
├── data/                    # Generated JSON data files
│   ├── states.json          # 38 provinces
│   ├── cities/              # 38 files (one per province)
│   ├── districts/           # 514 files (one per city)
│   └── villages/            # 7,284 files (one per district)
└── raw/                     # Downloaded raw data
    └── wilayah.sql          # Source SQL file from cahyadsn/wilayah
```

## Data Source

The data is sourced from [cahyadsn/wilayah](https://github.com/cahyadsn/wilayah) repository, which contains official Indonesian administrative region data based on Kepmendagri No 300.2.2-2138 Tahun 2025. The extraction script downloads the SQL file and converts it to JSON format.

Data includes:
- 38 Provinces (Provinsi)
- 514+ Regencies/Cities (Kabupaten/Kota)
- 7,000+ Districts (Kecamatan)
- 80,000+ Villages (Kelurahan/Desa)

To update the data:
```bash
./scripts/download_data.sh
```

## Known Issues

None. All endpoints are fully functional.

## Development

### Adding New Endpoints

1. Add the model in `internal/model/`
2. Implement the service logic in `internal/service/`
3. Create the handler in `internal/handler/`
4. Register the route in `main.go`

### Running Tests

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Acknowledgments

- Original project: [lokasi-id](https://github.com/ikhsanfalakh/lokasi-id) by ikhsanfalakh
- Framework: [Fiber](https://gofiber.io)
- Data source: [cahyadsn/wilayah](https://github.com/cahyadsn/wilayah) - Official Indonesian administrative region data

## Support

For issues and questions, please open an issue on GitHub.
