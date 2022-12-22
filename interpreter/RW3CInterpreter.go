package interpreter

import (
	"reflect"
	"strconv"

	parser "github.com/NekoMaru76/rw3c/build/RW3C"
	object "github.com/NekoMaru76/rw3c/object"
	serror "github.com/NekoMaru76/rw3c/serror"
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
	case *parser.LitExprContext:
		return v.EnterLitExpr(Obj, file)
	case *parser.ExprContext:
		return v.EnterExpr(Obj, file)
	case *parser.CallExprContext:
		panic("CallExprContext")
	case *parser.ChainExprContext:
		return v.EnterChainExpr(Obj, file)
	case *parser.AccessPropExprContext:
		panic("AccessPropExprContext")
	case *parser.AssignStmtContext:
		panic("AssignStmtContext")
	case *parser.AtomicExprContext:
		return v.EnterAtomicExpr(Obj, file)
	case *parser.ArgContext:
		panic("ArgContext")
	case *parser.DefTypeStmtContext:
		panic("DefTypeStmtContext")
	case *parser.DefVarStmtContext:
		panic("Def_varStmtContext")
	case *parser.FnExprContext:
		panic("FnExprContext")
	case *parser.IfExprContext:
		panic("IfExprContext")
	case *parser.IfStmtContext:
		panic("IfStmtContext")
	case *parser.NameExprContext:
		panic("NameExprContext")
	case *parser.RetStmtContext:
		panic("RetStmtContext")
	case *parser.ScopeContext:
		return v.EnterScope(Obj, file)
	case *parser.StmtContext:
		panic("StmtContext")
	case *parser.StructStmtContext:
		panic("StructStmtContext")
	case *parser.SwitchExprContext:
		panic("SwitchExprContext")
	case *parser.SwitchStmtContext:
		panic("SwitchStmtContext")
	case *parser.TypeArgContext:
		panic("TypeArgContext")
	case *parser.TypeExprContext:
		panic("TypeExprContext")
	case *parser.WhileExprContext:
		panic("WhileExprContext")
	case *parser.WhileStmtContext:
		panic("WhileStmtContext")
	case *parser.AccessPropTypeExprContext:
		panic("AccessPropTypeExprContext")
	case *parser.RunContext:
		return v.EnterRun(Obj, file)
	case *parser.WhileExprBodyContext:
		panic("WhileExprBodyContext")
	case *parser.WhileRunContext:
		panic("WhileRunContext")
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

	return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
		unwrapped := []object.UnwrappedObject{}

		for _, expr := range exprs {
			Obj := expr(locs, scope)

			unwrapped = append(unwrapped, Obj)
		}

		Type := object.NullType{}

		return object.UnwrappedObject{
			Raw: ctx.GetText,
			Test: func() *serror.Error {
				var err *serror.Error

				for _, unwrp := range unwrapped {
					err = unwrp.Test()

					if err != nil {
						return err
					}
				}

				return nil
			},
			Value: func() (*object.Object, *serror.Error) {
				null := object.NewNullObject()
				var last *object.Object = &null

				for _, unwrp := range unwrapped {
					v, err := unwrp.Value()

					if err != nil {
						return nil, err
					}

					last = v
				}

				return last, nil
			},
			Type: func() object.ObjectType {
				return Type
			},
		}
	}
}

func (v *TreeShapeListener) EnterAccessPropExpr(ctx *parser.AccessPropExprContext, file string, val object.UnwrappedObject) object.WrappedObject {
	return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
		child := ctx.GetChild(0)

		if tok, ok := child.(antlr.Token); ok {
			key := tok.GetText()

			return object.UnwrappedObject{
				Test: func() *serror.Error {
					return val.TestProp(key)
				},
				Value: func() (*object.Object, *serror.Error) {
					v, _ := val.Value()
					return (*v).GetProp(key), nil
				},
				Type: func() object.ObjectType {
					Type := val.GetPropType(key)
					return Type
				},
				Raw: ctx.GetText,
			}
		} else {
			unwrp := v.Visit(child.(antlr.Tree), file)(locs, scope)

			return object.UnwrappedObject{
				Test: func() *serror.Error {
					err := unwrp.Test()

					if err != nil {
						return err
					}

					v, e := unwrp.Value()

					if e != nil {
						return e
					}

					key := (*v).String()

					return unwrp.TestProp(key)
				},
				Value: func() (*object.Object, *serror.Error) {
					v, e := unwrp.Value()

					if e != nil {
						return nil, e
					}

					key := (*v).String()
					m, _ := val.Value()

					return (*m).GetProp(key), nil
				},
				Type: func() object.ObjectType {
					v, e := unwrp.Value()

					if e != nil {
						return reflect.TypeOf(object.NewNullObject())
					}

					key := (*v).String()
					return val.GetPropType(key)
				},
				Raw: ctx.GetText,
			}
		}

	}
}

