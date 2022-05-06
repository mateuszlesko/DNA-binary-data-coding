# DNA data coding algorithm

## Map table of binary values to DNA chromosomes:
|bits pair|chromosome|
|-----|--------|
|00 | A|
|01 | G|
|10 | C|
|11 | T|

Example:
66 = (..0)(0)1|00|00|10 = AGAAC

## Algorithm:

step 1 : set shift variable x to its max (value = 8)

step 2: shift right data's bytes by x positions

step 3: **getting pair of bits** - use & operation on shifted data and value 3 (0b00000011)

step 4: search in map key value that is equal to result of "&" operation, get its value

step 5: use string builder to append given value as substring to string

step 6: set x = x - 2

step 8: go to step 2 until x is greater or equal 0

