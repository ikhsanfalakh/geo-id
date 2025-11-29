# Geo-ID

Indonesian Administrative Region API built with Go and [gowok](https://github.com/gowok/gowok) framework.

## Overview

Geo-ID provides a REST API for accessing Indonesian administrative region data including provinces (states), cities, districts, and villages. This is a Go port of the [lokasi-id](https://github.com/ikhsanfalakh/lokasi-id) project.

## Features

- 🇮🇩 Complete Indonesian administrative region data
- 🚀 Built with Go for high performance
- 🔧 Uses gowok framework
- 📦 File-based data storage (JSON)
- 🔄 Easy data updates via download script

## Prerequisites

- Go 1.21 or higher
- Python 3.x (for data extraction script)

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

3. Download the data:
```bash
chmod +x scripts/download_data.sh
./scripts/download_data.sh
```

## Running the Server

### Development Mode

```bash
go run main.go --config=config.yaml
```

### Production Build

```bash
go build -o geo-id
./geo-id --config=config.yaml
```

The server will start on `http://localhost:8080` by default.

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
  }
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

Edit `config.yaml` to customize the server settings:

```yaml
web:
  enabled: true
  host: :8080
```

## Project Structure

```
.
├── config.yaml              # Server configuration
├── main.go                  # Application entry point
├── internal/
│   ├── model/
│   │   └── region.go       # Data models
│   ├── service/
│   │   └── location.go     # Business logic
│   └── handler/
│       └── location.go     # HTTP handlers
├── data/                    # JSON data files
│   ├── states.json
│   ├── cities/
│   ├── districts/
│   └── villages/
└── scripts/
    └── download_data.sh    # Data download script
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

⚠️ **Parameter Routes**: Currently, endpoints with URL parameters (e.g., `/states/:id`) are experiencing issues. The `/states` endpoint works correctly, but parameterized routes need debugging. This is being investigated.

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
- Framework: [gowok](https://github.com/gowok/gowok)
- Data source: [cahyadsn/wilayah](https://github.com/cahyadsn/wilayah) - Official Indonesian administrative region data

## Support

For issues and questions, please open an issue on GitHub.
