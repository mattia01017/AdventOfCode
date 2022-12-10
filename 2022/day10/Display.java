public class Display implements ClockListener {

    private boolean[][] matrix = new boolean[6][40];
    private Sprite sprite = new Sprite();

    public Display(Sprite sprite) {
        this.sprite = sprite;
        matrix[0][0] = sprite.isInPosition(0);
    }
    
    private void drawPixel(int cycle) {
        int l = cycle / 40;
        int c = cycle % 40;
        matrix[l][c] = sprite.isInPosition(c);
    }

    @Override
    public void update(int counter) {
        drawPixel(counter);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                sb.append(matrix[i][j] ? "\u25A0 " : ". ");
            }
            sb.append("\n");
        }
        return sb.toString();
    }
}
