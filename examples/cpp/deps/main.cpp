
#include "LuaLexer.h"
#include "LuaParser.h"
#include "LuaParserBaseListener.h"
#include "antlr4-runtime.h"

#include <iostream>
extern "C" {
#include <lua.h>
#include <luaconf.h>
}
#include <map>
#include <set>
#include <string>
#include <vector>

using namespace antlr4;

typedef std::pair<std::string, std::string> edge;
typedef std::set<edge> graph;

class deps : public LuaParserBaseListener {
  const std::string &current_package_;
  std::vector<std::string> &new_deps_;

public:
  deps(const std::string &current_package, std::vector<std::string> &new_deps)
      : current_package_(current_package), new_deps_(new_deps) {}

  void exitFunctioncall(LuaParser::FunctioncallContext *ctx) override {
    auto variable = ctx->variable();
    if (!variable) {
      return;
    }
    std::string new_dep;

    auto fn_name = variable->getText();
    if (fn_name == "pcall") {
      auto explist = ctx->nameAndArgs()->args()->explist();
      if (!explist || explist->exp().size() != 2) {
        return;
      }

      auto variable= explist->exp()[0]->variable();
      if (!variable) {
        return;
      }

      auto pcall_fn = variable->getText();
      if (pcall_fn != "require") {
        return;
      }

      auto string = explist->exp()[1]->lstring();
      if (!string) {
        return;
      }
      auto token = string->NORMALSTRING() ?: string->CHARSTRING();
      if (!token) {
        return;
      }
      new_dep = token->getSymbol()->getText();
    } else if (fn_name == "require") {
      auto string = ctx->nameAndArgs()->args()->lstring();
      if (!string) {
        auto explist = ctx->nameAndArgs()->args()->explist();
        if (!explist) {
          return;
        }
        string = explist->exp()[0]->lstring();

        if (!string) {
          return;
        }
      }

      auto token = string->NORMALSTRING() ?: string->CHARSTRING();
      if (!token) {
        return;
      }

      new_dep = token->getSymbol()->getText();
    } else {
      return;
    }
    // get rid of quotes
    new_dep.erase(new_dep.begin());
    new_dep.erase(new_dep.end() - 1);

    new_deps_.push_back(new_dep);
    return;
  }
};

static std::string lua_path_default = LUA_PATH_DEFAULT;
static std::string lua_cpath_default = LUA_CPATH_DEFAULT;

__attribute__((constructor)) void lua_paths() {
  auto lua_path = getenv("LUA_PATH");
  if (lua_path) {
    lua_path_default = lua_path;
  }
  auto lua_cpath = getenv("LUA_CPATH");
  if (lua_cpath) {
    lua_cpath_default = lua_cpath;
  }
}

static std::optional<std::string> lookup_dep(const std::string &dep) {

  std::vector<std::string> paths = {lua_path_default, lua_cpath_default};

  for (auto path : paths) {
    size_t pos;
    while ((pos = path.find("?")) != std::string::npos) {
      path.replace(pos, 1, dep);
    }

    std::string filename;
    std::istringstream ps(path);
    while (std::getline(ps, filename, ';')) {
      std::ifstream file(filename);
      if (file.is_open()) {
        file.close();
        return filename;
      }
    }
  }
  return std::nullopt;
}

static void get_deps(const std::string &filename, graph &graph,
                     std::set<std::string> &nodes) {

  auto dotpos = filename.find_last_of(".");
  if (dotpos != std::string::npos) {
    auto file_ext = filename.substr(dotpos);
    if (file_ext != ".lua") {
      return;
    }
  }
  std::ifstream file(filename);
  assert(file.is_open());

  ANTLRInputStream input(file);
  LuaLexer lexer(&input);
  CommonTokenStream tokens(&lexer);
  LuaParser parser(&tokens);

  auto tree = parser.chunk();
  std::vector<std::string> new_deps;
  deps listener(filename, new_deps);
  tree::ParseTreeWalker::DEFAULT.walk(&listener, tree);
  file.close();

  for (auto &new_dep : new_deps) {
    size_t pos;
    while ((pos = new_dep.find(".")) != std::string::npos) {
      new_dep.replace(pos, 1, "/");
    }

    auto dep_filename = lookup_dep(new_dep);
    if (dep_filename) {
      const auto recurse = (nodes.find(*dep_filename) == nodes.end());
      nodes.insert(*dep_filename);
      graph.insert(edge(filename, *dep_filename));
      if (recurse) {
        get_deps(*dep_filename, graph, nodes);
      }
    }
  }
}

int main(int argc, const char *argv[]) {
  assert(argc > 1);

  graph graph;
  std::set<std::string> nodes;
  get_deps(argv[1], graph, nodes);

  std::cout << "digraph G {" << std::endl;
  std::cout << "  overlap=false;" << std::endl;
  std::cout << "  splines=true;" << std::endl;
  std::cout << "  node [shape=\"rectangle\"]" << std::endl;
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
