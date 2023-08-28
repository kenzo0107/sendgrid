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
	r, err := c.AuthenticateDomain(context.TODO(), &sendgrid.InputAuthenticateDomain{
		Domain: "example.com",
	})
	if err != nil {
		return err
	}
	log.Printf("domain: %#v\n", r)

	return nil
}
