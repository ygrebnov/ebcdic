package ebcdic

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"empty string", "", "", false},
		{"single char", "A", "C1", false},
		{"all lowercase", "abc", "818283", false},
		{"all uppercase", "ABC", "C1C2C3", false},
		{"digits", "012", "F0F1F2", false},
		{"symbols", "!@#", "5A7C7B", false},
		{"mixed", "aZ9", "81E9F9", false},
		{"unknown char", "a€", "", true},
		{"multi unknown", "€€", "", true},
		{"partial unknown", "a€b", "", true},
		{"space", " ", "40", false},
		{"long string", "Hello, World!", "C8859393966B40E6969993845A", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.wantErr {
				t.Errorf("New() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"empty string", "", "", false},
		{"single char", "C1", "A", false},
		{"all lowercase", "818283", "abc", false},
		{"all uppercase", "C1C2C3", "ABC", false},
		{"digits", "F0F1F2", "012", false},
		{"symbols", "5A7C7B", "!@#", false},
		{"mixed", "81E9F9", "aZ9", false},
		{"space", "40", " ", false},
		{"long string", "C8859393966B40E6969993845A", "Hello, World!", false},
		{"odd length", "C1C", "", true},
		{"unknown code", "C1ZZ", "", true},
		{"partial unknown", "C1ZZC2", "", true},
		{"multi unknown", "ZZZZ", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.wantErr {
				t.Errorf("Parse() = %q, want %q", got, tt.want)
			}
		})
	}
}
