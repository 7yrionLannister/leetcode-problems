import java.util.Arrays;

public class TwoSumConstantSpace {
    public static void main(String[] args) {
        Solution2 s = new Solution2();
        System.out.println(Arrays.toString(s.twoSum(new int[]{-1,-1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, -2)));
        System.out.println(Arrays.toString(s.twoSum(new int[]{12,13,23,28,43,44,59,60,61,68,70,86,88,92,124,125,136,168,173,173,180,199,212,221,227,230,277,282,306,314,316,321,325,328,336,337,363,365,368,370,370,371,375,384,387,394,400,404,414,422,422,427,430,435,457,493,506,527,531,538,541,546,568,583,585,587,650,652,677,691,730,737,740,751,755,764,778,783,785,789,794,803,809,815,847,858,863,863,874,887,896,916,920,926,927,930,933,957,981,997}, 542)));
        System.out.println(Arrays.toString(s.twoSum(new int[]{3,24,50,79,88,150,345}, 200)));
        System.out.println(Arrays.toString(s.twoSum(new int[]{2,7,11,15}, 9)));
    }    
}

/*class Solution2 {
    public int[] twoSum(int[] numbers, int target) {
        int n = numbers.length;
        int leftIndex = 0;
        int rightIndex = n - 1;
        while (numbers[leftIndex] + numbers[rightIndex] != target) {
            int midIndex = (leftIndex + rightIndex) / 2;
            int midVal = numbers[midIndex];
            int leftValue = numbers[leftIndex];
            int rightValue = numbers[rightIndex];
            if ((midIndex != leftIndex && leftValue + midVal == target) || rightValue > target) {
                rightIndex = midIndex;
                continue;
            }
            if ((midIndex != rightIndex && rightValue + midVal == target) || leftValue < target) {
                leftIndex = midIndex;
                continue;
            }
        }
        return new int[] {leftIndex+1, rightIndex+1}; // +1 because it is 1-indexed
    }   
}
*/

// O(log n) binary search
class Solution2 {
    public int[] twoSum(int[] numbers, int target) {
        int n = numbers.length;
        int leftIndex = 0;
        int rightIndex = -1;

        int leftSearchIndex = 0;
        int rightSearchIndex = n - 1;
        while (rightIndex < 0 && leftSearchIndex != rightSearchIndex) {
            int midIndex = (leftSearchIndex + rightSearchIndex) / 2;
            int midVal = numbers[midIndex];
            if (midVal == target) {
                rightIndex = midIndex;
            } else if (midVal < target) {
                if (leftSearchIndex == midIndex) {
                    leftSearchIndex++;
                } else {
                    leftSearchIndex = midIndex;
                }
            } else {
                rightSearchIndex = midIndex;
            }
        }
        if (rightIndex < 0) { // rightSearchIndex == leftSearchIndex
            rightIndex = rightSearchIndex;
            if (rightSearchIndex == 0) {
                rightIndex++;
            }
        }

        int sum = numbers[leftIndex] + numbers[rightIndex];
        
        while (sum != target) {
            if (sum < target) {
                leftIndex++;
            } else { // sum > target
                rightIndex--;
            }
            sum = numbers[leftIndex] + numbers[rightIndex];
        }
        int[] result = new int[] {leftIndex+1, rightIndex+1};
        Arrays.sort(result);
        return result; // +1 because it is 1-indexed
    } 
}
