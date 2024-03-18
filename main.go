package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	helpCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Define subcommands
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Parse command-line arguments
	flag.Parse()

	// Check the subcommand and execute the corresponding function
	switch flag.Arg(0) {
	case "list":
		listCmd.Parse(os.Args[2:]) // Parse subcommand arguments
		listItems()
	case "help":
		helpCmd.Parse(os.Args[2:])
		showHelp()
	default:
		fmt.Println("asdasdasd")
	}
}

func showHelp() {
	fmt.Println("Usage: myapp [command]")
	fmt.Println("Commands:")
	fmt.Println("  list   List items")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func listItems() {
	// Your list command logic goes here
	fmt.Println("Listing items...")
}
