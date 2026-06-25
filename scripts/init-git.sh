#!/usr/bin/env bash
set -euo pipefail

git init
git add .
git commit -m "chore: initialize project structure"

echo "Next:"
echo "git remote add origin https://github.com/YOUR_USERNAME/device-file-transfer.git"
echo "git branch -M main"
echo "git push -u origin main"
