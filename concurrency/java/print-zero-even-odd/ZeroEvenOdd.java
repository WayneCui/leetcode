import java.util.concurrent.CountDownLatch;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.IntConsumer;

class ZeroEvenOdd {
    private int n;
    private AtomicInteger count = new AtomicInteger(0);
    private CountDownLatch latch = new CountDownLatch(3);

    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        int x = 0;
        while((x = count.get()) < 2 * n) {
            if(x % 2 == 0) {
                printNumber.accept(0);
                count.getAndIncrement();
            }
        }
        latch.countDown();
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        int x;
        while((x = count.get()) < 2 * n) {
            if(x % 2 != 0 && (x + 1) % 2 == 0) {
                printNumber.accept((x + 1) / 2);
                count.getAndIncrement();
            }
        }
        latch.countDown();
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        int x;
        while((x = count.get()) < 2 * n) {
            if(x % 2 != 0 && (x + 1) % 2 != 0) {
                printNumber.accept((x + 1) / 2);
                count.getAndIncrement();
            }
        }
        latch.countDown();
    }

    public static void main(String[] args) throws InterruptedException{
        ZeroEvenOdd zeo = new ZeroEvenOdd(5);
        new Thread() {
            @Override
            public void run() {
                try {
                    zeo.zero(x -> System.out.print(x));
                } catch (InterruptedException e) {

                }
            }
        }.start();

        new Thread() {
            @Override
            public void run() {
                try {
                    zeo.odd(x -> System.out.print(x));
                } catch (InterruptedException e) {

                }
            }
        }.start();

        new Thread() {
            @Override
            public void run() {
                try {
                    zeo.even(x -> System.out.print(x));
                } catch (InterruptedException e) {

                }
            }
        }.start();
        zeo.latch.await();
        System.out.println();

    }
}
