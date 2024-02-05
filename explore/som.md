## SOM.md

### Class
```
MyClass = SuperClass (
| field1 field2 field3 |
Class syntax
 aUnaryMethod = ( ... )
<<< argument = ( ... )
oneArgumentMethod: argument = ( ... )
twoArgumentMethod: argument1 otherArgument: argument2 = ( ... )
----
new = ( ... )
new: argument = ( ... )
)
```

### Method body basics

```
method: argument = (
| local1 local2 |
"body of method"
)
```

- Basic keywords and syntax:
  - _self_ factorial: i = ( ... self factorial: (i-1) ... )
  - _super_ testCondition = ( super testCondition && ... )
  - _assignment_ localOrField := 87
  - _sequencing_ doThis . doThat
  - _return_ identity: i = ( ^i )
  - _comments_ "This is a comment"
  - _literals_ 87, #symbol, ’This is a string’
- Useful basic messages:
  - _object instantiation_ MyClass new
  - _operators_ + - < = || && (and so on)

### Blocks and control structures
_blocks I_ [ __body__ ] (lexical scoping, LIFO)
_conditional_ (1<2) ifTrue: [ ... ] ifFalse: [ ... ]
_looping_ [i>0] whileTrue: [ i := i-1 ] (also whileFalse:)
_blocks II_ [ :parameter | ... body ... ] (any number)
*iteration* vector do: [ :element | element println ]
3

### Protocol for Object

_class_ get class of object
_=_ value equals
_==_ primitive equals
_isNil_ is this object nil? (yes, Nil is an object)
_asString_ string conversion
_,_ Vector construction
_value_ evaluate (interesting for blocks)
_print,println_ printing to std out
_error:_ error reporting
_subClassResponsibility_ used for abstract classes
_doesNotUnderstand:arguments:_ error handling; can do interesting things

and more...
