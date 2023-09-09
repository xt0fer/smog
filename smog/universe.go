package smog

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Universe struct {
	Globals       map[*Symbol]interface{}
	symboltable   *SymbolTable
	dumpBytecodes bool
	//
	PathSep string
	FileSep string

	classPath []string
	//
	NilObject   *Object
	TrueObject  *Object
	FalseObject *Object

	ObjectClass    *Class
	ClassClass     *Class
	MetaclassClass *Class

	NilClass        *Class
	IntegerClass    *Class
	BigIntegerClass *Class
	ArrayClass      *Class
	MethodClass     *Class
	SymbolClass     *Class
	FrameClass      *Class
	PrimitiveClass  *Class
	StringClass     *Class
	BlockClass      *Class
	DoubleClass     *Class
}

var instantiated *Universe
var interpreter *Interpreter

var once sync.Once

func GetUniverse() *Universe {
	once.Do(func() {
		instantiated = &Universe{}
		instantiated.initUniverse()
		interpreter = &Interpreter{}
	})
	return instantiated
}

// GetInterpreter returns the singleton interpreter
func GetInterpreter() *Interpreter {
	once.Do(func() {
		instantiated = &Universe{}
		instantiated.initUniverse()
		interpreter = &Interpreter{}
	})
	return interpreter

}

func (u *Universe) initUniverse() {
	u.symboltable = NewSymbolTable()
}

// func (u *Universe) NewArray(size int) *Array {
// 	return NewArray(size)
// }

func (u *Universe) SymbolFor(sym string) *Symbol {
	result := u.symboltable.lookup(sym)
	if result != nil {
		return result
	}
	result = u.NewSymbol(sym)
	return result
}

func (u *Universe) NewSymbol(sym string) *Symbol {
	result := NewSymbol(sym)
	result.setClass(u.SymbolClass)
	u.symboltable.insert(result)
	return result
}

// public static void main(java.lang.String[] arguments)
// {
//   // Setup the system path and file separators
//   pathSeparator = System.getProperty("path.separator");
//   fileSeparator = System.getProperty("file.separator");

//   // Check for command line switches
//   arguments = handleArguments(arguments);

//   // Initialize the known universe
//   Universe.initialize(arguments);

//	  // Exit with error code 0
//	  exit(0);
//	}
func (u *Universe) Initialize(arguments []string) {
	// Setup the system path and file separators
	u.PathSep = "/"
	u.FileSep = ":"

	// Check for command line switches
	arguments = u.handleArguments(arguments)

	// Initialize the known universe
	//u.initialize(arguments)
}

// public static void exit(int errorCode)
//
//	{
//	  // Exit from the Java system
//	  System.exit(errorCode);
//	}
func (u *Universe) Exit(errorCode int) {
	// Exit from the Java system
	os.Exit(errorCode)
}

func (u *Universe) ErrorExit(message string) {
	fmt.Println("Runtime Error: " + message)
	u.Exit(1)
}

func (u *Universe) handleArguments(arguments []string) []string {
	gotClasspath := false
	remainingArgs := make([]string, len(arguments))
	cnt := 0
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "-cp" {
			if i+1 >= len(arguments) {
				u.PrintUsageAndExit()
			}
			u.setupClassPath(arguments[i+1])
			i++ // skip class path
			gotClasspath = true
		} else if arguments[i] == "-d" {
			u.dumpBytecodes = true
		} else {
			remainingArgs[cnt] = arguments[i]
			cnt++
		}
	}
	if !gotClasspath {
		u.classPath = u.SetupDefaultClassPath(0)
	}
	// Copy the remaining elements from the original array into the new array
	arguments = make([]string, cnt)
	for i := 0; i < cnt; i++ {
		arguments[i] = remainingArgs[i]
	}
	// check remaining args for class paths, and strip file extension
	for i := 0; i < len(arguments); i++ {
		split := u.getPathClassExt(arguments[i])
		if split[0] != "" { // there was a path
			tmp := make([]string, len(u.classPath)+1)
			for j := 0; j < len(u.classPath); j++ {
				tmp[j+1] = u.classPath[j]
			}
			tmp[0] = split[0]
			u.classPath = tmp
		}
		arguments[i] = split[1]
	}
	return arguments
}

