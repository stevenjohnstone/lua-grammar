
#include "LuaLexer.h"
#include "LuaParser.h"
#include "LuaParserBaseListener.h"
#include "antlr4-runtime.h"

#include <iostream>
#include <map>
#include <set>
#include <stack>
#include <string>
#include <vector>

using namespace antlr4;

struct Scope {
  Scope *parent;
  std::string source;
  std::set<std::string> symbols;
  std::vector<Scope *> children;
};

class symbol_listener : public LuaParserBaseListener {
  Scope &global_scope_;
  CommonTokenStream &tokens_;
  std::stack<Scope *> parent_;
  std::set<std::string> variables_;

public:
  symbol_listener(Scope &global_scope, CommonTokenStream &tokens)
      : global_scope_(global_scope), tokens_(tokens) {
    parent_.push(&global_scope);
  }

  void enterBlock(LuaParser::BlockContext *ctx) override {
    auto scope = new (Scope);
    scope->source = tokens_.getText(ctx->start, ctx->stop);
    scope->parent = parent_.top();
    scope->parent->children.push_back(scope);
    for (auto const &var : variables_) {
      scope->symbols.insert(var);
    }
    variables_.clear();
    parent_.push(scope);
  }

  void exitBlock(LuaParser::BlockContext *ctx) override { parent_.pop(); }

  void enterVariablestat(LuaParser::VariablestatContext *ctx) override {
    auto variablelist = ctx->variablelist();
    for (auto const &var : variablelist->variable()) {
      // these are all global scope
      global_scope_.symbols.insert(var->getText());
    }
  }

  void enterNumericforstat(LuaParser::NumericforstatContext *ctx) override {
    variables_.insert(ctx->NAME()->getSymbol()->getText());
  }

  void exitNumericforstat(LuaParser::NumericforstatContext *ctx) override {
    variables_.clear();
  }

  void enterGenericforstat(LuaParser::GenericforstatContext *ctx) override {
    for (auto const &name : ctx->namelist()->NAME()) {
      variables_.insert(name->getSymbol()->getText());
    }
  }

  void exitGenericforstat(LuaParser::GenericforstatContext *ctx) override {
    variables_.clear();
  }

  void
  enterLocalfunctionstat(LuaParser::LocalfunctionstatContext *ctx) override {
    parent_.top()->symbols.insert(ctx->NAME()->getSymbol()->getText());
  }

  void enterFuncbody(LuaParser::FuncbodyContext *ctx) override {
    auto parlist = ctx->parlist();
    if (parlist) {
      auto namelist = parlist->namelist();
      if (namelist) {
        for (auto const &name : namelist->NAME()) {
          variables_.insert(name->getSymbol()->getText());
        }
      }
    }
  }

  void exitFuncbody(LuaParser::FuncbodyContext *ctx) override { variables_.clear(); }

  void enterFunctionstat(LuaParser::FunctionstatContext *ctx) override {
    auto funcname = ctx->funcname();

    size_t dot_count = funcname->DOT().size();
    bool has_colon = !!funcname->COLON();

    std::string namestr;

    for (auto const &name : funcname->NAME()) {
      namestr += name->getSymbol()->getText();
      if (dot_count) {
        namestr += ".";
        dot_count--;
      } else if (has_colon) {
        namestr += ":";
        has_colon = false;
      }
    }
    global_scope_.symbols.insert(namestr);
  }

  void enterLocalvariablestat(LuaParser::LocalvariablestatContext *ctx) override {
    auto alist = ctx->attnamelist();
    auto &parent = parent_.top();
    for (auto const &name : alist->NAME()) {
      parent->symbols.insert(name->getSymbol()->getText());
    }
  }
};

static void dump(Scope *s) {
  std::cerr << "scope: " << std::endl << std::quoted(s->source) << std::endl;
  std::cerr << "symbols: ";
  for (auto const &symbol : s->symbols) {
    std::cerr << symbol << " ";
  }
  std::cerr << std::endl;

  for (auto child : s->children) {
    dump(child);
  }
}

int main(int argc, const char *argv[]) {
  ANTLRInputStream input(std::cin);
  LuaLexer lexer(&input);
  CommonTokenStream tokens(&lexer);
  LuaParser parser(&tokens);

  auto tree = parser.chunk();

  Scope global_scope;
  global_scope.source = "global";

  symbol_listener listener(global_scope, tokens);
  tree::ParseTreeWalker::DEFAULT.walk(&listener, tree);

  dump(&global_scope);

  return 0;
}
