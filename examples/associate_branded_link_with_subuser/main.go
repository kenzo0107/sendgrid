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
	r, err := c.AssociateBrandedLinkWithSubuser(context.TODO(), 3401510, &sendgrid.InputAssociateBrandedLinkWithSubuser{
		Username: "subuser_name",
	})
	if err != nil {
		return err
	}
	log.Printf("%#v\n", r)

	return nil
}
