package patterns

import "fmt"

func Pattern2(in int) {
	for i := 1; i <= in; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
