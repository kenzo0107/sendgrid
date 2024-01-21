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
	template, err := c.GetTemplate(context.TODO(), "d-12345abcde")
	if err != nil {
		return err
	}

	log.Printf("template: %#v", template)

	return nil
}
