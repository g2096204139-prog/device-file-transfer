# Development Log

This document records the project decisions, implementation progress, mobile testing notes, and planned next steps for Device File Transfer.

## Project overview

- Project name: Device File Transfer
- GitHub repository: `g2096204139-prog/device-file-transfer`
- Main goal: provide a LAN-based cross-device file transfer tool with a browser-based UI.
- Backend language: Go
- Primary storage in v1.x: local `uploads/` directory
- Future storage direction: storage interface reserved for cloud object storage such as S3-compatible storage, Google Cloud Storage, or Azure Blob.

## Confirmed v1.0 scope

The v1.0 line focuses on local network file transfer and mobile browser usability.

Included features:

- Web UI
- Token-based access control
- Multiple file upload
- Upload progress bar
- File list
- File download
- File delete
- Local storage under `uploads/`
- Filename sanitization
- Path traversal protection
- Hidden system file filtering for `.gitkeep`, `.DS_Store`, and `Thumbs.db`
- File size configuration
- Extension blacklist
- Basic operation logging
- GitHub-ready project documentation and templates

Not included yet:

- Full upload resume
- Chunked upload
- Cloud sync
- P2P transfer
- HTTPS automation
- User account system

## Version progress

### v1.0.1

Initial corrected draft after project skeleton generation.

Completed:

- Go project structure
- README, docs, changelog, GitHub templates
- Basic API and web UI
- Token validation
- Local storage
- Basic tests
- `.env.example`
- Local testing checklist

### v1.0.2

Mobile file action fixes.

Completed:

- Hide `.gitkeep`, `.DS_Store`, and `Thumbs.db` from file list.
- Prevent download/delete for hidden system files.
- Improve mobile download handling.
- Improve delete button behavior.
- Update Go module path to `github.com/g2096204139-prog/device-file-transfer`.

### v1.0.3

Upload UI reset fixes.

Completed:

- Clear selected files after successful upload.
- Reset progress bar to `0%` after upload.
- Reset upload UI after delete.
- Verified on Android mobile browser.

### v1.0.4

Local environment and version display fixes.

Completed:

- Update displayed server version to `1.0.4`.
- Automatically load `.env` from the project root.
- Keep shell environment variables higher priority than `.env`.
- Remove the fixed UI hint that referenced `change-this-token`.
- Update README and local test checklist for `.env` usage.

### v1.0.5

Development log and documentation tracking.

Completed:

- Add `docs/development-log.md`.
- Record current project decisions and mobile testing status.
- Update displayed server version to `1.0.5`.
- Update changelog.

### v1.0.6

README configuration cleanup.

Completed:

- Clarify `.env` auto-loading behavior.
- Clarify configuration priority between shell environment variables and `.env`.
- Document Termux upload storage location.
- Clarify that delete only removes uploaded server-side copies.
- Update displayed server version to `1.0.6`.


## Mobile testing notes

Test environment:

- Android phone
- Termux
- Go installed through Termux
- Browser access through `http://localhost:8000`

Important notes:

- Running Go projects directly under `/sdcard/Download` can cause Go file-locking errors such as `RLock ... function not implemented`.
- The project should be copied to the Termux home directory and executed from:

```text
~/device-file-transfer
```

- Uploaded files are stored under:

```text
~/device-file-transfer/uploads/
```

- Deleting a file from the web UI only deletes the uploaded server-side copy under `uploads/`.
- It does not delete the original file from `/sdcard/Download` or any other phone folder.

## GitHub workflow notes

The repository was pushed from Termux using HTTPS and a GitHub Personal Access Token.

Useful commands:

```bash
git status
git add .
git commit -m "message"
git push
```

When updating from a regenerated zip package, the local folder may need to be reinitialized with Git and pushed back to the same remote.

## Current behavior notes

### Token

Use a local `.env` file for real use:

```env
ACCESS_TOKEN=your-own-token
```

The `.env` file should not be committed to GitHub.

### Browser URL

When opening the UI on the same device, use:

```text
http://localhost:8000
```

or:

```text
http://127.0.0.1:8000
```

Do not use `http://0.0.0.0:8000` in the browser. `0.0.0.0` is for server binding, not for normal browser access.

### Android Chrome download warning

Android Chrome may show a safety warning when downloading files from a local HTTP server. This is expected for local HTTP downloads. If the file is known and was uploaded by the user, the user can choose to keep it.

## Resume transfer discussion

Current status:

- Download resume: partially supported by the current streaming download foundation, but not yet formally tested and documented as a stable feature.
- Upload resume: not implemented yet.

Planned roadmap:

- v1.1.0: formal download resume support with HTTP Range validation.
- v2.0.0: upload resume through chunked upload.
- v2.0.1: stability fixes for chunk cleanup, merge recovery, duplicate complete protection, and status display.

## Recommended next steps

1. v1.0.6 or v1.1.0 planning decision:
   - Option A: continue polishing v1.0.x UX and security documentation.
   - Option B: start v1.1.0 and formally implement/test download resume.
2. Add handler/service tests for API behavior.
3. Add a user-facing explanation that delete only removes the server-side uploaded copy.
4. Improve download UX and avoid exposing `0.0.0.0` to users.
5. Add release tags on GitHub after stable verification.
