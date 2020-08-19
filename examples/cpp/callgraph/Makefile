OUTPUT=output
GENERATED=generated

CC=clang++
CCARGS=-I /usr/local/include/antlr4-runtime/ -I $(GENERATED) -std=c++17 -ggdb -lantlr4-runtime

all: callgraph

callgraph: dirs antlr4 main.cpp
	$(CC) $(CCARGS) main.cpp $(GENERATED)/LuaParserBaseListener.cpp $(GENERATED)/LuaParserListener.cpp $(GENERATED)/LuaParser.cpp $(GENERATED)/LuaLexer.cpp -o $@ 

antlr4:
	java -jar /usr/local/lib/antlr-4.8-complete.jar -Dlanguage=Cpp -o $(GENERATED) -lib ../../../ LuaLexer.g4 LuaParser.g4

dirs:
	mkdir -p $(OUTPUT) 
	mkdir -p $(GENERATED) 

clean:
	rm -rf $(OUTPUT)
	rm -rf $(GENERATED)