package dp

/**
BM62 斐波那契数列
快速幂
递归
描述
大家都知道斐波那契数列，现在要求输入一个正整数 n ，请你输出斐波那契数列的第 n 项。
斐波那契数列是一个满足 fib(x)=\left\{ \begin{array}{rcl} 1 & {x=1,2}\\ fib(x-1)+fib(x-2) &{x>2}\\ \end{array} \right.fib(x)={
1
fib(x−1)+fib(x−2)
​

x=1,2
x>2
​
  的数列
数据范围：1\leq n\leq 401≤n≤40
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n) ，本题也有时间复杂度 O(logn)O(logn) 的解法

输入描述：
一个正整数n
返回值描述：
输出一个正整数。
*/

func Fibonacci(n int) int {
	// res := make([]int, n)
	// return Fibonacci1(n, res)
	// return FibonacciDP(n)
	return FibonacciDP2(n)
}

func Fibonacci1(n int, res []int) int {
	if n == 1 || n == 2 {
		res[n] = 1
		return 1
	}
	if res[n] != 0 {
		return res[n]
	}
	l := Fibonacci1(n-1, res) + Fibonacci1(n-2, res)
	res[n] = l
	return l
}

func FibonacciDP(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func FibonacciDP2(n int) int {
	a, b, c := 1, 1, 1
	for i := 3; i <= n; i++ {
		c, a, b = a+b, b, a+b
	}
	return c
}
