/*
 * @lc app=leetcode id=146 lang=java
 *
 * [146] LRU Cache
 */

// @lc code=start
import java.util.*;
class LRUCache {
    private final int MAX_CACHE_SIZE;
    private final float DEFAULT_LOAD_FACTOR = 0.75f;
    LinkedHashMap<Integer, Integer> map;

    public LRUCache(int capacity) {
        this.MAX_CACHE_SIZE = capacity;
        int cap = (int) Math.ceil(MAX_CACHE_SIZE / DEFAULT_LOAD_FACTOR) + 1;
        map = new LinkedHashMap<Integer, Integer>(cap, DEFAULT_LOAD_FACTOR, true){
            @Override
            protected boolean removeEldestEntry(Map.Entry<Integer, Integer> eldlest) {
                return map.size() > MAX_CACHE_SIZE;
            }
        };
    }
    
    public int get(int key) {
        return map.getOrDefault(key, -1);
    }
    
    public void put(int key, int value) {
        map.put(key, value);
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
// @lc code=end


class LRUCache {
    private Map<Integer, Integer> map;
    private LinkedList<Integer> list;
    private int capacity;
    
    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.list = new LinkedList<Integer>();
        this.map = new HashMap<Integer, Integer>((int)Math.ceil(capacity / 0.75 + 1));
    }
    
    public int get(int key) {
        if(list.contains(key)){
            update(key);
        }
        return map.getOrDefault(key, -1);
    }
    
    public void put(int key, int value) {
        map.put(key, value);
        if(list.contains(key)) {
            list.remove(list.indexOf(key));
        }
        cleanup();
        list.add(key);
    }

    private void cleanup() {
        if(list.size() > capacity - 1) {
            Integer oldKey = list.removeFirst();
            map.remove(oldKey);
        }
    }

    private void update(Integer key) {
        list.remove(key);
        list.add(key);
    }
}