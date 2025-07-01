package ebcdic

type CodePage = uint16

const (
	CodePageInvariant CodePage = iota
)

func getTo(c CodePage) map[rune]string {
	switch c {
	case CodePageInvariant:
		return getTo0()
	default:
		panic("unknown code page")
	}
}

func getFrom(c CodePage) map[string]byte {
	switch c {
	case CodePageInvariant:
		return getFrom0()
	default:
		panic("unknown code page")
	}
}
