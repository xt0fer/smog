package vm

type Primitive struct {
	Object
	SignatureIndex          int
	HolderIndex             int
	NumberOfPrimitiveFields int
}

//   // Static field indices and number of primitive fields
//   static final int signatureIndex            = 1 + classIndex;
//   static final int holderIndex               = 1 + signatureIndex;
//   static final int numberOfPrimitiveFields = 1 + holderIndex;

func (p *Primitive) isPrimitive() bool {
	return true
}

func NewPrimitive(signatureString string) {
	np := &Primitive{}
	np.SignatureIndex = 1 + np.ClassIndex
	np.HolderIndex = 1 + np.SignatureIndex
	np.NumberOfObjectFields = 1 + np.HolderIndex
	np.setClass(GetUniverse().PrimitiveClass)
	np.setSignature(GetUniverse().symbolFor(signatureString))
}

func (p *Primitive) getSignature() *Symbol {
	return p.Fields[p.SignatureIndex]
}
func (p *Primitive) setSignature(value *Symbol) {
	p.Fields[p.SignatureIndex] = Object(*value)
}

func (p *Primitive) getHolder() *Object {
	return p.getField(p.HolderIndex)
}
func (p *Primitive) setHolder(value *Object) {
	p.setField(p.HolderIndex, value)
}

func (p *Primitive) getDefaultNumberOfFields() int {
	return p.NumberOfPrimitiveFields
}

func (p *Primitive) isEmpty() bool {
	return false
}

//   public static Primitive getEmptyPrimitive(java.lang.String signatureString)
//   {
//     // Return an empty primitive with the given signature
//     return
//       (new Primitive(signatureString)
//         {
//           public void invoke(Frame frame)
//           {
//             // Write a warning to the screen
//             System.out.println("Warning: undefined primitive " + this.getSignature().getString() +
//                                " called");
//           }

//           public boolean isEmpty()
//           {
//             // The empty primitives are empty
//             return true;
//           }
//         }
//        );
//   }
// }
