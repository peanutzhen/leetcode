import java.math.BigDecimal;
import java.util.Scanner;

public class __171_alibaba20220813_A {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] nums = new int[n];
        BigDecimal left_sum = new BigDecimal(0);
        BigDecimal right_sum = new BigDecimal(0);
        for (int i = 0; i < n; i++) {
            nums[i] = sc.nextInt();
            if (i == 0) {
                left_sum = left_sum.add(new BigDecimal(nums[0]));
            }
            else {
                right_sum = right_sum.add(new BigDecimal(nums[i]));
            }
        }
        BigDecimal max_value = left_sum.multiply(right_sum);
        int max_idx = 1;
        for (int i = 1; i < n-1; i++) {
            left_sum = left_sum.add(new BigDecimal(nums[i]));
            right_sum = right_sum.subtract(new BigDecimal(nums[i]));
            BigDecimal tmp = left_sum.multiply(right_sum);
            if (tmp.compareTo(max_value) > 0) {
                max_value = tmp;
                max_idx = i+1;
            }
        }
        System.out.println(max_idx);
    }
}
