import java.util.Random;
import java.util.concurrent.Semaphore;

public class App {
    private static final Random random = new Random();
    static Semaphore semaphore = new Semaphore(1);
    private static boolean semaphoreEnabled = true;

    static class CustomThread extends Thread {
        public void acquire() throws InterruptedException {
            if (semaphoreEnabled) {
                semaphore.acquire();
            }
        }

        public void release() throws InterruptedException {
            if (semaphoreEnabled) {
                semaphore.release();
            }
        }
    }

    static class BridgeConsumer extends CustomThread {
        private Bridge bridge;

        public BridgeConsumer(Bridge bridge) {
            this.bridge = bridge;
        }

        public void run() {
            try {
                while (true) {
                    this.acquire();
                    if (bridge.getCurrent() != null) {
                        Car c = bridge.getCurrent();
                        bridge.removeCurrent();
                        System.out.println("Consumed: " + c);
                    }
                    this.release();
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    static class CarProducer extends CustomThread {
        private Bridge bridge;
        private boolean direction;

        public CarProducer(Bridge bridge, boolean direction) {
            this.bridge = bridge;
            this.direction = direction;
        }

        public void run() {
            try {
                while (true) {
                    this.acquire();
                    if (bridge.getCurrent() == null) {
                        int id = random.nextInt(1000) + 1;
                        Car c = new Car(id, this.direction);
                        this.bridge.setCurrent(c);
                        System.out.println("Produced: " + c);
                    }
                    this.release();
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    /**
     * To change the semaphore on and off just change the variable semaphoreEnabled at the beginning of the class
     */
    public static void main(String[] args) {
        Bridge bridge = new Bridge();
        BridgeConsumer bc = new BridgeConsumer(bridge);
        CarProducer cp1 = new CarProducer(bridge, false);
        CarProducer cp2 = new CarProducer(bridge, true);
        bc.start();
        cp1.start();
        cp2.start();
    }
}
