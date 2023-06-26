/*
 * Copyright (c) 2013-2014 Kurt Jung (Gmail: kurt.w.jung)
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
package fpdf

var registeredIsRtl func(string) bool = nil

func RegisterIsRtl(method func(string) bool) {
	registeredIsRtl = method
}

// IsRtl checks if the text has rtl direction
func isRTL(text string) bool {
	if registeredIsRtl != nil {
		return registeredIsRtl(text)
	}
	if len(text) == 0 {
		return false
	}
	r := []rune(text)
	//Ranges are taken from : https://stackoverflow.com/questions/12006095/javascript-how-to-check-if-character-is-rtl
	if r[0] >= 0x0591 && 0x07FF >= r[0] {
		return true
	}
	if r[0] >= 0xFB1D && 0xFDFD >= r[0] {
		return true
	}
	if r[0] >= 0xFE70 && 0xFEFC >= r[0] {
		return true
	}
	if r[0] == 0x200F || r[0] == 0x202B || r[0] == 0x202E {
		return true
	}
	return false
}
