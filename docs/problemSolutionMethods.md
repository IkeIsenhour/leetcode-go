# Leetcode Solution Types and Corresponding Questions

## Two Pointer

### Description

### Problems

1. [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

## Recursion

### Description

### Problems

1. [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)

## HashMap/HashSet

### Description

Usually involves a situation in which you have to associate one thing with another thing. HashMap solutions
tend to be pretty intuitive, and can often be used in the brute force form of a solution. Can also be the optimal solution, like we see with
the word pattern problem. The key with HashMaps is the constant lookup time, making them a much more efficient option than other collection
types for a key:value lookup.

### Problems

1. [Ransom Note](https://leetcode.com/problems/ransom-note/description/)
2. [Word Pattern](https://leetcode.com/problems/word-pattern/description/)
3. [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/description/)
4. [Happy Number](https://leetcode.com/problems/happy-number/description/)

## Sliding Window

### Description

Left and Right boundaries, moving each separately to "slide" the window. Need to make sure that the window is always valid. Typically that means the left and
right boundaries are withing the bounds of the array/string that they are iterating. Equally, whatever else "valid" means for that question
(example: cannot have duplicates in string). When the window is valid, we increment "right". When the window is invalid, we increment "left".
The length of this window can always be calculated by "(right - left) + 1"

### Problems

1. [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/) uses the sliding
   window technique with a HashSet/HashMap for it's optimal solution. First, we begin with left and right at index 0. As we loop through the string, we check
   if the window is valid. If valid, we add the character at "right" to the HashSet, calculate the length, check if the length is greater than longest, and
   increment right. If invalid, we remove "left" from the HashMap and increment "left", repeating this until "right" is not in the map anymore.

## Running Count Algorithm

### Description

### Problems

1. [Majority Element](https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150)

## Traversing Linked List

### Description

Really just understanding how to traverse a linked list. Often requires keeping track of the current and/or previous while we set the next or val properties
equal to some temprary value that we will throw away on the next loop. Traversing specifically comes down to looping over the linked list until next is null.

### Problems

1. [Majority Element](https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150)

## Depth First Search

### Description

An algorithm used to search a Binary Tree or Graph. As opposed to Breadth First Search, DFS explores as far as possible on a branch before
backtracking. It is typically accomplished with recursion or a stack. It is useful for exploring all paths.

### Problems

1. [Maximum Depth of a Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/)
2. [Same Tree](https://leetcode.com/problems/same-tree/description/)
3. [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/)

## Breadth First Search

### Description

An algorithm used to search a Binary Tree or Graph. As opposed to Depth First Search, BFS explores all neighbors at the current depth before
moving onto the next depth. It is typically accomplished using a queue. It is useful for finding the shortest path.

### Problems

1. []()
