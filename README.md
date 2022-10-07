# Lemmy-Wails

This is a boilerplate generator/library for `wails` projects.

## Installation

Run the following in the folder you want to build your app:

```bash
go mod init example.com
go run github.com/LouisBrunner/lemmy-wails/cmd/lemmy init -name "App Name" -authorName "your name" -authorEmail "your email" -repo "github.com/your/repo"
```

## Update & Sync

You can ensure you are in-sync with the boilerplate by running `go run github.com/LouisBrunner/lemmy-wails/cmd/lemmy sync`. You can also use `go run github.com/LouisBrunner/lemmy-wails/cmd/lemmy update`, which will both update to the latest version of the boilerplate and sync it.
