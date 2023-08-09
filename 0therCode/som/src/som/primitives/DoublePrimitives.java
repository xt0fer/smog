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
import som.vmobjects.Double;
import som.vmobjects.Frame;
import som.vmobjects.Integer;
import som.vmobjects.Object;
import som.vmobjects.Primitive;

public class DoublePrimitives extends Primitives 
{
    private static Double coerceToDouble(Object o) {
        if(o instanceof Double)
            return (Double)o;
        if(o instanceof Integer)
            return Universe.newDouble((double)((Integer)o).getEmbeddedInteger());
        throw new ClassCastException("Cannot coerce to Double!");
    }
    
  public void installPrimitives() 
  {
    installInstancePrimitive
      (new Primitive("asString")
        {
          public void invoke(Frame frame)
          {
            Double self = (Double) frame.pop();
            frame.push(Universe.newString(java.lang.Double.toString(self.getEmbeddedDouble())));
          }
        }
       );
    
    installInstancePrimitive(
    			new Primitive("sqrt") {
    				public void invoke(Frame frame) {
    					Double self = (Double) frame.pop();
    					frame.push(Universe.newDouble(Math.sqrt(self.getEmbeddedDouble())));
    				}
    			}
    		);
    
    installInstancePrimitive
      (new Primitive("+")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            frame.push(Universe.newDouble(op1.getEmbeddedDouble() + op2.getEmbeddedDouble()) );
          }
        }
       );
    
    installInstancePrimitive
      (new Primitive("-")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            frame.push(Universe.newDouble(op2.getEmbeddedDouble() - op1.getEmbeddedDouble()));
          }
        }
       );
    
    installInstancePrimitive
      (new Primitive("*")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            frame.push(Universe.newDouble(op2.getEmbeddedDouble() * op1.getEmbeddedDouble()));
          }
        }
       );
    
    installInstancePrimitive
      (new Primitive("//")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            frame.push(Universe.newDouble(op2.getEmbeddedDouble() / op1.getEmbeddedDouble()));
          }
        }
       );

    installInstancePrimitive
      (new Primitive("%")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            frame.push(Universe.newDouble(op2.getEmbeddedDouble() % op1.getEmbeddedDouble()));
          }
        }
       );

    installInstancePrimitive
      (new Primitive("=")
        {
          public void invoke(Frame frame)
          {
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            if (op1.getEmbeddedDouble() == op2.getEmbeddedDouble())
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
            Double op1 = coerceToDouble(frame.pop());
            Double op2 = (Double) frame.pop();
            if (op2.getEmbeddedDouble() < op1.getEmbeddedDouble())
              frame.push(Universe.trueObject);
            else
              frame.push(Universe.falseObject);
          }
        }
       );
  }
}
