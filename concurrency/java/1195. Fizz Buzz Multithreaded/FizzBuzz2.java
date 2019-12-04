import java.util.function.IntConsumer;

class FizzBuzz2 {
    private int n;
    private int counter;

    public FizzBuzz2(int n) {
        this.n = n;
        this.counter = 1;
    }

    // printFizz.run() outputs "fizz".
    public synchronized void fizz(Runnable printFizz) throws InterruptedException {
        while(counter <= n) {
            if(counter % 3 == 0 && counter % 5 != 0) {
                printFizz.run();
                counter += 1;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public synchronized void buzz(Runnable printBuzz) throws InterruptedException {
        while(counter <= n) {
            if(counter % 3 != 0 && counter % 5 == 0) {
                printBuzz.run();
                counter += 1;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public synchronized void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        while(counter <= n) {
            if(counter % 3 == 0 && counter % 5 == 0) {
                printFizzBuzz.run();
                counter += 1;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public synchronized void number(IntConsumer printNumber) throws InterruptedException {
        while(counter <= n) {
            if(counter % 3 != 0 && counter % 5 != 0) {
                printNumber.accept(counter);
                counter += 1;
                notifyAll();
            } else {
                wait();
            }
        }
    }
}