# Getting Started & Basics in Go

## Table of Contents

1. [Initialize a Go Module](#initialize-a-go-module)

## Initialize a Go Module

To start a new Go project, you need to initialize a Go module. A Go module is a collection of Go packages stored in a file `go.mod` at its root. The `go.mod` file defines and lists the specific Go module and versions required by the project, just like a `package.json` file in Node.js projects.

To initialize a Go module, you can run the following command in the root directory of your project:

```bash
go mod init <module-name>
```

After that, you can start by creating a new file with the `.go` extension and write your Go code.

Be sure to include the `package main` statement at the beginning of the file. As the `go run .` command compiles and runs the main package by default.

```go
go run .
```

## Import and use external packages

You can import external packages from the internet, you can look at some of the popular packages at [pkg.go.dev](https://pkg.go.dev/) and import them in your project via the import statement.

```go
import "rsc.io/quote"
```

After importing the package, you may use the `go mod tidy` command to download the package and update the `go.mod` file with the required package and version.

```bash
go mod tidy
```

Now you can use the imported package in your code.

```go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Go())
}
```

## Testing

When testing, use of the `*_test.go` pattern is required for the compiler to identify testing files, and naming them with the same name as the file you want to test is a best practice.

When writing the **test functions** it is recommended to use the `Test` prefix to identify the test functions.

Finally, to run the tests, you can use the `go test` command.

```bash
go test
```
