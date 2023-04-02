package request

import "strings"

func getType(contentType string) string{
	if !strings.Contains(contentType,"application"){
		return contentType
	}

	spliced := strings.Split(contentType,"application/")
	var extension string
	if len(spliced) < 2 {
		extension = "undefined"
	}else{
		extension = spliced[1]
	}
	
	switch extension{
	case "octet-stream": return "binary"
	default: return extension
    }
}