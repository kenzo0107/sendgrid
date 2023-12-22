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
	r, err := c.UpdateVerifiedSender(context.TODO(), 12345678, &sendgrid.InputUpdateVerifiedSender{
		Nickname:  "dummy",
		Country:   "dummy",
		ReplyTo:   "dummy@example.com",
		FromName:  "dummy",
		FromEmail: "dummy@example.com",
		Address:   "dummy",
		Address2:  "dummy2",
		City:      "dummy",
	})
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", r)

	return nil
}
