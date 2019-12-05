// credits to https://leetcode.com/problems/the-dining-philosophers/discuss/407358/Semaphore-%2B-mutex

import java.util.concurrent.locks.ReentrantLock;
import java.util.concurrent.Semaphore;

class DiningPhilosophers {
    ReentrantLock[] locks;
    Semaphore sem;

    public DiningPhilosophers() {
        this.locks = new ReentrantLock[5];
        for(int i = 0; i < locks.length; i++) {
            locks[i] = new ReentrantLock();
        }
        this.sem = new Semaphore(4);
    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
                           Runnable pickLeftFork,
                           Runnable pickRightFork,
                           Runnable eat,
                           Runnable putLeftFork,
                           Runnable putRightFork) throws InterruptedException {
        int leftFork = philosopher;
        int rightFork = (philosopher + 4) % 5;
        //
        sem.acquire();

        //
        locks[leftFork].lock();
        pickLeftFork.run();

        locks[rightFork].lock();
        pickRightFork.run();

        //
        eat.run();

        //
        putRightFork.run();
        locks[rightFork].unlock();
        
        putLeftFork.run();
        locks[leftFork].unlock();
        
        //
        sem.release();
    }
}