/*
 * Copyright (c) 2016 Kurt Jung (Gmail: kurt.w.jung)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// Package tiff allows standard (LZW-compressed) TIFF images to be used in
// documents generated with gofpdf.
package tiff

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/go-pdf/fpdf"
	"golang.org/x/image/tiff"
)

// RegisterReader registers a TIFF image, adding it to the PDF file but not
// adding it to the page. imgName specifies the name that will be used in the
// call to Image() that actually places the image in the document. options
// specifies various image properties; in this case, the ImageType property
// should be set to "tiff". The TIFF image is a reader from the reader
// specified by r.
func RegisterReader(pdf *fpdf.Fpdf, imgName string, options fpdf.ImageOptions, r io.Reader) (info *fpdf.ImageInfoType) {
	var err error
	var img image.Image
	var buf bytes.Buffer
	if pdf.Ok() {
		if options.ImageType == "tiff" || options.ImageType == "tif" {
			img, err = tiff.Decode(r)
			if err == nil {
				err = png.Encode(&buf, img)
				if err == nil {
					options.ImageType = "png"
					info = pdf.RegisterImageOptionsReader(imgName, options, &buf)
				}
			}
		} else {
			err = fmt.Errorf("expecting \"tiff\" or \"tif\" as image type, got \"%s\"", options.ImageType)
		}
		if err != nil {
			pdf.SetError(err)
		}
	}
	return
}

// RegisterFile registers a TIFF image, adding it to the PDF file but not
// adding it to the page. imgName specifies the name that will be used in the
// call to Image() that actually places the image in the document. options
// specifies various image properties; in this case, the ImageType property
// should be set to "tiff". The TIFF image is read from the file specified by
// tiffFileStr.
func RegisterFile(pdf *fpdf.Fpdf, imgName string, options fpdf.ImageOptions, tiffFileStr string) (info *fpdf.ImageInfoType) {
	var f *os.File
	var err error

	if pdf.Ok() {
		f, err = os.Open(tiffFileStr)
		if err == nil {
			info = RegisterReader(pdf, imgName, options, f)
			f.Close()
		} else {
			pdf.SetError(err)
		}
	}
	return
}
