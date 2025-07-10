**ebcdic** is a small Go library for encoding/decoding strings into/from EBCDIC format. Written by [ygrebnov](https://github.com/ygrebnov).

---

[![GoDoc](https://pkg.go.dev/badge/github.com/ygrebnov/ebcdic)](https://pkg.go.dev/github.com/ygrebnov/ebcdic)
[![Build Status](https://github.com/ygrebnov/ebcdic/actions/workflows/build.yml/badge.svg)](https://github.com/ygrebnov/ebcdic/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ygrebnov/ebcdic)](https://goreportcard.com/report/github.com/ygrebnov/ebcdic)

## Character Sets

The library currently supports only the [EBCDIC invariant character set](https://www.ibm.com/docs/en/i/7.6.0?topic=sets-invariant-character-set-its-exceptions), augmented by the [C0 control characters](https://www.unicode.org/charts/nameslist/n_0000.html), which are mapped to the corresponding Unicode characters. Other character sets can be added in the future if needed.

## Usage

The usage is straightforward. `ebcdic.Encode` encodes a string into EBCDIC, and `ebcdic.Decode` decodes an EBCDIC-encoded string back to text. Both functions take an input string to encode or decode and an optional `ebcdic.CodePage` argument to specify the character set to use. If not provided, the default is `ebcdic.CodePageInvariant`.

## Installation

Compatible with Go 1.22 or later:

```shell
go get github.com/ygrebnov/ebcdic
```

## Contributing

Contributions are welcome!  
Please open an [issue](https://github.com/ygrebnov/ebcdic/issues) or submit a [pull request](https://github.com/ygrebnov/ebcdic/pulls).

## License

Distributed under the BSD-style license. See the [LICENSE](LICENSE) file for details.