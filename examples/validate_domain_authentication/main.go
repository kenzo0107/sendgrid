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
	r, err := c.ValidateDomainAuthentication(context.TODO(), 1234567)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", r)

	return nil
}
