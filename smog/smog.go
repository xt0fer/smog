package smog

import (
	"errors"
	"math"
	"reflect"
)

const (
	INTEGER_OBJ   = "Integer"
	FLOAT_OBJ     = "Float"
	BOOLEAN_OBJ   = "Boolean"
	STRING_OBJ    = "String"
	SYMBOL_OBJ    = "Symbol"
	BLOCK_OBJ     = "BLock"
	DEFERRED      = "Deferred"
	ARRAY_OBJ     = "Array"
	CLASS_OBJ     = "Class"
	OBJECT_OBJ    = "Object"
	UNDEFINED_OBJ = "Undefined"
)

var floatMessages = map[string]interface{}{
	`value`:            value,
	`=`:                equal,
	`~=`:               notEqual,
	`>`:                greater,
	`>=`:               greaterEqual,
	`<`:                lesser,
	`<=`:               lesserEqual,
	`+`:                plus,
	`-`:                minus,
	`*`:                mul,
	`/`:                div,
	`\\`:               mod,
	`//`:               intDiv,
	`rem:`:             rem,
	`max:`:             max,
	`min:`:             min,
	`abs`:              abs,
	`sqrt`:             sqrt,
	`sqr`:              sqr,
	`sin`:              sin,
	`cos`:              cos,
	`tan`:              tan,
	`arcSin`:           arcSin,
	`arcCos`:           arcCos,
	`arcTan`:           arcTan,
	`rounded`:          rounded,
	`truncated`:        truncated,
	`fractionPart`:     fractionPart,
	`floor`:            floor,
	`ceiling`:          ceiling,
	`negated`:          negated,
	`degreesToRadians`: degreesToRadians,
}

var booleanMessages = map[string]interface{}{
	`value`:           value,
	`=`:               boolEqual,
	`~=`:              boolNotEqual,
	`ifTrue:`:         ifTrue,
	`ifFalse:`:        ifFalse,
	`ifTrue:ifFalse:`: ifTrueIfFalse,
	`ifFalse:ifTrue:`: ifFalseIfTrue,
	`and:`:            and,
	`&`:               ampersand,
	`or:`:             or,
	`|`:               verticalBar,
	`xor:`:            xor,
	`not`:             not,
}

var blockMessages = map[string]interface{}{
	`value`:  value,
	`value:`: valueWith,
}

var arrayMessages = map[string]interface{}{
	`at:`: ValueAt,
	`+`:   arrPlus,
	`-`:   arrMinus,
	`*`:   arrMul,
	`/`:   arrDiv,
	`\\`:  arrMod,
	`//`:  arrIntDiv,
}

func value(receiver SmogObjectInterface) SmogObjectInterface {
	return receiver.Value()
}

func valueWith(receiver *SmogBlock, arg SmogObjectInterface) SmogObjectInterface {
	// scope := new(Scope).Initialize()
	// scope.OuterScope = receiver.scope
	// scope.SetVar(receiver.block.arguments[0].GetName(), arg)
	// return receiver.block.body.Eval(scope)
	// TODO: implement this
	return NewSmogUndefinedObject()
}

func equal(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() == arg.GetValue())
}

func notEqual(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() != arg.GetValue())
}

func greater(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() > arg.GetValue())
}

func greaterEqual(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() >= arg.GetValue())
}

func lesser(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() < arg.GetValue())
}

func lesserEqual(receiver *SmogFloat, arg *SmogFloat) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() <= arg.GetValue())
}

func plus(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.GetValue() + arg.GetValue())
}

func minus(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.GetValue() - arg.GetValue())
}

func mul(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.GetValue() * arg.GetValue())
}

func div(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.GetValue() / arg.GetValue())
}

func mod(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(float64(int64(receiver.GetValue()) % int64(arg.GetValue())))
}

func intDiv(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Floor(receiver.GetValue() / arg.GetValue()))
}

func rem(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	quo := math.Trunc(receiver.value / arg.value)
	return new(SmogFloat).SetValue(receiver.value - (quo * arg.value))
}

func max(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	if receiver.value > arg.value {
		return receiver
	} else {
		return arg
	}
}

func min(receiver *SmogFloat, arg *SmogFloat) *SmogFloat {
	if receiver.value > arg.value {
		return arg
	} else {
		return receiver
	}
}

func abs(receiver *SmogFloat) *SmogFloat {
	if receiver.value < 0 {
		return new(SmogFloat).SetValue(receiver.value * -1)
	} else {
		return receiver
	}
}

func sqrt(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Sqrt(receiver.value))
}

func sqr(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Pow(receiver.value, 2))
}

func sin(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Sin(receiver.value))
}

func cos(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Cos(receiver.value))
}

func tan(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Tan(receiver.value))
}

func arcSin(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Asin(receiver.value))
}

func arcCos(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Acos(receiver.value))
}

func arcTan(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Atan(receiver.value))
}

