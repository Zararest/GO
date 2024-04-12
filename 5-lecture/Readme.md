# Concurrency
In Go you can call a goroutine by adding `go` before a function call:
```go
func foo() {
  fmt.Print("From thread")
}

func main() {
  go foo()
}
```

## Michel Scott lockfree queue
### Push method
This algorithm is baed on the pattern, that we should capture the tail of the queue.
If thread fails to perform 2 non-atomic operaands consequently, it tries agaim.
 