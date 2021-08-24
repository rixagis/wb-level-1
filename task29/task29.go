package main

import (
	"fmt"
	"math/big"
)

func main(){
	var (
		a = big.NewInt(384354638384368684)
		b = big.NewInt(5438345735005435)
		c = big.NewInt(0)
	)

	fmt.Printf("a = %d, b = %d\n", a, b)
	c = c.Add(a, b)
	fmt.Printf("a + b = %d\n", c)
	c = c.Sub(a, b)
	fmt.Printf("a - b = %d\n", c)
	c = c.Mul(a, b)
	fmt.Printf("a * b = %d\n", c)
	c = c.Div(a, b)
	fmt.Printf("a / b = %d\n", c)


}