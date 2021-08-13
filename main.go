package main

import (
	"flag"
	"fmt"
	"os"
	
	"is_in/is_in"
	"is_in/models"
)

var env = "development"
var version = "0.0.0"
var help = `is_in | tpkn.me
A CLI tool for checking whether the parameters from first string are present in second string.
Usage: is_in [-options] <string1> <string2>`

func main() {
	var cli = models.CLI{}
	flag.StringVar(&cli.Delimiter, "d", ",", "Parameter delimiter")
	flag.BoolVar(&cli.IgnoreCase, "i", false, "Ignore case")
	flag.BoolVar(&cli.Prepare, "p", false, "Prepare parameters before comparison (trim spaces and replace wrapping single/double quotes)")
	flag.BoolVar(&cli.Verbose, "why", false, "Print comparison details")
	flag.BoolVar(&cli.Help, "help", false, "Help")
	flag.BoolVar(&cli.Version, "v", false, "Print current version number")
	flag.Parse()
	
	if cli.Help {
		fmt.Print(help)
		os.Exit(0)
	}
	
	if cli.Version {
		fmt.Print(version)
		os.Exit(0)
	}
	
	if len(flag.Args()) < 2 {
		fmt.Print("Error: not enough arguments!")
		os.Exit(1)
	}
	
	var string1 = flag.Args()[0]
	var string2 = flag.Args()[1]
	
	var result, err = is_in.Run(string1, string2, cli)
	
	if !cli.Verbose {
		fmt.Print(result)
	} else {
		// Or print the comparison details
		fmt.Println(err.Listify())
	}
}
