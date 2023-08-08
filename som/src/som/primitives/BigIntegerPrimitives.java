/**
 * Copyright (c) 2009 Michael Haupt, michael.haupt@hpi.uni-potsdam.de
 * Software Architecture Group, Hasso Plattner Institute, Potsdam, Germany
 * http://www.hpi.uni-potsdam.de/swa/
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package som.primitives;

import som.vm.Universe;
import som.vmobjects.BigInteger;
import som.vmobjects.Frame;
import som.vmobjects.Integer;
import som.vmobjects.Object;
import som.vmobjects.Primitive;

public class BigIntegerPrimitives extends Primitives 
{   
    public void installPrimitives() 
    {
	installInstancePrimitive
	    (new Primitive("asString")
		{
		    public void invoke(Frame frame)
		    {
			BigInteger self = (BigInteger) frame.pop();
			frame.push(Universe.newString(self.getEmbeddedBiginteger().toString()));
		    }
		}
	     );
	
	installInstancePrimitive(
			new Primitive("sqrt") {
				public void invoke(Frame frame) {
					BigInteger self = (BigInteger) frame.pop();
					frame.push(Universe.newDouble(Math.sqrt(self.getEmbeddedBiginteger().doubleValue())));
				}
			}
		);
    
	installInstancePrimitive
	    (new Primitive("+")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation and perform conversion to Integer if required
			java.math.BigInteger result = left.getEmbeddedBiginteger().add(right.getEmbeddedBiginteger());
			if(result.bitLength() > 31)
				frame.push(Universe.newBigInteger(result));
			else
				frame.push(Universe.newInteger(result.intValue()));
		    }
		}
	     );
    
	installInstancePrimitive
	    (new Primitive("-")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation and perform conversion to Integer if required
			java.math.BigInteger result = left.getEmbeddedBiginteger().subtract(right.getEmbeddedBiginteger());
			if(result.bitLength() > 31)
				frame.push(Universe.newBigInteger(result));
			else
				frame.push(Universe.newInteger(result.intValue()));
		    }
		}
	     );
    
	installInstancePrimitive
	    (new Primitive("*")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation and perform conversion to Integer if required
			java.math.BigInteger result = left.getEmbeddedBiginteger().multiply(right.getEmbeddedBiginteger());
			if(result.bitLength() > 31)
				frame.push(Universe.newBigInteger(result));
			else
				frame.push(Universe.newInteger(result.intValue()));
		    }
		}
	     );
    
	installInstancePrimitive
	    (new Primitive("/")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation and perform conversion to Integer if required
			java.math.BigInteger result = left.getEmbeddedBiginteger().divide(right.getEmbeddedBiginteger());
			if(result.bitLength() > 31)
				frame.push(Universe.newBigInteger(result));
			else
				frame.push(Universe.newInteger(result.intValue()));
		    }
		}
	     );

	installInstancePrimitive
	    (new Primitive("%")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation:
			frame.push(Universe.newBigInteger(left.getEmbeddedBiginteger().mod(right.getEmbeddedBiginteger())));
		    }
		}
	     );

	installInstancePrimitive
	    (new Primitive("&")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation:
			frame.push(Universe.newBigInteger(left.getEmbeddedBiginteger().and(right.getEmbeddedBiginteger())));
		    }
		}
	     );
    
	installInstancePrimitive
	    (new Primitive("=")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation:
			if (left.getEmbeddedBiginteger().compareTo(right.getEmbeddedBiginteger()) == 0)
			    frame.push(Universe.trueObject);
			else
			    frame.push(Universe.falseObject);
		    }
		}
	     );
    
	installInstancePrimitive
	    (new Primitive("<")
		{
		    public void invoke(Frame frame)
		    {
			Object rightObj  = frame.pop();
			BigInteger right = null;
			BigInteger left  = (BigInteger) frame.pop();
			
			// Check second parameter type:
			if (rightObj instanceof Integer) {
			    // Second operand was Integer
			    right = Universe.newBigInteger((long)((Integer)rightObj).getEmbeddedInteger());
			}
			else
			    right = (BigInteger)rightObj;
			
			// Do operation:
			if (left.getEmbeddedBiginteger().compareTo(right.getEmbeddedBiginteger()) < 0)
			    frame.push(Universe.trueObject);
			else
			    frame.push(Universe.falseObject);
		    }
		}
	     );
    }
}
