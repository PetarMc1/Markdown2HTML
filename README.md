# Markdown2HTML CLI Tool

Markdown2HTML is a command-line tool written in Go that converts Markdown files to HTML. It is designed to be simple and efficient for users looking to convert Markdown documents easily.

## Features

- Converts Markdown to HTML
- Supports headings, lists, paragraphs, and other common Markdown elements
- Wraps output in a complete HTML structure with DOCTYPE, <head>, and <body> tags
- Automatically generates an output file in the same directory as the input file

## Installation

### Prerequisites

Make sure you have [Go](https://golang.org/dl/) installed on your machine. You can verify the installation by running:

```bash
go version
```

### Build the Project

1. Clone the repository:

```bash
git clone https://github.com/PetarMc1/Markdown2HTML.git
```

2. Navigate to the project directory:

```bash
cd Markdown2HTML
```

3. Build the tool:

```bash
go build -o md2html OR go build -o md2html.exe (for Windows)
```

This will generate an executable file called md2html (or md2html.exe on Windows).

### Running the Tool

To convert a Markdown file to HTML, run the following command:

```bash
md2html <input-file>
```

For example:

```bash
md2html README.md
```

This will generate an `output.html` file in the same directory.

## Usage Example

If your Markdown file (`README.md`) contains:

```markdown
# Test Document

This is a test of Markdown to HTML conversion.

- Item 1
- Item 2
- Item 3
```

Running the command:

```bash
md2html sample.md
```

Will generate an HTML file with the following structure:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Markdown Output</title>
</head>
<body>
<h1>Test Document</h1>
<p>This is a test of Markdown to HTML conversion.</p>
<ul>
    <li>Item 1</li>
    <li>Item 2</li>
    <li>Item 3</li>
</ul>
</body>
</html>
```

## Contributing

Feel free to contribute to this project! Fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the [MIT License](/LICENSE).

