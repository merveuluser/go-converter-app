package main

import (
	"fmt"
	"github.com/merveuluser/go-converter-app/helper"
)

func main() {
	if !helper.Web() && !helper.Cli() {
		fmt.Println("Usage: ")
		return
	}

	if helper.Web() {
		helper.WebConversion()
	}

	if helper.Cli() {
		helper.CliConversion()
	}

}
