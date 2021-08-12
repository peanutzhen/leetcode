public class __166_majority_number {
    public static void main(String[] args) {
        __166_majority_number solution = new __166_majority_number();
        System.out.println(solution.majorityElement(new int[]{3,2,3}));
    }

    // 摩尔投票法
    public int majorityElement(int[] nums) {
        int major = nums[0];
        int count = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != major) {
                count--;
            } else {
                count++;
            }
            if (count == 0) {
                count = 1;
                major = nums[i];
            }
        }
        return major;
    }
}

