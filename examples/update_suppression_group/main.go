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
	r, err := c.UpdateSuppressionGroup(context.TODO(), 10000, &sendgrid.InputUpdateSuppressionGroup{
		Name:        "dummy 2",
		Description: "dummy description 2",
		IsDefault:   false,
	})
	if err != nil {
		return err
	}

	log.Printf("suppressionGroup: %#v", r)

	return nil
}
