package interpreter

import (
	"strconv"

	parser "github.com/NekoMaru76/rw3/build/RW3C"
	object "github.com/NekoMaru76/rw3/object"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseRW3CListener
	MainScope object.Scope
	Objs      []object.WrappedObject
}

func NewTreeShapeListener() TreeShapeListener {
	return TreeShapeListener{
		MainScope: object.Scope{},
		Objs:      []object.WrappedObject{},
	}
}

func (v *TreeShapeListener) Visit(ctx antlr.Tree, file string) object.WrappedObject {
	switch Obj := ctx.(type) {
	case *parser.ProgContext:
		return v.EnterProg(Obj, file)
	case *parser.Lit_exprContext:
		return v.EnterLit_expr(Obj, file)
	case *parser.ExprContext:
		return v.EnterExpr(Obj, file)
	case *parser.Call_exprContext:
		panic("Call_exprContext")
	case *parser.Chain_exprContext:
		return v.EnterChain_expr(Obj, file)
	case *parser.Access_prop_exprContext:
		panic("Access_prop_exprContext")
	case *parser.Assign_stmtContext:
		panic("Assign_stmtContext")
	case *parser.Atomic_exprContext:
		return v.EnterAtomic_expr(Obj, file)
	case *parser.ArgContext:
		panic("ArgContext")
	case *parser.Def_type_stmtContext:
		panic("Def_type_stmtContext")
	case *parser.Def_var_stmtContext:
		panic("Def_var_stmtContext")
	case *parser.Fn_exprContext:
		panic("Fn_exprContext")
	case *parser.If_exprContext:
		panic("If_exprContext")
	case *parser.If_stmtContext:
		panic("If_stmtContext")
	case *parser.Name_exprContext:
		panic("Name_exprContext")
	case *parser.Ret_stmtContext:
		panic("Ret_stmtContext")
	case *parser.ScopeContext:
		return v.EnterScope(Obj, file)
	case *parser.StmtContext:
		panic("StmtContext")
	case *parser.Struct_stmtContext:
		panic("Struct_stmtContext")
	case *parser.Switch_exprContext:
		panic("Switch_exprContext")
	case *parser.Switch_stmtContext:
		panic("Switch_stmtContext")
	case *parser.Type_argContext:
		panic("Type_argContext")
	case *parser.Type_exprContext:
		panic("Type_exprContext")
	case *parser.While_exprContext:
		panic("While_exprContext")
	case *parser.While_stmtContext:
		panic("While_stmtContext")
	case *parser.Access_prop_type_exprContext:
		panic("Access_prop_type_exprContext")
	case *parser.RunContext:
		return v.EnterRun(Obj, file)
	default:
		panic("Unknown context")
	}
}

func (v *TreeShapeListener) EnterRun(ctx *parser.RunContext, file string) object.WrappedObject {
	return v.Visit(ctx.GetChild(0), file)
}

func (v *TreeShapeListener) EnterScope(ctx *parser.ScopeContext, file string) object.WrappedObject {
	exprs := []object.WrappedObject{}

	for i := 1; i < ctx.GetChildCount()-1; i++ {
		exprs = append(exprs, v.Visit(ctx.GetChild(i), file))
	}

	return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
		null := object.NewNullObject()
		var last *object.Object = &null

		for _, expr := range exprs {
			obj, ok, err := expr(locs, scope)

			if !ok {
				return nil, ok, err
			}

			last = obj
		}

		return last, true, nil
	}
}

func (v *TreeShapeListener) EnterAccess_prop_expr(ctx *parser.Access_prop_exprContext, file string, val object.Object) object.WrappedObject {
	return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
		var key string
		child := ctx.GetChild(0)

		if tok, ok := child.(antlr.Token); ok {
			key = tok.GetText()
		} else {
			Obj, ok, err := v.Visit(child.(antlr.Tree), file)(locs, scope)

			if !ok {
				return nil, ok, err
			}

			key = (*Obj).String()
		}

		return (val).GetProp(key), true, nil
	}
}

func (v *TreeShapeListener) EnterCall_expr(ctx *parser.Call_exprContext, file string, val object.Object) object.WrappedObject {
	exprs := []object.WrappedObject{}

	for i := 1; i < ctx.GetChildCount()-1; i++ {
		exprs = append(exprs, v.Visit(ctx.GetChild(i), file))
	}

	return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
		f, ok := val.(object.FunctionObject)

		if !ok {
			return nil, ok, &object.Error{
				Message:  object.FN_NO_CALLABLE,
				Location: append(locs, object.TokToLoc(ctx.GetStart(), file)),
			}
		}

		args := []object.Object{}

		for i := 0; i < len(exprs); i++ {
			obj, ok, err := exprs[i](locs, scope)

			if !ok {
				return nil, ok, err
			}

			args = append(args, *obj)
		}

		return f.Value(args...)
	}
}

