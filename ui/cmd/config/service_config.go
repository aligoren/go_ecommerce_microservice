package config

import (
	"encoding/json"
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/ui/cmd/models"
	"io/ioutil"
	"log"
	"os"
)

var JsonConfig map[string]models.ServiceModel

func LoadServicesJson(fileName string) {
	jsonFile, err := os.Open(fmt.Sprintf("%s.json", fileName))

	if err != nil {
		log.Fatalf("Services file couldn't open %v", err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalf("Error while closing file %v", err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &JsonConfig)
	if err != nil {
		log.Fatalf("Error while unmarshaling json data %v", err)
	}

}
