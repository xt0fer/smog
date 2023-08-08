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
import som.vmobjects.Array;
import som.vmobjects.Frame;
import som.vmobjects.Object;
import som.vmobjects.Primitive;

public class ObjectPrimitives extends Primitives 
{   
  public void installPrimitives() 
  {
    installInstancePrimitive
      (new Primitive("==")
        {
          public void invoke(Frame frame)
          {
            Object op1 = frame.pop();
            Object op2 = frame.pop();
            if (op1 == op2)
              frame.push(Universe.trueObject);
            else
              frame.push(Universe.falseObject);
          }
        }
       );
    installInstancePrimitive
      (new Primitive("hashcode")
	{
	  public void invoke(Frame frame)
	  {
	    Object self = frame.pop();
	    frame.push(Universe.newInteger(self.hashCode()));
	  }
	}
       );
    installInstancePrimitive(
            new Primitive("objectSize") {
                public void invoke(Frame frame) {
                    Object self = frame.pop();
                    int size = self.getNumberOfFields();
                    if(self instanceof Array)
                        size += ((Array)self).getNumberOfIndexableFields();
                    frame.push(Universe.newInteger(size));
                }
            }
        );
  }
}
