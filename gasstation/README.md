**PROBLEM:**

Given two integer arrays A and B of size N. There are N gas stations along a circular route, where the amount of gas at station i is A[i].

You have a car with an unlimited gas tank and it costs B[i] of gas to travel from station i to its next station (i+1). You begin the journey with an empty tank at one of the gas stations.

Return the minimum starting gas station's index if you can travel around the circuit once, otherwise return -1.

You can only travel in one direction. i to i+1, i+2, ... n-1, 0, 1, 2.. Completing the circuit means starting at i and ending up at i again.

**Input:**

A = [1,2,3,4,5]

B = [3,4,5,1,2]

Output: 3

**SOLUTION:**

- In any station between A and B, let’s say C. C will have gas left in our tank, if we go from A to that station. We can’t reach B from A with some gas(maybe 0) left in the tank in C, so we can’t reach B from C with an empty tank.

- If the gas is more than the cost in total, there must be some stations we have enough gas to go through them.


This Concludes:

If the sum of gas is more than the sum of cost, then there must be a solution. And the question guaranteed that the solution is unique(The first one you found will the right one).
The tank should never be negative, so restart whenever there is a negative number.

**Solution Steps**

- Create a start to store the valid starting index from where the car could touch all the stations.
- For each station i , fill the fuel tank with gas[i] and burn the fuel by cost[i]
If at any point the tank is < 0 then, choose the next index as the starting point.
At last, check if the total fuel available at the gas stations is greater than the total fuel burnt during the travel.
Return the start

**RUN TEST**

cd practice

go test -v ./gasstation/
