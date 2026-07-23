1. TWO SUM (Hash Map)
   Uses a Hash Map to track encountered numbers and their index; we call this map "complements".
   As we iterate through the remaining list, we subtract each number from the target to find it's
   "complementary number" (i.e. the number that when added to our number sums to the target).
   We check our complements map to see if the complementary number already exists. If so, we return the two
   indices. Otherwise, we add it to the map.
