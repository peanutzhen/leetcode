import java.util.*;


public class __190_netease2022D {

    public int n;
    public int m;
    public int[][] board;
    public int[][] dp;

    public static void main(String[] args) {
        __190_netease2022D solution = new __190_netease2022D();
        int[][] test = {
                {2, 1, 1, 1, 2},
                {1, 1, 0, 1, 0},
                {1, 1, 2, 1, 1},
                {0, 2, 0, 0, 1},
        };
        System.out.println(solution.minSailCost(test));
    }

    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     * <p>
     * 计算最小航行费用
     *
     * @param input int整型二维数组 二维网格
     * @return int整型
     */
    public int minSailCost(int[][] input) {
        // write code here
        n = input.length;
        if (n == 0)
            return -1;
        m = input[0].length;
        if (input[n - 1][m - 1] == 2)
            return -1;
        dp = new int[n][m];
        board = input;
        for (int i = 0; i < n; i++) {
            Arrays.fill(dp[i], Integer.MAX_VALUE);
        }

        dfs(0, 0, board[0][0] == 1 ? -1 : -2);
        return dp[n - 1][m - 1] == Integer.MAX_VALUE ? -1 : dp[n - 1][m - 1];
    }

    public void dfs(int i, int j, int cost) {
        if (i == n || j == m || (board[i][j] == 2 && (i != 0 || j != 0))) {
            return;
        }
        cost = cost + (board[i][j] == 1 ? 1 : 2);
        if (cost < dp[i][j])
            dp[i][j] = cost;
        dfs(i + 1, j, cost);
        dfs(i, j + 1, cost);
    }
}