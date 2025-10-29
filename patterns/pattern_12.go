package patterns

import "fmt"

func Pattern12(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i; j <= 2*n-i-1; j++ {
			fmt.Print(" ")
		}

		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
