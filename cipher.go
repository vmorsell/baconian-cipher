package main

import "strings"

// Code length
const codeLength = 5

// Mapping from char to codes
var codeMapping = map[rune]string{
	'A': "aaaaa",
	'B': "aaaab",
	'C': "aaaba",
	'D': "aaabb",
	'E': "aabaa",
	'F': "aabab",
	'G': "aabba",
	'H': "aabbb",
	'I': "abaaa",
	'J': "abaab",
	'K': "ababa",
	'L': "ababb",
	'M': "abbaa",
	'N': "abbab",
	'O': "abbba",
	'P': "abbbb",
	'Q': "baaaa",
	'R': "baaab",
	'S': "baaba",
	'T': "baabb",
	'U': "babaa",
	'V': "babab",
	'W': "babba",
	'X': "babbb",
	'Y': "bbaaa",
	'Z': "bbaab",
}

// reverseMapping creates a reverse lookup table for decoding
func reverseMapping(mapping map[rune]string) map[string]rune {
	out := make(map[string]rune, len(mapping))

	for k, v := range mapping {
		out[v] = k
	}

	return out
}

// Encode a messasge
func Encode(msg string) string {
	var out strings.Builder

	for _, ch := range strings.ToUpper(msg) {
		if v, ok := codeMapping[ch]; ok {
			out.WriteString(v)
		}
	}

	return out.String()
}

// Decode a message
func Decode(msg string) string {
	var out strings.Builder

	m := reverseMapping(codeMapping)

	for i := 0; i < len(msg); i += codeLength {
		if i+codeLength > len(msg) {
			break
		}

		code := msg[i : i+codeLength]
		if v, ok := m[code]; ok {
			out.WriteRune(v)
		}
	}

	return out.String()
}
