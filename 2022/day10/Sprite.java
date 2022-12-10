public class Sprite {
    private int middlePos = 1;

    public void incrementPosition(int val) {
        middlePos += val;
    }

    public boolean isInPosition(int val) {
        return middlePos >= val - 1 && middlePos <= val + 1;
    }
}
