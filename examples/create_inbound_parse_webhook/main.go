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
	r, err := c.CreateInboundParsetWebhook(context.TODO(), &sendgrid.InputCreateInboundParsetWebhook{
		URL:       "https://example.com/sendgrid/inbound",
		Hostname:  "bar.foo",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err != nil {
		return err
	}
	log.Printf("%#v\n", r)

	return nil
}
