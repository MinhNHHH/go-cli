package fetch

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type AsciiArt struct {
	lines    []string
	maxWidth int
}

func NewAsciiArt(filePath string) (*AsciiArt, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	width := math.MinInt
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		asc := scanner.Text()
		lines = append(lines, asc)
		placeholders := []string{"${c1}", "${c2}", "${c3}"}
		for _, placeholder := range placeholders {
			asc = strings.ReplaceAll(asc, placeholder, "")
		}
		width = Max(len(asc), width)
	}
	return &AsciiArt{
		lines:    lines,
		maxWidth: width + 5,
	}, nil
}
