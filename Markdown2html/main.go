package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func convertHeader(line string) string {
	headerLevel := strings.Count(line, "#")
	if headerLevel > 0 && headerLevel <= 6 {
		content := strings.TrimSpace(line[headerLevel:])
		return fmt.Sprintf("<h%d>%s</h%d>", headerLevel, content, headerLevel)
	}
	return line
}

func convertBold(line string) string {
	return strings.ReplaceAll(line, "**", "<strong>")
}

func convertItalic(line string) string {
	return strings.ReplaceAll(line, "*", "<em>")
}

func convertLine(line string) string {
	// Process each line for Markdown syntax
	line = convertHeader(line)
	line = convertBold(line)
	line = convertItalic(line)
	return line
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: markdown_to_html <input.md>")
	}

	inputFile := os.Args[1]
	outputFile := strings.TrimSuffix(inputFile, ".md") + ".html"

	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	writer.WriteString("<html><body>\n")

	for scanner.Scan() {
		line := scanner.Text()
		convertedLine := convertLine(line)
		writer.WriteString(convertedLine + "<br/>\n")
	}

	writer.WriteString("</body></html>\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("Conversion successful! HTML file created: %s\n", outputFile)
}
