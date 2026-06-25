# Roadmap

## v1.0.0 - Basic LAN File Transfer

Goal: deliver a usable LAN-first web file transfer system.

Features:

- Web UI
- Multiple file upload
- File list
- File download
- File delete
- Fixed token authentication
- Local storage through `Storage` interface
- Configurable max upload size, default 500 MB
- `0` means unlimited upload size
- Filename sanitization
- Path traversal protection
- Extension blacklist
- Operation logs
- Mobile-friendly UI

## v1.1.0 - Download Resume

Goal: improve large file download reliability.

Planned work:

- HTTP Range support
- `206 Partial Content`
- `Accept-Ranges: bytes`
- Large file streaming optimization

Note: v1.0 already uses `http.ServeContent`, which provides a foundation for Range support.

## v1.2.0 - UX Improvements

Planned work:

- Better upload progress UI
- QR Code for LAN URL
- Server IP display improvement
- Drag-and-drop upload
- Better mobile layout
- Better error messages

## v1.3.0 - Management and Security

Planned work:

- Config file support
- Better log viewer or log format
- Delete policy refinement
- Optional extension allowlist
- Safer public config endpoint
- Basic rate limiting

## v2.0.0 - Resumable Upload

Goal: support interrupted large file upload recovery.

Planned APIs:

- `POST /api/uploads/init`
- `PUT /api/uploads/{upload_id}/chunks/{chunk_index}`
- `GET /api/uploads/{upload_id}/status`
- `POST /api/uploads/{upload_id}/complete`
- `DELETE /api/uploads/{upload_id}`

Planned features:

- Chunk upload
- Upload task metadata
- Missing chunk detection
- Chunk merge
- SHA-256 verification
- Temporary chunk cleanup

## v2.0.1 - Resumable Upload Stabilization

Planned work:

- Better chunk cleanup
- Safer complete operation
- Duplicate complete protection
- Invalid chunk index protection
- More tests
- Better frontend recovery state
