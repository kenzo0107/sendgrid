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
	users, err := c.GetPendingTeammates(context.TODO())
	if err != nil {
		return err
	}

	for _, u := range users {
		log.Printf("user: %#v", u)
	}

	return nil
}
