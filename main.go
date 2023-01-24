package main

import (
	"fmt"
	"js-strings/config"
	"js-strings/enums"
)

func main() {
	c := config.New()

	stringEnums := enums.New(c)

	fmt.Println(stringEnums)

}
