- In any station between A and B, let’s say C. C will have gas left in our tank, if we go from A to that station. We can’t reach B from A with some gas(maybe 0) left in the tank in C, so we can’t reach B from C with an empty tank.

- If the gas is more than the cost in total, there must be some stations we have enough gas to go through them.


This Concludes:

If the sum of gas is more than the sum of cost, then there must be a solution. And the question guaranteed that the solution is unique(The first one you found will the right one).
The tank should never be negative, so restart whenever there is a negative number.

Solution Steps

Create a start to store the valid starting index from where the car could touch all the stations.
For each station i , fill the fuel tank with gas[i] and burn the fuel by cost[i]
If at any point the tank is < 0 then, choose the next index as the starting point.
At last, check if the total fuel available at the gas stations is greater than the total fuel burnt during the travel.
Return the start

RUN TEST

cd practice
go test -v ./gasstation/