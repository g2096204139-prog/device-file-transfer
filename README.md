# Device File Transfer

Device File Transfer is a LAN-first file transfer tool for sharing files between devices on the same Wi-Fi or local network.

## v1.0 Scope

- Web interface
- Multiple file upload
- File list
- File download
- File delete
- Fixed token authentication
- Configurable upload size limit
- Local storage through a storage interface
- Filename sanitization and path traversal protection
- High-risk extension blacklist
- Operation logging
- Mobile-friendly single-page UI
- Future cloud storage and remote transfer expansion points

## Tech Stack

- Backend: Go
- Frontend: HTML, CSS, JavaScript
- Storage: local `uploads/` folder in v1.0
- Auth: `X-Access-Token`

## Project Structure

```text
device-file-transfer/
├── cmd/server/main.go
├── internal/
│   ├── auth/
│   ├── config/
│   ├── handler/
│   ├── logger/
│   ├── service/
│   ├── storage/
│   └── util/
├── web/
│   ├── static/
│   └── templates/
├── uploads/
├── docs/
└── .github/
```

## Run Locally

```bash
go run ./cmd/server
```

Open:

```text
http://localhost:8000
```

For another device on the same Wi-Fi, open:

```text
http://YOUR_LAPTOP_IP:8000
```

## Configuration and `.env`

The server automatically loads a local `.env` file from the project root when it starts.

Configuration priority:

```text
shell environment variables > .env values > built-in defaults
```

Copy `.env.example` to `.env` for local use:

```bash
cp .env.example .env
nano .env
go run ./cmd/server
```

Never commit your real `.env` file or real access token to GitHub. Only `.env.example` should be tracked.

| Variable | Default | Description |
|---|---:|---|
| `SERVER_HOST` | `0.0.0.0` | Server bind address. Use `localhost` or `127.0.0.1` in the browser for local testing. |
| `SERVER_PORT` | `8000` | Server port |
| `UPLOAD_DIR` | `uploads` | File storage directory |
| `ACCESS_TOKEN` | `change-this-token` | Fixed access token. Change this in `.env` before real use. |
| `AUTH_ENABLED` | `true` | Enable token auth |
| `MAX_UPLOAD_SIZE_MB` | `500` | Upload size limit. `0` means unlimited |
| `ALLOW_DELETE` | `true` | Allow delete API |
| `LOG_DIR` | `logs` | Log directory |

Uploaded files are stored on the device running the server, under the configured `UPLOAD_DIR`.

For the current Termux setup, that is usually:

```text
/data/data/com.termux/files/home/device-file-transfer/uploads/
```

Deleting a file in the web UI deletes only the uploaded server-side copy in `uploads/`. It does not delete the original file from Android Downloads or another source folder.


## API Summary

| Method | Path | Description |
|---|---|---|
| `GET` | `/` | Web UI |
| `GET` | `/api/health` | Health check |
| `GET` | `/api/files` | List files |
| `POST` | `/api/upload` | Upload one or more files |
| `GET` | `/api/download/{filename}` | Download file |
| `DELETE` | `/api/files?filename={filename}` | Delete file |
| `GET` | `/api/server-info` | Public server info |

All file APIs require:

```text
X-Access-Token: your-token
```

## Security Notes

This project is designed for trusted LAN use first. Do not expose it directly to the public internet without HTTPS, stronger authentication, rate limiting, and additional hardening.

v1.0 blocks common high-risk executable/script extensions such as `.exe`, `.bat`, `.cmd`, `.sh`, `.ps1`, `.msi`, `.dll`, `.js`, `.vbs`, `.jar`, and similar extensions.

## Roadmap

See [`docs/roadmap.md`](docs/roadmap.md).


## GitHub First Push

See [`docs/github-setup.md`](docs/github-setup.md).

Local testing checklist: [`docs/local-test-checklist.md`](docs/local-test-checklist.md).

Suggested first commit:

```bash
git init
git add .
git commit -m "chore: initialize project structure"
```
