package platypus

import "strings"

// Transformer applies some transformations to key
// before comparing to node.key
type Transformer func(pattern string) string

// NilTransformer returns the given without any change
func NilTransformer(key string) string {
	return key
}

// TrimTrailHash treams of the trailinf slash of cmd.Pattern last key.
func TrimTrailHash(key string) string {
	return strings.TrimSuffix(key, "#")
}
