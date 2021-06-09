package tiff_test

import (
	"github.com/go-pdf/fpdf"
	"github.com/go-pdf/fpdf/contrib/tiff"
	"github.com/go-pdf/fpdf/internal/example"
)

// ExampleRegisterFile demonstrates the loading and display of a TIFF image.
func ExampleRegisterFile() {
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPageFormat("L", fpdf.SizeType{Wd: 200, Ht: 200})
	opt := fpdf.ImageOptions{ImageType: "tiff", ReadDpi: false}
	_ = tiff.RegisterFile(pdf, "sample", opt, "../../image/golang-gopher.tiff")
	pdf.Image("sample", 0, 0, 200, 200, false, "", 0, "")
	fileStr := example.Filename("Fpdf_Contrib_Tiff")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/Fpdf_Contrib_Tiff.pdf
}
