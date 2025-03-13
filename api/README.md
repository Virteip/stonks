# Stonks API

Backend service for the Stonks application. Retrieves stock data from an external API, stores it in CockroachDB, and provides endpoints for the frontend.

## Requirements

- Go 1.24+
- CockroachDB

## Configuration

Copy `example.local.config.json` to `local.config.json` and update with your CockroachDB credentials and API key.

## Running the Service

```bash
# Start the API server
go run ./cmd
```

The service will start on port 8080 by default.

## Endpoints

### Get All Stocks

```
GET /api/v1/stonks-api/stocks
```

Query parameters:
- `page` - Page number (default: 1)
- `page_size` - Number of items per page (default: 20)

Response:
```json
{
  "stocks": [...],
  "total_count": 100,
  "page_size": 20,
  "page": 1,
  "total_pages": 5
}
```

### Search Stock by Ticker

```
GET /api/v1/stonks-api/stock/:ticker
```

Response:
```json
[
  {
    "id": "...",
    "ticker": "AAPL",
    "company": "Apple Inc.",
    "brokerage": "Example Brokerage",
    "action": "upgraded by",
    "rating_from": "Hold",
    "rating_to": "Buy",
    "target_from": 150.00,
    "target_to": 200.00,
    "time": "2025-01-01T00:00:00Z"
  }
]
```

### Get Recommendations

```
GET /api/v1/stonks-api/recommendations
```

Response:
```json
[
  {
    "stock": {
      "id": "...",
      "ticker": "AAPL",
      "company": "Apple Inc.",
      "brokerage": "Example Brokerage",
      "action": "upgraded by", 
      "rating_from": "Hold",
      "rating_to": "Buy",
      "target_from": 150.00,
      "target_to": 200.00,
      "time": "2025-01-01T00:00:00Z"
    },
    "score": 4.5,
    "reason": "Stock was upgraded, Target price increased significantly"
  }
]
```

## Authentication

All endpoints require an API key provided in the `X-API-Key` header.

Sergio Pietri