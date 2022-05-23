package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/broker-service/cmd/models"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"os"
)

var JsonConfig map[string]models.ServiceModel

func LoadServicesJson() {
	jsonFile, err := os.Open("services.json")

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

func IsPathValid(service string, ctx *fiber.Ctx) error {

	path := ctx.Params("path")

	cfg := JsonConfig[service]

	if cfg.BaseUrl != "" {
		if cfg.Path != path {
			return errors.New(fmt.Sprintf("Unregistered path /%s", path))
		}
	}

	return nil
}
