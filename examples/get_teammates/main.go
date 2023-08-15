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
	r, err := c.GetTeammates(context.TODO())
	if err != nil {
		return err
	}

	for _, u := range r.Teammates {
		log.Printf("user: %#v", u)
	}

	return nil
}
