package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func extractHeadings(mdContent string) []string {
	// Regular expression to match H1, H2, and H3 headings
	re := regexp.MustCompile(`(?m)^(#{1,3})\s+(.*)$`)

	matches := re.FindAllStringSubmatch(mdContent, -1)

	var headings []string
	for _, match := range matches {
		headings = append(headings, match[0]) // Entire matched heading
	}

	return headings
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

type SubSubHeader struct {
	H3 string `json:"h3"`
}

type SubHeader struct {
	H2       string         `json:"h2"`
	Children []SubSubHeader `json:"children"`
}

type Header struct {
	H1       string      `json:"h1"`
	Children []SubHeader `json:"children"`
}

func parseHeaders(headers []string) []Header {
	var result []Header
	var currentH1 *Header
	var currentH2 *SubHeader

	for _, header := range headers {
		switch {
		case len(header) > 3 && header[:3] == "###":
			if currentH2 != nil {
				currentH2.Children = append(currentH2.Children, SubSubHeader{H3: header[4:]})
			}
		case len(header) > 2 && header[:2] == "##":
			newH2 := SubHeader{H2: header[3:], Children: []SubSubHeader{}}
			if currentH1 != nil {
				currentH1.Children = append(currentH1.Children, newH2)
			}
			currentH2 = &currentH1.Children[len(currentH1.Children)-1]
		case len(header) > 1 && header[:1] == "#":
			newH1 := Header{H1: header[2:], Children: []SubHeader{}}
			result = append(result, newH1)
			currentH1 = &result[len(result)-1]
			currentH2 = nil
		}
	}
	return result
}

func writeJSONFile(headers []Header, filename string) {
	data := map[string]interface{}{"headers": headers}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func main() {
	ext := "../src/pages/ducks/"
	cmd := exec.Command("ls", "../src/pages/ducks")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		os.Exit(1)
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")

	file, _ := os.Open(ext + files[0])
	data, _ := io.ReadAll(file)
	headings := extractHeadings(string(data))
	headerData := parseHeaders(headings)

	newFile := "navbarSchema.json"
	if !fileExists(ext + newFile) {
		os.Create(ext + newFile)
	}
	file, err = os.OpenFile(ext+newFile, os.O_RDWR|os.O_CREATE, 0644)
	writeJSONFile(headerData, ext+newFile)
}
