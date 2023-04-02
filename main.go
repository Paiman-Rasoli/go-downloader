package main

import (
	"fmt"

	"github.com/Paiman-Rasoli/cmd-downloader/lib"
	"github.com/Paiman-Rasoli/cmd-downloader/request"
	"github.com/TwiN/go-color"
)

var pl = fmt.Println

func main(){
	pl("\nEnter"+ color.InBlue("(q)") +" for exit from application..\n")	
	URL := lib.GetURL()
	request.FetchData(URL)
}