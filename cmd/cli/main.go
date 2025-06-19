package main

import (
	"fmt"
	"os"
)

func main() {
	p, _ := os.Getwd()
	fmt.Printf("%s/cmd/cli", p)

}
