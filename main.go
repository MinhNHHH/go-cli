package main

import (
	"bufio"
	"cli-list/system"
	"flag"
	"fmt"
	"os"
	"runtime"
)

type Parse struct {
	lines []string
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewParse(filePath string) (*Parse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	red := "\033[91m"
	reset := "\033[0m"

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, fmt.Sprintf("%s%s%s", red, scanner.Text(), reset))
	}

	return &Parse{
		lines: lines,
	}, nil
}

type DataDisPlay [][]string

func InitDisplayData(asciiArt, sysInform []string) DataDisPlay {
	var data [][]string
	// Calculate the maximum number of lines between ascii and sysInform
	maxLines := max(len(asciiArt), len(sysInform))

	// Append ASCII art and sysInform to data
	for i := 0; i < maxLines; i++ {
		asciiLine, sysInformLine := "", ""

		if i < len(asciiArt) {
			asciiLine = asciiArt[i]
		}

		if i < len(sysInform) {
			sysInformLine = sysInform[i]
		}
		data = append(data, []string{asciiLine, sysInformLine})
	}
	return data
}

func PrintInfo(asciiArt, sysInfor []string) {
	data := InitDisplayData(asciiArt, sysInfor)
	for _, row := range data {
		fmt.Println(row[0], "\t", row[1])
	}
}

func main() {
	// Define flags

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	disableCmd := flag.NewFlagSet("disable", flag.ExitOnError)
	asciiCmd := flag.NewFlagSet("ascii", flag.ExitOnError)
	var asciiArtPath string
	var disableInfo []string

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
	case "disable":
		disableCmd.Parse(os.Args[2:])
		disableInfo = disableCmd.Args()
	case "ascii":
		asciiCmd.Parse(os.Args[2:])
		if len(asciiCmd.Args()) == 0 {
			asciiArtPath = "./system/ascii_art/linux.txt"
		} else {
			asciiArtPath = asciiCmd.Args()[0]
		}
	default:
		switch runtime.GOOS {
		case "linux":
			asciiArtPath = "./system/ascii_art/linux.txt"
		case "macos":
			asciiArtPath = "./system/ascii_art/linux.txt"
		case "windows":
			asciiArtPath = "./system/ascii_art/linux.txt"
		default:
			asciiArtPath = ""
		}
	}
	sysInfor := system.System()
	asciiArt, err := NewParse(asciiArtPath)
	if err != nil {
		fmt.Printf("Error open file: %s\n", err.Error())
		return
	}
	PrintInfo(asciiArt.lines, sysInfor.GenInfoSys(disableInfo))
}
