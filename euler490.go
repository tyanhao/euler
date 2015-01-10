// Binary euler490 is a naive implementation of the f function for Euler Project question 490.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/math"
)

func f(n int) int {
	stones := make([]bool, n)
	return fRecursive(0, stones, 0)
}

// fRecursive is a recursive function to compute f.
// curPos is the current index of stones the frog is at.
// stones[n] is true if the frog had already landed on the n-th stone.
// stoneCount is the number of stones the frog has landed before this, excluding curPos.
func fRecursive(curPos int, stones []bool, stoneCount int) int {
	stoneCount++
	if curPos == len(stones)-1 {
		if stoneCount == len(stones) {
			return 1
		} else {
			// If we are at the last stone but the frog has not landed on every other stone, this is an
			// invalid solution.
			return 0
		}
	}
	stones[curPos] = true
	sum := 0
	upper := math.Min(len(stones)-1, curPos+3)
	lower := math.Max(0, curPos-3)
	for i := lower; i <= upper; i++ {
		if stones[i] == false {
			sum += fRecursive(i, stones, stoneCount)
		}
	}
	stones[curPos] = false
	return sum
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
	fmt.Println(f(i))
}
