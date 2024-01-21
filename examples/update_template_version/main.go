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
	r, err := c.UpdateTemplateVersion(context.TODO(), "d-12345abcde", "aaaaaa-bbbb-0000-0000-aaaaaaaaa", &sendgrid.InputUpdateTemplateVersion{
		Name:   "dummy",
		Active: 0,
	})
	if err != nil {
		return err
	}

	log.Printf("template version: %#v", r)

	return nil
}
