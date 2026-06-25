# Test Cases v1.0

| ID | Scenario | Steps | Expected Result |
|---|---|---|---|
| TC-01 | Health check | Open `/api/health` | Returns `status=healthy` |
| TC-02 | Valid token list | Call `/api/files` with token | Returns file list |
| TC-03 | Invalid token list | Call `/api/files` without token | Returns 401 |
| TC-04 | Upload one file | Upload `photo.jpg` | File saved |
| TC-05 | Upload multiple files | Upload multiple files | All valid files processed |
| TC-06 | Duplicate filename | Upload same file twice | Second file renamed |
| TC-07 | Blocked extension | Upload `.exe` | Upload rejected |
| TC-08 | Path traversal filename | Upload `../../secret.txt` | Filename rejected or sanitized |
| TC-09 | Download existing file | Click download | File downloaded |
| TC-10 | Download missing file | Download non-existing file | Returns 404 |
| TC-11 | Delete existing file | Delete file with token | File removed |
| TC-12 | Delete without token | Delete file without token | Returns 401 |
| TC-13 | File too large | Upload file larger than limit | Upload rejected |
| TC-14 | Unlimited size config | Set `MAX_UPLOAD_SIZE_MB=0` | Size limit disabled |
| TC-15 | Operation log | Upload/download/delete | Log entries created |
| TC-16 | Mobile layout | Open on phone | UI remains usable |