func rounded(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Round(receiver.value))
}

func truncated(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Trunc(receiver.value))
}

func floor(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Floor(receiver.value))
}

func ceiling(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(math.Ceil(receiver.value))
}

func fractionPart(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.value - math.Trunc(receiver.value))
}

func negated(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.value * -1)
}

func degreesToRadians(receiver *SmogFloat) *SmogFloat {
	return new(SmogFloat).SetValue(receiver.value * math.Pi / 180.0)
}

// Boolean receiver messages section
func boolEqual(receiver *SmogBoolean, arg *SmogBoolean) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() == arg.GetValue())
}

func boolNotEqual(receiver *SmogBoolean, arg *SmogBoolean) *SmogBoolean {
	return new(SmogBoolean).SetValue(receiver.GetValue() != arg.GetValue())
}

func and(receiver *SmogBoolean, arg *SmogBlock) *SmogBoolean {
	if receiver.GetValue() {
		return arg.Value().(*SmogBoolean)
	} else {
		return receiver
	}
}

func ampersand(receiver *SmogBoolean, arg *SmogBoolean) *SmogBoolean {
	if receiver.GetValue() {
		return arg
	} else {
		return receiver
	}
}

func or(receiver *SmogBoolean, arg *SmogBlock) *SmogBoolean {
	if receiver.GetValue() {
		return receiver
	} else {
		return arg.Value().(*SmogBoolean)
	}
}

func verticalBar(receiver *SmogBoolean, arg *SmogBoolean) *SmogBoolean {
	if receiver.GetValue() {
		return receiver
	} else {
		return arg
	}
}

func xor(receiver *SmogBoolean, arg *SmogBoolean) *SmogBoolean {
	xor := !(receiver.GetValue() == arg.GetValue())
	return new(SmogBoolean).SetValue(xor)
}

func not(receiver *SmogBoolean) *SmogBoolean {
	if receiver.GetValue() {
		return new(SmogBoolean).SetValue(false)
	} else {
		return new(SmogBoolean).SetValue(true)
	}
}

func ifTrue(receiver *SmogBoolean, arg SmogObjectInterface) SmogObjectInterface {
	if receiver.GetValue() {
		return arg.Value()
	} else {
		return NewSmogUndefinedObject()
	}
}

func ifFalse(receiver *SmogBoolean, arg SmogObjectInterface) SmogObjectInterface {
	if receiver.GetValue() {
		return NewSmogUndefinedObject()
	} else {
		return arg.Value()
	}
}

func ifTrueIfFalse(receiver *SmogBoolean, argTrue SmogObjectInterface, argFalse SmogObjectInterface) SmogObjectInterface {
	if receiver.GetValue() {
		return argTrue.Value()
	} else {
		return argFalse.Value()
	}
}

func ifFalseIfTrue(receiver *SmogBoolean, argFalse SmogObjectInterface, argTrue SmogObjectInterface) SmogObjectInterface {
	if receiver.GetValue() {
		return argTrue.Value()
	} else {
		return argFalse.Value()
	}
}

// Array methods
func ValueAt(receiver *SmogArray, index *SmogFloat) SmogObjectInterface {
	return receiver.array[int64(index.value)-1]
}

func arrPlus(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, plus(each.(*SmogFloat), number))
	}

	return result
}

func arrMinus(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, minus(each.(*SmogFloat), number))
	}

	return result
}

func arrMul(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, mul(each.(*SmogFloat), number))
	}

	return result
}

func arrDiv(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, div(each.(*SmogFloat), number))
	}

	return result
}

func arrMod(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, mod(each.(*SmogFloat), number))
	}

	return result
}

func arrIntDiv(receiver *SmogArray, number *SmogFloat) SmogObjectInterface {
	result := new(SmogArray)
	for _, each := range receiver.array {
		result.array = append(result.array, intDiv(each.(*SmogFloat), number))
	}

	return result
}

/*******************************
 *
 * SmogObjectInterface methods
 *
 *******************************/
//////////////////////
//////////////////////
//////////////////////
//////////////////////
//////////////////////
///////////////////////////////////////

func Call(receiver SmogObjectInterface, m map[string]interface{}, name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	var receiverAndArgs []SmogObjectInterface
	if receiver.TypeOf() == DEFERRED {
		receiverAndArgs = append(receiverAndArgs, receiver.Value())
	} else {
		receiverAndArgs = append(receiverAndArgs, receiver)
	}
	for _, each := range params {
		if each.TypeOf() == DEFERRED {
			receiverAndArgs = append(receiverAndArgs, each.Value())
		} else {
			receiverAndArgs = append(receiverAndArgs, each)
		}
	}
	f, ok := m[name]
	if !ok {
		err := errors.New("does not understand: " + name)
		return nil, err
	}
	function := reflect.ValueOf(f)
	if len(receiverAndArgs) != function.Type().NumIn() {
		err := errors.New("wrong parameters length")
		return nil, err
	}
	in := make([]reflect.Value, len(receiverAndArgs))
	for k, param := range receiverAndArgs {
		in[k] = reflect.ValueOf(param)
	}
	result := function.Call(in)
	return result[0].Interface().(SmogObjectInterface), nil
}

