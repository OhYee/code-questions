const (
    mapSize = 1000005
)

type MyHashMap struct {
    m []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    m := make([]int, mapSize)
    for i := 0; i < mapSize; i++ {
        m[i] = -1
    }
    return MyHashMap {
        m: m,
    }
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    this.m[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    return this.m[key]
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    this.m[key] = -1
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
