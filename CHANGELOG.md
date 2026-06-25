# Changelog

## v1.0.6 - README configuration cleanup

- Clarify that the server automatically loads `.env` from the project root.
- Document configuration priority: shell environment variables override `.env`, and `.env` overrides built-in defaults.
- Clarify where uploaded files are stored on Termux.
- Clarify that web UI delete removes only the server-side uploaded copy under `uploads/`.
- Update displayed server version to `1.0.6`.

## v1.0.5 - Development log and documentation tracking

- Add `docs/development-log.md` to preserve project decisions, mobile testing notes, and roadmap discussion.
- Update displayed server version to `1.0.5`.
- Record current status of download resume and upload resume planning.

## v1.0.4 - Local environment and version display fixes

- Update server info version to `1.0.4`.
- Automatically load `.env` from the project root when starting the Go server.
- Keep shell environment variables higher priority than `.env` values.
- Replace the fixed default-token UI hint with a safer generic token message.
- Update README and local test checklist for `.env` usage.

## v1.0.3 - Upload UI reset fixes

- Reset selected file input after successful upload.
- Reset upload progress bar and percentage after successful upload.
- Reset upload input state after delete actions so stale file selections are not shown on mobile browsers.

## v1.0.2 - Mobile file action fixes

- Hide `.gitkeep`, `.DS_Store`, and `Thumbs.db` from file listings.
- Block download and delete operations for hidden system placeholder files.
- Improve mobile download behavior with attachment disposition and safer client-side download handling.
- Set button types explicitly in the web UI to avoid unintended navigation or form submission.
- Update Go module path to `github.com/g2096204139-prog/device-file-transfer`.


## v1.0.1

- Fixed README delete API path to match the implemented query parameter endpoint.
- Added `.env.example` for local configuration reference.
- Added `docs/local-test-checklist.md` for build, API, Web UI, security, and LAN smoke tests.


## 1.0.0 - Initial planned release

### Added

- LAN-first web file transfer
- Go backend skeleton
- Local storage implementation
- Token authentication
- Multi-file upload API
- Download, list, and delete APIs
- Upload size limit configuration
- Filename sanitization and path traversal protection
- Extension blacklist
- Operation logging
- Single-page web UI with progress bar and mobile-friendly layout
- Open-source project documentation templates