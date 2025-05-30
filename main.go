package main

import (
	"fmt"
	"github.com/aminzohrabifar/iranian-validation-go/src/mobile"
)

func main() {
	err, op := mobile.GetMobileNumberOprator("  sww 916ddw 123dw 4567  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(op)
}
