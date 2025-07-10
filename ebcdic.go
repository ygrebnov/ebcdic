// Copyright 2025 Yaroslav Grebnov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ebcdic provides functions to encode and decode text to and from EBCDIC format.
//
// EBCDIC format is defined by code pages, which are used to map characters to their EBCDIC representations.
// By default, the invariant code page is used.
package ebcdic

import (
	"strings"
	"unsafe"

	"github.com/ygrebnov/errorc"
)

var (
	ErrUnknownCharacter   = errorc.New("unknown character")
	ErrUnknownEBCDICCode  = errorc.New("unknown EBCDIC code")
	ErrInvalidEBCDICInput = errorc.New("invalid EBCDIC input")
)

// Encode encodes the given text to EBCDIC format using the character set corresponding to the specified code page.
func Encode(s string, cp ...CodePage) (string, error) {
	if s == "" {
		return "", nil
	}

	c := CodePageInvariant // Invariant character set is used by default.
	if len(cp) > 0 {
		c = cp[0]
	}

	to, err := getTo(c)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.Grow(len(s) * 2) // pre-allocate memory

	for _, r := range s {
		e, ok := to[r]
		if !ok {
			return "", errorc.With(
				ErrUnknownCharacter,
				errorc.Field("character", string(r)),
			)
		}
		b.WriteString(e)
	}

	return b.String(), nil
}

// Decode decodes an EBCDIC-encoded text to a string using the character set corresponding to the specified code page.
func Decode(s string, cp ...CodePage) (string, error) {
	n := len(s)

	if n == 0 {
		return "", nil
	}

	if n&1 != 0 {
		return "", errorc.With(
			ErrInvalidEBCDICInput,
			errorc.Field("", "input length must be even"),
			errorc.Field("input", s),
		)
	}

	c := CodePageInvariant // Invariant character set is used by default.
	if len(cp) > 0 {
		c = cp[0]
	}

	from, err := getFrom(c)
	if err != nil {
		return "", err
	}

	b := make([]byte, 0, n/2)

	for i := 0; i < n; i += 2 {
		a, ok := from[s[i:i+2]]
		if !ok {
			return "", errorc.With(
				ErrUnknownEBCDICCode,
				errorc.Field("code", s[i:i+2]),
			)
		}

		b = append(b, a)
	}

	// s has even length and at least 2 characters.
	return unsafe.String(&b[0], len(b)), nil
}
