package patterns

import "fmt"

func Pattern8(n int) {
	for i := n; i > 0; i-- {
		for j := n - 1; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 2*i - 1; j > 0; j-- {
			fmt.Print("*")
		}
		for j := n - 1; j >= i; j-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
