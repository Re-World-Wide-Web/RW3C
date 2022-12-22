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

// EnterWhileRun is called when production whileRun is entered.
func (s *BaseRW3CListener) EnterWhileRun(ctx *WhileRunContext) {}

// ExitWhileRun is called when production whileRun is exited.
func (s *BaseRW3CListener) ExitWhileRun(ctx *WhileRunContext) {}

// EnterWhileExprBody is called when production whileExprBody is entered.
func (s *BaseRW3CListener) EnterWhileExprBody(ctx *WhileExprBodyContext) {}

// ExitWhileExprBody is called when production whileExprBody is exited.
func (s *BaseRW3CListener) ExitWhileExprBody(ctx *WhileExprBodyContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseRW3CListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseRW3CListener) ExitExpr(ctx *ExprContext) {}

// EnterAtomicExpr is called when production atomicExpr is entered.
func (s *BaseRW3CListener) EnterAtomicExpr(ctx *AtomicExprContext) {}

// ExitAtomicExpr is called when production atomicExpr is exited.
func (s *BaseRW3CListener) ExitAtomicExpr(ctx *AtomicExprContext) {}

// EnterChainExpr is called when production chainExpr is entered.
func (s *BaseRW3CListener) EnterChainExpr(ctx *ChainExprContext) {}

// ExitChainExpr is called when production chainExpr is exited.
func (s *BaseRW3CListener) ExitChainExpr(ctx *ChainExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseRW3CListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseRW3CListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterAccessPropExpr is called when production accessPropExpr is entered.
func (s *BaseRW3CListener) EnterAccessPropExpr(ctx *AccessPropExprContext) {}

// ExitAccessPropExpr is called when production accessPropExpr is exited.
func (s *BaseRW3CListener) ExitAccessPropExpr(ctx *AccessPropExprContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseRW3CListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseRW3CListener) ExitStmt(ctx *StmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseRW3CListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseRW3CListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseRW3CListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseRW3CListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseRW3CListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseRW3CListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterSwitchExpr is called when production switchExpr is entered.
func (s *BaseRW3CListener) EnterSwitchExpr(ctx *SwitchExprContext) {}

// ExitSwitchExpr is called when production switchExpr is exited.
func (s *BaseRW3CListener) ExitSwitchExpr(ctx *SwitchExprContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseRW3CListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseRW3CListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterWhileExpr is called when production whileExpr is entered.
func (s *BaseRW3CListener) EnterWhileExpr(ctx *WhileExprContext) {}

// ExitWhileExpr is called when production whileExpr is exited.
func (s *BaseRW3CListener) ExitWhileExpr(ctx *WhileExprContext) {}

// EnterRetStmt is called when production retStmt is entered.
func (s *BaseRW3CListener) EnterRetStmt(ctx *RetStmtContext) {}

// ExitRetStmt is called when production retStmt is exited.
func (s *BaseRW3CListener) ExitRetStmt(ctx *RetStmtContext) {}

// EnterDefVarStmt is called when production defVarStmt is entered.
func (s *BaseRW3CListener) EnterDefVarStmt(ctx *DefVarStmtContext) {}

// ExitDefVarStmt is called when production defVarStmt is exited.
func (s *BaseRW3CListener) ExitDefVarStmt(ctx *DefVarStmtContext) {}

// EnterDefTypeStmt is called when production defTypeStmt is entered.
func (s *BaseRW3CListener) EnterDefTypeStmt(ctx *DefTypeStmtContext) {}

// ExitDefTypeStmt is called when production defTypeStmt is exited.
func (s *BaseRW3CListener) ExitDefTypeStmt(ctx *DefTypeStmtContext) {}

// EnterStructStmt is called when production structStmt is entered.
func (s *BaseRW3CListener) EnterStructStmt(ctx *StructStmtContext) {}

// ExitStructStmt is called when production structStmt is exited.
func (s *BaseRW3CListener) ExitStructStmt(ctx *StructStmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseRW3CListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseRW3CListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterFnExpr is called when production fnExpr is entered.
func (s *BaseRW3CListener) EnterFnExpr(ctx *FnExprContext) {}

// ExitFnExpr is called when production fnExpr is exited.
func (s *BaseRW3CListener) ExitFnExpr(ctx *FnExprContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseRW3CListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseRW3CListener) ExitScope(ctx *ScopeContext) {}

// EnterAccessPropTypeExpr is called when production accessPropTypeExpr is entered.
func (s *BaseRW3CListener) EnterAccessPropTypeExpr(ctx *AccessPropTypeExprContext) {}

// ExitAccessPropTypeExpr is called when production accessPropTypeExpr is exited.
func (s *BaseRW3CListener) ExitAccessPropTypeExpr(ctx *AccessPropTypeExprContext) {}

// EnterTypeExpr is called when production typeExpr is entered.
func (s *BaseRW3CListener) EnterTypeExpr(ctx *TypeExprContext) {}

// ExitTypeExpr is called when production typeExpr is exited.
func (s *BaseRW3CListener) ExitTypeExpr(ctx *TypeExprContext) {}

// EnterTypeArg is called when production typeArg is entered.
func (s *BaseRW3CListener) EnterTypeArg(ctx *TypeArgContext) {}

// ExitTypeArg is called when production typeArg is exited.
func (s *BaseRW3CListener) ExitTypeArg(ctx *TypeArgContext) {}

// EnterNameExpr is called when production nameExpr is entered.
func (s *BaseRW3CListener) EnterNameExpr(ctx *NameExprContext) {}

// ExitNameExpr is called when production nameExpr is exited.
func (s *BaseRW3CListener) ExitNameExpr(ctx *NameExprContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseRW3CListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseRW3CListener) ExitArg(ctx *ArgContext) {}

// EnterLitExpr is called when production litExpr is entered.
func (s *BaseRW3CListener) EnterLitExpr(ctx *LitExprContext) {}

// ExitLitExpr is called when production litExpr is exited.
func (s *BaseRW3CListener) ExitLitExpr(ctx *LitExprContext) {}
