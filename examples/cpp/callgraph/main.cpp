
#include "LuaLexer.h"
#include "LuaParser.h"
#include "LuaParserBaseListener.h"
#include "antlr4-runtime.h"

#include <iostream>
#include <map>
#include <set>
#include <string>

using namespace antlr4;

class dot : public LuaParserBaseListener {
  std::set<std::pair<std::string, std::string>> &graph_;
  std::string parent_;

public:
  dot(std::set<std::pair<std::string, std::string>> &graph)
      : graph_(graph), parent_("root") {}

  void
  enterLocalfunctionstat(LuaParser::LocalfunctionstatContext *ctx) override {
    parent_ = ctx->NAME()->getSymbol()->getText();
  }

  void enterFunctionstat(LuaParser::FunctionstatContext *ctx) override {
    std::ostringstream fn;
    auto names = ctx->funcname()->NAME();
    auto dotCount = ctx->funcname()->DOT().size();
    auto hasColon = !!ctx->funcname()->COLON();

    for (auto const &name : names) {
      fn << name->getSymbol()->getText();
      if (dotCount) {
        fn << ".";
        dotCount--;
      } else if (hasColon) {
        fn << ":";
        hasColon = false;
      }
    }
    parent_ = fn.str();
  }

  void exitFunctioncall(LuaParser::FunctioncallContext *ctx) override {
    auto variable = ctx->variable();

    std::string name = variable->getText();
    if (ctx->nameAndArgs()->COLON()) {
        name += ":";
        name += ctx->nameAndArgs()->NAME()->getSymbol()->getText();
    }
    graph_.insert(std::pair<std::string, std::string>(parent_, name));
    return;
  }
};

int main(int argc, const char *argv[]) {
  ANTLRInputStream input(std::cin);
  LuaLexer lexer(&input);
  CommonTokenStream tokens(&lexer);
  LuaParser parser(&tokens);

  auto tree = parser.chunk();

  std::set<std::pair<std::string, std::string>> graph;

  dot listener(graph);
  tree::ParseTreeWalker::DEFAULT.walk(&listener, tree);

  std::set<std::string> nodes;

  for (auto const &edge : graph) {
    nodes.insert(edge.first);
    nodes.insert(edge.second);
  }

  std::cout << "digraph G {" << std::endl;
  std::cout << "  ranksep=.25;" << std::endl;
  std::cout << "  edge [arrowsize=.5]" << std::endl;
  std::cout << "  node [shape=circle, fontname=\"ArialNarrow\"," << std::endl;
  std::cout << "        fontsize=12, fixedsize=false];" << std::endl;
  std::cout << "  ";
  for (auto const &node : nodes) {
    std::cout << std::quoted(node) << "; ";
  }
  std::cout << std::endl;
  for (auto const &edge : graph) {
    std::cout << "  ";
    std::cout << std::quoted(edge.first) << " -> " << std::quoted(edge.second)
              << ";" << std::endl;
  }
  std::cout << "}" << std::endl;

  return 0;
}
