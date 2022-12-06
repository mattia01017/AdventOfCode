import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.Set;
import java.util.TreeSet;

public class day6 {
    public static void main(String[] args) {
        Scanner s = new Scanner(System.in);
        String inpStr = s.nextLine();
        s.close();

        List<Character> chars = new ArrayList<>(inpStr.length());
        for (Character c : inpStr.toCharArray()) {
            chars.add(c);
        }

        System.out.println("start-of-packet: " + findStart(chars, 4));
        System.out.println("start-of-message: " + findStart(chars, 14));
    }

    private static int findStart(List<Character> chars, int range) {
        int i;
        for (i = 0; i < chars.size() - range; i++) {
            List<Character> subList = chars.subList(i, i+range);
            Set<Character> charSet = new TreeSet<>(subList);
            if (charSet.size() == range) {
                break;
            }
        }
        return i+range;
    }
}