package smog

import (
	"fmt"
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
	Shell           *Shell
	SystemClass     *Class
	SystemObject    *Object
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

func (u *Universe) SymbolFor(sym string) *Symbol {
	result := u.symboltable.lookup(sym)
	if result != nil {
		return result
	}
	result = u.NewSymbol(sym)
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
func (u *Universe) initialize(arguments []string) {
	// Allocate the nil object
	u.NilObject = NewObject(-1)
	// Allocate the Metaclass classes
	u.MetaclassClass = u.NewMetaclassClass()
	// Allocate the rest of the system classes
	u.ObjectClass = u.NewSystemClass()
	u.NilClass = u.NewSystemClass()
	u.ClassClass = u.NewSystemClass()
	u.ArrayClass = u.NewSystemClass()
	u.SymbolClass = u.NewSystemClass()
	u.MethodClass = u.NewSystemClass()
	u.IntegerClass = u.NewSystemClass()
	u.BigIntegerClass = u.NewSystemClass()
	u.FrameClass = u.NewSystemClass()
	u.PrimitiveClass = u.NewSystemClass()
	u.StringClass = u.NewSystemClass()
	u.DoubleClass = u.NewSystemClass()

	//   // Setup the class reference for the nil object
	u.NilObject.SetClass(u.NilClass)

	// // Initialize the system classes.
	// // Note: The order in which the system classes are initialized is important
	// //       since their names are used to bootstrap the system.
	// InitializeSystemClass(objectClass    , null         , "Object"    );
	// InitializeSystemClass(classClass     , objectClass , "Class"     );
	// InitializeSystemClass(metaclassClass , classClass  , "Metaclass" );
	// InitializeSystemClass(nilClass       , objectClass , "Nil"       );
	// InitializeSystemClass(arrayClass     , objectClass , "Array"     );
	u.InitializeSystemClass(u.ObjectClass, nil, "Object")
	u.InitializeSystemClass(u.ClassClass, u.ObjectClass, "Class")
	u.InitializeSystemClass(u.MetaclassClass, u.ClassClass, "Metaclass")
	u.InitializeSystemClass(u.NilClass, u.ObjectClass, "Nil")
	u.InitializeSystemClass(u.ArrayClass, u.ObjectClass, "Array")
	// InitializeSystemClass(methodClass    , arrayClass  , "Method"    );
	// InitializeSystemClass(symbolClass    , objectClass , "Symbol"    );
	// InitializeSystemClass(integerClass   , objectClass , "Integer"   );
	// InitializeSystemClass(bigintegerClass, objectClass , "BigInteger");
	u.InitializeSystemClass(u.MethodClass, u.ArrayClass, "Method")
	u.InitializeSystemClass(u.SymbolClass, u.ObjectClass, "Symbol")
	u.InitializeSystemClass(u.IntegerClass, u.ObjectClass, "Integer")
	u.InitializeSystemClass(u.BigIntegerClass, u.ObjectClass, "BigInteger")

	// InitializeSystemClass(frameClass     , arrayClass  , "Frame"     );
	// InitializeSystemClass(primitiveClass , objectClass , "Primitive" );
	// InitializeSystemClass(stringClass    , objectClass , "String"    );
	// InitializeSystemClass(doubleClass    , objectClass , "Double"    );
	u.InitializeSystemClass(u.FrameClass, u.ArrayClass, "Frame")
	u.InitializeSystemClass(u.PrimitiveClass, u.ObjectClass, "Primitive")
	u.InitializeSystemClass(u.StringClass, u.ObjectClass, "String")
	u.InitializeSystemClass(u.DoubleClass, u.ObjectClass, "Double")
	// // Load methods and fields into the system classes
	// loadSystemClass(objectClass);
	loadSystemClass(u.ObjectClass)
	// loadSystemClass(classClass);
	// loadSystemClass(metaclassClass);
	// loadSystemClass(nilClass);
	// loadSystemClass(arrayClass);
	loadSystemClass(u.ClassClass)
	loadSystemClass(u.MetaclassClass)
	loadSystemClass(u.NilClass)
	loadSystemClass(u.ArrayClass)
	// loadSystemClass(methodClass);
	// loadSystemClass(symbolClass);
	// loadSystemClass(integerClass);
	// loadSystemClass(bigintegerClass);
	loadSystemClass(u.MethodClass)
	loadSystemClass(u.SymbolClass)
	loadSystemClass(u.IntegerClass)
	loadSystemClass(u.BigIntegerClass)
	// loadSystemClass(frameClass);
	// loadSystemClass(primitiveClass);
	// loadSystemClass(stringClass);
	loadSystemClass(u.FrameClass)
	loadSystemClass(u.PrimitiveClass)
	loadSystemClass(u.StringClass)
	// loadSystemClass(doubleClass);
	loadSystemClass(u.DoubleClass)
	// // Load the generic block class
	// blockClass = loadClass(symbolFor("Block"));
	u.BlockClass = u.LoadClass(u.SymbolFor("Block"))
	// // Setup the true and false objects
	// trueObject = newInstance(loadClass(symbolFor("True")));
	u.TrueObject = instantiated.NewInstance(u.LoadClass(u.SymbolFor("True")))
	// falseObject = newInstance(loadClass(symbolFor("False")));
	u.FalseObject = instantiated.NewInstance(u.LoadClass(u.SymbolFor("False")))
	// // Load the system class and create an instance of it
	// systemClass = loadClass(symbolFor("System"));
	u.SystemClass = u.LoadClass(u.SymbolFor("System"))
	// Object systemObject = newInstance(systemClass);
	u.SystemObject = u.NewInstance(u.SystemClass)
	// // Put special objects and classes into the dictionary of globals
	// setGlobal(symbolFor("nil"), nilObject);
	u.SetGlobal(u.SymbolFor("nil"), u.NilObject)
	// setGlobal(symbolFor("true"), trueObject);
	u.SetGlobal(u.SymbolFor("true"), u.TrueObject)
	// setGlobal(symbolFor("false"), falseObject);
	u.SetGlobal(u.SymbolFor("false"), u.FalseObject)
	// setGlobal(symbolFor("system"), systemObject);
	u.SetGlobal(u.SymbolFor("system"), u.SystemObject)
	// setGlobal(symbolFor("System"), systemClass);
	u.SetGlobal(u.SymbolFor("System"), u.SystemClass)
	// setGlobal(symbolFor("Block"), blockClass);
	u.SetGlobal(u.SymbolFor("Block"), u.BlockClass)
	// // Create a fake bootstrap method to simplify later frame traversal
	// Method bootstrapMethod = newMethod(symbolFor("bootstrap"), 1, 0);
	bootstrapMethod := u.NewMethod(u.SymbolFor("bootstrap"), 1, 0)
	// bootstrapMethod.setBytecode(0, Bytecodes.halt);
	bootstrapMethod.SetBytecode(0, HALT)
	// bootstrapMethod.setNumberOfLocals(0);
	bootstrapMethod.SetNumberOfLocals(0)
	// bootstrapMethod.setMaximumNumberOfStackElements(2);
	bootstrapMethod.SetMaximumNumberOfStackElements(2)
	// bootstrapMethod.setHolder(systemClass);
	bootstrapMethod.SetHolder(u.SystemClass)
	// // Start the shell if no filename is given
	// if (arguments.length == 0)
	if len(arguments) == 0 {
		//	  {
		//		Shell.setBootstrapMethod(bootstrapMethod);
		u.Shell = &Shell{}
		u.Shell.SetBootstrapMethod(bootstrapMethod)
		// Shell.start();
		u.Shell.Start()
		// return;
		return
		// }
	}
	// // Convert the arguments into an array
	// Array argumentsArray = newArray(arguments);
	argumentsArray := u.NewArrayFromStrings(arguments)
	// // Create a fake bootstrap frame with the system object on the stack
	// Frame bootstrapFrame = Interpreter.pushNewFrame(bootstrapMethod);
	bootstrapFrame := GetInterpreter().PushNewFrame(bootstrapMethod)
	// bootstrapFrame.push(systemObject);
	bootstrapFrame.Push(u.SystemObject)
	// bootstrapFrame.push(argumentsArray);
	bootstrapFrame.Push(argumentsArray)

	// // Lookup the initialize invokable on the system class
	// Invokable initialize = systemClass.lookupInvokable(symbolFor("initialize:"));
	initialize := u.SystemClass.LookupInvokable(u.SymbolFor("initialize:"))
	// // Invoke the initialize invokable
	// initialize.invoke(bootstrapFrame);
	initialize.Invoke(bootstrapFrame)

	// // Start the interpreter
	// Interpreter.start();
	GetInterpreter().Start()
	// }
}

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
	result.SetClass(u.ArrayClass)

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
	result.SetClass(classClass)

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
	result.SetClass(u.FrameClass)
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
	result.SetClass(u.MethodClass)

	// Set the signature and the number of bytecodes
	result.SetSignature(signature)
	result.SetNumberOfBytecodes(numberOfBytecodes)
	result.SetNumberOfIndexableFields(numberOfLiterals)

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
	result := NewObject(instanceClass.GetNumberOfInstanceFields())
	result.SetClass(instanceClass)
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
	result := NewInteger(value)
	result.SetClass(u.IntegerClass)

	// Set the embedded integer of the newly allocated integer
	//result.setEmbeddedInteger(value)

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
	result := NewBigInteger(value)
	result.SetClass(u.BigIntegerClass)

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
	result := NewDouble(value)
	result.SetClass(u.DoubleClass)
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
	result := NewClass(0)
	result.SetClass(NewClass(0))

	// Setup the metaclass hierarchy
	result.GetSOMClass().SetClass(result)

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
	result := NewString(embeddedString)
	result.SetClass(u.StringClass)
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
	result := NewSymbol(str)
	result.SetClass(u.SymbolClass)
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
	systemClass := NewClass(0)
	// Setup the metaclass hierarchy
	systemClass.SetClass(NewClass(0))
	systemClass.GetSOMClass().SetClass(u.MetaclassClass)

	return systemClass
}

// public static void InitializeSystemClass(Class systemClass, Class superClass, java.lang.String name)
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
		systemClass.SetSuperClass(superClass)
		systemClass.GetSOMClass().SetSuperClass(superClass.GetSOMClass())
	} else {
		systemClass.GetSOMClass().SetSuperClass(u.ClassClass)
	}
	// Initialize the array of instance fields
	systemClass.SetInstanceFields(u.NewArray(0))
	systemClass.GetSOMClass().SetInstanceFields(u.NewArray(0))
	// Initialize the array of instance invokables
	systemClass.SetInstanceInvokables(u.NewArray(0))
	systemClass.GetSOMClass().SetInstanceInvokables(u.NewArray(0))
	// Initialize the name of the system class
	systemClass.SetName(u.SymbolFor(name))
	systemClass.GetSOMClass().SetName(u.SymbolFor(name + " class"))
	// Insert the system class into the dictionary of globals
	u.SetGlobal(systemClass.GetName(), systemClass)
}

