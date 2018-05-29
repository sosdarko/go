# About Aglorithm

We want to "add" memebers of smaller ordered sequence to the larger ordered sequence, so we get ordered sequence that has both of them.

Things will happen in place of larger sequence, so larger sequence must be big enough to take smaller into itself.

### ALGORITHM
**Goal**:
* trying to minimize shifting of elements of bigger sequence, using merge-like algorithm

Let's call bigger seq. A, and smaller B

**Two main operations**:
* Moving elements of A to make room for appropriate elements of B
* Copying appropriate elements from B to A

Assuming that both seq. are sorted in icreasing order
* start with the last element in B - the biggest one -> j
* find where Bj should lay in A
* all bigger elements of A (if there are any) shift j position to right
     * if we gone through all elements of A, copy the rest of B into beggining of the A and break out
* place Bj into extended A to it's final position
* repeat for all elements of B (j>=0)

**Example**:

A = [1 3 5 7 10 0 0 0 0 0 0]\
B = [-3 -1 6 6 13 25]

Copying B[5] to A[10]:\
[-3 -1 6 6 13 **25**] ->
[1 3 5 7 10 0 0 0 0 0 **25**]

Copying B[4] to A[9]:\
[-3 -1 6 6 **13** 25] ->
[1 3 5 7 10 0 0 0 0 **13** 25]

Shifting A:  3 - 5  to  7 - 9:\
[1 3 5 **7 10** 0 0 0 0 13 25] ->
[1 3 5 _7 10_ 0 0 **7 10** 13 25]

Copying B[3] to A[6]:\
[-3 -1 6 **6** 13 25] ->
[1 3 5 7 10 0 **6** 7 10 13 25]

Copying B[2] to A[5]:\
[-3 -1 **6** 6 13 25] ->
[1 3 5 7 10 **6** 6 7 10 13 25]

Shifting A:  0 - 3  to  2 - 5:\
[**1 3 5** 7 10 6 6 7 10 13 25] ->
[_1 3_ **1 3 5**  6 6 7 10 13 25]

Copying B[:2] to A[:2]:\
[**-3 -1** 6 6 13 25] ->
[**-3 -1** 1 3 5 6 6 7 10 13 25]
