package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pl = fmt.Println

type Data struct {
	URL string
}

func (data *Data) setURL(userInput string)   {
	data.URL = strings.TrimSpace(userInput)
}	

func (data *Data) ValidateURL()  bool {
	return WEBSITE_REGEX.MatchString(data.URL)
}

func readFromCMD() (string, error){
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	return userInput , err
}

func Start(){
	var data Data
	var err error
	var userInput string

	userInput, err = readFromCMD()
	for err != nil {
		userInput, err = readFromCMD()
	}
	data.setURL(userInput)
	isValidURL := data.ValidateURL()
	for !isValidURL{
		pl("Invalid URL, Enter the URL of your file: (e.g. https://abc.com/a.png) ")
		userInput, err =readFromCMD()
		data.setURL(userInput)
		isValidURL = data.ValidateURL()
	}
}
