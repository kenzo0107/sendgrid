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
	r, err := c.UpdateVerifiedSender(context.TODO(), 123456789, &sendgrid.InputUpdateVerifiedSender{
		FromName:  "dummy",
		FromEmail: "dummy@example.com",
		ReplyTo:   "dummy@example.com",
		Address:   "dummy",
		City:      "dummy",
		Country:   "dummy",
		Nickname:  "dummy",
	})
	if err != nil {
		return err
	}
	log.Printf("verified sender: %#v\n", r)

	return nil
}
