package main

import (
	"fmt"

	"gitlab.com/EndowTheGreat/spotman/cmd/cli"
)

func main() {
	if err := cli.Setup(); err != nil {
		fmt.Println(err)
	}
}
