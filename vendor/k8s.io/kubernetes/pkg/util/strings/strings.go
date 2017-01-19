/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package strings

import (
	"fmt"
	"path"
	"strings"
	"unicode"
)

// Splits a fully qualified name and returns its namespace and name.
// Assumes that the input 'str' has been validated.
func SplitQualifiedName(str string) (string, string) {
	parts := strings.Split(str, "/")
	if len(parts) < 2 {
		return "", str
	}
	return parts[0], parts[1]
}

// Joins 'namespace' and 'name' and returns a fully qualified name
// Assumes that the input is valid.
func JoinQualifiedName(namespace, name string) string {
	return path.Join(namespace, name)
}

// Returns the first N slice of a string.
func ShortenString(str string, n int) string {
	if len(str) <= n {
		return str
	} else {
		return str[:n]
	}
}

// GetArticleForNoun returns the article needed for the given noun.
func GetArticleForNoun(noun string, padding string) string {
	if noun[len(noun)-2:] != "ss" && noun[len(noun)-1:] == "s" {
		// Plurals don't have an article.
		// Don't catch words like class
		return fmt.Sprintf("%v", padding)
	}

	article := "a"
	if isVowel(rune(noun[0])) {
		article = "an"
	}

	return fmt.Sprintf("%s%s%s", padding, article, padding)
}

// isVowel returns true if the rune is a vowel (case insensitive).
func isVowel(c rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, value := range vowels {
		if value == unicode.ToLower(c) {
			return true
		}
	}
	return false
}