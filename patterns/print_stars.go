package patterns

import "fmt"

func PrintStars(in int) {
	for i := 0; i < in; i++ {
		for j := 0; j < in; j++ {
			fmt.Print("x")
		}
		fmt.Println("")
	}
}