// // take argument of the form "../foo/Test.som" and return
// // "../foo", "Test", "som"
func (u *Universe) getPathClassExt(arg string) []string {
	// Create a new tokenizer to split up the string of dirs
	tokenizer := strings.Split(arg, u.FileSep)
	cp := ""
	for i := 0; i < len(tokenizer)-2; i++ {
		cp = cp + tokenizer[i]
	}
	file := tokenizer[len(tokenizer)-1]
	tokenizer = strings.Split(file, ".")
	if len(tokenizer) > 2 {
		fmt.Println("Class with . in its name?")
		u.Exit(1)
	}
	result := make([]string, 3)
	result[0] = cp
	result[1] = tokenizer[0]
	if len(tokenizer) > 1 {
		result[2] = tokenizer[1]
	} else {
		result[2] = ""
	}
	return result
}

// private static void setupClassPath(java.lang.String cp)
//
//	{
//	  // Create a new tokenizer to split up the string of directories
//	  java.util.StringTokenizer tokenizer
//		= new java.util.StringTokenizer(cp, pathSeparator);
//	  // Get the default class path of the appropriate size
//	  classPath = setupDefaultClassPath(tokenizer.countTokens());
//		  // Get the dictories and put them into the class path array
//		  for (int i = 0; tokenizer.hasMoreTokens(); i++) {
//			classPath[i] = tokenizer.nextToken();
//		  }
//		}
func (u *Universe) setupClassPath(cp string) {
	// Create a new tokenizer to split up the string of directories
	tokenizer := strings.Split(cp, u.PathSep)
	// Get the default class path of the appropriate size
	u.classPath = u.SetupDefaultClassPath(len(tokenizer))
	// Get the dictories and put them into the class path array
	for i := 0; i < len(tokenizer); i++ {
		u.classPath[i] = tokenizer[i]
	}
}

// private static java.lang.String[] setupDefaultClassPath(int directories)
// {
//   // Get the default system class path
//   java.lang.String systemClassPath = System.getProperty("system.class.path");

//   // Compute the number of defaults
//   int defaults = (systemClassPath != null) ? 2 : 1;

//   // Allocate an array with room for the directories and the defaults
//   java.lang.String[] result = new java.lang.String[directories + defaults];

//   // Insert the system class path into the defaults section
//   if (systemClassPath != null) {
// 	result[directories] = systemClassPath;
//   }

//   // Insert the current directory into the defaults section
//   result[directories + defaults - 1] = ".";

//   // Return the class path
//   return result;
// }

func (u *Universe) SetupDefaultClassPath(directories int) []string {
	// Get the default system class path
	systemClassPath := ""

	// Compute the number of defaults
	defaults := 1
	if systemClassPath != "" {
		defaults++
	}

	// Allocate an array with room for the directories and the defaults
	result := make([]string, directories+defaults)

	// Insert the system class path into the defaults section
	if systemClassPath != "" {
		result[directories] = systemClassPath
	}

	// Insert the current directory into the defaults section
	result[directories+defaults-1] = "."

	// Return the class path
	return result
}

// private static void printUsageAndExit()
// {
//   // Print the usage
//   System.out.println("Usage: som [-options] [args...]                          ");
//   System.out.println("                                                         ");
//   System.out.println("where options include:                                   ");
//   System.out.println("    -cp <directories separated by " + pathSeparator + ">");
//   System.out.println("                  set search path for application classes");
//   System.out.println("    -d            enable disassembling");

//	  // Exit
//	  System.exit(0);
//	}
func (u *Universe) PrintUsageAndExit() {
	// Print the usage
	fmt.Println("Usage: som [-options] [args...]                          ")
	fmt.Println("                                                         ")
	fmt.Println("where options include:                                   ")
	fmt.Println("    -cp <directories separated by " + u.PathSep + ">")
	fmt.Println("                  set search path for application classes")
	fmt.Println("    -d            enable disassembling")

	os.Exit(0)
}

// private static void initialize(java.lang.String[] arguments)
// {
//   // Allocate the nil object
//   nilObject = new Object();

//   // Allocate the Metaclass classes
//   metaclassClass = newMetaclassClass();

//   // Allocate the rest of the system classes
//   objectClass    = newSystemClass();
//   nilClass       = newSystemClass();
//   classClass     = newSystemClass();
//   arrayClass     = newSystemClass();
//   symbolClass    = newSystemClass();
//   methodClass    = newSystemClass();
//   integerClass   = newSystemClass();
//   bigintegerClass= newSystemClass();
//   frameClass     = newSystemClass();
//   primitiveClass = newSystemClass();
//   stringClass    = newSystemClass();
//   doubleClass    = newSystemClass();

