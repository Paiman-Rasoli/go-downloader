package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

var pl = fmt.Println

type Data struct {
	URL string
}

func (data *Data) setURL(userInput string)   {
	var trimmedData string = strings.TrimSpace(userInput)
	if strings.ToLower(trimmedData) == "q"{
		pl("\nBye, see you later... :)")
		os.Exit(0)
	}
	data.URL = trimmedData
}	

func (data *Data) ValidateURL()  bool {
	return WEBSITE_REGEX.MatchString(data.URL)
}

func readFromCMD() (string, error){
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	return userInput , err
}

func GetURL() string{
	var data Data
	var err error
	var userInput string

	pl("Enter the URL of your file "+ color.InGreen("(e.g. https://abc.com/a.png)"))
	userInput, err = readFromCMD()
	for err != nil {
		userInput, err = readFromCMD()
	}
	data.setURL(userInput)
	isValidURL := data.ValidateURL()
	for !isValidURL{
		pl(color.InRed("Invalid URL")+" Please enter a valid URL: ")
		userInput, err =readFromCMD()
		data.setURL(userInput)
		isValidURL = data.ValidateURL()
	}
	return data.URL
}
