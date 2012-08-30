package com.feliperibeiro;
import java.util.Arrays;

public class MaximumSubarray {

    public static int[] max(int[] a) {
        int maxSum = 0,
            currentSum = 0,
            maxSumStart = 0,
            maxSumEnd = 0,
            currentSumStart = 0;
        for (int currentSumEnd = 0; currentSumEnd < a.length; currentSumEnd++) {
            currentSum += a[currentSumEnd];
            if (currentSum > maxSum) {
                maxSum = currentSum;
                maxSumEnd = currentSumEnd;
                maxSumStart = currentSumStart;
            }
            
            if (currentSum < 0) {
                currentSumStart = currentSumEnd + 1;
                currentSum = 0;
            }
        }

        return Arrays.copyOfRange(a, maxSumStart, maxSumEnd+1);
    }
}
