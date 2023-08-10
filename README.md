# smog
Simple Machine (Objects) in Go

[SOM (Simple Object Machine)](http://som-st.github.io) is 20-yr plus project for a simple OOP language and interpreter.

This is a passion project, re-implementing it in [Go](https://go.dev) (golang).

I'd like to make a simple interpreter (bytecode) for the SOM language.

## SOM

I have included the source for the java version of SOM in `som/`

### plan

- define the first set of SOM words
- simple lexer
- to AST or to a bytecode buffer?

the vm is proving interesting. 

- create simple text format for the bytecodes
- define interpreter for the bytecodes
- simple lexer and parser to turn SOM code into bytecodes