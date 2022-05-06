# the algorithm of binary data coding in DNA 
## About task:
Recently I saw a video about biotechnology idea to store data in moss. [link here](https://www.twistbioscience.com/resources/video/how-does-dna-data-storage-work)
Well it sounded to me like a crazy idea. 
But one thing got my attention the way of coding binary data into chromosomes. I came up with idea that it is excellent exercise to develop programming skill, or even maybe task for job interview.

So the problem is: **we are given a integer number [uint8], and we need to convert its value to set of DNA's chromosomes.**

Here is my solution. Freely share your own concept of solving this task by folking my repository. 

## My solution:
### Map table of binary values to DNA chromosomes:
|bits pair|chromosome|
|-----|--------|
|00 | A|
|01 | G|
|10 | C|
|11 | T|

Example:
66 = (..0)(0)1|00|00|10 = AGAAC

### Algorithm:

step 1 : set shift variable x to its max (value = 8)

step 2: shift right data's bytes by x positions

step 3: **getting pair of bits** - use & operation on shifted data and value 3 (0b00000011)

step 4: search in map key value that is equal to result of "&" operation, get its value

step 5: use string builder to append given value as substring to string

step 6: set x = x - 2

step 8: go to step 2 until x is greater or equal 0

