package rscript

import "strings"

func Tokenize(script string) {
	lines := strings.Split(script, "\n")
	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")

		switch words[0] {
			case "//"
				
		}
			
	}
}
