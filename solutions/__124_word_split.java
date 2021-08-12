import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

public class __124_word_split {
    public static void main(String[] args) {
        __124_word_split solution = new __124_word_split();
        System.out.println(solution.wordBreak("leetcode", Arrays.asList("leet", "code")));
    }
    public boolean wordBreak(String s, List<String> wordDict) {
        HashSet<String> set = new HashSet<>(wordDict);
        int n = s.length();
        boolean[] dp = new boolean[n+1];
        dp[0] = true;
        for (int i = 1; i <= n; i++) {
            for (int j = 0; j < i; j++) {
                dp[i] = dp[j] && set.contains(s.substring(j,i));
                if (dp[i]) {
                    break;
                }
            }
        }
        return dp[n];
    }
}

