import java.util.LinkedList;
import java.util.List;

public class Clock {
    private int counter = 1;
    private List<ClockListener> listeners = new LinkedList<>();

    public void cycle() {
        for (ClockListener listener : listeners) {
            listener.update(counter);
        }
        counter++;
    }

    public void addListener(ClockListener listener) {
        listeners.add(listener);
    }
}
