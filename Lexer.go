package Lexer

type Type int

const (
	WHITESPACE               Type = 0
	LINE_BREAK               Type = 1
	NEWLINE                  Type = 2
	LETTER                   Type = 3
	DIGIT                    Type = 4
	OPERATOR                 Type = 5
	STRING_QUOTE             Type = 6
	SINGLE_LINE_COMMENT      Type = 7
	MULTI_LINE_COMMENT_START Type = 8
	MULTI_LINE_COMMENT_END   Type = 9
	UNKNOWN                  Type = 10
)
