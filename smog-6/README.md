## Smog 6

The sixth attempt at smog system, the VM specifically, and also a new format for the bytecode, in some kind of human readable form.
Each phase will be tested as we go.

Need 5 things:
- heap
- global variables; a map of keys to values of objects kept
- literals
- execution stack
- call stack

each vmobject will have fields, clazz, and if a primitive, a value.

- literal (primitive) types
  - strings, symbols, arrays, integers * doubles

