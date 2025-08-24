package compiler

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenIdentifier
	TokenKeyword
	TokenSymbol
	TokenWhiteSpace
)

type Token struct {
	Type  TokenType
	value string
}

type Lexer interface {
	NextToken() string
}

type LexerSimple struct {
	input    string
	position int
}

func NewLexer(str string) *LexerSimple {
	return &LexerSimple{input: str, position: 0}
}

func (l *LexerSimple) consumeWhiteSpace() {
	for l.position < len(l.input) && (l.input == " " || l.input == "\n" || l.input == "\t") {
		l.position++
	}
}
func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *LexerSimple) consumeIdentifier() Token {
	start := l.position
	for l.position < len(l.input) && (isLetter(l.input[l.position]) || isNumber(l.input[l.position])) {
		l.position++
	}
	value := l.input[start:l.position]
	return Token{Type: TokenIdentifier, value: value}
}

func (l *LexerSimple) NextToken() Token {
	if l.position >= len(l.input) {
		return Token{Type: TokenEOF, value: ""}
	}
	c := l.input[l.position]
	switch {
	case c == ' ' || c == '\n' || c == '\t':
		l.consumeWhiteSpace()
		return l.NextToken()
	case c == ',' || c == ';' || c == '*' || c == '(' || c == ')' || c == '\'':
		l.position++
		return Token{Type: TokenSymbol, value: string(c)}
	case isLetter(c) || isNumber(c):
		return l.consumeIdentifier()
	default:
		return Token{Type: TokenError, value: string(c)}
	}
}
