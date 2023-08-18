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
	r, err := c.CreateSubuser(context.TODO(), &sendgrid.InputCreateSubuser{
		Username: "dummy",
		Email:    "dummy@example.com",
		Password: "dummydummy1!",
		Ips:      []string{"1.1.1.1"},
	})
	if err != nil {
		return err
	}

	log.Printf("subuser: %#v", r)

	return nil
}
