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
	r, err := c.CreateSuppressionGroup(context.TODO(), &sendgrid.InputCreateSuppressionGroup{
		Name:        "dummy",
		Description: "dummy description",
		IsDefault:   false,
	})
	if err != nil {
		return err
	}

	log.Printf("suppressionGroup: %#v", r)

	return nil
}
