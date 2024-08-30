package main

// In Go, you can use interfaces to achieve this polymorphism where a VMObject 
// (which I'll refer to as VMObj for clarity) can hold one of three possible types: int, float64, or string. 
// Here's how you can create such a structure along with the necessary interface for setting and getting the values.

// this essentially is a way to create a union type in Go; you use a `string` as the base type and then
// define a `Type` method that returns the type of the value as a string and a `Value` method that returns
// the string representation of the value. This way, you can store and retrieve values of different types
// using the same interface.

// The problem is, ints and floats are constantly being parsed and formatted as strings, which is inefficient.
// The Go compiler can't optimize this code because it doesn't know the concrete type of the value at compile time.
// This is where generics come in. With generics, you can define a type parameter for the value type,
// and the compiler can generate specialized code for each concrete type you use.

import (
    "fmt"
)

// Define the Interface
// First, define an interface that the types int, float64, and string will satisfy. 
// This interface will be used by the VMObj to store and retrieve values.

// VMValue is an interface that can be implemented by int, float64, or string.
type VMValue interface {
    Type() string  // Return the type of the value as a string
    Value() string // Return the string representation of the value
}

// IntValue wraps an int value.
type IntValue struct {
    val int
}

func (i IntValue) Type() string {
    return "int"
}

func (i IntValue) Value() string {
    return fmt.Sprintf("%d", i.val)
}

// FloatValue wraps a float64 value.
type FloatValue struct {
    val float64
}

func (f FloatValue) Type() string {
    return "float64"
}

func (f FloatValue) Value() string {
    return fmt.Sprintf("%f", f.val)
}

// StringValue wraps a string value.
type StringValue struct {
    val string
}

func (s StringValue) Type() string {
    return "string"
}

func (s StringValue) Value() string {
    return s.val
}


// VMObj holds a value of type VMValue.
type VMObj struct {
    value VMValue
}

// SetValue sets the value of the VMObj.
func (o *VMObj) SetValue(v VMValue) {
    o.value = v
}

// GetValue retrieves the value of the VMObj.
func (o *VMObj) GetValue() VMValue {
    return o.value
}

// GetType returns the type of the stored value.
func (o *VMObj) GetType() string {
    return o.value.Type()
}

func main() {
    // Create a VMObj instance
    var obj VMObj

    // Set an int value
    obj.SetValue(IntValue{val: 42})
    fmt.Printf("Value: %s, Type: %s\n", obj.GetValue().Value(), obj.GetType())

    // Set a float64 value
    obj.SetValue(FloatValue{val: 3.14})
    fmt.Printf("Value: %s, Type: %s\n", obj.GetValue().Value(), obj.GetType())

    // Set a string value
    obj.SetValue(StringValue{val: "Hello, SOM!"})
    fmt.Printf("Value: %s, Type: %s\n", obj.GetValue().Value(), obj.GetType())
}
