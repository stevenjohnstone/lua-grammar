OUTPUT=output
GENERATED=generated

CC=clang++
CCARGS=-I /usr/local/include/antlr4-runtime/ -I $(GENERATED) -std=c++17 -ggdb -lantlr4-runtime
GRAMMARDIR=$(abspath ../../../)

all: callgraph

callgraph: dirs antlr4 main.cpp
	$(CC) $(CCARGS) main.cpp $(GENERATED)/LuaParserBaseListener.cpp $(GENERATED)/LuaParserListener.cpp $(GENERATED)/LuaParser.cpp $(GENERATED)/LuaLexer.cpp -o $@ 

antlr4:
	java -jar /usr/local/lib/antlr-4.9-complete.jar -Dlanguage=Cpp -o $(GENERATED) $(GRAMMARDIR)/LuaLexer.g4 $(GRAMMARDIR)/LuaParser.g4

dirs:
	mkdir -p $(OUTPUT) 
	mkdir -p $(GENERATED) 

clean:
	rm -rf $(OUTPUT)
	rm -rf $(GENERATED)