package libtoml

import "fmt"

type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF

	tokenString
)

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}

	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}

	return fmt.Sprintf("%q", t.val)
}
