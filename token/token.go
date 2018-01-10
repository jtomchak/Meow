package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + lierals
	IDENT = "IDENT" //add, foobar, x, y
	INT   = "INT"   //1243456

	//Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTERISK = "*"
	SLASH  = "/"

	LT = "<"
	GT = ">"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EQ = "=="
	NOT_EQ = "!="

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE 		 = "TRUE"
	FALSE 	 = "FALSE"
	IF 		   = "IF"
	ELSE 		 = "ELSE"
	RETURN 	 = "RETURN"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if" :IF,
	"else": ELSE,
	"return": RETURN,
}


// LookupIdent  does a lookup for keywords, otherwise it's tokentype will be token.IDENT which is user defined. 
func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
