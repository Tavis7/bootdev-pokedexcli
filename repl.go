package main

import "strings"

func cleanInput(text string) []string {
	split_text := strings.Split(text, " ")
	result := []string{}
	for _, t := range split_text {
		if len(t) > 0 {
			result = append(result, t)
		}
	}

	return result
}
