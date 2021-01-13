package main

import "fmt"

func main()  {
	for i := 1; i <= 7; i++ {
        for j := i; j <= 7; j++ {
			fmt.Print(i)
		}
		fmt.Println()
    }
	
}