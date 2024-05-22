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
	r, err := c.UpdateDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901", &sendgrid.InputUpdateDesign{
		Name:        "example_updated",
		HTMLContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err != nil {
		return err
	}
	log.Printf("%#v\n", r)

	return nil
}
