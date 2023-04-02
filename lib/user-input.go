package lib

import (
	"bufio"
	"fmt"
	"os"
)

type Data struct {
	URL string
}

func (user *Data) SetURL(input string) {
	user.URL = input
}

func readFromCMD() (string, error){
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	return userInput , err
}

func Start(){
	var data Data
	var err error
	data.URL, err = readFromCMD()
	for err != nil {
		data.URL, err = readFromCMD()
	}
	fmt.Println(data)
}
