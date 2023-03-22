# PDF to Markdown Converter

PDF to Markdown Converter is a lightweight command-line tool developed in Go that allows users to easily convert PDF files into Markdown format. The tool leverages the go-fitz package for PDF to HTML conversion and the html-to-markdown package for transforming HTML into Markdown. However, due to varying PDF file structures, the conversion process may not be perfect and some styles could be lost during the conversion.

A simple command-line tool for converting PDF files to Markdown format in Go. This tool uses
the [go-fitz](https://github.com/gen2brain/go-fitz) package for converting PDF to HTML and
the [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) package for converting HTML to Markdown.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system
- [MuPDF](https://mupdf.com/docs/manual.html) library installed on your system

## Installation

1. Clone the repository:

```sh
git clone https://github.com/your_username/pdf-to-markdown.git
cd pdf-to-markdown
```

2. Install the required Go packages:

    ```sh
    go get github.com/gen2brain/go-fitz
    go get github.com/JohannesKaufmann/html-to-markdown
   ```
   
## Usage
1. Build the tool:
   ```shell
   go build -o pdf2md pdf2md.go
   ```
   2. Run the tool with the path to the PDF file:
   ```sh
   ./pdf2md <path_to_pdf_file>
   ```
This will generate a Markdown file with the same name as the input PDF file but with a .md extension in the same directory.

## Note
The quality of the text and style extraction depends on the PDF file's structure, and the conversion might not be perfect. Some styles might be lost during the conversion process.