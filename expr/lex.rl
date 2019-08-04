package main

import (
        "log"
        "fmt"
        "strconv"
)

var lineNumber int64 = 1

%%{ 
    machine mos_parser;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type parsLex struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

func newLexer(data []byte) *parsLex {
    lex := &parsLex{ 
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}

func debugPrint(a string) {
    log.Println(a)
}

func (lex *parsLex) Lex(out *parsSymType) int {
    eof := lex.pe
    tok := 0
    %%{ 
        main := |*
            '+' =>    { tok = ADD; debugPrint("+"); fbreak;};
            '-' =>    { tok = SUBT; debugPrint("-"); fbreak;};
            '*' =>    { tok = MULT; debugPrint("*"); fbreak;};
            '/' =>    { tok = DIV; debugPrint("/"); fbreak;};
            '(' =>    { tok = LPAREN; debugPrint("("); fbreak;};
            ')' =>    { tok = RPAREN; debugPrint(")"); fbreak;};

            digit+ => { tok = NUM; temp, _ := strconv.ParseInt(string(lex.data[lex.ts:lex.te]), 16, 64); out.val = int(temp); debugPrint("NUMBER"); fbreak;};
            space;
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *parsLex) Error(e string) {
    fmt.Println("error:", e)
}
