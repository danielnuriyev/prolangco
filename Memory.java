import java.util.ArrayList;
import java.util.List;

public class Memory {

    public static void main(String[] args) {

        final int rounds = 100;
        final int size = 1000000;
        // int[] list = new int[size];
        Integer[] list = new Integer[size];

        long min = Long.MAX_VALUE;
        long max = 0L;

        List<Long> memory = new ArrayList<>(rounds);

        long start = System.currentTimeMillis();

        for (int round = 0; round < rounds; round++) {

            for (int counter = 0; counter < size; counter++) {

                list[counter] = counter;

            }

            long usedMemory = Runtime.getRuntime().totalMemory() - Runtime.getRuntime().freeMemory();

            memory.add(usedMemory);

            if (usedMemory < min) {
                min = usedMemory;
            }
            if (usedMemory > max) {
                max = usedMemory;
            }

            // System.out.println(round + ": " + usedMemory);

        }

        System.out.println();

        System.out.println(System.currentTimeMillis() - start);
        System.out.println(min);
        System.out.println(
                (long) memory.stream().mapToDouble(a -> a).average().getAsDouble()
        );
        System.out.println(max);

    }

}
