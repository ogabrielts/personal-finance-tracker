package utils

import (
	"fmt"
	"strings"
)

func Tokenize(cmd string) ([]string, error) {
	var tokens []string // Store all tokens
	var current strings.Builder // Current will help hold each character of the current token, until the token is complete
	var quote rune // Quote helps understand whether the code is inside a quoted string
	/*
		quote = 0 -> not inside quotes
		qoute = '"' -> inside double quotes
		quote = '\'' -> inside single qoutes
	*/
	escaped := false // Escape tells whether the previous character is a backslash
	
	// Helper function that adds a token to the array of tokens
	flush := func() {
		if current.Len() > 0 {
			tokens = append(tokens, current.String())
			current.Reset()
		}
	}
	
	// Iterate over the command, character by character
	for _, r := range cmd {
		// If the previous character was a backslash, treat r literally
		if escaped {
			current.WriteRune(r)
			escaped = false
			continue
		}
		
		// Set escape to true, and treat the next r literally
		if r == '\\' {
			escaped = true
			continue
		}
		
		// If inside quotes
		if quote != 0 {
			// If r is equal the opened quote, quoting is complete
			// Else, add the character to the current token
			if r == quote {
				quote = 0
			} else {
				current.WriteRune(r)
			}
			continue
		}
		
		// When not inside quotes
		switch r {
		case '\'', '"': // If r is equal single or double quote, this is a new quoting (a new token)
			quote = r

		case ' ', '\t': // If r is a space or a tab, the token is complete and added to the tokens
			flush()

		default:
			current.WriteRune(r)
		}
	}

	// If the command ends immediately after a backslash, this is an incomplete escape sequence
	if escaped {
		return nil, fmt.Errorf("unfinished scape")
	}

	// If quote is not 0, it means a quote was open but never closed.
	if quote != 0 {
		return nil, fmt.Errorf("unfinished quote")
	}
	
	// Flush one last time in case a token ended without whitespace
	flush()

	return tokens, nil
}
