package go_ansi

import (
	"fmt"
	"os"
	"strconv"
)

var writer = os.Stdout

func ResetColor() {
	fmt.Fprintf(writer, "\x1B[0m")
}

func ChangeColor(fgcolor Color, fgbright bool, bgcolor Color, bgbright bool) {
	if fgcolor == NONE && bgcolor == NONE {
		return
	}
	//前面加0能够清除去前面的后景色
	  var ansi = []byte("\x1B[")
      if fgcolor != NONE {
		  text := []byte(strconv.Itoa(int(30+fgcolor-BLACK))+";")
		  if fgbright {
		  	text = append(text, "1;"...)
		  }
		  ansi = append(ansi, text...)
	  }else {
		  text := []byte(strconv.Itoa(int(30+WHITE-BLACK))+";")
		  ansi = append(ansi, text...)
	  }
	  if bgcolor != NONE {
		  text := []byte(strconv.Itoa(int(40+bgcolor-BLACK))+";")
		  if bgbright {
			  text = append(text, "1;"...)
		  }
		  ansi = append(ansi, text...)
	  }else{
		  text := []byte(strconv.Itoa(int(40))+";")
		  ansi = append(ansi, text...)
	  }
	  ansi = ansi[:len(ansi)-1]
	 // fmt.Println(string(ansi[1:]))
	  ansi = append(ansi, "m"...)
	  fmt.Fprintf(writer, string(ansi))
}