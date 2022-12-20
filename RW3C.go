package main

import (
	"fmt"
	"os"

	parser "github.com/NekoMaru76/rw3/build/RW3C"
	interpreter "github.com/NekoMaru76/rw3/interpreter"
	"github.com/NekoMaru76/rw3/object"
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
	logName := "log"
	logValue := func(args ...object.Object) (*object.Object, bool, *object.Error) {
		null := object.NewNullObject()
		str := ""

		for i, obj := range args {
			if i != 0 {
				str += " "
			}

			str += obj.String()
		}

		fmt.Println(str)

		return &null, true, nil
	}
	v, ok, e := wrapped([]object.Location{}, object.NewMainScope(object.RootVars{
		Log: object.FunctionObject{
			Name:  &logName,
			Value: logValue,
		},
	}))

	if !ok {
		panic(e.String())
	}

	fmt.Println((*v).String())
}
