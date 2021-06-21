// Copyright ©2021 The go-pdf Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fpdf

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func BenchmarkParsePNG_rgb(b *testing.B) {
	raw, err := ioutil.ReadFile("image/golang-gopher.png")
	if err != nil {
		b.Fatal(err)
	}

	pdf := New("P", "mm", "A4", "")
	pdf.AddPage()

	const readDPI = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pdf.parsepng(bytes.NewReader(raw), readDPI)
	}
}

func BenchmarkParsePNG_gray(b *testing.B) {
	raw, err := ioutil.ReadFile("image/logo-gray.png")
	if err != nil {
		b.Fatal(err)
	}

	pdf := New("P", "mm", "A4", "")
	pdf.AddPage()

	const readDPI = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pdf.parsepng(bytes.NewReader(raw), readDPI)
	}
}

func BenchmarkParsePNG_small(b *testing.B) {
	raw, err := ioutil.ReadFile("image/logo.png")
	if err != nil {
		b.Fatal(err)
	}

	pdf := New("P", "mm", "A4", "")
	pdf.AddPage()

	const readDPI = true
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pdf.parsepng(bytes.NewReader(raw), readDPI)
	}
}

func BenchmarkParseJPG(b *testing.B) {
	raw, err := ioutil.ReadFile("image/logo_gofpdf.jpg")
	if err != nil {
		b.Fatal(err)
	}

	pdf := New("P", "mm", "A4", "")
	pdf.AddPage()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pdf.parsejpg(bytes.NewReader(raw))
	}
}

func BenchmarkParseGIF(b *testing.B) {
	raw, err := ioutil.ReadFile("image/logo.gif")
	if err != nil {
		b.Fatal(err)
	}

	pdf := New("P", "mm", "A4", "")
	pdf.AddPage()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pdf.parsegif(bytes.NewReader(raw))
	}
}
