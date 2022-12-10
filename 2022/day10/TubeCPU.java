import java.util.Map;
import java.util.Scanner;
import java.util.TreeMap;

public class TubeCPU implements ClockListener {
    private int rax = 1;
    private int signalStrength = 0;
    private int currWait;
    private Scanner instructions;
    private String currInstruction;
    private Map<String, Integer> toWait;
    private Sprite sprite;

    public TubeCPU(Sprite sprite) {
        toWait = new TreeMap<>();
        toWait.put("noop", 0);
        toWait.put("addx", 1);

        this.sprite = sprite;
    }

    public void load(String program) {
        instructions = new Scanner(program);
    }

    public void execute() {
        String[] toks = currInstruction.split(" ");
        if (toks[0].equals("addx")) {
            int val = Integer.parseInt(toks[1]);
            rax += val;
            sprite.incrementPosition(val);
        }
    }

    private void calcSignalStrength(int counter) {
        if ((counter-20) % 40 == 0) {
            signalStrength += counter * rax;
        }
    }

    public int getSignalStrength() {
        return signalStrength;
    }

    @Override
    public void update(int counter) {
        calcSignalStrength(counter);
        if (currInstruction == null) {
            if (instructions.hasNext()) {
                currInstruction = instructions.nextLine();
                currWait = toWait.get(currInstruction.split(" ")[0]);
            } else {
                return;
            }
        }

        if (currWait == 0) {
            execute();
            currInstruction = null;
        } else {
            currWait--;
        }
    }
}
