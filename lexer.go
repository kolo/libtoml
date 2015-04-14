package libtoml

import "unicode/utf8"

const eof = -(iota + 1)

type stateFn func(*lexer) stateFn

type lexer struct {
	name   string
	input  string
	start  int
	pos    int
	width  int
	tokens chan token
}

func newLexer(name string, input string) (*lexer, chan token) {
	l := &lexer{
		name:   name,
		input:  input,
		tokens: make(chan token),
	}

	go l.run()
	return l, l.tokens
}

func (l *lexer) run() {
	for state := lexSkip; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(typ tokenType) {
	l.tokens <- token{typ, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	var r rune

	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width

	return r
}

// lexSkip skips input until it reaches a non space symbol or EOF.
func lexSkip(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			l.emit(tokenEOF)
			return nil
		}

		if !isSpace(r) && !isEndOfLine(r) {
			l.backup()
			l.ignore()
			return lexString
		}
	}
}

func lexString(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			l.emit(tokenEOF)
			return nil
		}

		if isSpace(r) || isEndOfLine(r) {
			l.backup()
			break
		}
	}

	if l.pos > l.start {
		l.emit(tokenString)
	}

	return lexSkip
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}
