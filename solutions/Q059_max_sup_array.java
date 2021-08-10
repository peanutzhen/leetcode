public class Q059_max_sup_array {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int res = solution.maxSubArray(new int[]{-2,1,-3,4,-1,2,1,-5,4});
        System.out.println(res);
    }
}

class Solution {
    public int maxSubArray(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        int maxSum = (-1<<31);
        int tmp = 0;
        for (int num : nums) {
            tmp += num;
            if (tmp > maxSum) {
                maxSum = tmp;
            }
            if (tmp < 0) {
                tmp = 0;
            }
        }
        return maxSum;
    }
}
