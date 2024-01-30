## Smog 6

The sixth attempt at smog system, the VM specifically, and also a new format for the bytecode, in some kind of human readable form.
Each phase will be tested as we go.

Need 5 things:
- heap
- global variables; a map of keys to values of objects kept
- literals (constant pool)
- execution (operand) stack
- call stack

each vmobject will have fields, clazz, and if a primitive, a value.

- literal (primitive) types
  - strings, symbols, arrays, integers & doubles

### Instructions

## Bytescodes

	0, halt, 1
	1, dup, 1
	2, push_local, 3
	3, push_argument 3
	4, push_field 2
	5, push_block 2
	6, push_constant 2
	7, push_global 2
	8, pop 1
	9, pop_local 3
	10, pop_argument 3
	11, pop_field 2
	12, send 2
	13, super_send 2
	14, return_local 1
	15, return_non_local 1

- HLT - halt the machine
- DUP - duplicate the top of the stack
- POP - pop the top of the stack
- SWP - swap the top two elements of the stack
- LIT i - push a constant value from the constant pool at index i onto the stack


from haskellprague

HALT exits the runtime environment gracefully without any exceptions.
DUP gets the value on top of the stack and pushes a duplicate of it. POP discards the top item of the stack.
PUSH LITERAL i pushes the literal on the given index to stack, transforming it to an object. If the literal is a block, it also captures the current top call frame in the call stack.
PUSH LOCAL i i pushes a variable from the current local scope on top of the stack. The first index is used for identifying the call frame to look in, 0 meaning the current call frame, 1 the captured call frame, 2 is the captured call frame of the current captured call frame and so on. The second index is then the field index in the given call frame.
PUSH FIELD i pushes to stack a field on the given index of the current object context (the object accessed as self).
PUSH GLOBAL i gets a global value on a given index and pushes it as an object to the stack. This means that if the global is an instance of a class, it’s pushed directly, and if it is a class, it pushes the object representing the class. If no global value of a given index is defined, the runtime exits with an exception.
SET LOCAL i i, SET FIELD i and SET GLOBAL i pop the top value from the stack and sets the local variable, current context object or a global to this value respectively.
CALL i calls a method of a given identifier on the object on top of the stack, pops the required amount of arguments and pushes a new call frame on the call stack with the popped valued as locals. It also reserves space for locally scoped variables and defaults their value to the nil value. This instruction expects that the arguments are ordered on the stack from bottom to top, meaning the last argument is popped next after the message target. This calling convention can be seen illustrated in image 2.3. If no method of the given identifier is defined on the receiving object, a runtime exception is raised and the execution of the virtual machine is halted.
SUPER CALL i calls a method with the same calling convention as CALL, but starts the search for the appropriate method in the superclass of the class where the currently executed method is defined.
RETURN exits the currently executed method, popping the top of the call stack. The return value is passed through the stack.
NONLOCAL RETURN is an exit expression executed inside of a block. It gets the currently captured call frame and validates that it is still present on the call stack. If it is, the call stack pops the frames until it is reached. If it is not, an exception is raised and the runtime is exited with error.

```
BLOCK i creates a block object from the block value on the given index and pushes it to the stack. The block captures the current call frame and the current object context.
```

a C++ impl opcodes

• LIT i retrieves a constant value from the constant pool at the index i and pushes it on the stack. The item can be either integer, double or string value.
• GET SLOT i pops a value from the operand stack, assuming it is an object. Then it retrieves a value with index i from the constatns pool, assuming it is a string. It then retrieves the value stored in the slot with the name specified by the string and pushes it onto the stack.
• SET SLOT i pops a value from the stack. This value is then assigned to an instance variable with identifier at index i in the constants pool.
• SEND i n sends a message to an object, which in most cases results in calling a method. A new frame is created on the execution stack, arguments are pushed and the execution jumps to the first instruction of the method.
• GET LOCAL i retrieves a local variable with an index i and pushes it to the top of the stack.
• SET LOCAL i pops a value x from the top of the stack and then assignes the x into a local variable with the index i.
• GET SELF retrieves the callee of executed method. The object is pushed to the top of the stack.
• GET ARG i retrieves the i–th argument of the current message from the stack and pushes it on top.
• BLOCK i creates a code block object. The argument i points to a block value in the constant pool. The block object is instantiated on the heap and pushed to the top of the stack.
• RET is used to return from a method call. The value from the top of the stack is returned. The address to return to is retrieved from the current frame, then the frame is popped and execution jumps to an instruction after the CALL that invoked the method.
• RETNL i - non local return. The value at the top of the current frame is used as the return value. Argument i specifies the type of non local return.