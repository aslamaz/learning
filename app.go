package main

import "fmt"

func main() {
	var sum int
	var a [5]int = [5]int{1, 2, 3, 4, 6}
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	avg := float64(sum) / float64(len(a))
	fmt.Println(avg)

}
func countDigits(num int) int {
	var count int
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}
func reverse(num int) int {
	var rev int
	var d int
	for num > 0 {
		d = num % 10
		num = num / 10
		rev = rev*10 + d
	}
	return rev
}
func isPalindrome(number int) bool {
	rev := reverse(number)
	if number == rev {
		return true
	} else {
		return false
	}
}
func isArmstrong(number int) bool {

	var d int
	var sum int
	numCopy := number
	for number > 0 {
		d = number % 10
		number = number / 10
		sum = sum + (d * d * d)
	}

	if sum == numCopy {
		return true
	} else {
		return false
	}

}
func isNarcissistic(number int) bool {

	var d int
	var sum int
	numCopy := number
	for number > 0 {
		d = number % 10
		number = number / 10
		sum = sum + (d * d * d * d)
	}

	if sum == numCopy {
		return true
	} else {
		return false
	}

}
func printTriangle(n int32) {

	var i, j int32
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i < j {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

}
func isPrime(number int) bool {
	var count int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}
func printUpToN(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(i)
	}
}
func printEvenNumber(number int) {
	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
func printOddNumber(number int) {
	for i := 0; i <= number; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}

	}
}
func primeNumber(number int) {
	for i := 1; i <= number; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}

	}
}
func fibonacci(number int) {
	a := 0
	b := 1
	var c int
	for i := 0; i < number; i++ {
		c = a + b
		fmt.Println(a)
		a = b
		b = c
	}
}
func fib(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fib(number-1) + fib(number-2)
	}

}
