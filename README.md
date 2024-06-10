This is a project about calculating the following
===================================================

### Average : 
The sum of the numbers divided by how many numbers are in the list.

For example, that mean average of the numbers 2, 3, 4, 7, and 9 (summing to 25) is 5.

### Median : 

_The median is the value in the middle of a data set, meaning that 50% of data points have a value smaller or equal to the median and 50% of data points have a value higher or equal to the median._

 #### For an odd data set : 

1. Sort data : Arrange the values of the data set from smallest to largest.

2. Find the central index : Calculate the central index of your data set using the following formula : 

    central_index=(n+1)/2, n is the total number of values

3. Identify the median: The median is the value located at the central index.


_For exemple : 1, 3, 5, 11, 12 : (5+1) / 2 = 3. The median is the value located at the central index, so it's 5_

#### For an even data set :

1. Sort data : Arrange the values of the data set from smallest to largest.

2. Find the two central index : Calculate the central index of your data set using the following fomula :

 - central_index_1 = n/2
 - central_index_2 = (n/2) + 1, n is the total number of values.

3. Identify the two central values : 

 - The two central values are those located at the central indices *central_index_1* and *central_index_2*

4. Calculate the median : The median is the average of the two central values.

_For exemple : 1, 3, 5, 11 : 4/2 = 2 and (4/2) + 1 = 3. (3+5) / 2 = 4. The median is 4_