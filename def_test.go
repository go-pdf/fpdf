// Copyright ©2023 The go-pdf Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fpdf

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestPDFVersionOrder(t *testing.T) {
	want := []pdfVersion{
		0,
		pdfVersionFrom(0, 1),
		pdfVersionFrom(0, 10),
		pdfVersionFrom(1, 0),
		pdfVersionFrom(1, 1),
		pdfVersionFrom(1, 2),
		pdfVers1_3,
		pdfVers1_4,
		pdfVers1_5,
		pdfVersionFrom(1, 10),
		pdfVersionFrom(2, 0),
		pdfVersionFrom(2, 1),
		pdfVersionFrom(2, 10),
	}
	sorted := sort.SliceIsSorted(want, func(i, j int) bool {
		return want[i] < want[j]
	})
	if !sorted {
		t.Fatalf("PDF-version ordering is flawed")
	}

	got := make([]pdfVersion, len(want))
	copy(got, want)

	rnd := rand.New(rand.NewSource(1234))
	rnd.Shuffle(len(got), func(i, j int) {
		got[i], got[j] = got[j], got[i]
	})

	if reflect.DeepEqual(got, want) {
		t.Fatalf("shuffling failed")
	}

	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PDF-version ordering is wrong:\ngot= %q\nwant=%q", got, want)
	}
}
