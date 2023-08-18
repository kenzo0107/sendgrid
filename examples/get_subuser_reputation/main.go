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
	reputations, err := c.GetSubuserReputations(context.TODO(), "dummy")
	if err != nil {
		return err
	}

	for _, reputation := range reputations {
		log.Printf("reputation: %#v", reputation)
	}

	return nil
}
