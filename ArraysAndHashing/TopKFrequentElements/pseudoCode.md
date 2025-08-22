# Pseudo Code for Top K Frequent Elements

## Input = nums = [1,2,2,3,3,3], k = 2

## Output = [2,3]

Create an empty hashmap which will measure the frequency of each number.
For each number in nums, if the num is not in the hashmap, add it as a key with a value 1. If it is, increase the frequency.
Sort the hashmap by value (frequency) in descending order.
Extract the keys of the first `k` items from the sorted list.
Return these keys as the result.
