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
	args := ctx.NameAndArgs().GetText()
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

func extractComments(program string) []string {
	is := antlr.NewInputStream(program)

	lexer := parser.NewLuaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, parser.LuaLexerCOMMENTS)

	stream.Consume()

	rv := []string{}

	for _, token := range stream.GetAllTokens() {
		switch token.GetTokenType() {
		case parser.LuaParserCOMMENT:
			fallthrough
		case parser.LuaParserLINE_COMMENT:
			rv = append(rv, token.GetText())
		}
	}
	return rv
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
		{
			name: "corner case: no parentheses call",
			args: args{
				program: "fn(\"foo\") \"bar\"",
			},
			want: []fnCall{{receiver: "fn(\"foo\")", args: "\"bar\""}, {receiver: "fn", args: "(\"foo\")"}},
		},
		{
			name: "corner case: calling a function returned by another function",
			args: args{
				program: "fn(\"foo\")(\"bar\")",
			},
			want: []fnCall{{receiver: "fn(\"foo\")", args: "(\"bar\")"}, {receiver: "fn", args: "(\"foo\")"}},
		},
		{
			name: "basic: commented out function",
			args: args{
				program: "--fn()",
			},
			want: nil,
		},
		{
			name: "basic: commented out function 2",
			args: args{
				program: "fn()--fn2()",
			},
			want: []fnCall{{receiver: "fn", args: "()"}},
		},
		{
			name: "basic: commented out arguments",
			args: args{
				program: "fn(--[[not_used]])",
			},
			want: []fnCall{{receiver: "fn", args: "()"}},
		},
		{
			name: "basic: commented out arguments 2",
			args: args{
				program: "fn(--[[not_used]] a)",
			},
			want: []fnCall{{receiver: "fn", args: "(a)"}},
		},
		{
			name: "basic: function result as function argument",
			args: args{
				program: "fn(a, b, fn2(c))",
			},
			want: []fnCall{{receiver: "fn", args: "(a,b,fn2(c))"}, {receiver: "fn2", args: "(c)"}},
		},
		{
			name: "basic: call result of indexing",
			args: args{
				program: "t[a](b)",
			},
			want: []fnCall{{receiver: "t[a]", args: "(b)"}},
		},
		{
			name: "basic: call result of indexing 2",
			args: args{
				program: "t.a(b)",
			},
			want: []fnCall{{receiver: "t.a", args: "(b)"}},
		},
		{
			name: "basic: index with function call result",
			args: args{
				program: "t[a(b)]",
			},
			want: []fnCall{{receiver: "a", args: "(b)"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectFunctions(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectFunctions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_extractComments(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic: single-line comment",
			args: args{
				program: "--this is a comment",
			},
			want: []string{"--this is a comment"},
		},
		{
			name: "basic: multiple single-line comments",
			args: args{
				program: "--this is a comment\n--this is also a comment",
			},
			want: []string{"--this is a comment\n", "--this is also a comment"},
		},
		{
			name: "basic: comment after code",
			args: args{
				program: "fn(a) --this is a comment",
			},
			want: []string{"--this is a comment"},
		},
		{
			name: "corner case: comment inside code",
			args: args{
				program: "fn(a --[[this is a comment]])",
			},
			want: []string{"--[[this is a comment]]"},
		},
		{
			name: "basic: multiline comment",
			args: args{
				program: "--[[\nthis is a\ncomment\n]]",
			},
			want: []string{"--[[\nthis is a\ncomment\n]]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractComments(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractComments() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
