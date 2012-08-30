package com.feliperibeiro;

public class InsertionSort {

    /**
     * In-place Insertion Sort
     */
    public void sort(int[] items) {
        for (int i = 1; i < items.length; i++) {
            int n = items[i];
            int pos = i;
            while (pos > 0 && items[pos-1] > n) {
                items[pos] = items[--pos];
            }
            items[pos] = n;
        }
    }

}

