package main

import (
	"fmt"

	"github.com/Paiman-Rasoli/cmd-downloader/lib"
)

var pl = fmt.Println

func main(){
	pl("Enter (q) for exit from application..\n")
	pl("Enter the URL of your file: (e.g. https://abc.com/a.png)")
	
	lib.Start()
}