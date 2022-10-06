package main

import "fmt"

func main() {

}
func countNumberOfOccurence() {
	var array = []int{1, 2, 3, 1, 2, 6, 8, 1}
	m := make(map[int]int)
	for _, v := range array {
		m[v]++
	}
	for num, count := range m {
		fmt.Printf("%d is found %d times\n", num, count)
	}

}
func bubbleSort() {
	var num int
	fmt.Println("enter a integer number")
	fmt.Scanf("%d", &num)
	array := make([]int, num)
	fmt.Printf("enter %d values", num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println(array)
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
func fact(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * fact(num-1)
	}
}

func factorial(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}
func productOfArray() {
	var num1 int
	var product int = 1
	var number1 [100]int
	fmt.Println("enter the values")
	fmt.Scanf("%d", &num1)
	fmt.Printf("enter %d values", num1)
	for i := 0; i < num1; i++ {
		fmt.Scanf("%d", &number1[i])
	}
	for i := 0; i < num1; i++ {
		product = product * number1[i]
	}
	fmt.Println(product)
}
func sumOfArray() {
	var n int
	var numbers [100]int
	var sum int
	fmt.Println("enter the values")
	fmt.Scanf("%d", &n)
	fmt.Printf("enter %d numbers", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	for i := 0; i < n; i++ {
		sum = sum + numbers[i]
	}
	fmt.Println(sum)
}

func array() {
	// WAP to get n values from user and print it
	var n int
	fmt.Println("enter the values of n")
	var numbers [100]int
	fmt.Scanf("%d", &n)
	fmt.Printf("enter %d numbers", n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(numbers[i])
	}

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
func luckyDraw(num int) bool {
	var d int
	var lastTwoDigits int
	var sum int
	lastTwoDigits = num % 100
	for lastTwoDigits > 0 {
		d = lastTwoDigits % 10
		lastTwoDigits = lastTwoDigits / 10
		sum = sum + d
	}
	if sum == 3 || sum == 8 {
		return true
	} else {
		return false
	}
}
func isLuckyDraw(num int) bool {
	var digit int
	var sum int
	for i := 0; i < 2; i++ {
		digit = num % 10
		num = num / 10
		sum = sum + digit
	}
	if sum == 3 || sum == 8 {
		return true
	} else {
		return false
	}
}
func grade(num int) string {
	if num >= 90 {
		return ("A")
	} else if num >= 80 {
		return ("B")
	} else if num >= 70 {
		return ("c")
	} else if num >= 60 {
		return ("D")
	} else if num >= 50 {
		return ("E")
	} else {
		return ("fail")
	}
}
func grade2(num int) string {
	switch {
	case num >= 90:
		return "A"
	case num >= 80:
		return "B"
	case num >= 70:
		return "c"
	case num >= 60:
		return "D"
	case num >= 50:
		return "E"
	default:
		return "fail"
	}
}
