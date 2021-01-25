package main

import (
	"reflect"
	"testing"

	"./parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type fnCall struct {
	receiver string
	args     string
}

// testListener will collect data from the parse tree
type testListener struct {
	*parser.BaseLuaParserListener
	functionReceivers []fnCall
}

func (tl *testListener) EnterFunctioncall(ctx *parser.FunctioncallContext) {
	args := ""
	arguments := ctx.AllNameAndArgs()
	for _, argument := range arguments {
		args += argument.GetText()
	}
	tl.functionReceivers = append(tl.functionReceivers, fnCall{
		receiver: ctx.Variable().GetText(),
		args:     args,
	})
}

func detectFunctions(program string) []fnCall {
	is := antlr.NewInputStream(program)

	lexer := parser.NewLuaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewLuaParser(stream)

	tl := &testListener{}

	antlr.ParseTreeWalkerDefault.Walk(tl, p.Chunk())
	return tl.functionReceivers
}

func TestDetectFunctions(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name string
		args args
		want []fnCall
	}{
		{
			name: "basic",
			args: args{
				program: "local a = fn()",
			},
			want: []fnCall{{receiver: "fn", args: "()"}},
		},
		{
			name: "basic: function result used",
			args: args{
				program: "local a = fn().b",
			},
			want: []fnCall{{receiver: "fn", args: "()"}},
		},
		{
			name: "basic: chained function calls",
			args: args{
				program: "local a = fn().b()",
			},
			want: []fnCall{{receiver: "fn().b", args: "()"}, {receiver: "fn", args: "()"}},
		},
		{
			name: "basic: function calls",
			args: args{
				program: "local r = require \"foo\"",
			},
			want: []fnCall{{receiver: "require", args: "\"foo\""}},
		},
		{
			name: "basic: nested function calls",
			args: args{
				program: "a(b())",
			},
			want: []fnCall{{receiver: "a", args: "(b())"}, {receiver: "b", args: "()"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectFunctions(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}
