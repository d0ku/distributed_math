%{
package main

/*
import (
    "os"
    "github.com/d0ku/distributed_math/base"
)
*/

// result is used to store parsing result
var result int
%}

%union {
        val int
}

%type <val> expression simple_expression

%token <val> NUM

%token ADD SUBT DIV MULT LPAREN RPAREN

%%
expression:
    simple_expression
    {
        result = $1
    }
|   NUM
    {
        result = $1
    }

simple_expression:
    expression ADD expression
    {
        $$ = $1 + $3
    }
|   expression SUBT expression
    {
        $$ = $1 - $3
    }
|   expression MULT expression
    {
        $$ = $1 * $3
    }
|   expression DIV expression
    {
        $$ = $1 / $3
    }
|   LPAREN expression RPAREN
    {
        $$ = $2
    }
