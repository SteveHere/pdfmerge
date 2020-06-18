package main

import (
	"fmt"
	"os"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Requires at least 3 arguments: output_path and 2 input paths\n")
		fmt.Printf("Usage: pdf_merge.exe output.pdf input1.pdf input2.pdf input3.pdf ...\n")
		os.Exit(0)
	}

	outputPath := os.Args[1]
	if err := pdf.MergeCreateFile(os.Args[2:], outputPath, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Complete, see output file: %s\n", outputPath)
	}
}
