package cmd

import "fmt"

func Sys(values []string) {
	// Print all values recieved
	for _, value := range values {
		fmt.Println(value)
	}

}
