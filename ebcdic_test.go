package ebcdic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	codePageInvariant = CodePageInvariant
	codePageInvalid   = CodePage(10) // An example of an invalid code page for testing purposes.
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		codePage      *CodePage
		expected      string
		expectedError error
	}{
		{
			"nominal default codepage",
			"Hello, World!",
			nil,
			"C8859393966B40E6969993845A",
			nil,
		},
		{
			"nominal invariant",
			"Hello, World!",
			&codePageInvariant,
			"C8859393966B40E6969993845A",
			nil,
		},
		{"empty string", "", nil, "", nil},
		{"single char", "A", nil, "C1", nil},
		{"all lowercase", "abc", nil, "818283", nil},
		{"all uppercase", "ABC", nil, "C1C2C3", nil},
		{"digits", "012", nil, "F0F1F2", nil},
		{"symbols", "!@#", nil, "5A7C7B", nil},
		{"mixed", "aZ9", nil, "81E9F9", nil},
		{"space", " ", nil, "40", nil},
		{"unknown char", "a€", nil, "", ErrUnknownCharacter},
		{"multi unknown", "€€", nil, "", ErrUnknownCharacter},
		{"partial unknown", "a€b", nil, "", ErrUnknownCharacter},
		{"unknown codepage", "a€b", &codePageInvalid, "", ErrUnknownCodePage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				encoded string
				err     error
			)

			if tt.codePage == nil {
				encoded, err = Encode(tt.input)
			} else {
				encoded, err = Encode(tt.input, *tt.codePage)
			}

			require.ErrorIs(t, err, tt.expectedError, "Encode() error mismatch")

			if err == nil {
				require.Equal(t, tt.expected, encoded, "Encode() result mismatch")
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		codePage      *CodePage
		expected      string
		expectedError error
	}{
		{
			"nominal default codepage",
			"C8859393966B40E6969993845A",
			nil,
			"Hello, World!",
			nil,
		},
		{
			"nominal invariant",
			"C8859393966B40E6969993845A",
			&codePageInvariant,
			"Hello, World!",
			nil,
		},
		{"empty string", "", nil, "", nil},
		{"single char", "C1", nil, "A", nil},
		{"all lowercase", "818283", nil, "abc", nil},
		{"all uppercase", "C1C2C3", nil, "ABC", nil},
		{"digits", "F0F1F2", nil, "012", nil},
		{"symbols", "5A7C7B", nil, "!@#", nil},
		{"mixed", "81E9F9", nil, "aZ9", nil},
		{"space", "40", nil, " ", nil},
		{"odd length", "C1C", nil, "", ErrInvalidEBCDICInput},
		{"unknown code", "C1ZZ", nil, "", ErrUnknownEBCDICCode},
		{"partial unknown", "C1ZZC2", nil, "", ErrUnknownEBCDICCode},
		{"multi unknown", "ZZZZ", nil, "", ErrUnknownEBCDICCode},
		{"unknown codepage", "ZZZZ", &codePageInvalid, "", ErrUnknownCodePage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				decoded string
				err     error
			)

			if tt.codePage == nil {
				decoded, err = Decode(tt.input)
			} else {
				decoded, err = Decode(tt.input, *tt.codePage)
			}

			require.ErrorIs(t, err, tt.expectedError, "Decode() error mismatch")

			if err == nil {
				require.Equal(t, tt.expected, decoded, "Decode() result mismatch")
			}
		})
	}
}
