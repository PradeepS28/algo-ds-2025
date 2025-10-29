package patterns

import "fmt"

func Pattern17(n int) {
	for i := 1; i <= n; i++ {
		var ch rune = 'A'
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("%c", ch)
			mid := (2*i - 1) / 2
			if j > mid {
				ch--
			} else {
				ch++
			}
		}
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
