package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	BANG     TokenType = "!"
	LT       TokenType = "<"
	GT       TokenType = ">"
	EQ       TokenType = "=="
	NOT_EQ   TokenType = "!="

	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "true"
	FALSE    TokenType = "false"
	IF       TokenType = "if"
	ELSE     TokenType = "else"
	RETURN   TokenType = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

// func (tt TokenType) Humanize() string {
// 	switch tt {
// 	case ILLEGAL:
// 		return "ILLEGAL"
// 	case EOF:
// 		return "EOF"
// 	case IDENT:
// 		return "IDENT"
// 	case INT:
// 		return "INT"
// 	case ASSIGN:
// 		return "="
// 	case PLUS:
// 		return "+"
// 	case COMMA:
// 		return ","
// 	case SEMICOLON:
// 		return ";"
// 	case LPAREN:
// 		return "("
// 	case RPAREN:
// 		return ")"
// 	case LBRACE:
// 		return "{"
// 	case RBRACE:
// 		return "}"
// 	case FUNCTION:
// 		return "FUNCTION"
// 	case LET:
// 		return "LET"
// 	default:
// 		return "UNKNOWN"
// 	}
// }
