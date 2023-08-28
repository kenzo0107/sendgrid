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
	domains, err := c.GetAuthenticatedDomains(context.TODO(), &sendgrid.InputGetAuthenticatedDomains{
		Limit: 1,
	})
	if err != nil {
		return err
	}
	for _, v := range domains {
		log.Printf("domain authentication: %#v\n", v)
	}

	return nil
}
