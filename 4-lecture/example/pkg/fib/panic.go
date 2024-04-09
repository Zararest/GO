//go:build noPanic

package fib

func init() {
	panicVar = false
}
