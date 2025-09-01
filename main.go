package main

import (
	"fmt"

	"igaleksus.com/spamMasker/internal/masker"
)

func main() {
	var stroka string = "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	fmt.Println(masker.CatchHttp(stroka))
}
