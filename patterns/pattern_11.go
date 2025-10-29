package patterns

import "fmt"

func Pattern11(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			if j%2 == 0 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Println("")
	}
}
