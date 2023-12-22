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
	senders, err := c.GetVerifiedSenders(context.TODO(), &sendgrid.InputGetVerifiedSenders{
		Limit:      1,
		ID:         705935,
		LastSeenID: 1000,
	})
	if err != nil {
		return err
	}
	for _, v := range senders {
		log.Printf("verified sender: %#v\n", v)
	}

	return nil
}
