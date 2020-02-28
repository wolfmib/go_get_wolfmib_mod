package main

import (
	"fmt"

	add "github.com/wolfmib/test_mod/add"
)

func main() {

	fmt.Println("[Jean]: C'est program est utiliser par testing go get module dans la github.com")
	var num1 int = 4
	var num2 int = 5
	add.MyNum(num1, num2)
	add.Show()

}
