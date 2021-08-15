import java.util.HashMap;
import java.util.Scanner;

public class __173_mihoyo2022r1_B {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] l = new int[n/2];
        int[] r = new int[n/2];
        for (int i = 0; i < n; i++) {
            if (i%2 == 0) {
                l[i/2] = sc.nextInt();
            }
            else {
                r[i/2] = sc.nextInt();
            }
        }
        int max1 = 0, max2 = 0;
        int x, y;
        HashMap<Integer, Integer> lmap = new HashMap<>();
        for (int i = 0; i < l.length; i++) {
            int count = lmap.getOrDefault(l[i], 0) + 1;
            lmap.put(l[i], count);
            if (count > max1) {
                max1 = count;
                x = l[i];
            }
        }

        HashMap<Integer, Integer> rmap = new HashMap<>();
        for (int i = 0; i < r.length; i++) {
            int count = rmap.getOrDefault(r[i], 0) + 1;
            rmap.put(r[i], count);
            if (count > max2) {
                max2 = count;
                y = r[i];
            }
        }

        System.out.println(n-max1-max2);
    }
}
