def quicksort(a, l = 0, r = a.length)
    return a if l >= r
    p = partition(a,l,r)
    quicksort(a, l, p)
    quicksort(a, p+1, r)
    a
end

def partition(a, l, r)
    # Random pivot
    p = rand(l..r-1)
    a[l],a[p] = a[p], a[l]
    pivot = a[l]
    i = l+1
    for j in i..r-1 do
        if a[j] < pivot
            a[j], a[i] = a[i], a[j]
            i += 1
        end
    end
    a[l], a[i-1] = a[i-1], a[l]
    i-1
end
a = [5,1,2,4,3,0,-1,10]
puts quicksort(a)
