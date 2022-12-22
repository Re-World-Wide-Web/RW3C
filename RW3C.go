package main

import (
	"fmt"
	"os"

	parser "github.com/NekoMaru76/rw3c/build/RW3C"
	interpreter "github.com/NekoMaru76/rw3c/interpreter"
	"github.com/NekoMaru76/rw3c/object"
	"github.com/NekoMaru76/rw3c/serror"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseRW3CListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewRW3CLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRW3CParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := interpreter.Interpret(tree, input.GetSourceName())
	wrapped := listener.Objs[len(listener.Objs)-1]
	Log := object.FunctionObject{
		Value: func(args ...object.Object) (*object.Object, *serror.Error) {
			null := object.NewNullObject()
			str := ""

			for i, obj := range args {
				if i != 0 {
					str += " "
				}

				str += obj.String()
			}

			fmt.Println(str)

			return &null, nil
		},
	}
	unwrapped := wrapped([]serror.Location{}, object.NewMainScope(object.RootVars{
		Log: Log,
	}))
	e := unwrapped.Test()

	if e != nil {
		panic(e.String())
	}

	v, er := unwrapped.Value()

	if er != nil {
		panic(e.String())
	}

	fmt.Println((*v).String())
}
