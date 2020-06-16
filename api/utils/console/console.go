package console

import (
	"encoding/json"
	"log"
)

func Pretty(data interface{})  {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(bytes))
}