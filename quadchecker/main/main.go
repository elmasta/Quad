package main

import (
	"fmt"
	"io"
	"os"
	"piscine"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	s := string(bytes)
	if s != "Nothing to print" {
		answerquad := piscine.Quadchecker(s)
		fmt.Println(answerquad)
	} else {
		fmt.Println("Not a quad function")
	}
}
