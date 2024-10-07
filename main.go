package main

import "fmt"

func numIsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		if numIsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
