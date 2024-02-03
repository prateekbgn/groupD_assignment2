package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	result := add(3, 7)
	fmt.Println("Sum:", result)
	var pnum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&pnum)
	if isPrime(pnum) {
		fmt.Printf("%d is a prime number.\n", pnum)
		} else {
		fmt.Printf("%d is not a prime number.\n", pnum)
		}

}
func add(a, b int) int {
	return a + b
}
func revString(str string) string {
	var theArray []string
	theArray = strings.Split(str, "")
	var strOutput string
	for i := len(str) - 1; i >= 0; i-- {
		strOutput += theArray[i]
	}
	return strOutput
}

func isPrime(num int) bool {
	if num < 2 {
	return false
	}
	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
		return false
		}
	}
	return true
}
