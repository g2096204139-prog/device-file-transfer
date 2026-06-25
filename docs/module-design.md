# Module Design

## Project Layout

```text
cmd/server/
  main.go

internal/config/
  config.go

internal/auth/
  token.go

internal/handler/
  handler.go
  response.go

internal/service/
  file_service.go

internal/storage/
  storage.go
  local.go
  cloud_future.go

internal/logger/
  logger.go

internal/util/
  filename.go

web/templates/
  index.html

web/static/
  app.js
  style.css
```

## Module Responsibilities

### `cmd/server`

Application entry point. It loads config, initializes logger, storage, service, handler, and starts HTTP server.

### `internal/config`

Centralizes environment-based configuration.

### `internal/auth`

Verifies fixed access token through `X-Access-Token`.

### `internal/handler`

HTTP layer. It parses requests, checks authentication, and returns JSON or file responses.

### `internal/service`

Business logic layer. It validates filenames, checks size policy, calls storage, and logs operations.

### `internal/storage`

Storage abstraction layer.

Current implementation:

- `LocalStorage`

Future implementation candidates:

- S3
- Google Cloud Storage
- Azure Blob
- MinIO

### `internal/logger`

Simple operation logger.

### `internal/util`

Filename sanitization, extension blacklist, and safety helpers.

### `web`

Frontend files.

## Dependency Direction

```text
cmd/server
  ↓
handler
  ↓
service
  ↓
storage + util + logger
```

The storage layer does not depend on HTTP logic. This makes future cloud storage integration easier.
