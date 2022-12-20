package object

import (
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ObjectType string

const (
	STRING_TYPE   = "STRING"
	INTEGER_TYPE  = "INTEGER"
	LONG_TYPE     = "LONG"
	FLOAT_TYPE    = "FLOAT"
	DOUBLE_TYPE   = "DOUBLE"
	NULL_TYPE     = "NULL"
	BOOLEAN_TYPE  = "BOOLEAN"
	FUNCTION_TYPE = "FUNCTION"
)

const (
	VAR_NO_EXIST   = "Var does not exist"
	FN_NO_CALLABLE = "Function is not callable"
	ERR_RT         = "Runtime"
)

type Object interface {
	Type() ObjectType
	String() string
	GetProp(key string) *Object
	SetProp(key string, value *Object)
}

type BaseObject struct {
	Object,
	Props map[string]*Object
}

func (v BaseObject) GetProp(key string) *Object {
	return v.Props[key]
}

func (v BaseObject) SetProp(key string, value *Object) {
	v.Props[key] = value
}

type WrappedObject = func(locs []Location, scope Scope) (*Object, bool, *Error)

func NewStringObject(Value any) Object {
	V, ok := Value.(string)

	if !ok {
		V = ""
	}

	return StringObject{
		Value: V,
	}
}

type StringObject struct {
	BaseObject
	Value string
}

func (v StringObject) Type() ObjectType {
	return STRING_TYPE
}

func (v StringObject) String() string {
	return "\"" + strings.Join(strings.Split(strings.Join(strings.Split(v.Value, "\\"), "\\\\"), "\""), "\\\"") + "\""
}

func NewNullObject() Object {
	return NullObject{}
}

type NullObject struct {
	BaseObject
}

func (v NullObject) Type() ObjectType {
	return NULL_TYPE
}

func (v NullObject) String() string {
	return "null"
}

func NewBooleanObject(Value any) Object {
	V, ok := Value.(bool)

	if !ok {
		V = false
	}

	return BooleanObject{
		Value: V,
	}
}

type BooleanObject struct {
	BaseObject
	Value bool
}

func (v BooleanObject) Type() ObjectType {
	return BOOLEAN_TYPE
}

func (v BooleanObject) String() string {
	return strconv.FormatBool(v.Value)
}

func NewIntegerObject(Value any) Object {
	V, ok := Value.(int)

	if !ok {
		V = 0
	}

	return IntegerObject{
		Value: V,
	}
}

type IntegerObject struct {
	BaseObject
	Value int
}

func (v IntegerObject) Type() ObjectType {
	return INTEGER_TYPE
}

func (v IntegerObject) String() string {
	return strconv.FormatInt(int64(v.Value), 18)
}

func NewLongObject(Value any) Object {
	V, ok := Value.(int64)

	if !ok {
		V = 0
	}

	return LongObject{
		Value: V,
	}
}

type LongObject struct {
	BaseObject
	Value int64
}

func (v LongObject) Type() ObjectType {
	return LONG_TYPE
}

func (v LongObject) String() string {
	return strconv.FormatInt(v.Value, 36)
}

func NewFloatObject(Value any) Object {
	V, ok := Value.(float32)

	if !ok {
		V = 0
	}

	return FloatObject{
		Value: V,
	}
}

type FloatObject struct {
	BaseObject
	Value float32
}

func (v FloatObject) Type() ObjectType {
	return FLOAT_TYPE
}

func (v FloatObject) String() string {
	return strconv.FormatFloat(float64(v.Value), 4, 32, 32)
}

func NewDoubleObject(Value any) Object {
	V, ok := Value.(float64)

	if !ok {
		V = 0
	}

	return DoubleObject{
		Value: V,
	}
}

type DoubleObject struct {
	BaseObject
	Value float64
}

func (v DoubleObject) Type() ObjectType {
	return DOUBLE_TYPE
}

func (v DoubleObject) String() string {
	return strconv.FormatFloat(v.Value, 8, 64, 64)
}

type FunctionObject struct {
	BaseObject
	Value func(...Object) (*Object, bool, *Error)
	Name  *string
}

func (v FunctionObject) Type() ObjectType {
	return FUNCTION_TYPE
}

func (v FunctionObject) String() string {
	return `Fn`
}

type Scope struct {
	Vars        map[string]Object
	Types       map[string]func(any) Object
	ParentScope *Scope
	Name        string
}

type Error struct {
	Location []Location
	Name     string `default:"Not Found"`
	Message  string
}

type Location struct {
	Line   int
	Column int
	File   string
}

func TokToLoc(tok antlr.Token, file string) Location {
	return Location{
		Column: tok.GetColumn(),
		Line:   tok.GetLine(),
		File:   file,
	}
}

func (loc Location) String() string {
	return loc.File + " " + strconv.Itoa(loc.Line) + ":" + strconv.Itoa(loc.Column)
}

func (err Error) String() string {
	str := err.Name + ": " + err.Message

	for _, pos := range err.Location {
		str += "\n\t" + pos.String()
	}

	return str
}

func (v Scope) HasVar(name string) bool {
	_, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return false
		}

		return v.ParentScope.HasVar(name)
	}

	return true
}

func (v Scope) GetVar(locs []Location, name string) (*Object, bool, *Error) {
	val, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return nil, false, &Error{
				Location: locs,
				Message:  VAR_NO_EXIST,
			}
		}

		return v.ParentScope.GetVar(locs, name)
	}

	return &val, true, nil
}

func (v Scope) DefVar(locs []Location, name string, val Object) (*Error, bool) {
	_, ok := v.Vars[name]

	if ok {
		return &Error{
			Location: locs,
			Message:  VAR_NO_EXIST,
		}, false
	}

	v.Vars[name] = val

	return nil, true
}

func (v Scope) SetVar(locs []Location, name string, val Object) (*Error, bool) {
	_, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return &Error{
				Location: locs,
				Message:  VAR_NO_EXIST,
			}, false
		}

		v.ParentScope.SetVar(locs, name, val)

		return nil, true
	}

	v.Vars[name] = val

	return nil, true
}

type RootVars struct {
	Log FunctionObject
}

func NewMainScope(rootVars RootVars) Scope {
	return Scope{
		Vars: map[string]Object{
			"log": rootVars.Log,
		},
		Types: map[string]func(any) Object{
			"int":    NewIntegerObject,
			"long":   NewLongObject,
			"bool":   NewBooleanObject,
			"float":  NewFloatObject,
			"double": NewDoubleObject,
			"string": NewStringObject,
		},
	}
}
