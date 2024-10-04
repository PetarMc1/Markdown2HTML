package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

// Function to read the Markdown file
func readMarkdown(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Function to convert Markdown to HTML using blackfriday with full extensions
func convertMarkdownToHTML(content []byte) []byte {
	// Use all common Markdown extensions
	extensions := blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs | blackfriday.NoEmptyLineBeforeBlock
	htmlRenderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})

	return blackfriday.Run(content, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(htmlRenderer))
}

// Function to wrap HTML content in proper HTML structure (head, body, etc.)
func wrapInHTML(content string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Markdown Output</title>
</head>
<body>
%s
</body>
</html>`, content)
}

// Function to write HTML to a file
func writeHTML(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// Function to generate the output file path
func generateOutputPath(input string) string {
	return filepath.Join(filepath.Dir(input), "output.html")
}

func main() {
	// Check for command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: md2html <filename>")
		return
	}

	input := os.Args[1]

	// Read the Markdown file
	markdownContent, err := readMarkdown(input)
	if err != nil {
		log.Fatalf("Error reading Markdown file: %v", err)
	}

	// Convert Markdown to HTML with full extensions
	htmlContent := convertMarkdownToHTML(markdownContent)

	// Wrap the converted content in an HTML structure
	fullHTML := wrapInHTML(string(htmlContent))

	// Generate output file path
	output := generateOutputPath(input)

	// Write the wrapped HTML to the output file
	err = writeHTML(output, []byte(fullHTML))
	if err != nil {
		log.Fatalf("Error writing HTML file: %v", err)
	}
	fmt.Printf("HTML written to %s\n", output)
}
