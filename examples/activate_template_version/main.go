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
	r, err := c.ActivateTemplateVersion(context.TODO(), "d-abcde12345", "aaaaaa-bbbb-0000-0000-aaaaaaaaa")
	if err != nil {
		return err
	}

	log.Printf("temlate version: %#v", r)

	return nil
}
