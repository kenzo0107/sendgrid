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
	r, err := c.CreateSSOTeammate(context.TODO(), &sendgrid.InputCreateSSOTeammate{
		Email:                      "dummy@example.com",
		FirstName:                  "hoge",
		LastName:                   "moge",
		IsSSO:                      true,
		IsAdmin:                    false,
		HasRestrictedSubuserAccess: true,
		SubuserAccess: []sendgrid.InputSubuserAccess{
			{
				ID:             12345678,
				PermissionType: "restricted",
				Scopes:         []string{"mail.send", "alerts.create"},
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("temlate: %#v", r)

	return nil
}
