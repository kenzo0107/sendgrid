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
	subusers, err := c.GetSubusers(context.TODO(), &sendgrid.InputGetSubusers{
		Username:      "dummy.dummy",
		Limit:         1,
		Offset:        0,
		IncludeRegion: true,
	})
	if err != nil {
		return err
	}

	for _, subuser := range subusers {
		log.Printf("subuser: %#v", subuser)
	}

	return nil
}
