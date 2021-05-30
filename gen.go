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

	componentsArr := gjson.Get(realm, "components.org.keycloak.services.clientregistration.policy.ClientRegistrationPolity")
	println(componentsArr.String())

}
