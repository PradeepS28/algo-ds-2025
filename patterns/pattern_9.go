package patterns

import "fmt"

func Pattern9(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}

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
