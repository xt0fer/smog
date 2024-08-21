package smog

import "log"

type Primitive struct {
	Object
	signature               *Symbol
	holder                  *Class
	primFields              []interface{}
	NumberOfPrimitiveFields int
}

func (p *Primitive) IsPrimitive() bool {
	return true
}

func NewPrimitive(signatureString string) *Primitive {
	np := &Primitive{}
	np.signature = GetUniverse().SymbolFor(signatureString)

	np.SetClass(GetUniverse().PrimitiveClass)
	return np
}

func (p *Primitive) GetSignature() *Symbol {
	return p.signature
}
func (p *Primitive) SetSignature(value *Symbol) {
	p.signature = value
}

func (p *Primitive) GetHolder() *Class {
	return p.holder
}
func (p *Primitive) SetHolder(value *Class) {
	p.holder = value
}

func (p *Primitive) GetDefaultNumberOfFields() int {
	return p.NumberOfPrimitiveFields
}

func (p *Primitive) IsEmpty() bool {
	return false
}

func (p *Primitive) Invoke(frame *Frame) {
	// Write a warning to the screen
	log.Println("Warning: undefined primitive " + p.signature.ToString() + " called")
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

//	        public boolean isEmpty()
//	        {
//	          // The empty primitives are empty
//	          return true;
//	        }
//	      }
//	     );
//	}
func GetEmptyPrimitive(signatureString string) *Primitive {
	log.Println("Warning: Get EMPTY primitive " + signatureString + " called")
	return NewPrimitive(signatureString)
}
