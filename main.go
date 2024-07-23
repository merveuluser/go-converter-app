package main

import (
	"fmt"
	"github.com/merveuluser/go-converter-app/helper"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helper.GeneralUsage)
		return
	}

	mode := os.Args[1]
	switch mode {
	case "web":
		helper.WebConversion()
	case "convert":
		if len(os.Args) != 4 {
			fmt.Println(helper.CliUsage)
			return
		}
		helper.CliConversion()
	default:
		fmt.Println(helper.GeneralUsage)
	}
}