//   // Setup the class reference for the nil object
//   nilObject.setClass(nilClass);

//   // Initialize the system classes.
//   initializeSystemClass(objectClass    , null         , "Object"    );
//   initializeSystemClass(classClass     , objectClass , "Class"     );
//   initializeSystemClass(metaclassClass , classClass  , "Metaclass" );
//   initializeSystemClass(nilClass       , objectClass , "Nil"       );
//   initializeSystemClass(arrayClass     , objectClass , "Array"     );
//   initializeSystemClass(methodClass    , arrayClass  , "Method"    );
//   initializeSystemClass(symbolClass    , objectClass , "Symbol"    );
//   initializeSystemClass(integerClass   , objectClass , "Integer"   );
//   initializeSystemClass(bigintegerClass, objectClass , "BigInteger");
//   initializeSystemClass(frameClass     , arrayClass  , "Frame"     );
//   initializeSystemClass(primitiveClass , objectClass , "Primitive" );
//   initializeSystemClass(stringClass    , objectClass , "String"    );
//   initializeSystemClass(doubleClass    , objectClass , "Double"    );

//   // Load methods and fields into the system classes
//   loadSystemClass(objectClass);
//   loadSystemClass(classClass);
//   loadSystemClass(metaclassClass);
//   loadSystemClass(nilClass);
//   loadSystemClass(arrayClass);
//   loadSystemClass(methodClass);
//   loadSystemClass(symbolClass);
//   loadSystemClass(integerClass);
//   loadSystemClass(bigintegerClass);
//   loadSystemClass(frameClass);
//   loadSystemClass(primitiveClass);
//   loadSystemClass(stringClass);
//   loadSystemClass(doubleClass);

//   // Load the generic block class
//   blockClass = loadClass(symbolFor("Block"));

//   // Setup the true and false objects
//   trueObject = newInstance(loadClass(symbolFor("True")));
//   falseObject = newInstance(loadClass(symbolFor("False")));

//   // Load the system class and create an instance of it
//   systemClass = loadClass(symbolFor("System"));
//   Object systemObject = newInstance(systemClass);

//   // Put special objects and classes into the dictionary of globals
//   setGlobal(symbolFor("nil"), nilObject);
//   setGlobal(symbolFor("true"), trueObject);
//   setGlobal(symbolFor("false"), falseObject);
//   setGlobal(symbolFor("system"), systemObject);
//   setGlobal(symbolFor("System"), systemClass);
//   setGlobal(symbolFor("Block"), blockClass);

//   // Create a fake bootstrap method to simplify later frame traversal
//   Method bootstrapMethod = newMethod(symbolFor("bootstrap"), 1, 0);
//   bootstrapMethod.setBytecode(0, Bytecodes.halt);
//   bootstrapMethod.setNumberOfLocals(0);
//   bootstrapMethod.setMaximumNumberOfStackElements(2);
//   bootstrapMethod.setHolder(systemClass);

//   // Start the shell if no filename is given
//   if (arguments.length == 0)
//   {
// 	Shell.setBootstrapMethod(bootstrapMethod);
// 	Shell.start();
// 	return;
//   }

//   // Convert the arguments into an array
//   Array argumentsArray = newArray(arguments);

//   // Create a fake bootstrap frame with the system object on the stack
//   Frame bootstrapFrame = Interpreter.pushNewFrame(bootstrapMethod);
//   bootstrapFrame.push(systemObject);
//   bootstrapFrame.push(argumentsArray);

//   // Lookup the initialize invokable on the system class
//   Invokable initialize = systemClass.lookupInvokable(symbolFor("initialize:"));

//   // Invoke the initialize invokable
//   initialize.invoke(bootstrapFrame);

//   // Start the interpreter
//   Interpreter.start();
// }

// public static void _assert(boolean value)
//
//	{
//	  if (!value) {
//		// For now we just print something whenever an assertion fails
//		System.out.println("Assertion failed");
//	  }
//	}
func (u *Universe) Assert(value bool) {
	if !value {
		// For now we just print something whenever an assertion fails
		fmt.Println("Assertion failed")
	}
}

