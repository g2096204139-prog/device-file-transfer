# GitHub Setup

## Create Repository

Create a new GitHub repository named:

```text
device-file-transfer
```

Do not initialize it with README, because this project already contains one.

## First Commit

From the project root:

```bash
git init
git add .
git commit -m "chore: initialize project structure"
```

## Push to GitHub

HTTPS:

```bash
git remote add origin https://github.com/YOUR_USERNAME/device-file-transfer.git
git branch -M main
git push -u origin main
```

SSH:

```bash
git remote add origin git@github.com:YOUR_USERNAME/device-file-transfer.git
git branch -M main
git push -u origin main
```

## Verify

```bash
git status
git log --oneline
```
