package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	
	"is_in/is_in"
	"is_in/models"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(`is_in, v%v | (c) 2021 tpkn.me
A CLI tool for checking whether the parameters from first string are present in second string.

Usage:
  is_in [-options] <string1> <string2>

Options:
  -i, --ignorecase    Ignore case (default: false)
  -d, --delimiter     Parameters delimiter (default: ",")
  -p, --prepare       Prepare parameters before comparison (default: false): trim spaces and replace wrapping single/double quotes
  -w, --why           Print comparsison details (default: false)
  --help              Help
  --version           Print current version
`, version)

func main() {
	var cli = models.CLI{}
	flag.StringVar(&cli.Delimiter, "delimiter", ",", "Parameters delimiter")
	flag.StringVar(&cli.Delimiter, "d", ",", "Alias for --delimiter")
	
	flag.BoolVar(&cli.IgnoreCase, "ignorecase", false, "Ignore case")
	flag.BoolVar(&cli.IgnoreCase, "i", false, "Alias for --ignorecase")
	
	flag.BoolVar(&cli.Prepare, "prepare", false, "Prepare parameters before comparison (trim spaces and replace wrapping single/double quotes)")
	flag.BoolVar(&cli.Prepare, "p", false, "Alias for --prepare")
	
	flag.BoolVar(&cli.Verbose, "why", false, "Print comparsison details")
	flag.BoolVar(&cli.Prepare, "w", false, "Alias for --why")
	
	flag.BoolVar(&cli.Help, "help", false, "Help")
	flag.BoolVar(&cli.Version, "version", false, "Print current version")
	flag.Parse()
	
	log.SetFlags(0)
	log.SetPrefix("Error: ")
	
	if cli.Help {
		fmt.Println(help)
		os.Exit(0)
	}
	
	if cli.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	
	var args = flag.Args()
	if len(args) < 2 {
		log.Fatalln("no strings to check")
	}
	
	var string1 = args[0]
	var string2 = args[1]
	
	var result, err = is_in.Run(string1, string2, cli)
	if !cli.Verbose {
		fmt.Print(result)
	} else {
		fmt.Println(err.Listify())
	}
}
