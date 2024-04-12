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
### Pop method
It have the same core as the push, but thread returns element only if it push the head.

## Fun fact
Loads in this algorithm SHOULD be atomic, since compiler isn't allowed to change their order. 

Tail can be moved only forward. Head can be moved only forward. Tail should be after head.
Therefore we should read them in this order:
```go
head := atomic.LoadPointer(&queue.Head)
tail := atomic.LoadPointer(&queue.Tail)
```
Otherwise:
```go
tail := atomic.LoadPointer(&queue.Tail)
//other thread moves head and now head is after the teil
head := atomic.LoadPointer(&queue.Head)
// broken invariant
```