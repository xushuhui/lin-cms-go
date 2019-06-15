package utils

import (
	"encoding/json"
	"fmt"
)

func BindJson(body []byte, obj interface{}) {
	var err error
	if err = json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}
}
