import java.util.Scanner;

public class day10 {
    public static void main(String[] args) {
        Sprite sprite = new Sprite();
        TubeCPU cpu = new TubeCPU(sprite);
        Display display = new Display(sprite);

        Clock clk = new Clock();
        clk.addListener(cpu);
        clk.addListener(display);

        Scanner sc = new Scanner(System.in).useDelimiter("\0");
        String program = sc.next();
        sc.close();

        cpu.load(program);

        for (int i = 0; i < 239; i++) {
            clk.cycle();
        }

        System.out.println(display);
        System.out.println("Signal stregth: " + cpu.getSignalStrength());
    }
}