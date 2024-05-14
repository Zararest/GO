# REST API
The main idea of REST API is a sateless architecture in client-server applications.
In this task http server will be implemented based on such API.

## Task
Server should contain phone numbers and user IDs. It should store them in files or DB.
The main operation it should implement are CRUD:
- Create
- Read
- Update
- Delete

## Go server implementation
In order to work with http requests Go has a package with http server interface implementation.

```go
import "net/http"
```
The main types are:
- http.Server - objecct that listens a socket and accepts new connections
- http.ResponseWriter - object that returns request from handler
- http.Request - object with information about request

Another package for efficient http server is router. 
It is needed for routing requests.
For example, this allows to specicy only handle function and regexp for particular request.

```go
import "github.com/gorilla/mux"
```

### General pipeline

#### Main goroutine
```go
// From http package
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func newHandler() *mux.Router {
  r := mux.NewRouter()

  // function that will handle perticular request
  route := r.HandleFunc("path", nil /*func(http.ResponseWriter, *http.Request)*/ )
  // route - path for a request

  // adding matches for http request
  route.Methods(http.MethodOptions, http.MethodGet)
}

func main() {
  handler := newHandler()
  server := http.Server{Addr : "host:8080", 
                        Handler : handler}
  // infinite loop 
  if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
```

## Note
Function can return `nil` if return value is a pointer or an interface!!!

