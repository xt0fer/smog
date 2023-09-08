package smog

type Primitive struct {
	Object
	signature               *Symbol
	holder                  *Object
	primFields              []interface{}
	NumberOfPrimitiveFields int
}

func (p *Primitive) isPrimitive() bool {
	return true
}

func NewPrimitive(signatureString string) {
	np := &Primitive{}
	np.signature = GetUniverse().SymbolFor(signatureString)

	np.setClass(GetUniverse().PrimitiveClass)
}

func (p *Primitive) GetSignature() *Symbol {
	return p.signature
}
func (p *Primitive) SetSignature(value *Symbol) {
	p.signature = value
}

func (p *Primitive) GetHolder() *Object {
	return p.holder
}
func (p *Primitive) SetHolder(value *Object) {
	p.holder = value
}

func (p *Primitive) GetDefaultNumberOfFields() int {
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
