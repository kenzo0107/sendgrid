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
	err := c.UpdateSubuserStatus(context.TODO(), "dummy", &sendgrid.InputUpdateSubuserStatus{
		Disabled: false,
	})
	if err != nil {
		return err
	}

	if err := c.UpdateSubuserIps(context.TODO(), "dummy", []string{"1.1.1.1"}); err != nil {
		return err
	}

	return nil
}
