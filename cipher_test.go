package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseMapping(t *testing.T) {
	mapping := map[rune]string{
		'a': "a",
		'b': "b",
	}
	want := map[string]rune{
		"a": 'a',
		"b": 'b',
	}

	require.EqualValues(t, want, reverseMapping(mapping))
}

func TestEncode(t *testing.T) {
	msg := "The quick brown fox jumps over the lazy dog"
	want := "baabbaabbbaabaabaaaababaaabaaaaaabaababaaaaabbaaababbbababbaabbabaabababbbababbbabaabbabaaabbaaabbbbbaabaabbbabababaabaabaaabbaabbaabbbaabaaababbaaaaabbaabbbaaaaaabbabbbaaabba"

	require.Equal(t, want, Encode(msg))
}

func TestDecode(t *testing.T) {
	msg := "baabbaabbbaabaabaaaababaaabaaaaaabaababaaaaabbaaababbbababbaabbabaabababbbababbbabaabbabaaabbaaabbbbbaabaabbbabababaabaabaaabbaabbaabbbaabaaababbaaaaabbaabbbaaaaaabbabbbaaabba"
	want := "THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG"

	require.Equal(t, want, Decode(msg))
}
