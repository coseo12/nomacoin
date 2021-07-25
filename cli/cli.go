package cli

import (
	"flag"
	"fmt"
	"os"

	explorer "github.com/coseo12/nomacoin/explorer/templates"
	"github.com/coseo12/nomacoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Nomacoin\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000: Set the PORT of the server\n")
	fmt.Printf("-mode=rest:	Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
}
