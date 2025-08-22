# Pseudo Code for Group Anagrams

## Input = ["act", "pots", "tops", "cat", "stop", "hat"]

## Output = [["hat"], ["act", "cat"], ["stop", "pots", "tops"]]

Create a hashmap.
For each string inside strs, sort the string alphabetically, and push this to the hashmap as a key.
Each value will be an array. Then for each value, push the non-sorted string to the array, using the spread operator.
Finally, return the values of the hashmap as an array of arrays (2D array).

# Big O Analysis

```
Time complexity: O(m * n)
Space complexity: O(m * n)
```

Where m is the number of strings in the input array, and n is the maximum length of a string in the input array.
