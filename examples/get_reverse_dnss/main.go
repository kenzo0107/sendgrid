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
	r, err := c.GetReverseDNSs(context.TODO(), &sendgrid.InputGetReverseDNSs{
		Limit:  1,
		Offset: 0,
		IP:     "149.72.168.239",
	})
	if err != nil {
		return err
	}

	for _, d := range r {
		log.Printf("%#v", d)
	}

	return nil
}
