# Go Course

This repository is a structured, self-paced learning path for mastering the **Go programming language**.  
It is organized into three progressive folders to separate beginner, intermediate, and advanced topics.

---

## 📁 Project Structure

```

go-course/
├── basic/          # Go fundamentals: variables, constants, data types, functions, control flow, arrays, slices, maps, structs, interfaces, imports, etc.
├── intermediate/   # (WIP) 
├── advanced/       # (WIP)
├── go.mod          # Go module definition
└── main.go         # Entry point (optional runner or aggregator)

```

---

## 🚀 How to Run Code

### Run the default entry point

```bash
go run main.go
```

### You can also compile instead of running:

```bash
go build ./...
```

### Format your code manually (like Prettier)

```bash
gofmt -w .
```

---

## 🛠️ Using the Makefile

You can use the included `Makefile` to simplify your workflow:

| Command           | What it does                                      |
|-------------------|---------------------------------------------------|
| `make run`        | Run main.go                    |
| `make fmt`        | Format all `.go` files in-place (`gofmt -w`)     |
| `make check`      | Run all CI checks (build, tidy, vet, formatting) |
| `make build`      | Build all packages in the repo                    |
| `make vet`        | Run `go vet` on all packages                      |

---

## ✅ Prerequisites

* Go `1.24.4` or higher
* Basic terminal/CLI usage

---
