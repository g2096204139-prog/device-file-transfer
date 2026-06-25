# API Design v1.0

## Authentication

All file APIs require:

```text
X-Access-Token: configured-token
```

## Response Format

Success:

```json
{
  "success": true,
  "message": "ok",
  "data": {}
}
```

Failure:

```json
{
  "success": false,
  "message": "unauthorized",
  "error_code": "UNAUTHORIZED"
}
```

## Endpoints

### `GET /`

Returns the web interface.

### `GET /api/health`

Returns server health.

Response:

```json
{
  "success": true,
  "message": "ok",
  "data": {
    "status": "healthy"
  }
}
```

### `GET /api/server-info`

Returns public server information.

Response data:

```json
{
  "name": "Device File Transfer",
  "version": "1.0.0",
  "host": "0.0.0.0",
  "port": "8000",
  "lan_ips": ["192.168.1.10"],
  "auth_enabled": true,
  "max_upload_size_mb": 500,
  "allow_delete": true
}
```

### `GET /api/files`

Returns file list.

Response data:

```json
[
  {
    "filename": "photo.jpg",
    "size": 204800,
    "uploaded_at": "2026-06-12T10:30:00+08:00"
  }
]
```

### `POST /api/upload`

Uploads one or more files.

Request:

```text
Content-Type: multipart/form-data
files: file[]
```

Response data:

```json
{
  "results": [
    {
      "filename": "photo.jpg",
      "size": 204800,
      "status": "success"
    }
  ]
}
```

### `GET /api/download/{filename}`

Downloads a file.

### `DELETE /api/files?filename={filename}`

Deletes a file.

Response:

```json
{
  "success": true,
  "message": "delete success",
  "data": {
    "filename": "photo.jpg"
  }
}
```

## Error Codes

| Code | Meaning |
|---|---|
| `UNAUTHORIZED` | Missing or invalid token |
| `NO_FILES` | Upload request has no file |
| `INVALID_UPLOAD` | Invalid upload request |
| `FILE_NOT_FOUND` | File not found |
| `DELETE_FAILED` | Delete failed |
| `LIST_FAILED` | List failed |
| `METHOD_NOT_ALLOWED` | HTTP method not allowed |
