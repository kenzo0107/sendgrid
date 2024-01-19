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
	suppressionGroups, err := c.GetSuppressionGroups(context.TODO())
	if err != nil {
		return err
	}

	for _, suppressionGroup := range suppressionGroups {
		log.Printf("suppressionGroup: %#v", suppressionGroup)
	}

	return nil
}
