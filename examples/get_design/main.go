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
	r, err := c.GetDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901")
	if err != nil {
		return err
	}
	log.Printf("%#v", r)

	return nil
}
