package fib

var panicVar = true

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func FibMain() {
	if panicVar {
		panic("AAAAAA from fib")
	}
	n := 10 // Change the value of n to generate Fibonacci sequence of different lengths
	seq := fibonacci(n)
	for i := 0; i < n; i++ {
		println(seq[i])
	}
}
