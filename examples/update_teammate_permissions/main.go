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
	u, err := c.UpdateTeammatePermissions(context.TODO(), "kenzo.tanaka", &sendgrid.InputUpdateTeammatePermissions{
		IsAdmin: sendgrid.Bool(false),
		Scopes: []*string{
			sendgrid.String("user.profile.read"),
		},
	})
	if err != nil {
		return err
	}
	log.Printf("user: %#v\n", u)
	return nil
}
