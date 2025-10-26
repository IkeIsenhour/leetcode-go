# Journal

## Friday, October 24, 2025

Tonight I did 3 binary tree questions and one hash set question. In order, these were: `Maximum Depth of Binary Tree`, `Same Tree`,
`Invert Binary Tree`, and `Happy Number`.

The first 3 were generate binary tree question, all of which used Depth First Search as the solution. The learning from these questions
was that DFS implemented via recursion can take you through many of the basic binary tree questions. It's an efficient way
to traverse a branch from top to bottom, and often is simply implemented by the same structure of code each time:

[BASE CASE]
This is usually a check if the current node is null, but could be anything. Usually some sort of return, whether
that be nothing, 0 or something else. You might also need a return for the positive case here as well, although sometimes that comes
later in the function. The ordering of your returns matters, it can effect if you hit a nil pointer issue or not.

[RECURSIVELY CALL THE FUNC FOR THE NEW LEFT AND RIGHT]
You will call whatever the recursive func is here for the left and right nodes of the current node, continuing all the way through the branch. There
are some slight variations on this, specifically regarding what gets passed into the func. `Same Tree` is a good example, it requires you
to pass in two different trees nodes at the same time, as opposed to left and right together. Instead left goes with left and right goes with right.

[IMPORTANT LOGIC]
Whatever needs to be done is done here. Whether that is a simple calculation, changing values of nodes, etc. Handle that information and then
return the positive case if that was handled already in the base case section

After those 3 questions, I decided to do a different topic for fun and ended up selecting `Happy Number`, which apparently is a
very popular Google question. This question is a bit of a math puzzle, although not really when you think about it. It requires a
hash set to solve efficiently.

Essentially, you need to realize what your base cases are to know if the number is happy or not. They give you the happy case: "if
the number eventually resolved to 1 then it is happy." But, this leaves you wondering: "how do I know when my number will never resolve
to 1? I need to avoid an endless loop". This is where it's a bit of a math problem, you need to realize that if you encounter the same number
twice, then you're in an endless loop and can kill the process. This really isn't a math problem though, because it's pretty straight forward.
If this number has been encountered before, that means you have already done all the calculations that come after it and you ended up at this
same number again so you have already proven this path will never resolve to 1.

After that realization, it is a pretty straightforward implementation of the business logic. Store the numbers you've seen in a hash set,
follow their rules to get to the next number, store that number and repeat. Check if you've hit your base cases and eventually return true
or false.

## Saturday, October 25, 2025

Today I worked on one problem: `Search Insert Position`. This problem involves searching through a sorted array for a target value. If
you find the value, return it's index. If you don't find it, return the index it should be at.

The big giveaway on this problem is firstly that the array is sorted and secondly that they want a solution in O(log n) time. Both
these are strong indicators that the problem is looking for a binary search solution, which is exactly what is implemented. A typical
binary search implementation is defining a low index and a high index. Add those two indexes together and divide by two to the middle
index of those two numbers. Check if target is equal to the middle index and keep dong this loop until you find the target.

The only unique thing about this problem is that it also wants you to return where a number would be if it is not in the list. You have to either
add our subtract 1 to the final result you get for middle in the case that you don't find the target, but that's about it. Pretty straightforward.
