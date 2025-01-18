# Timing Binary Search

As you may know, binary search is a search algorithm that works on sorted arrays and has a time complexity of O(logn) which is much faster than O(n). 

You may notice when timing binary search that the time is much faster than it seems to take.
For example if may say 24 microseconds for a random input of size 1,000,000,000 but it seems to take like 10-20 seconds to receive the time.

This is because when generating the random input, it takes O(n) to create an input array of size n. So much of the time taken to run the timing test is taken up by the random input generating. 

Just note that the time it returns is the accurate time it took to run the function, but not necessarily the total time you'll be waiting for results on the test
