public class Q082_search_word {
    public static void main(String[] args) {
        Q082_search_word main = new Q082_search_word();
        boolean result = main.exist(new char[][]{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "SEE");
        System.out.println(result);
    }

    private int idx;
    private char[][] b;
    private String s;
    private boolean[][] visited;
    private int n;
    private int m;

    public boolean exist(char[][] board, String word) {
        n = board.length;
        m = board[0].length;
        b = board;
        s = word;
        visited = new boolean[n][m];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (dfs(i, j)) {
                    return true;
                }
            }
        }
        return false;
    }

    private boolean dfs(int i, int j) {
        if (idx == s.length()) {
            return true;
        }
        if (i == -1 || j == -1 || i == n || j == m || visited[i][j]) {
            return false;
        }
        if (b[i][j] == s.charAt(idx)) {
            visited[i][j] = true;
            idx++;
            if (dfs(i+1, j) || dfs(i-1, j) || dfs(i, j-1) || dfs(i, j+1)) {
                return true;
            }
            visited[i][j] = false;
            idx--;
        }
        return false;
    }
}
