// Binary euler112 solves Euler Project question 112 using a naive brute force algorithm.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isInc(numStr string) bool {
	for i := 0; i < len(numStr)-1; i++ {
		if numStr[i+1] < numStr[i] {
			return false
		}
	}
	return true
}

func isDec(numStr string) bool {
	for i := 0; i < len(numStr)-1; i++ {
		if numStr[i+1] > numStr[i] {
			return false
		}
	}
	return true
}

func isBouncy(num int) bool {
	numStr := strconv.Itoa(num)
	return !isInc(numStr) && !isDec(numStr)
}

// bouncy returns num where the percentage of the bouncy number from 1 to num (inclusive)
// equals to p.
func bouncy(p int) int {
	num := 1
	for bouncyCount := 0; ; num++ {
		if isBouncy(num) {
			bouncyCount++
		}
		if bouncyCount*100/num == p {
			break
		}
	}
	return num
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Length of os.Args must be 2.")
		os.Exit(1)
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(bouncy(i))
}
