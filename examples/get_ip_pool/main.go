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
	r, err := c.GetIPPool(context.TODO(), "marketing")
	if err != nil {
		return err
	}
	log.Printf("pool name: %#v\n", r.PoolName)

	for _, ip := range r.IPs {
		log.Printf("%#v\n", ip)
	}
	return nil
}
