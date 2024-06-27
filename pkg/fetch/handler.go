package fetch

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

var CodeColor = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[1;31m",
	"green":  "\033[1;32m",
	"cyan":   "\033[1;33m",
	"yellow": "\033[1;34m",
	"purple": "\033[1;35m",
	"blue":   "\033[1;36m",
	"white":  "\033[1;37m",
}

var placeHolder = map[string]string{
	"${c0}": CodeColor["reset"],
	"${c1}": CodeColor["red"],
	"${c2}": CodeColor["green"],
	"${c3}": CodeColor["yellow"],
	"${c4}": CodeColor["blue"],
	"${c5}": CodeColor["purple"],
	"${c6}": CodeColor["cyan"],
	"${c7}": CodeColor["white"],
}

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

func DefaultArtSys() string {
	return runtime.GOOS
}

func HandleClient(cmd []string) {
	sysInfor := NewSysInfor()
	ascii := DefaultArt(DefaultArtSys())
	client := &ClientDetail{
		AsciiArt: ascii,
		SysInfor: sysInfor,
	}
	client.handleCommand(cmd)
}

func (c *ClientDetail) handleCommand(command []string) {
	disable := []string{}
	seemore := []string{}
	switch command[0] {
	case "list":
	case "source":
		if len(command) >= 2 {
			ascii, err := NewAsciiArt(command[1])
			if err != nil {
				log.Fatalf(err.Error())
			}
			c.AsciiArt = ascii
		}
	case "disable":
		disable = command[1:]
	case "ascii_distro":
		if len(command) >= 2 {
			ascii := DefaultArt(command[1])
			c.AsciiArt = ascii
		}
	case "ascii_color":
		if len(command) >= 2 {
			for i, color := range command[1:] {
				codeColor := fmt.Sprintf("${c%d}", i+1)
				placeHolder[codeColor] = CodeColor[color]
			}
		}
	default:
	}
	c.printInfor(disable, seemore)
}

func (c *ClientDetail) CountPattern(input string) int {
	count := 0
	for label := range placeHolder {
		if label == "${c0}" {
			count += strings.Count(input, label) * 4
		} else {
			count += strings.Count(input, label) * 7
		}
	}
	return count
}

func (c *ClientDetail) replacePlaceHolder(input string) string {
	for label, color := range placeHolder {
		input = strings.ReplaceAll(input, label, color)
	}
	return input
}

func (c *ClientDetail) printInfor(disable, seemore []string) {
	listInfor := c.SysInfor.ListSysInfor(disable, seemore)
	maxLines := Max(len(c.AsciiArt.Lines), len(listInfor))
	asciiLine, sysInformLine := "", ""
	for i := 0; i < maxLines; i++ {
		pattern := c.CountPattern(c.AsciiArt.Lines[i])
		if i < len(c.AsciiArt.Lines) {
			c.AsciiArt.Lines[i] = c.replacePlaceHolder(c.AsciiArt.Lines[i])
			asciiLine = c.AsciiArt.Lines[i]
		}

		if i < len(listInfor) {
			sysInformLine = listInfor[i]
		} else {
			sysInformLine = ""
		}
		originalDistance := c.AsciiArt.MaxCleanLen + pattern
		padding := 5
		fmt.Printf("%-*s %s\n", originalDistance+padding, asciiLine, sysInformLine)
	}
}
