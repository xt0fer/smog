package vmobj

type VMClass struct {
	VMObject
	classname string
	nCvars    int
	nMethods  int
}

func NewVMClass(name string, nCvars int, nMethods int) *VMClass {
	return &VMClass{NewObjInternal(0, nil), name, nCvars, nMethods}
}

func (c *VMClass) ToString() string {
	return c.classname
}

func (c *VMClass) Name() string {
	return c.classname
}

func (c *VMClass) NumberOfClassVars() int {
	return c.nCvars
}

func (c *VMClass) NumberOfMethods() int {
	return c.nMethods
}
