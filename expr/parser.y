%{
package main

import (
    "os"
    "github.com/d0ku/distributed_math/base"
)

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
        exp := base.ExpressionOperation{base.ArgumentUnion{Number: $1, IsNumber: true}, base.ArgumentUnion{Number: $3, IsNumber: true}, base.Add}
        $$ = base.SolveExpression(os.Getenv("ADDER_URL"), &exp)
    }
|   expression SUBT expression
    {
        exp := base.ExpressionOperation{base.ArgumentUnion{Number: $1, IsNumber: true}, base.ArgumentUnion{Number: $3, IsNumber: true}, base.Sub}
        $$ = base.SolveExpression(os.Getenv("SUBT_URL"), &exp)
    }
|   expression MULT expression
    {
        exp := base.ExpressionOperation{base.ArgumentUnion{Number: $1, IsNumber: true}, base.ArgumentUnion{Number: $3, IsNumber: true}, base.Mult}
        $$ = base.SolveExpression(os.Getenv("MULT_URL"), &exp)
    }
|   expression DIV expression
    {
        exp := base.ExpressionOperation{base.ArgumentUnion{Number: $1, IsNumber: true}, base.ArgumentUnion{Number: $3, IsNumber: true}, base.Div}
        $$ = base.SolveExpression(os.Getenv("DIV_URL"), &exp)

    }
|   LPAREN expression RPAREN
    {
        $$ = $2
    }
