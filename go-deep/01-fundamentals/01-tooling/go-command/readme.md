# The `go` Command

The `go` command is the primary CLI tool for the Go programming language. It handles everything from compiling and running code to managing dependencies and testing.

## Core Commands

**`go run`**
Compiles and runs a Go program in one step (doesn't leave a binary behind).

```bash
go run main.go
go run .          # runs the package in the current directory
```

**`go build`**
Compiles your code into an executable binary but doesn't run it.

```bash
go build main.go
go build .        # builds current package
go build -o myapp # custom output name
```

**`go install`**
Compiles and installs the binary into `$GOPATH/bin` (or `$GOBIN`), making it globally runnable.

```bash
go install
go install github.com/user/tool@latest
```

## Module & Dependency Management

**`go mod init`**
Initializes a new module (creates `go.mod`).

```bash
go mod init github.com/username/projectname
```

**`go mod tidy`**
Adds missing dependencies and removes unused ones. You'll run this constantly.

```bash
go mod tidy
```

**`go get`**
Adds or updates a dependency.

```bash
go get github.com/gin-gonic/gin
go get github.com/gin-gonic/gin@v1.9.0  # specific version
go get -u ./...                         # update all deps
```

**`go mod vendor`**
Creates a local `vendor/` folder with all dependencies bundled in.

## Testing & Quality

**`go test`**
Runs tests (files ending in `_test.go`).

```bash
go test ./...          # run all tests in project
go test -v             # verbose output
go test -cover         # show coverage
go test -run TestName  # run specific test
```

**`go vet`**
Statically analyzes code for suspicious constructs (like `printf` format mismatches).

```bash
go vet ./...
```

**`go fmt`**
Auto-formats your code to Go's standard style — everyone's code ends up looking the same.

```bash
go fmt ./...
```

## Info & Utility Commands

**`go doc`**
Shows documentation for a package or function.

```bash
go doc fmt.Println
```

**`go env`**
Prints (or sets) Go environment variables like `GOPATH`, `GOROOT`, `GOOS`, `GOARCH`.

```bash
go env
go env GOPATH
```

**`go version`**
Shows your installed Go version.

**`go clean`**
Removes build artifacts.

## Key Concepts Worth Knowing

- **`go.mod`** — declares your module name, Go version, and dependencies.
- **`go.sum`** — records checksums of dependencies for verification/security.
- **GOPATH vs Modules** — modern Go (1.16+) uses modules by default; you rarely need to think about GOPATH anymore.
- **Cross-compilation** — you can build for other platforms easily:

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux
GOOS=windows GOARCH=amd64 go build -o app.exe
```

## A Typical Workflow

```bash
go mod init myproject
# write code
go mod tidy          # pull in dependencies
go run .             # test it quickly
go test ./...        # run tests
go build             # produce final binary
```
