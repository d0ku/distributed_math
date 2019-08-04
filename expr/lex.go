
//line lex.rl:1
package main

import (
        "log"
        "fmt"
        "strconv"
)

var lineNumber int64 = 1


//line lex.go:15
var _mos_parser_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 
}

var _mos_parser_key_offsets []byte = []byte{
	0, 0, 11, 
}

var _mos_parser_trans_keys []byte = []byte{
	32, 40, 41, 42, 43, 45, 47, 9, 
	13, 48, 57, 48, 57, 
}

var _mos_parser_single_lengths []byte = []byte{
	0, 7, 0, 
}

var _mos_parser_range_lengths []byte = []byte{
	0, 2, 1, 
}

var _mos_parser_index_offsets []byte = []byte{
	0, 0, 10, 
}

var _mos_parser_trans_targs []byte = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 
	2, 0, 2, 1, 1, 
}

var _mos_parser_trans_actions []byte = []byte{
	17, 13, 15, 9, 5, 7, 11, 17, 
	0, 0, 0, 19, 19, 
}

var _mos_parser_to_state_actions []byte = []byte{
	0, 1, 0, 
}

var _mos_parser_from_state_actions []byte = []byte{
	0, 3, 0, 
}

var _mos_parser_eof_trans []byte = []byte{
	0, 0, 13, 
}

const mos_parser_start int = 1
const mos_parser_first_final int = 1
const mos_parser_error int = 0

const mos_parser_en_main int = 1


//line lex.rl:17


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
    
//line lex.go:87
	{
	 lex.cs = mos_parser_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lex.rl:31
    return lex
}

func debugPrint(a string) {
    log.Println(a)
}

func (lex *parsLex) Lex(out *parsSymType) int {
    eof := lex.pe
    tok := 0
    
//line lex.go:107
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	if  lex.cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_mos_parser_from_state_actions[ lex.cs])
	_nacts = uint(_mos_parser_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _mos_parser_actions[_acts - 1] {
		case 1:
//line NONE:1
 lex.ts = ( lex.p)

//line lex.go:130
		}
	}

	_keys = int(_mos_parser_key_offsets[ lex.cs])
	_trans = int(_mos_parser_index_offsets[ lex.cs])

	_klen = int(_mos_parser_single_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case  lex.data[( lex.p)] < _mos_parser_trans_keys[_mid]:
				_upper = _mid - 1
			case  lex.data[( lex.p)] > _mos_parser_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_mos_parser_range_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case  lex.data[( lex.p)] < _mos_parser_trans_keys[_mid]:
				_upper = _mid - 2
			case  lex.data[( lex.p)] > _mos_parser_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
_eof_trans:
	 lex.cs = int(_mos_parser_trans_targs[_trans])

	if _mos_parser_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_mos_parser_trans_actions[_trans])
	_nacts = uint(_mos_parser_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _mos_parser_actions[_acts-1] {
		case 2:
//line lex.rl:43
 lex.te = ( lex.p)+1
{ tok = ADD; debugPrint("+"); ( lex.p)++; goto _out
}
		case 3:
//line lex.rl:44
 lex.te = ( lex.p)+1
{ tok = SUBT; debugPrint("-"); ( lex.p)++; goto _out
}
		case 4:
//line lex.rl:45
 lex.te = ( lex.p)+1
{ tok = MULT; debugPrint("*"); ( lex.p)++; goto _out
}
		case 5:
//line lex.rl:46
 lex.te = ( lex.p)+1
{ tok = DIV; debugPrint("/"); ( lex.p)++; goto _out
}
		case 6:
//line lex.rl:47
 lex.te = ( lex.p)+1
{ tok = LPAREN; debugPrint("("); ( lex.p)++; goto _out
}
		case 7:
//line lex.rl:48
 lex.te = ( lex.p)+1
{ tok = RPAREN; debugPrint(")"); ( lex.p)++; goto _out
}
		case 8:
//line lex.rl:51
 lex.te = ( lex.p)+1

		case 9:
//line lex.rl:50
 lex.te = ( lex.p)
( lex.p)--
{ tok = NUM; temp, _ := strconv.ParseInt(string(lex.data[lex.ts:lex.te]), 16, 64); out.val = int(temp); debugPrint("NUMBER"); ( lex.p)++; goto _out
}
//line lex.go:239
		}
	}

_again:
	_acts = int(_mos_parser_to_state_actions[ lex.cs])
	_nacts = uint(_mos_parser_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _mos_parser_actions[_acts-1] {
		case 0:
//line NONE:1
 lex.ts = 0

//line lex.go:253
		}
	}

	if  lex.cs == 0 {
		goto _out
	}
	( lex.p)++
	if ( lex.p) != ( lex.pe) {
		goto _resume
	}
	_test_eof: {}
	if ( lex.p) == eof {
		if _mos_parser_eof_trans[ lex.cs] > 0 {
			_trans = int(_mos_parser_eof_trans[ lex.cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line lex.rl:55


    return tok;
}

func (lex *parsLex) Error(e string) {
    fmt.Println("error:", e)
}
