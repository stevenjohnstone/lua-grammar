// Code generated from LuaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // LuaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLuaParserListener is a complete listener for a parse tree produced by LuaParser.
type BaseLuaParserListener struct{}

var _ LuaParserListener = &BaseLuaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLuaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLuaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLuaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLuaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseLuaParserListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseLuaParserListener) ExitChunk(ctx *ChunkContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLuaParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLuaParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseLuaParserListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseLuaParserListener) ExitStat(ctx *StatContext) {}

// EnterBreakstat is called when production breakstat is entered.
func (s *BaseLuaParserListener) EnterBreakstat(ctx *BreakstatContext) {}

// ExitBreakstat is called when production breakstat is exited.
func (s *BaseLuaParserListener) ExitBreakstat(ctx *BreakstatContext) {}

// EnterGotostat is called when production gotostat is entered.
func (s *BaseLuaParserListener) EnterGotostat(ctx *GotostatContext) {}

// ExitGotostat is called when production gotostat is exited.
func (s *BaseLuaParserListener) ExitGotostat(ctx *GotostatContext) {}

// EnterDostat is called when production dostat is entered.
func (s *BaseLuaParserListener) EnterDostat(ctx *DostatContext) {}

// ExitDostat is called when production dostat is exited.
func (s *BaseLuaParserListener) ExitDostat(ctx *DostatContext) {}

// EnterWhilestat is called when production whilestat is entered.
func (s *BaseLuaParserListener) EnterWhilestat(ctx *WhilestatContext) {}

// ExitWhilestat is called when production whilestat is exited.
func (s *BaseLuaParserListener) ExitWhilestat(ctx *WhilestatContext) {}

// EnterRepeatstat is called when production repeatstat is entered.
func (s *BaseLuaParserListener) EnterRepeatstat(ctx *RepeatstatContext) {}

// ExitRepeatstat is called when production repeatstat is exited.
func (s *BaseLuaParserListener) ExitRepeatstat(ctx *RepeatstatContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaseLuaParserListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaseLuaParserListener) ExitIfstat(ctx *IfstatContext) {}

// EnterGenericforstat is called when production genericforstat is entered.
func (s *BaseLuaParserListener) EnterGenericforstat(ctx *GenericforstatContext) {}

// ExitGenericforstat is called when production genericforstat is exited.
func (s *BaseLuaParserListener) ExitGenericforstat(ctx *GenericforstatContext) {}

// EnterNumericforstat is called when production numericforstat is entered.
func (s *BaseLuaParserListener) EnterNumericforstat(ctx *NumericforstatContext) {}

// ExitNumericforstat is called when production numericforstat is exited.
func (s *BaseLuaParserListener) ExitNumericforstat(ctx *NumericforstatContext) {}

// EnterFunctionstat is called when production functionstat is entered.
func (s *BaseLuaParserListener) EnterFunctionstat(ctx *FunctionstatContext) {}

// ExitFunctionstat is called when production functionstat is exited.
func (s *BaseLuaParserListener) ExitFunctionstat(ctx *FunctionstatContext) {}

// EnterLocalfunctionstat is called when production localfunctionstat is entered.
func (s *BaseLuaParserListener) EnterLocalfunctionstat(ctx *LocalfunctionstatContext) {}

// ExitLocalfunctionstat is called when production localfunctionstat is exited.
func (s *BaseLuaParserListener) ExitLocalfunctionstat(ctx *LocalfunctionstatContext) {}

// EnterLocalvariablestat is called when production localvariablestat is entered.
func (s *BaseLuaParserListener) EnterLocalvariablestat(ctx *LocalvariablestatContext) {}

// ExitLocalvariablestat is called when production localvariablestat is exited.
func (s *BaseLuaParserListener) ExitLocalvariablestat(ctx *LocalvariablestatContext) {}

// EnterVariablestat is called when production variablestat is entered.
func (s *BaseLuaParserListener) EnterVariablestat(ctx *VariablestatContext) {}

// ExitVariablestat is called when production variablestat is exited.
func (s *BaseLuaParserListener) ExitVariablestat(ctx *VariablestatContext) {}

// EnterAttnamelist is called when production attnamelist is entered.
func (s *BaseLuaParserListener) EnterAttnamelist(ctx *AttnamelistContext) {}

// ExitAttnamelist is called when production attnamelist is exited.
func (s *BaseLuaParserListener) ExitAttnamelist(ctx *AttnamelistContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseLuaParserListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseLuaParserListener) ExitAttrib(ctx *AttribContext) {}