// public static Symbol symbolFor(java.lang.String string)
// {
//   // Lookup the symbol in the symbol table
//   Symbol result = SymbolTable.lookup(string);
//   if (result != null) return result;

//	  // Create a new symbol and return it
//	  result = newSymbol(string);
//	  return result;
//	}
// func (u *Universe) SymbolFor(str string) *Symbol {
// 	// Lookup the symbol in the symbol table
// 	result := u.symboltable.lookup(str)
// 	if result != nil {
// 		return result
// 	}

// 	// Create a new symbol and return it
// 	result = u.NewSymbol(str)
// 	return result
// }

// public static Array newArray(int length)
// {
//   // Allocate a new array and set its class to be the array class
//   Array result = new Array();
//   result.setClass(arrayClass);

//   // Set the number of indexable fields to the given value (length)
//   result.setNumberOfIndexableFields(length);

// // Return the freshly allocated array
// return result;
// }
func (u *Universe) NewArray(length int) *Array {
	// Allocate a new array and set its class to be the array class
	result := NewArray(length)
	result.setClass(u.ArrayClass)

	// Set the number of indexable fields to the given value (length)
	result.SetNumberOfIndexableFields(length)

	// Return the freshly allocated array
	return result
}

// public static Array newArray(java.util.List<?> list)
// {
//   // Allocate a new array with the same length as the list
//   Array result = newArray(list.size());

//   // Copy all elements from the list into the array
//   for (int i = 0; i < list.size(); i++) {
// 	result.setIndexableField(i, (Object) list.get(i));
//   }

// // Return the allocated and initialized array
// return result;
// }
// func (u *Universe) NewArray(list []interface{}) *Array {
// 	// Allocate a new array with the same length as the list
// 	result := u.NewArray(len(list))

// 	// Copy all elements from the list into the array
// 	for i := 0; i < len(list); i++ {
// 		result.SetIndexableField(i, list[i])
// 	}

// 	// Return the allocated and initialized array
// 	return result
// }

// public static Array newArray(java.lang.String[] stringArray)
// {
//   // Allocate a new array with the same length as the string array
//   Array result = newArray(stringArray.length);

//   // Copy all elements from the string array into the array
//   for (int i = 0; i < stringArray.length; i++) {
// 	result.setIndexableField(i, newString(stringArray[i]));
//   }

// // Return the allocated and initialized array
// return result;
// }
func (u *Universe) NewArrayFromStrings(stringArray []string) *Array {
	// Allocate a new array with the same length as the string array
	result := u.NewArray(len(stringArray))

	// Copy all elements from the string array into the array
	for i := 0; i < len(stringArray); i++ {
		result.SetIndexableField(i, u.NewString(stringArray[i]))
	}

	// Return the allocated and initialized array
	return result
}

// public static Block newBlock(Method method, Frame context, int arguments)
// {
//   // Allocate a new block and set its class to be the block class
//   Block result = new Block();
//   result.setClass(getBlockClass(arguments));

//   // Set the method and context of block
//   result.setMethod(method);
//   result.setContext(context);

// // Return the freshly allocated block
// return result;
// }
func (u *Universe) NewBlock(method *Method, context *Frame, arguments int) *Block {
	// Allocate a new block and set its class to be the block class
	result := NewBlock(arguments)
	//result.setClass(u.GetBlockClass())

	// Set the method and context of block
	result.SetMethod(method)
	result.SetContext(context)

	// Return the freshly allocated block
	return result
}

// public static Class newClass(Class classClass)
// {
//   // Allocate a new class and set its class to be the given class class
//   Class result = new Class(classClass.getNumberOfInstanceFields());
//   result.setClass(classClass);

// // Return the freshly allocated class
// return result;
// }
func (u *Universe) NewClass(classClass *Class) *Class {
	// Allocate a new class and set its class to be the given class class
	result := NewClass(classClass.GetNumberOfInstanceFields())
	result.setClass(classClass)

	// Return the freshly allocated class
	return result
}

