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
	key, err := c.UpdateAPIKeyName(context.TODO(), "dummy_api_key_id", &sendgrid.InputUpdateAPIKeyName{
		Name: "dummy-rewrite-name",
	})
	if err != nil {
		return err
	}
	log.Printf("api key: %#v\n", key)

	return nil
}
