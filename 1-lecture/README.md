# GO general
Folders:
- `cmd` - entery point of the app
- `pkg` - folder with packages that are visible from the outside
- `internal` - internal packages

## Start
First you should do is to initialize project:

```
go mod init `project-name`
```

After this you will see `go.mod`, which should be in your root directory.

## General program structure
At the beginning of each Go source file, you declare the package to which the file belongs.
This mechanism allows hiden certain things from other packages.

The names started from the lowercase letter are visible only within the current package. 
For example, in order to get formated I/O you can import "fmt" package:

```go
package main

import ("fmt")  // equivalent to import "fmt"

func main() {
  fmt.Println("Hi, Mark") // Println has the capital first letter
}
```
`main` package is special.
It should have `main` function which is an entry point for the program.

You also can import packages from your own module. 
For example, this is a structure of the GUI calculator module:
```
calculator/
|-- cmd/
|   |-- calculator/
|       |-- main.go
|
|-- internal/
|   |-- parser/
|       |-- parser.go
|
|   |-- gui/
|       |-- gui.go
|
|-- go.mod
|-- go.sum

```

Import means `import file` and after that you can use modules from it:

```go
package main

import "test/internal/GUI"

func main() {
	gui.Run()
}
```

## External modules
There is a CLI for package installation.
In order to add dependency from some package you should run:

```
go get github.com/rajveermalviya/gamen
```

Now you can add modules from the direct github links:
```
import jr "github.com/andygrunwald/go-jira"
```

`go.sum` - has information about dependencies (like conan.lock)

## Run
```
go mod tidy           #installs deps
go run ./cmd/main.go  #builds and runs without binary
```

```
cd build
go build -o project ../cmd/main.go #builds go file
./project
```

### Notes
- panic() - function for abortion
- each variable should be used (even err values!)
- `defer` - useful keyword which calls function after it after the current function is executed