// public static Frame newFrame(Frame previousFrame, Method method)
//
//	{
//	  // Allocate a new frame and set its class to be the frame class
//	  Frame result = new Frame();
//	  result.setClass(frameClass);
//	  // Compute the maximum number of stack locations (including arguments, locals and
//	  // extra buffer to support doesNotUnderstand) and set the number of
//	  // indexable fields accordingly
//	  int length = method.getNumberOfArguments() + method.getNumberOfLocals() + method.getMaximumNumberOfStackElements() + 2;
//	  result.setNumberOfIndexableFields(length);
//	  // Set the method of the frame and the previous frame
//	  result.setMethod(method);
//	  if (previousFrame != null) result.setPreviousFrame(previousFrame);
//	  // Reset the stack pointer and the bytecode index
//	  result.resetStackPointer();
//	  result.setBytecodeIndex(0);
//	  // Return the freshly allocated frame
//	  return result;
//	  }
func (u *Universe) NewFrame(previousFrame *Frame, method *Method) *Frame {
	// Allocate a new frame and set its class to be the frame class
	result := NewFrame()
	result.setClass(u.FrameClass)
	// Compute the maximum number of stack locations (including arguments, locals and
	// extra buffer to support doesNotUnderstand) and set the number of
	// indexable fields accordingly
	length := method.GetNumberOfArguments() + method.GetNumberOfLocals() + method.getMaximumNumberOfStackElements() + 2
	result.SetNumberOfIndexableFields(length)
	// Set the method of the frame and the previous frame
	result.SetMethod(method)
	if previousFrame != nil {
		result.SetPreviousFrame(previousFrame)
	}
	// Reset the stack pointer and the bytecode index
	result.ResetStackPointer()
	result.SetBytecodeIndex(0)
	// Return the freshly allocated frame
	return result
}

// public static Method newMethod(Symbol signature, int numberOfBytecodes, int numberOfLiterals)
// {
//   // Allocate a new method and set its class to be the method class
//   Method result = new Method();
//   result.setClass(methodClass);

//   // Set the signature and the number of bytecodes
//   result.setSignature(signature);
//   result.setNumberOfBytecodes(numberOfBytecodes);
//   result.setNumberOfIndexableFields(numberOfLiterals);

// // Return the freshly allocated method
// return result;
// }
func (u *Universe) NewMethod(signature *Symbol, numberOfBytecodes int, numberOfLiterals int) *Method {
	// Allocate a new method and set its class to be the method class
	result := NewMethod()
	result.setClass(u.MethodClass)

	// Set the signature and the number of bytecodes
	result.setSignature(signature)
	result.setNumberOfBytecodes(numberOfBytecodes)
	result.setNumberOfIndexableFields(numberOfLiterals)

	// Return the freshly allocated method
	return result
}

// public static Object newInstance(Class instanceClass)
//
//	{
//	  // Allocate a new instance and set its class to be the given class
//	  Object result = new Object(instanceClass.getNumberOfInstanceFields());
//	  result.setClass(instanceClass);
//	  // Return the freshly allocated instance
//	  return result;
//	  }
func (u *Universe) NewInstance(instanceClass *Class) *Object {
	// Allocate a new instance and set its class to be the given class
	result := NewObject(instanceClass.getNumberOfInstanceFields())
	result.setClass(instanceClass)
	// Return the freshly allocated instance
	return result
}

// public static Integer newInteger(int value)
// {
//   // Allocate a new integer and set its class to be the integer class
//   Integer result = new Integer();
//   result.setClass(integerClass);

//   // Set the embedded integer of the newly allocated integer
//   result.setEmbeddedInteger(value);

// // Return the freshly allocated integer
// return result;
// }
func (u *Universe) NewInteger(value int) *Integer {
	// Allocate a new integer and set its class to be the integer class
	result := NewInteger()
	result.setClass(u.IntegerClass)

	// Set the embedded integer of the newly allocated integer
	result.setEmbeddedInteger(value)

	// Return the freshly allocated integer
	return result
}

//   public static BigInteger newBigInteger(java.math.BigInteger value)
//   {
//   // Allocate a new integer and set its class to be the integer class
//   BigInteger result = new BigInteger();
//   result.setClass(bigintegerClass);

//   // Set the embedded integer of the newly allocated integer
//   result.setEmbeddedBiginteger(value);

// // Return the freshly allocated integer
// return result;
// }
func (u *Universe) NewBigInteger(value *big.Int) *BigInteger {
	// Allocate a new integer and set its class to be the integer class
	result := NewBigInteger()
	result.setClass(u.BigIntegerClass)

	// Set the embedded integer of the newly allocated integer
	result.setEmbeddedBiginteger(value)

	// Return the freshly allocated integer
	return result
}

//   public static BigInteger newBigInteger(long value)
//   {
//   // Allocate a new integer and set its class to be the integer class
//   BigInteger result = new BigInteger();
//   result.setClass(bigintegerClass);

