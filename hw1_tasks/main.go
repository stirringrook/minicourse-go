package main

import (
	"fmt"
	"math"
	"strings"
)

func helloWorld() {
	fmt.Println("Hello, World!")
}

func addTwoInts() {
	var x int
	var y int
	fmt.Print("Enter x: ")
	fmt.Scanln(&x)
	fmt.Print("Enter y: ")
	fmt.Scanln(&y)
	fmt.Println(x + y)
}

func oddEven() {
	var x int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&x)
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func maxOfThree() {
	var x int
	var y int
	var z int
	fmt.Print("Enter x: ")
	fmt.Scanln(&x)
	fmt.Print("Enter y: ")
	fmt.Scanln(&y)
	fmt.Print("Enter z: ")
	fmt.Scanln(&z)
	fmt.Println(max(x, y, z))
}

func factorial() {
	var x int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&x)
	ans := 1
	for x > 0 {
		ans = ans * x
		x -= 1
	}
	fmt.Println(ans)
}

func isVowel() {
	var x string
	fmt.Print("Enter a character: ")
	fmt.Scanln(&x)
	switch x {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Indeed a vowel")
	}
	fmt.Println("Not a vowel")
}

func isPrime(x int) bool {
	upperLimit := math.Sqrt(float64(x))
	for i := 2; i < int(upperLimit)+1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func genPrimes() {
	var x int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&x)
	primes := make([]int, 0, 10000)

	for i := 0; i < x; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	for _, x := range primes {
		fmt.Println(x)
	}
}

func reverseString() {
	var s string
	fmt.Print("Enter string: ")
	fmt.Scanln(&s)

	ans := []rune(s)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	fmt.Println(string(ans))
}

func sumOfArray() {
	len := 0
	fmt.Print("Enter the lenght of an array")
	fmt.Scanln(&len)
	input := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scanf("%f", &input[i])
	}
	sum := 0
	for i := 0; i < len; i++ {
		sum += input[i]
	}
	fmt.Println(sum)
}

type Rectangle struct {
	width, height int
}

func (a Rectangle) getArea() int {
	return a.height * a.width
}

func rectangleTester() {
	w := 0
	h := 0
	fmt.Println("Enter width:")
	fmt.Scanln(&w)
	fmt.Println("Enter height:")
	fmt.Scanln(&h)
	var rect Rectangle
	rect.width = w
	rect.height = h
	fmt.Printf("Area of a rectangle: %d", rect.getArea())

}

func celciusToFarenheit() {
	var degrees float64
	fmt.Print("Enter temperature in C: ")
	fmt.Scan(&degrees)
	ans := degrees*9/5 + 32
	fmt.Printf("Temperature in F: %.2f\n", ans)
}

func countdown() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}

func strLenNoBuiltin() {
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)
	length := 0
	for range s {
		length++
	}
	fmt.Printf("String length: %d", length)
}

func contains(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func isInSlice() {
	var n, elem int
	fmt.Print("Enter slice length: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter slice: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter a numver to find: ")
	fmt.Scan(&elem)
	if contains(arr, elem) {
		fmt.Println("Number found")
	} else {
		fmt.Println("Number not found")
	}
}

func sliceAvg() {
	var n int
	fmt.Print("Enter slice length: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter slice: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}

	fmt.Printf("Slice average: %.2f", float64(sum)/float64(len(arr)))
}

func multTable() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func isPalindrome() {
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)
	flag := true
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			flag = false
		}
	}

	if flag {
		fmt.Println("String is a palindrome")
	} else {
		fmt.Println("String is not a palindrome")
	}
}

func getMinMax() {
	var n int
	fmt.Print("Enter slice length: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter slice: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
}

func deleteByIdx() {
	var n, index int
	fmt.Print("Enter slice length: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter slice: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter index to remove: ")
	fmt.Scan(&index)
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Println("Resulting slice:", arr)
}

func linearSearch() {
	var n, element int
	fmt.Print("Enter slice length: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter slice: ")
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Element to search for: ")
	fmt.Scan(&element)

	index := -1
	for i, v := range arr {
		if v == element {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("Element found at %d.", index)
	} else {
		fmt.Println("Element not found")
	}
}

func main() {
	// to run a particular task --> uncomment a line below
	//
	//
	helloWorld()
	// addTwoInts()
	// oddEven()
	// maxOfThree()
	// factorial()
	// isVowel()
	// genPrimes()
	// reverseString()
	// sumOfArray()
	// rectangleTester()
	// celciusToFarenheit()
	// countdown()
	// strLenNoBuiltin()
	// isInSlice()
	// sliceAvg()
	// multTable()
	// isPalindrome()
	// getMinMax()
	// deleteByIdx()
	// linearSearch()
}
