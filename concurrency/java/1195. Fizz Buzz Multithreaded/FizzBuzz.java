import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.IntConsumer;

class FizzBuzz {
    private int n;
    private AtomicInteger counter;

    public FizzBuzz(int n) {
        this.n = n;
        this.counter = new AtomicInteger(1);
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        int c;
        while((c = counter.get()) <= n) {
            if(c % 3 == 0 && c % 5 != 0){
                printFizz.run();
                counter.compareAndSet(c, c + 1);
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        int c;
        while((c = counter.get()) <= n) {
            if(c % 3 != 0 && c % 5 != 0){
                printBuzz.run();
                counter.compareAndSet(c, c + 1);
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        int c;
        while((c = counter.get()) <= n) {
            if(c % 3 == 0 && c % 5 == 0){
                printFizzBuzz.run();
                counter.compareAndSet(c, c + 1);
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        int c;
        while((c = counter.get()) <= n) {
            if(c % 3 != 0 && c % 5 != 0){
                printNumber.accept(c);
                counter.compareAndSet(c, c + 1);
            }
        }
    }
}