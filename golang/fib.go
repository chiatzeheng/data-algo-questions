package main

func fib(n int) int {

	if n == 1 || n == 0 {
		return n
	} else {
		a := fib(n-1) + fib(n-2)
		return a

	}

}
