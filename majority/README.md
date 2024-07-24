1. Take a variable lastElement and set count = 0.
2. Iterate over the array.
3. For each element, check the value of count:
   
   		a. If count <= 0, update lastElement with the current array value.
   
   		b. If count > 0, compare lastElement with the current array value
   			b.1. If lastElement is not equal to the current array value, decrease count.
     		b.2. Otherwise, increase count.
5. Finally, return lastElement.