// EnterRetstat is called when production retstat is entered.
func (s *BaseLuaParserListener) EnterRetstat(ctx *RetstatContext) {}

// ExitRetstat is called when production retstat is exited.
func (s *BaseLuaParserListener) ExitRetstat(ctx *RetstatContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseLuaParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseLuaParserListener) ExitLabel(ctx *LabelContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseLuaParserListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseLuaParserListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterVariablelist is called when production variablelist is entered.
func (s *BaseLuaParserListener) EnterVariablelist(ctx *VariablelistContext) {}

// ExitVariablelist is called when production variablelist is exited.
func (s *BaseLuaParserListener) ExitVariablelist(ctx *VariablelistContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseLuaParserListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseLuaParserListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExplist is called when production explist is entered.
func (s *BaseLuaParserListener) EnterExplist(ctx *ExplistContext) {}

// ExitExplist is called when production explist is exited.
func (s *BaseLuaParserListener) ExitExplist(ctx *ExplistContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseLuaParserListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseLuaParserListener) ExitExp(ctx *ExpContext) {}

// EnterNamedvariable is called when production namedvariable is entered.
func (s *BaseLuaParserListener) EnterNamedvariable(ctx *NamedvariableContext) {}

// ExitNamedvariable is called when production namedvariable is exited.
func (s *BaseLuaParserListener) ExitNamedvariable(ctx *NamedvariableContext) {}

// EnterParenthesesvariable is called when production parenthesesvariable is entered.
func (s *BaseLuaParserListener) EnterParenthesesvariable(ctx *ParenthesesvariableContext) {}

// ExitParenthesesvariable is called when production parenthesesvariable is exited.
func (s *BaseLuaParserListener) ExitParenthesesvariable(ctx *ParenthesesvariableContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseLuaParserListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseLuaParserListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseLuaParserListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseLuaParserListener) ExitIndex(ctx *IndexContext) {}

// EnterNameAndArgs is called when production nameAndArgs is entered.
func (s *BaseLuaParserListener) EnterNameAndArgs(ctx *NameAndArgsContext) {}

// ExitNameAndArgs is called when production nameAndArgs is exited.
func (s *BaseLuaParserListener) ExitNameAndArgs(ctx *NameAndArgsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseLuaParserListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseLuaParserListener) ExitArgs(ctx *ArgsContext) {}

// EnterFunctiondef is called when production functiondef is entered.
func (s *BaseLuaParserListener) EnterFunctiondef(ctx *FunctiondefContext) {}

// ExitFunctiondef is called when production functiondef is exited.
func (s *BaseLuaParserListener) ExitFunctiondef(ctx *FunctiondefContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseLuaParserListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseLuaParserListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterParlist is called when production parlist is entered.
func (s *BaseLuaParserListener) EnterParlist(ctx *ParlistContext) {}

// ExitParlist is called when production parlist is exited.
func (s *BaseLuaParserListener) ExitParlist(ctx *ParlistContext) {}

// EnterTableconstructor is called when production tableconstructor is entered.
func (s *BaseLuaParserListener) EnterTableconstructor(ctx *TableconstructorContext) {}

// ExitTableconstructor is called when production tableconstructor is exited.
func (s *BaseLuaParserListener) ExitTableconstructor(ctx *TableconstructorContext) {}

// EnterFieldlist is called when production fieldlist is entered.
func (s *BaseLuaParserListener) EnterFieldlist(ctx *FieldlistContext) {}

// ExitFieldlist is called when production fieldlist is exited.
func (s *BaseLuaParserListener) ExitFieldlist(ctx *FieldlistContext) {}

// EnterField is called when production field is entered.
func (s *BaseLuaParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLuaParserListener) ExitField(ctx *FieldContext) {}

// EnterFieldsep is called when production fieldsep is entered.
func (s *BaseLuaParserListener) EnterFieldsep(ctx *FieldsepContext) {}

// ExitFieldsep is called when production fieldsep is exited.
func (s *BaseLuaParserListener) ExitFieldsep(ctx *FieldsepContext) {}

// EnterOperatorOr is called when production operatorOr is entered.
func (s *BaseLuaParserListener) EnterOperatorOr(ctx *OperatorOrContext) {}

// ExitOperatorOr is called when production operatorOr is exited.
func (s *BaseLuaParserListener) ExitOperatorOr(ctx *OperatorOrContext) {}

// EnterOperatorAnd is called when production operatorAnd is entered.
func (s *BaseLuaParserListener) EnterOperatorAnd(ctx *OperatorAndContext) {}

// ExitOperatorAnd is called when production operatorAnd is exited.
func (s *BaseLuaParserListener) ExitOperatorAnd(ctx *OperatorAndContext) {}

// EnterLessthan is called when production lessthan is entered.
func (s *BaseLuaParserListener) EnterLessthan(ctx *LessthanContext) {}

// ExitLessthan is called when production lessthan is exited.
func (s *BaseLuaParserListener) ExitLessthan(ctx *LessthanContext) {}

// EnterGreaterthan is called when production greaterthan is entered.
func (s *BaseLuaParserListener) EnterGreaterthan(ctx *GreaterthanContext) {}

// ExitGreaterthan is called when production greaterthan is exited.
func (s *BaseLuaParserListener) ExitGreaterthan(ctx *GreaterthanContext) {}

// EnterLessthanorequal is called when production lessthanorequal is entered.
func (s *BaseLuaParserListener) EnterLessthanorequal(ctx *LessthanorequalContext) {}

// ExitLessthanorequal is called when production lessthanorequal is exited.
func (s *BaseLuaParserListener) ExitLessthanorequal(ctx *LessthanorequalContext) {}

// EnterGreaterthanorequal is called when production greaterthanorequal is entered.
func (s *BaseLuaParserListener) EnterGreaterthanorequal(ctx *GreaterthanorequalContext) {}

// ExitGreaterthanorequal is called when production greaterthanorequal is exited.
func (s *BaseLuaParserListener) ExitGreaterthanorequal(ctx *GreaterthanorequalContext) {}

// EnterNotequal is called when production notequal is entered.
func (s *BaseLuaParserListener) EnterNotequal(ctx *NotequalContext) {}

// ExitNotequal is called when production notequal is exited.
func (s *BaseLuaParserListener) ExitNotequal(ctx *NotequalContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseLuaParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseLuaParserListener) ExitEqual(ctx *EqualContext) {}

// EnterOperatorStrcat is called when production operatorStrcat is entered.
func (s *BaseLuaParserListener) EnterOperatorStrcat(ctx *OperatorStrcatContext) {}

// ExitOperatorStrcat is called when production operatorStrcat is exited.
func (s *BaseLuaParserListener) ExitOperatorStrcat(ctx *OperatorStrcatContext) {}

// EnterOperatorAddSub is called when production operatorAddSub is entered.
func (s *BaseLuaParserListener) EnterOperatorAddSub(ctx *OperatorAddSubContext) {}

// ExitOperatorAddSub is called when production operatorAddSub is exited.
func (s *BaseLuaParserListener) ExitOperatorAddSub(ctx *OperatorAddSubContext) {}

// EnterOperatorMulDivMod is called when production operatorMulDivMod is entered.
func (s *BaseLuaParserListener) EnterOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// ExitOperatorMulDivMod is called when production operatorMulDivMod is exited.
func (s *BaseLuaParserListener) ExitOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// EnterOperatorBitwise is called when production operatorBitwise is entered.
func (s *BaseLuaParserListener) EnterOperatorBitwise(ctx *OperatorBitwiseContext) {}

// ExitOperatorBitwise is called when production operatorBitwise is exited.
func (s *BaseLuaParserListener) ExitOperatorBitwise(ctx *OperatorBitwiseContext) {}

// EnterOperatorUnary is called when production operatorUnary is entered.
func (s *BaseLuaParserListener) EnterOperatorUnary(ctx *OperatorUnaryContext) {}

// ExitOperatorUnary is called when production operatorUnary is exited.
func (s *BaseLuaParserListener) ExitOperatorUnary(ctx *OperatorUnaryContext) {}

// EnterOperatorPower is called when production operatorPower is entered.
func (s *BaseLuaParserListener) EnterOperatorPower(ctx *OperatorPowerContext) {}

// ExitOperatorPower is called when production operatorPower is exited.
func (s *BaseLuaParserListener) ExitOperatorPower(ctx *OperatorPowerContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseLuaParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseLuaParserListener) ExitNumber(ctx *NumberContext) {}

// EnterLstring is called when production lstring is entered.
func (s *BaseLuaParserListener) EnterLstring(ctx *LstringContext) {}

// ExitLstring is called when production lstring is exited.
func (s *BaseLuaParserListener) ExitLstring(ctx *LstringContext) {}
