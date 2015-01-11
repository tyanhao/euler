// Binary euler169 solves Euler Project question 169.
package main

import (
	"fmt"
	"math/big"
	"os"
)

// key is the format for the DP table.
type key struct {
	// num is the number we are interested in.
	Num string
	// exp is the largest power of 2 that can be used to form num. 2^k can only be used at most
	// twice for each k.
	Exp int
}

// dpTable maps key to the number of different ways key.Num can be expressed as a sum of powers of
// 2.
var dpTable = make(map[key]int)

func genKey(n *big.Int) string {
	return string(n.Bytes())
}

func f(num *big.Int, exp int) int {
	numBitLen := num.BitLen()
	// Invalid inputs, will not be able to form n with these constraints.
	if num.Cmp(big.NewInt(0)) < 0 || exp < numBitLen-2 {
		return 0
	}
	// Reduce the space complexity: k > n.BitLen() is equivalent to k == n.BitLen().
	if exp > numBitLen {
		exp = numBitLen
	}
	if exp == 0 {
		// Special cases for base conditions.
		if num.Cmp(big.NewInt(2)) <= 0 {
			return 1
		}
		// Will not be able to form n for larger n.
		return 0
	}
	dpKey := key{genKey(num), exp}
	if result, present := dpTable[dpKey]; present {
		return result
	}

	TwoExp := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(exp)), nil) // 2^exp
	// Find all ways without using 2^exp.
	caseA := f(num, exp-1)
	// Find all ways using 2^exp once: f(num - 2^exp, exp-1).
	caseB := f(new(big.Int).Sub(num, TwoExp), exp-1)
	// Find all ways using 2^exp twice: f(num - 2*2^exp, exp-1).
	caseC := f(new(big.Int).Sub(num, new(big.Int).Add(TwoExp, TwoExp)), exp-1)

	result := caseA + caseB + caseC
	dpTable[dpKey] = result
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Length of os.Args must be 2.")
		os.Exit(1)
	}
	i := new(big.Int)
	if _, ok := i.SetString(os.Args[1], 10); !ok {
		fmt.Printf("Invalid input: %v", os.Args[1])
		os.Exit(1)
	}
	exp := i.BitLen() - 1
	fmt.Printf("Result: %v\n", f(i, exp))
}
