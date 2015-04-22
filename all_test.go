package libtoml

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var tokenName = map[tokenType]string{
	tokenEOF:           "eof",
	tokenString:        "string",
	tokenComment:       "comment",
	tokenTable:         "table",
	tokenArrayOfTables: "array of tables",
}

func (typ tokenType) String() string {
	s := tokenName[typ]
	if s == "" {
		return fmt.Sprintf("item%d", int(typ))
	}

	return s
}

func Test_parseExample(t *testing.T) {
	example, err := ioutil.ReadFile("example.toml")
	if err != nil {
		t.Fatal("can't read example.toml")
	}

	tokens := Parse(string(example))
	for _, t := range tokens {
		fmt.Printf("<%v>:%s\n", t.typ, t)
	}
}