func (v *TreeShapeListener) EnterCallExpr(ctx *parser.CallExprContext, file string, val object.UnwrappedObject) object.WrappedObject {
	exprs := []object.WrappedObject{}

	for i := 1; i < ctx.GetChildCount()-1; i++ {
		exprs = append(exprs, v.Visit(ctx.GetChild(i), file))
	}

	return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
		argsU := []object.UnwrappedObject{}

		for _, expr := range exprs {
			argsU = append(argsU, expr(locs, scope))
		}

		return object.UnwrappedObject{
			Test: func() *serror.Error {
				f, ok := val.Type().(object.FunctionType)

				if !ok {
					err := &serror.Error{
						Message:  serror.FN_NO_CALLABLE,
						Location: append(locs, object.TokToLoc(ctx.GetStart(), file)),
					}
					return err
				}

				argsT := []object.ObjectType{}

				for _, argU := range argsU {
					err := argU.Test()

					if err != nil {
						return err
					}

					argsT = append(argsT, argU.Type())
				}

				return f.TestCall(locs, argsT)
			},
			Type: val.Type,
			Value: func() (*object.Object, *serror.Error) {
				f, e := val.Value()

				if e != nil {
					return nil, e
				}

				args := []object.Object{}

				for _, argU := range argsU {
					v, e := argU.Value()

					if e != nil {
						return nil, e
					}

					args = append(args, *v)
				}

				return (*f).(object.FunctionObject).Value(args...)
			},
			Raw: ctx.GetText,
		}
	}
}

func (v *TreeShapeListener) EnterChainExpr(ctx *parser.ChainExprContext, file string) object.WrappedObject {
	return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
		unwrapped := v.Visit(ctx.AtomicExpr(), file)(locs, scope)
		child := ctx.GetChild(1)
		var unwrp object.UnwrappedObject

		switch Obj := child.(type) {
		case *parser.AccessPropExprContext:
			{
				unwrp = v.EnterAccessPropExpr(Obj, file, unwrapped)(locs, scope)

				break
			}
		case *parser.CallExprContext:
			{
				unwrp = v.EnterCallExpr(Obj, file, unwrapped)(locs, scope)

				break
			}
		}

		return object.UnwrappedObject{
			Test: func() *serror.Error {
				return unwrp.Test()
			},
			Value: func() (*object.Object, *serror.Error) {
				return unwrp.Value()
			},
			Type: unwrp.Type,
			Raw:  ctx.GetText,
		}
	}
}

func (v *TreeShapeListener) EnterAtomicExpr(ctx *parser.AtomicExprContext, file string) object.WrappedObject {
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
				return v.Visit(ctx.ChainExpr(), file)
			}
		}
	}

	panic("")
}

func (v *TreeShapeListener) EnterLitExpr(ctx *parser.LitExprContext, file string) object.WrappedObject {
	var obj object.Object
	var err error
	var typ object.ObjectType
	loc := object.TokToLoc(ctx.GetStart(), file)
	stack := []serror.Location{loc}

	switch ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType() {
	case parser.RW3CLexerNULL:
		obj = object.NewNullObject()
		typ = object.NullType{}
		break
	case parser.RW3CLexerSTR:
		typ = object.StringType{}
		str := ctx.STR().GetText()
		obj = object.StringObject{
			Value: str[1 : len(str)-1],
		}
		break
	case parser.RW3CLexerLONG:
		var val int64

		val, err = strconv.ParseInt(ctx.LONG().GetText(), 0, 10)

		typ = object.LongType{}
		obj = object.LongObject{
			Value: val,
		}
		break
	case parser.RW3CLexerDOUBLE:
		var val float64

		val, err = strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)

		typ = object.DoubleType{}
		obj = object.DoubleObject{
			Value: val,
		}
		break
	case parser.RW3CLexerBOOL:
		var val bool

		val, err = strconv.ParseBool(ctx.BOOL().GetText())

		typ = object.BooleanType{}
		obj = object.BooleanObject{
			Value: val,
		}
		break
	case parser.RW3CLexerIDENTIF:
		return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
			key := ctx.GetText()
			stack = append(stack, locs...)

			return object.UnwrappedObject{
				Test: func() *serror.Error {
					isDefined := scope.HasVar(key)

					if !isDefined {
						return &serror.Error{
							Location: stack,
							Message:  serror.VAR_NO_EXIST,
						}
					}

					return nil
				},
				Value: func() (*object.Object, *serror.Error) {
					return scope.GetVar(stack, key)
				},
				Type: func() object.ObjectType {
					return *scope.GetVarType(key)
				},
				Raw: ctx.GetText,
			}
		}
	}

	return func(locs []serror.Location, scope object.Scope) object.UnwrappedObject {
		return object.UnwrappedObject{
			Test: func() *serror.Error {
				return nil
			},
			Value: func() (*object.Object, *serror.Error) {
				if err != nil {
					stack = append(stack, locs...)

					return nil, &serror.Error{
						Location: stack,
						Message:  err.Error(),
					}
				}

				return &obj, nil
			},
			Type: func() object.ObjectType {
				return typ
			},
			Raw: ctx.GetText,
		}
	}
}

// func (v *TreeShapeListener) Visit(ctx *parser.) Obj {

// }

func (v *TreeShapeListener) EnterProg(ctx *parser.ProgContext, file string) object.WrappedObject {
	for _, tree := range ctx.GetExprList() {
		v.Objs = append(v.Objs, v.Visit(tree, file))
	}

	return v.Objs[len(v.Objs)-1]
}

func Interpret(tree antlr.Tree, file string) TreeShapeListener {
	listener := NewTreeShapeListener()
	listener.Visit(tree, file)

	return listener
}
