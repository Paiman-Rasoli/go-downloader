package main

import (
	"fmt"

	"github.com/Paiman-Rasoli/cmd-downloader/lib"
)

var pl = fmt.Println

func main(){
	pl("Press (q) for exit from project..\n")
	pl("Enter the URL of your file: (e.g. https://abc.com/a.png)")
	
	lib.Start()
}