//   // Set the embedded integer of the newly allocated integer
//   result.setEmbeddedBiginteger(new java.math.BigInteger(java.lang.Long.valueOf(value).toString()));

// // Return the freshly allocated integer
// return result;
// }
func (u *Universe) NewBigInteger(value int64) *BigInteger {
	// Allocate a new integer and set its class to be the integer class
	result := NewBigInteger()
	result.setClass(u.BigIntegerClass)

	// Set the embedded integer of the newly allocated integer
	result.setEmbeddedBiginteger(big.NewInt(value))

	// Return the freshly allocated integer
	return result
}

//  public static Double newDouble(double value)
// {
//   // Allocate a new integer and set its class to be the double class
//   Double result = new Double();
//   result.setClass(doubleClass);

//   // Set the embedded double of the newly allocated double
//   result.setEmbeddedDouble(value);

// // Return the freshly allocated double
// return result;
// }
func (u *Universe) NewDouble(value float64) *Double {
	// Allocate a new integer and set its class to be the double class
	result := NewDouble()
	result.setClass(u.DoubleClass)

	// Set the embedded double of the newly allocated double
	result.setEmbeddedDouble(value)

	// Return the freshly allocated double
	return result
}

// public static Class newMetaclassClass()
// {
//   // Allocate the metaclass classes
//   Class result = new Class();
//   result.setClass(new Class());

//   // Setup the metaclass hierarchy
//   result.getSOMClass().setClass(result);

// // Return the freshly allocated metaclass class
// return result;
// }
func (u *Universe) NewMetaclassClass() *Class {
	// Allocate the metaclass classes
	result := NewClass()
	result.setClass(NewClass())

	// Setup the metaclass hierarchy
	result.getSOMClass().setClass(result)

	// Return the freshly allocated metaclass class
	return result
}

// public static String newString(java.lang.String embeddedString)
// {
//   // Allocate a new string and set its class to be the string class
//   String result = new String();
//   result.setClass(stringClass);

//   // Put the embedded string into the new string
//   result.setEmbeddedString(embeddedString);

// // Return the freshly allocated string
// return result;
// }
func (u *Universe) NewString(embeddedString string) *String {
	// Allocate a new string and set its class to be the string class
	result := NewString()
	result.setClass(u.StringClass)

	// Put the embedded string into the new string
	result.setEmbeddedString(embeddedString)

	// Return the freshly allocated string
	return result
}

// public static Symbol newSymbol(java.lang.String string)
// {
//   // Allocate a new symbol and set its class to be the symbol class
//   Symbol result = new Symbol();
//   result.setClass(symbolClass);

//   // Put the string into the symbol
//   result.setString(string);

//   // Insert the new symbol into the symbol table
//   SymbolTable.insert(result);

// // Return the freshly allocated symbol
// return result;
// }
func (u *Universe) NewSymbol(str string) *Symbol {
	// Allocate a new symbol and set its class to be the symbol class
	result := NewSymbol()
	result.setClass(u.SymbolClass)

	// Put the string into the symbol
	result.setString(str)

	// Insert the new symbol into the symbol table
	u.symboltable.insert(result)

	// Return the freshly allocated symbol
	return result
}

// public static Class newSystemClass()
//
//	{
//	  // Allocate the new system class
//	  Class systemClass = new Class();
//	  // Setup the metaclass hierarchy
//	  systemClass.setClass(new Class());
//	  systemClass.getSOMClass().setClass(metaclassClass);
//	  // Return the freshly allocated system class
//	  return systemClass;
//	  }
func (u *Universe) NewSystemClass() *Class {
	systemClass = NewClass()
	// Setup the metaclass hierarchy
	systemClass.setClass(NewClass())
	systemClass.getSOMClass().setClass(metaclassClass)

	return systemClass
}

// public static void initializeSystemClass(Class systemClass, Class superClass, java.lang.String name)
// {
//   // Initialize the superclass hierarchy
//   if (superClass != null) {
// 	systemClass.setSuperClass(superClass);
// 	systemClass.getSOMClass().setSuperClass(superClass.getSOMClass());
//   } else {
// 	systemClass.getSOMClass().setSuperClass(classClass);
//   }

//   // Initialize the array of instance fields
//   systemClass.setInstanceFields(newArray(0));
//   systemClass.getSOMClass().setInstanceFields(newArray(0));

