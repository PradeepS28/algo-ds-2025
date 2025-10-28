package patterns

import "fmt"

func Pattern10(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := n; i > 1; i-- {
		for j := i - 1; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
