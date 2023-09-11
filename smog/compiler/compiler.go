package compiler

import (
	"io"

	"github.com/xt0fer/smog"
)

type Compiler struct {
	Parser *Parser
}

func NewCompiler() *Compiler {
	c := &Compiler{}
	c.Parser = NewParser()
	return c
}

func (c *Compiler) CompileClassInFile(path string, file string, systemClass *smog.Class) *smog.Class {
	return c.Parser.CompileClass(path, file, systemClass)
}

func (c *Compiler) CompileClassInString(stmt string, systemClass *smog.Class) *smog.Class {
	return c.Parser.CompileClassString(stmt, systemClass)
}

// public static som.vmobjects.Class compileClass(String stmt, som.vmobjects.Class systemClass) {
// 	return new SourcecodeCompiler().compileClassString(stmt, systemClass);
// }

//	private som.vmobjects.Class compile(String path, String file, som.vmobjects.Class systemClass) throws IOException {
//		som.vmobjects.Class result = systemClass;
//		String fname = path + Universe.fileSeparator + file + ".som";
//		parser = new Parser(new FileReader(fname));
//		result = compile(systemClass);
//		som.vmobjects.Symbol cname = result.getName();
//		String cnameC = cname.getString();
//		if(file != cnameC)
//			throw new IllegalStateException("File name " + file + " does not match class name " + cnameC);
//		return result;
//	}
func (c *Compiler) Compile(path string, file string, systemClass *smog.Class) *smog.Class {
	result := systemClass
	fname := path + "/" + file + ".som"
	c.Parser = NewParser(io.FileReader(fname))
	result = c.Parser.Compile(systemClass)
	cname := result.GetName()
	cnameC := cname.GetString()
	if file != cnameC {
		panic("File name " + file + " does not match class name " + cnameC)
	}
	return result
}

//	private som.vmobjects.Class compileClassString(String stream, som.vmobjects.Class systemClass) {
//		parser = new Parser(new StringReader(stream));
//		som.vmobjects.Class result = compile(systemClass);
//		return result;
//	}
func (c *Compiler) CompileClassString(stream string, systemClass *smog.Class) *smog.Class {
	c.Parser = NewParser(io.StringReader(stream))
	result := c.Parser.Compile(systemClass)
	return result
}

//	private som.vmobjects.Class compile(som.vmobjects.Class systemClass) {
//		ClassGenerationContext cgc = new ClassGenerationContext();
//		som.vmobjects.Class result = systemClass;
//		parser.classdef(cgc);
//		if(systemClass == null)
//			result = cgc.assemble();
//		else
//			cgc.assembleSystemClass(result);
//		return result;
//	}
func (c *Compiler) Compile(systemClass *smog.Class) *smog.Class {
	cgc := NewClassGenerator()
	result := systemClass
	c.Parser.Classdef(cgc)
	if systemClass == nil {
		result = cgc.Assemble()
	} else {
		cgc.AssembleSystemClass(result)
	}
	return result
}
