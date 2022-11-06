package main

import "strings"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	return strings.Join(strs, ",")
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return strings.Split(strs, ",")
}
