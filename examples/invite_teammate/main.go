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
	u, err := c.InviteTeammate(context.TODO(), &sendgrid.InputInviteTeammate{
		Email:   sendgrid.String("kenzo.tanaka@example.com"),
		IsAdmin: sendgrid.Bool(false),
		Scopes: []*string{
			sendgrid.String("user.profile.read"),
			sendgrid.String("user.profile.update"),
		},
	})
	if err != nil {
		return err
	}
	log.Printf("invite user: %#v\n", u)
	return nil
}
