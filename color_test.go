package go_ansi

import (
	"fmt"
	"os"
	"testing"
)

func TestChangeColor(t *testing.T) {
	fmt.Println(WHITE)
}
func TestChangeForeColor(t *testing.T) {
	fmt.Fprintf(os.Stdout, "\x1B[32;1;#m")
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "\x1B[32;#m")
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "\x1B[30;3;#m")
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "\x1B[30;4;#m")
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "\x1B[0m")
	fmt.Println("normal text")
	fmt.Fprintf(os.Stdout, "\x1B[40;#m")
	fmt.Println("black ground")
	fmt.Fprintf(os.Stdout, "\x1B[40;3;#m")
	fmt.Println("black ground")
	fmt.Fprintf(os.Stdout, "\x1B[#m")
	fmt.Println("normal text again")
}

func TestNewChangeForeColor(t *testing.T) {
	fmt.Fprintf(os.Stdout, "\x1B[38;5;9;1;#m")
	fmt.Fprintf(os.Stdout, "\x1B[48;5;14;1;#m")
	fmt.Println("hello world")
}


func TestColorChange(t *testing.T) {
	ChangeColor(RED,true, WHITE, false)
	fmt.Println("red forecolor and black bgcolor")
	ChangeColor(RED,false, WHITE, true)
	fmt.Println("red forecolor and black bgcolor")
	ResetColor()
	fmt.Println("NORMAL text")
	ChangeColor(RED, false, NONE, true)
	fmt.Println("green forecolor and blue bgcolor")
	ChangeColor(NONE, false, BLUE, true)
	fmt.Println("green forecolor and blue bgcolor")
	ChangeColor(GREEN, true, BLUE,false )
	fmt.Println("green forecolor and blue bgcolor")
	ResetColor()
}