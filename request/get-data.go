package request

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func printDetails(response *http.Response) {
	var fileSizeInMB float64 = (float64(response.ContentLength) / 1024) / 1024
	var fileType string = getType(response.Header.Get("Content-Type"))
	fmt.Printf("\n { FileSize : %.2f MB  , fileType : %s }", fileSizeInMB,fileType)
}


func FetchData(URL string){
	httpClient := &http.Client{Timeout: 15 * time.Second}
	resp, err := httpClient.Head(URL)

	if err != nil {
	  log.Fatalf("error on HEAD request: %s", err.Error())
	}
	printDetails(resp)

}