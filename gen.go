package main

import (
	//"github.com/tidwall/sjson"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
)

func main() {
	// We are opening the json file as string into 'json' variable
	content, err := ioutil.ReadFile("realm-template.json")
	if err != nil {
		log.Fatal(err)
	}
	realm := string(content)

	realmId := gjson.Get(realm, "id")
	println(realmId.String())
}
