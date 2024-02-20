package Recursion

func Factorial(input int) int {
	if input < 0 {
		// 可以返回一個錯誤或特定值
		return -1 // 這裡返回-1作為錯誤的示例
	}
	if input == 0 || input == 1 {
		return 1
	}
	return input * Factorial(input-1)
}

func Fibonacci(input int) int {
	if input == 0 {
		return 0
	}
	if input == 1 {
		return 1
	}
	return Fibonacci(input-1) + Fibonacci(input-2)
}

func FibonacciOptimize(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b // Update the last two Fibonacci numbers
	}
	return b
}
