// Code generated from LuaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // LuaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LuaParserListener is a complete listener for a parse tree produced by LuaParser.
type LuaParserListener interface {
	antlr.ParseTreeListener

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterBreakstat is called when entering the breakstat production.
	EnterBreakstat(c *BreakstatContext)

	// EnterGotostat is called when entering the gotostat production.
	EnterGotostat(c *GotostatContext)

	// EnterDostat is called when entering the dostat production.
	EnterDostat(c *DostatContext)

	// EnterWhilestat is called when entering the whilestat production.
	EnterWhilestat(c *WhilestatContext)

	// EnterRepeatstat is called when entering the repeatstat production.
	EnterRepeatstat(c *RepeatstatContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterGenericforstat is called when entering the genericforstat production.
	EnterGenericforstat(c *GenericforstatContext)

	// EnterNumericforstat is called when entering the numericforstat production.
	EnterNumericforstat(c *NumericforstatContext)

	// EnterFunctionstat is called when entering the functionstat production.
	EnterFunctionstat(c *FunctionstatContext)

	// EnterLocalfunctionstat is called when entering the localfunctionstat production.
	EnterLocalfunctionstat(c *LocalfunctionstatContext)

	// EnterLocalvariablestat is called when entering the localvariablestat production.
	EnterLocalvariablestat(c *LocalvariablestatContext)

	// EnterVariablestat is called when entering the variablestat production.
	EnterVariablestat(c *VariablestatContext)

	// EnterAttnamelist is called when entering the attnamelist production.
	EnterAttnamelist(c *AttnamelistContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterRetstat is called when entering the retstat production.
	EnterRetstat(c *RetstatContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterVariablelist is called when entering the variablelist production.
	EnterVariablelist(c *VariablelistContext)

	// EnterNamelist is called when entering the namelist production.
	EnterNamelist(c *NamelistContext)

	// EnterExplist is called when entering the explist production.
	EnterExplist(c *ExplistContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterNamedvariable is called when entering the namedvariable production.
	EnterNamedvariable(c *NamedvariableContext)

	// EnterParenthesesvariable is called when entering the parenthesesvariable production.
	EnterParenthesesvariable(c *ParenthesesvariableContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterNameAndArgs is called when entering the nameAndArgs production.
	EnterNameAndArgs(c *NameAndArgsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterFunctiondef is called when entering the functiondef production.
	EnterFunctiondef(c *FunctiondefContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterParlist is called when entering the parlist production.
	EnterParlist(c *ParlistContext)

	// EnterTableconstructor is called when entering the tableconstructor production.
	EnterTableconstructor(c *TableconstructorContext)

	// EnterFieldlist is called when entering the fieldlist production.
	EnterFieldlist(c *FieldlistContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldsep is called when entering the fieldsep production.
	EnterFieldsep(c *FieldsepContext)

	// EnterOperatorOr is called when entering the operatorOr production.
	EnterOperatorOr(c *OperatorOrContext)

	// EnterOperatorAnd is called when entering the operatorAnd production.
	EnterOperatorAnd(c *OperatorAndContext)

	// EnterLessthan is called when entering the lessthan production.
	EnterLessthan(c *LessthanContext)

	// EnterGreaterthan is called when entering the greaterthan production.
	EnterGreaterthan(c *GreaterthanContext)

	// EnterLessthanorequal is called when entering the lessthanorequal production.
	EnterLessthanorequal(c *LessthanorequalContext)

	// EnterGreaterthanorequal is called when entering the greaterthanorequal production.
	EnterGreaterthanorequal(c *GreaterthanorequalContext)

	// EnterNotequal is called when entering the notequal production.
	EnterNotequal(c *NotequalContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterOperatorStrcat is called when entering the operatorStrcat production.
	EnterOperatorStrcat(c *OperatorStrcatContext)

	// EnterOperatorAddSub is called when entering the operatorAddSub production.
	EnterOperatorAddSub(c *OperatorAddSubContext)

	// EnterOperatorMulDivMod is called when entering the operatorMulDivMod production.
	EnterOperatorMulDivMod(c *OperatorMulDivModContext)

	// EnterOperatorBitwise is called when entering the operatorBitwise production.
	EnterOperatorBitwise(c *OperatorBitwiseContext)

	// EnterOperatorUnary is called when entering the operatorUnary production.
	EnterOperatorUnary(c *OperatorUnaryContext)

	// EnterOperatorPower is called when entering the operatorPower production.
	EnterOperatorPower(c *OperatorPowerContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterLstring is called when entering the lstring production.
	EnterLstring(c *LstringContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitBreakstat is called when exiting the breakstat production.
	ExitBreakstat(c *BreakstatContext)

	// ExitGotostat is called when exiting the gotostat production.
	ExitGotostat(c *GotostatContext)

	// ExitDostat is called when exiting the dostat production.
	ExitDostat(c *DostatContext)

	// ExitWhilestat is called when exiting the whilestat production.
	ExitWhilestat(c *WhilestatContext)

	// ExitRepeatstat is called when exiting the repeatstat production.
	ExitRepeatstat(c *RepeatstatContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitGenericforstat is called when exiting the genericforstat production.
	ExitGenericforstat(c *GenericforstatContext)

	// ExitNumericforstat is called when exiting the numericforstat production.
	ExitNumericforstat(c *NumericforstatContext)

	// ExitFunctionstat is called when exiting the functionstat production.
	ExitFunctionstat(c *FunctionstatContext)

	// ExitLocalfunctionstat is called when exiting the localfunctionstat production.
	ExitLocalfunctionstat(c *LocalfunctionstatContext)

	// ExitLocalvariablestat is called when exiting the localvariablestat production.
	ExitLocalvariablestat(c *LocalvariablestatContext)

	// ExitVariablestat is called when exiting the variablestat production.
	ExitVariablestat(c *VariablestatContext)

	// ExitAttnamelist is called when exiting the attnamelist production.
	ExitAttnamelist(c *AttnamelistContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitRetstat is called when exiting the retstat production.
	ExitRetstat(c *RetstatContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitVariablelist is called when exiting the variablelist production.
	ExitVariablelist(c *VariablelistContext)

	// ExitNamelist is called when exiting the namelist production.
	ExitNamelist(c *NamelistContext)

	// ExitExplist is called when exiting the explist production.
	ExitExplist(c *ExplistContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitNamedvariable is called when exiting the namedvariable production.
	ExitNamedvariable(c *NamedvariableContext)

	// ExitParenthesesvariable is called when exiting the parenthesesvariable production.
	ExitParenthesesvariable(c *ParenthesesvariableContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitNameAndArgs is called when exiting the nameAndArgs production.
	ExitNameAndArgs(c *NameAndArgsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitFunctiondef is called when exiting the functiondef production.
	ExitFunctiondef(c *FunctiondefContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitParlist is called when exiting the parlist production.
	ExitParlist(c *ParlistContext)

	// ExitTableconstructor is called when exiting the tableconstructor production.
	ExitTableconstructor(c *TableconstructorContext)

	// ExitFieldlist is called when exiting the fieldlist production.
	ExitFieldlist(c *FieldlistContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldsep is called when exiting the fieldsep production.
	ExitFieldsep(c *FieldsepContext)

	// ExitOperatorOr is called when exiting the operatorOr production.
	ExitOperatorOr(c *OperatorOrContext)

	// ExitOperatorAnd is called when exiting the operatorAnd production.
	ExitOperatorAnd(c *OperatorAndContext)

	// ExitLessthan is called when exiting the lessthan production.
	ExitLessthan(c *LessthanContext)

	// ExitGreaterthan is called when exiting the greaterthan production.
	ExitGreaterthan(c *GreaterthanContext)

	// ExitLessthanorequal is called when exiting the lessthanorequal production.
	ExitLessthanorequal(c *LessthanorequalContext)

	// ExitGreaterthanorequal is called when exiting the greaterthanorequal production.
	ExitGreaterthanorequal(c *GreaterthanorequalContext)

	// ExitNotequal is called when exiting the notequal production.
	ExitNotequal(c *NotequalContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitOperatorStrcat is called when exiting the operatorStrcat production.
	ExitOperatorStrcat(c *OperatorStrcatContext)

	// ExitOperatorAddSub is called when exiting the operatorAddSub production.
	ExitOperatorAddSub(c *OperatorAddSubContext)

	// ExitOperatorMulDivMod is called when exiting the operatorMulDivMod production.
	ExitOperatorMulDivMod(c *OperatorMulDivModContext)

	// ExitOperatorBitwise is called when exiting the operatorBitwise production.
	ExitOperatorBitwise(c *OperatorBitwiseContext)

	// ExitOperatorUnary is called when exiting the operatorUnary production.
	ExitOperatorUnary(c *OperatorUnaryContext)

	// ExitOperatorPower is called when exiting the operatorPower production.
	ExitOperatorPower(c *OperatorPowerContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitLstring is called when exiting the lstring production.
	ExitLstring(c *LstringContext)
}
