import java.util.Scanner;
import java.util.TreeSet;

public class Q165_prev_sum {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int[] nums = new int[n];
        for (int i = 0; i < n; i++) {
            nums[i] = scanner.nextInt();
        }
        int ans = 0;
        TreeSet<Integer> treeSet = new TreeSet<>();
        for (int i = 0; i < n; i++) {
            Integer prev = treeSet.floor(nums[i]);
            while (prev != null && prev == nums[i]) {
                prev = treeSet.floor(prev-1);
            }
            if (prev != null) {
                ans += prev * (i+1);
            }
            treeSet.add(nums[i]);
        }
        System.out.println(ans);
    }
}
