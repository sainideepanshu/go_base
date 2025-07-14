# 🧠 Running Go Files on Windows – Quick Guide

This project helps developers using **Go on Windows** to understand why `go run *.go` may not work, and how to fix it.

---

## ❌ Problem: `go run *.go` Not Working on Windows

When you run:

```bash
go run *.go
```

On **Windows CMD or PowerShell**, you’ll see an error like:

```
go: open *.go: The system cannot find the file specified.
```

### Why?
Windows **does not expand wildcards (`*.go`)** like Unix-based terminals. So the command passes the literal `*.go` to Go instead of the actual filenames.

---

## ✅ Solutions

### 1. **Manually list files**
```bash
go run main.go utils.go config.go
```

### 2. **Use Git Bash or WSL**
If you have [Git Bash](https://git-scm.com/download/win) or WSL (Windows Subsystem for Linux) installed:

```bash
go run *.go
```

> Wildcards will expand correctly in these environments.

### 3. **Use `go run .`** (Recommended for Modules)
If you're using Go modules and inside a directory with `package main`:
```bash
go run .
```

> This compiles and runs all Go files in the directory with `package main`.

---

## ✅ Example Folder Structure

```
myapp/
├── go.mod
├── main.go
├── helper.go
└── config.go
```

### To Run:
```bash
cd myapp
go run .
```

---

## ⚠️ Additional Notes

- Ensure only one `main()` function exists across `.go` files.
- You must have `package main` in at least one file with a `func main()`.

```go
// main.go
package main

import "fmt"

func main() {
  fmt.Println("Hello, Laado!")
}
```

---

