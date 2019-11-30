// credits to https://leetcode.com/problems/building-h2o/discuss/392237
class H2O {

    private Semaphore hsem, osem;
    private Phaser phaser;

    public H2O() {
        hsem = new Semaphore(2, true);
        osem = new Semaphore(1, true);
        phaser = new Phaser(3);
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        hsem.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        phaser.arriveAndAwaitAdvance();
        hsem.release();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        osem.acquire();
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
        releaseOxygen.run();
        phaser.arriveAndAwaitAdvance();
        osem.release();
    }
}