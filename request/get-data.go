package request

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)


func printDetailsAndGetType(URL string) string {
	httpClient := &http.Client{Timeout: 15 * time.Second}
	resp, err := httpClient.Head(URL)
	if err != nil {
		log.Fatalf("error on HEAD request: %s", err.Error())
	    }
	
	var fileSizeInMB float64 = math.Abs((float64(resp.ContentLength) / 1024) / 1024)
	var fileType string = getType(resp.Header.Get("Content-Type"))
	fmt.Printf("\n { FileSize : %.2f MB  , fileType : %s }\n", fileSizeInMB,fileType)
	return fileType
}

func generateFileName() string {
	seedSec := time.Now().Unix()
	rand.Seed(seedSec)
	randomValue := rand.Intn(1000) + 4000
	alphabets := "abcdefjhigklmnopqrstuvwxyz"
	var fileName string
	for i := 0; i < 6; i++{
		randomIndex := rand.Intn(26) // 0 - 25
            fileName += string(alphabets[randomIndex])
	}
	return fileName + strconv.Itoa(randomValue)
}
	


func FetchData(URL string){
	fileType := printDetailsAndGetType(URL)
	// Start fetching the content
	mySpinner := lunchSpinner()
	mySpinner.Start()
	resp, err := http.Get(URL)
	if err != nil{
		log.Fatal(err)
	}
      defer resp.Body.Close()
      out, err := os.Create(generateFileName()+"."+fileType)
      if err != nil {
        log.Fatal(err)
      }   
     defer out.Close()
     io.Copy(out, resp.Body)
     fmt.Println("\n\n  Your file has been downloaded. 😀\n\n")
}
