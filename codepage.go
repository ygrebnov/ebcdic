package ebcdic

import (
	"strconv"
	"sync"

	"github.com/ygrebnov/errorc"
)

type CodePage = uint16

const (
	CodePageInvariant CodePage = iota
)

var ErrUnknownCodePage = errorc.New("unknown code page")

var (
	to0                map[rune]string
	from0              map[string]byte
	onceTo0, onceFrom0 sync.Once
)

// initTo0 returns the mapping for the invariant code page.
func initTo0() {
	to0 = getTo0()
}

// initFrom0 returns the mapping for the invariant code page.
func initFrom0() {
	from0 = getFrom0()
}

func getTo(c CodePage) (map[rune]string, error) {
	switch c {
	case CodePageInvariant:
		onceTo0.Do(initTo0)
		return to0, nil
	default:
		return nil, errorc.With(ErrUnknownCodePage, errorc.Field("code_page", strconv.Itoa(int(c))))
	}
}

func getFrom(c CodePage) (map[string]byte, error) {
	switch c {
	case CodePageInvariant:
		onceFrom0.Do(initFrom0)
		return from0, nil
	default:
		return nil, errorc.With(ErrUnknownCodePage, errorc.Field("code_page", strconv.Itoa(int(c))))
	}
}
