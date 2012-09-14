def insertion_sort(a)
    for i in 1..a.length-1 do
        n = a[i]
        pos = i
        while pos > 0 and a[pos-1] > n do
            a[pos] = a[pos-1]
            pos -= 1
        end
        a[pos] = n
    end
    a
end
