package main

import (
	"context"
	"log"
	"os"

	"github.com/kenzo0107/sendgrid"
)

func main() {
	if err := handler(); err != nil {
		log.Fatal(err)
	}
}

func handler() error {
	apiKey := os.Getenv("SENDGRID_API_KEY")

	c := sendgrid.New(apiKey, sendgrid.OptionDebug(true))
	key, err := c.GetAPIKey(context.TODO(), "dummy_api_key_id")
	if err != nil {
		return err
	}
	log.Printf("api key: id: %s name: %s\n", key.ApiKeyId, key.Name)

	return nil
}
