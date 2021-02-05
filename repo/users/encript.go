package users

import (
	b64 "encoding/base64"
		"fmt"
)

func base64String(value string) (string){
	encodeString :=  b64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encodeString)
	return encodeString
}

