package patterns

import "fmt"

/*
Given an integer n. You need to recreate the pattern given below for any value of N. Let's say for N = 5, the pattern should look like as below:
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
**********

Print the pattern in the function given to you.
*/
func Pattern19(n int) {
	num := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= num; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= n-i+1; j++ {
			fmt.Print("*")
		}
		num = num + 2
		fmt.Println("")
	}
	num1 := n*2 - 2
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= num1; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		num1 = num1 - 2
		fmt.Println("")
	}
}
