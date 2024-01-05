package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/winebarrel/las"
)

var version string

func parseArgs() *las.Options {
	var CLI struct {
		las.Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Options
}

func main() {
	opts := parseArgs()
	c, err := las.NewClient(opts)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = c.ListAddSuppressedDestinations(func(sds []types.SuppressedDestinationSummary) {
		for _, sd := range sds {
			j, _ := json.Marshal(sd)
			fmt.Println(string(j))
		}
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
