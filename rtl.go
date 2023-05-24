package fpdf

func IsRtl(text string) bool {
	if len(text) == 0 {
		return false
	}
	r := []rune(text)
	if r[0] >= 0x0600 && 0x06FF >= r[0] {
		return true
	}
	return false
}
