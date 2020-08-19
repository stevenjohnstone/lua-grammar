# ANTLRv4 Lua parser and lexer grammars

Modified https://github.com/antlr/grammars-v4/tree/master/lua:

* split lexer and parser grammars
* added more statement types (varstat etc)
* added a COMMENTS channel
* whitespace goes to HIDDEN channel rather than simply skipping

## Usage

As for other ANTLRv4 grammars:

```shell
antlr4 LuaLexer.g4 LuaParser.g4
```

## Railroad Diagrams

See the grammar in the form of railroad diagrams [here](https://stevenjohnstone.github.io/lua-grammar/).
