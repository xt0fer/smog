// universe.go
package smog

import "github.com/xtofer/src/smog/internal/vm"

// see OldCode/som-src/som/src/som/vm/Universe.java for information on the functionality needed here

type Universe struct {
	// The universe is a collection of all the objects in the simulation
	objects []vm.Object
}

// public static Object nilObject;
// public static Object trueObject;
// public static Object falseObject;
var nilObject vm.Object
var trueObject vm.Object
var falseObject vm.Object

// public static Class objectClass;
// public static Class classClass;
// public static Class metaclassClass;
var objectClass vm.Class
var classClass vm.Class
var metaclassClass vm.Class

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
var (
	nilClass        vm.Class
	integerClass    vm.Class
	bigintegerClass vm.Class
	arrayClass      vm.Class
	methodClass     vm.Class
	symbolClass     vm.Class
	frameClass      vm.Class
	primitiveClass  vm.Class
	stringClass     vm.Class
	systemClass     vm.Class
	blockClass      vm.Class
	doubleClass     vm.Class
)
// private static java.util.HashMap<Symbol,som.vmobjects.Object> globals = new java.util.HashMap<Symbol,som.vmobjects.Object>();
var (
	globals map[vm.Symbol]vm.Object
)
// private static java.lang.String[] classPath;
// private static boolean dumpBytecodes;

func newUniverse() *Universe {
	return new(Universe)
}

func (u *Universe) Initialize(arguments []string) *Universe {
	// Initialize the known universe
	u.objects = make([]vm.Object, 0)
	return u
}
