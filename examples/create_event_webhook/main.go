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
	r, err := c.CreateEventWebhook(context.TODO(), &sendgrid.InputCreateEventWebhook{
		URL:     "https://example.com",
		Enabled: false,
	})
	if err != nil {
		return err
	}
	log.Printf("%#v\n", r)

	return nil
}
