package object

import (
	"strconv"
	"strings"

	serror "github.com/NekoMaru76/rw3c/serror"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Object interface {
	Type() ObjectType
	String() string
	GetProp(key string) *Object
	SetProp(key string, value *Object)
	Raw() string
}

type BaseObject struct {
	Object,
	Props map[string]*Object
	Text string
}

func (v BaseObject) Raw() string {
	return v.Text
}

func (v BaseObject) GetPropType(key string) ObjectType {
	return AnyType{}
}

func (v BaseObject) GetProp(key string) *Object {
	p, ok := v.Props[key]

	if !ok {
		null := NewNullObject()

		return &null
	}

	return p
}

func (v BaseObject) SetProp(key string, value *Object) {
	v.Props[key] = value
}

func (v BaseObject) TestProp(key string) *serror.Error {
	return nil
}

type UnwrappedObject struct {
	Test        func() *serror.Error
	Value       func() (*Object, *serror.Error)
	Type        func() ObjectType
	Raw         func() string
	TestProp    func(key string) *serror.Error
	GetPropType func(key string) ObjectType
}

type WrappedObject = func(locs []serror.Location, scope Scope) UnwrappedObject

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
	return StringType{}
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
	return NullType{}
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
	return BooleanType{}
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
	return IntegerType{}
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
	return LongType{}
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
	return FloatType{}
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
	return DoubleType{}
}

func (v DoubleObject) String() string {
	return strconv.FormatFloat(v.Value, 8, 64, 64)
}

type FunctionObject struct {
	BaseObject
	Value func(...Object) (*Object, *serror.Error)
	Name  *string
}

func (v FunctionObject) Type() ObjectType {
	return FunctionType{}
}

func (v FunctionObject) String() string {
	return v.Raw()
}

type Variable struct {
	Value Object
	Type  ObjectType
}

type Scope struct {
	Vars        map[string]Variable
	Types       map[string]func(any) Object
	ParentScope *Scope
	Name        string
}

func TokToLoc(tok antlr.Token, file string) serror.Location {
	return serror.Location{
		Column: tok.GetColumn(),
		Line:   tok.GetLine(),
		File:   file,
	}
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

func (v Scope) GetVarType(name string) *ObjectType {
	val, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return nil
		}

		return v.ParentScope.GetVarType(name)
	}

	return &val.Type
}

func (v Scope) GetVar(locs []serror.Location, name string) (*Object, *serror.Error) {
	val, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return nil, &serror.Error{
				Location: locs,
				Message:  serror.VAR_NO_EXIST,
			}
		}

		return v.ParentScope.GetVar(locs, name)
	}

	return &val.Value, nil
}

func (v Scope) DefVar(locs []serror.Location, name string, variable Variable) *serror.Error {
	_, ok := v.Vars[name]

	if ok {
		return &serror.Error{
			Location: locs,
			Message:  serror.VAR_NO_EXIST,
		}
	}

	v.Vars[name] = variable

	return nil
}

func (v Scope) SetVar(locs []serror.Location, name string, val Object) *serror.Error {
	u, ok := v.Vars[name]

	if !ok {
		if v.ParentScope == nil {
			return &serror.Error{
				Location: locs,
				Message:  serror.VAR_NO_EXIST,
			}
		}

		v.ParentScope.SetVar(locs, name, val)

		return nil
	}

	u.Value = val

	return nil
}

type RootVars struct {
	Log FunctionObject
}

