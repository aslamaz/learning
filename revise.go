package main

import (
	"fmt"
	"strings"
)

func main() {
	var country string
	var array = []string{"india:delhi", "srilanka:colombo", "bangladesh:dhaka"}
	m := make(map[string]string)

	fmt.Scanf("%s", &country)
	for _, v := range array {
		p := strings.Split((v), ":")
		m[p[0]] = p[1]
	}
	fmt.Println(m[country])
}
func mainP() {
	var num int
	var str [100]string
	var rslt []string
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	fmt.Println("enter the adress")
	for i := 0; i < num; i++ {
		fmt.Scanf("%s", &str[i])
	}
	for i := 0; i < num; i++ {
		v := str[i]

		rslt = append(rslt, strings.ToUpper(v))
	}
	fmt.Println(rslt)
}
func mainN() {
	arr := []string{"aslam", "azees", "muvattupuzha", "Ernakulam"}
	var result []string
	for _, v := range arr {

		result = append(result, strings.ToUpper(v))
		fmt.Println(result)
	}

}
func mainM() {
	var data string
	//array := []string{"paipra", "muvattupuzha", "ernakulam"}
	data = "aslamAzees"
	fmt.Println(strings.Fields(data))

}

func mainL() {
	var array = []int{1, 2, 3, 1, 4, 5, 2, 2, 7, 1}
	m := make(map[int]int)
	for _, v := range array {
		m[v]++
	}
	for num, count := range m {
		fmt.Printf("%d is found %d times\n", num, count)
	}

}
func mainK() {
	var num int
	fmt.Println("enter any integer number")
	fmt.Scanf("%d", &num)
	bubble := make([]int, num)
	fmt.Printf("enter the %d numbers", num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &bubble[i])
	}
	for i := 0; i < num-1; i++ {
		for j := 0; j < num-i-1; j++ {
			if bubble[j] > bubble[j+1] {
				bubble[j], bubble[j+1] = bubble[j+1], bubble[j]
			}
		}
	}
	fmt.Println(bubble)
}
func mainJ() {
	var num int
	var number [100]int
	var sum int
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	fmt.Printf("enter the %d numbers", num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &number[i])
	}
	for i := 0; i < num; i++ {
		sum = sum + number[i]
	}
	fmt.Println(sum)
}
func mainI() {
	var num int
	var array [100]int
	fmt.Println("enter the numbers")
	fmt.Scanf("%d", &num)
	fmt.Printf("enter the %d numbers", num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &array[i])
	}
	for i := 0; i < num; i++ {
		fmt.Println(array[i])
	}
}
func mainH() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	a := 0
	b := 1
	var c int
	for i := 1; i <= num; i++ {
		c = a + b
		fmt.Println(a)
		a = b
		b = c
	}
}
func mainG() {
	var num int
	var result int = 1
	fmt.Println("enter the number")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		result = result * i
	}
	fmt.Println(result)
}
func mainF() {
	var num int
	var d int
	var sum int
	fmt.Println("enter any integer number")
	fmt.Scanf("%d", &num)
	lastTwoDigits := num % 100
	for lastTwoDigits > 0 {
		d = lastTwoDigits % 10
		lastTwoDigits = lastTwoDigits / 10
		sum = sum + d
	}
	if sum == 3 || sum == 7 {
		fmt.Println("lucky winner")
	} else {
		fmt.Println("unlucky")
	}
}
func mainE() {
	var num int
	var count int
	fmt.Println("enter any integer number")
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
func maind() {
	var num int
	var sum int
	var d int
	fmt.Println("enter any integer number")
	fmt.Scanf("%d", &num)
	numcopy := num
	for num > 0 {
		d = num % 10
		num = num / 10
		sum = sum + (d * d * d)
	}
	if numcopy == sum {
		fmt.Println("this is armstrong")
	} else {
		fmt.Println("this is not armstrong")
	}

}

func c() {
	var num int
	var d int
	var rev int
	fmt.Println("enter any integer number")
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
func b() {
	var num int
	var d int
	var rev int
	fmt.Println("enter any integer number")
	fmt.Scanf("%d", &num)
	for num > 0 {
		d = num % 10
		num = num / 10
		rev = rev*10 + d
	}
	fmt.Println(rev)
}
func a() {
	var num int
	var count int
	fmt.Println("enter the integer number")
	fmt.Scanf("%d", &num)
	for num > 0 {
		num = num / 10
		count++
	}
	fmt.Println(count)
}
