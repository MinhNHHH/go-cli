// package main

// import (
// 	"cli-list/system"
// 	"flag"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Define flags
// 	helpCmd := flag.NewFlagSet("list", flag.ExitOnError)

// 	// Define subcommands
// 	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

// 	// Parse command-line arguments
// 	flag.Parse()

// 	// Check the subcommand and execute the corresponding function
// 	switch flag.Arg(0) {
// 	case "list":
// 		listCmd.Parse(os.Args[2:]) // Parse subcommand arguments
// 		listItems()
// 	case "help":
// 		helpCmd.Parse(os.Args[2:])
// 		showHelp()
// 	default:
// 		system.System()

// 	}
// }

// func showHelp() {
// 	fmt.Println("Usage: myapp [command]")
// 	fmt.Println("Commands:")
// 	fmt.Println("  list   List items")
// 	fmt.Println("Options:")
// 	flag.PrintDefaults()
// }

// func listItems() {
// 	// Your list command logic goes here
// 	fmt.Println("Listing items...")
// }

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open the ASCII art file
	file, err := os.Open("ascii_art.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	red := "\033[91m"
	reset := "\033[0m"
	// Read the content of the file
	asciiArtBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert bytes to string and print ASCII art
	asciiArt := string(asciiArtBytes)

	// Resize the ASCII art

	fmt.Printf("%s%s%s\n", red, asciiArt, reset)
}
