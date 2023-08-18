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
	key, err := c.CreateAPIKey(context.TODO(), &sendgrid.InputCreateAPIKey{
		Name: "dummy",
		Scopes: []string{
			"user.profile.read",
		},
	})
	if err != nil {
		return err
	}
	log.Printf("api key: %#v\n", key)

	return nil
}
