t (
    mapSize = 1000005
)

type MyHashSet struct {
    m []bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet {
        m: make([]bool, mapSize),
    }
}


func (this *MyHashSet) Add(key int)  {
    this.m[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    this.m[key] = false
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return this.m[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
