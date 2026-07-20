# Learning Golang in VS Code

Welcome to your Go development workspace! Everything you need to get started with Golang is configured.

## Workspace Layout
- [main.go](file:///Users/prankursharma/Developer/GOLANG/main.go) - Your entry point/first program.
- [go.mod](file:///Users/prankursharma/Developer/GOLANG/go.mod) - The Go module definition file.
- [.vscode/settings.json](file:///Users/prankursharma/Developer/GOLANG/.vscode/settings.json) - Settings for the Go extension (automatic formatting and imports sorting on save).

---

## Getting Started

### 1. How to run your program
Open your terminal inside this folder and run:
```bash
go run main.go
```

### 2. How to build your program into a binary
To compile your code into a standalone binary file:
```bash
go build main.go
./main
```

### 3. VS Code Features Enabled
- **Format on Save**: Every time you save a `.go` file, it will automatically format according to standard Go styling (`gofmt`/`gofumpt`).
- **Organize Imports on Save**: Unused imports will be deleted, and new imports will be auto-added when you save.
- **Language Server (`gopls`)**: Provides autocomplete, hover documentation, syntax highlighting, and code navigation (Cmd+Click).

---

## Helpful Resources for Learning Go
1. **[A Tour of Go](https://tour.golang.org/)**: The official interactive tour of Go basics.
2. **[Go by Example](https://gobyexample.com/)**: A hands-on introduction to Go using annotated example programs.
3. **[Effective Go](https://golang.org/doc/effective_go)**: Tips for writing clear, idiomatic Go code.