//   // Initialize the array of instance invokables
//   systemClass.setInstanceInvokables(newArray(0));
//   systemClass.getSOMClass().setInstanceInvokables(newArray(0));

//   // Initialize the name of the system class
//   systemClass.setName(symbolFor(name));
//   systemClass.getSOMClass().setName(symbolFor(name + " class"));

//	  // Insert the system class into the dictionary of globals
//	  setGlobal(systemClass.getName(), systemClass);
//	}
func (u *Universe) InitializeSystemClass(systemClass *Class, superClass *Class, name string) {
	if superClass != nil {
		systemClass.setSuperClass(superClass)
		systemClass.getSOMClass().setSuperClass(superClass.getSOMClass())
	} else {
		systemClass.getSOMClass().setSuperClass(classClass)
	}
	// Initialize the array of instance fields
	systemClass.setInstanceFields(u.NewArray(0))
	systemClass.getSOMClass().setInstanceFields(u.NewArray(0))
	// Initialize the array of instance invokables
	systemClass.setInstanceInvokables(u.NewArray(0))
	systemClass.getSOMClass().setInstanceInvokables(u.NewArray(0))
	// Initialize the name of the system class
	systemClass.setName(u.SymbolFor(name))
	systemClass.getSOMClass().setName(u.SymbolFor(name + " class"))
	// Insert the system class into the dictionary of globals
	u.setGlobal(systemClass.getName(), systemClass)
}

// public static Object getGlobal(Symbol name)
// {
//   // Return the global with the given name if it's in the dictionary of globals
//   if (hasGlobal(name)) return (Object) globals.get(name);

//	  // Global not found
//	  return null;
//	}
func (u *Universe) GetGlobal(name *Symbol) *Object {
	// Return the global with the given name if it's in the dictionary of globals
	if u.hasGlobal(name) {
		return u.globals[name]
	}

	// Global not found
	return nil
}

// public static void setGlobal(Symbol name, Object value)
//
//	{
//	  // Insert the given value into the dictionary of globals
//	  globals.put(name, value);
//	}
func (u *Universe) SetGlobal(name *Symbol, value *Object) {
	// Insert the given value into the dictionary of globals
	u.globals[name] = value
}

// public static boolean hasGlobal(Symbol name)
//
//	{
//	  // Returns if the universe has a value for the global of the given name
//	  return globals.containsKey(name);
//	}
func (u *Universe) HasGlobal(name *Symbol) bool {
	// Returns if the universe has a value for the global of the given name
	return u.globals[name] != nil
}

// public static Class getBlockClass()
//
//	{
//	  // Get the generic block class
//	  return blockClass;
//	}
func (u *Universe) GetBlockClass() *Class {
	// Get the generic block class
	return blockClass
}

// public static Class getBlockClass(int numberOfArguments)
//
//	{
//	  // Compute the name of the block class with the given number of arguments
//	  Symbol name = symbolFor("Block" + java.lang.Integer.toString(numberOfArguments));
//	  // Lookup the specific block class in the dictionary of globals and return it
//	  if (hasGlobal(name)) return (Class) getGlobal(name);
//	  // Get the block class for blocks with the given number of arguments
//	  Class result = loadClass(name, null);
//	  // Add the appropriate value primitive to the block class
//	  result.addInstancePrimitive(Block.getEvaluationPrimitive(numberOfArguments));
//	  // Insert the block class into the dictionary of globals
//	  setGlobal(name, result);
//	  // Return the loaded block class
//	  return result;
//	}
func (u *Universe) NewBlockClass(numberOfArguments int) *Class {
	// Compute the name of the block class with the given number of arguments
	name := u.SymbolFor("Block" + strconv.Itoa(numberOfArguments))
	// Lookup the specific block class in the dictionary of globals and return it
	if u.HasGlobal(name) {
		return u.GetGlobal(name)
	}
	// Get the block class for blocks with the given number of arguments
	result := u.LoadClass(name, nil)
	// Add the appropriate value primitive to the block class
	result.addInstancePrimitive(Block.getEvaluationPrimitive(numberOfArguments))
	// Insert the block class into the dictionary of globals
	u.SetGlobal(name, result)
	// Return the loaded block class
	return result
}

