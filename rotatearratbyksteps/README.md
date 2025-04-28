Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

 

**Example 1:**

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

**Example 2:**

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

**Follow up:**

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?

**Solution**

### Solution 1: Reverse three times

We can assume the length of the array is $n$ and calculate the actual number of steps needed by taking the module of $k$ and $n$, which is $k \bmod n$.

Next, let us reverse three times to get the final result:

1. Reverse the entire array.
2. Reverse the first $k$ elements.
3. Reverse the last $n - k$ elements.

For example, for the array $[1, 2, 3, 4, 5, 6, 7]$, $k = 3$, $n = 7$, $k \bmod n = 3$.

1. In the first reverse, reverse the entire array. We get $[7, 6, 5, 4, 3, 2, 1]$.
2. In the second reverse, reverse the first $k$ elements. We get $[5, 6, 7, 4, 3, 2, 1]$.
3. In the third reverse, reverse the last $n - k$ elements. We get $[5, 6, 7, 1, 2, 3, 4]$, which is the final result.

The time complexity is $O(n)$, where $n$ is the length of the array. The space complexity is $O(1)$.

