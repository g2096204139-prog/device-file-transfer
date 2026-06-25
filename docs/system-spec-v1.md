# System Specification v1.0

## 1. System Name

Device File Transfer

## 2. Purpose

Device File Transfer allows users to transfer files between devices on the same Wi-Fi or LAN through a browser-based interface.

## 3. Scope

v1.0 focuses on LAN file transfer and leaves expansion points for remote transfer and cloud storage.

Supported in v1.0:

- LAN access
- Web interface
- Multi-file upload
- Download
- File list
- Delete
- Fixed token authentication
- Local storage
- Safety checks
- Operation logs

Not supported in v1.0:

- P2P NAT traversal
- Public internet deployment
- Cloud storage implementation
- User accounts
- End-to-end encryption
- Resumable upload

## 4. User Scenario

1. User starts the server on a laptop.
2. Laptop and phone connect to the same Wi-Fi.
3. User opens `http://laptop-ip:8000` on the phone.
4. User enters the access token.
5. User uploads one or more files.
6. Other devices can download or delete files after authentication.

## 5. Functional Requirements

### FR-01 File Upload

Users can upload one or more files through the web UI.

Rules:

- Files are saved into `uploads/`.
- Duplicate names are renamed automatically.
- Dangerous filenames are rejected or sanitized.
- High-risk extensions are blocked.
- Single file size limit is configurable.

### FR-02 File Download

Users can download files from the file list.

Rules:

- Only files inside the storage layer can be accessed.
- Path traversal is blocked.
- Downloads require token authentication.

### FR-03 File List

Users can see:

- Filename
- File size
- Uploaded time
- Download button
- Delete button

### FR-04 File Delete

Users can delete files after token authentication.

Rules:

- Delete can be disabled through config.
- Filename must pass sanitization.
- Operation is logged.

### FR-05 Authentication

File APIs require:

```text
X-Access-Token: configured-token
```

### FR-06 Operation Log

The system records:

- Upload success
- Upload rejection
- Download
- Delete
- Authentication failure
- Internal errors

## 6. Non-functional Requirements

| Type | Requirement |
|---|---|
| Usability | Browser-based operation |
| Compatibility | Desktop and mobile browsers |
| Security | Token, filename safety, blacklist, path protection |
| Maintainability | Layered Go modules |
| Extensibility | Storage interface for future cloud support |
| Performance | Streaming file operations where possible |

## 7. Configuration

| Variable | Default | Description |
|---|---:|---|
| `SERVER_HOST` | `0.0.0.0` | Bind address |
| `SERVER_PORT` | `8000` | Port |
| `UPLOAD_DIR` | `uploads` | Upload folder |
| `ACCESS_TOKEN` | `change-this-token` | Token |
| `AUTH_ENABLED` | `true` | Enable auth |
| `MAX_UPLOAD_SIZE_MB` | `500` | Single file size limit, `0` means unlimited |
| `ALLOW_DELETE` | `true` | Enable delete |
| `LOG_DIR` | `logs` | Log directory |

## 8. Security Notes

v1.0 is intended for trusted LAN environments. Public internet exposure requires additional work:

- HTTPS
- Stronger authentication
- Rate limiting
- CSRF protection if cookie-based auth is added
- Malware scanning if untrusted users can upload files
- Audit log management
