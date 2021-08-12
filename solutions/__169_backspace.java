import java.io.BufferedInputStream;
import java.util.Scanner;

public class __169_backspace {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        for (int n = sc.nextInt(); n > 0; n--) {
            String s, t;
            s = sc.next();
            t = sc.next();
            if (s.length() < t.length()) {
                System.out.println("NO");
                continue;
            }
            int i = s.length()-1;
            int j = t.length()-1;
            while (i >= 0 && j >= 0) {
                if (s.charAt(i) == t.charAt(j)) {
                    i--;
                    j--;
                }
                else {
                    i -= 2;
                }
            }
            if (j == -1) {
                System.out.println("YES");
            }
            else {
                System.out.println("NO");
            }
        }
    }

}
