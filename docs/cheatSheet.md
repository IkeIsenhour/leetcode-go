## 1. TWO SUM (Hash Map)

Uses a Hash Map to track encountered numbers and their index; we call this map "complements".
As we iterate through the remaining list, we subtract each number from the target to find it's
"complementary number" (i.e. the number that when added to our number sums to the target).
We check our complements map to see if the complementary number already exists. If so, we return the two
indices. Otherwise, we add it to the map.

## 2. ADD TWO NUMBERS (Basic Carry Arithmetic and Remainder, Linked List traversal)

There's two important pieces to solving this optimally: first, traversing both linked lists at the
same time; second,performing math with the carry and remainder.
To iterate through the lists at the same time, just keep looping as long as both nodes are not null
and the carry isn't 0. Remember this is an OR not AND, if any of those things are not true then keep going.
Once a node is null, just return 0 for it. Otherwise, it's in-place addition where you carry over the carry and just
keep creating a new node with the two values + the carry. The leftover carry, if any, goes to the next iteration.
With this solution, you iterate two lists and create the new one all in place.

## 3. LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS (Sliding Window, Hash Set)

Uses a sliding window algorithm to focus on a subset of characters and a Hash Set to track what characters
have already been seen.The key here is to recognize that you don't need to initialize right to 0, you can just
use right as the iterator of the for loop, and then do a while loop within to remove characters until the next
right character is valid again.
At the end of the day, this is really just adding/removing from a set and incrementing indices accordingly.

## 4. MEDIAN OF TWO SORTED ARRAYS

## 5. LONGEST PALINDROMIC SUBSTRING

## 6.ZIGZAG CONVERSIONS (2-D Matrix, Top-Level Index Tracking w/ Directional Tracker)

Uses a 2-D matrix to create each sub-set of the string. Using a directional tracker that
oscillates between -1 and 1 depending on if we've hit the bottom. Bottom is determined by the
length of the number of rows. The key here is that we're no moving within each of the inner lists, but
just within the top-level list (matrix). Also, don't focus too much on the "zigzag" pattern,
at the end of the day it's more of a row based combination of the characters. The zigzag is just a
visual representation of how the rows are defined, but you don't need to focus on horizontal
spacing like that.

## 217. CONTAINS DUPLICATE (HashSet)

Use a set to track what as been seen. Return true if current number in already in seen; otherwise add it there.
If you make it through the whole list without having returned true then return false

## 268. Missing Number (Math)

Use the "n-th Triangular Number" formula to calculate what number is missing in a list of numbers. The trick here is to
realize that when given a list of consecutive natural number, if you can find the expected total then you can just subtract
from that to find the missing numbers. This is possible to do in O(n) time by using the formula mentioned earlier to get
the expected total (the total in the case that every number was present). From there it is a simple for loop with subtraction
to find the missing number.

## 448. Find All Numbers Disappeared in an Array

In this problem, you need to utilize the key fact that the numbers will be in the range of 1 to n (the length of the list). The brute
force solution is just to make a set out of the given numbers, then loop through each number in the range checking if it exists in the set
and adding to a return list if it does not exist in the set.

The more optimized approach is to avoid using extra memory by modifying the original list in place (extra memory does not prohibit making a
separate return list). In order to accomplish this, you just need to make numbers at an index negative and then use the absolute value
for the number at each index to see what number it originally was. Finally, you can loop back through the list one more time to see what is
not negative and build a return list out of that.