type SmogObjectInterface interface {
	TypeOf() string
	Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error)
	Value() SmogObjectInterface
}

type SmogObject struct {
}

func (obj *SmogObject) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return nil, nil
}

type SmogUndefinedObject struct {
	*SmogObject
}

func NewSmogUndefinedObject() *SmogUndefinedObject {
	return &SmogUndefinedObject{&SmogObject{}}
}

func (n *SmogUndefinedObject) Value() SmogObjectInterface {
	return n
}

func (n *SmogUndefinedObject) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return nil, errors.New("doesNotUnderstand")
}

func (n *SmogUndefinedObject) TypeOf() string {
	return UNDEFINED_OBJ
}

type SmogFloat struct {
	*SmogObject
	value float64
}

func NewSmogFloat(value float64) *SmogFloat {
	return &SmogFloat{&SmogObject{}, value}
}

func (n *SmogFloat) Value() SmogObjectInterface {
	return n
}

func (n *SmogFloat) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return Call(n, floatMessages, name, params)
}

func (n *SmogFloat) TypeOf() string {
	return FLOAT_OBJ
}

func (n *SmogFloat) GetValue() float64 {
	return n.value
}

func (n *SmogFloat) SetValue(val float64) *SmogFloat {
	n.value = val
	return n
}

type SmogSymbol struct {
	*SmogObject
	value string
}

func NewSmogSymbol(value string) *SmogSymbol {
	return &SmogSymbol{&SmogObject{}, value}
}

func (s *SmogSymbol) Value() SmogObjectInterface {
	return s
}

func (s *SmogSymbol) TypeOf() string {
	return SYMBOL_OBJ
}

func (s *SmogSymbol) GetValue() string {
	return s.value
}

func (s *SmogSymbol) SetValue(val string) *SmogSymbol {
	s.value = val
	return s
}

type SmogBoolean struct {
	*SmogObject
	value bool
}

func NewSmogBoolean(value bool) *SmogBoolean {
	return &SmogBoolean{&SmogObject{}, value}
}

func (b *SmogBoolean) Value() SmogObjectInterface {
	return b
}

func (b *SmogBoolean) TypeOf() string {
	return BOOLEAN_OBJ
}

func (b *SmogBoolean) GetValue() bool {
	return b.value
}

func (b *SmogBoolean) SetValue(val bool) *SmogBoolean {
	b.value = val
	return b
}

func (b *SmogBoolean) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return Call(b, booleanMessages, name, params)
}

type BlockNode struct {
}
type Scope struct {
}

type SmogBlock struct {
	*SmogObject
	block *BlockNode
	scope *Scope
}

func (b *SmogBlock) Value() SmogObjectInterface {
	return NewSmogUndefinedObject() //b.block.body.Eval(b.scope)
}

func (b *SmogBlock) TypeOf() string {
	return BLOCK_OBJ
}

func (b *SmogBlock) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return Call(b, blockMessages, name, params)
}

type SmogArray struct {
	SmogObject
	array []SmogObjectInterface
}

func (a *SmogArray) GetValueAt(index int64) SmogObjectInterface {
	return a.array[index]
}

func (a *SmogArray) GetValue() ([]interface{}, error) {
	var interfaceSlice = make([]interface{}, len(a.array))
	for i, each := range a.array {
		switch each.TypeOf() {
		case NUMBER_OBJ:
			interfaceSlice[i] = each.(*SmogFloat).GetValue()
		case STRING_OBJ:
			interfaceSlice[i] = each.(*SmogString).GetValue()
		case BOOLEAN_OBJ:
			interfaceSlice[i] = each.(*SmogBoolean).GetValue()
		case ARRAY_OBJ:
			innerArray, err := each.(*SmogArray).GetValue()
			if err != nil {
				return nil, err
			}
			interfaceSlice[i] = innerArray
		default:
			return nil, errors.New(`we do not support this type "` + each.TypeOf() + `" in array`)
		}
	}
	return interfaceSlice, nil
}

func (a *SmogArray) Value() SmogObjectInterface {
	return a
}

func (a *SmogArray) TypeOf() string {
	return ARRAY_OBJ
}

func (a *SmogArray) Send(name string, params []SmogObjectInterface) (SmogObjectInterface, error) {
	return Call(a, arrayMessages, name, params)
}

type Deferred struct {
	*SmogBlock
}

func (d *Deferred) TypeOf() string {
	return DEFERRED
}

func NewDeferred(blockNode *BlockNode, scope *Scope) *Deferred {
	return &Deferred{&SmogBlock{&SmogObject{}, blockNode, scope}}
}
