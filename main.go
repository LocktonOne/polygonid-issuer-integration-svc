package main

import (
	"os"

  "gitlab.com/tokene/polygonid-issuer-integration/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