func NewMainScope(rootVars RootVars) Scope {

	return Scope{
		Vars: map[string]Variable{
			"log": {
				Type:  FunctionType{},
				Value: rootVars.Log,
			},
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

const (
	STRING_TYPE   = "STRING"
	INTEGER_TYPE  = "INTEGER"
	LONG_TYPE     = "LONG"
	FLOAT_TYPE    = "FLOAT"
	DOUBLE_TYPE   = "DOUBLE"
	NULL_TYPE     = "NULL"
	BOOLEAN_TYPE  = "BOOLEAN"
	FUNCTION_TYPE = "FUNCTION"
	UNION_TYPE    = "UNION"
	STRUCT_TYPE   = "STRUCT"
	ANY_TYPE      = "ANY"
)

type ObjectType interface {
	String() string
}

type StringType struct {
	ObjectType
}

func (v StringType) String() string {
	return STRING_TYPE
}

type NumberType interface {
	IsFloating() bool
	Bit() int
}

type IntegerType struct {
	NumberType
}

func (v *IntegerType) IsFloating() bool {
	return false
}

func (v *IntegerType) Bit() int {
	return 32
}

func (v IntegerType) String() string {
	return INTEGER_TYPE
}

type LongType struct {
	NumberType
}

func (v *LongType) IsFloating() bool {
	return false
}

func (v *LongType) Bit() int {
	return 64
}

func (v LongType) String() string {
	return LONG_TYPE
}

type DoubleType struct {
	NumberType
}

func (v *DoubleType) IsFloating() bool {
	return true
}

func (v *DoubleType) Bit() int {
	return 32
}

func (v DoubleType) String() string {
	return DOUBLE_TYPE
}

type FloatType struct {
	NumberType
}

func (v *FloatType) IsFloating() bool {
	return true
}

func (v *FloatType) Bit() int {
	return 64
}

func (v FloatType) String() string {
	return FLOAT_TYPE
}

type NullType struct {
	ObjectType
}

func (v NullType) String() string {
	return NULL_TYPE
}

type BooleanType struct {
	ObjectType
}

func (v BooleanType) String() string {
	return BOOLEAN_TYPE
}

type AnyType struct {
	ObjectType
}

func (v AnyType) String() string {
	return ANY_TYPE
}

type FunctionType struct {
	ObjectType
	Args       []ObjectType
	ReturnType func() ObjectType
}

func (v FunctionType) TestCall(locs []serror.Location, args []ObjectType) *serror.Error {
	for i, arg := range v.Args {
		argV := args[i]

		arg_T := arg.String()
		argV_T := argV.String()

		if argV_T == arg_T {
			switch arg_T {
			case UNION_TYPE:
				panic("")
			case STRUCT_TYPE:
				panic("")
			}
		} else {
			return &serror.Error{
				Message:  serror.NO_CALL_FN_ARG_DIFF(arg_T, "[placeholder]", i, argV_T),
				Location: locs,
			}
		}
	}

	return nil
}

func (v FunctionType) String() string {
	return FUNCTION_TYPE
}

type StructType struct {
	ObjectType
}

func (v StructType) String() string {
	return STRUCT_TYPE
}

type t2 = map[string]map[string]func(Object, Object) Object

func newNum(isFloating bool, bit int) func(any) Object {
	switch bit {
	case 32:
		{
			switch isFloating {
			case true:
				return NewFloatObject
			case false:
				return NewIntegerObject
			}
		}
	case 64:
		{
			switch isFloating {
			case true:
				return NewDoubleObject
			case false:
				return NewLongObject
			}
		}
	}

	panic("Invalid bit and isFloating")
}

func noOverload2Err(locs []serror.Location, a, b, op string) *serror.Error {
	return &serror.Error{
		Location: locs,
		Message:  serror.NO_OVERLOAD_2(a, b, op),
	}
}

func TestSum(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	switch a.Type().(type) {
	case StringObject:
		{
			return nil
		}
	case FloatObject:
	case IntegerObject:
	case LongObject:
	case DoubleObject:
		{
			switch b.Type().(type) {
			case FloatObject:
			case IntegerObject:
			case LongObject:
			case DoubleObject:
				return nil
			default:
				return noOverload2Err(locs, a.Raw(), b.Raw(), "+")
			}
		}
	}

	return noOverload2Err(locs, a.Raw(), b.Raw(), "+")
}

func Sum(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	switch o1 := a.(type) {
	case StringObject:
		{
			str := NewStringObject(o1.Value + b.String())
			return &str, nil
		}
	case FloatObject:
	case IntegerObject:
	case LongObject:
	case DoubleObject:
		{
			switch o2 := b.(type) {
			case FloatObject:
			case IntegerObject:
			case LongObject:
			case DoubleObject:
				{
					o1t := o1.Type().(NumberType)
					o2t := o2.Type().(NumberType)
					isFloating := o1t.IsFloating() || o2t.IsFloating()
					bit := max(o1t.Bit(), o2t.Bit())
					val := newNum(isFloating, bit)(o1.Value + o2.Value)

					return &val, nil
				}
			}
		}
	}

	return nil, noOverload2Err(locs, a.Raw(), b.Raw(), "+")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
