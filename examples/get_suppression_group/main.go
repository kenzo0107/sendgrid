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
	suppressionGroup, err := c.GetSuppressionGroup(context.TODO(), 10000)
	if err != nil {
		return err
	}

	log.Printf("suppressionGroup: %#v", suppressionGroup)

	return nil
}