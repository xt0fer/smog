package compiler

import (
	"github.com/xt0fer/smog"
)

// private som.vmobjects.Symbol name;
// private som.vmobjects.Symbol superName;
// private boolean classSide;
// private List<som.vmobjects.Object> instanceFields = new ArrayList<som.vmobjects.Object>();
// private List<som.vmobjects.Invokable> instanceMethods = new ArrayList<som.vmobjects.Invokable>();
// private List<som.vmobjects.Object> classFields = new ArrayList<som.vmobjects.Object>();
// private List<som.vmobjects.Invokable> classMethods = new ArrayList<som.vmobjects.Invokable>();
type ClassGenerator struct {
	name            *smog.Symbol
	superName       *smog.Symbol
	classSide       bool
	instanceFields  []*smog.Object
	instanceMethods []smog.Invokable
	classFields     []*smog.Object
	classMethods    []smog.Invokable
}

//	public void setName(Symbol name) {
//		this.name = name;
//	}
func (cg *ClassGenerator) SetName(name *smog.Symbol) {
	cg.name = name
}

//	public void setSuperName(Symbol superName) {
//		this.superName = superName;
//	}
func (cg *ClassGenerator) SetSuperName(superName *smog.Symbol) {
	cg.superName = superName
}

//	public void addInstanceMethod(som.vmobjects.Invokable meth) {
//		instanceMethods.add(meth);
//	}
func (cg *ClassGenerator) AddInstanceMethod(meth smog.Invokable) {
	cg.instanceMethods = append(cg.instanceMethods, meth)
}

//	public void setClassSide(boolean b) {
//		classSide = b;
//	}
func (cg *ClassGenerator) SetClassSide(b bool) {
	cg.classSide = b
}

//	public void addClassMethod(som.vmobjects.Invokable meth) {
//		classMethods.add(meth);
//	}
func (cg *ClassGenerator) AddClassMethod(meth smog.Invokable) {
	cg.classMethods = append(cg.classMethods, meth)
}

//	public void addInstanceField(Symbol field) {
//		instanceFields.add(field);
//	}
func (cg *ClassGenerator) AddInstanceField(field *smog.Object) {
	cg.instanceFields = append(cg.instanceFields, field)
}

//	public void addClassField(Symbol field) {
//		classFields.add(field);
//	}
func (cg *ClassGenerator) AddClassField(field *smog.Object) {
	cg.classFields = append(cg.classFields, field)
}

//	public boolean findField(String field) {
//		return (isClassSide() ? classFields : instanceFields).indexOf(Universe.symbolFor(field)) != -1;
//	}
func (cg *ClassGenerator) FindField(field string) bool {
	var fields []*smog.Object
	if cg.classSide {
		fields = cg.classFields
	} else {
		fields = cg.instanceFields
	}
	for _, f := range fields {
		if f.GetSOMClass().Name.Name == field {
			return true
		}
	}
	return false
}

//	public boolean isClassSide() {
//		return classSide;
//	}
func (cg *ClassGenerator) IsClassSide() bool {
	return cg.classSide
}

//	public som.vmobjects.Class assemble() {
//		// build class class name
//		String ccname = name.getString() + " class";
//		// Load the super class
//		som.vmobjects.Class superClass = Universe.loadClass(superName);
//		// Allocate the class of the resulting class
//		som.vmobjects.Class resultClass = Universe.newClass(Universe.metaclassClass);
//		// Initialize the class of the resulting class
//		resultClass.setInstanceFields(Universe.newArray(classFields));
//		resultClass.setInstanceInvokables(Universe.newArray(classMethods));
//		resultClass.setName(Universe.symbolFor(ccname));
//		som.vmobjects.Class superMClass = superClass.getSOMClass();
//		resultClass.setSuperClass(superMClass);
//		// Allocate the resulting class
//		som.vmobjects.Class result = Universe.newClass(resultClass);
//		// Initialize the resulting class
//		result.setInstanceFields(Universe.newArray(instanceFields));
//		result.setInstanceInvokables(Universe.newArray(instanceMethods));
//		result.setName(name);
//		result.setSuperClass(superClass);
//		return result;
//	}
func (cg *ClassGenerator) Assemble() *smog.Class {
	// build class class name
	ccname := cg.name.String() + " class"
	u := smog.GetUniverse()
	// Load the super class
	superClass := u.LoadClass(cg.superName)
	// Allocate the class of the resulting class
	resultClass := u.NewClass(u.metaclassClass)
	// Initialize the class of the resulting class
	resultClass.SetInstanceFields(u.NewArray(cg.classFields))
	resultClass.SetInstanceInvokables(u.NewArray(cg.classMethods))
	resultClass.SetName(u.SymbolFor(ccname))
	superMClass := superClass.GetSOMClass()
	resultClass.SetSuperClass(superMClass)
	// Allocate the resulting class
	result := u.NewClass(resultClass)
	// Initialize the resulting class
	result.SetInstanceFields(u.NewArray(instanceFields))
	result.SetInstanceInvokables(u.NewArray(instanceMethods))
	result.SetName(cg.name)
	result.SetSuperClass(superClass)
	return result

}

// public void assembleSystemClass(som.vmobjects.Class systemClass) {
// 	systemClass.setInstanceInvokables(Universe.newArray(instanceMethods));
// 	systemClass.setInstanceFields(Universe.newArray(instanceFields));
// 	// class-bound == class-instance-bound
// 	som.vmobjects.Class superMClass = systemClass.getSOMClass();
// 	superMClass.setInstanceInvokables(Universe.newArray(classMethods));
// 	superMClass.setInstanceFields(Universe.newArray(classFields));
// }

// the NEWARRAY issue, need s special NewArray in Universe to handle these initializations
func (cg *ClassGenerator) AssembleSystemClass(systemClass *smog.Class) {
	u := smog.GetUniverse()
	systemClass.SetInstanceInvokables(u.NewArray(cg.instanceMethods))
	systemClass.SetInstanceFields(u.NewArray(cg.instanceFields))
	// class-bound == class-instance-bound
	superMClass := systemClass.GetSOMClass()
	superMClass.SetInstanceInvokables(u.NewArray(cg.classMethods))
	superMClass.SetInstanceFields(u.NewArray(cg.classFields))
}
