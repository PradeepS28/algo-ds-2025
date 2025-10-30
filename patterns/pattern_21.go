package patterns

import "fmt"

func Pattern21(n int) {
	for i := 1; i <= n; i++ {
		if i == 1 || i == n {
			for j := 1; j <= n; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j <= n; j++ {
				if j == 1 || j == n {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
