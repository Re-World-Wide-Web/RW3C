// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RW3C

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseRW3CListener is a complete listener for a parse tree produced by RW3CParser.
type BaseRW3CListener struct{}

var _ RW3CListener = &BaseRW3CListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRW3CListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRW3CListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRW3CListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRW3CListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseRW3CListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseRW3CListener) ExitProg(ctx *ProgContext) {}

// EnterRun is called when production run is entered.
func (s *BaseRW3CListener) EnterRun(ctx *RunContext) {}

// ExitRun is called when production run is exited.
func (s *BaseRW3CListener) ExitRun(ctx *RunContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseRW3CListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseRW3CListener) ExitExpr(ctx *ExprContext) {}

// EnterChain_expr is called when production chain_expr is entered.
func (s *BaseRW3CListener) EnterChain_expr(ctx *Chain_exprContext) {}

// ExitChain_expr is called when production chain_expr is exited.
func (s *BaseRW3CListener) ExitChain_expr(ctx *Chain_exprContext) {}

// EnterAtomic_expr is called when production atomic_expr is entered.
func (s *BaseRW3CListener) EnterAtomic_expr(ctx *Atomic_exprContext) {}

// ExitAtomic_expr is called when production atomic_expr is exited.
func (s *BaseRW3CListener) ExitAtomic_expr(ctx *Atomic_exprContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseRW3CListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseRW3CListener) ExitStmt(ctx *StmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseRW3CListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseRW3CListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterIf_expr is called when production if_expr is entered.
func (s *BaseRW3CListener) EnterIf_expr(ctx *If_exprContext) {}

// ExitIf_expr is called when production if_expr is exited.
func (s *BaseRW3CListener) ExitIf_expr(ctx *If_exprContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseRW3CListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseRW3CListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterSwitch_expr is called when production switch_expr is entered.
func (s *BaseRW3CListener) EnterSwitch_expr(ctx *Switch_exprContext) {}

// ExitSwitch_expr is called when production switch_expr is exited.
func (s *BaseRW3CListener) ExitSwitch_expr(ctx *Switch_exprContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseRW3CListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseRW3CListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterWhile_expr is called when production while_expr is entered.
func (s *BaseRW3CListener) EnterWhile_expr(ctx *While_exprContext) {}

// ExitWhile_expr is called when production while_expr is exited.
func (s *BaseRW3CListener) ExitWhile_expr(ctx *While_exprContext) {}

// EnterRet_stmt is called when production ret_stmt is entered.
func (s *BaseRW3CListener) EnterRet_stmt(ctx *Ret_stmtContext) {}

// ExitRet_stmt is called when production ret_stmt is exited.
func (s *BaseRW3CListener) ExitRet_stmt(ctx *Ret_stmtContext) {}

// EnterDef_var_stmt is called when production def_var_stmt is entered.
func (s *BaseRW3CListener) EnterDef_var_stmt(ctx *Def_var_stmtContext) {}

// ExitDef_var_stmt is called when production def_var_stmt is exited.
func (s *BaseRW3CListener) ExitDef_var_stmt(ctx *Def_var_stmtContext) {}

// EnterDef_type_stmt is called when production def_type_stmt is entered.
func (s *BaseRW3CListener) EnterDef_type_stmt(ctx *Def_type_stmtContext) {}

// ExitDef_type_stmt is called when production def_type_stmt is exited.
func (s *BaseRW3CListener) ExitDef_type_stmt(ctx *Def_type_stmtContext) {}

// EnterStruct_stmt is called when production struct_stmt is entered.
func (s *BaseRW3CListener) EnterStruct_stmt(ctx *Struct_stmtContext) {}

// ExitStruct_stmt is called when production struct_stmt is exited.
func (s *BaseRW3CListener) ExitStruct_stmt(ctx *Struct_stmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseRW3CListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseRW3CListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterCall_expr is called when production call_expr is entered.
func (s *BaseRW3CListener) EnterCall_expr(ctx *Call_exprContext) {}

// ExitCall_expr is called when production call_expr is exited.
func (s *BaseRW3CListener) ExitCall_expr(ctx *Call_exprContext) {}

// EnterAccess_prop_expr is called when production access_prop_expr is entered.
func (s *BaseRW3CListener) EnterAccess_prop_expr(ctx *Access_prop_exprContext) {}

// ExitAccess_prop_expr is called when production access_prop_expr is exited.
func (s *BaseRW3CListener) ExitAccess_prop_expr(ctx *Access_prop_exprContext) {}

// EnterFn_expr is called when production fn_expr is entered.
func (s *BaseRW3CListener) EnterFn_expr(ctx *Fn_exprContext) {}

// ExitFn_expr is called when production fn_expr is exited.
func (s *BaseRW3CListener) ExitFn_expr(ctx *Fn_exprContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseRW3CListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseRW3CListener) ExitScope(ctx *ScopeContext) {}

// EnterAccess_prop_type_expr is called when production access_prop_type_expr is entered.
func (s *BaseRW3CListener) EnterAccess_prop_type_expr(ctx *Access_prop_type_exprContext) {}

// ExitAccess_prop_type_expr is called when production access_prop_type_expr is exited.
func (s *BaseRW3CListener) ExitAccess_prop_type_expr(ctx *Access_prop_type_exprContext) {}

// EnterType_expr is called when production type_expr is entered.
func (s *BaseRW3CListener) EnterType_expr(ctx *Type_exprContext) {}

// ExitType_expr is called when production type_expr is exited.
func (s *BaseRW3CListener) ExitType_expr(ctx *Type_exprContext) {}

// EnterType_arg is called when production type_arg is entered.
func (s *BaseRW3CListener) EnterType_arg(ctx *Type_argContext) {}

// ExitType_arg is called when production type_arg is exited.
func (s *BaseRW3CListener) ExitType_arg(ctx *Type_argContext) {}

// EnterName_expr is called when production name_expr is entered.
func (s *BaseRW3CListener) EnterName_expr(ctx *Name_exprContext) {}

// ExitName_expr is called when production name_expr is exited.
func (s *BaseRW3CListener) ExitName_expr(ctx *Name_exprContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseRW3CListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseRW3CListener) ExitArg(ctx *ArgContext) {}

// EnterLit_expr is called when production lit_expr is entered.
func (s *BaseRW3CListener) EnterLit_expr(ctx *Lit_exprContext) {}

// ExitLit_expr is called when production lit_expr is exited.
func (s *BaseRW3CListener) ExitLit_expr(ctx *Lit_exprContext) {}