// public static Class loadClass(Symbol name)
//
//	{
//	  // Check if the requested class is already in the dictionary of globals
//	  if (hasGlobal(name)) return (Class) getGlobal(name);
//	  // Load the class
//	  Class result = loadClass(name, null);
//	  // Load primitives (if necessary) and return the resulting class
//	  if(result != null && result.hasPrimitives()) result.loadPrimitives();
//	  return result;
//	}
func loadClass(name *Symbol) *Class {
	// Check if the requested class is already in the dictionary of globals
	if u.HasGlobal(name) {
		return u.GetGlobal(name)
	}
	// Load the class
	result := u.LoadClass(name, nil)
	// Load primitives (if necessary) and return the resulting class
	if result != nil && result.hasPrimitives() {
		result.loadPrimitives()
	}
	return result
}

// public static void loadSystemClass(Class systemClass)
//
//	{
//	  // Load the system class
//	  Class result = loadClass(systemClass.getName(), systemClass);
//	  if (result == null) {
//		  System.out.println(systemClass.getName().getString());
//			System.out.println(" failed: loadClass(systemClass.getName(), systemClass)");
//	  }
//	  // Load primitives if necessary
//	  if(result.hasPrimitives()) result.loadPrimitives();
//	}
func loadSystemClass(systemClass *Class) {
	result := loadClass(systemClass.getName(), systemClass)
	if result == nil {
		fmt.Println(systemClass.getName().getString())
		fmt.Println("Failed: loadClass(systemClass.getName(), systemClass)")
	}
	if result.hasPrimitives() {
		result.loadPrimitives()
	}
}

// public static Class loadClass(Symbol name, Class systemClass)
//
//	{
//	  // Try loading the class from all different paths
//	  for(java.lang.String cpEntry : classPath) {
//		try {
//		  // Load the class from a file and return the loaded class
//		  Class result = som.compiler.SourcecodeCompiler.compileClass(cpEntry + fileSeparator,
//												  name.getString(),
//												  systemClass);
//		  if(dumpBytecodes) {
//			  Disassembler.dump(result.getSOMClass());
//			  Disassembler.dump(result);
//		  }
//		  return result;
//		} catch (IOException e) {
//								  // Continue trying different paths
//		}
//	  }
//	  // The class could not be found.
//	  System.out.println(name.getString());
//	  System.out.println(" The class could not be found");
//	  return null;
//	}
func loadClass(name *Symbol, systemClass *Class) *Class {
	// Try loading the class from all different paths
	for _, cpEntry := range classPath {
		// Load the class from a file and return the loaded class
		result := som.compiler.SourcecodeCompiler.compileClass(cpEntry+fileSeparator, name.getString(), systemClass)
		if dumpBytecodes {
			Disassembler.dump(result.getSOMClass())
			Disassembler.dump(result)
		}
		return result
	}
	// The class could not be found.
	fmt.Println(name.getString())
	fmt.Println(" The class could not be found")
	return nil
}

// public static Class loadShellClass(java.lang.String stmt) throws IOException
// {
// 	//java.io.ByteArrayInputStream in = new java.io.ByteArrayInputStream(stmt.getBytes());

//		// Load the class from a stream and return the loaded class
//		Class result = som.compiler.SourcecodeCompiler.compileClass(stmt, null);
//		if(dumpBytecodes)
//			Disassembler.dump(result);
//		return result;
//	}
func loadShellClass(stmt string) *Class {
	// Load the class from a stream and return the loaded class
	result := som.compiler.SourcecodeCompiler.compileClass(stmt, nil)
	if dumpBytecodes {
		Disassembler.dump(result)
	}
	return result
}

// public static Object nilObject;
// public static Object trueObject;
// public static Object falseObject;

// public static Class objectClass;
// public static Class classClass;
// public static Class metaclassClass;

// public static Class nilClass;
// public static Class integerClass;
// public static Class bigintegerClass;
// public static Class arrayClass;
// public static Class methodClass;
// public static Class symbolClass;
// public static Class frameClass;
// public static Class primitiveClass;
// public static Class stringClass;
// public static Class systemClass;
// public static Class blockClass;
// public static Class doubleClass;

// private static java.util.HashMap<Symbol,som.vmobjects.Object> globals = new java.util.HashMap<Symbol,som.vmobjects.Object>();
// private static java.lang.String[] classPath;
// private static boolean dumpBytecodes;

// public static java.lang.String pathSeparator;
// public static java.lang.String fileSeparator;
// }
