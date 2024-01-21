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
	r, err := c.UpdateTemplate(context.TODO(), "d-abcde12345", &sendgrid.InputUpdateTemplate{
		Name: "dummy",
	})
	if err != nil {
		return err
	}

	log.Printf("template: %#v", r)

	return nil
}
