import random

def quicksort(a, l, r):
    if l >= r:
        return a
    p = partition(a, l, r)
    quicksort(a, l, p)
    quicksort(a, p+1, r)
    return a

def partition(a, l, r):
    p = int(random.random() * (r-l))+l
    a[l], a[p] = a[p], a[l]
    pivot = a[l]
    i = l+1
    for j in xrange(i, r):
        if a[j] < pivot:
            a[j], a[i] = a[i], a[j]
            i += 1
    a[l], a[i-1] = a[i-1], a[l]
    return i-1

print quicksort([5,8,2,-1,10,0,3], 0, 7)
