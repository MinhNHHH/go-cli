package main

import (
	"cli-list/pkg/fetch"
	"os"
	"runtime"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// type DataDisPlay [][]string

// func InitDisplayData(asciiArt Parse, sysInform []string) DataDisPlay {
// 	var data [][]string
// 	// Calculate the maximum number of lines between ascii and sysInform
// 	maxLines := max(len(asciiArt.lines), len(sysInform))
// 	// Append ASCII art and sysInform to data
// 	for i := 0; i < maxLines; i++ {
// 		asciiLine, sysInformLine := "", ""

// 		if i < len(asciiArt.lines) {
// 			asciiArt.lines[i] = strings.ReplaceAll(asciiArt.lines[i], `${c1}`, Red)
// 			asciiArt.lines[i] = strings.ReplaceAll(asciiArt.lines[i], `${c2}`, Blue)
// 			asciiLine = asciiArt.lines[i]
// 		}
// 		if i < len(sysInform) {
// 			sysInformLine = sysInform[i]
// 		}

// 		data = append(data, []string{asciiLine, sysInformLine})
// 	}
// 	return data
// }

// func DisplayInfor(asciiArt Parse, sysInfor []string) {
// 	data := InitDisplayData(asciiArt, sysInfor)
// 	// Print the table
// 	for i := 0; i < len(data); i++ {
// 		// Left-align the image column (fixed width)
// 		// image := data[i][0]

// 		// Left-align the text column
// 		text := fmt.Sprintf("%20s", data[i][1])
// 		fmt.Println("||", text)
// 		// Print each row of the table
// 		// fmt.Printf("%s | %s\n", image, text)
// 	}
// }

func defaultArtSys() string {
	switch runtime.GOOS {
	case "linux":
		return "./pkg/fetch/ascii_art/linux.txt"
	case "darwin":
		return "./system/ascii_art/linux.txt"
	case "windows":
		return "./system/ascii_art/linux.txt"
	default:
		return "./system/ascii_art/default.txt"
	}
}

func main() {
	var cmd []string
	if len(os.Args) < 2 {
		cmd = []string{"info"}
	} else {
		cmd = os.Args[1:]
	}
	fetch.HandleClient("/home/minh/Desktop/go-cli/pkg/fetch/ascii_art/default.txt", cmd)
}
