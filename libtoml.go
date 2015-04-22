package libtoml

import "fmt"

func Parse(input string) []token {
	tokens := []token{}

	_, c := newLexer("toml", input)
	for t := range c {
		fmt.Println(t)
		tokens = append(tokens, t)
	}

	return tokens
}
