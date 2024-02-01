package lexer

import (
	"kong/token"
	"testing"
)

type ExpectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestSimpleCharNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []ExpectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	execute("TestSimpleCharNextToken", t, lexer, tests)
}

func TestSimpleGrammar1NextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []ExpectedToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	execute("TestSimpleGrammar1NextToken", t, lexer, tests)
}

func TestSimpleGrammar2NextToken(t *testing.T) {
	input := `!-/*5;
	5 < 10 > 5;
	`

	tests := []ExpectedToken{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	execute("TestSimpleGrammar2NextToken", t, lexer, tests)
}

func TestSimpleGrammar3NextToken(t *testing.T) {
	input := `if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []ExpectedToken{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := New(input)

	execute("TestSimpleGrammar3NextToken", t, lexer, tests)
}

func TestSimpleGrammar4NextToken(t *testing.T) {
	input := `10 == 10;
	10 != 9;
	`

	tests := []ExpectedToken{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	execute("TestSimpleGrammar4NextToken", t, lexer, tests)
}

func execute(name string, t *testing.T, l *Lexer, tests []ExpectedToken) {
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("%q - tests[%d] - tokentype wrong, expected=%q, got=%q", name, i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("%q - tests[%d] - tokenliteral wrong, expected=%q, got=%q", name, i, tt.expectedLiteral, tok.Literal)
		}
	}
}