// public static Object getGlobal(Symbol name)
// {
//   // Return the global with the given name if it's in the dictionary of globals
//   if (hasGlobal(name)) return (Object) globals.get(name);

//	  // Global not found
//	  return null;
//	}
func (u *Universe) GetGlobal(name *Symbol) interface{} {
	// Return the global with the given name if it's in the dictionary of globals
	if u.HasGlobal(name) {
		return u.Globals[name]
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
func (u *Universe) SetGlobal(name *Symbol, value interface{}) {
	// Insert the given value into the dictionary of globals
	u.Globals[name] = value
}

// public static boolean hasGlobal(Symbol name)
//
//	{
//	  // Returns if the universe has a value for the global of the given name
//	  return globals.containsKey(name);
//	}
func (u *Universe) HasGlobal(name *Symbol) bool {
	// Returns if the universe has a value for the global of the given name
	return u.Globals[name] != nil
}

// public static Class getBlockClass()
//
//	{
//	  // Get the generic block class
//	  return blockClass;
//	}
func (u *Universe) GetBlockClass() *Class {
	// Get the generic block class
	return u.BlockClass
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
		return u.GetGlobal(name).(*Class)
	}
	// Get the block class for blocks with the given number of arguments
	result := LoadClass(name, nil)
	// Add the appropriate value primitive to the block class
	result.AddInstancePrimitive(&GetEvaluationPrimitive(numberOfArguments).Primitive)
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
func (u *Universe) LoadClass(name *Symbol) *Class {
	// Check if the requested class is already in the dictionary of globals
	if u.HasGlobal(name) {
		return u.GetGlobal(name).(*Class)
	}
	// Load the class
	result := LoadClass(name, nil)
	// Load primitives (if necessary) and return the resulting class
	if result != nil && result.HasPrimitives() {
		result.LoadPrimitives()
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
	result := LoadClass(systemClass.GetName(), systemClass)
	if result == nil {
		fmt.Println(systemClass.GetName().String())
		fmt.Println("Failed: loadClass(systemClass.getName(), systemClass)")
	}
	if result.HasPrimitives() {
		result.LoadPrimitives()
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
func LoadClass(name *Symbol, systemClass *Class) *Class {
	u := GetUniverse()
	// Try loading the class from all different paths
	for _, cpEntry := range u.classPath {
		// Load the class from a file and return the loaded class
		// SourceCodeCompiler
		result := SourcecodeCompileClass(cpEntry+u.FileSep, name.String(), systemClass)
		if u.dumpBytecodes {
			DisassemblerDump(result.GetSOMClass())
			DisassemblerDump(result)
		}
		return result
	}
	// The class could not be found.
	fmt.Println(name.String())
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
func LoadShellClass(stmt string) *Class {
	u := GetUniverse()
	// Load the class from a stream and return the loaded class
	result := SourcecodeCompileClass("", stmt, nil)
	if u.dumpBytecodes {
		DisassemblerDump(result)
	}
	return result
}
