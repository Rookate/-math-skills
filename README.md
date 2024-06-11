This is a project about calculating the following
===================================================

> For run this program you need to use the following command : 

```sh 
go run main/main.go data.txt
```

### Average : 
---
The sum of the numbers divided by how many numbers are in the list.

For example, the mean average of the numbers 2, 3, 4, 7, and 9 (summing to 25) is 5.

### Median : 
---
The median is the value in the middle of a data set, meaning that 50% of data points have a value smaller or equal to the median and 50% of data points have a value higher or equal to the median.

#### For an odd data set : 

1. **Sort the data**: Arrange the values of the data set from smallest to largest.

2. **Find the central index**: Calculate the central index of your data :

    central_index=(n+1)/2, n is the total number of values

3. **Identify the median**: The median is the value located at the central index.

    _For exemple : 1, 3, 5, 11, 12 : (5+1) / 2 = 3. The median is the value located at the central index, so it's 5_

    The median is the value located at the central index, so it's 5.

#### For an even data set:

1. **Sort the data**: Arrange the values of the data set from smallest to largest.

2. **Find the two central indices**: Calculate the central indices of your data set using the following formulas:

 - central_index_1 = n/2
 - central_index_2 = (n/2) + 1, n is the total number of values.

3. **Identify the two central values**:  
- The two central values are those located at the central indices *central_index_1* and *central_index_2*

4. **Calculate the median**: The median is the average of the two central values.

_For exemple : 1, 3, 5, 11 : 4/2 = 2 and (4/2) + 1 = 3. (3+5) / 2 = 4. The median is 4_

### Variance :
---

#### Simple Definition:

Variance is a measure that tells us how spread out the values in a data set are from the average (mean). In other words, it shows whether the values are close to the mean or widely scattered around it.

#### Explanation with an Image:

Imagine you are throwing darts at a target. If all the darts are very close to the bullseye (the mean), the variance is low. If the darts are spread all over the target, the variance is high.

#### Why It's Useful:

- **Low variance**: The values are close to the mean, indicating little variation.
- **High variance**: The values are spread out around the mean, indicating a lot of variability.

#### How to Calculate It Simply:

1. **Find the mean**: Add up all the values and divide by the total number of values.
2. **Calculate the deviations**: Subtract the mean from each value (this gives the deviation from the mean).
3. **Square the deviations**: Multiply each deviation by itself (to make all deviations positive).
4. **Find the average of the squared deviations**: Add up all the squared deviations and divide by the total number of values.

#### Simple Example:

For the values 2, 4, 4, 6, 8:

1. **Mean**: (2 + 4 + 4 + 6 + 8) / 5 = 4.8
2. **Deviations**: 2-4.8, 4-4.8, 4-4.8, 6-4.8, 8-4.8 = -2.8, -0.8, -0.8, 1.2, 3.2
3. **Squared deviations**: (-2.8)^2, (-0.8)^2, (-0.8)^2, (1.2)^2, (3.2)^2 = 7.84, 0.64, 0.64, 1.44, 10.24
4. **Average of squared deviations**: (7.84 + 0.64 + 0.64 + 1.44 + 10.24) / 5 = 4.16

So, the variance is 4.16, which means that the values in this set are, on average, dispersed 4.16 squared units around the mean.

### Standard Deviation:
---
The standard deviation is a measure of the amount of variation or dispersion in a set of values. It indicates how much the values in a data set tend to deviate from the mean (average) value of the data set.

#### Why It's Useful:
- **Low standard deviation**: Indicates that the values tend to be close to the mean.
- **High standard deviation**: Indicates that the values are spread out over a wider range.

#### How to Calculate Standard Deviation:

1. **Find the mean**: Add up all the values and divide by the total number of values.
2. **Calculate the deviations**: Subtract the mean from each value (this gives the deviation from the mean).
3. **Square the deviations**: Multiply each deviation by itself (to make all deviations positive).
4. **Find the average of the squared deviations**: This is known as the variance. Add up all the squared deviations and divide by the total number of values.
5. **Take the square root**: The standard deviation is the square root of the variance.

_For exemple if we take the previous result "4.16" we just need to square root the result._

Square root of 4 is 2

So, the standard deviation is 2, indicating that the values are, on average, 2 units away from the mean