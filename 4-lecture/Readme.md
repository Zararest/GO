# Project build
Go compiler allows to dump AST and SSA froms of the program:
```bash
go tool compile -w ./pkg/fib/fib.go       #AST for a file without deps
go build -gcflags=-w ./cmd/main.go        #AST for a project 
GOSSAFUNC=fibonacci go tool compile ./pkg/fib/fib.go  #SSA for a single function
go build -C ./cmd                   # Build with panic
go build -C ./cmd -tags no-panic    # Build without panic
```

- `go tool compile` - tool for a single file
- `go build` - driver
- `go build -gcflags=` - flags that should be passed to go tool compile when it's invoked

This tool compiles a single file, so there should be no dependencies, because it can't handle it from a box.

## Building
In order to add multiple files to main package you should add all files to the build command.
AND TAGS ARE NOT WORKING IN MAIN PACKAGE.

## Artifacts
- `fib.ast` - AST for a single file
- `main.ast` - AST for a whole project
- `fibbonachi.ssa.html` - SSA for a fibonacci function

## Questions
How does `init` function in the other file but in the same module know that var name is valid? 
How to resolve name conflicts in multi-file module?
How to deal with a whole project and tags? 
```
var panic bool

func panic() {}
```
