import java.util.concurrent.Semaphore;

class H2O {

    private Semaphore hsem, osem;

    public H2O() {
        hsem = new Semaphore(2, true);
        osem = new Semaphore(0, true);
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        hsem.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        osem.release();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        osem.acquire(2);
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
        releaseOxygen.run();
        hsem.release(2);
    }
}