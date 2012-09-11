def insertion_sort(a):
    for i in xrange(1, len(a)):
        n = a[i]
        pos = i
        while pos > 0 and a[pos-1] > n:
            a[pos] = a[pos-1]
            pos -= 1
        a[pos] = n
    return a
