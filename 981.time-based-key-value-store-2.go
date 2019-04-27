/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */
type TimeMap struct {
	Store map[string]map[int]string
	Index map[string][]int
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
	m := TimeMap{}
	m.Store = make(map[string]map[int]string)
	m.Index = make(map[string][]int)
	return m
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	_, ok := this.Store[key]
	if !ok {
		this.Store[key] = make(map[int]string)
	}

	this.Store[key][timestamp] = value
	
	this.Index[key] = append(this.Index[key], timestamp)
}


func (this *TimeMap) Get(key string, timestamp int) string {
	v, ok := this.Store[key][timestamp]
	if ok {
		return v
	}

	idx := search(this.Index[key], timestamp)
	if idx < 0 {
		idx = -idx - 2
		if idx >= 0 {
			return this.Store[key][this.Index[key][idx]]
		}
	}

	return ""
}

func search(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid]{
			high = mid - 1
		} else {
			return mid
		}
	}

	return -(low + 1)
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

