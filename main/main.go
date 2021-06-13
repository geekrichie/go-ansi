package main

import (
	"fmt"
	ct "go-ansi"
)

func main() {
	ct.ResetColor()
	ct.ChangeColor(ct.RED,true, ct.BLACK, false)
	fmt.Println("red forecolor and black bgcolor")
	ct.ChangeColor(ct.NONE,false, ct.GREEN, false)
	fmt.Print("white forecolor and green bgcolor")
	ct.ResetColor()
	fmt.Println("NORMAL text")
	ct.ChangeColor(ct.RED, false, ct.NONE, true)
	fmt.Println("red forecolor and blue bgcolor")
	ct.ChangeColor(ct.NONE, false, ct.BLUE, true)
	fmt.Print("green forecolor and blue bgcolor")
	ct.ChangeColor(ct.GREEN, false, ct.CYAN,false )
	fmt.Println("green forecolor and blue bgcolor")
	ct.ResetColor()
}