func (v *TreeShapeListener) EnterChain_expr(ctx *parser.Chain_exprContext, file string) object.WrappedObject {
	return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
		obj, ok, err := v.Visit(ctx.Atomic_expr(), file)(locs, scope)

		if !ok {
			return nil, ok, err
		}

		val := *obj

		for i := 1; i < ctx.GetChildCount(); i++ {
			child := ctx.GetChild(i)
			var obj *object.Object
			var ok bool
			var err *object.Error

			switch Obj := child.(type) {
			case *parser.Access_prop_exprContext:
				{
					obj, ok, err = v.EnterAccess_prop_expr(Obj, file, val)(locs, scope)

					break
				}
			case *parser.Call_exprContext:
				{
					obj, ok, err = v.EnterCall_expr(Obj, file, val)(locs, scope)

					break
				}
			}

			if !ok {
				return nil, ok, err
			}

			val = *obj
		}

		return &val, true, nil
	}
}

func (v *TreeShapeListener) EnterAtomic_expr(ctx *parser.Atomic_exprContext, file string) object.WrappedObject {
	return v.Visit(ctx.GetChild(0), file)
}

func (v *TreeShapeListener) EnterExpr(ctx *parser.ExprContext, file string) object.WrappedObject {
	firstChild := ctx.GetChild(0)

	switch tok, ok := firstChild.(antlr.Token); ok {
	case true:
		{
			switch tok.GetText() {
			case "(":
				return v.Visit(ctx.Expr(0), file)
			case "!":
				panic("!")
			case "--":
				panic("--")
			case "++":
				panic("++")
			}
		}
	case false:
		{
			if ctx.GetChildCount() == 1 {
				return v.Visit(firstChild.(antlr.Tree), file)
			}

			secondChild := ctx.GetChild(1)

			switch tok, ok := secondChild.(antlr.Token); ok {
			case true:
				{
					switch tok.GetText() {
					case "++":
						panic("++")
					case "--":
						panic("--")
					case "+":
						panic("+")
					case "-":
						panic("-")
					case "*":
						panic("*")
					case "/":
						panic("/")
					case ">":
						panic(">")
					case "<":
						panic("<")
					case ">=":
						panic(">=")
					case "<=":
						panic("<=")
					case "==":
						panic("==")
					case "!=":
						panic("!=")
					case "&&":
						panic("&&")
					case "||":
						panic("||")
					}
				}
			case false:
				return v.Visit(ctx.Chain_expr(), file)
			}
		}
	}

	panic("")
}

func (v *TreeShapeListener) EnterLit_expr(ctx *parser.Lit_exprContext, file string) object.WrappedObject {
	start := ctx.GetStart()
	loc := object.TokToLoc(start, file)
	var obj object.Object = nil
	var err error

	switch ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType() {
	case parser.RW3CLexerSTR:
		str := ctx.STR().GetText()
		obj = object.StringObject{
			Value: str[1 : len(str)-1],
		}
		break
	case parser.RW3CLexerLONG:
		var val int64

		val, err = strconv.ParseInt(ctx.LONG().GetText(), 0, 10)

		obj = object.LongObject{
			Value: val,
		}
		break
	case parser.RW3CLexerDOUBLE:
		var val float64

		val, err = strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)

		obj = object.DoubleObject{
			Value: val,
		}
		break
	case parser.RW3CLexerBOOL:
		var val bool

		val, err = strconv.ParseBool(ctx.BOOL().GetText())

		obj = object.BooleanObject{
			Value: val,
		}
		break
	case parser.RW3CLexerIDENTIF:
		return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
			stack := []object.Location{loc}
			stack = append(stack, locs...)

			obj, ok, err := scope.GetVar(stack, ctx.IDENTIF().GetText())

			return obj, ok, err
		}
	}

	return func(locs []object.Location, scope object.Scope) (*object.Object, bool, *object.Error) {
		if err != nil {
			return nil, false, &object.Error{
				Message:  err.Error(),
				Location: locs,
			}
		}

		return &obj, true, nil
	}
}

// func (v *TreeShapeListener) Visit(ctx *parser.) Obj {

// }

func (v *TreeShapeListener) EnterProg(ctx *parser.ProgContext, file string) object.WrappedObject {
	for _, tree := range ctx.GetExpr_list() {
		v.Objs = append(v.Objs, v.Visit(tree, file))
	}

	return v.Objs[len(v.Objs)-1]
}

func Interpret(tree antlr.Tree, file string) TreeShapeListener {
	listener := NewTreeShapeListener()
	listener.Visit(tree, file)

	return listener
}
