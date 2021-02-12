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
	log := models.PushSwap()
	fmt.Printf("%v", log.Msgs)
}
