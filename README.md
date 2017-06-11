# Clockwise Matrix walk

## Assumptions
* For this project I am assuming the matrix will be consisting of int slices of all equal lengths. 

## Time
* I spent approxiamately 2 hours on this problem. 

## Algorithm
I tried a few different algorithms to solve this problem. I did not want a time complexity any slower than O(n). The current algorithm will walk through the multidimensional array using two functions `getTopDown` and `getBottomUp`. 

* getTopDown

The first function is run is getTopDown which will traverse the top of the slice from left to right appending each int found to the result string then removing this slice from the multidimonsional array. Next it will travere down the last index of each of the slices in the multidimensional array adding each found int to the result string and removing that index from the slice. 

Example:
```
1 2 3 4  ->  
x x x 5  ->  x x x 
x x x 6  ->  x x x 

result = "1 , 2 , 3, 4, 5, 6"
```

* getBottomUp

The second function is run is getBottomUp which will traverse the bottom of the slice from right to left appending each int found to the result string then removing this slice from the multidimonsional array. Next it will travere up the first index of each of the slices in the multidimensional array adding each found int to the result string and removing that index from the slice. 

Example:
```
6 x x x  ->
5 x x x  ->  x x x   
4 3 2 1  ->  x x x   

result = "1 , 2 , 3, 4, 5, 6"
```

This algorithm breaks the square down for each traversal removing '⅂' and 'L' until only a single line remains. If it finds that it is the last slice in the pattern it quits at the end of its walk.