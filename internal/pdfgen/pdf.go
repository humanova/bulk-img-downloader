// (c) 2020 Emir Erbasan (humanova)
// MIT License

package pdfgen

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(filename string, imagePaths []string, imageExtension string) error {

	fmt.Printf("[%d images] creating a pdf document...", len(imagePaths))
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, imagePath := range imagePaths {
		pdf.AddPage()

		pdf.ImageOptions(
			imagePath,
			0, 0,
			210, 297, // fill the page (for A4)
			false,
			gofpdf.ImageOptions{ImageType: imageExtension, ReadDpi: true},
			0,
			"",
		)
	}
	return pdf.OutputFileAndClose(filename)
}
