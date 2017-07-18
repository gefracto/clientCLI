package main

import (
	"fmt"

	"github.com/gefracto/clientCLI/src/cli"
)

func main() {
	if res, err := cli.Cli(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
