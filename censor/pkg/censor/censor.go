package censor

import "strings"

var badContent = [...]string{"qwerty", "йцуке", "123"}

// IsCensored returns true if text should be censored and false otherwise
func IsCensored(text string) bool {
	isCensored := false

	for _, subStr := range badContent {
		isCensored = strings.Contains(strings.ToLower(text), strings.ToLower(subStr))
		if isCensored {
			return isCensored
		}
	}

	return isCensored
}
