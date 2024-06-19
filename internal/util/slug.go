package util

import (
	"regexp"
	"strings"
	"unicode"
	
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Remove diacritics and normalize the string
func removeDiacritics(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

// Slugify converts a string into a slug
func Slugify(name string) string {
	// Remove diacritics
	name = removeDiacritics(name)
	// Convert to lowercase
	name = strings.ToLower(name)
	// Replace non-alphanumeric characters with hyphens
	re := regexp.MustCompile("[^a-z0-9]+")
	name = re.ReplaceAllString(name, "-")
	// Trim hyphens from the start and end
	name = strings.Trim(name, "-")
	return name
}
