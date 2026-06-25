# Local Test Checklist

Use this checklist before pushing or tagging a release.

## 1. Build and Unit Tests

```bash
go test ./...
go build ./cmd/server
```

Expected result:

```text
All tests pass and the server binary builds successfully.
```

## 2. Start Server

```bash
go run ./cmd/server
```

Open:

```text
http://localhost:8000
```

Local token:

```text
Use the ACCESS_TOKEN value from your local .env file.
```

## 3. API Smoke Test

### Health check

```bash
curl http://localhost:8000/api/health
```

Expected: HTTP 200.

### Auth check

```bash
curl -i http://localhost:8000/api/files
```

Expected: HTTP 401 when authentication is enabled.

### List files with token

```bash
curl -H "X-Access-Token: YOUR_LOCAL_TOKEN" http://localhost:8000/api/files
```

Expected: HTTP 200 with a JSON response.

## 4. Web UI Test

- Enter the token.
- Select one or more files.
- Upload files.
- Confirm progress is displayed.
- Confirm the file list refreshes.
- Download an uploaded file.
- Delete an uploaded file.
- Confirm the deleted file disappears from the list.

## 5. Security Test

- Try uploading a blocked extension such as `test.exe`.
- Try downloading a filename containing `../`.
- Try deleting without a token.
- Try uploading a file larger than `MAX_UPLOAD_SIZE_MB`.

Expected result: unsafe or unauthorized operations are rejected.

## 6. LAN Test

On the server laptop, find the LAN IP address.

Examples:

```bash
ipconfig
```

or:

```bash
ifconfig
```

From another device on the same Wi-Fi, open:

```text
http://YOUR_LAPTOP_IP:8000
```

Then repeat the Web UI upload, download, list, and delete tests.
