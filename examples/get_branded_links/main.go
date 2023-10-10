package main

import (
	"context"
	"fmt"
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
	brandedLinks, err := c.GetBrandedLinks(context.TODO(), &sendgrid.InputGetBrandedLinks{})
	if err != nil {
		return err
	}
	for _, brandedLink := range brandedLinks {
		fmt.Printf("branded link: %#v", brandedLink)
	}

	return nil
}
