// Popcount counts number of set bits using x&(x-1).
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	res := 0
	for x != 0 {
		res += 1
		x &= (x - 1)
		// fmt.Println(res, x)
	}
	fmt.Println(res)
}
