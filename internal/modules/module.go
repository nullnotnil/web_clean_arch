package modules

import (
	"fmt"
	"os"
)

func init() {
	p, _ := os.Getwd()
	fmt.Printf("%s/interal/modules", p)
}
