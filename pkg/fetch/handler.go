package fetch

import (
	"fmt"
	"log"
	"strings"
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

type ClientDetail struct {
	SysInfor SystemInfor
	AsciiArt *AsciiArt
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func HandleClient(filePath string, cmd []string) {
	ascii, err := NewAsciiArt(filePath)
	sysInfor := NewSysInfor()
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := &ClientDetail{
		AsciiArt: ascii,
		SysInfor: sysInfor,
	}
	client.handleCommand(cmd)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (c *ClientDetail) handleCommand(command []string) {
	switch command[0] {
	case "list":
	case "source":
	case "info":
		c.printInfor([]string{})
	default:
		fmt.Println("áđâs")
	}
}

func (c *ClientDetail) CountPattern(input string) int {
	count := 0
	placeholders := []string{"${c1}", "${c2}", "${c3}"}
	for _, placeholder := range placeholders {
		count += strings.Count(input, placeholder)
	}
	return count
}

func (c *ClientDetail) printInfor(disable []string) {
	listInfor := c.SysInfor.ListSysInfor(disable)
	maxLines := Max(len(c.AsciiArt.lines), len(listInfor))
	asciiLine, sysInformLine := "", ""
	for i := 0; i < maxLines; i++ {
		repeat := Abs(c.AsciiArt.maxWidth - c.CountPattern(c.AsciiArt.lines[i])*5)
		// fmt.Println(len(c.AsciiArt.lines[i]), c.CountPattern(c.AsciiArt.lines[i])*5, c.AsciiArt.maxWidth, repeat)
		if i < len(c.AsciiArt.lines) {
			c.AsciiArt.lines[i] = strings.ReplaceAll(c.AsciiArt.lines[i], `${c1}`, White)
			c.AsciiArt.lines[i] = strings.ReplaceAll(c.AsciiArt.lines[i], `${c2}`, Blue)
			c.AsciiArt.lines[i] = strings.ReplaceAll(c.AsciiArt.lines[i], `${c3}`, Red)
			asciiLine = c.AsciiArt.lines[i] + strings.Repeat(" ", repeat)
		}

		if i < len(listInfor) {
			sysInformLine = listInfor[i]
		} else {
			sysInformLine = ""
		}
		fmt.Printf("%s %5s\n", asciiLine, sysInformLine)
		// data = append(data, []string{asciiLine, sysInformLine})
	}
}
