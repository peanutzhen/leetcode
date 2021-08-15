import java.util.Scanner;
import java.util.Stack;

public class __172_mihoyo2022r1_A {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        while (n > 0) {
            boolean f = true;
            String s = sc.next();
            Stack<Character> stack = new Stack<>();
            for (int i = 0; i < s.length(); i++) {
                if (s.charAt(i) == 'a') {
                    stack.push(s.charAt(i));
                }
                else {
                    if (stack.isEmpty()) {
                        f = false;
                        break;
                    }
                    stack.pop();
                }
            }
            if (!stack.isEmpty()) {
                f = false;
            }
            if (f) {
                System.out.println("YES");
            }
            else {
                System.out.println("NO");
            }
            n--;
        }
    }
}
