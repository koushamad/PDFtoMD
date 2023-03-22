package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gen2brain/go-fitz"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run pdf2md.go <path_to_pdf_file>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	mdPath := strings.TrimSuffix(pdfPath, filepath.Ext(pdfPath)) + ".md"

	err := convertPDFToMDWithStyle(pdfPath, mdPath)
	if err != nil {
		fmt.Printf("Error converting PDF to Markdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s\n", pdfPath, mdPath)
}

func convertPDFToMD(pdfPath, mdPath string) error {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return err
	}
	defer doc.Close()

	numPages := doc.NumPage()
	var mdContent string

	for i := 0; i < numPages; i++ {
		text, err := doc.Text(i)
		if err != nil {
			return err
		}

		mdContent += text + "\n"
	}

	err = ioutil.WriteFile(mdPath, []byte(mdContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func convertPDFToMDWithStyle(pdfPath, mdPath string) error {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return err
	}
	defer doc.Close()

	numPages := doc.NumPage()
	var mdContent string

	for i := 0; i < numPages; i++ {
		html, err := doc.HTML(i, true)
		if err != nil {
			return err
		}

		converter := md.NewConverter("", true, nil)
		text, err := converter.ConvertString(html)
		if err != nil {
			return err
		}

		mdContent += text + "\n\n"
	}

	err = ioutil.WriteFile(mdPath, []byte(mdContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
