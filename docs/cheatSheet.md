1. TWO SUM (Hash Map)
   Uses a Hash Map to track encountered numbers and their index; we call this map "complements".
   As we iterate through the remaining list, we subtract each number from the target to find it's
   "complementary number" (i.e. the number that when added to our number sums to the target).
   We check our complements map to see if the complementary number already exists. If so, we return the two
   indices. Otherwise, we add it to the map.

2. ADD TWO NUMBERS (Basic Carry Arithmetic and Remainder, Linked List traversal)
   There's two important pieces to solving this optimally: first, traversing both linked lists at the
   same time; second,performing math with the carry and remainder.
   To iterate through the lists at the same time, just keep looping as long as both nodes are not null
   and the carry isn't 0. Remember this is an OR not AND, if any of those things are not true then keep going.
   Once a node is null, just return 0 for it. Otherwise, it's in-place addition where you carry over the carry and just
   keep creating a new node with the two values + the carry. The leftover carry, if any, goes to the next iteration.
   With this solution, you iterate two lists and create the new one all in place.

3. LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS (Sliding Window, Hash Set)
   Uses a sliding window algorithm to focus on a subset of characters and a Hash Set to track what characters
   have already been seen.The key here is to recognize that you don't need to initialize right to 0, you can just
   use right as the iterator of the for loop, and then do a while loop within to remove characters until the next
   right character is valid again.
   At the end of the day, this is really just adding/removing from a set and incrementing indices accordingly.
