---
layout: "../../layouts/LayoutSingle.astro"
title: "Doc2"
---

Here's the content from the slides transposed into markdown format with LaTeX for mathematical symbols and code snippets:

# Heap Operations

## SiftUp

Move an item up to its correct position by repeatedly swapping with its parent.

```python
def SiftUp(H, i):
    parent = (i - 1) // 2
    if (i > 0) and (H[parent].key < H[i].key):
        H[i], H[parent] = H[parent], H[i]
        SiftUp(H, parent)
```

Analysis: At most 1 comparison at each level, so total time is $O(\log n)$

## SiftDown

Move an item down to its correct position by swapping with the larger child.

```python
def SiftDown(H, i):
    n = H.size  # number of items in heap
    left = 2*i + 1
    right = 2*i + 2
    if (right < n) and (H[right].key > H[left].key):
        largerChild = right
    else:
        largerChild = left
    if (largerChild < n) and (H[i].key < H[largerChild].key):
        H[i], H[largerChild] = H[largerChild], H[i]
        SiftDown(H, largerChild)
```

Analysis: At most 2 comparisons at each level, so total time is $O(\log n)$

## Insert

```python
def Insert(H, x):
    H.size = H.size + 1
    k = H.size - 1  # index of last position
    H[k] = x  # insert x in last position
    SiftUp(H, k)
```

Analysis: SiftUp time dominates, so total time is $O(\log n)$

## Delete

```python
def Delete(H, i):
    k = H.size - 1  # index of last position
    H[i] = H[k]  # overwrite item being deleted with element in last position
    H.size = H.size - 1  # decrement number of items
    SiftUp(H, i)  # either SiftUp or SiftDown will do nothing
    SiftDown(H, i)
```

Analysis: SiftUp/SiftDown time dominates, so total time is $O(\log n)$

## ExtractMax

```python
def ExtractMax(H):
    x = H[0]
    Delete(H, 0)
    return x
```

Analysis: Delete time dominates, so total time is $O(\log n)$

# Heap Construction

## Heapify

```python
def Heapify(H, n):
    for i in range(n-1, -1, -1):
        SiftDown(H, i)
```

Correctness: After SiftDown(H,i) is executed, subtree rooted at node i satisfies heap invariant.

Running time: Heapify runs in $O(n)$ time.

Proof that Heapify runs in $O(n)$ time:
Suppose the tree has $n$ nodes and $d$ levels (so $2^d < n < 2^{d+1}$).

- If node $i$ is at level $j$, SiftDown(H, i) needs $\leq 2(d - j)$ comparisons.
- There are at most $2^j$ nodes at level $j$.
- Total number of comparisons is no more than:

$\sum_{j=0}^d 2^j \cdot 2(d-j) = 2d \cdot 2^{d+1} - 2 \cdot (d-1)2^{d+1} + 2 = 4 \cdot 2^d - 2d - 4 < 4n = O(n)$

# Heapsort

```python
def Heapsort(A, n):
    Heapify(A, n)  # form max heap using array A
    for k in range(n-1, 0, -1):
        A[k] = ExtractMax(A)
```

# Radix Sort

Useful for sorting multi-field keys in lexicographic order.

```python
def RadixSort(A, n):
    for field in range(rightmost, leftmost):  # least significant to most significant
        # sort A on field using a stable sort
<!--
This file contains notes for CS161, specifically focusing on the Radix Sort algorithm.

Radix Sort is a non-comparative sorting algorithm that sorts integers by processing individual digits. It processes digits from the least significant digit (LSD) to the most significant digit (MSD) or vice versa.

Time Complexity:
- Best Case: O(nk)
- Average Case: O(nk)
- Worst Case: O(nk)
  where n is the number of elements and k is the number of digits in the largest number.

Space Complexity:
- O(n + k)
  where n is the number of elements and k is the number of digits in the largest number.
-->
```

# Bucket Sort

```python
def BucketSort(A, n):
    for i in range(n):
        insert A[i] into bucket[key(A[i])]
    for i in range(b):
        sort bucket[i]
    combine the buckets
```

Running time:

- Worst case: $O(n^2)$
- Best case: $O(n)$
- Average case: $O(n)$ if certain assumptions are satisfied

Storage: $O(n + b)$

Based on the content provided, here is a continuation of the discussion on external sorting and replacement selection:
Based on the provided document excerpts, here's a summary of the variables used in external sorting:

1. n: The total number of records (items) in the file to be sorted.

2. m: The number of records that can fit in memory at once (m < n).

3. f: The number of input files that can be open at once.

4. x: Individual records or keys being sorted. These are the actual data elements being manipulated during the sorting process.

5. b: The base of the number system used (in radix sort).

6. d: The number of fields or digits being sorted (in radix sort).

In the context of replacement selection:

7. The current key being written out to a run.
8. The next key being read in.

For the polyphase merge sort:

9. Run: A sorted group of records.

10. Phase: The stages of the sorting process (initial phase and subsequent merge phases).

These variables are used throughout the external sorting process to manage the data, control the flow of the algorithm, and determine the efficiency of the sorting operation. The relationship between these variables (particularly n, m, and f) determines how the external sort will proceed and how many phases will be required to complete the sort.

# Replacement Selection

Replacement selection is a technique used to improve external sorting by creating longer initial runs. The key points are:

1. When a key is written out to a run, the next key is immediately read in.

2. If the new key is greater than the last key written, it becomes part of the current run.

3. If the new key is less than the last key written, it is saved for the next run.

This allows runs to potentially grow longer than the buffer size.

## Example

The document shows an example of applying replacement selection to sort the following numbers:

145 354 507 590 875 929 481 47 208 212 929 902 124 250 11 386 281 680 109 100 542 64 508 654 793 538 322 299 686 104 989 465 777 991 931 677 176 230 214 369 106 218 724 779 565 559 873 696 726 326 415 761 915

Using replacement selection resulted in 7 initial runs:

Run 1: 145 354 507 590 875
Run 2: 929 47 208 212 481 902 929
Run 3: 11 124 250 281 386 542 680
Run 4: 64 100 109 508 538 654 686 793 989
Run 5: 104 299 322 465 677 777 931 991
Run 6: 176 214 218 230 369 565 724 779 873
Run 7: 106 326 415 559 696 726 761 915

This is compared to 14 runs that would have been created without replacement selection.

## Benefits

Replacement selection produces longer initial runs compared to simple buffering. This reduces the number of runs that need to be merged in subsequent phases, potentially improving the overall efficiency of external sorting.

The document then goes on to show how these runs are merged to produce the final sorted output, demonstrating the complete external sorting process using replacement selection for the initial run creation.
