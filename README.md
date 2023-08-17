# smog
Simple Machine (Objects) in Go

_well, it *was* in go, until today when I decided to do the whole thing in Swift 5.7+_ But like many projects, once named, 

[SOM (Simple Object Machine)](http://som-st.github.io) is 20-yr plus project for a simple OOP language and interpreter.

This is a passion project, re-implementing it in [Go](https://go.dev) (golang).

I'd like to make a simple interpreter (bytecode) for the SOM language.

## SOM

I have included the source for the java version of SOM in `som/`

### plan 3

Do it all in Swift. (and then I can also use xcode!)

- implement vm/vmobjects as close to SOM as possible, but in Swift
- implement the interpreter
- convert compiler

### plan 1 and plan 2

See other Readme's in smog-v1 and smog-v2 for TODO lists that I spent much of the early part of the week working on.

The golang attempts. I couldn't get my head around go's generics (yet) and wasn't able to build a hierarchy of code that worked right.
The interpreter *really* wants to use an inherited class mechanism to count on "super class embedding".
Which, of course, Go does not do. It has it's Composition mechanism.

smog-v1:
```
      78     273    2148 array.go
      26      52     432 biginteger.go
     284    1116    8979 class.go
      23      49     385 double.go
     252     991    7028 frame.go
      23      49     364 integer.go
     439    1152   10564 interpreter.go
       9      17     152 invokable.go
     228     893    6801 method.go
     107     338    2660 object.go
      71     219    1905 primitive.go
       1       2      11 shell.go
      23      49     326 string.go
      54     199    1203 symbol.go
      27      69     449 symbol_test.go
      19      42     328 symboltable.go
      25      71     593 symboltable_test.go
      81     156    1439 universe.go
    1770    5737   45767 total
```

smog-v2:

```
       0       2      12 compiler.go
       1       2      13 primitives.go
     861    3033   24379 vm.go
     482    1923   15215 vmobjects.go
      51     142    1005 vmobjects_test.go
    1395    5102   40624 total
```

### plan (0)

- define the first set of SOM words
- simple lexer
- to AST or to a bytecode buffer?

the vm is proving interesting. 

- create simple text format for the bytecodes
- define interpreter for the bytecodes
- simple lexer and parser to turn SOM code into bytecodes