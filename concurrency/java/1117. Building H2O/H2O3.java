import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CyclicBarrier;

class H2O {

    private Semaphore hsem, osem;
    private CyclicBarrier barrier;

    public H2O() {
        hsem = new Semaphore(2, true);
        osem = new Semaphore(1, true);
        barrier = new CyclicBarrier(3);
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        hsem.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        try {
            barrier.await();
        } catch (BrokenBarrierException e) {}
        hsem.release();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        osem.acquire();
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
        releaseOxygen.run();
        try {
            barrier.await();
        } catch (BrokenBarrierException e) {}
        osem.release();
    }
}