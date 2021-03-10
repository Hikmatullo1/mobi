package main

import (
		"fmt"
		"mobi/pkg/commission"
	)
func main() {

		result :=commission.Calculate(500001)
		fmt.Println(result)
}