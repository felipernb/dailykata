import random

def quicksort(a):
    if len(a) <= 1:
        return a

    p = partition(a)
    sorted_list = quicksort(a[0:p])
    sorted_list.append(a[p])
    sorted_list.extend(quicksort(a[p+1:]))
    return sorted_list

def partition(a):
    p = int(random.random() * len(a))
    a[0], a[p] = a[p], a[0]
    pivot = a[0]
    i = 1
    for j in xrange(i, len(a)):
        if a[j] < pivot:
            a[j], a[i] = a[i], a[j]
            i += 1
    a[0], a[i-1] = a[i-1], a[0]
    return i-1

print quicksort([5,8,2,-1,10,0,3])
