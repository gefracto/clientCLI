package main

import (
	"fmt"

	"github.com/gefracto/clientCLI/src/cli"
)

func main() {
	res, err := cli.Cli()
	fmt.Println("Reason: " + fmt.Sprint(err))
	fmt.Println("Result:\n" + res)

}
