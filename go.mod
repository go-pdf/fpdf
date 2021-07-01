module github.com/go-pdf/fpdf

go 1.16

retract (
	// v1.4.3 was tagged to retract v1.4.2
	v1.4.3
	// v1.4.2 was imported from phpdave11/gofpdf by mistake.
	v1.4.2
)

require (
	github.com/boombuler/barcode v1.0.1
	github.com/phpdave11/gofpdi v1.0.13
	github.com/ruudk/golang-pdf417 v0.0.0-20201230142125-a7e3863a1245
	golang.org/x/image v0.0.0-20210607152325-775e3b0c77b9
)
