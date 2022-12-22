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

	// EnterWhileRun is called when entering the whileRun production.
	EnterWhileRun(c *WhileRunContext)

	// EnterWhileExprBody is called when entering the whileExprBody production.
	EnterWhileExprBody(c *WhileExprBodyContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAtomicExpr is called when entering the atomicExpr production.
	EnterAtomicExpr(c *AtomicExprContext)

	// EnterChainExpr is called when entering the chainExpr production.
	EnterChainExpr(c *ChainExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterAccessPropExpr is called when entering the accessPropExpr production.
	EnterAccessPropExpr(c *AccessPropExprContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterSwitchStmt is called when entering the switchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchExpr is called when entering the switchExpr production.
	EnterSwitchExpr(c *SwitchExprContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterWhileExpr is called when entering the whileExpr production.
	EnterWhileExpr(c *WhileExprContext)

	// EnterRetStmt is called when entering the retStmt production.
	EnterRetStmt(c *RetStmtContext)

	// EnterDefVarStmt is called when entering the defVarStmt production.
	EnterDefVarStmt(c *DefVarStmtContext)

	// EnterDefTypeStmt is called when entering the defTypeStmt production.
	EnterDefTypeStmt(c *DefTypeStmtContext)

	// EnterStructStmt is called when entering the structStmt production.
	EnterStructStmt(c *StructStmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterFnExpr is called when entering the fnExpr production.
	EnterFnExpr(c *FnExprContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterAccessPropTypeExpr is called when entering the accessPropTypeExpr production.
	EnterAccessPropTypeExpr(c *AccessPropTypeExprContext)

	// EnterTypeExpr is called when entering the typeExpr production.
	EnterTypeExpr(c *TypeExprContext)

	// EnterTypeArg is called when entering the typeArg production.
	EnterTypeArg(c *TypeArgContext)

	// EnterNameExpr is called when entering the nameExpr production.
	EnterNameExpr(c *NameExprContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterLitExpr is called when entering the litExpr production.
	EnterLitExpr(c *LitExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitRun is called when exiting the run production.
	ExitRun(c *RunContext)

	// ExitWhileRun is called when exiting the whileRun production.
	ExitWhileRun(c *WhileRunContext)

	// ExitWhileExprBody is called when exiting the whileExprBody production.
	ExitWhileExprBody(c *WhileExprBodyContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAtomicExpr is called when exiting the atomicExpr production.
	ExitAtomicExpr(c *AtomicExprContext)

	// ExitChainExpr is called when exiting the chainExpr production.
	ExitChainExpr(c *ChainExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitAccessPropExpr is called when exiting the accessPropExpr production.
	ExitAccessPropExpr(c *AccessPropExprContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitSwitchStmt is called when exiting the switchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchExpr is called when exiting the switchExpr production.
	ExitSwitchExpr(c *SwitchExprContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitWhileExpr is called when exiting the whileExpr production.
	ExitWhileExpr(c *WhileExprContext)

	// ExitRetStmt is called when exiting the retStmt production.
	ExitRetStmt(c *RetStmtContext)

	// ExitDefVarStmt is called when exiting the defVarStmt production.
	ExitDefVarStmt(c *DefVarStmtContext)

	// ExitDefTypeStmt is called when exiting the defTypeStmt production.
	ExitDefTypeStmt(c *DefTypeStmtContext)

	// ExitStructStmt is called when exiting the structStmt production.
	ExitStructStmt(c *StructStmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitFnExpr is called when exiting the fnExpr production.
	ExitFnExpr(c *FnExprContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitAccessPropTypeExpr is called when exiting the accessPropTypeExpr production.
	ExitAccessPropTypeExpr(c *AccessPropTypeExprContext)

	// ExitTypeExpr is called when exiting the typeExpr production.
	ExitTypeExpr(c *TypeExprContext)

	// ExitTypeArg is called when exiting the typeArg production.
	ExitTypeArg(c *TypeArgContext)

	// ExitNameExpr is called when exiting the nameExpr production.
	ExitNameExpr(c *NameExprContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitLitExpr is called when exiting the litExpr production.
	ExitLitExpr(c *LitExprContext)
}
