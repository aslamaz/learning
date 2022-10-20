package main

import "fmt"

// find the minimum number of points required to make the given sum
// you have given infinite these dinominations are 1,2,5,10,50
func main() {
	var array = []int{50, 10, 5, 2, 1}
	var num int
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	for _, v := range array {
		count := num / v
		fmt.Println(v, count)
		num = num % v
	}
}
func main3() {
	var num int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	array := make([]int, num)
	fmt.Printf("enter the %d numbers", num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &array[i])
	}
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
func factorial1() {
	var num int
	var result int = 1
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		result = result * i
	}
	fmt.Println(result)
}
func lucky() {
	var num int
	var d int
	var sum int
	var lastTwoDigits int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	lastTwoDigits = num % 100
	for lastTwoDigits > 0 {
		d = lastTwoDigits % 10
		lastTwoDigits = lastTwoDigits / 10
		sum = sum + d
	}
	if sum == 3 || sum == 8 {
		fmt.Println("lucky winner")
	} else {
		fmt.Println("unlucky")
	}
}

func main2() {
	var num int
	var count int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++

		}
	}
	if count == 2 {
		fmt.Println("this is a prime number")
	} else {
		fmt.Println("this is not a prime number")
	}
}

func armstrong1() {
	var num int
	var sum int
	var d int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	numcopy := num
	for num > 0 {
		d = num % 10
		num = num / 10
		sum = sum + (d * d * d)
	}
	if numcopy == sum {
		fmt.Println("this is armstrong number")
	} else {
		fmt.Println("this is not armstrong number")
	}
}

func palindrome1() {
	var num int
	var rev int
	var d int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	numcopy := num
	for num > 0 {
		d = num % 10
		num = num / 10
		rev = rev*10 + d
	}
	if numcopy == rev {
		fmt.Println("this is palindrome")
	} else {
		fmt.Println("this is not palindrome")
	}
}

func reverse2() {
	var num int
	var rev int
	var d int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	for num > 0 {
		d = num % 10
		num = num / 10
		rev = rev*10 + d
	}
	fmt.Println(rev)
}

func counting() {
	var count int
	var num int
	fmt.Println("enter any integer numbers")
	fmt.Scanf("%d", &num)
	for num > 0 {
		num = num / 10
		count++
	}
	fmt.Println(count)
}
