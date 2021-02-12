package main

import (
	"fmt"

	models "../../internal"
)

func main() {
	if err := models.ValidateInput(); err != nil {
		fmt.Print(err)
		return
	}
	instructions, err := models.ValidateInstructions()
	if err != nil {
		fmt.Print(err)
		return
	}
	result := models.CheckSolution(instructions)
	fmt.Println(result)
}
