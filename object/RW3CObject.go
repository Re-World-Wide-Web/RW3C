package object

import (
	"reflect"
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

type NumberObject interface {
	Type() ObjectType
	String() string
	IsFloating() bool
	Bit() int32
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
	return strconv.FormatFloat(float64(v.Value), 8, 64, 64)
}

func (v FloatObject) IsFloating() bool {
	return true
}

func (v FloatObject) Bit() int32 {
	return 32
}

func NewIntegerObject(Value any) Object {
	V, ok := Value.(int32)

	if !ok {
		V = 0
	}

	return IntegerObject{
		Value: V,
	}
}

type IntegerObject struct {
	BaseObject
	Value int32
}

func (v IntegerObject) Type() ObjectType {
	return IntegerType{}
}

func (v IntegerObject) String() string {
	return strconv.FormatInt(int64(v.Value), 10)
}

func (v IntegerObject) IsFloating() bool {
	return false
}

func (v IntegerObject) Bit() int32 {
	return 32
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

func (v LongObject) IsFloating() bool {
	return false
}

func (v LongObject) Bit() int32 {
	return 64
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

func (v DoubleObject) IsFloating() bool {
	return true
}

func (v DoubleObject) Bit() int32 {
	return 64
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
	Bit() int32
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

func newNum(isFloating bool, bit int32, v any) Object {
	switch bit {
	case 32:
		{
			switch isFloating {
			case true:
				return NewFloatObject(v)
			case false:
				return NewIntegerObject(v)
			}
		}
	case 64:
		{
			switch isFloating {
			case true:
				return NewDoubleObject(v)
			case false:
				return NewLongObject(v)
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

type BaseNumberObject struct {
	NumberObject
}

var operatorTypes = map[string][][]reflect.Type{
	"+": {
		{reflect.TypeOf(StringObject{}), reflect.TypeOf(BaseObject{})},
		{reflect.TypeOf(BaseNumberObject{}), reflect.TypeOf(BaseNumberObject{})},
	},
	"-": {
		{reflect.TypeOf(BaseNumberObject{}), reflect.TypeOf(BaseNumberObject{})},
	},
	"*": {
		{reflect.TypeOf(BaseNumberObject{}), reflect.TypeOf(BaseNumberObject{})},
	},
	"/": {
		{reflect.TypeOf(BaseNumberObject{}), reflect.TypeOf(BaseNumberObject{})},
	},
}

// checkTypes checks if the types of the operands are supported for the given operator.
// If the types are not supported, it returns an error.
func checkTypes(locs []serror.Location, a, b UnwrappedObject, op string) *serror.Error {
	aType, bType := reflect.TypeOf(a.Type()), reflect.TypeOf(b.Type())
	supportedTypesList, ok := operatorTypes[op]
	if !ok {
		return &serror.Error{
			Location: locs,
			Message:  serror.NO_OVERLOAD_2(a.Type().String(), b.Type().String(), op),
		}
	}

	for _, supportedTypes := range supportedTypesList {
		if containsType(supportedTypes, aType) && containsType(supportedTypes, bType) {
			return nil
		}
	}

	return noOverload2Err(locs, a.Raw(), b.Raw(), op)
}

// containsType checks if the given list of types contains the given type.
func containsType(types []reflect.Type, t reflect.Type) bool {
	for _, typ := range types {
		if typ == t {
			return true
		}
	}
	return false
}

func TestSum(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	return checkTypes(locs, a, b, "+")
}

func TestSub(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	return checkTypes(locs, a, b, "-")
}

func TestMul(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	return checkTypes(locs, a, b, "*")
}

func TestDiv(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	return checkTypes(locs, a, b, "/")
}

func TestMod(locs []serror.Location, a UnwrappedObject, b UnwrappedObject) *serror.Error {
	return checkTypes(locs, a, b, "%")
}

func Sum(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	return performArithmeticOperation(locs, a, b, "+")
}

func Sub(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	return performArithmeticOperation(locs, a, b, "-")
}

func Mul(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	return performArithmeticOperation(locs, a, b, "*")
}

func Div(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	return performArithmeticOperation(locs, a, b, "/")
}

func Mod(locs []serror.Location, a Object, b Object) (*Object, *serror.Error) {
	return performArithmeticOperation(locs, a, b, "%")
}

func performArithmeticOperation(locs []serror.Location, a Object, b Object, op string) (*Object, *serror.Error) {
	switch o1 := a.(type) {
	case StringObject:
		switch o2 := b.(type) {
		case StringObject:
			{
				obj := NewStringObject(o1.Value + o2.Value)

				return &obj, nil
			}
		default:
			{
				obj := NewStringObject(o1.Value + o2.String())

				return &obj, nil
			}
		}
	case NumberObject:
		switch o2 := b.(type) {
		case NumberObject:
			{
				isFloating := o1.IsFloating() || o2.IsFloating()
				bit := max(o1.Bit(), o2.Bit())

				var value1 float64
				switch v1 := o1.(type) {
				case FloatObject:
					value1 = float64(v1.Value)
				case IntegerObject:
					value1 = float64(v1.Value)
				case LongObject:
					value1 = float64(v1.Value)
				case DoubleObject:
					value1 = v1.Value
				}
				var value2 float64
				switch v2 := o2.(type) {
				case FloatObject:
					value2 = float64(v2.Value)
				case IntegerObject:
					value2 = float64(v2.Value)
				case LongObject:
					value2 = float64(v2.Value)
				case DoubleObject:
					value2 = v2.Value
				}

				var value float64

				switch op {
				case "+":
					value = value1 + value2
				case "-":
					value = value1 - value2
				case "*":
					value = value1 * value2
				case "/":
					value = value1 / value2
				}

				var obj = newNum(isFloating, bit, value)

				return &obj, nil
			}
		}
	}

	return nil, noOverload2Err(locs, a.Raw(), b.Raw(), op)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
