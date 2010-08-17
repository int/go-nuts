// playing with goyacc
// http://acm.tju.edu.cn/toj/showp3638.html

%{
package main

import (
	"fmt"
	"os"
	"unicode"
)
%}

%union {
	s string;
	val bool;
}

%token WRITE IF ELSE
%token<s> LITERAL
%token<val> COND
%type <s> stmt

%%

list:
    | list stmt { if ($2 != "") {fmt.Println($2)}}
    ;

stmt: WRITE '(' LITERAL ')' { $$ = $3 }
    | IF '(' COND ')' stmt { if ($3) {$$=$5} else {$$=""} }
    | IF '(' COND ')' stmt ELSE stmt { if ($3) {$$=$5} else {$$=$7}}
    ;

%%

func lookup(s string) int {
	if (s == "if") {
		return IF
	}
	if (s == "else") {
		return ELSE
	}
	if (s == "write") {
		return WRITE
	}
	return LITERAL
}

var peek = false
var pchar int
const buflen = 1024
var where = buflen
var buf []byte

func next() int {
	if peek {
		peek = false
		return pchar
	}
	if where < buflen {
		c := buf[where]
		where++
		return int(c)
	}
	n, err := os.Stdin.Read(buf)
	if n == 0 && err == os.EOF {
		return 0
	}
	where = 1
	return int(buf[0])
}

func unget(c int) {
	peek = true
	pchar = c
}

func getword() string {
	c := next()
	s := ""
	for (unicode.IsLower(c)) {
		s += string(c)
		c = next()
	}
	unget(c)
	return s
}

type dummy int

func (_ dummy) Lex(yylval *yySymType) int {
	c := next()
	for (unicode.IsSpace(c)) {
		c = next()
	}
	if (unicode.IsLower(c)) {
		unget(c)
		return lookup(getword())
	}
	if (c == '"') {
		yylval.s = getword()
		next() // '"'
		return LITERAL

	}
	if (c == '1' || c == '0') {
		yylval.val = c == '1'
		return COND
	}
	return c
}

func (_ dummy) Error(s string) {
	fmt.Println("got error", s)
}

func main() {
	buf = make([]byte, buflen)
	yyParse(dummy(0));
}
