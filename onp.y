// spoj problem ONP: https://www.spoj.pl/problems/ONP/
%{
package main

import "fmt"
import "unicode"
%}

%union {
	c int
}

%token <c> VAR

%left '+' '-'
%left '*' '/'
%right '^'

%%

list : expr { fmt.Print("\n") }
     ;

expr : expr '+' expr { fmt.Print("+") }
     | expr '-' expr { fmt.Printf("-") }
     | expr '*' expr { fmt.Printf("*") }
     | expr '/' expr { fmt.Printf("/") }
     | expr '^' expr { fmt.Printf("^") }
     | '(' expr ')'
     | VAR { fmt.Printf("%c", $1) }
     ;

%%

type dummy int

var line string;
var cur int;

func (_ dummy) Lex(yylval *yySymType) int {
	if cur == len(line) {
		return 0
	}
	c := int(line[cur])
	cur++
	if (unicode.IsLower(c)) {
		yylval.c = c
		return VAR
	}
	return c
}

func (_ dummy) Error(s string) {
	fmt.Println("got error: ", s)
}

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		cur = 0
		fmt.Scanln(&line)
		yyParse(dummy(0))
	}
}
