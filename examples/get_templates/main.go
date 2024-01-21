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
	r, err := c.GetTemplates(context.TODO(), &sendgrid.InputGetTemplates{
		Generations: "dynamic",
		PageSize:    10,
	})
	if err != nil {
		return err
	}

	for _, u := range r.Templates {
		log.Printf("template: %#v", u)
	}

	return nil
}
