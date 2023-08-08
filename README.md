# SendGrid API in Go

[![GoDoc](https://godoc.org/github.com/kenzo0107/sendgrid?status.svg)](https://godoc.org/github.com/kenzo0107/sendgrid) [![test](https://github.com/kenzo0107/sendgrid/actions/workflows/test.yml/badge.svg)](https://github.com/kenzo0107/sendgrid/actions/workflows/test.yml) [![lint](https://github.com/kenzo0107/sendgrid/actions/workflows/lint.yml/badge.svg)](https://github.com/kenzo0107/sendgrid/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/kenzo0107/sendgrid/branch/main/graph/badge.svg)](https://codecov.io/gh/kenzo0107/sendgrid)

This library supports most if not all of the [SendGrid](https://sendgrid.com) REST calls.

## Installing

### _go get_

    $ go get -u github.com/kenzo0107/sendgrid

## Example

### Get Teammate

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/kenzo0107/sendgrid"
)

func main() {
	apiKey := os.Getenv("SENDGRID_API_KEY")

	c := sendgrid.New(apiKey)
	u, err := c.GetTeammate(context.TODO(), "username")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	log.Printf("user: %#v\n", u)
}
```

## License

[MIT License](https://github.com/kenzo0107/sendgrid/blob/main/LICENSE)
