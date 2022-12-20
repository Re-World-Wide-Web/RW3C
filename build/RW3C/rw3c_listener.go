// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RW3C

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// RW3CListener is a complete listener for a parse tree produced by RW3CParser.
type RW3CListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterRun is called when entering the run production.
	EnterRun(c *RunContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterChain_expr is called when entering the chain_expr production.
	EnterChain_expr(c *Chain_exprContext)

	// EnterAtomic_expr is called when entering the atomic_expr production.
	EnterAtomic_expr(c *Atomic_exprContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterIf_expr is called when entering the if_expr production.
	EnterIf_expr(c *If_exprContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterSwitch_expr is called when entering the switch_expr production.
	EnterSwitch_expr(c *Switch_exprContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterWhile_expr is called when entering the while_expr production.
	EnterWhile_expr(c *While_exprContext)

	// EnterRet_stmt is called when entering the ret_stmt production.
	EnterRet_stmt(c *Ret_stmtContext)

	// EnterDef_var_stmt is called when entering the def_var_stmt production.
	EnterDef_var_stmt(c *Def_var_stmtContext)

	// EnterDef_type_stmt is called when entering the def_type_stmt production.
	EnterDef_type_stmt(c *Def_type_stmtContext)

	// EnterStruct_stmt is called when entering the struct_stmt production.
	EnterStruct_stmt(c *Struct_stmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterAccess_prop_expr is called when entering the access_prop_expr production.
	EnterAccess_prop_expr(c *Access_prop_exprContext)

	// EnterFn_expr is called when entering the fn_expr production.
	EnterFn_expr(c *Fn_exprContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterAccess_prop_type_expr is called when entering the access_prop_type_expr production.
	EnterAccess_prop_type_expr(c *Access_prop_type_exprContext)

	// EnterType_expr is called when entering the type_expr production.
	EnterType_expr(c *Type_exprContext)

	// EnterType_arg is called when entering the type_arg production.
	EnterType_arg(c *Type_argContext)

	// EnterName_expr is called when entering the name_expr production.
	EnterName_expr(c *Name_exprContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterLit_expr is called when entering the lit_expr production.
	EnterLit_expr(c *Lit_exprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitRun is called when exiting the run production.
	ExitRun(c *RunContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitChain_expr is called when exiting the chain_expr production.
	ExitChain_expr(c *Chain_exprContext)

	// ExitAtomic_expr is called when exiting the atomic_expr production.
	ExitAtomic_expr(c *Atomic_exprContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitIf_expr is called when exiting the if_expr production.
	ExitIf_expr(c *If_exprContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitSwitch_expr is called when exiting the switch_expr production.
	ExitSwitch_expr(c *Switch_exprContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitWhile_expr is called when exiting the while_expr production.
	ExitWhile_expr(c *While_exprContext)

	// ExitRet_stmt is called when exiting the ret_stmt production.
	ExitRet_stmt(c *Ret_stmtContext)

	// ExitDef_var_stmt is called when exiting the def_var_stmt production.
	ExitDef_var_stmt(c *Def_var_stmtContext)

	// ExitDef_type_stmt is called when exiting the def_type_stmt production.
	ExitDef_type_stmt(c *Def_type_stmtContext)

	// ExitStruct_stmt is called when exiting the struct_stmt production.
	ExitStruct_stmt(c *Struct_stmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitAccess_prop_expr is called when exiting the access_prop_expr production.
	ExitAccess_prop_expr(c *Access_prop_exprContext)

	// ExitFn_expr is called when exiting the fn_expr production.
	ExitFn_expr(c *Fn_exprContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitAccess_prop_type_expr is called when exiting the access_prop_type_expr production.
	ExitAccess_prop_type_expr(c *Access_prop_type_exprContext)

	// ExitType_expr is called when exiting the type_expr production.
	ExitType_expr(c *Type_exprContext)

	// ExitType_arg is called when exiting the type_arg production.
	ExitType_arg(c *Type_argContext)

	// ExitName_expr is called when exiting the name_expr production.
	ExitName_expr(c *Name_exprContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitLit_expr is called when exiting the lit_expr production.
	ExitLit_expr(c *Lit_exprContext)
}
