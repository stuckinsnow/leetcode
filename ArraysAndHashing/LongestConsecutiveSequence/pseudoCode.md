The secret to solving this with a hashmap is to assign the num as the key, not the value. That way we can look at index locations num-1 and num+1 in the map to see if there are sequences on either side. Then we combine and keep track of the longest one.

So to clarify --

Start with an empty hashmap
Set max, which will be the max sequence length, to 0
Loop through each number in the array

- If the number is not already in the map
  - check the left and right side of it by looking for num-1 and num+1 in the map. If they are not there, we treat as 0.
  - Add 1 to the left and right values to get the total sequence length for this number.
  - Set the map at the current number to the total sequence length.
  - Update the ends of the sequence, so we set map[num-left] and map[num+right] to the total sequence length.
  - Update max if the total sequence length is greater than max.

  Return max.
