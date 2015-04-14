package libtoml

import "fmt"

func Parse(input string) {
	_, tokens := newLexer("toml", input)
	for t := range tokens {
		fmt.Println(t)
	}
}
