# Go Modules

## What is a Module?

A **module** = one or more related Go packages, versioned and distributed together as a single unit. It's defined by a `go.mod` file at its root (declares module path, Go version, dependencies). Can be an app or a reusable library.

## Multiple `main` Packages

- One module **can** contain multiple `main` packages.
- Each `main` package must live in its **own directory** → compiles to its **own binary**.
- Common convention: put them under `cmd/`.

```
myproject/
├── go.mod
├── cmd/
│   ├── server/main.go   // package main → "server" binary
│   └── cli/main.go      // package main → "cli" binary
└── internal/
    └── helper/helper.go // shared logic
```

```bash
go build ./cmd/server
go build ./cmd/cli
```

## One Package Per Directory

- A directory = one package. All `.go` files in it must declare the same package name.
- **No sub-packages exist in Go.** A nested folder is just a separate, unrelated package — no parent/child relationship, no shared visibility.

```
greetings/
├── greetings.go     // package greetings
└── formal/
    └── formal.go    // package formal (independent, must be imported separately)
```

## The One Exception: `_test` Packages

A directory can hold **two** package names, but only this pair:

| File                | Package                     | Access                                              |
| ------------------- | --------------------------- | --------------------------------------------------- |
| `greetings.go`      | `greetings`                 | —                                                   |
| `greetings_test.go` | `greetings` (internal)      | Can see unexported identifiers                      |
| `greetings_test.go` | `greetings_test` (external) | Only exported/public API, like any outside importer |

External test example:

```go
package greetings_test

import (
	"testing"
	"modules/greetings"
)

func TestHello(t *testing.T) {
	got := greetings.Hello("World")
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

---

## Summary

1. **Module** = versioned collection of packages, defined by `go.mod`.
2. A module can hold **multiple `main` packages**, but each needs its own directory → own binary.
3. **Directory = package.** No sub-package inheritance — nesting is just physical location, not logical relationship.
4. **Only exception** to "one package per directory": `package foo` + `package foo_test` may coexist, for external (black-box) testing.
